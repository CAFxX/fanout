package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 uint8; F1 int16; F2 string }

func main() {
    var v1 bool = true
    v2 := make(map[string]uint32)
    var v3 int64 = -426
    var v4 uint32 = 232
    switch -684 {
    case 327:
        switch ("foo bar" + "\n\t") {
        case "a":
            v5 := GenS1{F0: 1, F1: 128}
            _ = v5
        case "foo bar":
            switch 127 {
            case 873:
                if v1 {
                    v3 -= v3
                } else {
                    v3 += (v3 << ((608) & 63))
                    var v6 float64 = 0.1
                    _ = v6
                }
            }
        default:
            fmt.Println(v1)
        }
    case 18:
        fmt.Println(v1, v3, v1)
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        fmt.Println(v4, v3)
        fmt.Println(fi7)
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 uint8 = 255
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    if v1 {
        fmt.Println(v1)
        var v10 uint16 = 42
        _ = v10
    } else {
        v4 -= (v4 % (((^(v4))) | 1))
        for fi11 := 0; fi11 < 16; fi11++ {
            v12 := []int{}
            if v1 {
                for fi13 := 0; fi13 < 15; fi13++ {
                    v3 -= v3
                    _ = fi13
                }
                if true {
                    v14 := v2[("hello" + ("hello" + ("" + "")))]
                    _ = v14
                } else {
                    v3 = (v3 << ((127) & 63))
                    if (!true) {
                        v15 := GenS2{F0: 104, F1: 32767}
                        var v16 int = -318
                        switch fi11 {
                        case -1:
                            if true {
                                v16 += 128
                                v3 = (v3 ^ (v3 - v3))
                            } else {
                                var v17 int = fi11
                                fmt.Println(v4)
                                _ = v17
                            }
                        }
                        _ = v15
                        _ = v16
                    } else {
                        var v18 int64 = (v3 << ((1) & 63))
                        if false {
                            v18 += v3
                            fmt.Println(v18, v4, v4)
                        } else {
                            if v1 {
                                if len(v12) > 0 {
                                    v12[0] = fi11
                                }
                                var v19 uint = 1
                                if v1 {
                                    var v20 uint64 = 0
                                    _ = v20
                                }
                                _ = v19
                            }
                            if len(v12) > 0 {
                                v12[0] = fi11
                            }
                        }
                        v2["a"] = 0
                        _ = v18
                    }
                    v21 := GenS2{F0: 2, F1: 1}
                    _ = v21
                }
                v22 := GenS3{F0: 128, F1: 32767, F2: ((("" + "foo bar") + "\n\t") + ("" + "hello"))}
                _ = v22
            } else {
                fmt.Println(v3, v3)
            }
            _ = fi11
            _ = v12
        }
    }
    fmt.Println(v3)
    v23 := GenS3{F0: 127, F1: 766, F2: "foo bar"}
    for fi24 := 0; fi24 < 6; fi24++ {
        var v25 float64 = 1.5
        _ = fi24
        _ = v25
    }
    var v26 uint = 2
    var v27 uint32 = v4
    var v28 int8 = 124
    for fi29 := 0; fi29 < 1; fi29++ {
        v30 := GenS3{F0: 128, F1: -1, F2: "hello"}
        _ = fi29
        _ = v30
    }
    v31 := []int{127, -772, 127}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
}
