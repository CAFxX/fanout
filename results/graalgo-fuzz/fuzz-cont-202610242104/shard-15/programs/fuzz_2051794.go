package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 int32; F1 int8; F2 int16 }
type GenS5 struct { F0 uint32; F1 bool; F2 int32 }
func genFn0(p0 bool, p1 int64, p2 string) uint32 {
    if (((2147483647 > -9223372036854775808) && ((!(32767 == 42)) || ("foo bar" >= p2))) || true) {
        v5 := GenS1{F0: (p0 && (!p0)), F1: 201}
        v6 := GenS1{F0: (86 <= -81), F1: 255}
        _ = v5
        _ = v6
    } else {
        switch "foo bar" {
        case "foo bar":
            var v7 uint16 = 32767
            _ = v7
        case "a":
            switch -543 {
            case 4294967295:
                var v8 float32 = 0.0
                _ = v8
            case 9223372036854775807:
                var v9 int = 398
                _ = v9
            case 127:
                switch (p2 + p2) {
                case "a":
                    fmt.Println(p0, p0, p2)
                case "hello":
                    var v10 float64 = 3.14159
                    _ = v10
                case "":
                    p2 = p2
                default:
                    fmt.Println(p2, p1)
                }
            default:
                for fi11 := 0; fi11 < 9; fi11++ {
                    var v12 int64 = (p1 + p1)
                    p0 = p0
                    _ = fi11
                    _ = v12
                }
            }
        default:
            var v13 int64 = (p1 + p1)
            _ = v13
        }
        v14 := GenS1{F0: true, F1: 2}
        _ = v14
    }
    p0 = p0
    v15 := []int{609, 0, -196, 32767}
    v16 := GenS1{F0: (586 == 0), F1: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return 2147483648
}
func genFn1(p0 uint16, p1 string, p2 uint64) int {
    p1 = (p1 + p1)
    v17 := []int{127}
    for fi18 := 0; fi18 < 7; fi18++ {
        p1 = (p1 + "hello")
        var v19 int32 = 179
        var v20 string = p1
        _ = fi18
        _ = v19
        _ = v20
    }
    v21 := GenS2{F0: false}
    _ = len(v17)
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v21
    return 4294967295
}

func main() {
    v1 := make(map[string]bool)
    var v2 string = "hello"
    var v3 bool = false
    var v4 int = 2147483647
    if v3 {
        var v22 uint = 255
        _ = v22
    }
    var v23 int64 = -767
    v2 = v2
    v24 := GenS2{F0: false}
    v25 := GenS3{F0: 307, F1: 2147483647}
    v26 := []int{}
    switch (-(v4)) {
    case 1:
        v2 = v2
    case 128:
        v23 = (-(v23))
    case -1:
        v27 := GenS4{F0: -2147483648, F1: 0, F2: -531}
        _ = v27
    default:
        switch v2 {
        case "a":
            var v28 string = v2
            _ = v28
        case "\n\t":
            switch v2 {
            case "a":
                v3 = ((v4 | 42) <= v4)
            }
        case "":
            v29 := GenS2{F0: v3}
            _ = v29
        default:
            switch "a" {
            case "a":
                var v30 uint = 1
                _ = v30
            case "\n\t":
                v3 = true
            default:
                var v31 uint64 = 895
                _ = v31
            }
        }
    }
    if (126 < -1) {
        if (!v3) {
            fmt.Println(v4, v2)
            v26 = append(v26, v4)
            v32 := GenS5{F0: 463, F1: true, F2: int32(v4)}
            _ = v32
        }
    } else {
        v33 := GenS4{F0: -2147483648, F1: 35, F2: 138}
        switch v4 {
        case 946:
            if len(v26) > 0 {
                v26[0] = (^(v4))
            }
        case 2147483647:
            var v34 int8 = -5
            _ = v34
        }
        _ = v33
    }
    for fi35 := 0; fi35 < 7; fi35++ {
        for fi36 := 0; fi36 < 20; fi36++ {
            v37 := GenS1{F0: v3, F1: 205}
            v38 := v1[""]
            _ = fi36
            _ = v37
            _ = v38
        }
        v39 := GenS4{F0: -2147483648, F1: 0, F2: 1000}
        _ = fi35
        _ = v39
    }
    switch v4 {
    case -1:
        if (-128 != 100) {
            for fi40 := 0; fi40 < 17; fi40++ {
                var v41 uint8 = 0
                _ = fi40
                _ = v41
            }
            v42 := GenS3{F0: int32(v23), F1: 65535}
            v43 := v1[v2]
            _ = v42
            _ = v43
        } else {
            v44 := GenS2{F0: v3}
            v2 = "hello"
            _ = v44
        }
    default:
        var v45 string = v2
        _ = v45
    }
    v46 := GenS4{F0: 1, F1: 0, F2: -1}
    v47 := GenS4{F0: -1, F1: -128, F2: 1000}
    fmt.Println(v23, v2, v3)
    for fi48 := 0; fi48 < 10; fi48++ {
        var v49 int64 = v23
        v50 := GenS1{F0: v3, F1: 2}
        _ = fi48
        _ = v49
        _ = v50
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v46
    _ = v47
}
