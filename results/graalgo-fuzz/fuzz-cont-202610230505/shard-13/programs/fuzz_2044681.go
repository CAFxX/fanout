package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 uint = 42
    var v2 int16 = 541
    var v3 uint32 = 1
    v4 := make(map[string]uint32)
    switch (("hello" + "") + "foo bar") {
    case "foo bar":
        v3 += (v3 | (v3 / (((v3 - (v3 % (((v3 >> (((v1 ^ v1)) & 63))) | 1)))) | 1)))
    case "":
        v5 := v4[""]
        _ = v5
    default:
        var v6 uint = (v1 >> ((567) & 63))
        _ = v6
    }
    if false {
        if (208 != 742) {
            for fi7 := 0; fi7 < 7; fi7++ {
                for fi8 := 0; fi8 < 6; fi8++ {
                    var v9 uint64 = 84
                    _ = fi8
                    _ = v9
                }
                _ = fi7
            }
            for fi10 := 0; fi10 < 20; fi10++ {
                switch ("hello" + ("" + ("\n\t" + ("a" + "foo bar")))) {
                case "\n\t":
                    for fi11 := 0; fi11 < 16; fi11++ {
                        switch (fi10 | fi10) {
                        case 1:
                            var v12 float64 = (0.0 * 1e10)
                            _ = v12
                        case 849:
                            var v13 int64 = 42
                            _ = v13
                        case -70:
                            var v14 uint32 = v3
                            _ = v14
                        }
                        _ = fi11
                    }
                default:
                    if (!true) {
                        v3 = v3
                        v15 := GenS1{F0: 0}
                        _ = v15
                    } else {
                        for fi16 := 0; fi16 < 13; fi16++ {
                            v2 = (v2 << (((v1 ^ (v1 + 742))) & 63))
                            _ = fi16
                        }
                        var v17 uint64 = 2
                        if ((^(v3)) == (v3 + (^(v3)))) {
                            v18 := GenS2{F0: (^(v1)), F1: uint8(v3), F2: (v3 + v3)}
                            for fi19 := 0; fi19 < 4; fi19++ {
                                var v20 float32 = (1.5 / 1e10)
                                _ = fi19
                                _ = v20
                            }
                            _ = v18
                        }
                        _ = v17
                    }
                }
                _ = fi10
            }
        } else {
            v21 := GenS2{F0: v1, F1: 0, F2: v3}
            v22 := []int{2147483647, 9223372036854775807}
            if (v2 == (-(v2))) {
                var v23 uint8 = 2
                _ = v23
            } else {
                v2 -= v2
            }
            _ = v21
            _ = v22
        }
    } else {
        v24 := GenS3{F0: 241}
        _ = v24
    }
    v3 += (v3 / (((v3 + uint32(v1))) | 1))
    switch "hello" {
    case "hello":
        fmt.Println(v2, v3, v2)
    case "foo bar":
        if true {
            var v25 string = "foo bar"
            switch 9223372036854775807 {
            case 9223372036854775807:
                if false {
                    var v26 int8 = 127
                    _ = v26
                } else {
                    if (4294967295 != 32768) {
                        v27 := GenS2{F0: v1, F1: 127, F2: v3}
                        delete(v4, "\n\t")
                        _ = v27
                    }
                    v28 := GenS2{F0: v1, F1: 1, F2: 42}
                    _ = v28
                }
            case 4294967295:
                switch 9223372036854775807 {
                case 128:
                    v3 = 815
                }
            }
            _ = v25
        } else {
            var v29 uint8 = 1
            _ = v29
        }
    case "\n\t":
        var v30 int64 = 1
        _ = v30
    }
    var v31 uint8 = 42
    v32 := []int{560, int(v1), -9223372036854775808}
    if len(v32) > 0 {
        v32[0] = 525
    }
    if (1 >= 9223372036854775807) {
        v33 := GenS3{F0: (v31 ^ 65)}
        _ = v33
    } else {
        fmt.Println(v1, v2)
    }
    v1 -= v1
    if true {
        fmt.Println(v2, v2)
        _ = len(v4)
        if false {
            for fi34 := 0; fi34 < 20; fi34++ {
                switch fi34 {
                case 4294967295:
                    var v35 uint16 = 255
                    _ = v35
                }
                _ = fi34
            }
        }
    }
    fmt.Println(v2, v2, v1)
    fmt.Println(v1)
    var v36 uint = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v36
    _ = v4
}
