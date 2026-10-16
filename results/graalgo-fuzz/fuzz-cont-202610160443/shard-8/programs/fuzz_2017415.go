package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint32 }
type GenS2 struct { F0 int16; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint64; F1 uint64; F2 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint }
func genFn0(p0 bool, p1 uint64, p2 string) uint64 {
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 int64 = -1
        if p0 {
            var v7 float32 = 3.14159
            _ = v7
        } else {
            p0 = p0
        }
        _ = fi5
        _ = v6
    }
    v8 := []int{9223372036854775807, 0, 239, 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 2
}
func genFn1(p0 uint8) int8 {
    switch "" {
    case "foo bar":
        if (!false) {
            for fi9 := 0; fi9 < 13; fi9++ {
                var v10 int = (fi9 % (((fi9 % ((fi9) | 1))) | 1))
                _ = fi9
                _ = v10
            }
            var v11 int8 = -128
            if true {
                v12 := GenS1{F0: 18446744073709551615, F1: -9223372036854775808, F2: 143}
                var v13 int16 = -333
                _ = v12
                _ = v13
            } else {
                v14 := []int{4294967295, -879, 32768}
                _ = v14
            }
            _ = v11
        }
    }
    if false {
        p0 = (^(p0))
    } else {
        var v15 uint16 = 228
        if false {
            p0 = (p0 << ((1) & 63))
            switch 4294967295 {
            case -1:
                var v16 int = 961
                _ = v16
            case 9223372036854775807:
                for fi17 := 0; fi17 < 12; fi17++ {
                    for fi18 := 0; fi18 < 1; fi18++ {
                        v15 = 65535
                        fmt.Println(v15, fi17, p0)
                        v19 := GenS2{F0: -319, F1: 1, F2: 1}
                        var v20 int = 127
                        v21 := []int{fi18, fi17, v20}
                        _ = fi18
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                    for fi22 := 0; fi22 < 4; fi22++ {
                        if false {
                            v23 := GenS3{F0: 452, F1: 580, F2: false}
                            _ = v23
                        }
                        var v24 int16 = 32767
                        _ = fi22
                        _ = v24
                    }
                    _ = fi17
                }
            case 4294967295:
                var v25 int = -9223372036854775808
                _ = v25
            }
        }
        v26 := GenS3{F0: 1, F1: 18446744073709551615, F2: (4294967295 < 42)}
        _ = v15
        _ = v26
    }
    p0 = (p0 / (((p0 >> ((18446744073709551615) & 63))) | 1))
    p0 = 255
    var v27 int16 = -219
    _ = p0
    _ = v27
    return 42
}

func main() {
    var v1 uint32 = 314
    v2 := make(map[int]uint32)
    var v3 string = ""
    var v4 uint32 = v1
    v28 := GenS3{F0: genFn0(false, 518, (v3 + ("foo bar" + v3))), F1: 0, F2: (true && (!(uint16(v4) <= 362)))}
    if ((true || true) && (!(true && false))) {
        v29 := GenS4{F0: ""}
        v30 := []int{-9223372036854775808, 38}
        v1 -= (^(v4))
        _ = v29
        _ = v30
    } else {
        v3 = v3
    }
    v31 := GenS3{F0: 2, F1: 4294967295, F2: true}
    v32 := GenS5{F0: -32}
    v4 = (^(v4))
    var v33 uint8 = 42
    v34 := []int{32768, 585, 32767}
    if (false && (true && true)) {
        v4 += v1
        if len(v34) >= 2 {
            v34 = v34[1:len(v34)]
        }
    } else {
        v2[605] = 9
        var v35 int64 = -752
        for fi36 := 0; fi36 < 16; fi36++ {
            v33 = (v33 / ((v33) | 1))
            for fi37 := 0; fi37 < 5; fi37++ {
                _ = len(v2)
                if true {
                    fmt.Println(v1, fi37)
                    var v38 float64 = 0.0
                    fmt.Println(v4, v1, v4)
                    _ = v38
                }
                var v39 int64 = (-(v35))
                _ = fi37
                _ = v39
            }
            v2[(-(fi36))] = v1
            _ = fi36
        }
        _ = v35
    }
    var v40 uint = 445
    v41 := GenS6{F0: (v40 + 880)}
    if (v1 < 1) {
        var v42 uint = (v40 + (v40 + (v40 * 974)))
        v1 = v4
        _ = v42
    }
    v33 = 189
    v1 = 2147483647
    for fi43 := 0; fi43 < 4; fi43++ {
        var v44 float32 = float32(1.5)
        if true {
            fmt.Println(fi43)
            v3 = v3
        }
        _ = fi43
        _ = v44
    }
    v45 := v2[127]
    if true {
        var v46 uint32 = v45
        v34 = append(v34, 127)
        _ = v46
    }
    var v47 bool = (true && (false || true))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v40
    _ = v41
    _ = v45
    _ = v47
}
