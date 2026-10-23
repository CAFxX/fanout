package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int64 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 int16 }
func genFn0(p0 int16, p1 string, p2 uint8) uint {
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 uint8 = p2
        v7 := GenS1{F0: 470, F1: p0, F2: -210}
        fmt.Println(p1, p0)
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: 103, F1: 831, F2: -9223372036854775808}
    v9 := GenS1{F0: 65535, F1: p0, F2: 1}
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 18446744073709551615
}
func genFn1() int8 {
    var v10 string = ("foo bar" + ("a" + ("foo bar" + "")))
    if ((3.14159 / float64((((1e10 / float32(-2.25)) + 1.5) + 1e10))) > -2.25) {
        v10 = (v10 + "a")
    }
    var v11 float64 = 1e10
    v12 := GenS2{F0: 0, F1: 0}
    var v13 int32 = 0
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return 42
}

func main() {
    var v1 int32 = -1
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]bool)
    var v4 float64 = ((((3.14159 + 1.5) + -2.25) * 0.1) + -2.25)
    v14 := GenS3{F0: 2147483647}
    v1 += 1
    for fi15 := 0; fi15 < 16; fi15++ {
        if ((!(389 == fi15)) && (4294967295 == 18446744073709551615)) {
            if (583 < 65535) {
                v1 = (v1 + v1)
                if false {
                    v1 = v1
                }
            } else {
                if true {
                    fmt.Println(v4)
                    if (!true) {
                        v4 -= (((float64(0.0) - v4) * v4) - (v4 + (v4 / 1.0)))
                        if false {
                            var v16 int = (fi15 + fi15)
                            delete(v2, 2147483648)
                            v1 = v1
                            _ = v16
                        } else {
                            v17 := GenS3{F0: -460}
                            if (!(true || (-9223372036854775808 >= 42))) {
                                var v18 int32 = 36
                                _ = v18
                            } else {
                                switch (fi15 * fi15) {
                                case 979:
                                    v19 := []int{(fi15 % ((fi15) | 1)), fi15, (-(fi15))}
                                    _ = v19
                                }
                            }
                            _ = v17
                        }
                        v4 = 0.0
                    } else {
                        var v20 int32 = int32(fi15)
                        _ = v20
                    }
                } else {
                    switch "" {
                    case "":
                        if (-1 != -32768) {
                            switch (fi15 >> ((966) & 63)) {
                            case 128:
                                switch -574 {
                                case -277:
                                    for fi21 := 0; fi21 < 7; fi21++ {
                                        v22 := GenS4{F0: 454, F1: 42}
                                        var v23 float32 = 0.0
                                        _ = fi21
                                        _ = v22
                                        _ = v23
                                    }
                                case 127:
                                    if true {
                                        v24 := []int{}
                                        v25 := GenS4{F0: 42, F1: 0}
                                        _ = v24
                                        _ = v25
                                    } else {
                                        v1 -= v1
                                    }
                                }
                            case -2147483648:
                                v26 := GenS4{F0: 42, F1: int16(v1)}
                                _ = v26
                            }
                        } else {
                            for fi27 := 0; fi27 < 10; fi27++ {
                                var v28 float32 = (0.0 + -2.25)
                                fmt.Println(v4)
                                var v29 int64 = 9223372036854775807
                                _ = fi27
                                _ = v28
                                _ = v29
                            }
                            for fi30 := 0; fi30 < 10; fi30++ {
                                var v31 float32 = 0.1
                                if (4294967295 > 0) {
                                    v1 -= (v1 | (v1 & 420))
                                    fmt.Println(v31, fi30)
                                    switch "hello" {
                                    case "hello":
                                        switch "hello" {
                                        case "\n\t":
                                            var v32 uint64 = 1
                                            _ = v32
                                        case "":
                                            v33 := v3[2147483647]
                                            _ = v33
                                        }
                                    case "":
                                        var v34 uint64 = 540
                                        switch (fi15 << ((255) & 63)) {
                                        case 880:
                                            switch ("a" + "hello") {
                                            case "":
                                                fmt.Println(fi30, v31, fi15)
                                            default:
                                                v35 := GenS2{F0: genFn1(), F1: 127}
                                                _ = v35
                                            }
                                        }
                                        _ = v34
                                    case "\n\t":
                                        v36 := []int{fi15, (fi30 + fi30)}
                                        _ = v36
                                    }
                                } else {
                                    fmt.Println(v1, v4)
                                }
                                v1 = v1
                                _ = fi30
                                _ = v31
                            }
                        }
                    default:
                        v37 := []int{(fi15 % ((-363) | 1)), -280, 0}
                        _ = v37
                    }
                    v38 := v2[383]
                    _ = v38
                }
            }
        } else {
            for fi39 := 0; fi39 < 7; fi39++ {
                var v40 int16 = 1
                _ = fi39
                _ = v40
            }
            v41 := []int{32767, (fi15 * 9223372036854775807), fi15}
            _ = v41
        }
        _ = fi15
    }
    if (!false) {
        v1 = 0
        v42 := GenS3{F0: 42}
        _ = v42
    }
    var v43 int8 = 100
    for fi44 := 0; fi44 < 20; fi44++ {
        v1 += -2147483648
        _ = fi44
    }
    var v45 float64 = v4
    switch ("hello" + "") {
    case "\n\t":
        switch "" {
        case "\n\t":
            var v46 float64 = (float64((3.14159 + 0.0)) / -2.25)
            _ = v46
        }
    case "":
        var v47 int16 = -32768
        _ = v47
    }
    v48 := GenS4{F0: -1, F1: 474}
    var v49 int8 = (v43 >> ((4294967295) & 63))
    var v50 int8 = -128
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v45
    _ = v48
    _ = v49
    _ = v50
}
