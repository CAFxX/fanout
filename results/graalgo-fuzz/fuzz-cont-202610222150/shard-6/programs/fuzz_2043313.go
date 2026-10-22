package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 string; F1 int }
func genFn0(p0 int, p1 int64) int16 {
    if (!(false && (int16(p1) >= 258))) {
        p1 += p1
        var v5 uint16 = 905
        _ = v5
    } else {
        switch ("hello" + (("hello" + "") + "hello")) {
        case "\n\t":
            fmt.Println(p1)
        }
        v6 := GenS1{F0: 65535, F1: p1}
        p1 += (p1 / ((924) | 1))
        _ = v6
    }
    switch "foo bar" {
    case "foo bar":
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 uint64 = 4294967295
            p1 = 9223372036854775807
            _ = fi7
            _ = v8
        }
    case "":
        var v9 int32 = -1
        _ = v9
    case "\n\t":
        switch "" {
        case "foo bar":
            p0 = -183
        default:
            fmt.Println(p1, p0)
        }
    default:
        if true {
            var v10 float64 = 0.0
            var v11 string = "hello"
            _ = v10
            _ = v11
        }
    }
    if true {
        if (!(false || true)) {
            var v12 uint8 = 1
            p1 = p1
            _ = v12
        }
        for fi13 := 0; fi13 < 7; fi13++ {
            var v14 uint32 = 768
            _ = fi13
            _ = v14
        }
    }
    v15 := []int{p0, 128}
    _ = p0
    _ = p1
    _ = v15
    return 817
}
func genFn1(p0 int32, p1 string) uint64 {
    v16 := []int{402}
    if false {
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
        if len(v16) > 0 {
            v16[0] = 2147483647
        }
        p1 = p1
    } else {
        var v17 uint8 = 127
        _ = v17
    }
    p0 += p0
    if ((true || false) && false) {
        p0 -= 0
        var v18 uint64 = 2
        _ = v18
    }
    _ = p0
    _ = p1
    _ = v16
    return 1
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]int)
    v3 := make(map[uint32]bool)
    var v4 float32 = float32(0.0)
    v19 := []int{-796, -440}
    v4 += (float32(3.14159) - v4)
    var v20 float32 = ((v4 + float32(0.1)) / -2.25)
    var v21 float64 = (1e10 / (-2.25 * 1.5))
    fmt.Println(v20, v21)
    v22 := GenS2{F0: 255, F1: -935}
    if true {
        for fi23 := 0; fi23 < 1; fi23++ {
            var v24 int32 = -1
            fmt.Println(v24, v4, v20)
            if true {
                v25 := GenS3{F0: ("\n\t" + (("hello" + ("" + ("" + "hello"))) + ("foo bar" + ("foo bar" + "a")))), F1: fi23}
                _ = v25
            } else {
                v26 := v3[765]
                for fi27 := 0; fi27 < 11; fi27++ {
                    fmt.Println(v26, v4, v24)
                    var v28 float32 = ((v4 * ((1.5 * v4) + ((-2.25 * (float32(v21) / 0.1)) + v4))) * v20)
                    if v26 {
                        v21 = v21
                        switch (fi27 & (fi27 % ((fi23) | 1))) {
                        case 1:
                            v24 = (-(v24))
                        case 32767:
                            switch (("\n\t" + "foo bar") + "a") {
                            case "foo bar":
                                v28 += ((v20 - (float32(0.1) / v4)) + 0.1)
                                var v29 int16 = 42
                                if (!(128 < 42)) {
                                    var v30 uint64 = 0
                                    var v31 int64 = 9223372036854775807
                                    v32 := GenS1{F0: 2147483647, F1: v31}
                                    _ = v30
                                    _ = v31
                                    _ = v32
                                }
                                _ = v29
                            case "\n\t":
                                v33 := GenS2{F0: 255, F1: (-(v24))}
                                _ = v33
                            default:
                                var v34 uint32 = 4294967295
                                _ = v34
                            }
                        }
                    }
                    _ = fi27
                    _ = v28
                }
                switch 128 {
                case -430:
                    if true {
                        v20 = v4
                    }
                case 2147483647:
                    if (0 <= 0) {
                        v35 := GenS2{F0: 65535, F1: (v24 << ((1) & 63))}
                        if v26 {
                            v19 = append(v19, (-(fi23)))
                            v36 := GenS3{F0: ("hello" + ("" + ((("hello" + ((("\n\t" + "\n\t") + ((("\n\t" + "a") + "\n\t") + "a")) + "hello")) + ((("" + "") + "") + "")) + ("" + "\n\t")))), F1: 9223372036854775807}
                            _ = v36
                        }
                        _ = v35
                    } else {
                        if ((((v24 ^ (v24 >> ((4294967295) & 63))) <= v24) && v26) || (-1 < genFn0((fi23 << ((1) & 63)), 9223372036854775807))) {
                            var v37 int = (fi23 - (fi23 + -2147483648))
                            switch v37 {
                            case -803:
                                var v38 bool = (2 >= 4294967295)
                                _ = v38
                            }
                            _ = v37
                        }
                    }
                case 4294967295:
                    var v39 float32 = (v4 * v20)
                    _ = v39
                }
                _ = v26
            }
            _ = fi23
            _ = v24
        }
    } else {
        var v40 int64 = 9223372036854775807
        _ = v40
    }
    for fi41 := 0; fi41 < 18; fi41++ {
        v42 := GenS3{F0: "foo bar", F1: (fi41 | fi41)}
        switch (fi41 >> ((4294967295) & 63)) {
        case 32768:
            v4 -= v20
        case 4294967295:
            v4 += float32(1.5)
        default:
            switch "a" {
            case "foo bar":
                v43 := GenS3{F0: "foo bar", F1: fi41}
                _ = v43
            }
        }
        _ = fi41
        _ = v42
    }
    for fi44 := 0; fi44 < 12; fi44++ {
        var v45 uint32 = 681
        _ = fi44
        _ = v45
    }
    for fi46 := 0; fi46 < 15; fi46++ {
        var v47 float64 = (0.1 + v21)
        v20 = (v4 + v20)
        _ = fi46
        _ = v47
    }
    v48 := GenS2{F0: 32767, F1: 2147483647}
    v4 = v20
    var v49 int16 = 32767
    if (-79 < v49) {
        var v50 float32 = v4
        fmt.Println(v49, v20)
        _ = v50
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v48
    _ = v49
}
