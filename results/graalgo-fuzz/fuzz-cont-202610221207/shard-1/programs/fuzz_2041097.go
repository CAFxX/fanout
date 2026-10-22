package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 uint = 0
    var v2 int = -478
    var v3 uint32 = 0
    var v4 int16 = 849
    v5 := make(map[uint32]uint32)
    v6 := make(map[string]uint32)
    if false {
        switch (v2 + -2147483648) {
        case 32768:
            v7 := GenS1{F0: 9223372036854775807, F1: (93 != 2)}
            _ = v7
        default:
            v8 := GenS2{F0: (v3 ^ 65535)}
            _ = v8
        }
        v4 -= v4
        v9 := []int{32768, (v2 * v2)}
        _ = v9
    } else {
        v1 += (v1 / ((v1) | 1))
        var v10 int32 = -2147483648
        var v11 float64 = -2.25
        _ = v10
        _ = v11
    }
    if (v2 != (v2 | (v2 | (v2 % ((0) | 1))))) {
        fmt.Println(v3, v1, v3)
        if true {
            var v12 string = ""
            v2 -= (v2 | v2)
            switch 32767 {
            case -2147483648:
                var v13 uint16 = 175
                _ = v13
            case 309:
                v2 = v2
            default:
                if false {
                    v14 := []int{(v2 + (v2 - v2)), 1}
                    if false {
                        v2 = (v2 / ((v2) | 1))
                        v12 = ""
                    }
                    if true {
                        v1 -= v1
                        var v15 uint16 = 686
                        _ = v15
                    }
                    _ = v14
                }
            }
            _ = v12
        }
    } else {
        for fi16 := 0; fi16 < 13; fi16++ {
            switch v2 {
            case 2147483647:
                v17 := GenS2{F0: v3}
                _ = v17
            case -1:
                v5[(^(v3))] = 65535
            default:
                v18 := []int{}
                _ = v18
            }
            v19 := v6[""]
            _ = fi16
            _ = v19
        }
    }
    switch v2 {
    case -344:
        for fi20 := 0; fi20 < 9; fi20++ {
            for fi21 := 0; fi21 < 13; fi21++ {
                fmt.Println(v4)
                v1 -= v1
                if ("" == "") {
                    var v22 uint8 = 27
                    var v23 int32 = 2147483647
                    var v24 uint16 = 65535
                    if ((18446744073709551615 < 4294967295) || true) {
                        var v25 int = -110
                        _ = v25
                    } else {
                        switch 32768 {
                        case 4294967295:
                            v26 := GenS2{F0: v3}
                            _ = v26
                        case -99:
                            fmt.Println(v1)
                        case 2147483647:
                            var v27 uint = 255
                            _ = v27
                        default:
                            var v28 int = (v2 ^ (fi21 - int(v1)))
                            _ = v28
                        }
                        if false {
                            switch "" {
                            case "a":
                                v24 += 262
                            default:
                                v29 := GenS2{F0: 930}
                                _ = v29
                            }
                            v2 = v2
                        }
                        _ = len(v6)
                    }
                    _ = v22
                    _ = v23
                    _ = v24
                }
                _ = fi21
            }
            var v30 int32 = 1
            _ = fi20
            _ = v30
        }
    default:
        _ = len(v6)
    }
    for fi31 := 0; fi31 < 2; fi31++ {
        var v32 uint32 = (v3 << ((255) & 63))
        _ = fi31
        _ = v32
    }
    if (-2.25 == (0.0 * 0.1)) {
        var v33 float32 = (3.14159 + 0.1)
        switch v2 {
        case 32767:
            v34 := []int{32767}
            _ = v34
        case 1:
            var v35 string = ""
            _ = v35
        }
        fmt.Println(v3, v4)
        _ = v33
    } else {
        var v36 uint32 = v3
        _ = v36
    }
    v37 := GenS1{F0: -9223372036854775808, F1: (false || false)}
    v1 -= 503
    v38 := []int{32768}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
