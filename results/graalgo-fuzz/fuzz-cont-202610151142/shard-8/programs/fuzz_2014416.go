package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16; F1 int16; F2 int64 }
type GenS5 struct { F0 int16; F1 uint64 }
func genFn0(p0 int16) int {
    var v7 uint = 0
    var v8 uint32 = 710
    v9 := GenS1{F0: 0, F1: 1}
    fmt.Println(p0, v8, p0)
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return -9223372036854775808
}
func genFn1(p0 uint8) string {
    p0 = (p0 | 127)
    var v10 int8 = 0
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 float64 = 1.5
        fmt.Println(v12, p0, fi11)
        _ = fi11
        _ = v12
    }
    v10 += (^(v10))
    _ = p0
    _ = v10
    return "hello"
}

func main() {
    var v1 uint16 = 1
    var v2 int = 878
    v3 := make(map[uint32]string)
    var v4 int16 = 32767
    var v5 bool = true
    v6 := []uint32{2147483647, 877}
    v13 := GenS1{F0: v1, F1: 622}
    v14 := v3[4294967295]
    var v15 string = v14
    v6 = append(v6, 450)
    v16 := v3[1]
    for fi17 := 0; fi17 < 14; fi17++ {
        v1 = 67
        if v5 {
            switch 4294967295 {
            case -1:
                fmt.Println(v5)
            case -2147483648:
                fmt.Println(fi17, v15, v16)
            case 9223372036854775807:
                v18 := GenS1{F0: v1, F1: 0}
                _ = v18
            }
        } else {
            v1 = 0
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            v16 = "a"
        }
        var v19 int = fi17
        _ = fi17
        _ = v19
    }
    var v20 int32 = int32(v2)
    fmt.Println(v20)
    var v21 string = (v16 + "foo bar")
    if v5 {
        v4 -= -1
    } else {
        for fi22 := 0; fi22 < 1; fi22++ {
            v23 := GenS1{F0: v1, F1: (v1 | (v1 / (((v1 % ((1) | 1))) | 1)))}
            if ((v20 ^ (v20 << ((65535) & 63))) != -1) {
                fmt.Println(v5, v5)
                switch (v16 + (("a" + v14) + "hello")) {
                case "hello":
                    fmt.Println(v4)
                case "\n\t":
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
            } else {
                if false {
                    var v24 uint16 = 0
                    _ = v24
                } else {
                    fmt.Println(v4)
                    if (1e10 > (((((0.1 - 0.1) * -2.25) / 1.5) + ((1.5 + 1e10) / 1.0)) / float64(3.14159))) {
                        if v5 {
                            fmt.Println(v14, v16)
                            v16 = v21
                            for fi25 := 0; fi25 < 2; fi25++ {
                                if (false && false) {
                                    switch (-(fi22)) {
                                    case 9223372036854775807:
                                        if ((v21 + "\n\t") < "\n\t") {
                                            if (381 >= 29) {
                                                if false {
                                                    var v26 string = v15
                                                    _ = v26
                                                } else {
                                                    v16 = v15
                                                    fmt.Println(fi22, v4)
                                                }
                                                v27 := GenS1{F0: v1, F1: v1}
                                                v4 += v4
                                                _ = v27
                                            }
                                            var v28 float32 = -2.25
                                            if (2 <= 2) {
                                                v5 = true
                                                var v29 float64 = 0.1
                                                _ = v29
                                            }
                                            var v30 int = (fi22 << ((2) & 63))
                                            _ = v28
                                            _ = v30
                                        }
                                    case 2147483647:
                                        var v31 string = ""
                                        _ = v31
                                    case 32768:
                                        fmt.Println(v1, v4)
                                    }
                                    var v32 int8 = 0
                                    _ = v32
                                } else {
                                    if ((v16 + (((((v15 + (v16 + (v15 + v16))) + v21) + ("a" + v21)) + v14) + "")) > ((v15 + v21) + v21)) {
                                        v33 := GenS2{F0: 0, F1: 42}
                                        _ = v33
                                    } else {
                                        var v34 uint32 = 42
                                        _ = v34
                                    }
                                    v1 = (v1 - (v1 >> ((0) & 63)))
                                }
                                v2 = 4294967295
                                var v35 int8 = -58
                                _ = fi25
                                _ = v35
                            }
                        }
                        for fi36 := 0; fi36 < 7; fi36++ {
                            v14 = "a"
                            fmt.Println(v5)
                            v4 = 383
                            _ = fi36
                        }
                        v37 := GenS3{F0: 1}
                        _ = v37
                    } else {
                        v21 = ""
                        for fi38 := 0; fi38 < 20; fi38++ {
                            fmt.Println(v21, fi22, v1)
                            if v5 {
                                var v39 uint64 = 263
                                _ = v39
                            }
                            var v40 uint = 18446744073709551615
                            _ = fi38
                            _ = v40
                        }
                    }
                }
                v41 := GenS4{F0: (v1 << ((4294967295) & 63)), F1: 32767, F2: 9223372036854775807}
                if len(v6) > 0 {
                    v6[0] = 42
                }
                _ = v41
            }
            _ = fi22
            _ = v23
        }
        v15 = (v21 + ("hello" + v15))
        fmt.Println(v21, v4, v20)
    }
    if v5 {
        var v42 int32 = 2147483647
        v43 := GenS3{F0: 42}
        var v44 float64 = 1.5
        _ = v42
        _ = v43
        _ = v44
    }
    if len(v6) > 0 {
        v6[0] = 2147483648
    }
    switch (v15 + "foo bar") {
    case "a":
        var v45 float32 = ((1.5 + (1.5 - 3.14159)) - 0.0)
        _ = v45
    case "":
        switch genFn0(32767) {
        case 32768:
            v46 := GenS4{F0: v1, F1: v4, F2: -9223372036854775808}
            _ = v46
        case -1:
            for fi47 := 0; fi47 < 15; fi47++ {
                fmt.Println(v15)
                v14 = "a"
                _ = fi47
            }
        case -2147483648:
            v3[889] = v15
        default:
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
    case "foo bar":
        for fi48 := 0; fi48 < 18; fi48++ {
            fmt.Println(v1, v16, v5)
            if v5 {
                v49 := GenS5{F0: (v4 / ((42) | 1)), F1: 1}
                _ = v49
            }
            _ = fi48
        }
    default:
        fmt.Println(v1, v4, v21)
    }
    var v50 int16 = (v4 + (-(v4)))
    var v51 int = (v2 * v2)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v6
}
