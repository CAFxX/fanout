package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int8; F1 int64; F2 uint8 }
type GenS3 struct { F0 int16; F1 uint64; F2 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool; F1 int32 }
type GenS6 struct { F0 uint32 }
func genFn0(p0 int32) uint64 {
    v5 := GenS1{F0: 42, F1: (p0 - p0)}
    v6 := GenS2{F0: 0, F1: 42, F2: 39}
    var v7 float64 = -2.25
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 116
}
func genFn1(p0 uint, p1 int16, p2 uint8) uint {
    var v8 float32 = (0.0 / 0.1)
    if (true && false) {
        v9 := []int{-181}
        switch -1 {
        case 9223372036854775807:
            p1 = (p1 * (p1 + p1))
        }
        v8 += 3.14159
        _ = v9
    } else {
        p1 = p1
    }
    v10 := GenS3{F0: p1, F1: 18446744073709551615, F2: ((65535 > 4294967295) && (!false))}
    var v11 uint8 = uint8(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v8
    return 844
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 float32 = 0.1
    var v3 uint16 = 433
    var v4 uint64 = 2
    v12 := GenS4{F0: (("foo bar" + "foo bar") + ("" + "foo bar"))}
    v3 = (v3 ^ v3)
    v2 += v2
    v13 := []int{}
    v1 += v1
    if true {
        switch int(v1) {
        case -2147483648:
            if (!true) {
                v4 -= v4
                v1 = (v1 + (v1 / ((v1) | 1)))
            }
        default:
            v13 = append(v13, 127)
        }
        v2 = v2
    } else {
        v13 = append(v13, 304)
        switch "foo bar" {
        case "hello":
            if len(v13) > 0 {
                v13[0] = 42
            }
        case "a":
            v1 = v1
        default:
            for fi14 := 0; fi14 < 9; fi14++ {
                v1 = 0
                v15 := GenS3{F0: 0, F1: (^(v4)), F2: true}
                if len(v13) > 0 {
                    v13[0] = (fi14 * int(v4))
                }
                _ = fi14
                _ = v15
            }
        }
    }
    var v16 bool = false
    var v17 float32 = ((float32((1e10 - 0.0)) + v2) / v2)
    if ((!v16) && v16) {
        v2 = (v17 / ((v17 / (-2.25 * (((v17 * (0.0 / v17)) + v2) + v2))) + (v2 + (float32(((0.0 / (((-2.25 + 3.14159) + -2.25) / float64(v2))) - 1e10)) - ((v17 * ((float32(0.0) * float32((1.5 + 1.5))) * 0.1)) - (float32(0.1) * float32(((float64(float32(0.1)) - 1.5) - 3.14159))))))))
        fmt.Println(v17)
    } else {
        switch "foo bar" {
        case "hello":
            var v18 string = "foo bar"
            _ = v18
        case "\n\t":
            var v19 int8 = 1
            _ = v19
        case "a":
            fmt.Println(v4, v17)
        }
        var v20 uint64 = v4
        var v21 int16 = 0
        _ = v20
        _ = v21
    }
    if (-588 > 0) {
        v4 = 18446744073709551615
        switch -296 {
        case -2147483648:
            if true {
                for fi22 := 0; fi22 < 10; fi22++ {
                    switch "" {
                    case "":
                        if (45 >= int16(v4)) {
                            v23 := GenS3{F0: -1, F1: v4, F2: v16}
                            v24 := GenS5{F0: (!(!v16)), F1: 1}
                            var v25 int64 = v1
                            _ = v23
                            _ = v24
                            _ = v25
                        } else {
                            fmt.Println(v3, v3, v2)
                            if ((42 >= 0) && false) {
                                if true {
                                    var v26 int16 = 42
                                    _ = v26
                                }
                            }
                        }
                    case "foo bar":
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                    }
                    var v27 int8 = 100
                    v3 += 65535
                    _ = fi22
                    _ = v27
                }
                v13 = append(v13, 0)
                v28 := GenS6{F0: 2}
                _ = v28
            }
        case 39:
            for fi29 := 0; fi29 < 8; fi29++ {
                var v30 uint32 = 65535
                if (-1 > -854) {
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                }
                _ = fi29
                _ = v30
            }
        case 128:
            var v31 uint16 = (^(v3))
            _ = v31
        default:
            _ = len(v13)
        }
    } else {
        switch ("foo bar" + "") {
        case "\n\t":
            for fi32 := 0; fi32 < 16; fi32++ {
                if false {
                    for fi33 := 0; fi33 < 5; fi33++ {
                        v16 = (-1 >= -2147483648)
                        var v34 int32 = 1
                        _ = fi33
                        _ = v34
                    }
                    _ = len(v13)
                    fmt.Println(v4)
                } else {
                    switch "foo bar" {
                    case "hello":
                        v17 += v2
                    default:
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                    }
                }
                fmt.Println(v17)
                v35 := GenS1{F0: -9223372036854775808, F1: 640}
                _ = fi32
                _ = v35
            }
        case "":
            if (v4 > 418) {
                if v16 {
                    v36 := GenS4{F0: (("\n\t" + ("" + ("\n\t" + "\n\t"))) + "")}
                    v17 = (v2 / v17)
                    fmt.Println(v17, v16, v3)
                    _ = v36
                }
                var v37 float64 = float64(0.0)
                if (v16 || (!v16)) {
                    _ = len(v13)
                } else {
                    var v38 uint = 162
                    _ = v38
                }
                _ = v37
            }
        }
        for fi39 := 0; fi39 < 1; fi39++ {
            var v40 int = 1
            for fi41 := 0; fi41 < 5; fi41++ {
                var v42 float32 = -2.25
                fmt.Println(v17, fi39, v2)
                _ = fi41
                _ = v42
            }
            v43 := GenS2{F0: 100, F1: 42, F2: 1}
            _ = fi39
            _ = v40
            _ = v43
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
