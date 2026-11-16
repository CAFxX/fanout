package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16; F1 int8; F2 int }
type GenS5 struct { F0 uint; F1 int8; F2 int32 }
func genFn0() bool {
    var v5 int8 = -8
    if (!(-2147483648 < -476)) {
        var v6 bool = (962 > 429)
        var v7 uint64 = 2
        _ = v6
        _ = v7
    }
    _ = v5
    return false
}
func genFn1() int16 {
    var v8 uint64 = 1
    switch 4294967295 {
    case 494:
        for fi9 := 0; fi9 < 7; fi9++ {
            fmt.Println(fi9, fi9)
            switch (^(fi9)) {
            case 472:
                for fi10 := 0; fi10 < 17; fi10++ {
                    var v11 int = fi10
                    if false {
                        var v12 int32 = 42
                        _ = v12
                    }
                    fmt.Println(fi9)
                    _ = fi10
                    _ = v11
                }
            case -929:
                var v13 int64 = -1
                _ = v13
            }
            _ = fi9
        }
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        v8 -= 322
        _ = fi14
    }
    v15 := []int{0, 32768, -626, -553}
    v16 := GenS1{F0: false, F1: 42, F2: 980}
    _ = v15
    _ = v16
    _ = v8
    return 979
}

func main() {
    var v1 uint16 = 42
    var v2 int = 1
    var v3 float64 = (((1e10 * ((0.1 - -2.25) * 3.14159)) / -2.25) - 0.0)
    v4 := make(map[string]uint32)
    fmt.Println(v1, v3)
    switch (v2 ^ 127) {
    case -1:
        v3 -= 1.5
    case 2147483647:
        v2 += (v2 & -126)
    }
    if (!true) {
        v1 = (v1 & v1)
        v1 = v1
    } else {
        v2 += -213
    }
    var v17 string = ""
    if ((-1 < 1) || (616 == (v1 % (((v1 * (v1 - v1))) | 1)))) {
        for fi18 := 0; fi18 < 17; fi18++ {
            v19 := GenS1{F0: (1 != 2147483647), F1: 1000, F2: (^(fi18))}
            v20 := GenS1{F0: ((v1 | v1) != (v1 | 65535)), F1: int16(fi18), F2: v2}
            switch fi18 {
            case 2147483647:
                var v21 float64 = 3.14159
                _ = v21
            case 127:
                var v22 int64 = -1
                _ = v22
            }
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    fmt.Println(v17, v17, v2)
    v1 = (v1 | v1)
    switch "\n\t" {
    case "a":
        var v23 uint = 18446744073709551615
        _ = v23
    }
    switch ("foo bar" + "a") {
    case "hello":
        if true {
            delete(v4, v17)
            for fi24 := 0; fi24 < 5; fi24++ {
                switch "hello" {
                case "hello":
                    var v25 int16 = 199
                    _ = v25
                default:
                    for fi26 := 0; fi26 < 18; fi26++ {
                        var v27 int64 = 266
                        if (!true) {
                            fmt.Println(v1, v3)
                            delete(v4, v17)
                        } else {
                            if true {
                                var v28 float64 = v3
                                fmt.Println(fi24)
                                fmt.Println(v28, v1)
                                if false {
                                    var v29 float32 = (((1e10 - 0.1) - (float32(v28) - (3.14159 + float32(1.5)))) + 0.1)
                                    v30 := GenS1{F0: (5 > -69), F1: 1000, F2: (v2 * (fi24 ^ fi26))}
                                    _ = v29
                                    _ = v30
                                }
                                _ = v28
                            }
                        }
                        _ = fi26
                        _ = v27
                    }
                }
                v31 := []int{311, (v2 * fi24)}
                _ = fi24
                _ = v31
            }
            for fi32 := 0; fi32 < 3; fi32++ {
                var v33 float32 = 1e10
                v34 := GenS2{F0: 100, F1: 121}
                var v35 int64 = 0
                _ = fi32
                _ = v33
                _ = v34
                _ = v35
            }
        }
    case "foo bar":
        v36 := GenS3{F0: 1000}
        _ = v36
    default:
        if (2 == 2) {
            v2 = (-(v2))
            v2 = (v2 | 0)
            if false {
                v3 = 1e10
                if true {
                    var v37 string = v17
                    fmt.Println(v17, v2, v37)
                    _ = v37
                } else {
                    fmt.Println(v17, v3)
                }
            }
        }
    }
    v38 := GenS4{F0: -105, F1: -128, F2: (v2 / ((v2) | 1))}
    for fi39 := 0; fi39 < 18; fi39++ {
        for fi40 := 0; fi40 < 4; fi40++ {
            switch -2147483648 {
            case -9223372036854775808:
                if false {
                    if (!true) {
                        fmt.Println(v3, v17, fi39)
                        var v41 bool = true
                        _ = v41
                    } else {
                        v1 += v1
                    }
                    v42 := GenS2{F0: 42, F1: 2}
                    _ = v42
                } else {
                    v2 = (fi39 >> ((4294967295) & 63))
                }
                var v43 uint32 = 682
                _ = v43
            case -376:
                v1 = (^(v1))
            default:
                fmt.Println(v17)
            }
            v44 := GenS5{F0: 2, F1: 42, F2: 819}
            _ = fi40
            _ = v44
        }
        _ = fi39
    }
    var v45 int8 = -14
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v45
}
