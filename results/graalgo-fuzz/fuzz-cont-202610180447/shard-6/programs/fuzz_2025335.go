package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int; F1 uint8; F2 uint32 }

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = 1.5
    v3 := make(map[int]string)
    v4 := make(map[int]string)
    v5 := []uint32{}
    var v6 int32 = 42
    for fi7 := 0; fi7 < 16; fi7++ {
        fmt.Println(v6, v2)
        v8 := GenS1{F0: -2147483648, F1: 1}
        _ = fi7
        _ = v8
    }
    var v9 float64 = 0.0
    if true {
        if false {
            switch -927 {
            case 4294967295:
                v10 := GenS2{F0: 42, F1: 144, F2: 0}
                _ = v10
            default:
                for fi11 := 0; fi11 < 11; fi11++ {
                    var v12 float32 = ((1.5 / ((v2 - -2.25) - v2)) - 3.14159)
                    var v13 uint8 = 1
                    switch (-(fi11)) {
                    case 136:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            }
        }
    } else {
        for fi14 := 0; fi14 < 8; fi14++ {
            v9 -= 3.14159
            _ = fi14
        }
        if (4294967295 < 1) {
            switch 128 {
            case 575:
                if (-9223372036854775808 != 1) {
                    switch "foo bar" {
                    case "foo bar":
                        v5 = append(v5, 42)
                    }
                    v9 = v9
                } else {
                    var v15 float64 = -2.25
                    _ = v15
                }
            case -2147483648:
                var v16 uint32 = 65535
                _ = v16
            default:
                v17 := GenS1{F0: v6, F1: 42}
                _ = v17
            }
            v9 = 1.5
        }
        switch "foo bar" {
        case "\n\t":
            if (1 > 65535) {
                for fi18 := 0; fi18 < 7; fi18++ {
                    var v19 string = "a"
                    fmt.Println(v19, fi18, fi18)
                    _ = fi18
                    _ = v19
                }
                switch ("" + "hello") {
                case "hello":
                    if ((-33 == 42) && true) {
                        _ = len(v3)
                        _ = len(v5)
                        v3[-456] = ""
                    } else {
                        v2 = v2
                    }
                }
                for fi20 := 0; fi20 < 16; fi20++ {
                    v21 := GenS1{F0: (v6 ^ (v6 << ((49) & 63))), F1: 128}
                    _ = fi20
                    _ = v21
                }
            }
        case "hello":
            fmt.Println(v6, v2)
        case "a":
            if (!false) {
                switch -2147483648 {
                case 414:
                    var v22 uint64 = 1
                    _ = v22
                case 991:
                    v9 -= v9
                case -1:
                    var v23 bool = (v6 != 1)
                    _ = v23
                default:
                    var v24 int64 = 905
                    _ = v24
                }
                var v25 uint16 = 32767
                _ = v25
            } else {
                var v26 int32 = (v6 / (((-(v6))) | 1))
                switch "hello" {
                case "hello":
                    v9 = v9
                case "a":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                case "foo bar":
                    v9 += float64(((float32(0.0) * v2) + (v2 - 3.14159)))
                }
                _ = v26
            }
        default:
            fmt.Println(v2, v6, v9)
        }
    }
    fmt.Println(v2, v9, v6)
    if false {
        var v27 float64 = 3.14159
        for fi28 := 0; fi28 < 14; fi28++ {
            switch "" {
            case "hello":
                var v29 uint64 = 948
                _ = v29
            default:
                for fi30 := 0; fi30 < 16; fi30++ {
                    if len(v5) > 0 {
                        v5[0] = 0
                    }
                    _ = fi30
                }
            }
            if false {
                var v31 uint64 = 922
                v9 = ((v9 / v9) * (v27 / float64(v2)))
                var v32 int16 = -32768
                _ = v31
                _ = v32
            } else {
                fmt.Println(v6, v6, v27)
            }
            var v33 int16 = -1
            _ = fi28
            _ = v33
        }
        _ = v27
    } else {
        var v34 uint8 = 42
        var v35 uint8 = 0
        _ = v34
        _ = v35
    }
    var v36 float32 = (1e10 - v2)
    var v37 int16 = 0
    for fi38 := 0; fi38 < 18; fi38++ {
        _ = len(v5)
        _ = fi38
    }
    if true {
        var v39 bool = ("" > "")
        var v40 int16 = (v37 - -205)
        _ = v39
        _ = v40
    }
    fmt.Println(v9, v36)
    var v41 float32 = float32((0.0 / (v9 + v9)))
    var v42 uint16 = 311
    var v43 uint32 = 4294967295
    v41 -= 1.5
    v44 := GenS2{F0: -316, F1: 0, F2: v43}
    v43 -= 272
    v36 = (-2.25 / v2)
    v45 := GenS2{F0: -2147483648, F1: 127, F2: v43}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
    _ = v9
}
