package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 string; F1 int32; F2 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 uint; F1 int64; F2 uint32 }
type GenS6 struct { F0 int8; F1 uint32 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint; F1 bool }
func genFn0(p0 string, p1 uint) uint16 {
    var v4 int = -1
    switch (^(v4)) {
    case 933:
        for fi5 := 0; fi5 < 18; fi5++ {
            for fi6 := 0; fi6 < 5; fi6++ {
                v7 := GenS1{F0: -115, F1: true}
                _ = fi6
                _ = v7
            }
            v8 := []int{v4, (v4 + (v4 % ((v4) | 1))), (v4 % (((fi5 >> ((403) & 63))) | 1))}
            var v9 int32 = 0
            _ = fi5
            _ = v8
            _ = v9
        }
    case 0:
        for fi10 := 0; fi10 < 3; fi10++ {
            var v11 uint8 = 80
            var v12 int = (v4 >> (((p1 | (p1 | p1))) & 63))
            _ = fi10
            _ = v11
            _ = v12
        }
    default:
        var v13 int32 = 1
        _ = v13
    }
    p0 = p0
    if (128 != -9223372036854775808) {
        for fi14 := 0; fi14 < 19; fi14++ {
            switch p0 {
            case "a":
                if (-5 < -32768) {
                    var v15 int32 = -1
                    v16 := GenS2{F0: p0, F1: v15, F2: 1}
                    _ = v15
                    _ = v16
                }
            default:
                v17 := GenS1{F0: -1, F1: ((!false) && false)}
                _ = v17
            }
            var v18 uint8 = 2
            _ = fi14
            _ = v18
        }
        p1 += p1
    }
    _ = p0
    _ = p1
    _ = v4
    return 255
}

func main() {
    var v1 int = 1
    v2 := []uint32{984, 750}
    var v3 uint32 = 2147483648
    if (!false) {
        for fi19 := 0; fi19 < 4; fi19++ {
            for fi20 := 0; fi20 < 15; fi20++ {
                v21 := GenS3{F0: -551}
                _ = fi20
                _ = v21
            }
            for fi22 := 0; fi22 < 16; fi22++ {
                v2 = append(v2, v3)
                _ = fi22
            }
            _ = fi19
        }
    } else {
        var v23 uint8 = 194
        v23 -= v23
        _ = v23
    }
    v24 := GenS2{F0: (("a" + "hello") + ("" + "\n\t")), F1: 495, F2: 959}
    v25 := GenS4{F0: 9223372036854775807, F1: (0.0 >= 1.5)}
    v26 := GenS5{F0: 18446744073709551615, F1: -9223372036854775808, F2: 1}
    var v27 uint8 = 128
    v1 = v1
    if (false || false) {
        v28 := GenS5{F0: 255, F1: 1, F2: (v3 >> ((1) & 63))}
        v1 -= -682
        for fi29 := 0; fi29 < 4; fi29++ {
            switch "\n\t" {
            case "hello":
                var v30 uint16 = 895
                _ = v30
            }
            v2 = append(v2, 65535)
            _ = fi29
        }
        _ = v28
    }
    for fi31 := 0; fi31 < 10; fi31++ {
        var v32 float64 = 1e10
        _ = fi31
        _ = v32
    }
    for fi33 := 0; fi33 < 10; fi33++ {
        switch "a" {
        case "foo bar":
            switch (fi33 | (fi33 + (v1 << ((18446744073709551615) & 63)))) {
            case 9223372036854775807:
                var v34 int8 = -13
                _ = v34
            case 128:
                v27 = (v27 ^ (v27 << ((533) & 63)))
            case 4294967295:
                v35 := GenS5{F0: 65535, F1: 42, F2: 720}
                _ = v35
            }
        case "a":
            var v36 uint16 = 65535
            _ = v36
        case "":
            switch ("a" + "hello") {
            case "a":
                if true {
                    v1 = v1
                    v37 := GenS6{F0: -128, F1: v3}
                    _ = v37
                } else {
                    v3 += v3
                    var v38 uint32 = (v3 ^ (v3 + v3))
                    _ = v38
                }
            case "foo bar":
                v3 = v3
            default:
                fmt.Println(v1, v3)
            }
        }
        if true {
            var v39 int = fi33
            v1 = 32768
            _ = v39
        } else {
            v3 = 4294967295
            for fi40 := 0; fi40 < 2; fi40++ {
                var v41 uint16 = 697
                fmt.Println(fi40, v41, v41)
                v41 -= v41
                _ = fi40
                _ = v41
            }
        }
        fmt.Println(v3, v1)
        _ = fi33
    }
    v27 = (v27 >> ((1) & 63))
    if (!true) {
        switch ("hello" + "") {
        case "a":
            v27 -= v27
        case "hello":
            fmt.Println(v3, v1, v1)
        }
        v27 = v27
    }
    v3 += (v3 - (v3 ^ (v3 + (v3 >> ((2) & 63)))))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v42 float32 = -2.25
    v43 := GenS7{F0: 541}
    switch 2147483647 {
    case -9223372036854775808:
        v44 := GenS8{F0: 960, F1: false}
        _ = v44
    case 128:
        v2 = append(v2, 65535)
    case -1:
        v42 += v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v42
    _ = v43
}
