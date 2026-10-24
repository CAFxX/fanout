package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 string; F2 int32 }
type GenS3 struct { F0 uint; F1 int32; F2 uint32 }
type GenS4 struct { F0 int64 }
func genFn0(p0 uint, p1 bool, p2 int32) string {
    for fi5 := 0; fi5 < 9; fi5++ {
        v6 := GenS1{F0: 42}
        for fi7 := 0; fi7 < 7; fi7++ {
            v8 := []int{(fi5 & (fi7 & fi5))}
            if (!p1) {
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                var v9 bool = (p0 == 4294967295)
                _ = v9
            }
            switch "hello" {
            case "a":
                var v10 uint16 = 32767
                p0 = p0
                p2 -= p2
                _ = v10
            case "\n\t":
                var v11 int16 = -32768
                _ = v11
            default:
                p2 -= 2147483647
                var v12 int8 = 1
                fmt.Println(p2, v12)
                _ = v12
            }
            _ = fi7
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    if p1 {
        var v13 uint16 = 2
        p0 = p0
        v14 := GenS2{F0: -1, F1: "\n\t", F2: (-(p2))}
        _ = v13
        _ = v14
    } else {
        if p1 {
            for fi15 := 0; fi15 < 11; fi15++ {
                for fi16 := 0; fi16 < 7; fi16++ {
                    v17 := GenS1{F0: 0}
                    if p1 {
                        if p1 {
                            var v18 uint8 = 255
                            switch 1 {
                            case 1:
                                switch ("\n\t" + "\n\t") {
                                case "a":
                                    var v19 bool = false
                                    _ = v19
                                default:
                                    var v20 float32 = 1e10
                                    _ = v20
                                }
                            case 32768:
                                fmt.Println(p0, p1)
                            case 32767:
                                p1 = ((4294967295 > 65535) && p1)
                            default:
                                p2 += (p2 ^ 42)
                            }
                            _ = v18
                        }
                        if ((((p1 && ((p0 ^ (p0 << (((^(p0))) & 63))) < (p0 / (((p0 + 0)) | 1)))) || p1) && (-88 <= 513)) && (!p1)) {
                            p0 -= p0
                            switch ("\n\t" + ((("a" + "hello") + ("a" + ("" + (("hello" + (("\n\t" + "") + "hello")) + "a")))) + ("a" + "hello"))) {
                            case "a":
                                v21 := []int{fi15, fi16, fi15}
                                _ = v21
                            case "hello":
                                p0 = (p0 >> ((p0) & 63))
                            case "foo bar":
                                if (42 == fi16) {
                                    p1 = (682 > 0)
                                } else {
                                    v22 := []int{}
                                    _ = v22
                                }
                                var v23 uint64 = 968
                                v24 := []int{-515, int(v23)}
                                _ = v23
                                _ = v24
                            default:
                                p2 = 129
                            }
                        } else {
                            var v25 float32 = 1.5
                            v26 := GenS2{F0: -91, F1: "hello", F2: 1}
                            _ = v25
                            _ = v26
                        }
                    }
                    _ = fi16
                    _ = v17
                }
                _ = fi15
            }
            fmt.Println(p2)
            switch ("\n\t" + ("" + "hello")) {
            case "":
                var v27 float32 = 0.0
                _ = v27
            case "\n\t":
                var v28 int16 = 42
                _ = v28
            default:
                p0 -= 255
            }
        }
    }
    v29 := []int{32767, -60}
    _ = p0
    _ = p1
    _ = p2
    _ = v29
    return "a"
}

func main() {
    var v1 bool = false
    var v2 string = (("a" + ("foo bar" + "foo bar")) + "hello")
    var v3 int = -661
    var v4 float32 = (3.14159 * 0.1)
    var v30 uint16 = 335
    v1 = (3.14159 != (((0.0 / float64(float32(float64(0.1)))) + (0.0 * float64((3.14159 / v4)))) / 1e10))
    var v31 int = (v3 - v3)
    for fi32 := 0; fi32 < 11; fi32++ {
        v30 -= (v30 << ((0) & 63))
        var v33 int8 = 22
        _ = fi32
        _ = v33
    }
    var v34 int32 = 1
    var v35 float64 = 1.5
    switch "hello" {
    case "a":
        v34 = (^(v34))
    case "\n\t":
        v4 -= 1.5
    }
    v36 := GenS3{F0: 255, F1: v34, F2: 2}
    v37 := []int{(v3 & 32767), (v3 + v31)}
    var v38 int8 = 100
    if (!((0.0 / (1.5 / v4)) >= (-2.25 / (0.1 / (((((v4 - (v4 / (v4 / 1.5))) + v4) - float32(v35)) * 3.14159) * ((((-2.25 - 1.5) - 3.14159) / (float32(0.1) * float32(float64(v4)))) * v4)))))) {
        switch v2 {
        case "hello":
            v37 = append(v37, v3)
        case "foo bar":
            _ = len(v37)
        }
        var v39 int8 = v38
        _ = v39
    } else {
        v38 += v38
        v40 := GenS4{F0: -1}
        _ = v40
    }
    v4 -= v4
    var v41 bool = true
    for fi42 := 0; fi42 < 6; fi42++ {
        v43 := GenS1{F0: (v30 ^ v30)}
        v44 := GenS3{F0: uint(v30), F1: (v34 ^ -2147483648), F2: 4294967295}
        _ = fi42
        _ = v43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v41
}
