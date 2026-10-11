package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 uint16; F1 uint16; F2 int64 }
type GenS4 struct { F0 string; F1 int16; F2 uint16 }
type GenS5 struct { F0 int32; F1 bool }
type GenS6 struct { F0 uint; F1 int32; F2 int }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 string }

func main() {
    var v1 bool = (254 < 0)
    var v2 int8 = -127
    var v3 uint = 18446744073709551615
    v3 -= (v3 + (v3 % ((v3) | 1)))
    v2 -= -33
    v1 = (!(3.14159 != 1.5))
    v3 = v3
    var v4 int8 = v2
    v5 := []int{1, 42}
    if (-2.25 == (((float32(-2.25) + 3.14159) - float32((float64((-2.25 - 0.0)) + (0.0 + (1e10 - 0.1))))) / float32((3.14159 + 3.14159)))) {
        switch "" {
        case "foo bar":
            switch 806 {
            case 0:
                v6 := GenS1{F0: "\n\t"}
                _ = v6
            case -2147483648:
                for fi7 := 0; fi7 < 7; fi7++ {
                    v8 := GenS2{F0: fi7, F1: 2}
                    _ = fi7
                    _ = v8
                }
            case 9223372036854775807:
                var v9 uint32 = 65535
                _ = v9
            }
        }
        v10 := GenS3{F0: 716, F1: 569, F2: 1}
        _ = v10
    } else {
        if v1 {
            if (!v1) {
                for fi11 := 0; fi11 < 2; fi11++ {
                    for fi12 := 0; fi12 < 10; fi12++ {
                        v1 = (!((!(966 > 2)) && v1))
                        v1 = v1
                        _ = fi12
                    }
                    fmt.Println(fi11, fi11)
                    switch "a" {
                    case "\n\t":
                        v2 -= v4
                    default:
                        var v13 uint32 = 587
                        _ = v13
                    }
                    _ = fi11
                }
                var v14 int8 = (v2 * (^(v2)))
                var v15 int8 = 0
                _ = v14
                _ = v15
            }
            for fi16 := 0; fi16 < 7; fi16++ {
                v4 = (v4 * int8(v3))
                _ = fi16
            }
        }
        for fi17 := 0; fi17 < 2; fi17++ {
            for fi18 := 0; fi18 < 18; fi18++ {
                var v19 string = ("" + "a")
                v3 = (v3 | v3)
                _ = fi18
                _ = v19
            }
            if v1 {
                fmt.Println(fi17, v4)
                fmt.Println(v1, fi17)
                var v20 int = (fi17 + -429)
                _ = v20
            }
            v2 -= v4
            _ = fi17
        }
    }
    if v1 {
        if v1 {
            v21 := GenS4{F0: (("a" + "hello") + ("hello" + "foo bar")), F1: -1, F2: 379}
            fmt.Println(v1, v4)
            v22 := GenS4{F0: "foo bar", F1: -805, F2: 0}
            _ = v21
            _ = v22
        }
    }
    v1 = v1
    if (!(v1 && false)) {
        var v23 int = 0
        _ = v23
    }
    v24 := GenS2{F0: 0, F1: 2}
    var v25 uint = 0
    _ = len(v5)
    fmt.Println(v1, v25, v2)
    fmt.Println(v3)
    switch 4294967295 {
    case -2147483648:
        var v26 uint8 = 149
        _ = v26
    case -923:
        switch 32768 {
        case 0:
            v27 := GenS5{F0: 0, F1: false}
            _ = v27
        case 42:
            switch "hello" {
            case "a":
                var v28 uint32 = 65535
                _ = v28
            }
        default:
            for fi29 := 0; fi29 < 10; fi29++ {
                for fi30 := 0; fi30 < 5; fi30++ {
                    var v31 uint64 = 203
                    _ = fi30
                    _ = v31
                }
                v4 += v4
                switch "\n\t" {
                case "a":
                    for fi32 := 0; fi32 < 14; fi32++ {
                        var v33 uint8 = 255
                        _ = fi32
                        _ = v33
                    }
                case "\n\t":
                    if ((v1 || true) || true) {
                        fmt.Println(v1, v3, v4)
                    } else {
                        for fi34 := 0; fi34 < 6; fi34++ {
                            v4 = (v4 << ((v3) & 63))
                            if v1 {
                                var v35 int64 = 9223372036854775807
                                if (-1 == -2147483648) {
                                    var v36 int8 = (v4 + (v2 + (v4 + v2)))
                                    _ = v36
                                } else {
                                    v1 = v1
                                    var v37 int64 = v35
                                    _ = v37
                                }
                                var v38 uint32 = 1
                                _ = v35
                                _ = v38
                            } else {
                                fmt.Println(v25, v4, v2)
                            }
                            _ = fi34
                        }
                        switch fi29 {
                        case 2147483647:
                            for fi39 := 0; fi39 < 8; fi39++ {
                                v25 += (v3 >> ((v3) & 63))
                                v40 := GenS6{F0: (v3 << ((v3) & 63)), F1: 42, F2: (fi39 | 1)}
                                v41 := GenS4{F0: "hello", F1: 802, F2: 255}
                                v25 = v3
                                v5 = append(v5, (fi29 + 0))
                                fmt.Println(v3, v2, v2)
                                _ = fi39
                                _ = v40
                                _ = v41
                            }
                        default:
                            switch "" {
                            case "a":
                                if (1 != 65535) {
                                    v42 := GenS7{F0: v25}
                                    var v43 int16 = 32767
                                    fmt.Println(v2, v3)
                                    _ = v42
                                    _ = v43
                                } else {
                                    v2 += v2
                                    fmt.Println(v1, v3)
                                    fmt.Println(v25, v1, v3)
                                }
                            default:
                                v25 += v25
                            }
                        }
                        v44 := GenS8{F0: "\n\t"}
                        _ = v44
                    }
                }
                _ = fi29
            }
        }
    case 9223372036854775807:
        fmt.Println(v25, v1, v25)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
