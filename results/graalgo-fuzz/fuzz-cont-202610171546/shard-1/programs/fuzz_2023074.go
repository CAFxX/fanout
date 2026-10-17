package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 int8; F1 uint32; F2 string }
type GenS5 struct { F0 uint32; F1 uint }
type GenS6 struct { F0 int8 }
func genFn0(p0 bool, p1 int16) int8 {
    p1 = (p1 - p1)
    p1 = p1
    v4 := GenS1{F0: true, F1: 42}
    _ = p0
    _ = p1
    _ = v4
    return 1
}

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[string]uint32)
    v3 := []string{(("hello" + ("hello" + ("\n\t" + ""))) + "\n\t"), ("foo bar" + "hello")}
    var v5 uint32 = (v1 + 4294967295)
    v5 = v5
    v6 := GenS2{F0: 128, F1: -626}
    for fi7 := 0; fi7 < 3; fi7++ {
        switch -704 {
        case -2147483648:
            v8 := GenS3{F0: 9223372036854775807, F1: (fi7 | fi7)}
            _ = v8
        case 127:
            for fi9 := 0; fi9 < 12; fi9++ {
                v5 += v5
                if (1 != 1) {
                    v3 = append(v3, "\n\t")
                }
                _ = fi9
            }
        case 417:
            switch 0 {
            case 4294967295:
                v5 -= 748
            case 42:
                for fi10 := 0; fi10 < 15; fi10++ {
                    fmt.Println(fi10)
                    v1 -= v1
                    _ = fi10
                }
            }
        default:
            if true {
                var v11 uint64 = 4294967295
                v2[("" + "foo bar")] = (v5 ^ v1)
                v1 += 2147483648
                _ = v11
            } else {
                for fi12 := 0; fi12 < 9; fi12++ {
                    v13 := GenS3{F0: 208, F1: 185}
                    if false {
                        fmt.Println(fi7)
                        v14 := GenS2{F0: 127, F1: fi12}
                        _ = v14
                    }
                    var v15 int = -534
                    _ = fi12
                    _ = v13
                    _ = v15
                }
            }
        }
        _ = fi7
    }
    var v16 int32 = 42
    v17 := GenS4{F0: -19, F1: v1, F2: ""}
    _ = len(v2)
    for fi18 := 0; fi18 < 5; fi18++ {
        var v19 uint32 = 1
        switch (fi18 & (fi18 % (((fi18 << ((584) & 63))) | 1))) {
        case -116:
            for fi20 := 0; fi20 < 15; fi20++ {
                var v21 int = fi20
                var v22 uint64 = 18446744073709551615
                v23 := GenS5{F0: (v1 >> ((0) & 63)), F1: 0}
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
            }
        case -1:
            v24 := GenS6{F0: -128}
            _ = v24
        default:
            v25 := GenS3{F0: -1, F1: fi18}
            _ = v25
        }
        _ = fi18
        _ = v19
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        for fi27 := 0; fi27 < 9; fi27++ {
            v5 -= (v5 + (v5 | (v5 / (((v1 & v5)) | 1))))
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v28 int32 = -210
            _ = fi27
            _ = v28
        }
        v5 = v5
        v16 = v16
        _ = fi26
    }
    for fi29 := 0; fi29 < 6; fi29++ {
        switch (fi29 - fi29) {
        case 128:
            if false {
                if (-2147483648 != v16) {
                    switch "a" {
                    case "\n\t":
                        var v30 float32 = (3.14159 - float32(0.0))
                        _ = v30
                    default:
                        for fi31 := 0; fi31 < 19; fi31++ {
                            v1 += v5
                            _ = fi31
                        }
                    }
                    v3 = append(v3, "foo bar")
                } else {
                    v32 := GenS3{F0: 724, F1: (fi29 << ((745) & 63))}
                    fmt.Println(fi29, v5, fi29)
                    delete(v2, "")
                    _ = v32
                }
                fmt.Println(v1, v1, fi29)
                fmt.Println(v1, v1)
            }
        case 32768:
            _ = len(v2)
        case -675:
            for fi33 := 0; fi33 < 11; fi33++ {
                if false {
                    _ = len(v3)
                } else {
                    fmt.Println(fi33, v16)
                }
                _ = fi33
            }
        }
        for fi34 := 0; fi34 < 6; fi34++ {
            fmt.Println(v1)
            if false {
                var v35 bool = (18446744073709551615 > 316)
                _ = v35
            }
            _ = fi34
        }
        var v36 bool = (true && (!true))
        _ = fi29
        _ = v36
    }
    var v37 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v37
    _ = v5
    _ = v6
}
