package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint64; F1 uint8; F2 int }
type GenS3 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 uint64; F1 bool }
type GenS6 struct { F0 uint16; F1 bool; F2 int8 }
type GenS7 struct { F0 uint32; F1 int32; F2 string }
func genFn0(p0 uint, p1 string, p2 uint) uint {
    p0 = p0
    p0 += p2
    p2 = p0
    v4 := []int{-700, 2147483647, 32768, -198}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return p0
}
func genFn1(p0 int32, p1 int16) uint64 {
    for fi5 := 0; fi5 < 18; fi5++ {
        if true {
            switch (fi5 / ((-96) | 1)) {
            case -1:
                var v6 uint16 = 255
                _ = v6
            default:
                var v7 uint8 = 13
                _ = v7
            }
        } else {
            var v8 int64 = 0
            v9 := GenS1{F0: (p1 % (((p1 - p1)) | 1)), F1: 160}
            for fi10 := 0; fi10 < 11; fi10++ {
                var v11 uint16 = 205
                v11 += 42
                _ = fi10
                _ = v11
            }
            _ = v8
            _ = v9
        }
        switch "" {
        case "":
            fmt.Println(fi5, fi5)
        default:
            p1 = -290
        }
        for fi12 := 0; fi12 < 14; fi12++ {
            if (!(867 > 65535)) {
                var v13 uint64 = 42
                v14 := []int{-222, fi12, (fi12 >> ((1) & 63))}
                v15 := GenS1{F0: p1, F1: 5}
                _ = v13
                _ = v14
                _ = v15
            }
            v16 := GenS1{F0: 32767, F1: 128}
            _ = fi12
            _ = v16
        }
        _ = fi5
    }
    v17 := GenS1{F0: (p1 & 0), F1: 1}
    v18 := GenS2{F0: 42, F1: 42, F2: -9223372036854775808}
    for fi19 := 0; fi19 < 11; fi19++ {
        switch (-(fi19)) {
        case -1:
            if true {
                v20 := []int{fi19, -373, fi19}
                for fi21 := 0; fi21 < 14; fi21++ {
                    var v22 uint32 = 1
                    _ = fi21
                    _ = v22
                }
                p1 -= p1
                _ = v20
            } else {
                p1 = p1
                switch (fi19 * fi19) {
                case 61:
                    switch "" {
                    case "\n\t":
                        var v23 bool = (((3.14159 - float64((1.5 - ((-2.25 + 0.1) - 3.14159)))) == (3.14159 * float64(1.5))) || true)
                        _ = v23
                    case "hello":
                        var v24 uint = 65535
                        _ = v24
                    case "foo bar":
                        v25 := GenS3{F0: 6, F1: 141, F2: 18446744073709551615}
                        _ = v25
                    }
                case -415:
                    p1 = (p1 + (^(p1)))
                case -862:
                    v26 := GenS4{F0: 127, F1: 2}
                    _ = v26
                default:
                    v27 := []int{(fi19 >> ((4294967295) & 63)), fi19}
                    _ = v27
                }
            }
        case 882:
            var v28 uint = 2
            _ = v28
        default:
            var v29 int64 = 29
            _ = v29
        }
        for fi30 := 0; fi30 < 8; fi30++ {
            v31 := GenS5{F0: 2, F1: (3.14159 <= ((1.5 + -2.25) - 1.5))}
            _ = fi30
            _ = v31
        }
        _ = fi19
    }
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return 213
}

func main() {
    var v1 int32 = 2147483647
    var v2 float64 = 0.1
    var v3 int64 = -1
    v32 := []int{431, 32768, -1}
    v2 += 1.5
    switch -451 {
    case 9223372036854775807:
        var v33 int64 = v3
        _ = v33
    default:
        if len(v32) > 0 {
            v32[0] = 4294967295
        }
    }
    switch ("" + "\n\t") {
    case "hello":
        v34 := GenS3{F0: 149, F1: 42, F2: 4294967295}
        _ = v34
    case "a":
        v35 := GenS6{F0: 42, F1: true, F2: 1}
        _ = v35
    case "foo bar":
        switch 32767 {
        case -77:
            v3 -= v3
        case 0:
            _ = len(v32)
        }
    }
    if len(v32) >= 2 {
        v32 = v32[1:len(v32)]
    }
    fmt.Println(v1, v3)
    if ((2 == 1) || true) {
        if false {
            v32 = append(v32, -1)
            for fi36 := 0; fi36 < 5; fi36++ {
                v37 := GenS6{F0: 0, F1: (v3 > v3), F2: -128}
                _ = fi36
                _ = v37
            }
            fmt.Println(v3)
        } else {
            if true {
                _ = len(v32)
                var v38 int64 = 0
                _ = v38
            }
            if len(v32) >= 2 {
                v32 = v32[1:len(v32)]
            }
        }
        v39 := GenS5{F0: 1, F1: true}
        if false {
            v40 := GenS7{F0: 0, F1: v1, F2: ""}
            v3 -= (v3 >> ((genFn0(704, "", 42)) & 63))
            _ = v40
        } else {
            v3 -= v3
            _ = len(v32)
            var v41 int = 1
            _ = v41
        }
        _ = v39
    } else {
        for fi42 := 0; fi42 < 6; fi42++ {
            v32 = append(v32, -2147483648)
            _ = fi42
        }
        v3 -= int64(v1)
    }
    v3 += (v3 % ((v3) | 1))
    var v43 int8 = 42
    switch ("foo bar" + "\n\t") {
    case "foo bar":
        v3 = (^(v3))
    default:
        for fi44 := 0; fi44 < 12; fi44++ {
            v43 = (^(v43))
            v32 = append(v32, fi44)
            v32 = append(v32, (fi44 ^ fi44))
            _ = fi44
        }
    }
    var v45 bool = false
    for fi46 := 0; fi46 < 16; fi46++ {
        if (255 == 255) {
            var v47 string = ""
            for fi48 := 0; fi48 < 13; fi48++ {
                v45 = v45
                _ = fi48
            }
            _ = v47
        }
        _ = fi46
    }
    fmt.Println(v2, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v43
    _ = v45
}
