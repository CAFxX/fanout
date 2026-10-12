package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 int64; F1 int; F2 int32 }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 int64; F1 uint8 }
func genFn0(p0 uint, p1 int16) uint {
    var v5 int16 = (p1 & 1)
    for fi6 := 0; fi6 < 14; fi6++ {
        v5 += v5
        _ = fi6
    }
    fmt.Println(v5, p1)
    _ = p0
    _ = p1
    _ = v5
    return p0
}
func genFn1(p0 int64, p1 int8) int16 {
    v7 := GenS1{F0: ""}
    v8 := []int{-9223372036854775808, 9223372036854775807, 2147483647}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 238
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = -1
    var v3 float64 = (1e10 * 1e10)
    var v4 float32 = 1e10
    if false {
        v3 = (v3 - 1.5)
        v4 = ((v4 + 3.14159) + (-2.25 - 1.5))
        v4 -= v4
    }
    if (!((((3.14159 - float32(v3)) <= v4) || (160 > 32768)) && (755 < 833))) {
        var v9 bool = (true || true)
        _ = v9
    }
    _ = len(v1)
    v2 = (v2 >> ((255) & 63))
    switch -267 {
    case -282:
        v10 := GenS2{F0: 0, F1: 2147483648}
        _ = v10
    case -738:
        if (!true) {
            v4 = 1e10
        }
    default:
        for fi11 := 0; fi11 < 17; fi11++ {
            var v12 int64 = v2
            _ = len(v1)
            if (!true) {
                switch (fi11 ^ (fi11 - fi11)) {
                case 9223372036854775807:
                    v3 = (v3 * v3)
                default:
                    var v13 uint8 = 2
                    _ = v13
                }
            }
            _ = fi11
            _ = v12
        }
    }
    v14 := GenS3{F0: v2, F1: 788, F2: -2147483648}
    v3 += float64((1.5 - v4))
    var v15 float32 = v4
    var v16 int16 = 1
    var v17 int8 = 42
    v18 := GenS1{F0: ("\n\t" + "a")}
    var v19 float64 = 0.0
    for fi20 := 0; fi20 < 3; fi20++ {
        if false {
            switch ("foo bar" + "a") {
            case "foo bar":
                if false {
                    if true {
                        for fi21 := 0; fi21 < 17; fi21++ {
                            v4 = 0.0
                            fmt.Println(fi21)
                            var v22 int8 = v17
                            var v23 int16 = (v16 << ((255) & 63))
                            _ = fi21
                            _ = v22
                            _ = v23
                        }
                        if (false && true) {
                            if false {
                                switch (("" + "") + "foo bar") {
                                case "hello":
                                    v17 -= v17
                                case "\n\t":
                                    var v24 float32 = (float32(1e10) / 1.5)
                                    _ = v24
                                case "a":
                                    _ = len(v1)
                                default:
                                    for fi25 := 0; fi25 < 12; fi25++ {
                                        var v26 int32 = 372
                                        var v27 int32 = (v26 | v26)
                                        _ = fi25
                                        _ = v26
                                        _ = v27
                                    }
                                }
                                if (true && (255 >= 194)) {
                                    for fi28 := 0; fi28 < 6; fi28++ {
                                        switch "foo bar" {
                                        case "\n\t":
                                            if (!false) {
                                                v17 += v17
                                                var v29 uint = 2
                                                if false {
                                                    v17 = v17
                                                    var v30 int = (fi20 & (fi20 * fi20))
                                                    if true {
                                                        _ = len(v1)
                                                        switch "a" {
                                                        case "\n\t":
                                                            v31 := GenS4{F0: -2147483648, F1: ""}
                                                            _ = v31
                                                        default:
                                                            fmt.Println(v29)
                                                        }
                                                    }
                                                    _ = v30
                                                }
                                                _ = v29
                                            }
                                        case "hello":
                                            v32 := GenS5{F0: v2, F1: 128}
                                            _ = v32
                                        case "":
                                            switch (fi28 * fi20) {
                                            case -9223372036854775808:
                                                v33 := []int{-9223372036854775808, fi28}
                                                _ = v33
                                            }
                                        }
                                        fmt.Println(v17)
                                        _ = fi28
                                    }
                                    v34 := GenS2{F0: (v16 ^ genFn1(1, (^(v17)))), F1: 4294967295}
                                    _ = v34
                                }
                            }
                            v35 := GenS3{F0: 9223372036854775807, F1: (fi20 - fi20), F2: -2147483648}
                            _ = v35
                        } else {
                            v36 := []int{}
                            _ = v36
                        }
                    }
                    for fi37 := 0; fi37 < 4; fi37++ {
                        var v38 uint64 = 18446744073709551615
                        _ = fi37
                        _ = v38
                    }
                }
            }
            if (!(v19 != float64(v4))) {
                v39 := GenS2{F0: genFn1(0, 0), F1: 65535}
                var v40 uint64 = 0
                _ = v39
                _ = v40
            }
        }
        var v41 uint = 255
        _ = fi20
        _ = v41
    }
    var v42 uint8 = 199
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v42
}
