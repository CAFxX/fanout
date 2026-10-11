package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint8; F1 string }
type GenS6 struct { F0 int }

func main() {
    v1 := []int{128, -2147483648, -9223372036854775808}
    var v2 int16 = 0
    var v3 uint = 42
    v4 := make(map[int]uint32)
    for fi5 := 0; fi5 < 7; fi5++ {
        if (((1.5 / 0.1) == (-2.25 - 0.1)) || (float32((3.14159 * float64(1.5))) == float32(0.0))) {
            v6 := GenS1{F0: int8(v3)}
            v2 -= 0
            if (((658 == 510) || true) && (!false)) {
                v3 = 0
                var v7 uint32 = 321
                var v8 int8 = 42
                _ = v7
                _ = v8
            }
            _ = v6
        }
        var v9 float64 = -2.25
        _ = fi5
        _ = v9
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        _ = len(v1)
        _ = fi10
    }
    var v11 uint = 65535
    var v12 bool = false
    v13 := GenS2{F0: 237, F1: (^(v3))}
    switch 644 {
    case -9223372036854775808:
        v14 := GenS2{F0: 32767, F1: v3}
        _ = v14
    }
    if v12 {
        fmt.Println(v2, v11)
    } else {
        for fi15 := 0; fi15 < 5; fi15++ {
            v11 -= (^(v3))
            v16 := GenS3{F0: -128}
            _ = fi15
            _ = v16
        }
        if ((^(v2)) > v2) {
            v17 := GenS2{F0: 0, F1: (v11 * (v3 * v3))}
            fmt.Println(v11)
            var v18 int64 = 743
            _ = v17
            _ = v18
        } else {
            switch "a" {
            case "":
                switch "a" {
                case "hello":
                    var v19 float32 = 3.14159
                    _ = v19
                case "\n\t":
                    if ((!(!(((1 < int64(v2)) && (423 <= 42)) && false))) && (uint64(v11) < uint64(v11))) {
                        v20 := v4[127]
                        if v12 {
                            switch "a" {
                            case "foo bar":
                                v2 += (v2 - v2)
                            }
                        }
                        v21 := GenS3{F0: -1}
                        _ = v20
                        _ = v21
                    }
                default:
                    v11 += 65535
                }
            case "a":
                if false {
                    switch 32767 {
                    case 0:
                        switch -68 {
                        case 9223372036854775807:
                            v12 = true
                        }
                    }
                    if false {
                        if v12 {
                            if (0.0 >= 0.1) {
                                fmt.Println(v3, v12)
                                for fi22 := 0; fi22 < 2; fi22++ {
                                    var v23 int16 = (^(v2))
                                    var v24 int32 = 1
                                    if v12 {
                                        if false {
                                            v11 = v11
                                            for fi25 := 0; fi25 < 17; fi25++ {
                                                v11 += (^(v11))
                                                v26 := GenS4{F0: 616}
                                                v11 = 4294967295
                                                _ = fi25
                                                _ = v26
                                            }
                                        }
                                        if true {
                                            var v27 float64 = -2.25
                                            _ = v27
                                        } else {
                                            var v28 uint8 = 2
                                            for fi29 := 0; fi29 < 19; fi29++ {
                                                switch ("a" + "\n\t") {
                                                case "a":
                                                    fmt.Println(fi22)
                                                }
                                                _ = fi29
                                            }
                                            _ = v28
                                        }
                                    }
                                    _ = fi22
                                    _ = v23
                                    _ = v24
                                }
                                v2 = (v2 ^ (-(v2)))
                            }
                            if len(v1) > 0 {
                                v1[0] = -9223372036854775808
                            }
                            var v30 int64 = -223
                            _ = v30
                        }
                        v31 := GenS5{F0: 146, F1: "a"}
                        _ = v31
                    }
                }
            case "\n\t":
                switch "foo bar" {
                case "a":
                    v32 := GenS6{F0: 2147483647}
                    _ = v32
                default:
                    var v33 uint8 = 128
                    _ = v33
                }
            default:
                v4[32768] = 784
            }
            fmt.Println(v12)
        }
    }
    v34 := v4[32768]
    if (!(!false)) {
        if false {
            var v35 uint = 42
            _ = v35
        }
    }
    switch 127 {
    case 42:
        var v36 float64 = 0.0
        _ = v36
    case 9223372036854775807:
        var v37 uint64 = 964
        _ = v37
    case -9223372036854775808:
        for fi38 := 0; fi38 < 15; fi38++ {
            var v39 uint64 = 4294967295
            var v40 int32 = 42
            _ = fi38
            _ = v39
            _ = v40
        }
    default:
        for fi41 := 0; fi41 < 18; fi41++ {
            v2 = 32767
            var v42 uint32 = (v34 / (((v34 << (((v3 ^ v11)) & 63))) | 1))
            _ = fi41
            _ = v42
        }
    }
    for fi43 := 0; fi43 < 5; fi43++ {
        v44 := GenS3{F0: -97}
        _ = fi43
        _ = v44
    }
    var v45 uint8 = 117
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v45
}
