package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int }
func genFn0() uint64 {
    var v6 uint16 = 135
    if true {
        v6 = 65535
        if true {
            switch "" {
            case "":
                if (!true) {
                    fmt.Println(v6)
                } else {
                    if (927 == 23) {
                        if false {
                            v7 := GenS1{F0: 509}
                            v8 := []int{}
                            _ = v7
                            _ = v8
                        } else {
                            switch ("a" + "\n\t") {
                            case "foo bar":
                                if ((2147483647 >= 2) && (((^(v6)) <= v6) || false)) {
                                    v9 := []int{4294967295}
                                    v6 -= (v6 / (((v6 * 0)) | 1))
                                    _ = v9
                                }
                            case "a":
                                fmt.Println(v6)
                            case "\n\t":
                                for fi10 := 0; fi10 < 9; fi10++ {
                                    v11 := GenS1{F0: -1}
                                    var v12 uint = 18446744073709551615
                                    var v13 bool = (!true)
                                    _ = fi10
                                    _ = v11
                                    _ = v12
                                    _ = v13
                                }
                            }
                            var v14 uint64 = 2
                            for fi15 := 0; fi15 < 10; fi15++ {
                                var v16 float32 = 1e10
                                _ = fi15
                                _ = v16
                            }
                            _ = v14
                        }
                        fmt.Println(v6)
                    }
                }
            case "\n\t":
                fmt.Println(v6)
            default:
                v17 := []int{}
                _ = v17
            }
            v6 -= 0
        } else {
            if (!("a" >= (("hello" + ("foo bar" + "\n\t")) + (("foo bar" + "") + "")))) {
                if (-391 == 42) {
                    if (2 > 769) {
                        var v18 uint = 255
                        _ = v18
                    }
                    v19 := GenS2{F0: 42, F1: 717}
                    _ = v19
                } else {
                    for fi20 := 0; fi20 < 8; fi20++ {
                        var v21 uint = 0
                        var v22 uint16 = v6
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                }
                var v23 uint8 = 153
                _ = v23
            } else {
                var v24 int = 694
                v6 = (v6 >> ((1) & 63))
                _ = v24
            }
            var v25 float64 = (1.5 - 3.14159)
            _ = v25
        }
    } else {
        fmt.Println(v6)
        if (!(493 <= 1)) {
            v6 -= v6
            var v26 int64 = 42
            v27 := []int{-729, 127, 4294967295}
            _ = v26
            _ = v27
        } else {
            if (1 <= 537) {
                v6 = (v6 / (((v6 + v6)) | 1))
            }
            var v28 int = 2147483647
            _ = v28
        }
    }
    v6 -= v6
    _ = v6
    return 42
}
func genFn1(p0 uint8) uint {
    p0 = 255
    for fi29 := 0; fi29 < 5; fi29++ {
        v30 := []int{fi29, fi29, (fi29 ^ fi29)}
        for fi31 := 0; fi31 < 8; fi31++ {
            var v32 float32 = 3.14159
            p0 = (p0 + p0)
            if len(v30) >= 2 {
                v30 = v30[1:len(v30)]
            }
            if len(v30) >= 2 {
                v30 = v30[1:len(v30)]
            }
            fmt.Println(fi31, fi29, p0)
            _ = fi31
            _ = v32
        }
        _ = fi29
        _ = v30
    }
    _ = p0
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]string)
    var v3 int = -549
    var v4 float32 = 1.5
    var v5 bool = false
    fmt.Println(v4)
    v33 := v1[478]
    fmt.Println(v4, v3, v5)
    for fi34 := 0; fi34 < 15; fi34++ {
        if v5 {
            var v35 float64 = 0.1
            _ = v35
        } else {
            v33 -= 65535
            var v36 int16 = 1000
            v3 = (v3 - fi34)
            _ = v36
        }
        _ = fi34
    }
    fmt.Println(v3, v3, v4)
    for fi37 := 0; fi37 < 17; fi37++ {
        var v38 float64 = 0.0
        v39 := GenS3{F0: 0}
        if false {
            fmt.Println(v5, v4)
            for fi40 := 0; fi40 < 17; fi40++ {
                v2["a"] = ((("" + "\n\t") + "a") + "hello")
                v3 = fi40
                v41 := v1[(v33 % ((v33) | 1))]
                _ = fi40
                _ = v41
            }
            for fi42 := 0; fi42 < 4; fi42++ {
                v43 := GenS4{F0: v3}
                _ = fi42
                _ = v43
            }
        }
        _ = fi37
        _ = v38
        _ = v39
    }
    for fi44 := 0; fi44 < 16; fi44++ {
        var v45 float64 = (0.1 * 0.1)
        _ = len(v2)
        _ = fi44
        _ = v45
    }
    var v46 uint64 = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v46
    _ = v5
}
