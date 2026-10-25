package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 string }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 int; F1 int32; F2 string }

func main() {
    var v1 int32 = 2147483647
    v2 := []uint32{2147483647, 2, 0, 2}
    var v3 uint8 = 0
    v4 := []string{}
    var v5 uint64 = 108
    switch (("a" + "") + "a") {
    case "\n\t":
        v6 := GenS1{F0: 127, F1: int16(v1), F2: "a"}
        _ = v6
    case "":
        v7 := GenS2{F0: 1, F1: v5}
        _ = v7
    default:
        for fi8 := 0; fi8 < 9; fi8++ {
            v9 := GenS3{F0: int(v5), F1: 1, F2: ""}
            var v10 float64 = (3.14159 / 1.0)
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    v11 := GenS2{F0: (v5 % (((v5 | v5)) | 1)), F1: (v5 << ((826) & 63))}
    v5 += v5
    v5 = (^(v5))
    var v12 bool = ((v5 >> ((0) & 63)) > v5)
    var v13 float64 = 0.0
    v5 = 1
    _ = len(v4)
    if ("foo bar" == "") {
        switch ("hello" + "foo bar") {
        case "":
            v3 = v3
        case "hello":
            if true {
                v5 = v5
            } else {
                for fi14 := 0; fi14 < 13; fi14++ {
                    if len(v2) > 0 {
                        v2[0] = 617
                    }
                    _ = fi14
                }
                v15 := GenS1{F0: -76, F1: -842, F2: "a"}
                _ = v15
            }
        default:
            v12 = false
        }
    }
    fmt.Println(v5, v3)
    if false {
        var v16 uint32 = 10
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if v12 {
            fmt.Println(v1, v13, v5)
        }
        _ = v16
    }
    for fi17 := 0; fi17 < 9; fi17++ {
        for fi18 := 0; fi18 < 4; fi18++ {
            fmt.Println(v5, v5)
            v5 += 893
            _ = fi18
        }
        var v19 int8 = 104
        _ = fi17
        _ = v19
    }
    var v20 bool = v12
    v1 = v1
    var v21 bool = false
    var v22 float64 = float64(1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
