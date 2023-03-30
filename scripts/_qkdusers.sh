#!/bin/bash

set -Eeuo pipefail

term() {
  export TERM=/usr/bin/gnome-terminal
  if [ -f $TERM ]; then
    $TERM -- bash -c "$@"
  else
    export TERM=/usr/bin/osascript
    if [ -f $TERM ]; then
      export CMD="tell app \"Terminal\" to do script \"$@\""
      $TERM -e "$CMD"
    else
      export TERM=/cygdrive/c/Windows/System32/cmd.exe
      if [ -f $TERM ]; then
       $TERM /c start bash -c "$@"
      fi
    fi
  fi
}

export MY_DIR=`dirname $0`
pushd $MY_DIR/..

export PROJ_ROOT=$PWD
#export LD_LIBRARY_PATH=$PROJ_ROOT/.libs:/opt/oqs/lib
#export DYLD_LIBRARY_PATH=$PROJ_ROOT/.libs:/opt/oqs/lib
touch src/test/java/lv/lumii/test/QkdTestUser1.java
touch src/test/java/lv/lumii/test/QkdTestUser2.java
./gradlew compileTestJava

export JAVA_CP="$PROJ_ROOT/build/classes/java/main:$PROJ_ROOT/build/classes/java/test:$PROJ_ROOT/.jars/\\\*"
export JAVA_LP="/opt/oqs/lib:/usr/local/lib:$PROJ_ROOT/.libs"

export MAIN_CLASS=lv.lumii.test.QkdTestUser1
#term "DYLD_LIBRARYPATH=$JAVA_LP java -cp $JAVA_CP -Djava.library.path=$JAVA_LP $MAIN_CLASS"

export MAIN_CLASS=lv.lumii.test.QkdTestUser2
term "java -cp $JAVA_CP -Djava.library.path=$JAVA_LP $MAIN_CLASS"

popd

