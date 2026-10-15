package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint; F1 int; F2 int64 }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 int8; F1 int64 }
func genFn0(p0 uint8, p1 uint16) uint16 {
    v7 := GenS1{F0: (p0 & (p0 * 0))}
    p1 -= 42
    _ = p0
    _ = p1
    _ = v7
    return (^(p1))
}

func main() {
    v1 := []string{"hello"}
    v2 := make(map[uint32]uint32)
    var v3 int8 = 127
    var v4 int = 2147483647
    v5 := []string{"foo bar", "hello", "a", ("a" + "hello")}
    v6 := []int{(v4 % ((v4) | 1)), v4, (v4 << ((608) & 63)), 0}
    _ = len(v1)
    if (!(416 <= 1000)) {
        v8 := GenS2{F0: 4294967295}
        switch (v4 << ((4294967295) & 63)) {
        case -9223372036854775808:
            fmt.Println(v4, v3)
        case -944:
            v9 := GenS2{F0: 0}
            _ = v9
        default:
            var v10 int64 = 9223372036854775807
            _ = v10
        }
        switch "\n\t" {
        case "\n\t":
            delete(v2, 42)
        }
        _ = v8
    }
    v11 := GenS3{F0: 4294967295}
    _ = len(v2)
    if true {
        for fi12 := 0; fi12 < 16; fi12++ {
            switch int(v3) {
            case 127:
                fmt.Println(fi12, fi12, fi12)
            case 257:
                fmt.Println(fi12)
            case 2147483647:
                fmt.Println(fi12, v3, fi12)
            }
            switch ("" + "a") {
            case "a":
                var v13 int64 = -1
                _ = v13
            case "":
                v3 = v3
            default:
                var v14 int16 = 42
                _ = v14
            }
            var v15 int64 = 9223372036854775807
            _ = fi12
            _ = v15
        }
    }
    _ = len(v2)
    if false {
        var v16 float32 = 3.14159
        var v17 float32 = v16
        _ = len(v2)
        _ = v16
        _ = v17
    }
    v5 = append(v5, "")
    if false {
        v3 += -1
    } else {
        var v18 float64 = (-2.25 + 1e10)
        var v19 bool = true
        _ = v18
        _ = v19
    }
    _ = len(v2)
    for fi20 := 0; fi20 < 2; fi20++ {
        var v21 string = ((("foo bar" + "") + "") + "")
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        for fi23 := 0; fi23 < 11; fi23++ {
            if false {
                if true {
                    var v24 int16 = -32768
                    _ = v24
                } else {
                    v25 := GenS4{F0: 1, F1: fi23, F2: -1}
                    switch (v4 - v4) {
                    case 1:
                        v26 := GenS1{F0: 1}
                        _ = v26
                    case -206:
                        if len(v1) > 0 {
                            v1[0] = ""
                        }
                    default:
                        switch (^(fi22)) {
                        case 4294967295:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case 191:
                            if len(v5) > 0 {
                                v5[0] = ("\n\t" + "")
                            }
                        case -414:
                            v4 = 1
                        default:
                            v27 := v2[42]
                            _ = v27
                        }
                    }
                    delete(v2, 521)
                    if (255 > 127) {
                        if true {
                            switch fi22 {
                            case 238:
                                var v28 uint = 590
                                v2[uint32(v28)] = 0
                                v4 += 42
                                _ = v28
                            case -1:
                                v29 := v2[2147483648]
                                _ = v29
                            case 43:
                                v4 -= v4
                            default:
                                v3 -= 1
                            }
                            v3 = v3
                            if true {
                                var v30 uint16 = 0
                                _ = v30
                            }
                            switch fi22 {
                            case -2147483648:
                                v31 := GenS3{F0: 255}
                                _ = v31
                            default:
                                v4 += (fi23 % (((fi23 / (((^(fi22))) | 1))) | 1))
                            }
                            switch 32767 {
                            case 127:
                                var v32 int16 = -1
                                _ = v32
                            case 138:
                                v4 -= (fi23 << ((4294967295) & 63))
                            case -350:
                                v3 -= 76
                            default:
                                if true {
                                    v4 = 1
                                    fmt.Println(fi23, v4, v3)
                                    if len(v1) > 0 {
                                        v1[0] = "a"
                                    }
                                }
                            }
                        } else {
                            v33 := GenS3{F0: 4294967295}
                            _ = v33
                        }
                        switch (fi22 % ((fi22) | 1)) {
                        case 128:
                            if (false || false) {
                                switch "foo bar" {
                                case "foo bar":
                                    var v34 uint64 = 1
                                    _ = v34
                                default:
                                    v4 = fi22
                                }
                            }
                        case -9223372036854775808:
                            if true {
                                if (2 <= 174) {
                                    switch (-(v4)) {
                                    case 127:
                                        _ = len(v6)
                                        v35 := GenS5{F0: 0, F1: 127}
                                        _ = v35
                                    case 4294967295:
                                        fmt.Println(fi22, fi23)
                                    case -9223372036854775808:
                                        fmt.Println(v4, v3, fi22)
                                    default:
                                        fmt.Println(fi23, fi23, fi23)
                                    }
                                    fmt.Println(fi23, v4)
                                } else {
                                    v3 = (v3 << ((126) & 63))
                                }
                                v4 = (-(fi22))
                                if true {
                                    fmt.Println(v4, fi22, v4)
                                    v36 := GenS2{F0: 210}
                                    _ = v36
                                }
                            } else {
                                fmt.Println(v4, v3, fi22)
                                fmt.Println(v4, v4)
                            }
                        }
                        if false {
                            fmt.Println(v3, v3)
                            v3 += v3
                            fmt.Println(v4, fi23)
                            switch -689 {
                            case 2147483647:
                                if (!true) {
                                    v4 = 1
                                    var v37 float32 = 0.1
                                    _ = v37
                                }
                            case 4294967295:
                                v3 = (-(v3))
                            case 9223372036854775807:
                                v3 = v3
                            }
                            v3 += 127
                        }
                    }
                    _ = v25
                }
            } else {
                var v38 int8 = v3
                fmt.Println(v4, fi22, fi23)
                _ = v38
            }
            if (true || true) {
                var v39 uint16 = 65535
                v40 := GenS6{F0: 127, F1: 42}
                _ = v39
                _ = v40
            }
            _ = fi23
        }
        _ = fi22
    }
    var v41 float64 = float64(1e10)
    var v42 uint = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
