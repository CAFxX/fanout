package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int32 }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint }
type GenS3 struct { F0 uint; F1 int32; F2 bool }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 int16 }
func genFn0() uint32 {
    var v5 bool = ((!((-32768 <= 1) || false)) || (false || (false || (231 <= 2))))
    switch 0 {
    case -1:
        if v5 {
            fmt.Println(v5)
            if (1.5 != float64((((float32((1e10 + -2.25)) - 0.0) / float32(((-2.25 / 1.5) / (3.14159 + -2.25)))) / -2.25))) {
                v5 = (0 >= 32767)
                var v6 uint64 = 0
                if (296 <= -451) {
                    if true {
                        if (v5 || ("" < ((((("foo bar" + "hello") + "") + "\n\t") + "") + "hello"))) {
                            v7 := GenS1{F0: 42, F1: "", F2: -458}
                            for fi8 := 0; fi8 < 17; fi8++ {
                                if true {
                                    fmt.Println(v5)
                                    var v9 int32 = -255
                                    fmt.Println(v9, v5)
                                    _ = v9
                                } else {
                                    v5 = (231 >= 42)
                                    v6 += 42
                                    fmt.Println(v5, fi8)
                                }
                                if (21 != 1) {
                                    var v10 uint32 = 42
                                    var v11 bool = (1 < 9223372036854775807)
                                    _ = v10
                                    _ = v11
                                }
                                _ = fi8
                            }
                            _ = v7
                        } else {
                            if v5 {
                                var v12 float64 = 1.5
                                v13 := GenS2{F0: 51, F1: 1, F2: 255}
                                _ = v12
                                _ = v13
                            }
                            var v14 float32 = 1.5
                            v15 := GenS2{F0: 841, F1: 0, F2: 65535}
                            _ = v14
                            _ = v15
                        }
                    }
                } else {
                    if true {
                        var v16 float64 = 0.1
                        var v17 uint64 = (^(v6))
                        _ = v16
                        _ = v17
                    }
                    var v18 uint64 = (v6 & (v6 & (v6 % ((v6) | 1))))
                    var v19 uint = 65535
                    _ = v18
                    _ = v19
                }
                _ = v6
            }
            v5 = v5
        } else {
            v20 := GenS2{F0: 2, F1: 255, F2: 65535}
            fmt.Println(v5)
            _ = v20
        }
    case 0:
        var v21 int32 = -749
        _ = v21
    }
    var v22 int8 = -85
    _ = v22
    _ = v5
    return 2147483647
}

func main() {
    var v1 uint16 = 127
    var v2 uint32 = 0
    var v3 uint8 = 187
    var v4 uint = uint(v2)
    var v23 uint64 = 1
    v24 := GenS2{F0: 4294967295, F1: (^(v1)), F2: v4}
    v3 -= v3
    v25 := GenS3{F0: uint(v2), F1: 0, F2: false}
    fmt.Println(v2, v3, v2)
    if (42 >= 32767) {
        if true {
            var v26 uint64 = (v23 * v23)
            v27 := GenS2{F0: uint32(v26), F1: (v1 - v1), F2: (v4 >> ((255) & 63))}
            _ = v26
            _ = v27
        } else {
            v28 := GenS2{F0: v2, F1: 700, F2: v4}
            _ = v28
        }
        switch (("hello" + "a") + "") {
        case "hello":
            v29 := GenS4{F0: uint32(v23), F1: 65535}
            _ = v29
        case "\n\t":
            v30 := []int{}
            _ = v30
        default:
            v4 = (v4 + v4)
        }
        var v31 bool = (!(false && (!(uint8(v23) != 235))))
        _ = v31
    }
    var v32 float32 = 0.0
    var v33 int16 = 1000
    if false {
        switch ("" + "hello") {
        case "hello":
            for fi34 := 0; fi34 < 19; fi34++ {
                if false {
                    for fi35 := 0; fi35 < 2; fi35++ {
                        var v36 int32 = 600
                        _ = fi35
                        _ = v36
                    }
                    if true {
                        v23 = 42
                        fmt.Println(v3, v2, v23)
                    }
                    switch ("foo bar" + "a") {
                    case "a":
                        var v37 int8 = 127
                        _ = v37
                    case "foo bar":
                        v3 = (v3 | (v3 - v3))
                    case "":
                        if (true || ((uint64(v1) <= 0) && ((((-(fi34)) == (^(fi34))) || true) && true))) {
                            v38 := []int{fi34}
                            var v39 string = ""
                            var v40 float32 = ((float32(-2.25) / 1e10) * 0.0)
                            _ = v38
                            _ = v39
                            _ = v40
                        }
                    default:
                        v41 := GenS1{F0: 9223372036854775807, F1: "\n\t", F2: 42}
                        _ = v41
                    }
                } else {
                    var v42 int8 = int8(fi34)
                    var v43 uint = v4
                    v3 = 2
                    _ = v42
                    _ = v43
                }
                _ = fi34
            }
        case "":
            var v44 int8 = 127
            _ = v44
        case "\n\t":
            v33 -= v33
        }
    }
    v45 := GenS5{F0: v33}
    for fi46 := 0; fi46 < 14; fi46++ {
        for fi47 := 0; fi47 < 5; fi47++ {
            v48 := []int{(fi46 - fi46), fi47, 19}
            v49 := GenS1{F0: -1, F1: "\n\t", F2: -1}
            if len(v48) > 0 {
                v48[0] = fi47
            }
            _ = fi47
            _ = v48
            _ = v49
        }
        var v50 int32 = 1
        _ = fi46
        _ = v50
    }
    v2 = uint32(v23)
    v23 -= v23
    var v51 int8 = -128
    var v52 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v45
    _ = v51
    _ = v52
}
