package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 string; F1 int64; F2 int8 }
type GenS4 struct { F0 int16; F1 string }
func genFn0(p0 uint32) string {
    p0 = (^(p0))
    v5 := GenS1{F0: -454}
    _ = p0
    _ = v5
    return "\n\t"
}
func genFn1(p0 uint64, p1 uint32, p2 int16) int8 {
    p1 -= p1
    v6 := []int{42, 4294967295, 4294967295, 9223372036854775807}
    var v7 uint8 = 127
    var v8 uint = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return -1
}

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = 363
    var v3 uint16 = 32767
    v4 := []uint32{0, 4294967295, 640, 4294967295}
    var v9 uint8 = 153
    var v10 string = "\n\t"
    v11 := GenS1{F0: -1}
    v2 = (v2 - -259)
    if true {
        for fi12 := 0; fi12 < 4; fi12++ {
            var v13 uint64 = uint64(v9)
            if (false && false) {
                for fi14 := 0; fi14 < 15; fi14++ {
                    if false {
                        fmt.Println(v10)
                        v10 = v10
                    }
                    var v15 bool = ((v10 + "foo bar") == "a")
                    var v16 string = (v10 + "foo bar")
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                v17 := GenS1{F0: 42}
                v9 += v9
                _ = v17
            } else {
                fmt.Println(fi12, v2)
                var v18 uint8 = v9
                _ = v18
            }
            v2 = v2
            _ = fi12
            _ = v13
        }
        for fi19 := 0; fi19 < 3; fi19++ {
            var v20 uint = 65535
            v21 := GenS1{F0: 137}
            var v22 uint64 = 786
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
    } else {
        v23 := GenS1{F0: 9223372036854775807}
        v24 := GenS2{F0: uint32(v9), F1: (^(v3))}
        v10 = v10
        _ = v23
        _ = v24
    }
    v2 = v2
    for fi25 := 0; fi25 < 14; fi25++ {
        v2 = (v2 >> ((65535) & 63))
        v26 := GenS1{F0: 42}
        _ = fi25
        _ = v26
    }
    var v27 uint16 = (^(v3))
    if false {
        for fi28 := 0; fi28 < 10; fi28++ {
            v2 += v2
            _ = fi28
        }
    }
    v3 = v27
    v9 = (v9 & (v9 >> ((1) & 63)))
    v27 -= 65535
    var v29 uint8 = (v9 % ((v9) | 1))
    fmt.Println(v10, v27, v2)
    v30 := GenS2{F0: 49, F1: (v27 / ((v3) | 1))}
    v9 = v9
    if true {
        for fi31 := 0; fi31 < 18; fi31++ {
            for fi32 := 0; fi32 < 4; fi32++ {
                if false {
                    switch (fi32 << ((255) & 63)) {
                    case 127:
                        if false {
                            v33 := GenS1{F0: 1}
                            var v34 int16 = (v2 ^ -32768)
                            _ = len(v4)
                            _ = v33
                            _ = v34
                        } else {
                            fmt.Println(fi31, v10, v27)
                        }
                    case 1:
                        var v35 uint16 = 515
                        _ = v35
                    case 32767:
                        switch fi31 {
                        case 999:
                            if false {
                                v29 = v9
                            } else {
                                _ = len(v1)
                                if (!true) {
                                    v36 := GenS1{F0: int64(v2)}
                                    _ = v36
                                } else {
                                    v1[v10] = true
                                    var v37 uint8 = (v9 << ((42) & 63))
                                    switch -9223372036854775808 {
                                    case -920:
                                        var v38 uint8 = v37
                                        _ = v38
                                    default:
                                        v39 := GenS2{F0: 2, F1: (v27 >> ((18446744073709551615) & 63))}
                                        _ = v39
                                    }
                                    _ = v37
                                }
                                fmt.Println(v9, v9)
                            }
                        case 9:
                            var v40 float64 = float64((float32(0.1) * 0.0))
                            _ = v40
                        case -9223372036854775808:
                            _ = len(v4)
                        default:
                            var v41 int8 = 40
                            _ = v41
                        }
                    }
                    var v42 int8 = 100
                    var v43 uint16 = (v3 ^ v3)
                    _ = v42
                    _ = v43
                }
                _ = fi32
            }
            var v44 uint = 42
            _ = fi31
            _ = v44
        }
        v45 := GenS3{F0: (v10 + v10), F1: -1, F2: 52}
        _ = v45
    }
    v46 := GenS4{F0: (v2 | -32768), F1: "a"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v46
    _ = v9
}
