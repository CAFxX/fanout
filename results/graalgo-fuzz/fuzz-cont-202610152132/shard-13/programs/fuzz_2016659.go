package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint; F1 int16; F2 uint8 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 920
    var v3 uint8 = 1
    var v4 int8 = 0
    var v5 uint32 = 0
    var v6 float64 = 1.5
    v7 := GenS1{F0: 42, F1: 471}
    if true {
        var v8 int32 = 1
        for fi9 := 0; fi9 < 12; fi9++ {
            v10 := GenS1{F0: -1, F1: 146}
            _ = fi9
            _ = v10
        }
        v3 -= 20
        _ = v8
    } else {
        for fi11 := 0; fi11 < 1; fi11++ {
            v12 := GenS2{F0: 32768, F1: v4, F2: uint8(v5)}
            switch "a" {
            case "\n\t":
                v13 := GenS1{F0: -1, F1: 2}
                _ = v13
            case "a":
                var v14 bool = true
                _ = v14
            case "foo bar":
                delete(v1, v5)
            default:
                if true {
                    switch 9223372036854775807 {
                    case -9223372036854775808:
                        var v15 int64 = int64(v3)
                        _ = v15
                    default:
                        v16 := []int{(fi11 >> ((v2) & 63)), fi11, (fi11 ^ 1)}
                        _ = v16
                    }
                }
            }
            var v17 float32 = 0.1
            _ = fi11
            _ = v12
            _ = v17
        }
    }
    if (!(!false)) {
        v3 -= v3
        v3 = (v3 ^ v3)
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        v5 += (^(v5))
        _ = fi18
    }
    var v19 uint64 = 18446744073709551615
    if (v6 == v6) {
        for fi20 := 0; fi20 < 11; fi20++ {
            var v21 uint64 = 124
            v22 := []int{(-(fi20)), fi20, 32768, fi20}
            _ = fi20
            _ = v21
            _ = v22
        }
        var v23 uint8 = (^(v3))
        _ = v23
    }
    if false {
        v5 = (v5 + uint32(v19))
        if true {
            var v24 int64 = -9223372036854775808
            _ = v24
        } else {
            fmt.Println(v19, v19, v2)
            v2 += (v2 % (((v2 / (((^(v2))) | 1))) | 1))
            switch (("\n\t" + (("foo bar" + (("" + ("hello" + ("\n\t" + "\n\t"))) + "\n\t")) + "foo bar")) + "\n\t") {
            case "\n\t":
                fmt.Println(v3, v5, v4)
            default:
                v4 = -128
            }
        }
    } else {
        v4 = v4
    }
    v25 := GenS1{F0: 32767, F1: 1}
    var v26 float32 = (-2.25 - 1e10)
    if (!(595 <= -553)) {
        switch ("a" + "foo bar") {
        case "":
            var v27 uint64 = v19
            _ = v27
        case "\n\t":
            if true {
                var v28 uint32 = 780
                _ = v28
            } else {
                for fi29 := 0; fi29 < 13; fi29++ {
                    v30 := GenS2{F0: fi29, F1: (v4 >> ((65535) & 63)), F2: (v3 << ((v2) & 63))}
                    if true {
                        var v31 float64 = v6
                        if (49 >= (v5 | uint32(v3))) {
                            for fi32 := 0; fi32 < 10; fi32++ {
                                v19 += v19
                                var v33 uint32 = 516
                                _ = fi32
                                _ = v33
                            }
                        }
                        for fi34 := 0; fi34 < 12; fi34++ {
                            v35 := GenS3{F0: (v2 - v2), F1: -653, F2: v3}
                            v36 := GenS2{F0: (fi34 / (((fi29 << (((v2 % ((uint(v19)) | 1))) & 63))) | 1)), F1: v4, F2: (v3 + (^(v3)))}
                            v37 := []int{}
                            _ = fi34
                            _ = v35
                            _ = v36
                            _ = v37
                        }
                        _ = v31
                    } else {
                        var v38 int8 = 63
                        v38 = 0
                        _ = v38
                    }
                    _ = fi29
                    _ = v30
                }
                v2 += v2
            }
        }
    } else {
        v1[(v5 ^ (v5 * v5))] = -174
    }
    if true {
        for fi39 := 0; fi39 < 6; fi39++ {
            v40 := GenS1{F0: 0, F1: 921}
            _ = fi39
            _ = v40
        }
        v41 := []int{}
        for fi42 := 0; fi42 < 9; fi42++ {
            var v43 uint8 = 61
            switch (^(fi42)) {
            case -1:
                fmt.Println(v4, v5)
            default:
                v44 := GenS1{F0: 0, F1: uint16(v5)}
                _ = v44
            }
            switch ("a" + "hello") {
            case "\n\t":
                v45 := GenS3{F0: 935, F1: 1000, F2: (v43 ^ v3)}
                _ = v45
            case "hello":
                v46 := v1[v5]
                _ = v46
            default:
                for fi47 := 0; fi47 < 6; fi47++ {
                    var v48 float64 = ((float64(((0.0 - (v26 - 0.0)) * v26)) - -2.25) + ((v6 + 3.14159) - v6))
                    _ = fi47
                    _ = v48
                }
            }
            _ = fi42
            _ = v43
        }
        _ = v41
    } else {
        v49 := v1[(v5 >> ((42) & 63))]
        _ = v49
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
