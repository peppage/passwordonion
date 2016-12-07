/*
Package passwordonion provides a way to store passwords in the same was as dropbox.

This is all based on this blog post https://blogs.dropbox.com/tech/2016/09/how-dropbox-securely-stores-your-passwords

TLDR; SHA512 to make up for some bcrypt shortcomings before bcrypt. Then AES256 the
bcrypted password so if the db leaks in some cases they might not get the pepper which
is not stored in the db.

To start you'll want to generate a "pepper" or key that is used for AES but not stored in the
database. The key needs to be 32 characters long.

Then call Encrypt when you need to store a user's password and store the result in your database.

When you need to check their password use Compare and send the pepper along with the password
from the db and the string the user entered. The func will return nil if the user can continue.

    // Encrypt password, store p.
    p, err := passwordonion.Encrypt(pepper, password)

    // Check a user entered password. Verify err is nil.
    err := passwordonion.Compare(pepper, string(p), password)

*/
package passwordonion
