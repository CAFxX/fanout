package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS2 struct { F0 uint32; F1 int16; F2 int32 }
type GenS3 struct { F0 uint8; F1 int; F2 int32 }
type GenS4 struct { F0 bool; F1 int32; F2 uint8 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 int) uint {
    v6 := GenS1{F0: 2, F1: 127, F2: -69}
    if false {
        for fi7 := 0; fi7 < 20; fi7++ {
            var v8 float32 = 0.1
            _ = fi7
            _ = v8
        }
        fmt.Println(p0)
    } else {
        var v9 uint64 = 42
        for fi10 := 0; fi10 < 15; fi10++ {
            for fi11 := 0; fi11 < 1; fi11++ {
                v12 := []int{fi11, (fi10 % (((fi11 ^ (fi11 & -2147483648))) | 1)), 476}
                if ("\n\t" <= (("a" + (("a" + (("hello" + ("foo bar" + "a")) + "\n\t")) + "")) + ("foo bar" + "foo bar"))) {
                    var v13 uint64 = (v9 ^ (v9 >> ((0) & 63)))
                    _ = v13
                }
                var v14 float32 = 3.14159
                _ = fi11
                _ = v12
                _ = v14
            }
            _ = fi10
        }
        _ = v9
    }
    if (171 > 2147483647) {
        for fi15 := 0; fi15 < 4; fi15++ {
            var v16 uint16 = 42
            _ = fi15
            _ = v16
        }
        var v17 uint = 2
        _ = v17
    }
    _ = p0
    _ = v6
    return 42
}
func genFn1() int8 {
    for fi18 := 0; fi18 < 4; fi18++ {
        v19 := []int{(-(fi18))}
        if len(v19) > 0 {
            v19[0] = (fi18 | 32768)
        }
        _ = fi18
        _ = v19
    }
    var v20 float64 = 1e10
    v20 -= (1e10 * (3.14159 - 1e10))
    var v21 uint32 = 65535
    _ = v20
    _ = v21
    return 100
}

func main() {
    v1 := []uint32{42, 42, 1, 0}
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]uint32)
    var v4 int64 = 9223372036854775807
    var v5 bool = true
    fmt.Println(v4)
    switch 42 {
    case 163:
        switch -1 {
        case 75:
            switch "" {
            case "a":
                var v22 int64 = (v4 | (v4 - v4))
                _ = v22
            case "foo bar":
                v4 = v4
            case "":
                for fi23 := 0; fi23 < 7; fi23++ {
                    v24 := GenS2{F0: 344, F1: 195, F2: 0}
                    if false {
                        for fi25 := 0; fi25 < 6; fi25++ {
                            v5 = (546 != 255)
                            v5 = ((!(!v5)) && ((false || v5) && true))
                            var v26 int16 = 42
                            var v27 uint64 = 2
                            v4 = 1
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                        delete(v3, 0)
                        v4 = (v4 & -9223372036854775808)
                    }
                    v4 = (v4 ^ 42)
                    _ = fi23
                    _ = v24
                }
            default:
                fmt.Println(v5)
            }
        default:
            var v28 uint8 = 2
            _ = v28
        }
    default:
        var v29 float32 = (3.14159 + float32(0.0))
        _ = v29
    }
    v5 = false
    fmt.Println(v5)
    switch "" {
    case "hello":
        v30 := GenS2{F0: 42, F1: -1, F2: 0}
        _ = v30
    case "foo bar":
        if len(v1) > 0 {
            v1[0] = 1
        }
    default:
        fmt.Println(v4, v4)
    }
    v4 += (v4 * v4)
    v31 := GenS3{F0: 128, F1: 1, F2: -2147483648}
    fmt.Println(v4, v4)
    for fi32 := 0; fi32 < 2; fi32++ {
        v5 = v5
        _ = fi32
    }
    if (!v5) {
        v33 := GenS3{F0: 0, F1: -2147483648, F2: -754}
        _ = v33
    } else {
        v34 := GenS1{F0: 18446744073709551615, F1: 255, F2: (v4 >> ((44) & 63))}
        _ = v34
    }
    if v5 {
        if v5 {
            var v35 uint64 = 4294967295
            var v36 uint64 = (v35 % ((v35) | 1))
            v37 := GenS4{F0: v5, F1: 1, F2: 77}
            _ = v35
            _ = v36
            _ = v37
        } else {
            v38 := GenS5{F0: 348}
            v5 = (255 != 178)
            var v39 string = "foo bar"
            _ = v38
            _ = v39
        }
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = len(v3)
    }
    if (v5 && true) {
        var v40 uint8 = 128
        _ = v40
    }
    var v41 uint64 = 42
    v42 := GenS3{F0: 204, F1: 32767, F2: 42}
    v43 := GenS3{F0: 128, F1: -1, F2: int32(v4)}
    for fi44 := 0; fi44 < 8; fi44++ {
        if len(v1) > 0 {
            v1[0] = 0
        }
        if v5 {
            for fi45 := 0; fi45 < 3; fi45++ {
                v4 = (v4 * v4)
                if (3.14159 >= float32(3.14159)) {
                    v46 := v3[4294967295]
                    _ = v46
                } else {
                    var v47 int64 = v4
                    var v48 uint = 18446744073709551615
                    _ = v47
                    _ = v48
                }
                _ = fi45
            }
        }
        _ = fi44
    }
    v41 = 4294967295
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
