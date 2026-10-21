package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 string }
type GenS2 struct { F0 bool; F1 uint; F2 int16 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 uint; F1 uint8; F2 uint8 }
type GenS5 struct { F0 uint16; F1 int16; F2 uint64 }

func main() {
    var v1 uint32 = 65535
    v2 := make(map[string]string)
    var v3 string = ""
    var v4 uint8 = 0
    v5 := make(map[string]string)
    var v6 uint16 = 0
    for fi7 := 0; fi7 < 9; fi7++ {
        if (1.5 < 0.1) {
            for fi8 := 0; fi8 < 17; fi8++ {
                if true {
                    v5[v3] = ("hello" + "")
                    switch v3 {
                    case "":
                        v4 = v4
                    default:
                        switch (fi8 << ((42) & 63)) {
                        case -9223372036854775808:
                            v9 := GenS1{F0: -185, F1: true, F2: v3}
                            _ = v9
                        case 32768:
                            if true {
                                if ((0.0 / 1e10) != 1e10) {
                                    fmt.Println(v1, v6, fi8)
                                    v10 := []int{fi8, (fi7 << ((245) & 63)), fi8}
                                    _ = v10
                                } else {
                                    var v11 string = ("foo bar" + v3)
                                    _ = v11
                                }
                                v12 := GenS1{F0: -1, F1: false, F2: ("foo bar" + ((v3 + v3) + v3))}
                                _ = v12
                            } else {
                                if true {
                                    var v13 int16 = -32768
                                    switch v3 {
                                    case "foo bar":
                                        v6 = (v6 | (v6 + 1))
                                    default:
                                        switch v3 {
                                        case "a":
                                            fmt.Println(v4, v13, v13)
                                        case "foo bar":
                                            v14 := GenS2{F0: true, F1: 2, F2: v13}
                                            _ = v14
                                        default:
                                            v15 := GenS2{F0: false, F1: uint(v4), F2: (v13 * (v13 % ((v13) | 1)))}
                                            _ = v15
                                        }
                                    }
                                    _ = v13
                                } else {
                                    v16 := []int{fi7, 32767}
                                    _ = v16
                                }
                                v1 += (v1 | (^(v1)))
                            }
                        }
                    }
                    var v17 int16 = 329
                    _ = v17
                }
                v6 = (v6 + (v6 * v6))
                switch "foo bar" {
                case "foo bar":
                    if false {
                        var v18 float32 = float32(float64((-2.25 * float32(-2.25))))
                        _ = v18
                    } else {
                        var v19 float32 = 0.0
                        _ = v19
                    }
                default:
                    fmt.Println(v6, v1, v1)
                }
                _ = fi8
            }
            v1 += (^(v1))
        }
        v20 := []int{(fi7 % ((42) | 1)), fi7, (^(fi7)), fi7}
        _ = fi7
        _ = v20
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        if (("foo bar" + v3) >= "a") {
            switch v3 {
            case "foo bar":
                if (v6 <= v6) {
                    var v22 bool = (!(0 >= 42))
                    switch fi21 {
                    case -9223372036854775808:
                        for fi23 := 0; fi23 < 6; fi23++ {
                            v1 -= v1
                            if (646 >= 1) {
                                v22 = false
                                v1 -= 2
                            }
                            v22 = (1 == 308)
                            _ = fi23
                        }
                    }
                    v24 := GenS3{F0: -1, F1: -942}
                    _ = v22
                    _ = v24
                } else {
                    v1 += v1
                    var v25 int16 = 42
                    v26 := []int{}
                    _ = v25
                    _ = v26
                }
            case "a":
                v27 := v5[(v3 + v3)]
                _ = v27
            case "\n\t":
                v4 = v4
            }
            v28 := []int{(^(fi21)), fi21, fi21, fi21}
            v29 := GenS4{F0: 414, F1: v4, F2: (v4 | (^(v4)))}
            _ = v28
            _ = v29
        } else {
            var v30 int64 = 1
            _ = v30
        }
        var v31 uint64 = 900
        _ = fi21
        _ = v31
    }
    v32 := GenS4{F0: 64, F1: v4, F2: v4}
    switch 9223372036854775807 {
    case 517:
        var v33 int32 = 43
        _ = v33
    case -643:
        var v34 uint = 255
        _ = v34
    case 512:
        v1 += v1
    }
    _ = len(v2)
    var v35 uint64 = 18446744073709551615
    var v36 int64 = 42
    v37 := []int{32767, 128, -1}
    v1 += v1
    v38 := GenS3{F0: 36, F1: 42}
    v39 := v2[(v3 + v3)]
    var v40 uint16 = (^(v6))
    _ = len(v5)
    v6 = 255
    if true {
        v41 := GenS5{F0: v40, F1: -32768, F2: v35}
        v3 = v39
        _ = v41
    } else {
        var v42 uint64 = 2
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
