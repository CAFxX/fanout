package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 string; F1 bool; F2 int32 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS8 struct { F0 string; F1 int64; F2 uint }
func genFn0(p0 int8) int32 {
    fmt.Println(p0)
    v5 := GenS1{F0: -1, F1: p0, F2: 606}
    p0 = (p0 << ((782) & 63))
    v6 := GenS1{F0: -682, F1: p0, F2: -370}
    v7 := []int{0, 32767, 128}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 uint32) uint8 {
    v8 := []int{-5, 128, 9223372036854775807, -9223372036854775808}
    fmt.Println(p0)
    if (0.1 != 1.5) {
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 uint32 = p0
            p0 += 1
            for fi11 := 0; fi11 < 9; fi11++ {
                switch ("hello" + ("hello" + "a")) {
                case "hello":
                    v12 := GenS1{F0: 42, F1: -128, F2: 1000}
                    _ = v12
                case "\n\t":
                    _ = len(v8)
                case "a":
                    v13 := GenS1{F0: -787, F1: 127, F2: 387}
                    _ = v13
                default:
                    v10 = 0
                }
                v14 := GenS1{F0: -1, F1: 96, F2: 1}
                if (-136 == 827) {
                    v15 := GenS2{F0: genFn0(-95)}
                    _ = v15
                } else {
                    p0 -= p0
                    fmt.Println(fi11, v10, fi9)
                    v10 = 1
                }
                _ = fi11
                _ = v14
            }
            _ = fi9
            _ = v10
        }
        for fi16 := 0; fi16 < 16; fi16++ {
            for fi17 := 0; fi17 < 1; fi17++ {
                var v18 string = ""
                v19 := GenS3{F0: false, F1: true}
                _ = len(v8)
                _ = fi17
                _ = v18
                _ = v19
            }
            _ = fi16
        }
        for fi20 := 0; fi20 < 18; fi20++ {
            var v21 uint64 = 0
            var v22 uint = 0
            _ = fi20
            _ = v21
            _ = v22
        }
    } else {
        v23 := GenS3{F0: ((65535 == 255) && (4294967295 == 4294967295)), F1: true}
        if len(v8) > 0 {
            v8[0] = -1
        }
        switch ("hello" + ("foo bar" + ("hello" + "hello"))) {
        case "foo bar":
            switch "a" {
            case "hello":
                p0 = p0
            }
        case "a":
            for fi24 := 0; fi24 < 7; fi24++ {
                v25 := GenS4{F0: genFn0(106), F1: 127}
                fmt.Println(fi24, fi24)
                _ = fi24
                _ = v25
            }
        }
        _ = v23
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        if true {
            for fi27 := 0; fi27 < 11; fi27++ {
                v28 := GenS5{F0: (("a" + ("hello" + "\n\t")) + "a"), F1: (255 < 65535), F2: genFn0(-1)}
                var v29 uint16 = 210
                p0 -= (p0 - (p0 - (p0 / (((p0 & p0)) | 1))))
                var v30 uint32 = (p0 + 463)
                switch "foo bar" {
                case "hello":
                    fmt.Println(v29)
                case "a":
                    var v31 uint32 = (p0 >> ((255) & 63))
                    _ = v31
                case "foo bar":
                    switch 128 {
                    case -2147483648:
                        switch (fi26 * fi27) {
                        case -400:
                            fmt.Println(v30)
                        case 2147483647:
                            v8 = append(v8, (fi26 * (fi27 - fi26)))
                        }
                    case 1:
                        var v32 int32 = -1
                        _ = v32
                    default:
                        if true {
                            v30 += v30
                        }
                    }
                }
                _ = fi27
                _ = v28
                _ = v29
                _ = v30
            }
        } else {
            var v33 string = "\n\t"
            v33 = (v33 + v33)
            if false {
                var v34 float32 = ((((3.14159 / 1.0) + 1e10) * -2.25) * 1e10)
                v35 := GenS4{F0: genFn0(0), F1: 2}
                for fi36 := 0; fi36 < 18; fi36++ {
                    var v37 uint = 255
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    _ = fi36
                    _ = v37
                }
                _ = v34
                _ = v35
            }
            _ = v33
        }
        _ = fi26
    }
    p0 = p0
    _ = p0
    _ = v8
    return 42
}

func main() {
    var v1 int64 = 42
    var v2 int8 = 42
    var v3 uint64 = 644
    v4 := make(map[string]string)
    v2 = v2
    var v38 uint32 = 42
    v3 -= v3
    v39 := []int{}
    fmt.Println(v38, v3)
    var v40 uint64 = v3
    switch "" {
    case "foo bar":
        switch "a" {
        case "a":
            var v41 float64 = -2.25
            _ = v41
        case "foo bar":
            var v42 uint64 = (v40 << ((0) & 63))
            _ = v42
        case "\n\t":
            if (!(false || false)) {
                v40 = uint64(v38)
                v3 -= v3
                v43 := GenS6{F0: v3}
                _ = v43
            }
        }
    default:
        v44 := GenS7{F0: uint8(v40), F1: (v3 - 4294967295), F2: 1}
        _ = v44
    }
    v1 -= (v1 << ((255) & 63))
    if (v1 == (^(v1))) {
        switch 957 {
        case 42:
            v45 := GenS5{F0: "a", F1: false, F2: int32(v40)}
            _ = v45
        case 32768:
            v2 -= v2
        }
        var v46 uint64 = v40
        if true {
            if len(v39) > 0 {
                v39[0] = 42
            }
            for fi47 := 0; fi47 < 12; fi47++ {
                v4["a"] = "foo bar"
                if (1.5 <= 0.1) {
                    v1 -= 42
                } else {
                    v3 -= v3
                    var v48 uint16 = 255
                    v40 += (^(v46))
                    _ = v48
                }
                _ = fi47
            }
            v49 := GenS8{F0: "hello", F1: (v1 * (v1 * v1)), F2: 42}
            _ = v49
        } else {
            for fi50 := 0; fi50 < 3; fi50++ {
                var v51 int64 = 640
                _ = fi50
                _ = v51
            }
            for fi52 := 0; fi52 < 6; fi52++ {
                v53 := GenS4{F0: -48, F1: 17}
                var v54 uint64 = 42
                v46 = (v40 | (^(v54)))
                _ = fi52
                _ = v53
                _ = v54
            }
            fmt.Println(v46)
        }
        _ = v46
    }
    switch ("hello" + ("a" + "foo bar")) {
    case "\n\t":
        if (!((v2 > (v2 - v2)) || true)) {
            if ((true || false) || true) {
                for fi55 := 0; fi55 < 17; fi55++ {
                    for fi56 := 0; fi56 < 19; fi56++ {
                        v40 += 2
                        _ = len(v39)
                        var v57 uint8 = 214
                        _ = fi56
                        _ = v57
                    }
                    v38 -= (v38 / (((^(v38))) | 1))
                    _ = fi55
                }
                v1 -= v1
            } else {
                fmt.Println(v2, v40, v38)
                var v58 int32 = -174
                _ = v58
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
