package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 int32; F2 uint16 }
type GenS4 struct { F0 int16; F1 int64 }
func genFn0() int16 {
    var v7 uint32 = 230
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := []int{128, (fi8 - 128), fi8}
        _ = fi8
        _ = v9
    }
    var v10 int = -163
    _ = v10
    _ = v7
    return -32768
}

func main() {
    var v1 bool = true
    v2 := make(map[string]string)
    var v3 int32 = 42
    var v4 uint8 = 127
    var v5 int8 = 1
    var v6 int8 = v5
    v11 := []int{-2147483648, 872}
    if v1 {
        var v12 float64 = 0.0
        if (!(42 >= uint16(v4))) {
            v11 = append(v11, 32768)
        }
        _ = v12
    } else {
        for fi13 := 0; fi13 < 11; fi13++ {
            fmt.Println(fi13, v5)
            v5 = v5
            _ = fi13
        }
        v4 += (v4 - (v4 - (v4 ^ (v4 & v4))))
        for fi14 := 0; fi14 < 16; fi14++ {
            if false {
                if (v1 || (259 <= 1000)) {
                    v15 := GenS1{F0: v1}
                    fmt.Println(v1, v4)
                    _ = v15
                } else {
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                }
            }
            _ = fi14
        }
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        var v17 int16 = -1
        v18 := GenS1{F0: v1}
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := GenS2{F0: (v1 || v1)}
    v20 := GenS3{F0: 79, F1: (v3 & 42), F2: 553}
    _ = len(v11)
    _ = len(v2)
    if true {
        v21 := GenS1{F0: (!v1)}
        v22 := GenS3{F0: -1, F1: -874, F2: 922}
        _ = v21
        _ = v22
    }
    v6 = v6
    switch ("hello" + "\n\t") {
    case "a":
        var v23 uint16 = 872
        _ = v23
    case "\n\t":
        for fi24 := 0; fi24 < 6; fi24++ {
            switch (-(fi24)) {
            case 128:
                v4 -= 1
            case -9223372036854775808:
                var v25 uint32 = 4294967295
                _ = v25
            case 42:
                v6 = v5
            default:
                v26 := GenS4{F0: -1, F1: -1}
                _ = v26
            }
            delete(v2, ("\n\t" + "a"))
            _ = fi24
        }
    case "foo bar":
        v11 = append(v11, -4)
    default:
        v1 = false
    }
    var v27 int64 = -572
    var v28 uint8 = v4
    fmt.Println(v6, v1)
    v3 += -1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
