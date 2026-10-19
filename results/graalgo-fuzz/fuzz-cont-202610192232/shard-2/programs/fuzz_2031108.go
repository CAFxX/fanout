package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int8 }
type GenS2 struct { F0 uint; F1 int16; F2 uint64 }
type GenS3 struct { F0 uint; F1 int8; F2 string }
func genFn0(p0 uint8) int32 {
    p0 += p0
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := GenS1{F0: -563, F1: false, F2: -128}
        var v8 int32 = -2147483648
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    return 1
}

func main() {
    v1 := []uint32{0, 0, 110}
    var v2 float64 = 0.0
    var v3 uint = 65535
    v4 := []uint32{42}
    var v5 int64 = -936
    var v9 int8 = 100
    switch 42 {
    case -439:
        var v10 uint = (v3 << (((v3 / ((18446744073709551615) | 1))) & 63))
        _ = v10
    case -605:
        var v11 int = -111
        _ = v11
    }
    v12 := GenS1{F0: 0, F1: (!true), F2: v9}
    if ((v9 >= 42) && (0 >= 0)) {
        var v13 bool = (((("a" + (("foo bar" + "a") + "hello")) + "a") + ("a" + "a")) <= "a")
        var v14 int16 = -437
        _ = v13
        _ = v14
    } else {
        v15 := GenS1{F0: 1, F1: (!false), F2: -14}
        var v16 uint = 131
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 5; fi17++ {
        var v18 uint = 255
        v9 = 35
        for fi19 := 0; fi19 < 19; fi19++ {
            switch "" {
            case "":
                var v20 uint = v18
                _ = v20
            }
            _ = fi19
        }
        _ = fi17
        _ = v18
    }
    if (0 != 2) {
        if len(v1) > 0 {
            v1[0] = 0
        }
        var v21 int64 = v5
        var v22 uint32 = 395
        _ = v21
        _ = v22
    } else {
        switch 32767 {
        case 722:
            var v23 int16 = -32768
            _ = v23
        }
    }
    v9 = v9
    for fi24 := 0; fi24 < 19; fi24++ {
        if len(v1) > 0 {
            v1[0] = 678
        }
        v25 := GenS1{F0: -1, F1: true, F2: v9}
        _ = fi24
        _ = v25
    }
    var v26 bool = (true && (810 >= 32767))
    v27 := GenS2{F0: 919, F1: 1000, F2: 42}
    var v28 uint32 = 165
    v9 -= -128
    v4 = append(v4, v28)
    switch 293 {
    case 1:
        if (0 >= (v28 % ((4294967295) | 1))) {
            var v29 uint = 1
            var v30 uint32 = (^(v28))
            switch 0 {
            case 1:
                var v31 float32 = 1.5
                _ = v31
            }
            _ = v29
            _ = v30
        } else {
            for fi32 := 0; fi32 < 3; fi32++ {
                v33 := GenS2{F0: (v3 >> (((v3 >> ((v3) & 63))) & 63)), F1: 627, F2: 4294967295}
                var v34 uint16 = 372
                switch 32767 {
                case 9223372036854775807:
                    v35 := GenS1{F0: 1000, F1: false, F2: -57}
                    _ = v35
                case 2147483647:
                    v2 += (float64(float32(-2.25)) + v2)
                default:
                    switch "hello" {
                    case "":
                        v36 := GenS1{F0: -131, F1: (!false), F2: -1}
                        _ = v36
                    case "foo bar":
                        if v26 {
                            v37 := GenS2{F0: (v3 * v3), F1: -32768, F2: 1}
                            for fi38 := 0; fi38 < 9; fi38++ {
                                var v39 int32 = 2147483647
                                var v40 int16 = 1
                                v3 += uint(v28)
                                if false {
                                    var v41 uint32 = (v28 & (v28 >> ((v3) & 63)))
                                    var v42 uint8 = 127
                                    _ = v41
                                    _ = v42
                                } else {
                                    _ = len(v4)
                                }
                                _ = fi38
                                _ = v39
                                _ = v40
                            }
                            _ = v37
                        }
                    case "hello":
                        v34 = v34
                    default:
                        var v43 int16 = -32768
                        _ = v43
                    }
                }
                _ = fi32
                _ = v33
                _ = v34
            }
            switch "" {
            case "":
                v9 += (v9 % ((-128) | 1))
            default:
                v44 := GenS3{F0: v3, F1: v9, F2: ("" + "a")}
                _ = v44
            }
            fmt.Println(v26, v28, v26)
        }
    default:
        if (!false) {
            fmt.Println(v3, v2, v2)
            var v45 int8 = (v9 + (-(v9)))
            _ = v45
        } else {
            _ = len(v1)
            v5 -= (^(v5))
        }
    }
    _ = len(v1)
    var v46 uint8 = 128
    var v47 uint8 = v46
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v9
}
