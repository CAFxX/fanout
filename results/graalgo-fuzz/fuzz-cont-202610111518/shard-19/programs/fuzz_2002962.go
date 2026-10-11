package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 int64; F2 int8 }
func genFn0() int64 {
    var v6 uint64 = 2
    var v7 int8 = 1
    switch 32768 {
    case 4294967295:
        var v8 uint32 = 367
        _ = v8
    case 2147483647:
        if false {
            var v9 int8 = v7
            _ = v9
        }
    }
    v6 = v6
    _ = v6
    _ = v7
    return -848
}
func genFn1(p0 uint, p1 string, p2 int32) uint8 {
    if (9223372036854775807 <= genFn0()) {
        p1 = "a"
        p1 = ("" + p1)
        p0 = (^(p0))
    } else {
        p2 += p2
        v10 := GenS1{F0: (p2 + p2)}
        _ = v10
    }
    v11 := []int{678, 32768, int(p0), 2147483647}
    fmt.Println(p1, p2)
    for fi12 := 0; fi12 < 17; fi12++ {
        v13 := GenS1{F0: -2147483648}
        _ = fi12
        _ = v13
    }
    p2 -= (-(p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 38
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[int]bool)
    v3 := []uint32{42, 65535}
    var v4 int = 42
    v5 := []string{}
    switch "a" {
    case "\n\t":
        switch "a" {
        case "foo bar":
            switch (v4 | v4) {
            case 452:
                v1 -= (v1 % (((v1 * v1)) | 1))
            case 709:
                var v14 uint = 18446744073709551615
                _ = v14
            default:
                switch "foo bar" {
                case "a":
                    if len(v3) > 0 {
                        v3[0] = 0
                    }
                }
            }
        case "\n\t":
            for fi15 := 0; fi15 < 7; fi15++ {
                if false {
                    for fi16 := 0; fi16 < 12; fi16++ {
                        switch "hello" {
                        case "":
                            v4 -= 127
                        }
                        _ = fi16
                    }
                    v4 = fi15
                }
                v17 := GenS1{F0: 2147483647}
                v4 = v4
                _ = fi15
                _ = v17
            }
        }
    default:
        if false {
            v4 = v4
            var v18 float64 = (float64((-2.25 / 0.1)) * ((3.14159 + 1e10) + (3.14159 + 0.0)))
            _ = v18
        }
    }
    var v19 uint32 = 2147483647
    for fi20 := 0; fi20 < 8; fi20++ {
        var v21 float64 = 3.14159
        _ = fi20
        _ = v21
    }
    switch v4 {
    case -1:
        var v22 uint8 = genFn1(936, "", 2147483647)
        _ = v22
    case 32767:
        if true {
            v19 = (v19 + (v19 % ((v19) | 1)))
        }
    case 0:
        v2[(v4 & 127)] = true
    }
    switch (v4 | v4) {
    case 32768:
        _ = len(v2)
    case 4294967295:
        var v23 bool = true
        _ = v23
    default:
        v1 += (v1 / ((42) | 1))
    }
    if false {
        if ((-32768 != 0) || ((128 < 1) || false)) {
            _ = len(v5)
            v2[32768] = (!(true || (false && false)))
        }
    }
    switch (("hello" + "foo bar") + ("foo bar" + "hello")) {
    case "\n\t":
        if len(v3) > 0 {
            v3[0] = 2147483648
        }
    }
    if false {
        var v24 int64 = 1
        v24 -= -1
        _ = v24
    } else {
        v25 := GenS2{F0: v19}
        if (true || true) {
            fmt.Println(v19)
        }
        var v26 uint64 = (^(v1))
        _ = v25
        _ = v26
    }
    switch "foo bar" {
    case "":
        fmt.Println(v1)
    case "hello":
        v27 := GenS3{F0: (v19 << ((255) & 63)), F1: 624, F2: -55}
        _ = v27
    case "foo bar":
        if (0 > 0) {
            var v28 string = ""
            var v29 int64 = genFn0()
            _ = v28
            _ = v29
        } else {
            v4 = v4
            var v30 uint16 = 1
            _ = v30
        }
    default:
        _ = len(v5)
    }
    for fi31 := 0; fi31 < 8; fi31++ {
        if false {
            fmt.Println(v4, v4, v19)
            v1 += v1
            var v32 float64 = (1.5 * 3.14159)
            _ = v32
        }
        switch "\n\t" {
        case "":
            var v33 float64 = 1.5
            _ = v33
        case "foo bar":
            var v34 float64 = 1e10
            _ = v34
        }
        _ = fi31
    }
    v35 := v2[513]
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
