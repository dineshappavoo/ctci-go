GO PACKAGE:
1.create a folder in the name same as the package name u wanted
2.Add the code to the folder
3.Inside the code the package name should be the folder name
    for example 'binary tree' package
        /src
            trees/
                binarytree/
                    binarytree.go
    Inside binarytree.go first line should be 'package binarytree'
4. Set the go path and export for the src directory of the go project
    export GOPATH=$Home/Documents/FALL-2013-COURSES/Imp_Data_structures/workspace/ctci-go/src
5. go to the binarytree folder cd $Home/Documents/FALL-2013-COURSES/Imp_Data_structures/workspace/ctci-go/src/trees/binarytree
6. go build
7. go install - this will add a library under package folder
    	/Users/Dany/Documents/FALL-2013-COURSES/Imp_Data_structures/workspace/ctci-go/pkg/darwin_amd64/trees/binarytree.a

