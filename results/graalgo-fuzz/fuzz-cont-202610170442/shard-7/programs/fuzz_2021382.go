package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 string }
type GenS2 struct { F0 int32; F1 int32; F2 int16 }
func genFn0(p0 int) uint32 {
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 uint32 = 2147483647
        v8 := []int{fi6, -782, 32767}
        var v9 uint64 = 2
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        p0 -= (p0 * 0)
        p0 = 4294967295
        _ = fi10
    }
    var v11 float32 = 1.5
    var v12 int64 = 9223372036854775807
    _ = p0
    _ = v11
    _ = v12
    return 9
}
func genFn1() int64 {
    for fi13 := 0; fi13 < 9; fi13++ {
        var v14 uint16 = 42
        _ = fi13
        _ = v14
    }
    var v15 int32 = 42
    var v16 int32 = v15
    fmt.Println(v15)
    _ = v15
    _ = v16
    return -9223372036854775808
}

func main() {
    var v1 int8 = 0
    v2 := []string{("hello" + "a"), "hello"}
    var v3 int8 = (v1 - (v1 * v1))
    var v4 uint32 = 2147483648
    var v5 uint32 = (^(v4))
    v4 -= 75
    v5 = (v5 / ((genFn0(1)) | 1))
    var v17 uint = 637
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v18 int32 = 246
    var v19 int64 = -9223372036854775808
    var v20 bool = (127 != 228)
    if (uint16(v5) != 832) {
        if len(v2) > 0 {
            v2[0] = "hello"
        }
        for fi21 := 0; fi21 < 3; fi21++ {
            v2 = append(v2, "\n\t")
            _ = fi21
        }
    } else {
        var v22 float64 = 0.0
        _ = v22
    }
    fmt.Println(v4)
    for fi23 := 0; fi23 < 14; fi23++ {
        for fi24 := 0; fi24 < 18; fi24++ {
            v25 := GenS1{F0: 18446744073709551615, F1: (v18 >> ((736) & 63)), F2: ""}
            var v26 uint32 = (v4 ^ v4)
            _ = fi24
            _ = v25
            _ = v26
        }
        var v27 float64 = 1e10
        switch "hello" {
        case "":
            for fi28 := 0; fi28 < 13; fi28++ {
                if (!false) {
                    var v29 int8 = v3
                    _ = v29
                }
                var v30 int16 = -136
                switch ((("\n\t" + "a") + "\n\t") + "hello") {
                case "\n\t":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case "foo bar":
                    var v31 int8 = v3
                    _ = v31
                }
                _ = fi28
                _ = v30
            }
        case "a":
            fmt.Println(fi23)
        default:
            _ = len(v2)
        }
        _ = fi23
        _ = v27
    }
    if len(v2) > 0 {
        v2[0] = ("" + "\n\t")
    }
    v32 := GenS1{F0: 42, F1: int32(v1), F2: "a"}
    for fi33 := 0; fi33 < 1; fi33++ {
        if ((!(fi33 == (fi33 << ((v17) & 63)))) && (1e10 == (1e10 - 3.14159))) {
            v34 := GenS2{F0: (v18 << ((v17) & 63)), F1: 1, F2: 1}
            v4 = (v5 ^ (v5 << (((v17 + v17)) & 63)))
            _ = v34
        }
        _ = fi33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
