package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 uint8; F2 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32; F1 uint8 }
func genFn0(p0 int64, p1 int8, p2 uint) int32 {
    v7 := []int{32768, 734}
    if len(v7) > 0 {
        v7[0] = 8
    }
    var v8 float64 = 1.5
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 0
}
func genFn1() uint8 {
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 float32 = 0.1
        fmt.Println(v10, fi9)
        v11 := []int{}
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := []int{42, -9223372036854775808}
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 int8 = 59
        for fi15 := 0; fi15 < 17; fi15++ {
            fmt.Println(v14, fi13, fi15)
            v14 = (v14 ^ (-(v14)))
            var v16 float32 = 0.0
            if len(v12) > 0 {
                v12[0] = fi15
            }
            _ = fi15
            _ = v16
        }
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        _ = fi13
        _ = v14
    }
    _ = v12
    return 255
}

func main() {
    v1 := []int{2147483647, 127, 708, -238}
    v2 := []int{4294967295, 2147483647}
    v3 := []string{"foo bar", ("hello" + "a")}
    v4 := []int{-2147483648}
    var v5 int8 = -128
    var v6 float64 = 1e10
    v6 -= v6
    if false {
        switch 693 {
        case -302:
            v2 = append(v2, -9223372036854775808)
        case 32768:
            if false {
                if (1 > 0) {
                    var v17 int8 = (v5 & (v5 + v5))
                    fmt.Println(v17, v6)
                    switch -117 {
                    case 4294967295:
                        v18 := GenS1{F0: "a"}
                        _ = v18
                    }
                    _ = v17
                }
                v4 = append(v4, 32767)
            } else {
                if false {
                    if (false && true) {
                        fmt.Println(v5)
                    } else {
                        for fi19 := 0; fi19 < 2; fi19++ {
                            if (9223372036854775807 < -9223372036854775808) {
                                var v20 string = ""
                                _ = v20
                            } else {
                                v1 = append(v1, fi19)
                                v1 = append(v1, fi19)
                            }
                            v21 := GenS1{F0: ""}
                            for fi22 := 0; fi22 < 11; fi22++ {
                                fmt.Println(fi19, v5)
                                v23 := GenS2{F0: (-(fi19)), F1: 2, F2: 223}
                                v5 = v5
                                _ = fi22
                                _ = v23
                            }
                            _ = fi19
                            _ = v21
                        }
                        switch 563 {
                        case 2147483647:
                            v5 -= (v5 & (v5 | v5))
                        default:
                            var v24 int32 = -2147483648
                            _ = v24
                        }
                    }
                    if true {
                        if false {
                            var v25 int64 = -264
                            _ = v25
                        } else {
                            v5 -= v5
                            switch "" {
                            case "\n\t":
                                v1 = append(v1, 128)
                            }
                            var v26 int64 = 9223372036854775807
                            _ = v26
                        }
                        var v27 float32 = 1e10
                        _ = v27
                    } else {
                        _ = len(v1)
                        for fi28 := 0; fi28 < 20; fi28++ {
                            for fi29 := 0; fi29 < 17; fi29++ {
                                var v30 int32 = int32(v5)
                                var v31 uint16 = 864
                                var v32 int = fi29
                                switch fi29 {
                                case -780:
                                    fmt.Println(v32, fi28, v5)
                                }
                                _ = fi29
                                _ = v30
                                _ = v31
                                _ = v32
                            }
                            _ = fi28
                        }
                    }
                    var v33 bool = (42 != 0)
                    _ = v33
                }
            }
        default:
            var v34 float32 = float32(v6)
            _ = v34
        }
        for fi35 := 0; fi35 < 6; fi35++ {
            v5 = v5
            v36 := GenS3{F0: "\n\t"}
            _ = fi35
            _ = v36
        }
        var v37 uint8 = 128
        _ = v37
    } else {
        _ = len(v4)
    }
    _ = len(v2)
    var v38 uint16 = 2
    for fi39 := 0; fi39 < 5; fi39++ {
        var v40 int8 = v5
        var v41 uint = uint(v38)
        if (!(false || false)) {
            for fi42 := 0; fi42 < 18; fi42++ {
                var v43 uint32 = 31
                var v44 int16 = 1
                var v45 uint64 = 0
                _ = fi42
                _ = v43
                _ = v44
                _ = v45
            }
            v46 := GenS3{F0: "hello"}
            var v47 bool = (!false)
            _ = v46
            _ = v47
        } else {
            v48 := GenS4{F0: 65535, F1: 255}
            _ = len(v1)
            _ = v48
        }
        _ = fi39
        _ = v40
        _ = v41
    }
    v6 = (v6 * float64((-2.25 + (float32(((0.1 / 1.5) - v6)) / 1e10))))
    v49 := GenS3{F0: ("hello" + "foo bar")}
    fmt.Println(v5)
    var v50 int16 = 1
    v2 = append(v2, -2147483648)
    v38 -= (v38 >> ((825) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
