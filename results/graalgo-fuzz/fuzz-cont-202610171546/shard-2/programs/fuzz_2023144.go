package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 float64 = -2.25
    var v2 int = 2147483647
    v3 := []uint32{65535, 2, 65535}
    v4 := []string{("\n\t" + "\n\t"), "\n\t", (("a" + "a") + "hello"), ("hello" + "foo bar")}
    v5 := []int{}
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 int = fi6
        var v8 uint8 = 2
        fmt.Println(fi6, v2, v7)
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint16 = 255
    fmt.Println(v9, v9, v1)
    var v10 int32 = 1
    var v11 int16 = 1000
    v12 := GenS1{F0: (v2 & (^(v2))), F1: -1}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v13 string = "a"
    switch "foo bar" {
    case "":
        if (((v10 % (((v10 & v10)) | 1)) >= v10) || false) {
            fmt.Println(v2, v13)
        } else {
            v11 = (v11 ^ v11)
        }
    case "foo bar":
        v10 = (v10 & v10)
    case "hello":
        v11 = (v11 / (((v11 + v11)) | 1))
    default:
        v14 := GenS1{F0: v2, F1: 0}
        _ = v14
    }
    var v15 int16 = (v11 << ((1) & 63))
    for fi16 := 0; fi16 < 17; fi16++ {
        v17 := GenS2{F0: 127}
        if true {
            if (v1 > v1) {
                var v18 uint16 = v9
                _ = v18
            }
            v1 += v1
        }
        _ = fi16
        _ = v17
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v19 := GenS1{F0: (-(v2)), F1: 0}
    v20 := GenS2{F0: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
