package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 int32; F2 int32 }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 uint; F1 bool; F2 int }
type GenS7 struct { F0 int8 }
func genFn0(p0 uint64, p1 bool, p2 uint64) int16 {
    v7 := GenS1{F0: 0, F1: 18446744073709551615}
    if p1 {
        var v8 int32 = 1
        p0 = 42
        _ = v8
    }
    v9 := []int{1}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v9
    return 1000
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]bool)
    v3 := []uint32{4294967295, 42}
    var v4 string = (("" + "hello") + ("a" + "hello"))
    v5 := []int{-590, 32768}
    var v6 uint = 421
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v6 = (^(v6))
    var v10 int32 = 737
    var v11 float64 = (((-2.25 - 3.14159) + 0.1) / 1.0)
    switch "hello" {
    case "foo bar":
        _ = len(v5)
    case "":
        for fi12 := 0; fi12 < 8; fi12++ {
            for fi13 := 0; fi13 < 1; fi13++ {
                switch fi12 {
                case -944:
                    var v14 uint8 = 168
                    _ = v14
                case 127:
                    _ = len(v3)
                case -1:
                    var v15 uint = v6
                    _ = v15
                default:
                    switch (fi12 * fi12) {
                    case 900:
                        v10 = 122
                        fmt.Println(v11)
                        fmt.Println(v10, v6, v6)
                    case 2147483647:
                        fmt.Println(fi13, v6)
                    default:
                        if false {
                            v4 = v4
                            v10 = -1
                        } else {
                            v16 := v1[1]
                            fmt.Println(v6)
                            fmt.Println(fi13, v4, v16)
                            _ = v16
                        }
                    }
                }
                _ = fi13
            }
            if false {
                switch v4 {
                case "a":
                    v6 -= 0
                case "foo bar":
                    fmt.Println(fi12)
                case "":
                    var v17 uint32 = 2147483647
                    _ = v17
                }
            } else {
                switch fi12 {
                case 1:
                    if false {
                        _ = len(v1)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if true {
                            for fi18 := 0; fi18 < 14; fi18++ {
                                switch "hello" {
                                case "\n\t":
                                    var v19 uint32 = 42
                                    _ = v19
                                case "":
                                    switch (v4 + v4) {
                                    case "foo bar":
                                        var v20 int64 = -1
                                        _ = v20
                                    case "hello":
                                        var v21 float64 = v11
                                        _ = v21
                                    }
                                default:
                                    v10 = 222
                                }
                                if (v11 < v11) {
                                    v6 = (^(v6))
                                    var v22 int16 = -32768
                                    if (true && (-128 <= 100)) {
                                        v23 := GenS2{F0: (v22 << (((v6 << ((756) & 63))) & 63))}
                                        _ = v23
                                    } else {
                                        v24 := GenS2{F0: genFn0(42, ((((-9223372036854775808 >= 0) && true) && false) && ((-(v10)) > v10)), 4294967295)}
                                        var v25 int = (fi18 - (fi18 ^ -1))
                                        switch (fi12 ^ 0) {
                                        case 127:
                                            if ((-(v10)) > 159) {
                                                var v26 float32 = 0.1
                                                v27 := GenS3{F0: ("foo bar" + v4), F1: 2147483647, F2: v10}
                                                _ = v26
                                                _ = v27
                                            } else {
                                                v10 = v10
                                                v28 := GenS2{F0: v22}
                                                _ = v28
                                            }
                                        case 0:
                                            var v29 uint32 = 100
                                            _ = v29
                                        default:
                                            var v30 uint = v6
                                            _ = v30
                                        }
                                        _ = v24
                                        _ = v25
                                    }
                                    _ = len(v1)
                                    var v31 int32 = v10
                                    _ = v22
                                    _ = v31
                                }
                                delete(v2, 465)
                                _ = fi18
                            }
                            var v32 uint32 = 42
                            v32 -= (v32 + (^(v32)))
                            _ = v32
                        }
                    } else {
                        var v33 uint32 = 2
                        var v34 int32 = (v10 ^ (v10 % ((42) | 1)))
                        fmt.Println(fi12)
                        _ = v33
                        _ = v34
                    }
                case 2147483647:
                    for fi35 := 0; fi35 < 20; fi35++ {
                        switch fi12 {
                        case 2147483647:
                            var v36 int32 = -2147483648
                            _ = v36
                        case -2147483648:
                            var v37 int16 = 727
                            _ = v37
                        default:
                            if true {
                                fmt.Println(fi35)
                                var v38 string = "hello"
                                switch v4 {
                                case "hello":
                                    v39 := GenS4{F0: v6, F1: 18446744073709551615}
                                    _ = v39
                                case "a":
                                    fmt.Println(fi35, v4)
                                case "\n\t":
                                    v2[4294967295] = (((v6 << ((v6) & 63)) <= v6) || true)
                                }
                                _ = v38
                            } else {
                                switch (fi12 ^ (fi35 % ((int(v10)) | 1))) {
                                case 128:
                                    var v40 uint64 = 18446744073709551615
                                    _ = v40
                                case 42:
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                            }
                        }
                        switch 2147483647 {
                        case -9223372036854775808:
                            var v41 int32 = 0
                            var v42 int = 2147483647
                            _ = v41
                            _ = v42
                        case 9223372036854775807:
                            _ = len(v1)
                        default:
                            var v43 string = (v4 + "\n\t")
                            _ = v43
                        }
                        switch (fi12 / ((fi35) | 1)) {
                        case -1:
                            var v44 uint16 = 1
                            _ = v44
                        case -2147483648:
                            v6 += v6
                        case -9223372036854775808:
                            if (!true) {
                                var v45 uint16 = 32767
                                var v46 uint32 = 65535
                                v6 = v6
                                _ = v45
                                _ = v46
                            } else {
                                v10 = int32(v6)
                                v4 = (v4 + v4)
                                var v47 int8 = 118
                                _ = v47
                            }
                        }
                        _ = fi35
                    }
                case 32767:
                    var v48 int = (fi12 * fi12)
                    _ = v48
                }
            }
            _ = fi12
        }
    }
    v10 = v10
    var v49 uint32 = 2
    v11 = v11
    v50 := GenS5{F0: (v10 | v10), F1: (-(v10))}
    v2[(v49 | (v49 - (v49 * 1)))] = true
    fmt.Println(v49, v4, v11)
    var v51 uint64 = 2
    v10 -= (v10 / ((-2147483648) | 1))
    v52 := GenS5{F0: (-(v10)), F1: (^(v10))}
    for fi53 := 0; fi53 < 19; fi53++ {
        switch 388 {
        case 326:
            var v54 uint8 = 52
            _ = v54
        case -630:
            v55 := GenS6{F0: 4294967295, F1: false, F2: (fi53 & fi53)}
            _ = v55
        default:
            v11 -= v11
        }
        delete(v1, fi53)
        v10 -= -1
        _ = fi53
    }
    v49 = 2
    v56 := GenS2{F0: 160}
    v57 := GenS7{F0: -6}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v56
    _ = v57
    _ = v6
}
