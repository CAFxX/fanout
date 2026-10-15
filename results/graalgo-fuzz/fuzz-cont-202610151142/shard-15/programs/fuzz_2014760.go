package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 int32, p1 uint64, p2 bool) bool {
    switch "foo bar" {
    case "\n\t":
        fmt.Println(p1, p2)
    case "a":
        p0 += p0
    case "foo bar":
        if (2 >= 128) {
            v4 := GenS1{F0: -1}
            v5 := []int{308, 32768, 2147483647}
            _ = v4
            _ = v5
        } else {
            v6 := []int{0, -9223372036854775808, 355}
            var v7 int64 = -298
            _ = v6
            _ = v7
        }
    default:
        var v8 float64 = (1.5 / 3.14159)
        _ = v8
    }
    p2 = (42 != 9223372036854775807)
    p1 -= 641
    _ = p0
    _ = p1
    _ = p2
    return (0 != 647)
}
func genFn1(p0 int64, p1 int32, p2 string) int8 {
    fmt.Println(p1, p1, p0)
    switch "a" {
    case "":
        for fi9 := 0; fi9 < 19; fi9++ {
            for fi10 := 0; fi10 < 13; fi10++ {
                p1 = 515
                switch fi9 {
                case 562:
                    p1 += (p1 * p1)
                case -9223372036854775808:
                    v11 := []int{(fi10 | (fi9 % (((fi9 >> ((1) & 63))) | 1))), 4294967295, (fi10 / (((fi10 - -1)) | 1))}
                    _ = v11
                case 767:
                    fmt.Println(p0, p1)
                }
                var v12 uint = 130
                _ = fi10
                _ = v12
            }
            _ = fi9
        }
    case "foo bar":
        switch 0 {
        case -1:
            fmt.Println(p1)
        case 971:
            if (-106 < 4294967295) {
                var v13 float32 = -2.25
                _ = v13
            }
        }
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        for fi15 := 0; fi15 < 12; fi15++ {
            if (42 != 9223372036854775807) {
                var v16 int = 32768
                if (!true) {
                    var v17 uint16 = 32767
                    if false {
                        fmt.Println(p1)
                        var v18 int8 = 127
                        var v19 string = p2
                        _ = v18
                        _ = v19
                    } else {
                        v20 := GenS1{F0: -113}
                        _ = v20
                    }
                    switch (^(fi15)) {
                    case 2147483647:
                        if (!(!(p2 >= ""))) {
                            v21 := GenS1{F0: -1}
                            var v22 uint16 = (v17 ^ (^(v17)))
                            _ = v21
                            _ = v22
                        }
                    case 0:
                        v23 := []int{(v16 - int(p1)), v16}
                        _ = v23
                    }
                    _ = v17
                }
                _ = v16
            } else {
                p1 = 379
            }
            p1 -= p1
            _ = fi15
        }
        if (!true) {
            var v24 uint = 432
            v25 := GenS1{F0: 74}
            _ = v24
            _ = v25
        }
        var v26 uint32 = 720
        _ = fi14
        _ = v26
    }
    p2 = "a"
    var v27 uint16 = 553
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    return 100
}

func main() {
    v1 := []string{}
    var v2 int32 = 0
    var v3 uint32 = 245
    fmt.Println(v3, v3)
    if false {
        fmt.Println(v2, v2)
        var v28 int = 2147483647
        _ = v28
    }
    if (42 == int16(v3)) {
        _ = len(v1)
    }
    v1 = append(v1, "foo bar")
    v29 := GenS1{F0: 42}
    var v30 float32 = 0.0
    v30 = 1.5
    for fi31 := 0; fi31 < 4; fi31++ {
        for fi32 := 0; fi32 < 10; fi32++ {
            fmt.Println(fi31, v30, v3)
            v2 = (v2 / ((v2) | 1))
            if false {
                switch -9223372036854775808 {
                case 32767:
                    var v33 int32 = (v2 % (((v2 / (((v2 - v2)) | 1))) | 1))
                    _ = v33
                }
                var v34 int = (fi31 + fi31)
                _ = v34
            }
            v2 += (v2 - (v2 << ((750) & 63)))
            v35 := GenS1{F0: 0}
            _ = fi32
            _ = v35
        }
        switch "" {
        case "hello":
            fmt.Println(v2, fi31, v30)
        case "\n\t":
            switch (fi31 & -355) {
            case 42:
                v2 -= v2
            default:
                switch (fi31 ^ fi31) {
                case -1:
                    v30 = 1e10
                }
            }
        default:
            v30 += v30
        }
        _ = fi31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
}
