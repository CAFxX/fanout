package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint16 }
func genFn0() uint8 {
    v6 := []int{9223372036854775807, 32768, 32767, -9223372036854775808}
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 int64 = 1
        _ = fi7
        _ = v8
    }
    _ = v6
    return 127
}
func genFn1(p0 int64, p1 uint16) uint16 {
    v9 := GenS1{F0: 65535, F1: (p0 ^ (-(p0))), F2: 4294967295}
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 int16 = 0
        _ = fi10
        _ = v11
    }
    p0 -= (p0 + p0)
    _ = p0
    _ = p1
    _ = v9
    return 32767
}

func main() {
    var v1 int8 = 127
    var v2 int8 = v1
    var v3 int16 = 42
    v4 := make(map[uint32]uint32)
    var v5 float64 = 0.0
    switch "foo bar" {
    case "a":
        var v12 int8 = v1
        _ = v12
    case "":
        switch ("\n\t" + "hello") {
        case "a":
            if false {
                v13 := GenS2{F0: v1}
                v14 := GenS3{F0: v1}
                _ = v13
                _ = v14
            } else {
                var v15 bool = (!false)
                _ = v15
            }
        default:
            if false {
                var v16 uint = 18446744073709551615
                for fi17 := 0; fi17 < 14; fi17++ {
                    var v18 int64 = 658
                    if true {
                        if true {
                            v2 = (v1 & v2)
                            fmt.Println(v3, fi17, v3)
                        } else {
                            v19 := GenS4{F0: "\n\t"}
                            for fi20 := 0; fi20 < 13; fi20++ {
                                v4[982] = 2147483648
                                _ = fi20
                            }
                            var v21 uint16 = 1
                            _ = v19
                            _ = v21
                        }
                        var v22 uint = v16
                        v23 := v4[2147483648]
                        _ = v22
                        _ = v23
                    } else {
                        switch -1 {
                        case -978:
                            var v24 int64 = v18
                            _ = v24
                        case 0:
                            var v25 int64 = v18
                            _ = v25
                        default:
                            for fi26 := 0; fi26 < 16; fi26++ {
                                var v27 float64 = v5
                                v1 = v2
                                if (0 >= 810) {
                                    v28 := []int{fi26, 175, fi17}
                                    if false {
                                        var v29 uint32 = 4294967295
                                        delete(v4, (^(v29)))
                                        _ = v29
                                    } else {
                                        v1 -= 0
                                    }
                                    _ = len(v4)
                                    _ = v28
                                }
                                _ = fi26
                                _ = v27
                            }
                        }
                        fmt.Println(v2, fi17, v16)
                    }
                    v30 := []int{(fi17 + fi17), 128, (fi17 & 4294967295)}
                    _ = fi17
                    _ = v18
                    _ = v30
                }
                v2 = (v1 & v1)
                _ = v16
            }
        }
    case "hello":
        var v31 uint8 = 155
        _ = v31
    default:
        v2 = 42
    }
    var v32 bool = (73 >= 4294967295)
    var v33 float64 = (3.14159 * v5)
    var v34 string = "\n\t"
    var v35 int = 2147483647
    var v36 int32 = -1
    v1 = int8(v3)
    var v37 uint = 42
    var v38 uint64 = 890
    var v39 uint64 = uint64(v37)
    var v40 string = ""
    var v41 int32 = (v36 % (((v36 & (v36 - v36))) | 1))
    if (false || (!v32)) {
        if v32 {
            var v42 float32 = float32(((3.14159 - 3.14159) - v33))
            _ = v42
        } else {
            if v32 {
                _ = len(v4)
            } else {
                if v32 {
                    v4[65535] = 65
                    var v43 int8 = v2
                    if ((^(v3)) >= 42) {
                        v44 := GenS5{F0: (v3 % ((v3) | 1))}
                        if (42 != (v36 * v41)) {
                            _ = len(v4)
                            var v45 uint8 = genFn0()
                            v32 = false
                            _ = v45
                        }
                        _ = v44
                    } else {
                        fmt.Println(v32, v40, v39)
                    }
                    _ = v43
                } else {
                    v40 = ((((("foo bar" + "foo bar") + v34) + "a") + v40) + v34)
                    var v46 uint32 = 65535
                    v47 := GenS5{F0: -507}
                    _ = v46
                    _ = v47
                }
            }
            fmt.Println(v41)
            fmt.Println(v33, v2, v34)
        }
        v48 := []int{-619, (v35 ^ (-(v35)))}
        _ = v48
    }
    v49 := GenS2{F0: 0}
    var v50 int32 = 2147483647
    v51 := GenS6{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v49
    _ = v5
    _ = v50
    _ = v51
}
