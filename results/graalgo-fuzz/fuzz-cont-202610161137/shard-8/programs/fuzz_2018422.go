package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 bool; F1 uint64; F2 uint }

func main() {
    var v1 int32 = 42
    var v2 int64 = 42
    v3 := make(map[string]uint32)
    var v4 float64 = 3.14159
    var v5 uint16 = 971
    v6 := []int{-53, 32767, 128}
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 uint = 754
        fmt.Println(v4, v1)
        _ = len(v6)
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        v1 = v1
        v4 -= v4
        switch fi9 {
        case -142:
            v10 := GenS1{F0: v1, F1: v5}
            _ = v10
        case -303:
            if false {
                var v11 string = (((("foo bar" + "a") + "\n\t") + "") + (("hello" + "") + ("\n\t" + "")))
                _ = v11
            } else {
                for fi12 := 0; fi12 < 3; fi12++ {
                    fmt.Println(fi9)
                    _ = len(v6)
                    v1 -= (v1 / ((v1) | 1))
                    _ = fi12
                }
            }
        case 945:
            if true {
                if (0 > (v2 / ((v2) | 1))) {
                    for fi13 := 0; fi13 < 1; fi13++ {
                        switch "" {
                        case "a":
                            if (!true) {
                                v2 = int64(v1)
                                _ = len(v6)
                                var v14 int32 = int32(v5)
                                _ = v14
                            } else {
                                v2 = (v2 | (v2 - 0))
                                var v15 float64 = v4
                                _ = v15
                            }
                        }
                        fmt.Println(v5, fi9, v4)
                        _ = fi13
                    }
                    var v16 int8 = -8
                    var v17 int64 = (-(v2))
                    _ = v16
                    _ = v17
                }
            }
        default:
            for fi18 := 0; fi18 < 9; fi18++ {
                v5 -= v5
                _ = fi18
            }
        }
        _ = fi9
    }
    v2 = (v2 + (v2 << ((18446744073709551615) & 63)))
    fmt.Println(v4, v4)
    if true {
        switch "hello" {
        case "a":
            v4 += v4
        }
        switch 32767 {
        case 33:
            var v19 string = ("\n\t" + "foo bar")
            _ = v19
        case 4294967295:
            var v20 uint64 = 0
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 18; fi21++ {
        switch -2147483648 {
        case -9223372036854775808:
            var v22 int8 = 5
            _ = v22
        case 32767:
            v4 = v4
        case 0:
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        default:
            _ = len(v3)
        }
        _ = fi21
    }
    switch 127 {
    case 4294967295:
        for fi23 := 0; fi23 < 11; fi23++ {
            switch int(v1) {
            case 34:
                var v24 bool = (167 == 42)
                _ = v24
            default:
                v1 += 902
            }
            if false {
                for fi25 := 0; fi25 < 12; fi25++ {
                    v5 = (v5 << ((2) & 63))
                    v26 := GenS2{F0: true, F1: 621, F2: 201}
                    _ = len(v3)
                    _ = fi25
                    _ = v26
                }
            }
            _ = fi23
        }
    }
    if (v1 == -999) {
        if (127 <= 1) {
            var v27 uint8 = 255
            _ = v27
        }
        v3[("foo bar" + "\n\t")] = 1
    }
    var v28 int = -9223372036854775808
    var v29 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
