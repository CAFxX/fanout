package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint16; F1 int64; F2 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint8; F1 uint8; F2 uint }
func genFn0(p0 int16, p1 string, p2 int64) int16 {
    var v5 uint8 = 123
    p1 = "\n\t"
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p0
}

func main() {
    var v1 uint8 = 164
    var v2 uint64 = 228
    var v3 float64 = float64(float32((float64(0.1) + 1.5)))
    var v4 float32 = (1e10 / float32((v3 - -2.25)))
    v4 = v4
    fmt.Println(v3)
    if true {
        var v6 uint = 2
        v3 = v3
        if false {
            var v7 int = 127
            _ = v7
        }
        _ = v6
    }
    var v8 uint = 65535
    v4 -= 0.0
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 string = ("" + "foo bar")
        var v11 int32 = -336
        v8 += (v8 >> ((v8) & 63))
        _ = fi9
        _ = v10
        _ = v11
    }
    if (false || true) {
        v12 := GenS1{F0: "hello", F1: 0}
        _ = v12
    }
    if true {
        fmt.Println(v3, v2, v3)
        v13 := GenS2{F0: 525, F1: -1, F2: (!(!false))}
        _ = v13
    }
    v14 := []int{536, 0, 42}
    if false {
        if true {
            var v15 int64 = 9223372036854775807
            _ = v15
        } else {
            switch -2147483648 {
            case 32767:
                for fi16 := 0; fi16 < 3; fi16++ {
                    if false {
                        fmt.Println(v4, v8)
                        if false {
                            if ((3.14159 - (v4 / 1e10)) < v4) {
                                fmt.Println(v8)
                            } else {
                                for fi17 := 0; fi17 < 20; fi17++ {
                                    v2 = 18446744073709551615
                                    _ = fi17
                                }
                            }
                            v14 = append(v14, fi16)
                            for fi18 := 0; fi18 < 10; fi18++ {
                                v1 = (v1 ^ v1)
                                _ = fi18
                            }
                        } else {
                            for fi19 := 0; fi19 < 8; fi19++ {
                                if (v4 >= ((((v4 - float32(((v3 + v3) - v3))) / v4) + v4) + v4)) {
                                    var v20 uint16 = 255
                                    switch "foo bar" {
                                    case "":
                                        v21 := GenS3{F0: 1}
                                        _ = v21
                                    default:
                                        var v22 uint = uint(v2)
                                        _ = v22
                                    }
                                    v8 -= (v8 + 432)
                                    switch -9223372036854775808 {
                                    case 0:
                                        fmt.Println(v20, v20, v20)
                                    case -9223372036854775808:
                                        v2 -= (v2 * (v2 & (v2 + (v2 * v2))))
                                    default:
                                        v1 = 151
                                    }
                                    _ = v20
                                }
                                if false {
                                    var v23 uint8 = 255
                                    v24 := GenS3{F0: -1}
                                    v14 = append(v14, 1)
                                    _ = v23
                                    _ = v24
                                }
                                var v25 bool = true
                                _ = fi19
                                _ = v25
                            }
                            v26 := GenS1{F0: (((("" + "a") + "hello") + ("foo bar" + "a")) + "hello"), F1: 2}
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                            _ = v26
                        }
                        for fi27 := 0; fi27 < 12; fi27++ {
                            var v28 uint = v8
                            v29 := GenS4{F0: -9223372036854775808, F1: "a"}
                            fmt.Println(v3, v8)
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                    }
                    _ = fi16
                }
            case -236:
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
            default:
                v30 := GenS5{F0: 200}
                _ = v30
            }
        }
        if true {
            v3 = 1e10
            switch 127 {
            case 100:
                var v31 uint64 = 1
                _ = v31
            case 32767:
                var v32 uint32 = 2147483648
                _ = v32
            case 4294967295:
                for fi33 := 0; fi33 < 11; fi33++ {
                    var v34 uint16 = 32767
                    if ((!true) || (v3 > ((0.1 - v3) / (0.0 - 3.14159)))) {
                        if len(v14) > 0 {
                            v14[0] = fi33
                        }
                    } else {
                        v2 = (v2 + (v2 & (v2 << ((v8) & 63))))
                    }
                    for fi35 := 0; fi35 < 12; fi35++ {
                        var v36 float64 = v3
                        var v37 int64 = -1
                        fmt.Println(v37, v2, v1)
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                    _ = fi33
                    _ = v34
                }
            }
        }
        v38 := GenS6{F0: uint8(v2), F1: (v1 & (v1 / ((86) | 1))), F2: (v8 & v8)}
        _ = v38
    } else {
        fmt.Println(v3, v3, v1)
        v14 = append(v14, 1)
    }
    var v39 float64 = (v3 / 3.14159)
    var v40 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v8
}
