package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 bool; F1 uint64; F2 uint }
type GenS4 struct { F0 uint64; F1 int; F2 uint64 }
type GenS5 struct { F0 uint64; F1 uint64; F2 int32 }
type GenS6 struct { F0 string; F1 int16; F2 uint8 }
func genFn0(p0 uint16, p1 uint64, p2 uint32) int8 {
    v4 := GenS1{F0: 0, F1: -1}
    v5 := []int{-813, 127, 2147483647, 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 string, p1 int16) uint8 {
    var v6 int64 = -792
    var v7 bool = false
    p1 = (p1 << ((545) & 63))
    var v8 int32 = -841
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 int64 = 42
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 128
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]int)
    var v3 string = ""
    v3 = v3
    for fi11 := 0; fi11 < 14; fi11++ {
        v3 = "a"
        v12 := GenS2{F0: 670, F1: -585}
        for fi13 := 0; fi13 < 10; fi13++ {
            v14 := GenS3{F0: (4294967295 == 88), F1: 18446744073709551615, F2: 2}
            v15 := GenS4{F0: 18446744073709551615, F1: (fi13 ^ (fi11 << ((318) & 63))), F2: 42}
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi11
        _ = v12
    }
    v16 := []int{4294967295, 32767, 0}
    switch v3 {
    case "hello":
        delete(v1, -171)
    case "a":
        v3 = ("a" + "\n\t")
    case "foo bar":
        var v17 int32 = -1
        _ = v17
    default:
        v3 = "foo bar"
    }
    v18 := v1[4294967295]
    var v19 float32 = -2.25
    v18 -= 2
    v20 := GenS4{F0: 18446744073709551615, F1: -9223372036854775808, F2: 4294967295}
    v18 -= (^(v18))
    for fi21 := 0; fi21 < 12; fi21++ {
        v22 := GenS5{F0: 663, F1: 42, F2: int32(v18)}
        v23 := GenS2{F0: 65535, F1: -1}
        for fi24 := 0; fi24 < 7; fi24++ {
            v19 = v19
            var v25 int32 = 0
            _ = fi24
            _ = v25
        }
        _ = fi21
        _ = v22
        _ = v23
    }
    for fi26 := 0; fi26 < 20; fi26++ {
        v27 := GenS5{F0: 42, F1: 0, F2: -2147483648}
        for fi28 := 0; fi28 < 8; fi28++ {
            v2[(fi28 + 42)] = -2147483648
            _ = fi28
        }
        _ = fi26
        _ = v27
    }
    if (true && true) {
        var v29 uint = 4294967295
        v29 = (^(v29))
        switch v3 {
        case "a":
            if true {
                v30 := GenS4{F0: 4294967295, F1: 1, F2: 4294967295}
                switch v3 {
                case "foo bar":
                    var v31 float32 = float32(3.14159)
                    _ = v31
                default:
                    var v32 int = 2147483647
                    _ = v32
                }
                v18 = v18
                _ = v30
            }
        case "":
            if (!true) {
                v1[0] = (v18 & (v18 * (v18 / ((65535) | 1))))
                switch 1 {
                case 0:
                    v18 -= (^(v18))
                default:
                    for fi33 := 0; fi33 < 12; fi33++ {
                        var v34 int16 = 345
                        fmt.Println(fi33, v34)
                        _ = fi33
                        _ = v34
                    }
                }
                v29 -= v29
            } else {
                for fi35 := 0; fi35 < 18; fi35++ {
                    for fi36 := 0; fi36 < 15; fi36++ {
                        v2[(fi36 | fi36)] = -9223372036854775808
                        _ = fi36
                    }
                    if false {
                        v29 = v29
                        switch fi35 {
                        case 32767:
                            switch v3 {
                            case "foo bar":
                                if len(v16) > 0 {
                                    v16[0] = fi35
                                }
                            default:
                                var v37 float64 = 1e10
                                _ = v37
                            }
                        case -2147483648:
                            var v38 string = ""
                            _ = v38
                        case 4294967295:
                            fmt.Println(v3, v3)
                        default:
                            if (v3 >= v3) {
                                if false {
                                    for fi39 := 0; fi39 < 9; fi39++ {
                                        var v40 float32 = (1e10 / v19)
                                        v41 := GenS3{F0: (498 >= 42), F1: 892, F2: 42}
                                        if ((!(1e10 == 1.5)) && false) {
                                            v42 := GenS6{F0: v3, F1: 870, F2: 42}
                                            _ = v42
                                        }
                                        _ = fi39
                                        _ = v40
                                        _ = v41
                                    }
                                    v3 = ("foo bar" + v3)
                                }
                            }
                        }
                    } else {
                        var v43 int16 = -1
                        fmt.Println(v18)
                        v19 = (v19 + v19)
                        _ = v43
                    }
                    v29 += (v29 ^ v29)
                    _ = fi35
                }
                v16 = append(v16, -137)
                v16 = append(v16, 9223372036854775807)
            }
        case "foo bar":
            fmt.Println(v18, v19, v18)
        default:
            v18 -= (v18 / (((v18 & v18)) | 1))
        }
        _ = v29
    } else {
        v19 = v19
        _ = len(v1)
        if false {
            v18 = 870
            fmt.Println(v18)
        }
    }
    v19 = v19
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
