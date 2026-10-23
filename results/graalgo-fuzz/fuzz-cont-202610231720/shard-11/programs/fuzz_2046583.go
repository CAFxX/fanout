package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
func genFn0() int {
    var v4 uint32 = 0
    var v5 uint64 = 0
    var v6 int = -354
    fmt.Println(v4)
    _ = v4
    _ = v5
    _ = v6
    return (v6 * (-(v6)))
}

func main() {
    var v1 int64 = 42
    v2 := make(map[string]string)
    var v3 bool = false
    var v7 string = "a"
    fmt.Println(v7, v1)
    v8 := GenS1{F0: 0, F1: -9223372036854775808}
    v2[v7] = "a"
    v9 := []int{}
    v1 = -1
    v1 = 9223372036854775807
    if ("foo bar" < v7) {
        var v10 uint8 = 255
        _ = v10
    } else {
        var v11 bool = (65535 < 308)
        switch 1 {
        case 32768:
            v11 = v3
        case -2147483648:
            v11 = (v11 && ((false && ((2 < 1) || (false || (((v7 + ("hello" + (v7 + ("\n\t" + (v7 + v7))))) + v7) == v7)))) && true))
        default:
            v12 := GenS1{F0: 0, F1: (v1 % (((v1 ^ -1)) | 1))}
            _ = v12
        }
        _ = v11
    }
    v1 -= 200
    var v13 bool = v3
    var v14 int8 = 42
    delete(v2, v7)
    var v15 int = -2147483648
    for fi16 := 0; fi16 < 10; fi16++ {
        var v17 uint = 255
        _ = fi16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
