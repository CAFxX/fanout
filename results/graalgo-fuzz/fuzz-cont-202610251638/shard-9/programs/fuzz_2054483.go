package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint16 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 int8; F2 int64 }
func genFn0(p0 uint16, p1 uint16, p2 uint64) int8 {
    var v5 string = "hello"
    var v6 float32 = (0.1 - 3.14159)
    var v7 uint64 = p2
    v8 := GenS1{F0: 394, F1: -1, F2: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 uint16, p1 int64, p2 bool) string {
    var v9 bool = p2
    var v10 int8 = -1
    var v11 float64 = 1.5
    for fi12 := 0; fi12 < 19; fi12++ {
        v13 := GenS2{F0: fi12, F1: 40}
        for fi14 := 0; fi14 < 4; fi14++ {
            v15 := []int{(fi12 & 32767), (fi12 >> ((228) & 63)), (fi14 | fi12), (fi14 >> ((18446744073709551615) & 63))}
            switch fi14 {
            case -2147483648:
                var v16 uint = 42
                _ = v16
            case -659:
                if (!(p2 || (42 == 42))) {
                    p0 = p0
                    if false {
                        if ("hello" >= (("" + "foo bar") + "")) {
                            if (("foo bar" + ("hello" + "foo bar")) != "") {
                                v17 := GenS1{F0: 573, F1: -1, F2: (p0 >> ((734) & 63))}
                                v18 := GenS2{F0: -608, F1: 2}
                                if ((fi12 < 817) || (v11 < v11)) {
                                    var v19 int16 = 0
                                    _ = len(v15)
                                    fmt.Println(fi14, p1, fi12)
                                    v20 := GenS1{F0: 408, F1: v19, F2: p0}
                                    _ = v19
                                    _ = v20
                                } else {
                                    switch ("" + "foo bar") {
                                    case "foo bar":
                                        p1 -= p1
                                        if (0 == 65535) {
                                            fmt.Println(p2)
                                            fmt.Println(p2, v11)
                                            p1 = p1
                                            var v21 uint32 = 2
                                            var v22 float32 = 1e10
                                            v22 -= v22
                                            _ = v21
                                            _ = v22
                                        }
                                    }
                                }
                                if ((v10 - v10) > -53) {
                                    var v23 int32 = 2147483647
                                    switch fi14 {
                                    case -9223372036854775808:
                                        if len(v15) >= 2 {
                                            v15 = v15[1:len(v15)]
                                        }
                                    case -2147483648:
                                        var v24 uint = 4294967295
                                        _ = v24
                                    }
                                    v23 = (v23 % ((42) | 1))
                                    _ = v23
                                } else {
                                    if len(v15) > 0 {
                                        v15[0] = (fi14 >> ((0) & 63))
                                    }
                                    if len(v15) > 0 {
                                        v15[0] = (fi12 ^ fi12)
                                    }
                                    if len(v15) >= 2 {
                                        v15 = v15[1:len(v15)]
                                    }
                                }
                                var v25 int = 42
                                var v26 uint64 = 351
                                switch "\n\t" {
                                case "hello":
                                    if ((p0 % (((p0 ^ (p0 ^ 42))) | 1)) == 613) {
                                        v9 = (!(4294967295 == 65535))
                                        v15 = append(v15, fi14)
                                    } else {
                                        var v27 float64 = float64(0.1)
                                        switch "" {
                                        case "\n\t":
                                            var v28 int32 = 42
                                            _ = v28
                                        case "":
                                            var v29 uint64 = (v26 | v26)
                                            _ = v29
                                        }
                                        _ = v27
                                    }
                                case "foo bar":
                                    var v30 int32 = 42
                                    _ = v30
                                default:
                                    var v31 uint = 4294967295
                                    _ = v31
                                }
                                v9 = true
                                _ = v17
                                _ = v18
                                _ = v25
                                _ = v26
                            }
                        }
                    } else {
                        if p2 {
                            v32 := GenS3{F0: -83}
                            _ = v32
                        } else {
                            switch ("hello" + ("foo bar" + "foo bar")) {
                            case "a":
                                if true {
                                    if v9 {
                                        var v33 uint16 = 255
                                        _ = v33
                                    }
                                    if len(v15) > 0 {
                                        v15[0] = -737
                                    }
                                } else {
                                    var v34 uint = 1
                                    _ = v34
                                }
                            case "hello":
                                v11 -= float64(1e10)
                            case "\n\t":
                                v10 = genFn0(2, (^(p0)), 1)
                            }
                        }
                        if p2 {
                            v9 = ((p0 + 42) == p0)
                        }
                        fmt.Println(fi14, p1, fi12)
                    }
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                }
            default:
                _ = len(v15)
            }
            switch "hello" {
            case "foo bar":
                v35 := GenS3{F0: 42}
                _ = v35
            case "hello":
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
            default:
                var v36 bool = v9
                _ = v36
            }
            v37 := GenS1{F0: 0, F1: -757, F2: 144}
            _ = fi14
            _ = v15
            _ = v37
        }
        _ = fi12
        _ = v13
    }
    var v38 float32 = float32(0.1)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v38
    _ = v9
    return "hello"
}

func main() {
    var v1 uint8 = 42
    v2 := make(map[uint32]bool)
    var v3 int16 = 42
    v4 := []uint32{}
    v3 = v3
    v39 := GenS3{F0: (v3 + v3)}
    switch "foo bar" {
    case "a":
        for fi40 := 0; fi40 < 1; fi40++ {
            var v41 string = ("\n\t" + "foo bar")
            _ = fi40
            _ = v41
        }
    default:
        var v42 uint64 = 18446744073709551615
        _ = v42
    }
    fmt.Println(v3)
    _ = len(v4)
    var v43 float64 = 1e10
    var v44 bool = ((false || (!true)) && (!true))
    v45 := GenS4{F0: 9223372036854775807, F1: 127, F2: 1}
    var v46 int8 = 0
    var v47 int8 = 127
    switch ("a" + ((("a" + "a") + (((("a" + "hello") + ("a" + "hello")) + "") + "foo bar")) + "a")) {
    case "\n\t":
        var v48 int = 42
        _ = v48
    default:
        for fi49 := 0; fi49 < 20; fi49++ {
            v46 -= (v47 - v46)
            _ = fi49
        }
    }
    v50 := GenS3{F0: 1000}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v50
}
