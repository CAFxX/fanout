package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }

func main() {
    var v1 string = "a"
    var v2 int16 = 1000
    var v3 float64 = 0.1
    var v4 uint16 = 65535
    var v5 int = 2147483647
    v6 := []int{v5}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v7 float64 = (0.0 + v3)
    var v8 int8 = 88
    if len(v6) > 0 {
        v6[0] = (v5 | 127)
    }
    var v9 int = v5
    v7 -= v7
    v3 = (v3 + -2.25)
    if true {
        if false {
            var v10 float64 = v3
            _ = v10
        } else {
            v2 = v2
            v3 = ((float64(1e10) / ((v3 * v7) + -2.25)) * 1.5)
            v11 := GenS1{F0: false, F1: -533}
            _ = v11
        }
        if ((("\n\t" + (v1 + v1)) + (v1 + (v1 + ""))) < ((v1 + v1) + v1)) {
            var v12 uint32 = 1
            _ = v12
        } else {
            var v13 int = 128
            _ = v13
        }
        if (false || (v9 != (v5 / ((v9) | 1)))) {
            fmt.Println(v3)
            v3 -= 1e10
            v9 = v5
        }
    }
    v14 := GenS1{F0: true, F1: v2}
    for fi15 := 0; fi15 < 7; fi15++ {
        var v16 int16 = -1
        var v17 string = v1
        _ = fi15
        _ = v16
        _ = v17
    }
    if (32768 <= (v9 % (((v9 ^ v9)) | 1))) {
        var v18 int = -1
        _ = v18
    }
    v8 -= int8(v9)
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
