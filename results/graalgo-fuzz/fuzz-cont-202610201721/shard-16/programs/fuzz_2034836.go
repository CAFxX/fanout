package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }

func main() {
    var v1 uint = 340
    var v2 uint16 = 0
    var v3 uint32 = 317
    var v4 bool = (-638 > 42)
    v5 := make(map[string]uint32)
    v6 := make(map[int]bool)
    if ((v1 % ((42) | 1)) >= 65535) {
        if ((!(!(v4 || v4))) || v4) {
            var v7 int16 = 32767
            _ = v7
        }
    } else {
        var v8 uint64 = 0
        var v9 uint32 = (v3 % (((v3 * 56)) | 1))
        switch "foo bar" {
        case "\n\t":
            var v10 bool = ((2 < v2) && false)
            _ = v10
        case "hello":
            if (v4 && v4) {
                var v11 int32 = 1
                v6[-9223372036854775808] = (!((v2 == 65535) && (((("\n\t" + "hello") + "foo bar") + "\n\t") == "")))
                _ = v11
            }
        case "foo bar":
            v12 := GenS1{F0: 42, F1: v4}
            _ = v12
        default:
            v13 := GenS1{F0: 68, F1: true}
            _ = v13
        }
        _ = v8
        _ = v9
    }
    v2 += v2
    var v14 int32 = 825
    v14 += (v14 % ((v14) | 1))
    var v15 uint64 = 42
    if (!false) {
        if (v4 || (18446744073709551615 > 0)) {
            if (v4 && (("" + "hello") >= ("hello" + "hello"))) {
                v16 := GenS1{F0: 127, F1: (!v4)}
                _ = v16
            }
            v17 := v5["a"]
            fmt.Println(v2, v15, v15)
            _ = v17
        }
    } else {
        var v18 string = (("foo bar" + "a") + "\n\t")
        _ = v18
    }
    v15 = v15
    if (v14 == 2147483647) {
        v3 = (v3 | (v3 | v3))
        for fi19 := 0; fi19 < 8; fi19++ {
            fmt.Println(v15, v3, v14)
            _ = fi19
        }
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        v14 += int32(v1)
        _ = fi20
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        v1 = (^(v1))
        v1 = 4294967295
        _ = fi21
    }
    v22 := GenS1{F0: -6, F1: (("hello" + "") < ("a" + ""))}
    var v23 uint32 = v3
    v24 := []int{-24, 775, 32767, 127}
    var v25 bool = (!(-32768 != 42))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
