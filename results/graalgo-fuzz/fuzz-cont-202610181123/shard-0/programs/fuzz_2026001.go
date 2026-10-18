package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint16; F1 uint16 }
type GenS5 struct { F0 uint8; F1 int16; F2 int16 }
type GenS6 struct { F0 uint16; F1 uint16 }
type GenS7 struct { F0 int16; F1 uint16 }
func genFn0(p0 uint64) uint16 {
    fmt.Println(p0)
    for fi4 := 0; fi4 < 8; fi4++ {
        p0 = (^(p0))
        _ = fi4
    }
    fmt.Println(p0)
    _ = p0
    return 1
}
func genFn1() uint {
    var v5 int64 = -241
    v6 := GenS1{F0: 285}
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 uint = 18446744073709551615
        switch ((("\n\t" + "a") + "hello") + ("\n\t" + (("" + "\n\t") + "foo bar"))) {
        case "foo bar":
            v5 = v5
        case "a":
            var v9 string = "foo bar"
            _ = v9
        }
        var v10 string = "hello"
        _ = fi7
        _ = v10
        _ = v8
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        for fi12 := 0; fi12 < 16; fi12++ {
            var v13 int64 = 9223372036854775807
            fmt.Println(v5, fi12, fi11)
            _ = fi12
            _ = v13
        }
        v14 := GenS1{F0: 32767}
        _ = fi11
        _ = v14
    }
    var v15 int16 = -32768
    _ = v15
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 int = -2147483648
    var v2 uint8 = 241
    var v3 string = "hello"
    var v16 uint16 = 2
    v17 := []int{v1}
    var v18 uint64 = 0
    v18 -= (^(v18))
    v19 := GenS2{F0: 1, F1: v1}
    v2 += (v2 ^ 223)
    var v20 int64 = 145
    v21 := GenS3{F0: 128}
    v2 -= (v2 ^ (^(v2)))
    if (-2147483648 == (v1 << ((18446744073709551615) & 63))) {
        v2 = (v2 / (((v2 / ((v2) | 1))) | 1))
    }
    switch "hello" {
    case "hello":
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
    default:
        switch "hello" {
        case "":
            for fi22 := 0; fi22 < 19; fi22++ {
                var v23 int64 = (v20 | (v20 / ((v20) | 1)))
                for fi24 := 0; fi24 < 6; fi24++ {
                    switch "hello" {
                    case "a":
                        v17 = append(v17, fi22)
                    }
                    v25 := GenS2{F0: v23, F1: (v1 * (^(fi24)))}
                    switch fi22 {
                    case 4294967295:
                        var v26 uint8 = v2
                        if (genFn1() < 255) {
                            v27 := GenS4{F0: v16, F1: (v16 << ((4294967295) & 63))}
                            v1 -= -834
                            v17 = append(v17, v1)
                            var v28 uint16 = (v16 / ((v16) | 1))
                            _ = v27
                            _ = v28
                        } else {
                            switch fi22 {
                            case 127:
                                var v29 int8 = 12
                                _ = v29
                            case 2147483647:
                                v1 = (v1 + (v1 >> ((470) & 63)))
                            default:
                                _ = len(v17)
                            }
                            var v30 uint32 = 2147483648
                            v30 = (v30 % ((uint32(v18)) | 1))
                            switch "" {
                            case "\n\t":
                                var v31 uint64 = v18
                                _ = v31
                            }
                            _ = v30
                        }
                        var v32 int8 = -128
                        _ = v26
                        _ = v32
                    case -610:
                        var v33 uint8 = v2
                        _ = v33
                    case 320:
                        var v34 int8 = -1
                        _ = v34
                    }
                    _ = fi24
                    _ = v25
                }
                fmt.Println(fi22, fi22, fi22)
                _ = fi22
                _ = v23
            }
        default:
            if (-2.25 < (3.14159 - -2.25)) {
                for fi35 := 0; fi35 < 3; fi35++ {
                    for fi36 := 0; fi36 < 13; fi36++ {
                        switch v3 {
                        case "foo bar":
                            v3 = "a"
                        case "":
                            v37 := GenS1{F0: v16}
                            _ = v37
                        default:
                            v38 := GenS4{F0: (v16 & v16), F1: v16}
                            _ = v38
                        }
                        if (v16 != v16) {
                            v39 := GenS2{F0: (v20 - v20), F1: fi35}
                            if true {
                                v40 := GenS5{F0: (v2 | (v2 - (v2 >> ((uint(v2)) & 63)))), F1: 1, F2: -1}
                                _ = v40
                            } else {
                                v16 = (v16 | uint16(v18))
                                if len(v17) > 0 {
                                    v17[0] = fi35
                                }
                            }
                            var v41 uint16 = genFn0((^(v18)))
                            _ = v39
                            _ = v41
                        } else {
                            if len(v17) > 0 {
                                v17[0] = int(v2)
                            }
                        }
                        fmt.Println(v2)
                        _ = fi36
                    }
                    v42 := GenS6{F0: genFn0(0), F1: (v16 + v16)}
                    _ = fi35
                    _ = v42
                }
            } else {
                for fi43 := 0; fi43 < 15; fi43++ {
                    v44 := GenS7{F0: 648, F1: (v16 >> ((18446744073709551615) & 63))}
                    _ = fi43
                    _ = v44
                }
                var v45 int8 = -1
                var v46 float32 = float32(1e10)
                _ = v45
                _ = v46
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
