package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8 }
func genFn0() uint {
    var v7 int = 801
    v7 += 127
    fmt.Println(v7)
    var v8 float32 = 0.0
    _ = v7
    _ = v8
    return 963
}
func genFn1(p0 uint, p1 int32, p2 int8) int8 {
    switch ("" + "foo bar") {
    case "\n\t":
        switch 95 {
        case 4294967295:
            p2 -= (p2 * p2)
        }
    case "":
        fmt.Println(p1, p0)
    }
    if false {
        fmt.Println(p1, p1)
        fmt.Println(p0, p1, p0)
    } else {
        var v9 int64 = -1
        fmt.Println(p0)
        if ("a" != ("" + "foo bar")) {
            switch "a" {
            case "a":
                if true {
                    if true {
                        v9 = -648
                        if (!true) {
                            v10 := GenS1{F0: p2, F1: 2, F2: 0}
                            _ = v10
                        }
                    }
                    var v11 uint8 = 245
                    _ = v11
                }
            case "foo bar":
                for fi12 := 0; fi12 < 11; fi12++ {
                    fmt.Println(fi12, v9)
                    for fi13 := 0; fi13 < 7; fi13++ {
                        fmt.Println(p1)
                        var v14 uint64 = 217
                        var v15 float32 = float32(0.0)
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    switch "a" {
                    case "hello":
                        var v16 uint = p0
                        _ = v16
                    default:
                        p0 = p0
                    }
                    _ = fi12
                }
            default:
                var v17 uint16 = 42
                _ = v17
            }
        }
        _ = v9
    }
    fmt.Println(p0, p2)
    p2 += p2
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 int16 = -217
    var v2 float64 = 1.5
    var v3 int8 = -1
    var v4 uint = 2
    v5 := make(map[string]bool)
    var v6 int = 245
    for fi18 := 0; fi18 < 6; fi18++ {
        v19 := []int{(^(fi18)), v6}
        v20 := GenS1{F0: genFn1(v4, 0, 24), F1: (^(v4)), F2: fi18}
        _ = fi18
        _ = v19
        _ = v20
    }
    v3 = v3
    fmt.Println(v1, v4)
    if true {
        var v21 int8 = (v3 * int8(v1))
        _ = v21
    } else {
        for fi22 := 0; fi22 < 15; fi22++ {
            var v23 float32 = -2.25
            v3 = genFn1(v4, -1, v3)
            _ = fi22
            _ = v23
        }
    }
    fmt.Println(v1)
    var v24 float32 = 1.5
    v25 := GenS1{F0: int8(v6), F1: 255, F2: (v6 ^ v6)}
    for fi26 := 0; fi26 < 12; fi26++ {
        if false {
            switch "hello" {
            case "foo bar":
                fmt.Println(v4)
            case "hello":
                v6 -= (v6 >> (((v4 - (v4 & v4))) & 63))
            case "":
                v27 := []int{}
                _ = v27
            }
            for fi28 := 0; fi28 < 4; fi28++ {
                switch (fi28 << ((genFn0()) & 63)) {
                case 4294967295:
                    if false {
                        var v29 int64 = -268
                        if false {
                            var v30 uint = 18446744073709551615
                            _ = v30
                        } else {
                            v24 += 3.14159
                            v1 = (-(v1))
                        }
                        if false {
                            var v31 float32 = (1e10 / v24)
                            if false {
                                switch fi26 {
                                case 32768:
                                    v24 = (v24 * (v24 / v31))
                                    var v32 bool = ((!false) || false)
                                    _ = v32
                                case -299:
                                    v33 := []int{fi28}
                                    _ = v33
                                default:
                                    var v34 bool = (v31 >= (0.0 / (float32((v2 * v2)) / v24)))
                                    _ = v34
                                }
                                _ = len(v5)
                            }
                            _ = v31
                        }
                        var v35 string = ("hello" + ("a" + ("foo bar" + ("hello" + ("a" + "\n\t")))))
                        v29 += v29
                        delete(v5, "a")
                        _ = v29
                        _ = v35
                    } else {
                        v24 = 1.5
                        v36 := GenS2{F0: 114}
                        var v37 int8 = v3
                        var v38 string = "foo bar"
                        _ = v36
                        _ = v37
                        _ = v38
                    }
                default:
                    if (((!false) || (2 != 128)) && (1 > (v3 & v3))) {
                        v4 -= (v4 + v4)
                        v39 := GenS1{F0: (v3 * int8(fi28)), F1: (v4 & (v4 + (v4 << (((v4 | genFn0())) & 63)))), F2: -2147483648}
                        v40 := []int{fi26, v6, (fi28 + (fi26 + (^(fi26)))), fi26}
                        _ = v39
                        _ = v40
                    }
                }
                _ = fi28
            }
            var v41 uint16 = 0
            _ = v41
        }
        switch ((("a" + ("\n\t" + "a")) + "\n\t") + ("" + ("foo bar" + "hello"))) {
        case "\n\t":
            v42 := GenS3{F0: v3}
            _ = v42
        }
        v43 := v5[(("foo bar" + "a") + ("hello" + "a"))]
        _ = fi26
        _ = v43
    }
    v44 := GenS2{F0: 42}
    if (278 > 0) {
        v45 := GenS2{F0: 125}
        if false {
            fmt.Println(v2, v2)
            var v46 int32 = 42
            _ = v46
        } else {
            var v47 uint8 = 1
            _ = v47
        }
        _ = len(v5)
        _ = v45
    } else {
        if (42 != 2) {
            if true {
                v4 = (v4 >> (((v4 >> ((v4) & 63))) & 63))
                v4 = v4
            } else {
                for fi48 := 0; fi48 < 14; fi48++ {
                    v3 -= 82
                    var v49 int16 = (v1 ^ (v1 >> ((520) & 63)))
                    _ = fi48
                    _ = v49
                }
                v50 := []int{(-(v6)), (v6 << (((v4 - 747)) & 63)), v6}
                if ((350 <= 1) && false) {
                    v2 = float64(float32(float64(v24)))
                    var v51 int64 = 42
                    _ = v51
                }
                _ = v50
            }
            fmt.Println(v6, v4, v1)
        }
        v2 += (1e10 - v2)
    }
    var v52 int8 = 7
    v52 = 100
    if (v2 != (v2 * 3.14159)) {
        switch (v6 & 42) {
        case -162:
            var v53 uint8 = 0
            _ = v53
        case 1:
            fmt.Println(v52, v1)
        }
        v54 := GenS1{F0: (v52 % ((100) | 1)), F1: genFn0(), F2: v6}
        _ = v54
    } else {
        v55 := []int{446}
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v52
    _ = v6
}
