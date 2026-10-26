package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint }
type GenS3 struct { F0 int8; F1 uint16 }

func main() {
    v1 := []uint32{4294967295}
    v2 := make(map[int]uint32)
    var v3 int64 = 1
    v4 := []uint32{2147483648, 2147483647, 1}
    var v5 int = -889
    var v6 uint64 = 18446744073709551615
    if true {
        var v7 bool = (2 == 644)
        _ = v7
    } else {
        v5 += (v5 ^ int(v6))
        for fi8 := 0; fi8 < 12; fi8++ {
            if true {
                switch "foo bar" {
                case "a":
                    var v9 int32 = 722
                    _ = v9
                case "":
                    v10 := v2[(^(v5))]
                    _ = v10
                default:
                    for fi11 := 0; fi11 < 13; fi11++ {
                        fmt.Println(v3, fi11)
                        _ = fi11
                    }
                }
                for fi12 := 0; fi12 < 1; fi12++ {
                    if false {
                        _ = len(v2)
                        v13 := GenS1{F0: (4294967295 >= (v5 / ((-130) | 1)))}
                        v5 -= int(v6)
                        _ = v13
                    }
                    _ = fi12
                }
            } else {
                fmt.Println(v3, v5, v6)
            }
            v5 = (fi8 << ((65535) & 63))
            _ = fi8
        }
    }
    switch 42 {
    case 1:
        var v14 int64 = -1
        _ = v14
    case 42:
        switch (v5 >> ((18446744073709551615) & 63)) {
        case 128:
            var v15 int32 = -985
            _ = v15
        case 734:
            if (4294967295 >= v6) {
                delete(v2, (v5 + (v5 | -2147483648)))
                if true {
                    var v16 int32 = 2147483647
                    _ = v16
                } else {
                    if true {
                        fmt.Println(v3)
                    } else {
                        if false {
                            var v17 int32 = -1
                            _ = v17
                        }
                        switch (v5 | (v5 - (v5 % ((v5) | 1)))) {
                        case 128:
                            v2[int(v3)] = 2147483648
                        }
                    }
                    v6 -= (^(v6))
                }
                var v18 float64 = 0.1
                _ = v18
            }
        case 0:
            if ((!(false && true)) || false) {
                var v19 float64 = 1.5
                var v20 uint32 = 271
                _ = v19
                _ = v20
            } else {
                for fi21 := 0; fi21 < 2; fi21++ {
                    var v22 float64 = -2.25
                    _ = fi21
                    _ = v22
                }
                v23 := GenS2{F0: 32767, F1: 102, F2: 601}
                _ = v23
            }
        }
    case 4294967295:
        var v24 bool = false
        _ = v24
    default:
        if false {
            v25 := GenS3{F0: -1, F1: 42}
            _ = v25
        }
    }
    var v26 float32 = 3.14159
    if true {
        if (-1 < -2147483648) {
            var v27 int32 = 1
            _ = v27
        } else {
            if true {
                v6 += 4294967295
                v28 := GenS1{F0: (uint16(v6) < 786)}
                v4 = append(v4, 2)
                _ = v28
            }
        }
        var v29 uint = 485
        var v30 float64 = (3.14159 - ((1e10 * 1.5) + 1e10))
        _ = v29
        _ = v30
    } else {
        var v31 bool = ((^(v6)) == v6)
        if v31 {
            v5 -= (v5 + 32767)
        } else {
            for fi32 := 0; fi32 < 2; fi32++ {
                if v31 {
                    switch 9223372036854775807 {
                    case 128:
                        v26 = ((0.0 * float32(1.5)) + (v26 * (((((3.14159 - v26) - (v26 * 3.14159)) + 0.1) + v26) + v26)))
                    case 469:
                        fmt.Println(v26, fi32)
                    case 275:
                        switch "" {
                        case "":
                            var v33 int64 = (v3 / ((v3) | 1))
                            _ = v33
                        default:
                            v34 := GenS2{F0: 42, F1: 52, F2: 2}
                            _ = v34
                        }
                    }
                }
                _ = fi32
            }
            v35 := GenS1{F0: (("hello" + "hello") == "")}
            _ = v35
        }
        v26 -= (v26 * (1.5 / 3.14159))
        _ = v31
    }
    v5 -= 127
    var v36 int = (v5 * 9223372036854775807)
    var v37 int8 = 100
    var v38 int = v36
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
