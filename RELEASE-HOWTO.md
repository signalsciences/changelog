# Release HOWTO

since I forget.

Most if not all of this should be semi-automated.

1. Commit any changes and make sure Travis-ci passes

2. Review existing tags and pick new release number


    TODO: from changelog get last version and make sure it's not
    tagged already.

    ```bash
    git tag
    ```

3. Add to CHANGELOG.md, commit, push, make sure travis-ci passes

4. Tag locally 

    ```bash
    git tag -a vNEWVERSION -m "tag version XYZ"
    ```

5. Push

    ```bash
    git push origin vNEWVERSION
    ```

