package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 int; F2 int8 }
func genFn0(p0 uint64, p1 int) uint64 {
    if false {
        if (!((2147483647 > 2147483648) && ((false && false) || true))) {
            var v4 int16 = 560
            _ = v4
        }
        var v5 uint32 = 4294967295
        var v6 int16 = 1
        _ = v5
        _ = v6
    } else {
        var v7 uint8 = 127
        _ = v7
    }
    var v8 int16 = 86
    for fi9 := 0; fi9 < 7; fi9++ {
        p0 -= p0
        if (-542 > -1) {
            switch ("foo bar" + "foo bar") {
            case "":
                var v10 int64 = 1
                _ = v10
            case "hello":
                fmt.Println(p0, v8)
            case "\n\t":
                v8 = v8
            }
        } else {
            var v11 string = "a"
            _ = v11
        }
        v12 := []int{(^(p1)), 891, (p1 << ((255) & 63))}
        _ = fi9
        _ = v12
    }
    if (0.0 <= ((3.14159 / float64((1.5 - (float32(1e10) * 1.5)))) + (0.0 / 1.0))) {
        if ("a" <= ("foo bar" + "a")) {
            fmt.Println(p0)
            v13 := GenS1{F0: 100, F1: v8, F2: 128}
            _ = v13
        }
    }
    if (-9223372036854775808 > 1) {
        if true {
            v14 := GenS1{F0: -25, F1: (v8 & v8), F2: 208}
            fmt.Println(v8, p1)
            var v15 int64 = -9223372036854775808
            _ = v14
            _ = v15
        } else {
            var v16 int64 = 0
            var v17 uint8 = 255
            _ = v16
            _ = v17
        }
        if (1.5 == 3.14159) {
            if false {
                var v18 int64 = -1
                fmt.Println(p0, p0, p0)
                for fi19 := 0; fi19 < 11; fi19++ {
                    for fi20 := 0; fi20 < 8; fi20++ {
                        v18 = (v18 >> ((255) & 63))
                        fmt.Println(p0)
                        _ = fi20
                    }
                    var v21 int8 = 4
                    _ = fi19
                    _ = v21
                }
                _ = v18
            }
        } else {
            if ((1.5 + 0.1) > (3.14159 + float64(0.1))) {
                var v22 float32 = 0.0
                for fi23 := 0; fi23 < 16; fi23++ {
                    if true {
                        p1 -= 733
                    } else {
                        fmt.Println(p0, p0, fi23)
                        var v24 float32 = 3.14159
                        var v25 int64 = -796
                        _ = v24
                        _ = v25
                    }
                    _ = fi23
                }
                v26 := GenS2{F0: 0}
                _ = v22
                _ = v26
            }
            switch "\n\t" {
            case "hello":
                if false {
                    var v27 float64 = 0.1
                    v28 := GenS1{F0: 0, F1: v8, F2: 1}
                    _ = v27
                    _ = v28
                }
            }
            fmt.Println(p0, p1)
        }
    }
    _ = p0
    _ = p1
    _ = v8
    return p0
}

func main() {
    var v1 float64 = 1e10
    var v2 uint = 208
    v3 := []string{("hello" + "a"), "a"}
    var v29 int8 = 1
    var v30 float64 = (v1 * ((1.5 + (((-2.25 + ((v1 / float64(0.1)) / v1)) - (0.1 - v1)) + 1.5)) - -2.25))
    if len(v3) > 0 {
        v3[0] = ""
    }
    var v31 uint64 = 704
    var v32 uint16 = uint16(v31)
    v33 := GenS3{F0: -2147483648, F1: 1, F2: (v29 << (((v2 << ((65535) & 63))) & 63))}
    fmt.Println(v1, v31)
    v32 += v32
    var v34 float64 = 0.1
    if (250 >= 0) {
        fmt.Println(v2, v1)
    } else {
        switch "" {
        case "a":
            var v35 uint16 = v32
            _ = v35
        case "\n\t":
            if len(v3) > 0 {
                v3[0] = "\n\t"
            }
        }
        switch "a" {
        case "\n\t":
            v29 -= 120
        default:
            var v36 float32 = (0.1 + 1.5)
            _ = v36
        }
        for fi37 := 0; fi37 < 2; fi37++ {
            for fi38 := 0; fi38 < 8; fi38++ {
                var v39 uint32 = 65535
                _ = fi38
                _ = v39
            }
            _ = fi37
        }
    }
    v2 = uint(v32)
    var v40 uint = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v40
}
