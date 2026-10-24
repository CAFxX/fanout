package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 string }
func genFn0(p0 int8) uint8 {
    v4 := []int{-64, 128, -1, 0}
    p0 += (p0 | (p0 << ((145) & 63)))
    var v5 uint32 = 902
    v4 = append(v4, 32767)
    _ = p0
    _ = v4
    _ = v5
    return 2
}

func main() {
    v1 := []string{"", ("\n\t" + ""), "hello"}
    v2 := []string{"foo bar", "", "", "hello"}
    var v3 uint32 = 4294967295
    v3 = 927
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 uint64 = 869
        var v8 int64 = 9223372036854775807
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 int64 = 1
    v10 := GenS1{F0: 100, F1: 0, F2: "hello"}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v11 int16 = -1
    var v12 uint = 65535
    fmt.Println(v3, v11)
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v13 float32 = (0.1 + (((0.1 + 3.14159) - 0.0) + ((3.14159 * ((0.0 * 1e10) * 0.0)) + (1e10 - 0.1))))
        v14 := GenS1{F0: -97, F1: uint16(v3), F2: ("\n\t" + "a")}
        _ = v13
        _ = v14
    } else {
        v2 = append(v2, "")
        if ((-890 == 42) && false) {
            var v15 float64 = (0.0 / (-2.25 - 1.5))
            v2 = append(v2, (("a" + ("hello" + "hello")) + "hello"))
            _ = v15
        }
        v16 := GenS1{F0: -128, F1: 36, F2: "\n\t"}
        _ = v16
    }
    if true {
        _ = len(v2)
    } else {
        v9 = (^(v9))
        _ = len(v2)
    }
    v9 += v9
    for fi17 := 0; fi17 < 3; fi17++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v9 = (v9 & v9)
        _ = fi17
    }
    v9 += -281
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v9
}
