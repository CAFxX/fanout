package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int; F1 int64; F2 uint }
func genFn0(p0 string, p1 uint, p2 uint) int {
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := GenS1{F0: (fi7 << (((p2 & 2)) & 63))}
        v9 := []int{fi7, 9223372036854775807}
        switch p0 {
        case "":
            switch p0 {
            case "\n\t":
                var v10 bool = true
                _ = v10
            case "foo bar":
                if true {
                    for fi11 := 0; fi11 < 15; fi11++ {
                        switch p0 {
                        case "":
                            var v12 uint = (p1 / (((p2 & (^(p1)))) | 1))
                            _ = v12
                        case "\n\t":
                            p1 = (^(p1))
                        default:
                            var v13 int = 42
                            _ = v13
                        }
                        v14 := GenS1{F0: 0}
                        var v15 uint16 = 255
                        var v16 int8 = 1
                        fmt.Println(p2)
                        v17 := GenS2{F0: -137, F1: fi11}
                        _ = fi11
                        _ = v14
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                }
            }
        case "a":
            p0 = p0
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    if true {
        var v18 float32 = 3.14159
        v19 := []int{}
        _ = v18
        _ = v19
    } else {
        p1 -= (^(p2))
    }
    var v20 float32 = ((1.5 / -2.25) / (1.5 + (-2.25 * (((-2.25 + (0.0 / 1.5)) + -2.25) * (0.1 - (3.14159 * (float32(1.5) - (0.1 - 1e10))))))))
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return 127
}
func genFn1(p0 uint16, p1 bool, p2 uint8) string {
    p0 -= 1
    var v21 bool = (p0 >= (p0 >> ((4294967295) & 63)))
    p0 += p0
    fmt.Println(p2)
    p1 = v21
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return "foo bar"
}

func main() {
    var v1 int = 127
    v2 := []uint32{374, 2147483647, 2, 2147483647}
    v3 := make(map[int]string)
    var v4 int64 = -9223372036854775808
    v5 := make(map[int]int)
    v6 := make(map[string]bool)
    v1 += v1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v2 = append(v2, 2)
    for fi22 := 0; fi22 < 7; fi22++ {
        if true {
            fmt.Println(v1, v1, v1)
            for fi23 := 0; fi23 < 19; fi23++ {
                switch (fi22 | 4294967295) {
                case 718:
                    var v24 uint32 = 2147483647
                    _ = v24
                case -2147483648:
                    v5[genFn0("foo bar", 4294967295, 42)] = -917
                }
                _ = fi23
            }
            if (116 != 347) {
                delete(v5, 42)
            }
        }
        switch v1 {
        case 173:
            fmt.Println(v1, v1, fi22)
        }
        _ = fi22
    }
    v25 := GenS3{F0: 18446744073709551615, F1: 518}
    for fi26 := 0; fi26 < 19; fi26++ {
        for fi27 := 0; fi27 < 18; fi27++ {
            var v28 uint64 = 18446744073709551615
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    v4 += (-(v4))
    var v29 int = 42
    fmt.Println(v4, v4, v1)
    for fi30 := 0; fi30 < 17; fi30++ {
        if (((((false || (1000 == -673)) && false) && (-32768 != 1000)) && true) && true) {
            var v31 uint32 = 445
            _ = v31
        }
        v32 := GenS4{F0: 161}
        var v33 int32 = 0
        _ = fi30
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 12; fi34++ {
        v1 = genFn0(("foo bar" + "hello"), 255, 1)
        if (!true) {
            for fi35 := 0; fi35 < 13; fi35++ {
                v36 := GenS5{F0: -1, F1: -1, F2: 735}
                var v37 float64 = -2.25
                _ = fi35
                _ = v36
                _ = v37
            }
        }
        v1 += (-(fi34))
        _ = fi34
    }
    switch (v1 ^ genFn0(((("hello" + "foo bar") + "foo bar") + ("foo bar" + ((((("" + "foo bar") + "hello") + "foo bar") + ("\n\t" + "")) + ("hello" + ("" + "foo bar"))))), 4294967295, 455)) {
    case 4294967295:
        fmt.Println(v29, v4)
    case 1:
        fmt.Println(v29, v29, v4)
    case 128:
        fmt.Println(v29)
    }
    if len(v2) > 0 {
        v2[0] = 2
    }
    switch "a" {
    case "":
        var v38 int32 = -2147483648
        _ = v38
    case "a":
        var v39 float64 = 1e10
        _ = v39
    default:
        fmt.Println(v4, v29)
    }
    if true {
        switch "foo bar" {
        case "a":
            var v40 int = v1
            _ = v40
        case "foo bar":
            if len(v2) > 0 {
                v2[0] = 1
            }
        }
    }
    for fi41 := 0; fi41 < 15; fi41++ {
        v29 -= v1
        var v42 bool = (1 == 722)
        _ = fi41
        _ = v42
    }
    var v43 uint = 606
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
