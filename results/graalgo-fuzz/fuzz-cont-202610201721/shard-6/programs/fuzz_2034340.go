package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 uint32; F2 int64 }
type GenS3 struct { F0 int8; F1 uint64; F2 string }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 int16; F1 uint16; F2 string }
type GenS6 struct { F0 int8; F1 bool; F2 string }
type GenS7 struct { F0 int32; F1 uint8; F2 int64 }

func main() {
    v1 := make(map[string]string)
    var v2 string = ("" + "")
    v3 := []int{0, 42}
    var v4 string = v2
    var v5 int8 = 118
    var v6 int8 = (v5 + v5)
    if (false || (true && (751 != 216))) {
        var v7 float32 = 1.5
        var v8 uint32 = 960
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 int = (fi9 >> ((755) & 63))
        _ = fi9
        _ = v10
    }
    switch -2147483648 {
    case 32768:
        var v11 uint16 = 255
        _ = v11
    case -2147483648:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if (65535 != 4294967295) {
        for fi12 := 0; fi12 < 8; fi12++ {
            if len(v3) > 0 {
                v3[0] = -2147483648
            }
            for fi13 := 0; fi13 < 13; fi13++ {
                fmt.Println(v6)
                switch (v4 + "a") {
                case "":
                    v14 := GenS1{F0: 65535}
                    _ = v14
                case "a":
                    var v15 int64 = -9223372036854775808
                    _ = v15
                }
                _ = fi13
            }
            var v16 float64 = (1e10 * 1e10)
            _ = fi12
            _ = v16
        }
        var v17 int64 = 9223372036854775807
        _ = v17
    } else {
        if true {
            var v18 int8 = (v5 << ((65535) & 63))
            switch "\n\t" {
            case "\n\t":
                v19 := GenS2{F0: "", F1: 2147483647, F2: 804}
                _ = v19
            case "a":
                if (!true) {
                    v5 -= (^(v5))
                } else {
                    var v20 int16 = 42
                    for fi21 := 0; fi21 < 12; fi21++ {
                        v22 := GenS3{F0: v18, F1: 18446744073709551615, F2: "\n\t"}
                        for fi23 := 0; fi23 < 6; fi23++ {
                            v2 = v2
                            _ = fi23
                        }
                        _ = fi21
                        _ = v22
                    }
                    v24 := GenS4{F0: 1, F1: 656}
                    _ = v20
                    _ = v24
                }
            case "foo bar":
                switch -350 {
                case 32767:
                    for fi25 := 0; fi25 < 11; fi25++ {
                        v26 := GenS5{F0: -32768, F1: 1, F2: v4}
                        var v27 bool = false
                        if true {
                            var v28 int8 = (v5 ^ -36)
                            _ = v28
                        }
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                case -146:
                    v6 = 127
                case -284:
                    switch 4294967295 {
                    case -2147483648:
                        fmt.Println(v4)
                    default:
                        v5 -= (v6 * v6)
                    }
                default:
                    v29 := GenS6{F0: (-(v6)), F1: false, F2: v2}
                    _ = v29
                }
            }
            var v30 bool = (false && false)
            _ = v18
            _ = v30
        }
    }
    v2 = (v4 + v4)
    _ = len(v1)
    for fi31 := 0; fi31 < 14; fi31++ {
        for fi32 := 0; fi32 < 1; fi32++ {
            v33 := GenS5{F0: 849, F1: 255, F2: v4}
            _ = fi32
            _ = v33
        }
        fmt.Println(fi31, fi31, fi31)
        _ = fi31
    }
    delete(v1, v2)
    for fi34 := 0; fi34 < 17; fi34++ {
        v5 = v5
        if (255 > 65535) {
            v2 = v2
        }
        _ = fi34
    }
    var v35 uint32 = 4294967295
    if true {
        v36 := GenS7{F0: 2147483647, F1: 42, F2: -881}
        v3 = append(v3, 2147483647)
        var v37 uint32 = v35
        _ = v36
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
