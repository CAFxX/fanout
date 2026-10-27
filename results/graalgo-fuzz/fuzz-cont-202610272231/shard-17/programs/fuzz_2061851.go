package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 int64; F2 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 int; F2 int8 }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint; F1 string }
type GenS8 struct { F0 uint16; F1 int8 }
type GenS9 struct { F0 uint }

func main() {
    var v1 int16 = -723
    v2 := make(map[string]int)
    var v3 int8 = 1
    v4 := make(map[string]uint32)
    var v5 uint8 = 2
    var v6 uint16 = 317
    var v7 float64 = 3.14159
    for fi8 := 0; fi8 < 9; fi8++ {
        v5 = (v5 - v5)
        for fi9 := 0; fi9 < 14; fi9++ {
            v3 += v3
            var v10 int16 = int16(v5)
            v5 -= 0
            switch (fi9 & fi8) {
            case -9223372036854775808:
                v4["foo bar"] = uint32(v5)
            case 0:
                v11 := GenS1{F0: true}
                _ = v11
            case 127:
                v12 := v4[("hello" + "\n\t")]
                _ = v12
            default:
                var v13 int32 = 2147483647
                _ = v13
            }
            v6 = 65535
            if true {
                v14 := []int{}
                _ = v14
            } else {
                if (2147483647 > uint32(v5)) {
                    v7 -= ((float64(0.1) * ((v7 / (1e10 + float64(0.1))) - float64((0.1 - 1.5)))) - (1e10 - (0.1 - v7)))
                    var v15 int = (^(fi9))
                    _ = v15
                } else {
                    v16 := v2[""]
                    switch (v16 % (((fi9 - 81)) | 1)) {
                    case 0:
                        v17 := GenS2{F0: (^(v10)), F1: 396, F2: ("a" + "foo bar")}
                        _ = v17
                    case -1:
                        var v18 string = ("" + "a")
                        switch (fi9 - (-(v16))) {
                        case 9223372036854775807:
                            v19 := GenS1{F0: false}
                            _ = v19
                        case -579:
                            v3 = 0
                        case 4294967295:
                            fmt.Println(v10, v6)
                        }
                        _ = v18
                    default:
                        fmt.Println(v1, v5, v1)
                    }
                    _ = v16
                }
                var v20 uint16 = (v6 * (^(v6)))
                _ = v20
            }
            v21 := GenS2{F0: (v1 % ((v10) | 1)), F1: 1, F2: "foo bar"}
            _ = fi9
            _ = v10
            _ = v21
        }
        _ = fi8
    }
    var v22 float64 = v7
    v22 = (v22 + 3.14159)
    if (88 > 18446744073709551615) {
        if (((1.5 / (3.14159 / 1.0)) - 0.1) > 1.5) {
            delete(v4, ((("\n\t" + "a") + "") + ((("\n\t" + ("foo bar" + ("" + "\n\t"))) + "") + "a")))
            v23 := GenS2{F0: (v1 % ((int16(v5)) | 1)), F1: 0, F2: "\n\t"}
            switch -9223372036854775808 {
            case 54:
                var v24 uint16 = 32767
                _ = v24
            case 0:
                fmt.Println(v7)
            case 32767:
                for fi25 := 0; fi25 < 13; fi25++ {
                    for fi26 := 0; fi26 < 5; fi26++ {
                        v27 := GenS3{F0: true}
                        fmt.Println(v6)
                        v1 = 209
                        _ = fi26
                        _ = v27
                    }
                    v28 := GenS2{F0: 1000, F1: 0, F2: ("hello" + ("a" + ("foo bar" + "a")))}
                    _ = fi25
                    _ = v28
                }
            }
            _ = v23
        } else {
            v6 = v6
            if (((279 > 62) || (!false)) || (4294967295 < 4294967295)) {
                v29 := GenS3{F0: true}
                v6 -= v6
                for fi30 := 0; fi30 < 12; fi30++ {
                    v6 -= v6
                    _ = fi30
                }
                _ = v29
            }
        }
        var v31 uint64 = 0
        _ = v31
    }
    var v32 uint = 93
    for fi33 := 0; fi33 < 15; fi33++ {
        v34 := GenS4{F0: (^(v6)), F1: fi33, F2: (v3 % ((v3) | 1))}
        switch 42 {
        case 4294967295:
            v35 := []int{}
            _ = v35
        default:
            v3 += v3
        }
        if false {
            var v36 uint = (v32 & 96)
            var v37 uint32 = 4294967295
            _ = v36
            _ = v37
        }
        _ = fi33
        _ = v34
    }
    fmt.Println(v32, v22, v22)
    fmt.Println(v22, v7, v5)
    switch 127 {
    case 2147483647:
        v38 := GenS5{F0: 2, F1: v32}
        _ = v38
    default:
        for fi39 := 0; fi39 < 6; fi39++ {
            if false {
                v3 = v3
            }
            switch "hello" {
            case "":
                if true {
                    v40 := GenS6{F0: 683}
                    v41 := []int{(fi39 - fi39), fi39, fi39}
                    _ = v40
                    _ = v41
                }
            default:
                if (!false) {
                    if (-1 >= 9223372036854775807) {
                        v32 -= (v32 % ((v32) | 1))
                    }
                }
            }
            v42 := GenS3{F0: ((^(v5)) == v5)}
            _ = fi39
            _ = v42
        }
    }
    for fi43 := 0; fi43 < 9; fi43++ {
        v44 := v4["hello"]
        if true {
            for fi45 := 0; fi45 < 3; fi45++ {
                var v46 int8 = v3
                if false {
                    var v47 float64 = v7
                    var v48 int = fi45
                    if (((v44 & (v44 >> ((v32) & 63))) <= v44) || false) {
                        if true {
                            fmt.Println(v7)
                            v49 := GenS7{F0: (^(v32)), F1: "\n\t"}
                            _ = v49
                        }
                        v50 := GenS8{F0: (v6 ^ (v6 * 255)), F1: (v3 | v46)}
                        _ = v50
                    }
                    _ = v47
                    _ = v48
                }
                v6 = v6
                _ = fi45
                _ = v46
            }
        }
        v51 := GenS9{F0: 0}
        _ = fi43
        _ = v44
        _ = v51
    }
    fmt.Println(v3, v7, v5)
    if false {
        v52 := []int{0}
        _ = v52
    }
    fmt.Println(v6, v3)
    fmt.Println(v6)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
