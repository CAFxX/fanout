package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int64 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS4 struct { F0 uint32; F1 int16; F2 uint16 }
func genFn0(p0 uint64) bool {
    v7 := []int{-9223372036854775808}
    p0 = p0
    v8 := GenS1{F0: 42, F1: 65535, F2: -298}
    _ = p0
    _ = v7
    _ = v8
    return false
}
func genFn1(p0 int8, p1 uint32) uint32 {
    v9 := GenS1{F0: -2147483648, F1: 186, F2: -9223372036854775808}
    p1 -= (p1 & (p1 - p1))
    var v10 int16 = -1
    v10 -= v10
    p1 += (^(p1))
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return (^(p1))
}

func main() {
    var v1 int = -9223372036854775808
    var v2 uint64 = 2
    var v3 uint32 = 0
    v4 := make(map[int]string)
    var v5 float32 = (3.14159 / float32(3.14159))
    var v6 int = v1
    switch "foo bar" {
    case "\n\t":
        v3 = v3
    case "foo bar":
        for fi11 := 0; fi11 < 19; fi11++ {
            switch (^(fi11)) {
            case -1:
                var v12 int = 2147483647
                _ = v12
            case 9223372036854775807:
                v13 := GenS1{F0: -2147483648, F1: 833, F2: -1}
                _ = v13
            case 42:
                if (901 != 1) {
                    if false {
                        var v14 float32 = 0.0
                        var v15 uint8 = 128
                        _ = v14
                        _ = v15
                    } else {
                        v6 -= (fi11 ^ (v1 | 128))
                        v16 := []int{9223372036854775807, v6, fi11}
                        switch int(v3) {
                        case 32767:
                            _ = len(v16)
                        case -398:
                            for fi17 := 0; fi17 < 18; fi17++ {
                                if false {
                                    var v18 int8 = -1
                                    var v19 int32 = 0
                                    _ = v18
                                    _ = v19
                                }
                                switch (v1 + (v1 * (v6 * fi11))) {
                                case 1:
                                    if true {
                                        if ((^(v3)) < v3) {
                                            v2 += (v2 | 0)
                                            v20 := GenS2{F0: (v6 % ((v1) | 1)), F1: 127}
                                            v3 = v3
                                            _ = v20
                                        } else {
                                            var v21 uint16 = 42
                                            _ = v21
                                        }
                                    } else {
                                        v22 := GenS3{F0: 0, F1: 0, F2: 42}
                                        switch fi11 {
                                        case 621:
                                            var v23 uint64 = (v2 / ((v2) | 1))
                                            _ = v23
                                        default:
                                            if (!(771 >= 713)) {
                                                v1 = v1
                                                v1 -= (^(fi17))
                                            }
                                            if true {
                                                v5 = 1.5
                                                if (!((-1 < 1) && (2 > 391))) {
                                                    if false {
                                                        fmt.Println(fi11)
                                                        v24 := GenS3{F0: 2, F1: uint16(v2), F2: 2147483647}
                                                        _ = len(v4)
                                                        _ = v24
                                                    } else {
                                                        var v25 uint8 = uint8(v2)
                                                        if (true || true) {
                                                            var v26 int = (fi11 >> ((4294967295) & 63))
                                                            _ = v26
                                                        }
                                                        _ = v25
                                                    }
                                                    v27 := GenS4{F0: (v3 + (v3 + v3)), F1: 1, F2: 2}
                                                    var v28 uint8 = 42
                                                    _ = v27
                                                    _ = v28
                                                }
                                                v29 := GenS3{F0: 2, F1: uint16(v3), F2: -2147483648}
                                                _ = v29
                                            } else {
                                                switch fi17 {
                                                case 89:
                                                    v16 = append(v16, v6)
                                                case 32768:
                                                    v30 := GenS2{F0: (v6 * -211), F1: 242}
                                                    v31 := GenS1{F0: 896, F1: uint(v2), F2: -340}
                                                    switch (fi11 ^ 261) {
                                                    case -2147483648:
                                                        v32 := GenS2{F0: 1, F1: 170}
                                                        _ = v32
                                                    default:
                                                        switch ("hello" + ("foo bar" + "hello")) {
                                                        case "a":
                                                            var v33 string = "hello"
                                                            _ = v33
                                                        default:
                                                            var v34 bool = true
                                                            _ = v34
                                                        }
                                                    }
                                                    _ = v30
                                                    _ = v31
                                                case -849:
                                                    v35 := v4[(^(fi17))]
                                                    fmt.Println(fi11, v1)
                                                    var v36 int16 = -32768
                                                    _ = v35
                                                    _ = v36
                                                }
                                                switch "hello" {
                                                case "foo bar":
                                                    v37 := GenS1{F0: -1, F1: 1, F2: 1}
                                                    _ = v37
                                                case "\n\t":
                                                    v38 := GenS4{F0: genFn1(0, (v3 * 65535)), F1: 1000, F2: 255}
                                                    var v39 float32 = -2.25
                                                    _ = v38
                                                    _ = v39
                                                }
                                            }
                                            v1 -= (^(v1))
                                        }
                                        _ = v22
                                    }
                                default:
                                    if ((!((3.14159 + 1.5) != (float64((0.1 * ((1.5 / float32((1.5 * (1.5 * -2.25)))) / ((v5 + v5) / 0.1)))) / 3.14159))) && (0 != 22)) {
                                        var v40 uint32 = (v3 - (v3 & 1))
                                        var v41 float32 = (1.5 - (v5 + (v5 / v5)))
                                        _ = v40
                                        _ = v41
                                    }
                                }
                                _ = fi17
                            }
                        }
                        _ = v16
                    }
                }
            default:
                v2 += uint64(v3)
            }
            _ = fi11
        }
    case "hello":
        var v42 uint64 = (v2 % ((4294967295) | 1))
        _ = v42
    }
    fmt.Println(v3, v2, v1)
    for fi43 := 0; fi43 < 8; fi43++ {
        v44 := []int{(v1 & fi43), (fi43 - fi43), (v6 * v6), fi43}
        _ = fi43
        _ = v44
    }
    v2 += v2
    v45 := GenS1{F0: 1, F1: 652, F2: -1}
    var v46 uint = 18446744073709551615
    var v47 int8 = 0
    var v48 uint64 = uint64(v46)
    var v49 string = ("\n\t" + "foo bar")
    var v50 float64 = 1.5
    switch v49 {
    case "foo bar":
        var v51 int16 = -52
        _ = v51
    default:
        if true {
            v52 := GenS4{F0: v3, F1: 0, F2: 32767}
            _ = v52
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
