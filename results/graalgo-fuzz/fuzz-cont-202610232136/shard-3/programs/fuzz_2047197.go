package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int16; F1 int; F2 int64 }
type GenS3 struct { F0 int32; F1 uint; F2 string }
type GenS4 struct { F0 uint16; F1 int }
func genFn0(p0 uint64, p1 int8) uint16 {
    p0 += p0
    var v7 float32 = (float32(0.1) / 0.1)
    v7 = float32((1e10 - float64(v7)))
    switch "hello" {
    case "a":
        var v8 uint16 = 1
        _ = v8
    case "hello":
        v9 := GenS1{F0: 18446744073709551615, F1: 1000}
        _ = v9
    case "\n\t":
        var v10 bool = false
        _ = v10
    default:
        for fi11 := 0; fi11 < 17; fi11++ {
            for fi12 := 0; fi12 < 13; fi12++ {
                v13 := []int{}
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return 383
}
func genFn1(p0 uint8, p1 int, p2 uint) uint64 {
    fmt.Println(p0)
    fmt.Println(p1, p0)
    if false {
        v14 := GenS1{F0: 796, F1: 0}
        _ = v14
    } else {
        p1 -= 0
        for fi15 := 0; fi15 < 2; fi15++ {
            switch (-(p1)) {
            case 786:
                for fi16 := 0; fi16 < 9; fi16++ {
                    if true {
                        switch (p1 & fi16) {
                        case -829:
                            switch fi15 {
                            case 0:
                                v17 := GenS2{F0: 42, F1: (fi15 ^ (-(fi15))), F2: 604}
                                if (!(((p2 | (p2 >> ((235) & 63))) != (p2 | 255)) && false)) {
                                    p2 = p2
                                }
                                v18 := GenS3{F0: -164, F1: (p2 << (((^(p2))) & 63)), F2: (("hello" + (("foo bar" + "hello") + "a")) + "hello")}
                                v19 := GenS2{F0: 20, F1: (fi16 + (fi16 * (-(fi16)))), F2: -696}
                                v20 := []int{}
                                _ = v17
                                _ = v18
                                _ = v19
                                _ = v20
                            case 32767:
                                var v21 int8 = 100
                                _ = v21
                            default:
                                p2 += p2
                            }
                        default:
                            if (!false) {
                                p2 = (p2 & (p2 - (p2 ^ p2)))
                                p2 = p2
                            } else {
                                fmt.Println(p2)
                            }
                        }
                        v22 := []int{}
                        if false {
                            _ = len(v22)
                            fmt.Println(p0, p2, fi16)
                            var v23 float64 = (float64(1e10) / 3.14159)
                            _ = v23
                        }
                        _ = v22
                    }
                    var v24 int8 = -78
                    _ = fi16
                    _ = v24
                }
            case 127:
                v25 := []int{32768, 9223372036854775807}
                _ = v25
            }
            _ = fi15
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 464
}

func main() {
    var v1 int8 = 42
    v2 := make(map[string]bool)
    var v3 float64 = 1e10
    var v4 uint32 = 2147483648
    var v5 uint64 = 0
    v6 := make(map[string]uint32)
    for fi26 := 0; fi26 < 16; fi26++ {
        v3 -= v3
        v27 := []int{146, 32767, fi26, fi26}
        _ = fi26
        _ = v27
    }
    v4 += (^(v4))
    var v28 uint32 = v4
    var v29 int8 = 1
    var v30 float32 = ((-2.25 + (((1.5 - (-2.25 + (-2.25 + float32((0.0 + 0.0))))) * float32(v3)) - 0.0)) * 3.14159)
    v31 := GenS4{F0: 32767, F1: -1}
    if (int16(v5) < 1) {
        for fi32 := 0; fi32 < 13; fi32++ {
            var v33 int64 = 9223372036854775807
            var v34 bool = true
            switch "foo bar" {
            case "hello":
                var v35 int64 = (v33 / ((v33) | 1))
                _ = v35
            }
            _ = fi32
            _ = v33
            _ = v34
        }
    } else {
        if (v3 >= float64((v30 / v30))) {
            for fi36 := 0; fi36 < 7; fi36++ {
                var v37 string = "\n\t"
                v38 := GenS4{F0: 1, F1: 42}
                _ = fi36
                _ = v37
                _ = v38
            }
            var v39 uint = 30
            _ = v39
        } else {
            var v40 uint16 = 768
            _ = v40
        }
        switch "a" {
        case "hello":
            if false {
                var v41 uint16 = uint16(v5)
                var v42 string = ("hello" + "\n\t")
                _ = v41
                _ = v42
            }
        case "a":
            fmt.Println(v30, v4, v4)
        case "":
            v43 := GenS2{F0: -32768, F1: 251, F2: 42}
            _ = v43
        }
    }
    fmt.Println(v3, v1)
    if false {
        for fi44 := 0; fi44 < 13; fi44++ {
            v45 := GenS3{F0: 231, F1: 345, F2: ("hello" + ("hello" + ("a" + ("foo bar" + "\n\t"))))}
            v46 := GenS2{F0: -32768, F1: (fi44 - (fi44 >> ((65535) & 63))), F2: -1}
            _ = fi44
            _ = v45
            _ = v46
        }
        var v47 int64 = 42
        var v48 float64 = (0.1 / (1e10 - (-2.25 - v3)))
        _ = v47
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
