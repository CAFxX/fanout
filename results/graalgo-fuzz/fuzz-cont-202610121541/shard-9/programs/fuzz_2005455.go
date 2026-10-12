package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint64 }
func genFn0(p0 uint16) uint32 {
    p0 = p0
    var v6 float32 = (-2.25 * 1e10)
    _ = p0
    _ = v6
    return 10
}
func genFn1(p0 int16, p1 int64) uint16 {
    var v7 string = "a"
    v8 := []int{32767}
    if (false || (!false)) {
        p1 = (-(p1))
        fmt.Println(v7, p1)
        v9 := GenS1{F0: p0, F1: 65535}
        _ = v9
    } else {
        p1 = p1
        v7 = (v7 + v7)
        var v10 int32 = 0
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 32767
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]string)
    var v3 uint8 = 42
    var v4 int = 1
    var v5 string = "\n\t"
    var v11 uint = 2
    v4 -= v4
    for fi12 := 0; fi12 < 14; fi12++ {
        v11 = (v11 << ((v11) & 63))
        _ = fi12
    }
    if false {
        v4 += v4
    }
    v4 += (v4 | v4)
    fmt.Println(v3)
    var v13 int8 = -76
    v14 := GenS2{F0: genFn0(32767)}
    v15 := GenS1{F0: -10, F1: 2}
    if (!(9223372036854775807 == 0)) {
        var v16 int = (^(v4))
        _ = v16
    } else {
        v11 = (v11 % ((v11) | 1))
        delete(v2, v5)
        v5 = "foo bar"
    }
    switch (v5 + "") {
    case "":
        v3 -= (v3 - (v3 + v3))
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        fmt.Println(v5)
        for fi18 := 0; fi18 < 17; fi18++ {
            v4 += (v4 % ((fi17) | 1))
            _ = fi18
        }
        v19 := GenS1{F0: 204, F1: genFn1(-901, 42)}
        _ = fi17
        _ = v19
    }
    var v20 uint16 = 32767
    if (v5 > v5) {
        var v21 uint8 = v3
        for fi22 := 0; fi22 < 19; fi22++ {
            if (v11 == (v11 | (v11 / ((v11) | 1)))) {
                var v23 float64 = float64(0.0)
                _ = v23
            } else {
                if false {
                    for fi24 := 0; fi24 < 7; fi24++ {
                        if true {
                            var v25 bool = false
                            var v26 string = v5
                            v5 = v26
                            _ = v25
                            _ = v26
                        }
                        var v27 float64 = float64(1.5)
                        v28 := []int{(fi22 | v4), (fi22 + v4)}
                        _ = fi24
                        _ = v27
                        _ = v28
                    }
                }
                _ = len(v1)
                for fi29 := 0; fi29 < 10; fi29++ {
                    if false {
                        switch 4294967295 {
                        case -2147483648:
                            _ = len(v2)
                        case -236:
                            v30 := GenS2{F0: 2}
                            _ = v30
                        }
                        v31 := GenS3{F0: 42, F1: fi29}
                        var v32 bool = (1.5 >= -2.25)
                        v13 = v13
                        if (v32 || ((1 == 4294967295) && (749 == 1))) {
                            v5 = v5
                            if (v5 == (v5 + "a")) {
                                var v33 float64 = float64(3.14159)
                                _ = v33
                            } else {
                                if (!v32) {
                                    var v34 uint64 = 928
                                    _ = v34
                                }
                                var v35 float32 = 1e10
                                var v36 int64 = -1
                                _ = v35
                                _ = v36
                            }
                        }
                        var v37 uint = (v11 >> ((673) & 63))
                        _ = v31
                        _ = v32
                        _ = v37
                    }
                    _ = fi29
                }
            }
            if true {
                v2[v5] = (("foo bar" + "\n\t") + (v5 + v5))
            }
            _ = fi22
        }
        _ = v21
    } else {
        delete(v1, 2147483648)
    }
    for fi38 := 0; fi38 < 18; fi38++ {
        switch v4 {
        case -9223372036854775808:
            v13 = (v13 + (v13 ^ (-(v13))))
        case 1:
            v20 = v20
        case -152:
            if false {
                for fi39 := 0; fi39 < 2; fi39++ {
                    if true {
                        if (!(879 <= 18446744073709551615)) {
                            switch v5 {
                            case "hello":
                                var v40 int64 = 498
                                _ = v40
                            case "":
                                fmt.Println(v3)
                            case "\n\t":
                                fmt.Println(v3, v5, fi38)
                            }
                        } else {
                            v20 -= (v20 / ((v20) | 1))
                            if (!true) {
                                fmt.Println(v5)
                                fmt.Println(v20, v4, v3)
                            } else {
                                switch "\n\t" {
                                case "":
                                    var v41 uint = (v11 >> ((976) & 63))
                                    _ = v41
                                }
                                var v42 uint8 = 128
                                _ = v42
                            }
                            v43 := GenS4{F0: 1}
                            _ = v43
                        }
                        if (-567 > 1) {
                            v3 += v3
                        } else {
                            v44 := v1[65535]
                            v45 := GenS5{F0: 42}
                            _ = v44
                            _ = v45
                        }
                    }
                    fmt.Println(v20)
                    _ = fi39
                }
            }
        default:
            var v46 int16 = int16(v11)
            _ = v46
        }
        v3 = 158
        _ = fi38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
