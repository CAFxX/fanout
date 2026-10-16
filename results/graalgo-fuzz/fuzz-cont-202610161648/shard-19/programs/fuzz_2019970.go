package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 uint32; F1 bool; F2 bool }
func genFn0() uint16 {
    switch "hello" {
    case "\n\t":
        var v6 uint = 940
        _ = v6
    case "foo bar":
        var v7 int8 = 67
        _ = v7
    case "":
        var v8 float64 = (((0.0 / 1e10) - ((1e10 - (-2.25 * 0.1)) + -2.25)) * 3.14159)
        _ = v8
    }
    var v9 int32 = 0
    var v10 uint32 = 42
    v10 = 4294967295
    v11 := GenS1{F0: -9223372036854775808}
    _ = v10
    _ = v11
    _ = v9
    return 65535
}
func genFn1(p0 bool, p1 int32) uint16 {
    v12 := GenS1{F0: 1}
    if p0 {
        for fi13 := 0; fi13 < 20; fi13++ {
            if p0 {
                var v14 int16 = -854
                v15 := []int{725, (fi13 << ((374) & 63)), (fi13 * (fi13 ^ -523)), 596}
                if p0 {
                    v14 -= v14
                    if p0 {
                        if p0 {
                            var v16 int8 = 100
                            var v17 int64 = 172
                            _ = v16
                            _ = v17
                        }
                        if p0 {
                            p1 = -990
                            for fi18 := 0; fi18 < 15; fi18++ {
                                if true {
                                    v19 := GenS1{F0: 42}
                                    v20 := GenS1{F0: -826}
                                    var v21 float64 = (0.0 - float64(-2.25))
                                    v14 = v14
                                    switch "" {
                                    case "foo bar":
                                        var v22 int16 = -32768
                                        _ = v22
                                    default:
                                        var v23 float32 = float32(0.0)
                                        _ = v23
                                    }
                                    _ = v19
                                    _ = v20
                                    _ = v21
                                }
                                _ = fi18
                            }
                        }
                        v14 = (v14 >> ((42) & 63))
                    } else {
                        p1 = p1
                        for fi24 := 0; fi24 < 15; fi24++ {
                            v14 += v14
                            _ = fi24
                        }
                        p1 = -2147483648
                    }
                    var v25 int64 = 9223372036854775807
                    _ = v25
                }
                _ = v14
                _ = v15
            } else {
                p1 -= 80
                var v26 float64 = ((1.5 * -2.25) + (1.5 / 1.0))
                _ = v26
            }
            v27 := []int{(^(fi13)), fi13, fi13, (^(fi13))}
            _ = fi13
            _ = v27
        }
        fmt.Println(p0, p1)
    }
    var v28 float32 = (1e10 + ((1.5 - float32(float64(1e10))) / 1e10))
    switch "\n\t" {
    case "":
        p1 = (p1 << ((465) & 63))
    case "\n\t":
        for fi29 := 0; fi29 < 10; fi29++ {
            v28 -= (1e10 - v28)
            v30 := GenS1{F0: -464}
            _ = fi29
            _ = v30
        }
    case "a":
        switch "" {
        case "a":
            if (32767 != 0) {
                var v31 uint16 = 1
                if false {
                    var v32 uint32 = 2147483648
                    fmt.Println(p0)
                    _ = v32
                } else {
                    var v33 uint8 = 217
                    _ = v33
                }
                _ = v31
            } else {
                fmt.Println(p0, v28, p1)
                var v34 uint = 255
                v35 := GenS2{F0: 0, F1: 2, F2: 255}
                _ = v34
                _ = v35
            }
        }
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v28
    return 65535
}

func main() {
    v1 := make(map[string]int)
    v2 := []string{("" + "\n\t"), "foo bar", "a"}
    var v3 string = "a"
    v4 := []uint32{2, 826, 75, 65535}
    var v5 uint32 = 816
    v36 := GenS3{F0: 1}
    var v37 uint16 = 255
    if (v5 <= v5) {
        for fi38 := 0; fi38 < 20; fi38++ {
            var v39 int32 = -1
            _ = fi38
            _ = v39
        }
        v37 += (^(v37))
        for fi40 := 0; fi40 < 19; fi40++ {
            v37 = 65535
            v5 = (v5 << ((348) & 63))
            v41 := GenS4{F0: false, F1: -1}
            _ = fi40
            _ = v41
        }
    } else {
        for fi42 := 0; fi42 < 10; fi42++ {
            v43 := GenS3{F0: -307}
            v5 -= v5
            _ = fi42
            _ = v43
        }
    }
    v5 = 4294967295
    v44 := GenS4{F0: (!(-392 <= 2147483647)), F1: -1}
    if (2147483647 <= 1) {
        switch v3 {
        case "\n\t":
            if ((!false) && (!true)) {
                v5 -= v5
                v45 := GenS4{F0: (!true), F1: 602}
                switch 0 {
                case 992:
                    v1[("a" + (v3 + "\n\t"))] = 4294967295
                case 6:
                    var v46 int64 = 0
                    _ = v46
                }
                _ = v45
            }
        case "hello":
            var v47 uint8 = 127
            _ = v47
        case "foo bar":
            switch "hello" {
            case "hello":
                var v48 int32 = 1
                _ = v48
            default:
                if false {
                    var v49 uint8 = 127
                    _ = v49
                } else {
                    delete(v1, "\n\t")
                }
            }
        }
        v50 := GenS3{F0: 0}
        _ = v50
    } else {
        var v51 uint8 = 2
        _ = v51
    }
    if true {
        v5 += v5
    }
    fmt.Println(v5, v5, v3)
    v52 := GenS5{F0: (v5 | (v5 + 65535)), F1: (!(1 >= 1000)), F2: (0 <= 179)}
    for fi53 := 0; fi53 < 11; fi53++ {
        v1[v3] = fi53
        _ = fi53
    }
    switch v3 {
    case "hello":
        fmt.Println(v5, v5)
    case "a":
        fmt.Println(v3)
    case "":
        if false {
            v3 = (((v3 + ((v3 + v3) + v3)) + ("" + "\n\t")) + v3)
        }
    }
    v4 = append(v4, 403)
    fmt.Println(v37)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v44
    _ = v5
    _ = v52
}
