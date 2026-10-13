package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint32; F1 uint64 }

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{284}
    v3 := make(map[string]uint32)
    var v4 int = 9223372036854775807
    if false {
        switch -984 {
        case 4294967295:
            v5 := GenS1{F0: 65535, F1: 32767}
            _ = v5
        case 42:
            v4 = v4
        }
        v4 += (v4 >> ((297) & 63))
        var v6 int32 = 505
        _ = v6
    }
    v7 := GenS2{F0: 833, F1: 2}
    var v8 uint32 = 2147483647
    var v9 float32 = float32((0.0 + 0.1))
    fmt.Println(v4, v4)
    v8 = v8
    if ("a" == "\n\t") {
        switch ("hello" + "hello") {
        case "\n\t":
            v10 := GenS1{F0: 42, F1: 32767}
            _ = v10
        default:
            var v11 bool = (v9 >= v9)
            _ = v11
        }
        v8 += (v8 % ((v8) | 1))
        var v12 int = (v4 << ((18446744073709551615) & 63))
        _ = v12
    } else {
        v13 := GenS1{F0: 0, F1: -197}
        _ = v13
    }
    if true {
        for fi14 := 0; fi14 < 10; fi14++ {
            if len(v2) > 0 {
                v2[0] = 2
            }
            _ = fi14
        }
        var v15 string = (("\n\t" + ("" + "hello")) + ("a" + "\n\t"))
        for fi16 := 0; fi16 < 6; fi16++ {
            v4 -= (fi16 + v4)
            delete(v1, v15)
            var v17 uint16 = 2
            _ = fi16
            _ = v17
        }
        _ = v15
    } else {
        fmt.Println(v8, v8, v8)
        if true {
            var v18 int8 = -128
            v2 = append(v2, (v8 * v8))
            _ = v18
        } else {
            switch "hello" {
            case "\n\t":
                v4 += v4
            }
            switch v4 {
            case 32767:
                v3["hello"] = v8
            default:
                if true {
                    var v19 bool = false
                    _ = v19
                } else {
                    v9 = -2.25
                }
            }
            v8 -= (v8 & (v8 * (v8 | (v8 / ((4294967295) | 1)))))
        }
    }
    var v20 bool = false
    v21 := GenS2{F0: 42, F1: 1}
    var v22 bool = v20
    var v23 uint16 = 42
    if v22 {
        var v24 int8 = -21
        _ = v24
    } else {
        v25 := GenS2{F0: v8, F1: uint64(v23)}
        v26 := GenS1{F0: 794, F1: 935}
        v23 += (v23 / ((v23) | 1))
        _ = v25
        _ = v26
    }
    switch ("\n\t" + "hello") {
    case "a":
        if (v22 && ((v8 > (v8 ^ 4294967295)) && (2 < 2))) {
            v9 -= ((float32(0.0) - float32(0.1)) * 1e10)
            v2 = append(v2, v8)
        }
    case "":
        v23 += v23
    default:
        if (((v23 / ((v23) | 1)) <= v23) || v20) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            v27 := GenS1{F0: 776, F1: 42}
            _ = v27
        }
    }
    switch "\n\t" {
    case "a":
        var v28 uint16 = v23
        _ = v28
    case "foo bar":
        v29 := GenS1{F0: 18446744073709551615, F1: -32768}
        _ = v29
    default:
        v3[(("\n\t" + "a") + ("\n\t" + "foo bar"))] = (^(v8))
    }
    delete(v1, "\n\t")
    switch (v4 & 9223372036854775807) {
    case 9223372036854775807:
        var v30 int64 = -487
        _ = v30
    default:
        for fi31 := 0; fi31 < 6; fi31++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi31
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
