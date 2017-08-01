# Release HOWTO

since I forget.

Most if not all of this should be semi-automated.

1. Review existing tags and pick new release number


    TODO: from changelog get last version and make sure it's not
    tagged already.

    ```bash
    git tag
    ```

2. Add to CHANGELOG.md

2. Tag locally 

    ```bash
    git tag -a vNEWVERSION -m "tag version XYZ"
    ```

3. Push

    ```bash
    git push origin NEWVERSION
    ```

