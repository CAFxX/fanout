package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 int16 }
type GenS5 struct { F0 int32; F1 uint8; F2 uint64 }
func genFn0(p0 uint64) uint32 {
    if false {
        for fi4 := 0; fi4 < 7; fi4++ {
            v5 := GenS1{F0: 187}
            _ = fi4
            _ = v5
        }
        p0 += (p0 & (p0 * (p0 << ((0) & 63))))
    }
    var v6 int16 = 32767
    switch -9223372036854775808 {
    case 42:
        var v7 string = ("hello" + "a")
        _ = v7
    case 32767:
        fmt.Println(p0, v6)
    default:
        switch 705 {
        case -306:
            v8 := GenS1{F0: 1}
            _ = v8
        case 128:
            v6 -= v6
        }
    }
    p0 = p0
    v9 := []int{-9223372036854775808, 127, -2147483648, 4294967295}
    _ = p0
    _ = v6
    _ = v9
    return 42
}
func genFn1(p0 uint32, p1 int) string {
    v10 := []int{p1, p1, p1}
    p1 += (^(p1))
    if (419 == 65535) {
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 uint64 = 488
            _ = fi11
            _ = v12
        }
        p0 -= (p0 ^ p0)
    }
    if len(v10) > 0 {
        v10[0] = p1
    }
    _ = p0
    _ = p1
    _ = v10
    return ("hello" + "\n\t")
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 65535
    v3 := []string{(((("\n\t" + "") + "a") + "a") + "\n\t"), "a", "", "hello"}
    var v13 int8 = 0
    v14 := GenS1{F0: 32767}
    for fi15 := 0; fi15 < 16; fi15++ {
        if (true || (!false)) {
            v13 = (v13 | -128)
            fmt.Println(v13)
            for fi16 := 0; fi16 < 20; fi16++ {
                _ = len(v1)
                v13 = (-(v13))
                _ = fi16
            }
        }
        if false {
            v17 := GenS2{F0: v2, F1: fi15}
            for fi18 := 0; fi18 < 14; fi18++ {
                var v19 uint = 18446744073709551615
                v20 := GenS3{F0: (("\n\t" + "a") + "\n\t")}
                fmt.Println(fi18)
                _ = fi18
                _ = v19
                _ = v20
            }
            v2 += (^(v2))
            _ = v17
        }
        fmt.Println(v13, v2)
        _ = fi15
    }
    var v21 int8 = v13
    if true {
        v13 += v13
    }
    v13 = (^(v21))
    v22 := GenS4{F0: 1, F1: 876}
    for fi23 := 0; fi23 < 18; fi23++ {
        for fi24 := 0; fi24 < 17; fi24++ {
            switch fi23 {
            case 128:
                var v25 bool = (0 > 130)
                _ = v25
            case 2147483647:
                if (("foo bar" + "a") > "foo bar") {
                    var v26 string = "foo bar"
                    v27 := GenS3{F0: (v26 + "")}
                    _ = v26
                    _ = v27
                }
            case 32767:
                delete(v1, "\n\t")
                v21 = (v13 + (v21 * (v13 << ((239) & 63))))
                var v28 int8 = (^(v13))
                _ = v28
            }
            _ = fi24
        }
        for fi29 := 0; fi29 < 12; fi29++ {
            v2 += genFn0(781)
            if (930 > 4294967295) {
                v2 -= genFn0(2)
            }
            _ = fi29
        }
        _ = fi23
    }
    v30 := GenS5{F0: -671, F1: 127, F2: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v30
}
