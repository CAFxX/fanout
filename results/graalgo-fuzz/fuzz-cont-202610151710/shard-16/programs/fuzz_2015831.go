package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 string; F1 uint; F2 uint32 }
type GenS4 struct { F0 string; F1 int64; F2 string }
func genFn0(p0 int) int32 {
    p0 = p0
    for fi6 := 0; fi6 < 7; fi6++ {
        for fi7 := 0; fi7 < 16; fi7++ {
            var v8 int8 = 100
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    var v9 uint16 = 32767
    p0 += (-(p0))
    _ = p0
    _ = v9
    return 0
}

func main() {
    var v1 uint = 0
    var v2 int16 = -1
    v3 := make(map[uint32]int)
    var v4 string = "hello"
    var v5 float64 = 0.1
    v10 := GenS1{F0: 42, F1: 2147483648}
    v2 = (v2 + 654)
    for fi11 := 0; fi11 < 16; fi11++ {
        v3[466] = -545
        _ = fi11
    }
    var v12 uint64 = 666
    v13 := GenS2{F0: (v12 | v12), F1: (v1 << ((v1) & 63))}
    v2 -= (^(v2))
    if (!true) {
        fmt.Println(v12)
    } else {
        fmt.Println(v4, v12)
        var v14 uint = (v1 | v1)
        _ = v14
    }
    v2 += (-(v2))
    v1 = (v1 ^ 0)
    if (v5 == 0.1) {
        if (((v1 - v1) <= (v1 * v1)) || false) {
            switch 1 {
            case 763:
                delete(v3, 2)
            case 127:
                var v15 int64 = 9223372036854775807
                _ = v15
            default:
                v16 := v3[65535]
                _ = v16
            }
            v17 := GenS1{F0: 1, F1: 2147483647}
            var v18 bool = (!true)
            _ = v17
            _ = v18
        } else {
            v19 := GenS2{F0: (v12 / ((711) | 1)), F1: v1}
            fmt.Println(v12)
            _ = v19
        }
        if true {
            for fi20 := 0; fi20 < 18; fi20++ {
                fmt.Println(v1, v1)
                for fi21 := 0; fi21 < 2; fi21++ {
                    var v22 string = v4
                    v23 := GenS1{F0: v1, F1: 4294967295}
                    switch ((v22 + ((v22 + "") + "hello")) + "foo bar") {
                    case "":
                        var v24 int32 = -2147483648
                        _ = v24
                    default:
                        v5 = float64(1e10)
                    }
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                if true {
                    fmt.Println(v1, v5)
                    for fi25 := 0; fi25 < 10; fi25++ {
                        if false {
                            var v26 uint = v1
                            fmt.Println(v1)
                            switch (fi20 & (-(fi25))) {
                            case 32768:
                                v27 := GenS3{F0: v4, F1: (v26 >> (((^(v26))) & 63)), F2: 65535}
                                _ = v27
                            default:
                                var v28 uint32 = 348
                                _ = v28
                            }
                            _ = v26
                        }
                        fmt.Println(fi20, v1, fi20)
                        v29 := GenS2{F0: 2, F1: 37}
                        _ = fi25
                        _ = v29
                    }
                    for fi30 := 0; fi30 < 11; fi30++ {
                        fmt.Println(v1, v5)
                        v2 = v2
                        var v31 uint8 = 127
                        _ = fi30
                        _ = v31
                    }
                } else {
                    v32 := []int{(fi20 & (fi20 << ((v1) & 63)))}
                    _ = v32
                }
                _ = fi20
            }
        }
    } else {
        switch v4 {
        case "":
            v33 := GenS2{F0: (v12 << (((v1 ^ 42)) & 63)), F1: 0}
            _ = v33
        case "foo bar":
            for fi34 := 0; fi34 < 6; fi34++ {
                v35 := v3[0]
                v36 := GenS4{F0: "\n\t", F1: -1, F2: v4}
                var v37 uint = (v1 / (((v1 + v1)) | 1))
                _ = fi34
                _ = v35
                _ = v36
                _ = v37
            }
        case "a":
            var v38 string = v4
            _ = v38
        default:
            if true {
                var v39 bool = (v4 < ("foo bar" + v4))
                fmt.Println(v4, v5)
                v4 = v4
                _ = v39
            } else {
                if true {
                    v40 := []int{}
                    var v41 uint32 = 448
                    _ = v40
                    _ = v41
                }
                v12 = v12
            }
        }
        v42 := []int{2147483647}
        switch ((v4 + "\n\t") + v4) {
        case "hello":
            v42 = append(v42, 1)
        case "\n\t":
            if false {
                var v43 uint32 = 333
                for fi44 := 0; fi44 < 16; fi44++ {
                    fmt.Println(fi44)
                    _ = fi44
                }
                _ = v43
            }
        default:
            v45 := GenS3{F0: (v4 + v4), F1: v1, F2: 42}
            _ = v45
        }
        _ = v42
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
