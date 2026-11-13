package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 string; F1 uint8; F2 int64 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 2
    var v3 float64 = 0.0
    v4 := []uint32{4294967295, 65535}
    var v5 string = "a"
    var v6 float64 = (3.14159 / v3)
    var v7 uint64 = 4294967295
    switch 966 {
    case -1:
        if false {
            switch v5 {
            case "":
                delete(v1, 2)
            case "hello":
                if (!(!false)) {
                    switch -188 {
                    case 2147483647:
                        if (0 != v2) {
                            if (v3 == float64(3.14159)) {
                                v8 := GenS1{F0: -258, F1: 467}
                                fmt.Println(v6, v6, v5)
                                _ = len(v4)
                                _ = v8
                            }
                        }
                    default:
                        if (!(138 < uint8(v2))) {
                            switch (v5 + v5) {
                            case "":
                                v2 -= v2
                            default:
                                var v9 bool = (!(65535 == 2147483647))
                                _ = v9
                            }
                            var v10 bool = ((42 >= 42) || (!false))
                            _ = v10
                        } else {
                            for fi11 := 0; fi11 < 2; fi11++ {
                                switch v5 {
                                case "":
                                    _ = len(v4)
                                case "foo bar":
                                    fmt.Println(v6, v2)
                                }
                                switch "foo bar" {
                                case "\n\t":
                                    for fi12 := 0; fi12 < 6; fi12++ {
                                        switch v5 {
                                        case "a":
                                            var v13 float64 = (v3 - (v3 - v6))
                                            _ = v13
                                        case "\n\t":
                                            v2 = (v2 ^ (^(v2)))
                                        }
                                        _ = fi12
                                    }
                                case "":
                                    if false {
                                        v3 = v6
                                        for fi14 := 0; fi14 < 1; fi14++ {
                                            v15 := GenS2{F0: -2147483648, F1: 2}
                                            v16 := GenS3{F0: "", F1: 42, F2: 324}
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            v17 := GenS3{F0: ("\n\t" + v5), F1: 0, F2: -1}
                                            _ = len(v4)
                                            v5 = v5
                                            switch ("a" + v5) {
                                            case "\n\t":
                                                var v18 int32 = 1
                                                _ = v18
                                            case "hello":
                                                var v19 int8 = 127
                                                _ = v19
                                            default:
                                                var v20 uint16 = 255
                                                _ = v20
                                            }
                                            _ = fi14
                                            _ = v15
                                            _ = v16
                                            _ = v17
                                        }
                                    }
                                case "hello":
                                    v6 -= float64(1e10)
                                default:
                                    for fi21 := 0; fi21 < 4; fi21++ {
                                        v7 += (v7 >> ((v2) & 63))
                                        _ = fi21
                                    }
                                }
                                _ = fi11
                            }
                        }
                    }
                }
            }
            var v22 uint16 = 1
            var v23 float32 = 3.14159
            _ = v22
            _ = v23
        } else {
            var v24 uint = 1
            v25 := GenS2{F0: -2147483648, F1: 128}
            v6 += float64(1e10)
            _ = v24
            _ = v25
        }
    case 261:
        v2 = 42
    }
    v26 := v1[970]
    v2 -= 42
    var v27 int16 = -1
    var v28 int8 = -128
    if (!v26) {
        if v26 {
            for fi29 := 0; fi29 < 19; fi29++ {
                v30 := GenS3{F0: ("\n\t" + (v5 + (((("hello" + ("" + v5)) + v5) + "foo bar") + v5))), F1: 42, F2: 1}
                for fi31 := 0; fi31 < 11; fi31++ {
                    if (v26 && (int16(v2) == -32768)) {
                        fmt.Println(v6)
                    }
                    var v32 uint64 = (^(v7))
                    if v26 {
                        v33 := GenS1{F0: 1, F1: -1}
                        v34 := GenS2{F0: -2147483648, F1: 127}
                        _ = v33
                        _ = v34
                    } else {
                        v32 += 2
                        var v35 float64 = v3
                        _ = v35
                    }
                    var v36 uint = (^(v2))
                    v2 -= (v36 | v2)
                    _ = fi31
                    _ = v32
                    _ = v36
                }
                _ = fi29
                _ = v30
            }
        }
        v26 = v26
    }
    for fi37 := 0; fi37 < 5; fi37++ {
        v27 += (v27 << (((v2 & v2)) & 63))
        switch fi37 {
        case 640:
            v6 = v3
        }
        fmt.Println(v5, v5)
        _ = fi37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
