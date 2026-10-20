package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 bool }
type GenS2 struct { F0 int8; F1 bool; F2 bool }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 int32; F1 uint8; F2 int16 }

func main() {
    v1 := []int{706, 4294967295, -9223372036854775808, 1}
    var v2 uint64 = 2
    var v3 uint8 = 127
    for fi4 := 0; fi4 < 15; fi4++ {
        switch "hello" {
        case "":
            v5 := GenS1{F0: (-(fi4)), F1: 0, F2: false}
            _ = v5
        }
        v6 := GenS2{F0: -92, F1: (!false), F2: ((false && (false && (446 <= (v2 & v2)))) && false)}
        v1 = append(v1, fi4)
        _ = fi4
        _ = v6
    }
    if (68 == 42) {
        fmt.Println(v2, v3)
    } else {
        for fi7 := 0; fi7 < 2; fi7++ {
            for fi8 := 0; fi8 < 8; fi8++ {
                var v9 int64 = 427
                var v10 int32 = 0
                _ = fi8
                _ = v10
                _ = v9
            }
            v11 := GenS3{F0: 18446744073709551615, F1: 0}
            v3 = (v3 ^ uint8(v2))
            _ = fi7
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        if len(v1) > 0 {
            v1[0] = fi12
        }
        _ = fi12
    }
    switch ("\n\t" + "hello") {
    case "a":
        if (42 < (^(v3))) {
            v3 = v3
            var v13 uint = 18446744073709551615
            for fi14 := 0; fi14 < 7; fi14++ {
                v2 = v2
                _ = fi14
            }
            _ = v13
        }
    case "hello":
        var v15 bool = (42 <= 604)
        _ = v15
    default:
        v2 = v2
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        if (false && false) {
            var v17 int = (fi16 - fi16)
            _ = v17
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v18 := GenS1{F0: int(v2), F1: 0, F2: false}
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v18
        }
        _ = fi16
    }
    v19 := GenS2{F0: 100, F1: false, F2: (1 >= 42)}
    switch "a" {
    case "hello":
        v20 := GenS2{F0: -7, F1: false, F2: true}
        _ = v20
    case "\n\t":
        switch 382 {
        case -1:
            for fi21 := 0; fi21 < 1; fi21++ {
                var v22 int64 = 1
                v23 := GenS3{F0: 255, F1: 42}
                _ = fi21
                _ = v22
                _ = v23
            }
        case -415:
            switch "foo bar" {
            case "a":
                var v24 int8 = 127
                _ = v24
            default:
                v2 = (^(v2))
            }
        }
    default:
        v2 -= (v2 << ((4294967295) & 63))
    }
    if true {
        v2 = 4294967295
        fmt.Println(v3)
    }
    if (true && true) {
        if true {
            var v25 float64 = 1e10
            if true {
                var v26 string = "a"
                if ((v3 % ((127) | 1)) != (v3 | v3)) {
                    v2 += (v2 ^ 1)
                } else {
                    if (v26 <= "") {
                        v27 := GenS1{F0: 358, F1: 2147483647, F2: (3.14159 > 3.14159)}
                        fmt.Println(v26, v25)
                        v1 = append(v1, 875)
                        _ = v27
                    }
                    v25 -= v25
                    var v28 float32 = (-2.25 + (0.0 / 0.1))
                    _ = v28
                }
                var v29 string = ""
                _ = v26
                _ = v29
            } else {
                for fi30 := 0; fi30 < 14; fi30++ {
                    for fi31 := 0; fi31 < 20; fi31++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        fmt.Println(v25, v3, v25)
                        v25 += 3.14159
                        _ = fi31
                    }
                    _ = fi30
                }
            }
            _ = v25
        } else {
            for fi32 := 0; fi32 < 19; fi32++ {
                var v33 float64 = ((0.1 * (((1e10 * 3.14159) - 1e10) - 0.1)) / 1e10)
                for fi34 := 0; fi34 < 19; fi34++ {
                    var v35 bool = (255 < 609)
                    _ = fi34
                    _ = v35
                }
                _ = fi32
                _ = v33
            }
            if (false && ((0.1 * 3.14159) == 0.1)) {
                if len(v1) > 0 {
                    v1[0] = 128
                }
            } else {
                if (true || ("" <= ("\n\t" + "hello"))) {
                    if false {
                        var v36 int16 = 1000
                        _ = v36
                    }
                }
            }
        }
        if ((false && true) || (v3 < v3)) {
            var v37 float64 = (float64((float32((3.14159 - 3.14159)) * 1e10)) - 1e10)
            v1 = append(v1, -9223372036854775808)
            _ = len(v1)
            _ = v37
        } else {
            fmt.Println(v3)
        }
    } else {
        v3 = v3
        if (0.1 > 0.0) {
            fmt.Println(v2, v2)
            for fi38 := 0; fi38 < 12; fi38++ {
                fmt.Println(v2, v3)
                _ = fi38
            }
        }
        v39 := GenS2{F0: 100, F1: (!(!false)), F2: (!true)}
        _ = v39
    }
    if true {
        v2 -= (v2 % (((v2 | (v2 ^ (v2 & 18446744073709551615)))) | 1))
        if len(v1) > 0 {
            v1[0] = 9223372036854775807
        }
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if true {
            v40 := GenS4{F0: 0, F1: 128, F2: -844}
            fmt.Println(v3)
            fmt.Println(v3)
            _ = v40
        }
    }
    fmt.Println(v3, v3)
    if len(v1) > 0 {
        v1[0] = -2147483648
    }
    var v41 uint32 = 4294967295
    _ = len(v1)
    fmt.Println(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v41
}
