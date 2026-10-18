package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 bool; F1 uint; F2 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 string; F1 string }
func genFn0() uint16 {
    var v6 uint32 = 2147483647
    v6 = 42
    v6 -= 42
    _ = v6
    return uint16(v6)
}

func main() {
    v1 := []uint32{42, 2147483647, 2147483647}
    var v2 uint64 = 42
    var v3 int8 = -128
    v4 := []int{127, 4294967295, 950, 826}
    var v5 int = -2147483648
    if (false && (!true)) {
        var v7 int = v5
        _ = v7
    } else {
        if (true || false) {
            var v8 int32 = 2147483647
            fmt.Println(v2)
            v9 := GenS1{F0: (v5 / ((v5) | 1))}
            _ = v8
            _ = v9
        } else {
            var v10 uint64 = v2
            v10 = 72
            _ = v10
        }
        var v11 int64 = 94
        fmt.Println(v11, v3, v5)
        _ = v11
    }
    var v12 int8 = v3
    switch v5 {
    case -9223372036854775808:
        fmt.Println(v5)
    case 128:
        v5 -= (v5 + v5)
    default:
        var v13 uint32 = 73
        _ = v13
    }
    v14 := GenS2{F0: 32767, F1: ""}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if ((false && false) || (42 > 120)) {
        switch "foo bar" {
        case "":
            fmt.Println(v3, v3)
        case "a":
            for fi15 := 0; fi15 < 13; fi15++ {
                var v16 uint16 = 2
                v5 = 32767
                _ = fi15
                _ = v16
            }
        case "foo bar":
            var v17 int16 = -1
            _ = v17
        default:
            if true {
                var v18 uint64 = v2
                _ = v18
            } else {
                fmt.Println(v5)
                var v19 int64 = 0
                _ = v19
            }
        }
        switch int(v3) {
        case 545:
            fmt.Println(v5, v2)
        }
    } else {
        if false {
            for fi20 := 0; fi20 < 13; fi20++ {
                var v21 string = "hello"
                _ = fi20
                _ = v21
            }
        }
        for fi22 := 0; fi22 < 10; fi22++ {
            v12 += 27
            var v23 uint8 = 35
            _ = fi22
            _ = v23
        }
    }
    var v24 uint = 446
    if (255 >= 0) {
        for fi25 := 0; fi25 < 3; fi25++ {
            var v26 float32 = (-2.25 * -2.25)
            v27 := GenS3{F0: false, F1: 255, F2: 0}
            _ = fi25
            _ = v26
            _ = v27
        }
        switch ("foo bar" + ("\n\t" + "a")) {
        case "a":
            switch "a" {
            case "":
                v2 = v2
            case "hello":
                _ = len(v4)
            }
        default:
            for fi28 := 0; fi28 < 2; fi28++ {
                v2 += 18446744073709551615
                var v29 int16 = 1
                for fi30 := 0; fi30 < 11; fi30++ {
                    var v31 int64 = 9223372036854775807
                    _ = fi30
                    _ = v31
                }
                _ = fi28
                _ = v29
            }
        }
        v24 = (v24 % ((v24) | 1))
    } else {
        fmt.Println(v12, v12, v5)
        v5 = v5
        v32 := GenS1{F0: (v5 * 4294967295)}
        _ = v32
    }
    v33 := GenS4{F0: (42 >= -561)}
    var v34 uint32 = 494
    v35 := GenS2{F0: 1, F1: "a"}
    switch -2147483648 {
    case -1:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case 32768:
        var v36 uint = 1
        _ = v36
    case 649:
        if false {
            fmt.Println(v24, v34, v5)
            switch (v5 / (((v5 + (v5 * 9223372036854775807))) | 1)) {
            case 2147483647:
                if len(v4) > 0 {
                    v4[0] = v5
                }
            case 128:
                v37 := GenS3{F0: false, F1: (v24 / (((v24 - v24)) | 1)), F2: 42}
                _ = v37
            default:
                for fi38 := 0; fi38 < 15; fi38++ {
                    var v39 string = "a"
                    _ = fi38
                    _ = v39
                }
            }
            for fi40 := 0; fi40 < 8; fi40++ {
                v41 := GenS4{F0: (((748 == uint16(v34)) && (83 != uint8(v34))) && false)}
                _ = fi40
                _ = v41
            }
        } else {
            v1 = append(v1, 963)
            if (-2147483648 == -1) {
                for fi42 := 0; fi42 < 18; fi42++ {
                    v5 += (^(fi42))
                    _ = fi42
                }
                for fi43 := 0; fi43 < 16; fi43++ {
                    switch "foo bar" {
                    case "":
                        if (!(genFn0() < genFn0())) {
                            v1 = append(v1, 4294967295)
                            if false {
                                if (true && true) {
                                    if (true && (false && (-567 > 2147483647))) {
                                        v44 := GenS5{F0: genFn0()}
                                        fmt.Println(v12, v3)
                                        _ = v44
                                    }
                                } else {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                v2 -= (v2 | (v2 / ((245) | 1)))
                            }
                            v45 := GenS6{F0: ("hello" + "a"), F1: "hello"}
                            _ = v45
                        }
                    case "\n\t":
                        v5 += fi43
                    default:
                        var v46 int16 = 1000
                        _ = v46
                    }
                    _ = fi43
                }
                fmt.Println(v12)
            }
            fmt.Println(v24, v2, v12)
        }
    default:
        v2 += (v2 ^ (v2 ^ 42))
    }
    if (!(!(genFn0() < 0))) {
        fmt.Println(v12, v3, v5)
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (1e10 >= -2.25) {
        if len(v4) > 0 {
            v4[0] = (v5 - v5)
        }
        var v47 uint = (v24 - v24)
        _ = v47
    } else {
        v48 := GenS5{F0: 32767}
        var v49 int16 = -32768
        _ = v48
        _ = v49
    }
    fmt.Println(v3, v12, v24)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v24
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
