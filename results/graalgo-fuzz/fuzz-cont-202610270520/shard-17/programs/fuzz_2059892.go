package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 string; F1 uint8; F2 uint }
type GenS4 struct { F0 int32 }
func genFn0(p0 int16, p1 int32, p2 uint8) int32 {
    v5 := []int{1}
    var v6 int = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (p1 >> ((18446744073709551615) & 63))
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []int{-9223372036854775808, 32768}
    var v3 bool = (false || (127 >= 42))
    v4 := []int{-59}
    for fi7 := 0; fi7 < 3; fi7++ {
        switch (fi7 * fi7) {
        case 1:
            var v8 uint = 65535
            _ = v8
        case 2147483647:
            for fi9 := 0; fi9 < 4; fi9++ {
                v10 := GenS1{F0: "", F1: ("hello" + ""), F2: 0}
                if (46 == 644) {
                    v3 = v3
                }
                _ = fi9
                _ = v10
            }
        }
        _ = fi7
    }
    fmt.Println(v1, v1)
    v11 := GenS1{F0: "foo bar", F1: "\n\t", F2: 18446744073709551615}
    fmt.Println(v3, v3)
    if (-2147483648 < 280) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if (42 > 103) {
            v12 := GenS1{F0: "hello", F1: "foo bar", F2: 4294967295}
            _ = v12
        } else {
            v3 = v3
        }
        if (-694 >= -309) {
            switch -625 {
            case 32768:
                if v3 {
                    fmt.Println(v3, v3)
                } else {
                    switch ((("a" + "hello") + "a") + "a") {
                    case "foo bar":
                        v2 = append(v2, 458)
                    case "\n\t":
                        for fi13 := 0; fi13 < 19; fi13++ {
                            var v14 int64 = (v1 % (((v1 << ((255) & 63))) | 1))
                            v3 = ((^(fi13)) < (fi13 - (fi13 * 326)))
                            _ = fi13
                            _ = v14
                        }
                    default:
                        v15 := GenS1{F0: "a", F1: ("" + "a"), F2: 0}
                        _ = v15
                    }
                    v16 := GenS2{F0: "\n\t", F1: 1}
                    _ = v16
                }
            default:
                switch "a" {
                case "":
                    v1 = (v1 + 42)
                case "hello":
                    for fi17 := 0; fi17 < 9; fi17++ {
                        var v18 float32 = 0.1
                        var v19 uint32 = 4294967295
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                }
            }
            for fi20 := 0; fi20 < 1; fi20++ {
                if true {
                    v3 = (true && (v3 || (32768 <= 42)))
                }
                var v21 uint = 4294967295
                fmt.Println(v21)
                _ = fi20
                _ = v21
            }
            v22 := GenS1{F0: "a", F1: "a", F2: 739}
            _ = v22
        } else {
            for fi23 := 0; fi23 < 8; fi23++ {
                var v24 uint8 = 202
                _ = fi23
                _ = v24
            }
            if v3 {
                v25 := GenS1{F0: "\n\t", F1: "foo bar", F2: 1}
                _ = v25
            } else {
                v26 := GenS3{F0: ("" + "hello"), F1: 199, F2: 18446744073709551615}
                var v27 string = "\n\t"
                _ = v26
                _ = v27
            }
            for fi28 := 0; fi28 < 2; fi28++ {
                v29 := GenS2{F0: "a", F1: -1}
                _ = fi28
                _ = v29
            }
        }
    }
    if v3 {
        v30 := GenS4{F0: -2147483648}
        _ = v30
    } else {
        var v31 bool = (466 > 0)
        if len(v2) > 0 {
            v2[0] = 1
        }
        _ = v31
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch "foo bar" {
    case "a":
        var v32 uint64 = 16
        _ = v32
    case "":
        var v33 float64 = 3.14159
        _ = v33
    case "\n\t":
        fmt.Println(v3)
    default:
        v34 := GenS4{F0: -189}
        _ = v34
    }
    if (972 >= 65535) {
        switch 32768 {
        case -2147483648:
            _ = len(v2)
        }
        for fi35 := 0; fi35 < 7; fi35++ {
            v36 := GenS2{F0: "hello", F1: -774}
            var v37 uint32 = 65535
            var v38 float32 = float32(1e10)
            _ = fi35
            _ = v36
            _ = v37
            _ = v38
        }
        switch "\n\t" {
        case "hello":
            v3 = v3
        case "a":
            v39 := GenS3{F0: "a", F1: 207, F2: 2}
            _ = v39
        default:
            for fi40 := 0; fi40 < 16; fi40++ {
                fmt.Println(v1, v3, fi40)
                _ = fi40
            }
        }
    } else {
        var v41 uint16 = 252
        if (42 < 100) {
            if (false || (v3 && false)) {
                if len(v4) > 0 {
                    v4[0] = 161
                }
                var v42 uint32 = 0
                _ = v42
            } else {
                v43 := GenS4{F0: 42}
                if (v41 != 334) {
                    fmt.Println(v41)
                    var v44 uint32 = 4294967295
                    _ = v44
                } else {
                    var v45 int32 = -2147483648
                    var v46 int16 = 32767
                    _ = v45
                    _ = v46
                }
                _ = v43
            }
        }
        _ = v41
    }
    v3 = v3
    v1 -= (v1 & (^(v1)))
    v47 := GenS1{F0: "foo bar", F1: "foo bar", F2: 174}
    fmt.Println(v3)
    _ = len(v4)
    v48 := GenS3{F0: "", F1: 128, F2: 0}
    if len(v2) > 0 {
        v2[0] = 42
    }
    if (!((((float32(1e10) / -2.25) * 0.0) == (1e10 * -2.25)) || ((!v3) || (0 >= 32767)))) {
        var v49 uint32 = 0
        _ = v49
    } else {
        var v50 int8 = 0
        _ = v50
    }
    for fi51 := 0; fi51 < 7; fi51++ {
        v1 = 9223372036854775807
        _ = fi51
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
}
