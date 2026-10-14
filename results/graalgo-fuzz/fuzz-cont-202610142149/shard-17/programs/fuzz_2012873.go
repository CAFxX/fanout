package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint16; F1 string; F2 uint64 }

func main() {
    var v1 uint64 = 381
    v2 := make(map[int]string)
    var v3 uint32 = 900
    var v4 uint16 = 32767
    var v5 float64 = (1e10 + 3.14159)
    if true {
        var v6 int32 = -2147483648
        _ = v6
    } else {
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 int32 = 0
            v9 := GenS1{F0: 97, F1: 18446744073709551615}
            v5 = v5
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    v4 += 65535
    if (((v5 * (v5 * 1e10)) >= 1e10) && (255 == 128)) {
        if false {
            if false {
                var v10 uint = 913
                _ = v10
            }
            for fi11 := 0; fi11 < 12; fi11++ {
                if (!true) {
                    for fi12 := 0; fi12 < 5; fi12++ {
                        var v13 float64 = v5
                        v13 = ((1e10 / ((v13 / v5) + v13)) / v13)
                        if true {
                            v5 += (((float64(0.1) * (v13 / (1e10 / (v5 - (v13 + ((((((v5 / float64(3.14159)) / v13) - v5) * (v13 + v13)) + (float64(1e10) - float64(0.1))) / v13)))))) / (v5 - (v5 - v13))) - v13)
                        }
                        _ = fi12
                        _ = v13
                    }
                    var v14 uint8 = 231
                    v3 += 495
                    _ = v14
                } else {
                    switch (fi11 ^ fi11) {
                    case 9223372036854775807:
                        var v15 uint32 = v3
                        _ = v15
                    case 959:
                        var v16 int8 = -1
                        _ = v16
                    case 909:
                        v4 = 319
                    }
                }
                var v17 uint = 18446744073709551615
                _ = fi11
                _ = v17
            }
        } else {
            fmt.Println(v3, v1, v4)
        }
        var v18 uint32 = (v3 * v3)
        _ = v18
    } else {
        v2[-686] = "\n\t"
        v19 := []int{744, 32768, 0, 1}
        _ = v19
    }
    switch 127 {
    case 723:
        var v20 uint16 = v4
        _ = v20
    default:
        if (!(0.1 < float32(0.0))) {
            for fi21 := 0; fi21 < 9; fi21++ {
                fmt.Println(fi21, v4, v5)
                for fi22 := 0; fi22 < 8; fi22++ {
                    v23 := []int{1}
                    if len(v23) > 0 {
                        v23[0] = int(v1)
                    }
                    _ = fi22
                    _ = v23
                }
                fmt.Println(v4, v3)
                _ = fi21
            }
        }
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        fmt.Println(v1)
        for fi25 := 0; fi25 < 6; fi25++ {
            v26 := GenS1{F0: 127, F1: 2}
            var v27 int8 = 127
            var v28 uint64 = v1
            _ = fi25
            _ = v26
            _ = v27
            _ = v28
        }
        var v29 int16 = 0
        _ = fi24
        _ = v29
    }
    if (127 <= 255) {
        if (!true) {
            for fi30 := 0; fi30 < 5; fi30++ {
                var v31 uint8 = 255
                switch ("foo bar" + "") {
                case "\n\t":
                    var v32 bool = true
                    _ = v32
                case "":
                    v3 -= v3
                default:
                    v4 = v4
                }
                v31 = (^(v31))
                _ = fi30
                _ = v31
            }
        } else {
            var v33 int16 = -32768
            v4 -= (v4 >> ((2) & 63))
            _ = v33
        }
        var v34 int64 = 582
        _ = v34
    } else {
        var v35 int32 = 1
        v3 -= v3
        _ = v35
    }
    if true {
        fmt.Println(v5, v3)
        v1 = 636
        v3 += v3
    } else {
        var v36 uint8 = 128
        _ = v36
    }
    for fi37 := 0; fi37 < 16; fi37++ {
        v38 := GenS2{F0: true, F1: 128}
        for fi39 := 0; fi39 < 12; fi39++ {
            switch (fi39 + (fi39 - 2147483647)) {
            case 2147483647:
                fmt.Println(v1, fi37)
            case -513:
                v4 = v4
                v5 -= 3.14159
                v1 = 18446744073709551615
            default:
                if true {
                    v2[(fi37 + 845)] = "a"
                    if false {
                        var v40 int8 = 127
                        v41 := []int{(fi37 - 127), fi37, fi37}
                        v42 := GenS3{F0: uint16(v3), F1: "hello", F2: 0}
                        _ = v40
                        _ = v41
                        _ = v42
                    }
                    var v43 bool = true
                    _ = len(v2)
                    _ = v43
                } else {
                    v1 = 470
                }
            }
            v1 = 0
            _ = fi39
        }
        _ = fi37
        _ = v38
    }
    fmt.Println(v5, v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
