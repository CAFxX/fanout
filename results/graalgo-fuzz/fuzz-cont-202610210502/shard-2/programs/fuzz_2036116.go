package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 int64; F2 uint16 }
type GenS4 struct { F0 int8; F1 uint; F2 uint16 }
type GenS5 struct { F0 uint16; F1 int }
type GenS6 struct { F0 int32 }
func genFn0(p0 int) int16 {
    v7 := []int{p0, (^(p0)), p0}
    v7 = append(v7, -9223372036854775808)
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := GenS1{F0: 165}
        _ = fi8
        _ = v9
    }
    var v10 int64 = 42
    _ = p0
    _ = v10
    _ = v7
    return -1
}

func main() {
    var v1 uint64 = 390
    v2 := []string{((("" + "hello") + "") + ""), ("a" + "hello")}
    v3 := make(map[uint32]string)
    var v4 int8 = 42
    var v5 uint8 = 128
    v6 := []string{("hello" + "\n\t"), "", (("" + "a") + "\n\t")}
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 string = "a"
        var v13 int16 = -621
        fmt.Println(v4, v13)
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 bool = false
    v15 := GenS1{F0: genFn0(873)}
    fmt.Println(v4)
    var v16 uint16 = 42
    if v14 {
        v17 := GenS1{F0: 1}
        v2 = append(v2, ((((("hello" + ("hello" + "\n\t")) + "a") + (("hello" + "foo bar") + "")) + (("a" + "\n\t") + "a")) + "a"))
        _ = v17
    } else {
        v1 = v1
        if (false || v14) {
            v18 := GenS2{F0: 0}
            v5 = (v5 % ((55) | 1))
            _ = v18
        }
    }
    v1 += v1
    var v19 int64 = 42
    if false {
        v14 = v14
        v20 := GenS2{F0: 18446744073709551615}
        _ = v20
    } else {
        v21 := GenS1{F0: 1000}
        switch "\n\t" {
        case "hello":
            v22 := GenS2{F0: (v1 - 68)}
            _ = v22
        case "a":
            if (!(65535 == 316)) {
                if v14 {
                    var v23 uint8 = (v5 * (v5 & 255))
                    v4 += v4
                    _ = v23
                }
                fmt.Println(v19, v1, v4)
                switch 2147483647 {
                case -476:
                    v24 := GenS3{F0: "foo bar", F1: -1, F2: (v16 + (v16 << ((1) & 63)))}
                    _ = v24
                case 627:
                    if v14 {
                        var v25 uint64 = v1
                        switch -147 {
                        case 2147483647:
                            v19 -= v19
                        case 127:
                            var v26 uint32 = 0
                            _ = v26
                        default:
                            for fi27 := 0; fi27 < 11; fi27++ {
                                switch "" {
                                case "":
                                    var v28 string = ("foo bar" + "a")
                                    _ = v28
                                default:
                                    fmt.Println(v25)
                                }
                                _ = fi27
                            }
                        }
                        _ = v25
                    } else {
                        v14 = v14
                        delete(v3, 42)
                        v29 := GenS4{F0: 127, F1: 18446744073709551615, F2: 559}
                        _ = v29
                    }
                }
            } else {
                v30 := GenS5{F0: (v16 % ((v16) | 1)), F1: -1}
                for fi31 := 0; fi31 < 14; fi31++ {
                    var v32 float32 = -2.25
                    switch ("hello" + (("foo bar" + ("\n\t" + ((("" + "a") + "hello") + ("foo bar" + "")))) + "\n\t")) {
                    case "\n\t":
                        for fi33 := 0; fi33 < 10; fi33++ {
                            delete(v3, 2147483647)
                            if true {
                                _ = len(v3)
                            }
                            _ = fi33
                        }
                    case "hello":
                        var v34 uint = uint(v5)
                        _ = v34
                    }
                    switch int(v4) {
                    case 0:
                        fmt.Println(fi31, v16, v5)
                    case -1:
                        if (0 == 0) {
                            v1 = (v1 << ((0) & 63))
                            if false {
                                switch ("foo bar" + "foo bar") {
                                case "foo bar":
                                    v3[4294967295] = "a"
                                case "a":
                                    switch 32768 {
                                    case -9223372036854775808:
                                        for fi35 := 0; fi35 < 9; fi35++ {
                                            var v36 int32 = 2147483647
                                            _ = fi35
                                            _ = v36
                                        }
                                    case 0:
                                        for fi37 := 0; fi37 < 3; fi37++ {
                                            v1 -= (^(v1))
                                            v19 -= (-(v19))
                                            v38 := GenS4{F0: (v4 - (v4 - 53)), F1: 4294967295, F2: 347}
                                            _ = fi37
                                            _ = v38
                                        }
                                    case 42:
                                        v16 += 255
                                    default:
                                        var v39 int8 = v4
                                        _ = v39
                                    }
                                }
                            }
                            delete(v3, 0)
                        }
                    case 513:
                        fmt.Println(v19)
                    default:
                        for fi40 := 0; fi40 < 14; fi40++ {
                            v14 = v14
                            v3[282] = "foo bar"
                            var v41 float64 = 0.1
                            _ = fi40
                            _ = v41
                        }
                    }
                    _ = fi31
                    _ = v32
                }
                _ = v30
            }
        case "foo bar":
            v2 = append(v2, ("a" + "foo bar"))
        }
        v5 += (v5 & uint8(v1))
        _ = v21
    }
    var v42 int8 = (-(v4))
    v43 := GenS6{F0: 195}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
