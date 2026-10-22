package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint32 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 bool }
func genFn0() uint32 {
    v7 := []int{-2147483648}
    v8 := GenS1{F0: 754, F1: -1, F2: 65535}
    var v9 int = -2147483648
    v10 := GenS1{F0: 255, F1: -1, F2: 2147483647}
    _ = len(v7)
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 2
}
func genFn1(p0 int16, p1 uint32, p2 uint32) uint64 {
    for fi11 := 0; fi11 < 8; fi11++ {
        p1 += genFn0()
        _ = fi11
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        if (true && (false && (true && false))) {
            for fi13 := 0; fi13 < 18; fi13++ {
                var v14 int32 = -1
                _ = fi13
                _ = v14
            }
        } else {
            if ((p2 | p2) == p2) {
                var v15 float64 = (1.5 / 0.1)
                for fi16 := 0; fi16 < 19; fi16++ {
                    if (65535 < 182) {
                        v17 := GenS2{F0: 2147483647, F1: ("" + "")}
                        _ = v17
                    } else {
                        v18 := []int{(fi16 + (-(fi16))), fi12, fi12}
                        _ = v18
                    }
                    var v19 uint16 = 0
                    _ = fi16
                    _ = v19
                }
                if (!false) {
                    v20 := GenS1{F0: 255, F1: -74, F2: (p1 - (^(p1)))}
                    var v21 uint64 = 858
                    _ = v20
                    _ = v21
                }
                _ = v15
            } else {
                p0 = 1000
                fmt.Println(p1)
            }
        }
        _ = fi12
    }
    p1 -= (p2 + 504)
    var v22 float64 = (0.0 - 0.0)
    p1 -= (p2 & (^(p1)))
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    return 0
}

func main() {
    var v1 int16 = 1000
    var v2 float32 = 1e10
    v3 := []string{"foo bar", ("\n\t" + ("a" + ""))}
    var v4 int8 = 115
    v5 := make(map[string]string)
    v6 := []uint32{2, 65535, 4294967295, 42}
    for fi23 := 0; fi23 < 10; fi23++ {
        for fi24 := 0; fi24 < 5; fi24++ {
            if (0 != 0) {
                var v25 bool = (352 == 65535)
                var v26 int = (fi23 & (fi24 << ((4294967295) & 63)))
                _ = len(v3)
                switch "a" {
                case "foo bar":
                    fmt.Println(v2, fi24)
                case "a":
                    switch ((("hello" + ("foo bar" + ((("foo bar" + "\n\t") + "hello") + "\n\t"))) + "foo bar") + "foo bar") {
                    case "foo bar":
                        switch ("a" + ("hello" + "")) {
                        case "foo bar":
                            v2 = float32(-2.25)
                        default:
                            var v27 uint8 = 0
                            _ = v27
                        }
                    case "":
                        v2 = v2
                    case "a":
                        var v28 int16 = (-(v1))
                        _ = v28
                    default:
                        v25 = (true && ((1 == -1) && (42 >= 958)))
                    }
                default:
                    var v29 uint = 0
                    var v30 uint64 = 0
                    var v31 uint8 = 201
                    _ = v29
                    _ = v30
                    _ = v31
                }
                _ = v25
                _ = v26
            } else {
                fmt.Println(v2)
                v2 = v2
            }
            v2 -= float32(3.14159)
            fmt.Println(v4)
            _ = fi24
        }
        switch fi23 {
        case 4294967295:
            var v32 int8 = (^(v4))
            _ = v32
        case -33:
            v1 += v1
        case -9223372036854775808:
            for fi33 := 0; fi33 < 7; fi33++ {
                switch "a" {
                case "hello":
                    fmt.Println(v2)
                case "a":
                    fmt.Println(v2, fi33, fi23)
                case "":
                    if false {
                        fmt.Println(fi23, fi23)
                        v2 -= float32(0.1)
                    }
                default:
                    v1 += (^(v1))
                }
                _ = fi33
            }
        }
        for fi34 := 0; fi34 < 14; fi34++ {
            v35 := GenS2{F0: 2147483647, F1: "hello"}
            _ = fi34
            _ = v35
        }
        _ = fi23
    }
    for fi36 := 0; fi36 < 17; fi36++ {
        v1 -= (v1 % (((-(v1))) | 1))
        _ = fi36
    }
    v4 -= v4
    _ = len(v6)
    v37 := GenS3{F0: 21, F1: 127}
    v38 := GenS1{F0: 1, F1: -48, F2: 407}
    if (480 != 4294967295) {
        switch "hello" {
        case "foo bar":
            v39 := GenS2{F0: -2147483648, F1: ""}
            _ = v39
        }
    }
    switch "foo bar" {
    case "foo bar":
        v1 -= v1
    case "a":
        switch "a" {
        case "foo bar":
            var v40 uint32 = 65535
            _ = v40
        case "\n\t":
            for fi41 := 0; fi41 < 19; fi41++ {
                v2 = (v2 - 1.5)
                var v42 int64 = 1
                _ = fi41
                _ = v42
            }
        case "hello":
            switch "\n\t" {
            case "foo bar":
                for fi43 := 0; fi43 < 2; fi43++ {
                    for fi44 := 0; fi44 < 20; fi44++ {
                        v45 := GenS2{F0: 465, F1: ("\n\t" + "foo bar")}
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi44
                        _ = v45
                    }
                    var v46 uint8 = 128
                    _ = fi43
                    _ = v46
                }
            }
        }
    }
    var v47 string = "a"
    for fi48 := 0; fi48 < 5; fi48++ {
        var v49 bool = (false || false)
        v2 -= (v2 + (float32(float64((float32(3.14159) - (1.5 - 1e10)))) - float32(1e10)))
        if len(v6) > 0 {
            v6[0] = 760
        }
        _ = fi48
        _ = v49
    }
    v4 += -1
    for fi50 := 0; fi50 < 5; fi50++ {
        v51 := GenS4{F0: true}
        _ = fi50
        _ = v51
    }
    v1 -= v1
    v1 = v1
    v47 = "foo bar"
    v52 := GenS3{F0: 629, F1: 1}
    if len(v6) > 0 {
        v6[0] = 2
    }
    for fi53 := 0; fi53 < 7; fi53++ {
        for fi54 := 0; fi54 < 12; fi54++ {
            fmt.Println(fi53)
            var v55 int8 = 100
            var v56 string = v47
            _ = fi54
            _ = v55
            _ = v56
        }
        var v57 int64 = -9223372036854775808
        for fi58 := 0; fi58 < 8; fi58++ {
            v1 -= (v1 / (((v1 + 0)) | 1))
            fmt.Println(fi53)
            var v59 uint64 = 2
            _ = fi58
            _ = v59
        }
        _ = fi53
        _ = v57
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v47
    _ = v5
    _ = v52
    _ = v6
}
