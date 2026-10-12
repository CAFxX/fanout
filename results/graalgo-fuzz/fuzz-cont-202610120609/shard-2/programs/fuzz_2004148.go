package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 int32 }

func main() {
    var v1 int = 32767
    v2 := make(map[int]int)
    var v3 int32 = 42
    var v4 int = 245
    var v5 float64 = ((0.0 + 0.1) - 0.1)
    v6 := make(map[uint32]string)
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 uint16 = 793
        switch fi7 {
        case 127:
            var v9 uint16 = (v8 ^ (v8 / ((v8) | 1)))
            _ = v9
        case 1:
            for fi10 := 0; fi10 < 9; fi10++ {
                if (971 == 598) {
                    switch "hello" {
                    case "a":
                        v4 = fi7
                    case "":
                        var v11 float64 = ((0.0 * v5) + (1e10 - float64(0.0)))
                        _ = v11
                    }
                    v12 := []int{216, (fi7 * (fi7 / ((1) | 1))), (v4 % ((128) | 1))}
                    v13 := GenS1{F0: 1000, F1: 153, F2: fi10}
                    _ = v12
                    _ = v13
                } else {
                    fmt.Println(v5, v5, v5)
                    if false {
                        v5 += v5
                        v14 := v6[1]
                        v3 -= -2147483648
                        _ = v14
                    } else {
                        if false {
                            v1 = v4
                            if (988 < 4294967295) {
                                v15 := GenS1{F0: 785, F1: 2, F2: (v1 & fi7)}
                                _ = v15
                            }
                            if (2 < 2147483647) {
                                v16 := GenS1{F0: 0, F1: uint8(v8), F2: (fi7 << ((255) & 63))}
                                _ = v16
                            } else {
                                switch 2147483647 {
                                case 360:
                                    var v17 uint32 = 264
                                    _ = v17
                                }
                                if (1 != 37) {
                                    var v18 int16 = -480
                                    _ = v18
                                } else {
                                    v8 = 522
                                }
                                fmt.Println(v4, v5)
                            }
                        }
                    }
                    fmt.Println(v1, v3)
                }
                var v19 int64 = 9223372036854775807
                var v20 float32 = 0.1
                _ = fi10
                _ = v19
                _ = v20
            }
        case -239:
            v21 := GenS1{F0: -32768, F1: 82, F2: v1}
            _ = v21
        default:
            switch "a" {
            case "":
                var v22 uint8 = 2
                _ = v22
            }
        }
        var v23 float32 = 3.14159
        _ = fi7
        _ = v23
        _ = v8
    }
    if (!(float32(-2.25) >= 3.14159)) {
        v24 := GenS2{F0: 2}
        v3 += v3
        for fi25 := 0; fi25 < 7; fi25++ {
            v5 += (-2.25 - 0.1)
            switch "a" {
            case "":
                v4 -= 4294967295
            case "hello":
                v26 := GenS2{F0: 42}
                _ = v26
            default:
                v27 := GenS3{F0: 4294967295, F1: -2147483648}
                _ = v27
            }
            v28 := GenS3{F0: 2, F1: 931}
            _ = fi25
            _ = v28
        }
        _ = v24
    } else {
        for fi29 := 0; fi29 < 13; fi29++ {
            v1 = v4
            _ = fi29
        }
    }
    v4 -= v1
    if (v1 <= (v4 << ((529) & 63))) {
        switch v1 {
        case -2147483648:
            var v30 uint64 = 18446744073709551615
            _ = v30
        case -578:
            v31 := GenS3{F0: 280, F1: 11}
            _ = v31
        default:
            switch v1 {
            case 127:
                var v32 uint = 42
                _ = v32
            }
        }
        var v33 float32 = 0.1
        _ = v33
    }
    var v34 uint8 = 117
    fmt.Println(v5, v4, v3)
    var v35 float64 = float64(((1e10 - ((-2.25 * (0.1 - 3.14159)) / -2.25)) / 1.0))
    var v36 uint32 = 177
    v36 -= v36
    v37 := GenS1{F0: -32768, F1: 0, F2: (v4 ^ v4)}
    switch v4 {
    case 1:
        v35 = v5
    default:
        var v38 float32 = 0.1
        _ = v38
    }
    fmt.Println(v4, v36, v36)
    var v39 uint32 = (v36 + v36)
    v40 := GenS2{F0: v39}
    fmt.Println(v39, v5, v36)
    switch v1 {
    case -157:
        v41 := GenS1{F0: -399, F1: 252, F2: v1}
        _ = v41
    default:
        v1 -= v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
