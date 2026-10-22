package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8; F1 int8; F2 int32 }
type GenS5 struct { F0 uint8; F1 uint32 }
func genFn0(p0 uint16, p1 int32) uint16 {
    fmt.Println(p0, p0)
    for fi5 := 0; fi5 < 17; fi5++ {
        v6 := []int{fi5, fi5, (fi5 << ((347) & 63)), fi5}
        v7 := GenS1{F0: (p0 * 42), F1: 0, F2: (0 != p1)}
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 int16 = 0
        p0 -= (^(p0))
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    return p0
}

func main() {
    var v1 float32 = (float32(1.5) / float32(float64(1.5)))
    v2 := make(map[string]bool)
    var v3 float64 = 3.14159
    var v4 uint64 = 2
    for fi10 := 0; fi10 < 12; fi10++ {
        delete(v2, ("\n\t" + ("\n\t" + "")))
        var v11 uint8 = 251
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        switch (^(fi12)) {
        case -1:
            v4 = (^(v4))
        }
        v2[""] = (255 != 0)
        _ = fi12
    }
    fmt.Println(v1, v1)
    var v13 int16 = 42
    for fi14 := 0; fi14 < 1; fi14++ {
        v1 -= ((1.5 - 0.0) * 1e10)
        switch fi14 {
        case 4294967295:
            if false {
                var v15 int16 = -124
                if (true || ((-(fi14)) <= fi14)) {
                    v1 += -2.25
                    fmt.Println(v1)
                    switch (fi14 + fi14) {
                    case 128:
                        v13 -= v15
                    case 55:
                        fmt.Println(v1, v1, v4)
                    default:
                        if (0 == 168) {
                            if (-9223372036854775808 != 947) {
                                if (true || (!(((((0.0 + (v3 - (1.5 - v3))) * ((v3 * 1e10) * 1.5)) != ((v3 * v3) - v3)) && ((32767 < v15) && (true || (v4 >= 791)))) && false))) {
                                    v16 := GenS1{F0: 65535, F1: 18446744073709551615, F2: false}
                                    _ = v16
                                }
                            }
                            v4 = 1
                        }
                    }
                } else {
                    v1 += v1
                    v2["hello"] = (true && false)
                }
                v15 = (v13 % ((32767) | 1))
                _ = v15
            }
        case 9223372036854775807:
            var v17 int8 = 127
            _ = v17
        }
        _ = fi14
    }
    v18 := GenS2{F0: 42}
    switch 120 {
    case -2147483648:
        if (genFn0(32767, 591) <= 65535) {
            v19 := GenS2{F0: v4}
            for fi20 := 0; fi20 < 11; fi20++ {
                v3 = v3
                _ = fi20
            }
            _ = v19
        }
    case -262:
        if true {
            v3 = 1e10
            for fi21 := 0; fi21 < 16; fi21++ {
                v22 := GenS3{F0: (v4 - (v4 - (v4 + (v4 ^ (v4 % (((v4 / (((v4 % (((v4 << ((401) & 63))) | 1))) | 1))) | 1))))))}
                var v23 uint32 = 966
                _ = fi21
                _ = v22
                _ = v23
            }
            v4 -= (v4 / ((2) | 1))
        } else {
            var v24 int16 = 1
            for fi25 := 0; fi25 < 12; fi25++ {
                v4 += v4
                _ = fi25
            }
            v1 = v1
            _ = v24
        }
    }
    for fi26 := 0; fi26 < 12; fi26++ {
        v27 := GenS2{F0: v4}
        v13 = v13
        _ = fi26
        _ = v27
    }
    v28 := GenS4{F0: -128, F1: 1, F2: 0}
    if true {
        v29 := []int{-9223372036854775808}
        var v30 uint8 = 128
        _ = v29
        _ = v30
    }
    switch "" {
    case "\n\t":
        for fi31 := 0; fi31 < 6; fi31++ {
            v1 = 1.5
            v1 -= 0.1
            switch "\n\t" {
            case "foo bar":
                v32 := []int{(fi31 * fi31), 128, 42, fi31}
                _ = v32
            case "hello":
                for fi33 := 0; fi33 < 10; fi33++ {
                    var v34 uint32 = 2
                    var v35 float64 = (0.0 * ((((((v3 / 1.5) - (v3 / (v3 * (v3 / v3)))) * ((v3 * v3) / 3.14159)) * v3) - (v3 * v3)) + float64(0.0)))
                    _ = fi33
                    _ = v34
                    _ = v35
                }
            case "":
                if (42 > 217) {
                    switch ("hello" + "foo bar") {
                    case "a":
                        v13 = v13
                    }
                } else {
                    v13 += v13
                    v36 := GenS5{F0: 2, F1: 371}
                    _ = v36
                }
            default:
                var v37 uint16 = genFn0(209, -2147483648)
                _ = v37
            }
            _ = fi31
        }
    case "foo bar":
        switch -277 {
        case -524:
            switch 887 {
            case 49:
                var v38 uint32 = 42
                _ = v38
            }
        }
    case "a":
        var v39 int64 = 96
        _ = v39
    }
    v40 := GenS3{F0: (v4 - v4)}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v40
}
