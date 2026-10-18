package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 uint64; F1 uint64 }

func main() {
    var v1 float32 = (float32((1e10 - 0.1)) - float32(0.1))
    v2 := make(map[uint32]int)
    v3 := []int{418, -1, 4294967295}
    var v4 uint = 229
    var v5 uint16 = 945
    v6 := make(map[int]int)
    switch -2147483648 {
    case 32767:
        _ = len(v6)
    case 32768:
        for fi7 := 0; fi7 < 16; fi7++ {
            v8 := GenS1{F0: 100, F1: 42, F2: 0}
            _ = fi7
            _ = v8
        }
    default:
        switch -1 {
        case 687:
            if false {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v9 float32 = float32(1.5)
                var v10 int8 = -1
                _ = v10
                _ = v9
            } else {
                if true {
                    var v11 string = "a"
                    v1 += v1
                    for fi12 := 0; fi12 < 5; fi12++ {
                        if (false && (fi12 == fi12)) {
                            var v13 bool = true
                            _ = v13
                        }
                        var v14 uint = (v4 & v4)
                        _ = fi12
                        _ = v14
                    }
                    _ = v11
                } else {
                    for fi15 := 0; fi15 < 4; fi15++ {
                        if false {
                            fmt.Println(v4, fi15, v1)
                            if false {
                                var v16 uint = v4
                                switch "" {
                                case "a":
                                    v17 := GenS1{F0: 84, F1: 128, F2: v16}
                                    _ = v17
                                default:
                                    for fi18 := 0; fi18 < 9; fi18++ {
                                        v19 := GenS2{F0: -1, F1: -1}
                                        v5 = v5
                                        _ = fi18
                                        _ = v19
                                    }
                                }
                                _ = v16
                            }
                        } else {
                            v20 := GenS1{F0: 0, F1: 1, F2: 2}
                            switch "foo bar" {
                            case "":
                                fmt.Println(v4)
                            case "a":
                                v21 := GenS2{F0: 1, F1: -450}
                                _ = v21
                            case "foo bar":
                                if true {
                                    for fi22 := 0; fi22 < 18; fi22++ {
                                        v23 := GenS2{F0: 0, F1: -799}
                                        v1 = v1
                                        _ = fi22
                                        _ = v23
                                    }
                                    v5 += v5
                                } else {
                                    switch ((("a" + ("foo bar" + "a")) + "\n\t") + "") {
                                    case "a":
                                        var v24 uint64 = 0
                                        _ = v24
                                    case "foo bar":
                                        v25 := GenS2{F0: 0, F1: -411}
                                        _ = v25
                                    case "hello":
                                        var v26 uint32 = 2147483648
                                        _ = v26
                                    }
                                    var v27 float32 = 0.1
                                    _ = v27
                                }
                            }
                            _ = v20
                        }
                        v5 = (v5 * (v5 | 1))
                        if (0.1 == (float64(v1) / (((1.5 / (3.14159 * 3.14159)) / 3.14159) + float64(3.14159)))) {
                            v6[fi15] = 32768
                            if (1.5 < v1) {
                                switch -1 {
                                case -188:
                                    v28 := GenS1{F0: 33, F1: 17, F2: v4}
                                    _ = v28
                                }
                                _ = len(v6)
                                fmt.Println(v1, v1, v1)
                            }
                        } else {
                            var v29 uint64 = 1
                            v30 := GenS1{F0: -1, F1: 128, F2: v4}
                            fmt.Println(v4, fi15)
                            _ = v29
                            _ = v30
                        }
                        _ = fi15
                    }
                    var v31 uint8 = 0
                    _ = v31
                }
                switch 128 {
                case 9223372036854775807:
                    fmt.Println(v4)
                case -9223372036854775808:
                    v32 := v2[589]
                    _ = v32
                case 128:
                    for fi33 := 0; fi33 < 2; fi33++ {
                        v2[1] = (-(fi33))
                        var v34 int = (fi33 - fi33)
                        _ = fi33
                        _ = v34
                    }
                }
            }
        case 1:
            v1 = ((v1 * v1) / (((((3.14159 + 1.5) + ((1e10 * 1.5) * 3.14159)) / (v1 / 1.5)) * (v1 / 3.14159)) + v1))
        default:
            var v35 int32 = -2147483648
            _ = v35
        }
    }
    v36 := GenS3{F0: uint64(v5), F1: 1}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v6[860] = 32768
    v1 += 3.14159
    var v37 uint64 = uint64(v4)
    v5 = v5
    fmt.Println(v37, v37)
    switch "" {
    case "\n\t":
        for fi38 := 0; fi38 < 14; fi38++ {
            fmt.Println(v37, v5)
            v4 = (v4 | 42)
            _ = fi38
        }
    case "":
        for fi39 := 0; fi39 < 6; fi39++ {
            var v40 int64 = -9223372036854775808
            fmt.Println(v5)
            _ = fi39
            _ = v40
        }
    }
    var v41 string = "a"
    var v42 int64 = -724
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
