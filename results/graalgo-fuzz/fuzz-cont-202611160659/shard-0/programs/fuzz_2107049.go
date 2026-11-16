package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 uint32; F2 int64 }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 uint64; F1 uint8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 string; F1 bool }
func genFn0(p0 int32, p1 int16) bool {
    var v5 int16 = 918
    fmt.Println(v5, p1, p0)
    var v6 int64 = 1
    var v7 uint8 = 127
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return false
}
func genFn1() int {
    var v8 int16 = -956
    v8 = (v8 << ((42) & 63))
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 string = "foo bar"
        switch 2147483647 {
        case 127:
            switch v10 {
            case "foo bar":
                var v11 uint8 = 192
                _ = v11
            default:
                var v12 uint32 = 0
                _ = v12
            }
        }
        switch (fi9 - (fi9 >> ((1) & 63))) {
        case 32768:
            v8 -= v8
        case 4294967295:
            if true {
                v10 = (v10 + v10)
            }
        default:
            var v13 uint8 = 0
            _ = v13
        }
        _ = fi9
        _ = v10
    }
    _ = v8
    return -1
}

func main() {
    var v1 uint16 = 255
    var v2 string = "foo bar"
    var v3 float64 = (0.0 / ((1e10 - (1.5 / 0.1)) * (1.5 / (1e10 - 0.1))))
    v4 := make(map[uint32]string)
    if false {
        v2 = ((v2 + "") + v2)
        var v14 uint8 = 255
        _ = v14
    }
    var v15 int64 = -1
    v16 := GenS1{F0: 0, F1: 2147483647}
    if false {
        var v17 int8 = 59
        _ = v17
    }
    v18 := GenS1{F0: 0, F1: 9223372036854775807}
    v2 = v2
    var v19 int8 = 1
    switch 9223372036854775807 {
    case -9223372036854775808:
        switch v2 {
        case "foo bar":
            var v20 int = 128
            _ = v20
        case "":
            v2 = ("\n\t" + v2)
        case "\n\t":
            v2 = v2
        default:
            v21 := GenS1{F0: (v19 << ((65535) & 63)), F1: 414}
            _ = v21
        }
    case 568:
        if false {
            for fi22 := 0; fi22 < 3; fi22++ {
                switch "\n\t" {
                case "":
                    if true {
                        delete(v4, 0)
                        switch ("a" + ((v2 + ("" + "\n\t")) + "foo bar")) {
                        case "foo bar":
                            v23 := []int{(fi22 - genFn1())}
                            _ = v23
                        case "hello":
                            if false {
                                var v24 uint16 = (^(v1))
                                _ = v24
                            }
                        case "":
                            for fi25 := 0; fi25 < 5; fi25++ {
                                fmt.Println(v2)
                                v26 := GenS1{F0: v19, F1: (fi22 / (((fi22 >> ((65535) & 63))) | 1))}
                                v27 := v4[709]
                                _ = fi25
                                _ = v26
                                _ = v27
                            }
                        default:
                            v19 = 127
                        }
                        fmt.Println(v2, v19)
                    }
                case "\n\t":
                    switch (fi22 | fi22) {
                    case -9223372036854775808:
                        v4[580] = (v2 + "hello")
                    default:
                        if false {
                            var v28 int16 = 32767
                            v29 := GenS2{F0: -1}
                            v4[2147483648] = v2
                            _ = v28
                            _ = v29
                        } else {
                            if (!true) {
                                var v30 uint32 = 42
                                _ = v30
                            } else {
                                v19 += v19
                                if false {
                                    v3 += v3
                                    for fi31 := 0; fi31 < 14; fi31++ {
                                        v32 := GenS3{F0: (v19 | (v19 ^ v19)), F1: 2147483647, F2: v15}
                                        v15 = v15
                                        switch -1 {
                                        case 9223372036854775807:
                                            var v33 uint = 387
                                            _ = v33
                                        case -1:
                                            v34 := GenS1{F0: v19, F1: (fi31 | (fi31 << ((16) & 63)))}
                                            _ = v34
                                        }
                                        _ = fi31
                                        _ = v32
                                    }
                                }
                                var v35 uint = 2
                                _ = v35
                            }
                        }
                    }
                }
                v19 = (v19 / ((127) | 1))
                if true {
                    v36 := []int{(fi22 + (fi22 >> ((2) & 63)))}
                    v37 := GenS4{F0: v2, F1: (v1 >> ((797) & 63))}
                    var v38 uint16 = (v1 >> ((18446744073709551615) & 63))
                    _ = v36
                    _ = v37
                    _ = v38
                }
                _ = fi22
            }
            var v39 uint32 = 685
            _ = v39
        }
    default:
        v1 -= v1
    }
    switch ("" + v2) {
    case "\n\t":
        fmt.Println(v2, v19, v1)
    default:
        v40 := GenS5{F0: 2, F1: 2}
        _ = v40
    }
    v3 = v3
    switch ("hello" + v2) {
    case "a":
        for fi41 := 0; fi41 < 7; fi41++ {
            switch (fi41 << ((2) & 63)) {
            case -647:
                var v42 float32 = float32(float64(0.0))
                _ = v42
            case 4294967295:
                var v43 uint64 = 574
                _ = v43
            case 912:
                if false {
                    if true {
                        var v44 uint = 65535
                        _ = v44
                    }
                    if (!(v2 > "")) {
                        var v45 uint32 = 42
                        _ = v45
                    } else {
                        var v46 float64 = -2.25
                        _ = v46
                    }
                    v1 = 42
                }
            default:
                var v47 uint8 = 42
                _ = v47
            }
            fmt.Println(v1)
            _ = fi41
        }
    case "":
        switch -305 {
        case 128:
            var v48 int32 = 1
            _ = v48
        case 668:
            if (((2 >= 42) || false) && false) {
                v49 := GenS6{F0: -1}
                _ = v49
            }
        default:
            v50 := GenS7{F0: ((v2 + v2) + v2), F1: ((v19 + (v19 % ((v19) | 1))) > v19)}
            _ = v50
        }
    default:
        v51 := []int{32767, -1, 32768, genFn1()}
        _ = v51
    }
    v52 := []int{}
    v53 := GenS6{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v52:", len(v52), v52)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v52
    _ = v53
}
