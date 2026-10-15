package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint16; F1 uint64; F2 int }
type GenS3 struct { F0 int8; F1 int8; F2 uint64 }
func genFn0(p0 int32, p1 int8) uint {
    if ((p1 << ((304) & 63)) < 127) {
        p0 = (p0 ^ p0)
        if (!true) {
            var v5 uint = 694
            for fi6 := 0; fi6 < 18; fi6++ {
                var v7 uint64 = 0
                v5 = v5
                v8 := GenS1{F0: -1, F1: 32767, F2: (v7 | 42)}
                _ = fi6
                _ = v7
                _ = v8
            }
            _ = v5
        } else {
            p1 += 42
            var v9 float64 = 0.1
            for fi10 := 0; fi10 < 6; fi10++ {
                v11 := []int{(fi10 & fi10), (fi10 >> ((18446744073709551615) & 63))}
                _ = fi10
                _ = v11
            }
            _ = v9
        }
        v12 := []int{-575, 128}
        _ = v12
    } else {
        v13 := []int{-332, -1, -9223372036854775808, 294}
        _ = v13
    }
    p0 = (p0 & (p0 - (-(p0))))
    _ = p0
    _ = p1
    return 42
}
func genFn1() int8 {
    v14 := GenS1{F0: -1, F1: 255, F2: 0}
    var v15 int = 2147483647
    _ = v14
    _ = v15
    return 1
}

func main() {
    v1 := []string{("hello" + ""), ("foo bar" + "hello")}
    var v2 uint32 = 260
    var v3 uint64 = 0
    var v4 bool = (!true)
    switch 2147483647 {
    case 4294967295:
        var v16 int = 9223372036854775807
        _ = v16
    case 128:
        var v17 uint16 = 65535
        _ = v17
    case 82:
        var v18 int64 = 384
        _ = v18
    }
    fmt.Println(v4)
    switch "foo bar" {
    case "":
        if (!(9223372036854775807 <= 42)) {
            fmt.Println(v3, v2, v4)
        }
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    fmt.Println(v2, v4)
    v19 := GenS2{F0: 32767, F1: (v3 % ((18446744073709551615) | 1)), F2: 2147483647}
    var v20 float64 = 1e10
    if ((v2 - (v2 | v2)) > v2) {
        for fi21 := 0; fi21 < 20; fi21++ {
            if len(v1) > 0 {
                v1[0] = ("foo bar" + "hello")
            }
            _ = fi21
        }
        var v22 int8 = 1
        for fi23 := 0; fi23 < 12; fi23++ {
            fmt.Println(v22, v4, v4)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi23
        }
        _ = v22
    }
    var v24 string = "foo bar"
    v1 = append(v1, v24)
    v2 = 763
    switch 127 {
    case -9223372036854775808:
        var v25 int64 = -103
        _ = v25
    default:
        fmt.Println(v4, v2, v3)
    }
    v26 := GenS1{F0: 42, F1: 32767, F2: v3}
    for fi27 := 0; fi27 < 6; fi27++ {
        var v28 int64 = 270
        if true {
            v4 = v4
            var v29 int8 = 100
            v28 = v28
            _ = v29
        }
        switch "" {
        case "\n\t":
            var v30 int = 258
            _ = v30
        case "a":
            var v31 uint32 = 2147483647
            _ = v31
        default:
            for fi32 := 0; fi32 < 7; fi32++ {
                var v33 uint16 = 777
                var v34 float32 = 0.1
                if (153 > 1000) {
                    fmt.Println(v20, fi27, v20)
                    fmt.Println(v33, v4)
                }
                _ = fi32
                _ = v33
                _ = v34
            }
        }
        _ = fi27
        _ = v28
    }
    var v35 float32 = (3.14159 * -2.25)
    var v36 int8 = -82
    for fi37 := 0; fi37 < 5; fi37++ {
        if v4 {
            var v38 uint32 = v2
            _ = v38
        }
        v3 = (v3 + v3)
        v39 := GenS3{F0: -92, F1: (v36 << ((4294967295) & 63)), F2: (v3 + v3)}
        _ = fi37
        _ = v39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v26
    _ = v3
    _ = v35
    _ = v36
    _ = v4
}
