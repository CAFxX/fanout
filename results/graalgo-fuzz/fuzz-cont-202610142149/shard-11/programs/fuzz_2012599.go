package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 int16; F1 uint64 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 411
    var v3 float64 = 1.5
    var v4 uint = 4294967295
    var v5 int64 = -929
    v6 := make(map[string]string)
    fmt.Println(v4, v4)
    for fi7 := 0; fi7 < 6; fi7++ {
        switch 167 {
        case 2147483647:
            var v8 uint16 = v2
            _ = v8
        default:
            var v9 int32 = -1
            _ = v9
        }
        _ = fi7
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        v5 = -1
        _ = fi10
    }
    switch 4294967295 {
    case 371:
        for fi11 := 0; fi11 < 14; fi11++ {
            v12 := GenS1{F0: 0, F1: 42}
            _ = fi11
            _ = v12
        }
    default:
        v13 := GenS1{F0: -568, F1: -320}
        _ = v13
    }
    var v14 uint8 = 0
    if (true && (-2.25 <= 0.0)) {
        if (-2.25 == 1.5) {
            var v15 string = "foo bar"
            switch v15 {
            case "hello":
                v4 -= v4
            case "":
                var v16 int32 = 652
                _ = v16
            }
            v3 = v3
            _ = v15
        }
        fmt.Println(v2)
        for fi17 := 0; fi17 < 17; fi17++ {
            var v18 int16 = int16(v2)
            v19 := []int{(fi17 >> (((v4 - v4)) & 63))}
            v5 += v5
            _ = fi17
            _ = v18
            _ = v19
        }
    } else {
        v14 = v14
        var v20 float32 = (0.1 - (1.5 * 1e10))
        fmt.Println(v20, v14)
        _ = v20
    }
    v14 += (v14 % ((41) | 1))
    switch -619 {
    case 32767:
        v5 = v5
    case -844:
        v21 := GenS1{F0: -1, F1: 338}
        _ = v21
    }
    if (v2 < 399) {
        for fi22 := 0; fi22 < 15; fi22++ {
            switch (fi22 | fi22) {
            case -145:
                if ("" <= ((("" + "a") + "\n\t") + "foo bar")) {
                    v23 := GenS1{F0: int16(v2), F1: 1}
                    var v24 int = (fi22 >> (((v4 - v4)) & 63))
                    fmt.Println(v3)
                    _ = v23
                    _ = v24
                } else {
                    v25 := v1[2147483648]
                    _ = v25
                }
            case 32767:
                fmt.Println(v2)
            }
            if (715 == 2) {
                if true {
                    var v26 int64 = (v5 << ((1) & 63))
                    fmt.Println(v3, v5)
                    _ = v26
                }
                var v27 bool = true
                v27 = (970 != -1)
                _ = v27
            }
            var v28 bool = true
            _ = fi22
            _ = v28
        }
    }
    for fi29 := 0; fi29 < 14; fi29++ {
        var v30 uint32 = 65535
        _ = fi29
        _ = v30
    }
    switch -400 {
    case 419:
        _ = len(v6)
    case 773:
        switch "\n\t" {
        case "":
            v31 := []int{9223372036854775807, -57, 32768}
            _ = v31
        case "\n\t":
            v32 := []int{1, 553, 127, -1}
            _ = v32
        case "foo bar":
            v33 := []int{-987}
            _ = v33
        default:
            if true {
                if false {
                    v14 = v14
                    var v34 int8 = -1
                    _ = v34
                }
                if false {
                    if (v4 != (^(v4))) {
                        var v35 uint8 = v14
                        v36 := GenS2{F0: 42, F1: 42}
                        switch 2147483647 {
                        case 127:
                            v5 = v5
                        case -991:
                            if (0.1 != -2.25) {
                                v37 := []int{}
                                _ = v37
                            } else {
                                if ("\n\t" == (("\n\t" + "\n\t") + "foo bar")) {
                                    v2 = v2
                                    v5 -= int64(v2)
                                }
                            }
                        }
                        _ = v35
                        _ = v36
                    }
                }
                var v38 uint64 = 42
                _ = v38
            }
        }
    }
    v39 := []int{-9223372036854775808}
    var v40 int16 = -165
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
