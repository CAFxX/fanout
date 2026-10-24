package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 uint16; F1 int8 }
type GenS6 struct { F0 int32; F1 int16 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint64; F1 uint }
type GenS9 struct { F0 uint32; F1 uint; F2 int8 }
func genFn0() uint {
    v5 := GenS1{F0: 0, F1: true}
    if (!false) {
        for fi6 := 0; fi6 < 4; fi6++ {
            fmt.Println(fi6)
            for fi7 := 0; fi7 < 1; fi7++ {
                var v8 int64 = -1
                v8 -= (v8 / ((int64(fi6)) | 1))
                var v9 int16 = 599
                _ = fi7
                _ = v8
                _ = v9
            }
            _ = fi6
        }
    } else {
        v10 := []int{128, 2, 32768}
        var v11 uint = 255
        fmt.Println(v11)
        _ = v10
        _ = v11
    }
    if true {
        var v12 uint32 = 65535
        v13 := []int{-1, 128, -273, 1}
        _ = v12
        _ = v13
    }
    _ = v5
    return 0
}
func genFn1(p0 uint16, p1 uint8, p2 bool) int {
    v14 := GenS2{F0: p0, F1: 4294967295}
    var v15 uint = genFn0()
    var v16 string = ("" + "a")
    v17 := GenS3{F0: (!false)}
    var v18 float32 = (-2.25 * float32((1e10 - 0.1)))
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    return -630
}

func main() {
    var v1 uint16 = 65535
    var v2 int = 148
    v3 := []uint32{42, 1, 1}
    v4 := make(map[uint32]uint32)
    v19 := GenS4{F0: 1, F1: 1}
    if true {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch "\n\t" {
        case "":
            fmt.Println(v2, v2)
        case "foo bar":
            switch "foo bar" {
            case "foo bar":
                switch genFn1(v1, 0, true) {
                case 9223372036854775807:
                    fmt.Println(v2, v2)
                case 32768:
                    v20 := GenS2{F0: v1, F1: 0}
                    _ = v20
                case -9223372036854775808:
                    v21 := GenS3{F0: (v1 == v1)}
                    _ = v21
                }
            case "\n\t":
                if (!(false || false)) {
                    var v22 uint64 = 1
                    if ((v2 | v2) != (-(v2))) {
                        for fi23 := 0; fi23 < 4; fi23++ {
                            var v24 int64 = -9223372036854775808
                            _ = fi23
                            _ = v24
                        }
                        v2 -= genFn1((^(v1)), 200, true)
                    } else {
                        v25 := GenS5{F0: v1, F1: 0}
                        if false {
                            var v26 uint8 = 2
                            _ = v26
                        }
                        if (true || ((v22 % ((v22) | 1)) != 2)) {
                            for fi27 := 0; fi27 < 12; fi27++ {
                                for fi28 := 0; fi28 < 14; fi28++ {
                                    switch "foo bar" {
                                    case "\n\t":
                                        if true {
                                            _ = len(v4)
                                            v22 -= 2
                                        } else {
                                            switch ("a" + ((("" + "\n\t") + "foo bar") + ("foo bar" + ("a" + (("" + "hello") + ""))))) {
                                            case "foo bar":
                                                v29 := GenS2{F0: v1, F1: 649}
                                                _ = v29
                                            default:
                                                if true {
                                                    var v30 uint64 = v22
                                                    v31 := GenS6{F0: 1, F1: 1}
                                                    var v32 string = ("" + "")
                                                    _ = v30
                                                    _ = v31
                                                    _ = v32
                                                } else {
                                                    var v33 uint64 = v22
                                                    _ = v33
                                                }
                                            }
                                        }
                                    case "foo bar":
                                        v34 := v4[2147483648]
                                        _ = v34
                                    default:
                                        _ = len(v3)
                                    }
                                    switch v2 {
                                    case 5:
                                        v35 := GenS7{F0: 1}
                                        _ = v35
                                    }
                                    var v36 bool = ((-(fi28)) <= fi28)
                                    _ = fi28
                                    _ = v36
                                }
                                _ = fi27
                            }
                            switch v2 {
                            case 2147483647:
                                var v37 float32 = 1e10
                                _ = v37
                            default:
                                if len(v3) > 0 {
                                    v3[0] = 971
                                }
                            }
                        }
                        _ = v25
                    }
                    v38 := v4[4294967295]
                    _ = v22
                    _ = v38
                }
            case "":
                v1 = (v1 ^ v1)
            }
        case "hello":
            fmt.Println(v2, v1)
        }
    }
    var v39 uint32 = 42
    v4[2147483647] = (v39 ^ v39)
    for fi40 := 0; fi40 < 7; fi40++ {
        for fi41 := 0; fi41 < 9; fi41++ {
            v39 = v39
            v42 := GenS8{F0: 42, F1: 2}
            if (true && true) {
                v43 := GenS9{F0: v39, F1: 1, F2: 42}
                v2 += 253
                v39 -= (v39 * (^(v39)))
                _ = v43
            }
            fmt.Println(fi40, fi40)
            if (false && (964 > -1)) {
                v44 := GenS2{F0: 861, F1: 231}
                _ = v44
            }
            var v45 uint64 = 848
            _ = fi41
            _ = v42
            _ = v45
        }
        _ = fi40
    }
    for fi46 := 0; fi46 < 15; fi46++ {
        var v47 uint8 = 127
        _ = fi46
        _ = v47
    }
    if len(v3) > 0 {
        v3[0] = v39
    }
    v1 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v39
    _ = v4
}
