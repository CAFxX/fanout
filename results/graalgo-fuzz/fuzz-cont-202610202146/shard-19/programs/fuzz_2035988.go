package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 int8; F2 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string; F1 bool }
type GenS6 struct { F0 string; F1 string }
type GenS7 struct { F0 uint8; F1 int8; F2 int8 }
func genFn0() uint8 {
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := GenS1{F0: -2147483648, F1: -1, F2: 2}
        fmt.Println(fi6)
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 2147483647, F1: 42, F2: 698}
    if (1.5 > (float64(3.14159) - 3.14159)) {
        var v9 int8 = 72
        switch "a" {
        case "foo bar":
            v9 += (v9 << ((581) & 63))
        case "":
            v10 := GenS2{F0: true}
            _ = v10
        }
        for fi11 := 0; fi11 < 16; fi11++ {
            fmt.Println(fi11, fi11)
            v9 -= (v9 & (v9 >> ((2) & 63)))
            _ = fi11
        }
        _ = v9
    } else {
        var v12 int = -2147483648
        if ("a" <= "") {
            v12 -= (v12 % ((-9223372036854775808) | 1))
            v13 := []int{v12, v12, (^(v12))}
            var v14 uint64 = 1
            _ = v13
            _ = v14
        }
        v15 := GenS1{F0: -27, F1: 127, F2: 492}
        _ = v12
        _ = v15
    }
    var v16 uint16 = 32767
    var v17 float64 = (-2.25 - (3.14159 / (float64(0.1) / 3.14159)))
    _ = v16
    _ = v17
    _ = v8
    return 0
}
func genFn1() uint64 {
    for fi18 := 0; fi18 < 12; fi18++ {
        v19 := GenS3{F0: 42, F1: 1, F2: 42}
        _ = fi18
        _ = v19
    }
    var v20 uint = 0
    v21 := []int{4294967295, 128, -2147483648, -9223372036854775808}
    var v22 uint16 = 1
    _ = v20
    _ = v21
    _ = v22
    return 2
}

func main() {
    var v1 uint = 65535
    var v2 uint32 = 42
    v3 := []string{("foo bar" + "foo bar")}
    var v4 float64 = -2.25
    var v5 int32 = -2147483648
    switch "foo bar" {
    case "a":
        for fi23 := 0; fi23 < 4; fi23++ {
            var v24 int32 = 42
            v25 := GenS4{F0: 558}
            v26 := GenS4{F0: 1}
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
    case "hello":
        v2 -= v2
    default:
        v4 += v4
    }
    var v27 int = 32768
    if ((1e10 + (3.14159 - (((0.1 + float32(v4)) - 1e10) * 1.5))) != ((0.0 * 1.5) / (3.14159 + (3.14159 * ((float32(v4) / 3.14159) / (((3.14159 + 3.14159) * 1e10) - float32((v4 * v4)))))))) {
        if true {
            for fi28 := 0; fi28 < 14; fi28++ {
                var v29 bool = (v2 != v2)
                _ = fi28
                _ = v29
            }
            for fi30 := 0; fi30 < 7; fi30++ {
                switch (v27 >> ((v1) & 63)) {
                case 2147483647:
                    if len(v3) > 0 {
                        v3[0] = ("hello" + ("hello" + (("a" + ("foo bar" + ("" + ("a" + "a")))) + "\n\t")))
                    }
                case 42:
                    fmt.Println(v4, v5)
                default:
                    v31 := GenS3{F0: 18446744073709551615, F1: 0, F2: (v1 ^ (v1 + v1))}
                    _ = v31
                }
                _ = len(v3)
                v2 = (v2 & 106)
                _ = fi30
            }
            switch -194 {
            case -1:
                var v32 int8 = -128
                _ = v32
            case 32767:
                for fi33 := 0; fi33 < 17; fi33++ {
                    v4 = v4
                    _ = fi33
                }
            }
        } else {
            var v34 uint16 = 2
            _ = v34
        }
        for fi35 := 0; fi35 < 19; fi35++ {
            var v36 uint32 = v2
            switch ("hello" + "\n\t") {
            case "a":
                var v37 uint16 = 0
                _ = v37
            case "":
                v38 := GenS5{F0: "\n\t", F1: false}
                _ = v38
            }
            _ = fi35
            _ = v36
        }
    }
    if false {
        if false {
            switch "a" {
            case "\n\t":
                fmt.Println(v27, v2)
            case "hello":
                var v39 int8 = 100
                _ = v39
            case "a":
                _ = len(v3)
            default:
                v40 := GenS6{F0: "foo bar", F1: "foo bar"}
                _ = v40
            }
            for fi41 := 0; fi41 < 15; fi41++ {
                if len(v3) > 0 {
                    v3[0] = (((("foo bar" + "\n\t") + ("\n\t" + (("foo bar" + "") + "a"))) + "hello") + "foo bar")
                }
                v42 := GenS2{F0: ((0.0 != 0.1) && false)}
                _ = fi41
                _ = v42
            }
            var v43 uint64 = 18446744073709551615
            _ = v43
        }
    }
    v4 += (1.5 / float64(-2.25))
    if len(v3) > 0 {
        v3[0] = "a"
    }
    if (!(4294967295 == 42)) {
        if (0 == 70) {
            v2 += 884
        }
        for fi44 := 0; fi44 < 6; fi44++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi44
        }
        v1 -= (v1 | 42)
    }
    switch -1 {
    case -1:
        fmt.Println(v2)
    case 32768:
        fmt.Println(v27, v1)
    case -219:
        var v45 uint16 = 32767
        _ = v45
    }
    v2 -= 4294967295
    v46 := GenS7{F0: 255, F1: 1, F2: 98}
    switch 42 {
    case 488:
        v2 += (v2 & (v2 & (v2 ^ (v2 * (v2 + v2)))))
    case 127:
        switch v27 {
        case 32767:
            fmt.Println(v1)
        case 2147483647:
            _ = len(v3)
        default:
            if len(v3) > 0 {
                v3[0] = (("a" + "foo bar") + "\n\t")
            }
        }
    default:
        v3 = append(v3, "a")
    }
    var v47 int16 = 0
    fmt.Println(v5, v27)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
}
