package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int32; F1 uint32; F2 string }
type GenS3 struct { F0 bool; F1 int64; F2 uint }
func genFn0(p0 int32, p1 uint32) uint16 {
    for fi5 := 0; fi5 < 8; fi5++ {
        var v6 int32 = 347
        _ = fi5
        _ = v6
    }
    v7 := []int{387, 4294967295}
    _ = p0
    _ = p1
    _ = v7
    return 2
}

func main() {
    var v1 uint64 = 42
    v2 := []string{"", "a", "\n\t"}
    v3 := make(map[string]string)
    var v4 int8 = 42
    v8 := GenS1{F0: "a", F1: 1000}
    v4 -= 100
    var v9 int32 = 2147483647
    for fi10 := 0; fi10 < 13; fi10++ {
        switch "foo bar" {
        case "a":
            for fi11 := 0; fi11 < 9; fi11++ {
                v4 = 42
                _ = fi11
            }
        case "hello":
            fmt.Println(v4, v4, v9)
        case "":
            v12 := GenS2{F0: v9, F1: 65535, F2: ("\n\t" + ("foo bar" + ("a" + "a")))}
            _ = v12
        }
        delete(v3, "")
        _ = fi10
    }
    var v13 bool = false
    v4 -= -128
    for fi14 := 0; fi14 < 2; fi14++ {
        switch "hello" {
        case "":
            var v15 int64 = -965
            _ = v15
        default:
            for fi16 := 0; fi16 < 17; fi16++ {
                var v17 uint32 = 2147483647
                if v13 {
                    var v18 int64 = 0
                    _ = v18
                }
                _ = fi16
                _ = v17
            }
        }
        if (true && (293 > 899)) {
            var v19 bool = (!false)
            _ = v19
        }
        if false {
            fmt.Println(v9)
        } else {
            fmt.Println(v1, v1)
            v20 := GenS2{F0: (v9 & v9), F1: 2147483648, F2: "foo bar"}
            _ = v20
        }
        _ = fi14
    }
    v21 := GenS3{F0: (!v13), F1: -471, F2: uint(v1)}
    if (375 == 18446744073709551615) {
        for fi22 := 0; fi22 < 9; fi22++ {
            v1 = v1
            _ = fi22
        }
    } else {
        if (42 > 1) {
            _ = len(v2)
            v23 := v3["\n\t"]
            v9 = (v9 % ((0) | 1))
            _ = v23
        } else {
            v9 = (v9 * (v9 | v9))
        }
    }
    v24 := GenS1{F0: (("a" + ("hello" + "hello")) + "\n\t"), F1: -311}
    v25 := GenS2{F0: 2147483647, F1: 2147483648, F2: "foo bar"}
    var v26 uint32 = 4294967295
    var v27 int8 = (v4 ^ 0)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
