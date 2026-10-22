package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int64; F1 uint }
func genFn0(p0 string, p1 int32, p2 int) bool {
    v7 := []int{}
    for fi8 := 0; fi8 < 17; fi8++ {
        fmt.Println(p1, p1, fi8)
        var v9 int32 = 0
        _ = fi8
        _ = v9
    }
    p1 = p1
    p2 += p2
    var v10 int = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    return false
}

func main() {
    var v1 uint8 = 1
    v2 := []uint32{65535, 4294967295}
    var v3 uint32 = 2147483648
    v4 := []int{-138, -9223372036854775808, -961}
    v5 := make(map[uint32]bool)
    var v6 int32 = 1
    var v11 int = -508
    fmt.Println(v1, v1, v6)
    for fi12 := 0; fi12 < 11; fi12++ {
        v13 := GenS1{F0: (v6 / ((v6) | 1)), F1: 1}
        var v14 uint32 = v3
        var v15 uint8 = (v1 >> ((255) & 63))
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    var v16 string = "hello"
    for fi17 := 0; fi17 < 18; fi17++ {
        v18 := GenS1{F0: 0, F1: 0}
        _ = fi17
        _ = v18
    }
    v1 = v1
    var v19 uint32 = (^(v3))
    var v20 uint32 = (v19 % (((v19 * (v19 + 2147483648))) | 1))
    for fi21 := 0; fi21 < 9; fi21++ {
        switch v11 {
        case -2147483648:
            switch (v16 + v16) {
            case "foo bar":
                var v22 uint32 = v20
                _ = v22
            default:
                var v23 bool = (-218 == -179)
                _ = v23
            }
        case 0:
            if (false || true) {
                for fi24 := 0; fi24 < 6; fi24++ {
                    v25 := GenS2{F0: -308, F1: 682}
                    var v26 uint32 = v20
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                v27 := GenS1{F0: v6, F1: 127}
                _ = v27
            } else {
                if false {
                    for fi28 := 0; fi28 < 1; fi28++ {
                        v20 = v19
                        v19 -= (v3 / ((v3) | 1))
                        v29 := GenS2{F0: -1, F1: 18446744073709551615}
                        _ = fi28
                        _ = v29
                    }
                    var v30 uint32 = (v3 << ((2) & 63))
                    var v31 int64 = -1
                    _ = v30
                    _ = v31
                } else {
                    for fi32 := 0; fi32 < 4; fi32++ {
                        var v33 uint16 = 65535
                        v6 = (v6 & 2147483647)
                        _ = fi32
                        _ = v33
                    }
                    if true {
                        for fi34 := 0; fi34 < 9; fi34++ {
                            if false {
                                if (!false) {
                                    var v35 uint8 = 0
                                    v36 := GenS1{F0: 42, F1: 0}
                                    fmt.Println(v20, v1)
                                    _ = v35
                                    _ = v36
                                }
                                var v37 uint64 = 1
                                if true {
                                    var v38 int64 = -368
                                    v38 -= 1
                                    if len(v2) > 0 {
                                        v2[0] = (v3 >> ((189) & 63))
                                    }
                                    _ = v38
                                } else {
                                    v16 = ("" + v16)
                                }
                                switch v11 {
                                case 619:
                                    if true {
                                        switch 4294967295 {
                                        case -1:
                                            var v39 bool = true
                                            _ = v39
                                        case 9223372036854775807:
                                            if ((-(fi34)) >= fi21) {
                                                v37 -= (^(v37))
                                                fmt.Println(v3)
                                                var v40 int16 = 1
                                                _ = v40
                                            }
                                        case -746:
                                            fmt.Println(v3, v3, v3)
                                        default:
                                            var v41 int16 = 276
                                            _ = v41
                                        }
                                        v16 = ((v16 + v16) + (v16 + v16))
                                        switch (fi34 << ((48) & 63)) {
                                        case -687:
                                            switch (v11 & fi21) {
                                            case 4294967295:
                                                switch "foo bar" {
                                                case "\n\t":
                                                    v4 = append(v4, v11)
                                                }
                                            case -549:
                                                _ = len(v5)
                                            default:
                                                var v42 int8 = 127
                                                _ = v42
                                            }
                                        case -13:
                                            v19 += v3
                                        case -1:
                                            v3 += 65535
                                        }
                                    } else {
                                        fmt.Println(fi21)
                                    }
                                case 895:
                                    var v43 int16 = 1
                                    _ = v43
                                }
                                var v44 uint16 = 829
                                _ = v37
                                _ = v44
                            } else {
                                _ = len(v5)
                            }
                            _ = fi34
                        }
                        var v45 float64 = 0.0
                        _ = v45
                    }
                }
                v11 -= (v11 * -244)
                var v46 uint32 = 2
                _ = v46
            }
        case 9223372036854775807:
            fmt.Println(v3, v6)
        }
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
