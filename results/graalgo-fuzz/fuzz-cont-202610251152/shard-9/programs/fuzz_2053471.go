package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 bool; F2 uint8 }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 uint16; F1 int8; F2 string }
func genFn0(p0 int8) string {
    switch 127 {
    case 9223372036854775807:
        switch "\n\t" {
        case "hello":
            var v5 uint64 = 351
            _ = v5
        case "foo bar":
            for fi6 := 0; fi6 < 13; fi6++ {
                v7 := GenS1{F0: "a"}
                p0 -= p0
                switch "" {
                case "\n\t":
                    p0 = 100
                case "a":
                    v8 := GenS2{F0: 18446744073709551615}
                    _ = v8
                default:
                    var v9 uint8 = 31
                    _ = v9
                }
                _ = fi6
                _ = v7
            }
        default:
            var v10 int8 = p0
            _ = v10
        }
    case 2147483647:
        v11 := GenS1{F0: "a"}
        _ = v11
    case -9223372036854775808:
        switch ("\n\t" + ("" + "foo bar")) {
        case "foo bar":
            p0 = p0
        case "a":
            p0 = 100
        default:
            for fi12 := 0; fi12 < 19; fi12++ {
                p0 = p0
                if false {
                    v13 := []int{(fi12 % ((fi12) | 1)), fi12, 2147483647}
                    p0 = p0
                    var v14 float32 = (1e10 - 1e10)
                    _ = v13
                    _ = v14
                } else {
                    var v15 int8 = p0
                    _ = v15
                }
                _ = fi12
            }
        }
    default:
        var v16 int = -205
        _ = v16
    }
    v17 := []int{0, 32767, 9223372036854775807}
    _ = p0
    _ = v17
    return "a"
}
func genFn1(p0 uint8, p1 uint16) int {
    fmt.Println(p0)
    if false {
        var v18 float32 = ((1.5 - float32(3.14159)) * -2.25)
        for fi19 := 0; fi19 < 11; fi19++ {
            switch "a" {
            case "hello":
                var v20 uint64 = 1
                _ = v20
            case "foo bar":
                v21 := GenS2{F0: 42}
                _ = v21
            }
            p0 = (^(p0))
            switch fi19 {
            case -1:
                var v22 float64 = 0.1
                _ = v22
            case 42:
                v18 = float32(((0.1 / 3.14159) * (1.5 + float64(v18))))
            }
            _ = fi19
        }
        switch -552 {
        case 127:
            fmt.Println(p0)
        case 2147483647:
            for fi23 := 0; fi23 < 4; fi23++ {
                if (!true) {
                    var v24 int32 = 1
                    p1 = (p1 + (p1 + p1))
                    v25 := GenS3{F0: (!true), F1: true, F2: (p0 / (((^(p0))) | 1))}
                    _ = v24
                    _ = v25
                } else {
                    var v26 int16 = 1000
                    _ = v26
                }
                _ = fi23
            }
        }
        _ = v18
    }
    p0 += uint8(p1)
    var v27 int64 = 0
    _ = p0
    _ = p1
    _ = v27
    return 32768
}

func main() {
    v1 := []string{"a"}
    var v2 uint64 = 492
    v3 := make(map[uint32]bool)
    v4 := make(map[string]bool)
    fmt.Println(v2)
    v2 = v2
    v2 = 42
    if false {
        v2 += 4294967295
    } else {
        v2 -= 58
        for fi28 := 0; fi28 < 16; fi28++ {
            for fi29 := 0; fi29 < 4; fi29++ {
                v2 = (v2 * (v2 ^ 1))
                var v30 int16 = 936
                _ = fi29
                _ = v30
            }
            for fi31 := 0; fi31 < 6; fi31++ {
                if true {
                    v2 -= (v2 << ((42) & 63))
                }
                _ = len(v1)
                _ = fi31
            }
            for fi32 := 0; fi32 < 6; fi32++ {
                switch fi28 {
                case 128:
                    switch (fi32 << ((999) & 63)) {
                    case 127:
                        var v33 bool = (9223372036854775807 >= -931)
                        _ = v33
                    }
                case -9223372036854775808:
                    fmt.Println(fi32)
                case 2147483647:
                    v34 := GenS3{F0: ((^(v2)) == v2), F1: (0 <= 69), F2: 2}
                    _ = v34
                }
                _ = fi32
            }
            _ = fi28
        }
        for fi35 := 0; fi35 < 15; fi35++ {
            if ((1.5 <= -2.25) && false) {
                if false {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v3[0] = false
                }
                var v36 int64 = int64(v2)
                v36 -= 9223372036854775807
                _ = v36
            } else {
                delete(v3, 2)
                var v37 int = fi35
                var v38 int = (v37 & (fi35 | (v37 - (fi35 - -218))))
                _ = v37
                _ = v38
            }
            _ = fi35
        }
    }
    v2 = v2
    if ((42 <= 18446744073709551615) || true) {
        v2 = v2
        v39 := GenS4{F0: 241, F1: false}
        _ = v39
    }
    for fi40 := 0; fi40 < 19; fi40++ {
        for fi41 := 0; fi41 < 15; fi41++ {
            v2 = v2
            v2 += v2
            _ = fi41
        }
        _ = fi40
    }
    for fi42 := 0; fi42 < 2; fi42++ {
        if true {
            var v43 uint32 = 1
            v44 := GenS5{F0: 28, F1: 100, F2: ("foo bar" + "foo bar")}
            _ = v43
            _ = v44
        } else {
            var v45 int8 = 75
            var v46 uint64 = v2
            var v47 uint32 = 391
            _ = v45
            _ = v46
            _ = v47
        }
        _ = fi42
    }
    v2 += v2
    var v48 int8 = -26
    fmt.Println(v48)
    var v49 int16 = 1000
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
}
