package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int16; F1 uint16; F2 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint32; F1 uint8 }
type GenS8 struct { F0 string; F1 int8; F2 uint32 }
type GenS9 struct { F0 uint64; F1 uint32; F2 int64 }
func genFn0(p0 string, p1 uint64, p2 uint) string {
    switch (p0 + "foo bar") {
    case "a":
        p1 = (^(p1))
    case "foo bar":
        v6 := []int{}
        _ = v6
    }
    switch p0 {
    case "hello":
        var v7 string = ("" + (p0 + p0))
        _ = v7
    case "a":
        var v8 int32 = 1
        _ = v8
    case "foo bar":
        p1 -= (p1 | (^(p1)))
    default:
        v9 := GenS1{F0: 0, F1: p1, F2: 48}
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return "hello"
}
func genFn1(p0 bool) int16 {
    p0 = true
    var v10 int32 = 1
    switch 128 {
    case -9223372036854775808:
        v10 = (^(v10))
    case -2147483648:
        fmt.Println(p0, v10)
    default:
        v11 := GenS2{F0: 762, F1: 361, F2: -32768}
        _ = v11
    }
    _ = p0
    _ = v10
    return 0
}

func main() {
    var v1 int64 = -1
    var v2 float64 = ((3.14159 + float64(0.0)) * 3.14159)
    var v3 string = (("a" + "") + "")
    v4 := make(map[int]bool)
    v5 := make(map[uint32]uint32)
    switch -1 {
    case -1:
        if true {
            if (!true) {
                if true {
                    v12 := GenS3{F0: (127 > 398)}
                    _ = v12
                }
                switch -2147483648 {
                case 128:
                    v3 = "\n\t"
                default:
                    v2 -= v2
                }
                fmt.Println(v3)
            } else {
                v13 := GenS4{F0: 944}
                v14 := v5[2]
                for fi15 := 0; fi15 < 4; fi15++ {
                    var v16 uint8 = 79
                    _ = fi15
                    _ = v16
                }
                _ = v13
                _ = v14
            }
            var v17 uint = 1
            _ = v17
        }
    case 9223372036854775807:
        v1 = v1
    case -151:
        v18 := []int{128}
        _ = v18
    }
    var v19 float64 = (3.14159 + v2)
    v1 = 1
    v2 = v19
    switch (v3 + v3) {
    case "a":
        switch -741 {
        case 312:
            switch 2147483647 {
            case -2147483648:
                if (!false) {
                    v20 := []int{538, 32767, 32768, 127}
                    _ = v20
                } else {
                    v21 := GenS5{F0: -9223372036854775808}
                    for fi22 := 0; fi22 < 3; fi22++ {
                        v19 += v19
                        if (1 != 65535) {
                            switch (v3 + ("a" + (v3 + v3))) {
                            case "a":
                                for fi23 := 0; fi23 < 5; fi23++ {
                                    fmt.Println(v19, fi23, v3)
                                    delete(v4, 0)
                                    _ = fi23
                                }
                            default:
                                if (0 != 2) {
                                    for fi24 := 0; fi24 < 8; fi24++ {
                                        if true {
                                            v25 := GenS6{F0: 42}
                                            v3 = v3
                                            fmt.Println(v19)
                                            _ = v25
                                        } else {
                                            var v26 int = fi24
                                            _ = v26
                                        }
                                        _ = fi24
                                    }
                                }
                            }
                            fmt.Println(v19)
                            v27 := GenS4{F0: 2}
                            _ = v27
                        } else {
                            _ = len(v4)
                            v3 = ""
                        }
                        _ = fi22
                    }
                    fmt.Println(v2, v19)
                    _ = v21
                }
            case 127:
                v28 := GenS3{F0: false}
                _ = v28
            }
        case 497:
            for fi29 := 0; fi29 < 18; fi29++ {
                v3 = v3
                switch (v3 + "a") {
                case "hello":
                    for fi30 := 0; fi30 < 17; fi30++ {
                        delete(v4, (fi29 + 128))
                        v3 = ((v3 + "hello") + (v3 + v3))
                        _ = fi30
                    }
                case "\n\t":
                    fmt.Println(v3, fi29)
                case "foo bar":
                    var v31 int = 127
                    _ = v31
                default:
                    v32 := GenS7{F0: 408, F1: 42}
                    _ = v32
                }
                _ = fi29
            }
        case 9223372036854775807:
            var v33 uint = 1
            _ = v33
        default:
            if false {
                var v34 uint8 = 2
                _ = v34
            } else {
                v3 = ("" + ((v3 + (v3 + (v3 + ""))) + ""))
                fmt.Println(v2)
            }
        }
    case "\n\t":
        var v35 float64 = (v19 - float64(1e10))
        _ = v35
    }
    v36 := GenS8{F0: (v3 + v3), F1: 43, F2: 2}
    v1 += -1
    v19 = (v2 / v2)
    var v37 int16 = 32767
    v37 -= genFn1((-128 > 100))
    if false {
        v1 += -9223372036854775808
        v2 = (float64((-2.25 * 3.14159)) / 1e10)
        fmt.Println(v1, v1)
    }
    var v38 float64 = float64(0.1)
    var v39 uint16 = 21
    if true {
        if (!(v39 == (v39 & (v39 ^ (v39 >> ((0) & 63)))))) {
            for fi40 := 0; fi40 < 4; fi40++ {
                var v41 bool = (true && ((v39 * (^(v39))) >= (v39 - v39)))
                var v42 bool = true
                v41 = v41
                _ = fi40
                _ = v41
                _ = v42
            }
            var v43 uint64 = 4294967295
            if false {
                switch 909 {
                case 1:
                    if (false && true) {
                        var v44 int16 = -1
                        v45 := GenS9{F0: v43, F1: 42, F2: v1}
                        _ = v44
                        _ = v45
                    }
                default:
                    if (((v3 + "a") < (v3 + (((v3 + (v3 + (v3 + (v3 + (v3 + (v3 + (((v3 + "hello") + ("a" + v3)) + v3))))))) + "") + v3))) && (2 <= v39)) {
                        if (!true) {
                            var v46 uint32 = 4294967295
                            var v47 uint64 = 42
                            _ = v46
                            _ = v47
                        }
                    }
                }
            }
            _ = v43
        }
        var v48 uint = 65535
        v49 := GenS5{F0: 127}
        _ = v48
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
