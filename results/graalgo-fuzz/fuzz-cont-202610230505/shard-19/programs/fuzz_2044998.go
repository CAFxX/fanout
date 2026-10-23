package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint }
type GenS2 struct { F0 int32; F1 int64; F2 int16 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]uint32)
    var v3 uint = 255
    var v4 uint = (v3 / ((v3) | 1))
    for fi5 := 0; fi5 < 17; fi5++ {
        for fi6 := 0; fi6 < 13; fi6++ {
            v7 := v2["a"]
            v8 := GenS1{F0: (fi6 | (fi5 + fi5)), F1: uint64(v4), F2: (v3 & (^(v3)))}
            v9 := []int{-2147483648, (-(fi5))}
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
        v10 := GenS1{F0: (fi5 | -9223372036854775808), F1: 42, F2: v4}
        _ = fi5
        _ = v10
    }
    var v11 int8 = -128
    v4 -= (^(v3))
    fmt.Println(v11, v3)
    delete(v1, 676)
    var v12 int16 = 1000
    var v13 uint32 = 995
    v4 += 2
    var v14 uint32 = v13
    var v15 float64 = (1e10 * -2.25)
    v16 := GenS1{F0: int(v11), F1: 302, F2: uint(v13)}
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 int = -2147483648
        _ = fi17
        _ = v18
    }
    if (!false) {
        for fi19 := 0; fi19 < 6; fi19++ {
            v3 = v3
            _ = fi19
        }
    }
    switch "hello" {
    case "\n\t":
        var v20 string = "foo bar"
        _ = v20
    case "a":
        if false {
            switch 482 {
            case -790:
                v21 := GenS1{F0: -368, F1: 0, F2: v3}
                _ = v21
            default:
                for fi22 := 0; fi22 < 5; fi22++ {
                    v23 := []int{(^(fi22))}
                    if (42 != 0) {
                        v13 += v14
                        if (true || true) {
                            for fi24 := 0; fi24 < 2; fi24++ {
                                var v25 uint64 = 888
                                v12 += (-(v12))
                                v11 -= v11
                                v12 -= (-(v12))
                                var v26 int16 = v12
                                _ = fi24
                                _ = v25
                                _ = v26
                            }
                            v11 += v11
                            if true {
                                fmt.Println(v4, v15, v14)
                                var v27 int = fi22
                                var v28 uint8 = 128
                                _ = v27
                                _ = v28
                            } else {
                                fmt.Println(fi22, v13)
                            }
                        } else {
                            var v29 uint = (^(v4))
                            _ = v29
                        }
                    } else {
                        switch (-(fi22)) {
                        case -2147483648:
                            if (!(-547 != fi22)) {
                                fmt.Println(fi22)
                                var v30 int64 = int64(v13)
                                v4 += (v4 % ((v4) | 1))
                                _ = v30
                            } else {
                                v3 += 759
                            }
                        case 42:
                            for fi31 := 0; fi31 < 8; fi31++ {
                                var v32 float64 = 1.5
                                if (1 <= (v4 & (v4 - v4))) {
                                    var v33 int = fi31
                                    _ = v33
                                } else {
                                    v34 := GenS1{F0: fi22, F1: 2, F2: (^(v4))}
                                    if false {
                                        if (v15 > ((v32 / float64(1.5)) + (((0.0 + v15) - (v15 / 1.5)) * v15))) {
                                            fmt.Println(v4)
                                            var v35 uint32 = 760
                                            _ = v35
                                        }
                                        var v36 bool = true
                                        switch (^(fi31)) {
                                        case 9223372036854775807:
                                            fmt.Println(v36, v12)
                                        default:
                                            var v37 string = "\n\t"
                                            _ = v37
                                        }
                                        if (!true) {
                                            v38 := GenS1{F0: fi22, F1: 1, F2: (^(v4))}
                                            _ = v38
                                        } else {
                                            var v39 int16 = v12
                                            _ = v39
                                        }
                                        _ = v36
                                    } else {
                                        var v40 uint = v3
                                        _ = v40
                                    }
                                    _ = v34
                                }
                                _ = fi31
                                _ = v32
                            }
                        }
                        v41 := GenS1{F0: fi22, F1: 2, F2: (v3 & v4)}
                        _ = v41
                    }
                    _ = fi22
                    _ = v23
                }
            }
            v42 := GenS2{F0: 0, F1: -116, F2: v12}
            _ = v42
        }
    default:
        v13 -= v13
    }
    if false {
        var v43 int = 32767
        for fi44 := 0; fi44 < 8; fi44++ {
            switch (-(fi44)) {
            case 360:
                v45 := GenS1{F0: 42, F1: 2, F2: v3}
                _ = v45
            }
            _ = fi44
        }
        _ = v43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
