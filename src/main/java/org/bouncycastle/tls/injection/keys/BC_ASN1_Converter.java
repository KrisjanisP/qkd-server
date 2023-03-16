package org.bouncycastle.tls.injection;

import org.bouncycastle.asn1.ASN1Set;
import org.bouncycastle.asn1.pkcs.PrivateKeyInfo;
import org.bouncycastle.asn1.x509.SubjectPublicKeyInfo;
import org.bouncycastle.crypto.params.AsymmetricKeyParameter;

import java.io.IOException;

/**
 * The interface for converting public and private keys
 * between
 * <ul>
 * <li>
 *   the lightweight (internal) BC representation
 *   ("key parameter" in BC terms, represented by the AsymmetricKeyParameter class)
 * </li>
 * and
 * <li>
 *   the ASN.1 notation
 *   (represented by the PrivateKeyInfo and SubjectPublicKeyInfo BC classes).
 * </li>
 * </ul>
 * <p>
 * #pqc-tls #injection
 *
 * @author Sergejs Kozlovics
 */
public interface BC_ASN1_Converter {
    /**
     * Checks whether the given BC key (public or private) can be converted to ASN.1.
     *
     * @param bcKey an internal BC representation of a public or a private key
     *              that has to be converted to ASN.1
     * @return returns true, iff bcKey is of known type and can be converted to ASN.1
     * (i.e., a PrivateKeyInfo or SubjectPublicKeyInfo instance)
     */
    boolean isSupportedParameter(AsymmetricKeyParameter bcKey);

    /**
     * Converts the given private key from ASN.1 to the internal BC representation.
     *
     * @param asnPrivateKey private key in the ASN.1 notation
     * @return internal BC representation of the private key
     * @throws IOException
     */
    AsymmetricKeyParameter createPrivateKeyParameter(PrivateKeyInfo asnPrivateKey) throws IOException;


    /**
     * Converts the given private key from the internal BC representation to the ASN.1 notation.
     *
     * @param bcPrivateKey internal BC representation of a private key
     * @param attributes   ASN.1 attributes to be embedded into the ASN.1 representation
     * @return ASN.1 representation of the private key
     * @throws IOException
     */
    PrivateKeyInfo createPrivateKeyInfo(AsymmetricKeyParameter bcPrivateKey, ASN1Set attributes) throws IOException;


    /**
     * Converts the given public key from ASN.1 to the internal BC representation.
     *
     * @param ansPublicKey  public key in the ASN.1 notation
     * @param defaultParams some default parameters (currently, null is passed)
     * @return internal BC representation of the public key
     * @throws IOException
     */
    AsymmetricKeyParameter createPublicKeyParameter(SubjectPublicKeyInfo ansPublicKey, Object defaultParams) throws IOException;

    /**
     * Converts the given public key from the internal BC representation to the ASN.1 notation.
     *
     * @param bcPublicKey internal BC representation of a public key
     * @return ASN.1 representation of the public key
     * @throws IOException
     */
    SubjectPublicKeyInfo createSubjectPublicKeyInfo(AsymmetricKeyParameter bcPublicKey) throws IOException;
}
