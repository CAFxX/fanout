package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint; F1 uint32 }
func genFn0(p0 int16, p1 uint8) bool {
    if false {
        for fi7 := 0; fi7 < 12; fi7++ {
            var v8 bool = true
            var v9 uint32 = 269
            fmt.Println(p1)
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    var v10 uint8 = (p1 | p1)
    _ = p0
    _ = p1
    _ = v10
    return (255 == 4294967295)
}

func main() {
    var v1 int = 1
    v2 := []string{}
    var v3 int32 = 740
    v4 := make(map[int]int)
    var v5 int8 = -34
    var v6 uint64 = 18446744073709551615
    var v11 int16 = 32767
    var v12 string = "foo bar"
    v6 -= 42
    for fi13 := 0; fi13 < 3; fi13++ {
        delete(v4, 753)
        _ = fi13
    }
    v4[-2147483648] = 128
    var v14 uint8 = uint8(v6)
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 uint32 = 25
        var v17 float32 = 0.0
        v5 = v5
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := GenS1{F0: v14, F1: v1, F2: ((("" + v12) + ((v12 + "a") + v12)) + v12)}
    fmt.Println(v14, v6, v11)
    if (2 >= 1) {
        var v19 int = v1
        switch 0 {
        case -102:
            var v20 int16 = v11
            _ = v20
        case -841:
            v21 := GenS2{F0: true}
            _ = v21
        default:
            v12 = (((v12 + "") + (((v12 + "a") + v12) + ((v12 + "foo bar") + ("foo bar" + v12)))) + v12)
        }
        _ = v19
    }
    var v22 uint = 341
    v23 := GenS3{F0: -125, F1: 544}
    switch "hello" {
    case "":
        for fi24 := 0; fi24 < 17; fi24++ {
            var v25 bool = (32768 != (-(fi24)))
            v25 = v25
            _ = fi24
            _ = v25
        }
    case "\n\t":
        for fi26 := 0; fi26 < 5; fi26++ {
            var v27 float64 = 0.0
            var v28 bool = true
            var v29 string = ("foo bar" + v12)
            _ = fi26
            _ = v27
            _ = v28
            _ = v29
        }
    case "foo bar":
        if ((v5 - v5) < 123) {
            v11 -= v11
            if ((!(!true)) || true) {
                v1 = int(v11)
                v30 := GenS3{F0: 809, F1: 891}
                for fi31 := 0; fi31 < 13; fi31++ {
                    if (int8(v1) >= v5) {
                        var v32 bool = (1 == v3)
                        fmt.Println(v11)
                        v14 = (^(v14))
                        _ = v32
                    }
                    if (true && ((v5 | v5) == 127)) {
                        var v33 int8 = (v5 % (((v5 % ((-72) | 1))) | 1))
                        fmt.Println(v5, v1)
                        var v34 uint8 = v14
                        _ = v33
                        _ = v34
                    }
                    _ = fi31
                }
                _ = v30
            } else {
                v6 = v6
                var v35 float32 = 1.5
                v36 := v4[(v1 & v1)]
                _ = v35
                _ = v36
            }
            v1 = int(v11)
        }
    default:
        v37 := GenS3{F0: -585, F1: 32767}
        _ = v37
    }
    if true {
        for fi38 := 0; fi38 < 10; fi38++ {
            v39 := GenS1{F0: v14, F1: fi38, F2: ((v12 + "a") + v12)}
            _ = fi38
            _ = v39
        }
    } else {
        if false {
            v22 = (^(v22))
            if (((v1 + v1) >= v1) || ((v22 - (v22 << ((4294967295) & 63))) <= (v22 + (v22 * 1)))) {
                var v40 int = int(v6)
                _ = v40
            }
        }
    }
    v41 := GenS4{F0: 899}
    if (v6 >= 1) {
        switch "\n\t" {
        case "hello":
            var v42 uint16 = 2
            _ = v42
        case "foo bar":
            fmt.Println(v11, v6, v5)
        default:
            for fi43 := 0; fi43 < 13; fi43++ {
                for fi44 := 0; fi44 < 2; fi44++ {
                    v2 = append(v2, v12)
                    v22 = (v22 % (((^(v22))) | 1))
                    v45 := GenS2{F0: (false && false)}
                    switch (v12 + "foo bar") {
                    case "\n\t":
                        v46 := GenS3{F0: 0, F1: 1}
                        _ = v46
                    }
                    _ = fi44
                    _ = v45
                }
                for fi47 := 0; fi47 < 16; fi47++ {
                    if ((!false) && false) {
                        switch "hello" {
                        case "foo bar":
                            v48 := GenS2{F0: (!(!(v6 <= (^(v6)))))}
                            _ = v48
                        default:
                            if (!((0.0 + (3.14159 * (1.5 * float32((1e10 * -2.25))))) < 0.0)) {
                                v49 := GenS5{F0: 65535, F1: uint32(v22)}
                                _ = v49
                            }
                        }
                        v2 = append(v2, "hello")
                    } else {
                        var v50 uint32 = 2
                        v51 := GenS5{F0: 108, F1: v50}
                        _ = v50
                        _ = v51
                    }
                    _ = fi47
                }
                v22 = 2
                _ = fi43
            }
        }
        var v52 int16 = (v11 % (((v11 ^ (v11 >> ((501) & 63)))) | 1))
        _ = v52
    } else {
        v2 = append(v2, v12)
        fmt.Println(v22)
    }
    fmt.Println(v12, v14)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
