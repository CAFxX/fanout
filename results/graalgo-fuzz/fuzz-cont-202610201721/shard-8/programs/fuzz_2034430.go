package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool; F1 uint16; F2 uint64 }
func genFn0(p0 int32) int8 {
    v7 := []int{9223372036854775807, 32767}
    if true {
        for fi8 := 0; fi8 < 13; fi8++ {
            v9 := GenS1{F0: 128}
            if false {
                var v10 uint64 = 4294967295
                for fi11 := 0; fi11 < 10; fi11++ {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    fmt.Println(fi11)
                    _ = fi11
                }
                _ = v10
            } else {
                var v12 int64 = 9223372036854775807
                if (float32((-2.25 + 0.0)) > float32(0.0)) {
                    var v13 int16 = 0
                    _ = v13
                }
                var v14 float64 = (-2.25 + 3.14159)
                _ = v12
                _ = v14
            }
            _ = len(v7)
            _ = fi8
            _ = v9
        }
    }
    switch -2147483648 {
    case 9223372036854775807:
        fmt.Println(p0)
    case -492:
        var v15 float32 = (-2.25 + ((1e10 - 0.0) + 3.14159))
        _ = v15
    }
    _ = p0
    _ = v7
    return -128
}
func genFn1(p0 bool, p1 uint8, p2 uint64) bool {
    v16 := []int{-9223372036854775808, 9223372036854775807, 128, 2147483647}
    v17 := GenS2{F0: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return p0
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]int)
    var v3 string = "\n\t"
    var v4 uint32 = 0
    v5 := make(map[string]string)
    var v6 uint = 0
    for fi18 := 0; fi18 < 18; fi18++ {
        v19 := []int{(fi18 & 32767), fi18, fi18}
        var v20 int8 = genFn0(709)
        for fi21 := 0; fi21 < 3; fi21++ {
            var v22 string = v3
            v23 := GenS3{F0: -32768}
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi18
        _ = v19
        _ = v20
    }
    v24 := v5[("foo bar" + v3)]
    for fi25 := 0; fi25 < 2; fi25++ {
        switch v24 {
        case "\n\t":
            var v26 uint64 = 0
            _ = v26
        case "a":
            v24 = v24
        }
        _ = fi25
    }
    _ = len(v1)
    for fi27 := 0; fi27 < 3; fi27++ {
        v24 = v3
        for fi28 := 0; fi28 < 8; fi28++ {
            var v29 int16 = 1
            _ = fi28
            _ = v29
        }
        _ = fi27
    }
    for fi30 := 0; fi30 < 2; fi30++ {
        switch -805 {
        case 42:
            var v31 int = (fi30 >> (((^(v6))) & 63))
            _ = v31
        case 376:
            var v32 int64 = 9223372036854775807
            _ = v32
        default:
            switch fi30 {
            case 1:
                v3 = v24
            case -408:
                v33 := GenS3{F0: 1}
                _ = v33
            case -1:
                v34 := []int{1, 32768, (fi30 + 4294967295), (fi30 | (fi30 & (^(fi30))))}
                _ = v34
            }
        }
        v35 := []int{(fi30 % (((fi30 + (fi30 / ((fi30) | 1)))) | 1)), (fi30 - (fi30 * 32768)), 4294967295, 1}
        _ = fi30
        _ = v35
    }
    delete(v1, (v3 + "hello"))
    v36 := []int{-279}
    v37 := GenS2{F0: -347}
    fmt.Println(v4, v24, v3)
    for fi38 := 0; fi38 < 16; fi38++ {
        v39 := GenS4{F0: (0 == -228), F1: uint16(v4), F2: 1}
        if (uint16(v4) <= 550) {
            v40 := GenS3{F0: 0}
            _ = v40
        } else {
            if true {
                if false {
                    for fi41 := 0; fi41 < 14; fi41++ {
                        v6 += 18446744073709551615
                        _ = fi41
                    }
                    for fi42 := 0; fi42 < 4; fi42++ {
                        var v43 bool = (1e10 >= float32(0.0))
                        v6 = v6
                        _ = fi42
                        _ = v43
                    }
                } else {
                    v24 = (v3 + v24)
                    if len(v36) >= 2 {
                        v36 = v36[1:len(v36)]
                    }
                    if true {
                        if (1 != 1) {
                            v5[v3] = v3
                            v44 := GenS2{F0: -9223372036854775808}
                            _ = v44
                        }
                        v45 := v1["\n\t"]
                        var v46 uint16 = 735
                        _ = v45
                        _ = v46
                    } else {
                        if (618 > 1) {
                            if true {
                                switch v24 {
                                case "foo bar":
                                    if ((268 > (v4 << ((v6) & 63))) || false) {
                                        v24 = ((((v24 + ("foo bar" + (v3 + v3))) + ((("a" + (v24 + (v24 + ""))) + v3) + (("\n\t" + v3) + v24))) + (v3 + v3)) + (v24 + "a"))
                                        switch v24 {
                                        case "\n\t":
                                            v4 = 0
                                        default:
                                            fmt.Println(v4, v4)
                                        }
                                        fmt.Println(v6)
                                    }
                                default:
                                    var v47 uint8 = 105
                                    _ = v47
                                }
                            }
                        }
                        v36 = append(v36, 984)
                        v3 = "foo bar"
                    }
                }
            } else {
                _ = len(v1)
            }
            var v48 int16 = 136
            _ = v48
        }
        _ = fi38
        _ = v39
    }
    v49 := GenS4{F0: false, F1: 0, F2: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v49
    _ = v5
    _ = v6
}
