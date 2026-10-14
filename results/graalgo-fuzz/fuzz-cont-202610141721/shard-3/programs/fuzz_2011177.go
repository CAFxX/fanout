package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 int64; F1 int64; F2 int32 }
func genFn0(p0 uint64, p1 uint16) uint {
    for fi4 := 0; fi4 < 8; fi4++ {
        var v5 int32 = -2147483648
        p0 -= 851
        var v6 uint = 1
        _ = fi4
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        fmt.Println(p0)
        v8 := []int{fi7, fi7, (fi7 / ((-475) | 1))}
        var v9 float32 = 1e10
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 2147483648}
    var v11 bool = (-2.25 != (-2.25 * float64(-2.25)))
    for fi12 := 0; fi12 < 17; fi12++ {
        v13 := GenS2{F0: (fi12 - fi12), F1: 2}
        if (0.1 != (3.14159 * 0.0)) {
            p1 -= p1
            if v11 {
                var v14 bool = false
                _ = v14
            } else {
                p1 = p1
                switch (fi12 % (((fi12 | fi12)) | 1)) {
                case 387:
                    p1 = 2
                default:
                    fmt.Println(p0)
                }
            }
        }
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 42
}

func main() {
    var v1 uint32 = 165
    v2 := make(map[int]int)
    var v3 uint16 = 623
    var v15 float64 = 0.1
    for fi16 := 0; fi16 < 5; fi16++ {
        v3 -= v3
        _ = fi16
    }
    v17 := []int{4294967295}
    for fi18 := 0; fi18 < 11; fi18++ {
        var v19 int32 = 42
        _ = fi18
        _ = v19
    }
    var v20 float32 = (1.5 / -2.25)
    switch -93 {
    case -473:
        for fi21 := 0; fi21 < 10; fi21++ {
            var v22 uint8 = 255
            for fi23 := 0; fi23 < 11; fi23++ {
                v24 := GenS1{F0: (v1 >> ((4294967295) & 63))}
                _ = fi23
                _ = v24
            }
            fmt.Println(v1, v15, v3)
            _ = fi21
            _ = v22
        }
    case 660:
        v25 := GenS3{F0: 1, F1: 58, F2: -2147483648}
        _ = v25
    case 32768:
        _ = len(v17)
    default:
        for fi26 := 0; fi26 < 14; fi26++ {
            var v27 int = (fi26 >> ((255) & 63))
            fmt.Println(v15, fi26, v20)
            _ = fi26
            _ = v27
        }
    }
    if (("" + "") < "foo bar") {
        if true {
            if (32767 > 1000) {
                if false {
                    delete(v2, -1)
                }
                var v28 string = (("" + "hello") + "foo bar")
                var v29 int8 = 1
                _ = v28
                _ = v29
            }
        }
        for fi30 := 0; fi30 < 17; fi30++ {
            if (127 < 116) {
                v1 += v1
                var v31 int = -218
                var v32 string = "hello"
                _ = v31
                _ = v32
            }
            _ = fi30
        }
        v17 = append(v17, 2147483647)
    }
    v33 := GenS1{F0: 0}
    if ("a" < "hello") {
        switch 929 {
        case 2147483647:
            for fi34 := 0; fi34 < 13; fi34++ {
                if false {
                    fmt.Println(v15)
                    if true {
                        v1 -= v1
                        var v35 bool = false
                        _ = v35
                    } else {
                        var v36 uint8 = 1
                        var v37 int16 = -448
                        _ = v36
                        _ = v37
                    }
                }
                _ = fi34
            }
        case 9223372036854775807:
            v15 += float64((-2.25 + v20))
        default:
            v3 = (^(v3))
        }
    } else {
        var v38 int = 0
        _ = v38
    }
    switch "foo bar" {
    case "foo bar":
        v15 = 1.5
    }
    v2[285] = -245
    for fi39 := 0; fi39 < 16; fi39++ {
        var v40 int16 = 32767
        v2[(fi39 & fi39)] = fi39
        _ = fi39
        _ = v40
    }
    v20 -= (v20 * v20)
    var v41 float32 = ((float32(1.5) * v20) * 0.0)
    var v42 uint8 = 1
    var v43 int16 = 32767
    for fi44 := 0; fi44 < 5; fi44++ {
        fmt.Println(v1, v15)
        for fi45 := 0; fi45 < 11; fi45++ {
            v46 := GenS1{F0: uint32(v3)}
            _ = fi45
            _ = v46
        }
        _ = fi44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v41
    _ = v42
    _ = v43
}
