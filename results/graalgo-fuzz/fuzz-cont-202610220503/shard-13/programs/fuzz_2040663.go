package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 bool }
type GenS2 struct { F0 int; F1 int32; F2 int32 }
type GenS3 struct { F0 string; F1 string }
func genFn0() uint64 {
    v6 := []int{}
    var v7 uint32 = 65535
    if true {
        v8 := GenS1{F0: 277, F1: 4294967295, F2: false}
        _ = len(v6)
        _ = v8
    }
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 string = ""
    var v2 float64 = 3.14159
    var v3 bool = true
    var v4 string = v1
    var v5 bool = v3
    for fi9 := 0; fi9 < 19; fi9++ {
        if true {
            for fi10 := 0; fi10 < 13; fi10++ {
                v3 = (!v5)
                var v11 int32 = -2147483648
                _ = fi10
                _ = v11
            }
        }
        v12 := []int{(fi9 + fi9)}
        _ = fi9
        _ = v12
    }
    var v13 uint64 = 1
    var v14 uint64 = v13
    v15 := []int{}
    v15 = append(v15, -605)
    v16 := GenS2{F0: 32768, F1: -2147483648, F2: -778}
    var v17 int8 = 42
    fmt.Println(v4)
    v13 -= v14
    var v18 uint32 = 2
    if (!(v1 >= ("\n\t" + v4))) {
        switch -1 {
        case 1:
            v17 = (v17 >> ((65535) & 63))
        case 128:
            for fi19 := 0; fi19 < 20; fi19++ {
                var v20 uint = uint(v18)
                _ = fi19
                _ = v20
            }
        case 889:
            v21 := GenS3{F0: "\n\t", F1: (((((v4 + (("\n\t" + v4) + v4)) + "a") + ("a" + ((v4 + v1) + ("a" + v4)))) + (((v1 + "") + v1) + "foo bar")) + ("foo bar" + v1))}
            _ = v21
        default:
            v15 = append(v15, 2147483647)
        }
        v17 = 0
        fmt.Println(v4)
    }
    var v22 bool = (((("\n\t" + "a") + ("" + v4)) >= ((v4 + v4) + ("hello" + ""))) && v3)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
