let sequenceType = 0x30;
let integerType = 0x02;
let byteArrayType = 0x04;
let objectType = 0x06;

export function encodeSequence(elements: Uint8Array[]) {
    let length = 0;
    for (let i = 0; i < elements.length; i++) {
        length += elements[i].length;
    }
    const result = new Uint8Array(2 + length);
    result[0] = 0x30;
    result[1] = length;
    let offset = 2;
    for (let i = 0; i < elements.length; i++) {
        for (let j = 0; j < elements[i].length; j++) {
            result[offset++] = elements[i][j];
        }
    }
    return result;
}

export function encodeElement(element: any): Uint8Array {
    if (element instanceof Uint8Array) {
        return encodeByteArray(element);
    } else if (typeof element === "number") {
        return encodeInteger(element);
    } else {
        throw new Error("Unknown element type");
    }
}

export function encodeByteArray(element: Uint8Array): Uint8Array {
    const result = new Uint8Array(2 + element.length);
    result[0] = byteArrayType;
    result[1] = element.length;
    for (let i = 0; i < element.length; i++) {
        result[2 + i] = element[i];
    }
    return result;
}

export function encodeInteger(element: number): Uint8Array {
    const result = new Uint8Array(2);
    result[0] = integerType;
    result[1] = 1;
    result[2] = element;
    return result;
}