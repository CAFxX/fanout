package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int8 }
type GenS2 struct { F0 int8; F1 uint; F2 string }
func genFn0(p0 bool, p1 int32) int8 {
    v5 := GenS1{F0: 42, F1: 2, F2: -128}
    v6 := []int{-9223372036854775808, 4294967295, 0, 32768}
    var v7 float64 = 1.5
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return -88
}
func genFn1(p0 int32, p1 uint32) uint8 {
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := []int{fi8, fi8}
        var v10 int16 = 1000
        _ = fi8
        _ = v10
        _ = v9
    }
    if false {
        if (("" + (("a" + "") + ("a" + ("foo bar" + "foo bar")))) <= "a") {
            var v11 int16 = -771
            v12 := GenS1{F0: v11, F1: p1, F2: 42}
            v13 := []int{}
            _ = v11
            _ = v12
            _ = v13
        } else {
            v14 := GenS1{F0: 42, F1: p1, F2: genFn0(false, -2147483648)}
            _ = v14
        }
    }
    _ = p0
    _ = p1
    return 8
}

func main() {
    var v1 int8 = 42
    v2 := []uint32{0, 65535, 2147483648, 65535}
    var v3 uint8 = 172
    v4 := make(map[uint32]string)
    switch "" {
    case "a":
        if (1e10 < float64(1e10)) {
            for fi15 := 0; fi15 < 2; fi15++ {
                v16 := GenS1{F0: 980, F1: 65535, F2: v1}
                var v17 uint8 = 127
                _ = fi15
                _ = v16
                _ = v17
            }
        } else {
            for fi18 := 0; fi18 < 2; fi18++ {
                v19 := GenS2{F0: (v1 / ((genFn0(false, -1)) | 1)), F1: 65535, F2: "hello"}
                delete(v4, 266)
                var v20 int32 = 0
                _ = fi18
                _ = v19
                _ = v20
            }
            fmt.Println(v3)
        }
    case "foo bar":
        v1 -= 42
    case "\n\t":
        for fi21 := 0; fi21 < 5; fi21++ {
            v1 += genFn0(true, 254)
            v22 := GenS1{F0: 1000, F1: 2, F2: v1}
            v23 := GenS2{F0: v1, F1: 18446744073709551615, F2: ("a" + "\n\t")}
            _ = fi21
            _ = v22
            _ = v23
        }
    }
    switch -1 {
    case -9223372036854775808:
        v3 = genFn1(-1, 2)
    case 127:
        var v24 float64 = -2.25
        _ = v24
    case -564:
        fmt.Println(v1, v1)
    default:
        if true {
            fmt.Println(v3, v1)
        }
    }
    var v25 int32 = -1
    if true {
        v1 = -126
        for fi26 := 0; fi26 < 10; fi26++ {
            v25 += -864
            v3 = (v3 * v3)
            var v27 int64 = 298
            _ = fi26
            _ = v27
        }
    }
    var v28 int = 1
    v25 = (v25 << ((0) & 63))
    if false {
        switch "\n\t" {
        case "hello":
            for fi29 := 0; fi29 < 20; fi29++ {
                v28 = 2147483647
                if (!true) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                } else {
                    var v30 string = "\n\t"
                    fmt.Println(fi29)
                    var v31 float32 = 0.1
                    _ = v30
                    _ = v31
                }
                _ = fi29
            }
        case "a":
            var v32 float64 = float64(float32(((1.5 - (((0.1 / float64((0.1 / ((-2.25 / (3.14159 * (1.5 / 1e10))) + 3.14159)))) / -2.25) / ((3.14159 + 0.1) - (((3.14159 - float64(3.14159)) - -2.25) + 1.5)))) - (3.14159 - 3.14159))))
            _ = v32
        default:
            if true {
                v28 += v28
                v1 = (v1 | v1)
            } else {
                v28 -= (-(v28))
                fmt.Println(v1, v3, v3)
            }
        }
        switch v28 {
        case -390:
            fmt.Println(v3, v28)
        case 127:
            for fi33 := 0; fi33 < 6; fi33++ {
                v28 = -9223372036854775808
                _ = fi33
            }
        case 4294967295:
            var v34 string = "hello"
            _ = v34
        }
    } else {
        switch "" {
        case "a":
            if (true && false) {
                if (1.5 > 1.5) {
                    v3 += (v3 ^ v3)
                }
                for fi35 := 0; fi35 < 13; fi35++ {
                    var v36 float32 = 0.0
                    _ = fi35
                    _ = v36
                }
                v3 = v3
            }
        default:
            var v37 bool = false
            _ = v37
        }
        _ = len(v4)
    }
    v28 = -595
    for fi38 := 0; fi38 < 9; fi38++ {
        v39 := v4[42]
        if (v25 <= v25) {
            var v40 uint64 = 2
            if true {
                switch fi38 {
                case -9223372036854775808:
                    switch ((v39 + ("" + (v39 + (v39 + v39)))) + v39) {
                    case "hello":
                        var v41 int8 = genFn0((((v39 + v39) + v39) < ("foo bar" + ((v39 + (v39 + v39)) + "hello"))), v25)
                        _ = v41
                    }
                default:
                    switch v39 {
                    case "":
                        var v42 int64 = 0
                        _ = v42
                    }
                }
                var v43 int64 = 989
                _ = v43
            } else {
                for fi44 := 0; fi44 < 12; fi44++ {
                    v25 += (-(v25))
                    _ = fi44
                }
                fmt.Println(v40, fi38, v1)
                _ = len(v4)
            }
            var v45 string = v39
            _ = v40
            _ = v45
        }
        _ = fi38
        _ = v39
    }
    v46 := GenS1{F0: -32768, F1: 42, F2: -1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v46
}
