package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS3 struct { F0 int64; F1 int32 }
func genFn0(p0 uint8, p1 int8) int16 {
    if true {
        var v4 int = 4294967295
        switch (v4 >> ((334) & 63)) {
        case 2147483647:
            v5 := GenS1{F0: 42, F1: 4294967295}
            _ = v5
        case 42:
            var v6 string = "\n\t"
            _ = v6
        default:
            for fi7 := 0; fi7 < 12; fi7++ {
                fmt.Println(p0, p0, v4)
                v8 := []int{128}
                var v9 string = ("\n\t" + ("\n\t" + "foo bar"))
                _ = fi7
                _ = v8
                _ = v9
            }
        }
        switch ("a" + "foo bar") {
        case "":
            for fi10 := 0; fi10 < 6; fi10++ {
                fmt.Println(p1, p0, p1)
                _ = fi10
            }
        default:
            if true {
                v4 -= 32768
                var v11 uint16 = 2
                var v12 uint16 = (v11 ^ uint16(p0))
                _ = v11
                _ = v12
            } else {
                var v13 int16 = 1
                fmt.Println(p1, v13)
                _ = v13
            }
        }
        _ = v4
    } else {
        v14 := GenS1{F0: -536, F1: int(p0)}
        _ = v14
    }
    fmt.Println(p0, p0)
    switch "foo bar" {
    case "\n\t":
        v15 := []int{-1}
        _ = v15
    default:
        for fi16 := 0; fi16 < 3; fi16++ {
            var v17 uint16 = 255
            v18 := GenS1{F0: 0, F1: -551}
            for fi19 := 0; fi19 < 12; fi19++ {
                switch fi19 {
                case 75:
                    v20 := GenS1{F0: -694, F1: (fi16 | fi19)}
                    _ = v20
                case 324:
                    p0 -= p0
                case 4294967295:
                    fmt.Println(p1)
                default:
                    var v21 uint16 = 255
                    _ = v21
                }
                p0 = (p0 * 181)
                _ = fi19
            }
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    var v22 float32 = (1.5 * 0.1)
    _ = p0
    _ = p1
    _ = v22
    return 42
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 255
    var v3 uint8 = (^(v2))
    var v23 uint8 = (v3 * v2)
    if (1 >= 0) {
        v1["\n\t"] = false
        fmt.Println(v23)
        fmt.Println(v3, v23)
    }
    v24 := GenS2{F0: true, F1: 143, F2: 1}
    for fi25 := 0; fi25 < 15; fi25++ {
        switch fi25 {
        case -607:
            if false {
                var v26 int = 0
                v3 -= v3
                _ = v26
            }
        default:
            v2 -= (v3 >> ((42) & 63))
        }
        _ = fi25
    }
    v27 := GenS1{F0: -1, F1: 0}
    if (!(false && false)) {
        v3 = v2
    }
    if ((!("hello" < "foo bar")) && true) {
        fmt.Println(v3, v3, v23)
        v28 := GenS1{F0: -1, F1: 4294967295}
        _ = v28
    }
    var v29 int = 32768
    v30 := GenS3{F0: 546, F1: 822}
    var v31 float32 = float32(0.0)
    v2 = v2
    if (false || (-839 <= 74)) {
        v31 -= (1.5 + float32(float64(1e10)))
        if (226 >= 272) {
            _ = len(v1)
            v2 += (v23 / ((v2) | 1))
            var v32 uint8 = v3
            _ = v32
        } else {
            v33 := GenS1{F0: 42, F1: v29}
            var v34 uint8 = v2
            v23 -= v34
            _ = v33
            _ = v34
        }
    }
    for fi35 := 0; fi35 < 10; fi35++ {
        v2 += (v3 & (^(v2)))
        for fi36 := 0; fi36 < 12; fi36++ {
            fmt.Println(v29)
            var v37 uint32 = 2147483647
            _ = fi36
            _ = v37
        }
        if true {
            v38 := GenS3{F0: 892, F1: -2147483648}
            var v39 int32 = 486
            switch "\n\t" {
            case "foo bar":
                if (611 < 779) {
                    var v40 int16 = 1
                    _ = v40
                }
            default:
                var v41 bool = (32767 == 255)
                _ = v41
            }
            _ = v38
            _ = v39
        }
        _ = fi35
    }
    for fi42 := 0; fi42 < 12; fi42++ {
        var v43 uint8 = v23
        if ("foo bar" != ("hello" + ("\n\t" + ("" + (("a" + "a") + "a"))))) {
            var v44 int64 = -286
            v45 := GenS2{F0: (42 > -128), F1: 1, F2: 2}
            _ = v44
            _ = v45
        }
        _ = fi42
        _ = v43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
