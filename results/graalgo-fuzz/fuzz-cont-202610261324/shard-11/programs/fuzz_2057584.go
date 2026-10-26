package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 uint8; F1 uint32; F2 int8 }
type GenS5 struct { F0 uint32; F1 int8 }
type GenS6 struct { F0 int32; F1 string }
type GenS7 struct { F0 uint16; F1 int; F2 int }
func genFn0(p0 int8) int32 {
    p0 = p0
    if (202 < 352) {
        switch "hello" {
        case "":
            p0 = (p0 >> ((820) & 63))
        default:
            var v4 int8 = (-(p0))
            _ = v4
        }
        for fi5 := 0; fi5 < 12; fi5++ {
            var v6 string = ((("foo bar" + "foo bar") + "hello") + "")
            v7 := GenS1{F0: 2147483647}
            if true {
                p0 -= -1
            } else {
                v8 := []int{(^(fi5))}
                _ = v8
            }
            _ = fi5
            _ = v6
            _ = v7
        }
    } else {
        v9 := []int{127, -1, 4294967295}
        switch "hello" {
        case "":
            var v10 int16 = -549
            _ = v10
        case "hello":
            if (42 != 233) {
                p0 -= 127
                switch -380 {
                case 213:
                    p0 += (p0 - (p0 >> ((2) & 63)))
                }
            }
        case "a":
            p0 = p0
        }
        switch "\n\t" {
        case "":
            if len(v9) > 0 {
                v9[0] = -9223372036854775808
            }
        }
        _ = v9
    }
    _ = p0
    return 42
}
func genFn1(p0 int32) int8 {
    v11 := GenS1{F0: genFn0(42)}
    p0 -= (p0 + p0)
    _ = p0
    _ = v11
    return 42
}

func main() {
    var v1 string = "foo bar"
    v2 := []string{}
    var v3 uint16 = 255
    for fi12 := 0; fi12 < 6; fi12++ {
        if (-2147483648 > -2147483648) {
            var v13 uint32 = 0
            _ = v13
        } else {
            v3 += 1
            v14 := GenS2{F0: 67}
            for fi15 := 0; fi15 < 2; fi15++ {
                var v16 float32 = 3.14159
                _ = fi15
                _ = v16
            }
            _ = v14
        }
        var v17 uint64 = 540
        v18 := GenS2{F0: 422}
        _ = fi12
        _ = v17
        _ = v18
    }
    _ = len(v2)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 -= v3
    for fi19 := 0; fi19 < 8; fi19++ {
        for fi20 := 0; fi20 < 6; fi20++ {
            switch (fi20 % (((fi20 ^ fi19)) | 1)) {
            case 127:
                v1 = (v1 + v1)
            default:
                var v21 int32 = 0
                _ = v21
            }
            var v22 string = "foo bar"
            var v23 bool = true
            _ = fi20
            _ = v22
            _ = v23
        }
        var v24 uint = 632
        _ = fi19
        _ = v24
    }
    var v25 int8 = -100
    for fi26 := 0; fi26 < 14; fi26++ {
        v1 = ((((((("a" + v1) + v1) + "a") + (v1 + "a")) + "foo bar") + v1) + "hello")
        _ = fi26
    }
    if (!(-1 <= -2147483648)) {
        v3 -= v3
        if (!true) {
            v27 := GenS3{F0: -1, F1: 42}
            var v28 int32 = 648
            _ = v27
            _ = v28
        } else {
            v3 = 0
        }
    } else {
        if (!false) {
            v2 = append(v2, (("foo bar" + v1) + ("\n\t" + v1)))
            var v29 int = 128
            for fi30 := 0; fi30 < 5; fi30++ {
                v31 := GenS4{F0: 127, F1: 65535, F2: genFn1(0)}
                _ = fi30
                _ = v31
            }
            _ = v29
        } else {
            v25 -= (v25 >> ((4294967295) & 63))
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v32 float64 = 3.14159
    if len(v2) > 0 {
        v2[0] = (v1 + v1)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v33 := GenS1{F0: genFn0((v25 % (((-(v25))) | 1)))}
    if (32767 == 595) {
        fmt.Println(v25, v3, v32)
        var v34 bool = (true || ((1.5 * ((3.14159 - ((3.14159 * (float32(v32) / 0.1)) + (-2.25 - 0.0))) + 1.5)) != (float32((1.5 - (v32 + ((-2.25 - v32) * 3.14159)))) / 1.0)))
        var v35 uint64 = 492
        _ = v34
        _ = v35
    } else {
        v36 := GenS5{F0: 65535, F1: (v25 ^ v25)}
        v37 := GenS6{F0: -973, F1: "a"}
        _ = v36
        _ = v37
    }
    var v38 uint8 = 91
    if (!true) {
        var v39 int8 = (v25 ^ (^(v25)))
        for fi40 := 0; fi40 < 11; fi40++ {
            v3 = v3
            v25 = v25
            _ = fi40
        }
        _ = v39
    } else {
        for fi41 := 0; fi41 < 15; fi41++ {
            v42 := GenS6{F0: genFn0((v25 * (v25 % ((genFn1(genFn0(int8(v38)))) | 1)))), F1: v1}
            v43 := GenS1{F0: genFn0(v25)}
            _ = fi41
            _ = v42
            _ = v43
        }
    }
    v44 := GenS7{F0: v3, F1: -2147483648, F2: -941}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v32
    _ = v33
    _ = v38
    _ = v44
}
