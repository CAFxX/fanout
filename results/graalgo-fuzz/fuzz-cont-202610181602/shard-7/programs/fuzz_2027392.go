package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 int16; F1 int64; F2 int16 }
func genFn0(p0 int) string {
    fmt.Println(p0)
    for fi5 := 0; fi5 < 4; fi5++ {
        p0 -= 1
        for fi6 := 0; fi6 < 2; fi6++ {
            p0 -= 127
            v7 := []int{127}
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    p0 = (-(p0))
    for fi8 := 0; fi8 < 19; fi8++ {
        v9 := []int{p0, (p0 * p0), (^(p0))}
        _ = fi8
        _ = v9
    }
    _ = p0
    return ("foo bar" + "a")
}
func genFn1(p0 int, p1 uint32) uint {
    switch "foo bar" {
    case "a":
        switch (("hello" + ("hello" + "")) + (((("foo bar" + "foo bar") + "a") + "") + (("a" + ("a" + "")) + "\n\t"))) {
        case "a":
            v10 := []int{107}
            _ = v10
        case "":
            fmt.Println(p1)
        case "foo bar":
            p1 = p1
        default:
            switch "\n\t" {
            case "foo bar":
                switch 128 {
                case 1:
                    switch ("" + ("" + "foo bar")) {
                    case "hello":
                        p1 = 0
                    }
                case 128:
                    p0 = (^(p0))
                case 32767:
                    v11 := GenS1{F0: -800, F1: 255, F2: 42}
                    _ = v11
                default:
                    p1 += (p1 >> ((688) & 63))
                }
            case "":
                for fi12 := 0; fi12 < 12; fi12++ {
                    v13 := []int{}
                    p0 = (-(fi12))
                    _ = fi12
                    _ = v13
                }
            case "a":
                var v14 float32 = (float32(float64(0.1)) * (1.5 / 0.1))
                _ = v14
            }
        }
    case "":
        if (!(("" + "") >= "\n\t")) {
            v15 := GenS1{F0: 42, F1: 255, F2: (^(p1))}
            _ = v15
        }
    }
    var v16 int32 = -2147483648
    _ = p0
    _ = p1
    _ = v16
    return 2
}

func main() {
    var v1 bool = true
    var v2 int = 32768
    var v3 uint64 = 42
    v4 := make(map[int]int)
    v4[v2] = v2
    v2 = (v2 + -9223372036854775808)
    if ((^(v3)) >= (v3 * 49)) {
        if (!((((-2.25 / float64((float32(0.1) + (3.14159 - (0.1 + (0.1 - 1.5)))))) * 0.0) - (1.5 / 1.5)) > 0.1)) {
            var v17 int32 = -979
            _ = v17
        }
    }
    if (3.14159 >= 3.14159) {
        fmt.Println(v1, v1)
        var v18 bool = ((v1 && (151 > 128)) || (127 == 128))
        if ((0 > -24) || ((690 > 18446744073709551615) && false)) {
            v19 := GenS2{F0: (v2 & v2), F1: v18}
            v2 = (v2 + 805)
            switch (v2 | (v2 & v2)) {
            case 1:
                v20 := GenS2{F0: v2, F1: (1 >= 1)}
                _ = v20
            default:
                switch ("a" + ("\n\t" + "foo bar")) {
                case "\n\t":
                    v18 = ((float64(1.5) - ((0.1 - (1.5 / 0.1)) / 3.14159)) > 1.5)
                case "":
                    switch "a" {
                    case "a":
                        delete(v4, (v2 + 395))
                    case "":
                        var v21 bool = false
                        _ = v21
                    default:
                        if (0 >= 65535) {
                            v22 := GenS2{F0: -1, F1: v1}
                            for fi23 := 0; fi23 < 4; fi23++ {
                                if v18 {
                                    switch "foo bar" {
                                    case "\n\t":
                                        for fi24 := 0; fi24 < 20; fi24++ {
                                            v18 = v1
                                            v1 = false
                                            if (2147483647 >= 65535) {
                                                if (v18 && (645 < 484)) {
                                                    var v25 float64 = (3.14159 / 1.0)
                                                    _ = v25
                                                } else {
                                                    var v26 string = (("hello" + "a") + "a")
                                                    v27 := []int{fi23, 703, (^(fi24)), -1}
                                                    switch v2 {
                                                    case 4294967295:
                                                        fmt.Println(fi23)
                                                    default:
                                                        var v28 string = (((((("" + (v26 + "")) + v26) + v26) + v26) + "") + v26)
                                                        _ = v28
                                                    }
                                                    v18 = v18
                                                    _ = v26
                                                    _ = v27
                                                }
                                                v18 = (!true)
                                                v29 := []int{}
                                                _ = v29
                                            }
                                            _ = fi24
                                        }
                                    case "foo bar":
                                        for fi30 := 0; fi30 < 11; fi30++ {
                                            var v31 uint64 = 445
                                            _ = fi30
                                            _ = v31
                                        }
                                    case "a":
                                        switch 9223372036854775807 {
                                        case 1:
                                            for fi32 := 0; fi32 < 14; fi32++ {
                                                var v33 int16 = -32768
                                                _ = fi32
                                                _ = v33
                                            }
                                        case 500:
                                            v1 = false
                                        case 128:
                                            v34 := []int{fi23, fi23, -1}
                                            _ = v34
                                        }
                                    default:
                                        v1 = v18
                                    }
                                } else {
                                    fmt.Println(v18)
                                    v18 = v1
                                    v4[fi23] = 1
                                }
                                _ = fi23
                            }
                            _ = v22
                        }
                    }
                default:
                    for fi35 := 0; fi35 < 18; fi35++ {
                        for fi36 := 0; fi36 < 5; fi36++ {
                            _ = len(v4)
                            _ = fi36
                        }
                        _ = fi35
                    }
                }
            }
            _ = v19
        }
        _ = v18
    }
    var v37 uint = genFn1(v2, 2)
    var v38 int16 = 32767
    v39 := []int{-575, (v2 | v2), (v2 & v2), -38}
    v40 := GenS2{F0: (v2 - (v2 & (v2 % (((v2 - v2)) | 1)))), F1: v1}
    _ = len(v4)
    v41 := GenS1{F0: -1, F1: 984, F2: 65535}
    if ((v3 >> ((4294967295) & 63)) <= v3) {
        v39 = append(v39, (v2 ^ -1))
        if (v1 && (v1 && (true && (("a" >= "foo bar") || (0 >= 4294967295))))) {
            var v42 float64 = 3.14159
            _ = v42
        } else {
            _ = len(v4)
            if (v1 && true) {
                for fi43 := 0; fi43 < 17; fi43++ {
                    switch v2 {
                    case 128:
                        fmt.Println(v37)
                    default:
                        for fi44 := 0; fi44 < 2; fi44++ {
                            fmt.Println(v37, v37, v2)
                            _ = fi44
                        }
                    }
                    _ = fi43
                }
                v45 := GenS3{F0: (v3 & 18446744073709551615), F1: ""}
                _ = v45
            } else {
                if len(v39) >= 2 {
                    v39 = v39[1:len(v39)]
                }
            }
            v46 := GenS4{F0: v38, F1: 9223372036854775807, F2: 1}
            _ = v46
        }
    }
    v2 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
