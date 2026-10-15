package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int }

func main() {
    var v1 uint64 = 1
    var v2 int64 = -140
    var v3 int64 = (v2 - 0)
    v4 := make(map[int]int)
    if (float32((0.0 + ((-2.25 / 0.1) / ((-2.25 / float64(0.1)) / 1e10)))) >= -2.25) {
        switch 411 {
        case -53:
            var v5 string = ("foo bar" + "a")
            _ = v5
        case 0:
            switch 42 {
            case -359:
                v6 := GenS1{F0: 255, F1: -940}
                _ = v6
            default:
                if (float64(3.14159) != 0.0) {
                    v1 += 135
                } else {
                    if (1 != -1) {
                        var v7 int = 4294967295
                        _ = v7
                    }
                    if true {
                        var v8 int = 2147483647
                        for fi9 := 0; fi9 < 6; fi9++ {
                            v10 := GenS1{F0: 0, F1: (fi9 << ((612) & 63))}
                            _ = fi9
                            _ = v10
                        }
                        _ = v8
                    }
                }
            }
        default:
            var v11 int8 = -128
            _ = v11
        }
        v2 -= -174
        switch 9223372036854775807 {
        case -1:
            v12 := []int{272, 1, 32768}
            _ = v12
        case 1:
            if (!(3.14159 <= float32(0.0))) {
                switch ("hello" + ("a" + "\n\t")) {
                case "a":
                    if ((-298 < 1) || (uint16(v1) == 65535)) {
                        v4[1] = 32768
                        v2 += (v3 % ((42) | 1))
                    }
                }
                _ = len(v4)
            }
        }
    }
    v1 += v1
    var v13 int = 534
    v14 := GenS1{F0: 255, F1: v13}
    var v15 uint32 = 1
    fmt.Println(v3, v3)
    v2 += int64(v15)
    if true {
        for fi16 := 0; fi16 < 6; fi16++ {
            v17 := GenS2{F0: 42}
            switch (v13 - 2147483647) {
            case 405:
                switch -2147483648 {
                case -335:
                    if ("" <= "\n\t") {
                        var v18 uint32 = v15
                        v4[(^(v13))] = fi16
                        for fi19 := 0; fi19 < 8; fi19++ {
                            v13 += 9223372036854775807
                            v1 = (v1 - v1)
                            switch (v13 % ((fi19) | 1)) {
                            case 50:
                                v20 := []int{(fi19 / ((4294967295) | 1)), fi16, fi19, fi16}
                                _ = v20
                            }
                            _ = fi19
                        }
                        _ = v18
                    } else {
                        v21 := v4[fi16]
                        if true {
                            v1 -= (v1 % ((0) | 1))
                        } else {
                            for fi22 := 0; fi22 < 6; fi22++ {
                                if false {
                                    v21 -= fi22
                                    var v23 uint8 = 171
                                    _ = v23
                                }
                                v24 := []int{(v21 >> ((4294967295) & 63))}
                                if len(v24) > 0 {
                                    v24[0] = (v13 & (fi16 - (^(v21))))
                                }
                                _ = fi22
                                _ = v24
                            }
                            v3 = -485
                            fmt.Println(v3, fi16)
                        }
                        for fi25 := 0; fi25 < 20; fi25++ {
                            var v26 int8 = 0
                            _ = fi25
                            _ = v26
                        }
                        _ = v21
                    }
                default:
                    v27 := v4[fi16]
                    _ = v27
                }
            }
            var v28 uint64 = 962
            _ = fi16
            _ = v17
            _ = v28
        }
        v13 -= v13
    }
    fmt.Println(v13, v3)
    for fi29 := 0; fi29 < 4; fi29++ {
        switch "foo bar" {
        case "hello":
            if (42 < 1000) {
                v30 := GenS2{F0: 0}
                _ = v30
            }
        default:
            var v31 uint16 = 255
            _ = v31
        }
        v32 := GenS3{F0: 18446744073709551615}
        v33 := GenS4{F0: (fi29 + (v13 | (v13 + fi29)))}
        _ = fi29
        _ = v32
        _ = v33
    }
    v34 := GenS1{F0: 18446744073709551615, F1: -1}
    v3 -= (v3 + 9223372036854775807)
    v35 := GenS3{F0: 42}
    var v36 uint16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
