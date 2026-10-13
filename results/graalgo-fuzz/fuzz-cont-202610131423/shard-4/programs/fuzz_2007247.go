package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 uint }
func genFn0(p0 int32, p1 int8, p2 uint8) uint {
    p2 = 0
    p0 = int32(p2)
    for fi5 := 0; fi5 < 14; fi5++ {
        if false {
            v6 := []int{fi5, (^(fi5))}
            var v7 float64 = 0.0
            _ = v6
            _ = v7
        }
        var v8 string = ("foo bar" + "a")
        _ = fi5
        _ = v8
    }
    v9 := GenS1{F0: 945, F1: 1000}
    var v10 uint64 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 2
}
func genFn1(p0 uint64, p1 int16) int32 {
    for fi11 := 0; fi11 < 14; fi11++ {
        p1 = (p1 - p1)
        _ = fi11
    }
    switch 1 {
    case 4294967295:
        fmt.Println(p1, p1)
    case -683:
        fmt.Println(p1, p0)
    case 128:
        var v12 float64 = -2.25
        _ = v12
    }
    _ = p0
    _ = p1
    return -2147483648
}

func main() {
    var v1 uint = 2
    var v2 bool = false
    var v3 int16 = -563
    v4 := make(map[uint32]string)
    v13 := GenS1{F0: 0, F1: v3}
    v14 := GenS2{F0: 42}
    if v2 {
        for fi15 := 0; fi15 < 8; fi15++ {
            v1 -= 18446744073709551615
            switch "a" {
            case "":
                var v16 string = "a"
                _ = v16
            case "hello":
                var v17 int = fi15
                _ = v17
            case "foo bar":
                v1 = 18446744073709551615
            default:
                var v18 string = "hello"
                _ = v18
            }
            if v2 {
                if ((v2 && v2) && (!(447 < 42))) {
                    if (!v2) {
                        v2 = (((v3 * v3) >= -156) || false)
                    } else {
                        switch "hello" {
                        case "foo bar":
                            v19 := []int{}
                            _ = v19
                        case "":
                            if ((((313 < 42) || (42 > 9223372036854775807)) && v2) || v2) {
                                v2 = v2
                                v20 := GenS2{F0: 42}
                                v21 := v4[42]
                                _ = v20
                                _ = v21
                            }
                        case "hello":
                            if (true || v2) {
                                var v22 int16 = (v3 & v3)
                                _ = v22
                            }
                        }
                    }
                    v2 = false
                    var v23 bool = true
                    _ = v23
                } else {
                    for fi24 := 0; fi24 < 13; fi24++ {
                        var v25 int32 = 638
                        _ = fi24
                        _ = v25
                    }
                }
                if false {
                    v1 -= v1
                } else {
                    var v26 uint64 = 127
                    var v27 int8 = 20
                    _ = v26
                    _ = v27
                }
            }
            _ = fi15
        }
    }
    v3 -= (v3 ^ v3)
    var v28 int = 32768
    switch "foo bar" {
    case "\n\t":
        var v29 uint32 = 42
        _ = v29
    case "hello":
        if false {
            for fi30 := 0; fi30 < 13; fi30++ {
                if false {
                    v28 -= 128
                }
                for fi31 := 0; fi31 < 15; fi31++ {
                    v3 = v3
                    _ = fi31
                }
                _ = fi30
            }
        }
    }
    if ((((0.0 + 0.0) - 1.5) / 1.0) > 0.1) {
        v4[56] = "hello"
        var v32 uint64 = 2
        fmt.Println(v3, v1, v2)
        _ = v32
    }
    switch "foo bar" {
    case "foo bar":
        switch v28 {
        case 128:
            for fi33 := 0; fi33 < 16; fi33++ {
                switch ("hello" + "a") {
                case "hello":
                    v2 = ((-(fi33)) >= (v28 + fi33))
                default:
                    for fi34 := 0; fi34 < 4; fi34++ {
                        v35 := []int{(^(v28)), fi34, v28, (v28 >> ((genFn0(genFn1(0, 613), 38, 0)) & 63))}
                        _ = fi34
                        _ = v35
                    }
                }
                _ = fi33
            }
        case -9223372036854775808:
            v28 -= (v28 ^ -16)
        }
    default:
        for fi36 := 0; fi36 < 20; fi36++ {
            for fi37 := 0; fi37 < 5; fi37++ {
                var v38 uint16 = 780
                _ = fi37
                _ = v38
            }
            for fi39 := 0; fi39 < 19; fi39++ {
                switch ("hello" + ("\n\t" + "\n\t")) {
                case "":
                    var v40 uint32 = 0
                    v2 = v2
                    var v41 uint32 = 1
                    _ = v40
                    _ = v41
                case "a":
                    var v42 int8 = 100
                    _ = v42
                }
                _ = fi39
            }
            for fi43 := 0; fi43 < 12; fi43++ {
                v1 -= (^(v1))
                _ = fi43
            }
            _ = fi36
        }
    }
    fmt.Println(v1)
    for fi44 := 0; fi44 < 19; fi44++ {
        var v45 uint8 = 255
        v28 += fi44
        _ = fi44
        _ = v45
    }
    v46 := []int{(v28 >> (((^(v1))) & 63)), 127, (v28 & v28), (v28 << ((genFn0(-84, -128, 0)) & 63))}
    var v47 int64 = 0
    v2 = (1e10 < float64(-2.25))
    var v48 uint = v1
    if v2 {
        fmt.Println(v48)
        _ = len(v4)
    } else {
        if (true && v2) {
            v49 := GenS1{F0: -9223372036854775808, F1: (v3 | (v3 + v3))}
            var v50 int64 = 0
            _ = v49
            _ = v50
        } else {
            v51 := GenS3{F0: v2, F1: (^(v48))}
            v47 += -9223372036854775808
            _ = v51
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
}
