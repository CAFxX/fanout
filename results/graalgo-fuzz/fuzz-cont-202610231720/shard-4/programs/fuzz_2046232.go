package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int32; F1 int32; F2 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 int16; F2 int }

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{"foo bar"}
    var v3 bool = false
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]uint32)
    var v6 string = "foo bar"
    switch v6 {
    case "a":
        fmt.Println(v6, v6)
    case "foo bar":
        fmt.Println(v3)
    case "":
        var v7 int8 = -128
        _ = v7
    default:
        v3 = ((float64(1.5) - 1.5) < ((float64((((((1e10 * 3.14159) * float32(1e10)) - (0.1 / 0.1)) * -2.25) / 1e10)) / float64(3.14159)) * float64(-2.25)))
    }
    var v8 string = ""
    v8 = "hello"
    if v3 {
        var v9 float32 = 1.5
        _ = v9
    } else {
        for fi10 := 0; fi10 < 5; fi10++ {
            v11 := GenS1{F0: 189, F1: 0}
            switch "\n\t" {
            case "foo bar":
                for fi12 := 0; fi12 < 17; fi12++ {
                    v6 = ("foo bar" + "")
                    v8 = "a"
                    var v13 int16 = 1
                    _ = fi12
                    _ = v13
                }
            default:
                if (-128 >= 127) {
                    for fi14 := 0; fi14 < 13; fi14++ {
                        v3 = v3
                        if true {
                            var v15 int = fi10
                            v16 := GenS1{F0: 42, F1: 1}
                            v5[7] = 4294967295
                            v15 = fi10
                            if (65535 >= 1) {
                                v6 = v6
                                _ = len(v2)
                                switch ((v6 + v6) + ("hello" + (v8 + v6))) {
                                case "":
                                    v17 := v4["foo bar"]
                                    _ = v17
                                case "foo bar":
                                    v6 = v6
                                case "a":
                                    v15 -= (fi10 | (fi14 + (fi10 - fi14)))
                                    v3 = (54 != 72)
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                default:
                                    v1[128] = 2
                                }
                            }
                            _ = v15
                            _ = v16
                        } else {
                            var v18 uint64 = 85
                            _ = v18
                        }
                        v19 := GenS1{F0: 128, F1: -272}
                        _ = fi14
                        _ = v19
                    }
                    v8 = ((("a" + v8) + "foo bar") + "\n\t")
                } else {
                    var v20 int64 = 836
                    fmt.Println(v8, v3, v3)
                    for fi21 := 0; fi21 < 20; fi21++ {
                        var v22 uint64 = 42
                        v22 = (^(v22))
                        _ = fi21
                        _ = v22
                    }
                    _ = v20
                }
            }
            for fi23 := 0; fi23 < 7; fi23++ {
                var v24 int32 = 280
                if (!false) {
                    switch (v6 + "a") {
                    case "hello":
                        switch fi10 {
                        case 219:
                            if v3 {
                                v25 := GenS2{F0: (v24 + -1), F1: 1, F2: 65535}
                                v26 := GenS3{F0: 9223372036854775807}
                                _ = v25
                                _ = v26
                            }
                        case 32768:
                            if v3 {
                                if ((127 <= -70) && (482 < 469)) {
                                    v27 := GenS4{F0: 365, F1: -740, F2: (fi23 / (((fi10 % ((fi10) | 1))) | 1))}
                                    _ = v27
                                } else {
                                    fmt.Println(v6, v6, v8)
                                    var v28 int32 = (v24 >> ((255) & 63))
                                    _ = v28
                                }
                                var v29 int8 = 0
                                _ = v29
                            }
                        default:
                            v30 := GenS2{F0: (v24 | (v24 * v24)), F1: v24, F2: 299}
                            _ = v30
                        }
                    case "foo bar":
                        v31 := GenS3{F0: 9223372036854775807}
                        _ = v31
                    case "a":
                        var v32 uint16 = 1
                        _ = v32
                    }
                    if (255 < 242) {
                        v33 := GenS2{F0: 1, F1: (v24 - v24), F2: 739}
                        v24 = (v24 ^ v24)
                        v34 := GenS3{F0: -425}
                        _ = v33
                        _ = v34
                    } else {
                        var v35 uint64 = 18446744073709551615
                        _ = v35
                    }
                    switch (v6 + v8) {
                    case "\n\t":
                        v36 := GenS3{F0: -9223372036854775808}
                        _ = v36
                    default:
                        var v37 int8 = 44
                        _ = v37
                    }
                }
                v3 = true
                v24 = 0
                _ = fi23
                _ = v24
            }
            _ = fi10
            _ = v11
        }
    }
    v5[972] = 263
    v6 = "a"
    var v38 string = "a"
    if (!false) {
        v38 = ("" + v38)
        var v39 float32 = 1.5
        if v3 {
            for fi40 := 0; fi40 < 7; fi40++ {
                v8 = (((v38 + "") + v38) + v38)
                _ = fi40
            }
        } else {
            var v41 uint8 = 128
            _ = v41
        }
        _ = v39
    }
    if false {
        var v42 float64 = 1.5
        v43 := GenS1{F0: 217, F1: 399}
        if len(v2) > 0 {
            v2[0] = v8
        }
        _ = v42
        _ = v43
    } else {
        fmt.Println(v6)
    }
    v44 := v4[(("" + (((("hello" + v6) + "\n\t") + v38) + ((v8 + "a") + "hello"))) + "a")]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v44
    _ = v5
    _ = v6
    _ = v8
}
