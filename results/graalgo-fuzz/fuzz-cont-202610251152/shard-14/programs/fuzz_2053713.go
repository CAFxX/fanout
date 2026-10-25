package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 uint8; F2 uint }
type GenS3 struct { F0 int32; F1 string; F2 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int; F1 int; F2 uint32 }
func genFn0(p0 int) int {
    if true {
        p0 += (-(p0))
    } else {
        p0 = (p0 | p0)
    }
    switch "" {
    case "hello":
        v5 := GenS1{F0: 237}
        _ = v5
    case "":
        fmt.Println(p0)
    default:
        var v6 float32 = 0.1
        _ = v6
    }
    if true {
        v7 := []int{(p0 % ((2147483647) | 1))}
        p0 += 0
        _ = v7
    } else {
        var v8 float32 = (1.5 / float32(float64(3.14159)))
        _ = v8
    }
    _ = p0
    return (p0 >> ((4294967295) & 63))
}
func genFn1(p0 uint16) uint64 {
    var v9 int8 = 127
    v10 := GenS1{F0: 42}
    _ = p0
    _ = v10
    _ = v9
    return 0
}

func main() {
    var v1 int32 = 229
    var v2 int32 = (^(v1))
    var v3 string = "foo bar"
    var v4 bool = (true || (((!(v3 > (v3 + v3))) || true) && (286 != 753)))
    for fi11 := 0; fi11 < 15; fi11++ {
        v12 := GenS2{F0: (v2 - v1), F1: 2, F2: 659}
        v3 = v3
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 int8 = 42
        _ = fi13
        _ = v14
    }
    if v4 {
        for fi15 := 0; fi15 < 13; fi15++ {
            if v4 {
                var v16 int32 = v2
                _ = v16
            }
            for fi17 := 0; fi17 < 20; fi17++ {
                v3 = (v3 + (v3 + v3))
                switch fi15 {
                case 418:
                    if (65535 <= 448) {
                        var v18 uint32 = 42
                        var v19 float64 = 1.5
                        _ = v18
                        _ = v19
                    } else {
                        if (0 != 901) {
                            var v20 bool = v4
                            _ = v20
                        } else {
                            v21 := []int{}
                            v1 = v1
                            _ = v21
                        }
                        fmt.Println(v3, fi17)
                        v22 := []int{(fi17 - fi15), fi15, genFn0(fi15)}
                        _ = v22
                    }
                default:
                    switch v3 {
                    case "":
                        if v4 {
                            switch (fi17 % (((fi17 << ((1) & 63))) | 1)) {
                            case -746:
                                v4 = (42 <= 32767)
                            }
                        } else {
                            var v23 int8 = 1
                            var v24 float32 = ((-2.25 + (1.5 - 0.1)) / ((3.14159 + ((1.5 * ((0.1 * (float32(0.1) * (((1e10 / 1e10) / ((((1e10 + 0.0) + 3.14159) + (0.0 * -2.25)) / 0.1)) - (0.1 * 0.1)))) / 1.5)) / 1.0)) + 0.1))
                            v25 := GenS1{F0: (^(v1))}
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                    default:
                        v4 = v4
                    }
                }
                if (((-2.25 / float32(float64((1.5 * (((1e10 + (-2.25 + 1e10)) - ((1e10 - 3.14159) * 3.14159)) + -2.25))))) - 1e10) >= 1.5) {
                    v3 = "hello"
                }
                _ = fi17
            }
            _ = fi15
        }
    } else {
        for fi26 := 0; fi26 < 7; fi26++ {
            v27 := GenS3{F0: v2, F1: "", F2: 9223372036854775807}
            if (v4 && v4) {
                var v28 uint = 255
                if v4 {
                    switch v3 {
                    case "\n\t":
                        var v29 uint8 = 255
                        _ = v29
                    case "":
                        v4 = true
                    }
                } else {
                    v3 = "hello"
                    fmt.Println(v4)
                    for fi30 := 0; fi30 < 12; fi30++ {
                        var v31 int8 = int8(v2)
                        v4 = (v4 || v4)
                        switch fi30 {
                        case 42:
                            var v32 uint32 = 0
                            _ = v32
                        case -581:
                            switch (^(fi26)) {
                            case -2147483648:
                                var v33 int64 = 0
                                _ = v33
                            case 0:
                                v34 := GenS4{F0: (v28 - v28)}
                                _ = v34
                            default:
                                v31 += v31
                            }
                        case 93:
                            var v35 string = v3
                            _ = v35
                        default:
                            if v4 {
                                var v36 uint16 = 413
                                var v37 uint64 = 428
                                _ = v36
                                _ = v37
                            }
                        }
                        _ = fi30
                        _ = v31
                    }
                }
                _ = v28
            } else {
                if (42 <= 687) {
                    v38 := GenS4{F0: 0}
                    if v4 {
                        fmt.Println(v4, v4)
                        fmt.Println(fi26, v2, v1)
                        v39 := []int{}
                        _ = v39
                    }
                    _ = v38
                }
            }
            v40 := []int{9223372036854775807, (fi26 - genFn0((^(fi26))))}
            _ = fi26
            _ = v27
            _ = v40
        }
    }
    v3 = v3
    if (1 >= 86) {
        switch 4294967295 {
        case 9223372036854775807:
            var v41 uint8 = 127
            _ = v41
        case -2147483648:
            var v42 int16 = -101
            _ = v42
        case -976:
            var v43 float32 = 1e10
            _ = v43
        default:
            switch 482 {
            case -1:
                var v44 int64 = 9223372036854775807
                _ = v44
            case 128:
                v4 = false
            }
        }
    }
    var v45 uint16 = 200
    for fi46 := 0; fi46 < 2; fi46++ {
        v47 := GenS5{F0: 331, F1: fi46, F2: 0}
        var v48 bool = v4
        v1 = 471
        _ = fi46
        _ = v47
        _ = v48
    }
    fmt.Println(v45)
    var v49 uint8 = 127
    v1 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v49
}
