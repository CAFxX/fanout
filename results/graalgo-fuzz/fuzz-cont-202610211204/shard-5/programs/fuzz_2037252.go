package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 bool; F1 int8; F2 int }
type GenS3 struct { F0 string; F1 int16 }

func main() {
    v1 := []int{-184, -2147483648}
    var v2 uint8 = 2
    var v3 bool = true
    var v4 bool = ("hello" > (("foo bar" + "\n\t") + "foo bar"))
    var v5 uint32 = 4294967295
    if len(v1) > 0 {
        v1[0] = -70
    }
    v6 := GenS1{F0: ("hello" + "foo bar"), F1: 1}
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            v9 := GenS2{F0: false, F1: -1, F2: (fi8 + fi8)}
            _ = fi8
            _ = v9
        }
        var v10 int16 = -319
        _ = fi7
        _ = v10
    }
    if ((32767 <= 902) || (v5 != (v5 * v5))) {
        fmt.Println(v5, v4, v2)
        v11 := GenS3{F0: ("a" + ("a" + "")), F1: 1}
        v5 += v5
        _ = v11
    }
    v4 = (!(!(51 != -63)))
    var v12 int = -902
    var v13 int32 = 42
    var v14 float32 = 1.5
    fmt.Println(v2)
    var v15 int16 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
