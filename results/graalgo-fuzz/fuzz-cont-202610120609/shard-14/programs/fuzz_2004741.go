package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint64 }
type GenS2 struct { F0 int8; F1 string; F2 int64 }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS5 struct { F0 bool; F1 uint32 }

func main() {
    v1 := make(map[int]bool)
    v2 := []int{}
    v3 := []uint32{42, 616}
    var v4 uint32 = 2147483648
    v4 = v4
    fmt.Println(v4)
    v5 := GenS1{F0: false, F1: 194, F2: 18446744073709551615}
    if (255 == 2) {
        fmt.Println(v4)
        v4 -= 42
    } else {
        _ = len(v1)
    }
    if (0 != 1000) {
        fmt.Println(v4)
    } else {
        v6 := GenS1{F0: (42 != 42), F1: 203, F2: 18446744073709551615}
        v4 += (v4 | v4)
        switch 0 {
        case 32768:
            switch ("\n\t" + "\n\t") {
            case "":
                v7 := v1[282]
                _ = v7
            }
        }
        _ = v6
    }
    var v8 string = ""
    if false {
        var v9 int32 = -416
        var v10 uint8 = 128
        _ = v10
        _ = v9
    } else {
        switch 4294967295 {
        case 0:
            switch 42 {
            case -1:
                v11 := GenS2{F0: -35, F1: (((("\n\t" + v8) + "foo bar") + v8) + (v8 + v8)), F2: -857}
                _ = v11
            }
        case -187:
            switch 128 {
            case -2147483648:
                v4 -= (v4 >> ((42) & 63))
            default:
                delete(v1, -2147483648)
            }
        default:
            for fi12 := 0; fi12 < 18; fi12++ {
                var v13 float64 = 1e10
                _ = fi12
                _ = v13
            }
        }
        v4 -= 951
        v8 = v8
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        for fi15 := 0; fi15 < 3; fi15++ {
            if (!true) {
                v4 = (v4 | (v4 / ((v4) | 1)))
                if false {
                    fmt.Println(fi15)
                    fmt.Println(fi14)
                    var v16 uint = 2
                    _ = v16
                } else {
                    v8 = v8
                    v8 = v8
                    var v17 uint8 = 1
                    var v18 uint32 = v4
                    _ = v17
                    _ = v18
                }
                var v19 int32 = -1
                _ = v19
            }
            v20 := v1[fi15]
            _ = fi15
            _ = v20
        }
        var v21 float64 = (((1e10 * 1e10) - (float64(((3.14159 / (0.1 * (-2.25 - 0.0))) * 3.14159)) - (float64((-2.25 * float32(1.5))) / 1.0))) * (-2.25 / 1e10))
        _ = fi14
        _ = v21
    }
    if true {
        v22 := GenS3{F0: 1, F1: false}
        switch ((v8 + v8) + (("a" + "a") + v8)) {
        case "foo bar":
            for fi23 := 0; fi23 < 6; fi23++ {
                v24 := GenS2{F0: -1, F1: "a", F2: 371}
                fmt.Println(fi23, fi23)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi23
                _ = v24
            }
        default:
            switch ("foo bar" + v8) {
            case "":
                v8 = v8
            case "hello":
                if false {
                    if ((!true) || (!false)) {
                        v25 := GenS4{F0: 42, F1: -488, F2: 66}
                        for fi26 := 0; fi26 < 7; fi26++ {
                            v27 := GenS5{F0: true, F1: 42}
                            switch (fi26 | -17) {
                            case 42:
                                v4 += (^(v4))
                            case 32767:
                                v8 = v8
                            }
                            v4 = v4
                            _ = fi26
                            _ = v27
                        }
                        v28 := GenS2{F0: -104, F1: "a", F2: 42}
                        _ = v25
                        _ = v28
                    }
                    var v29 bool = true
                    var v30 float64 = (3.14159 / 0.1)
                    _ = v29
                    _ = v30
                } else {
                    for fi31 := 0; fi31 < 7; fi31++ {
                        v2 = append(v2, fi31)
                        _ = fi31
                    }
                    fmt.Println(v4, v8)
                    v4 -= (v4 & 164)
                }
            default:
                switch ("a" + v8) {
                case "hello":
                    switch "" {
                    case "":
                        v8 = v8
                    }
                case "foo bar":
                    for fi32 := 0; fi32 < 20; fi32++ {
                        delete(v1, (fi32 + (fi32 << ((uint(v4)) & 63))))
                        fmt.Println(fi32)
                        v33 := GenS5{F0: (1 >= -9223372036854775808), F1: 1}
                        _ = fi32
                        _ = v33
                    }
                case "\n\t":
                    v4 += v4
                }
            }
        }
        _ = v22
    }
    fmt.Println(v8)
    var v34 string = ("\n\t" + v8)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v8
}
