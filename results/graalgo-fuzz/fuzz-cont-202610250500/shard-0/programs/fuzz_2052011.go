package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8; F1 uint32 }

func main() {
    var v1 float32 = 0.0
    v2 := make(map[int]bool)
    var v3 string = ("\n\t" + "\n\t")
    v4 := []int{683}
    v1 = v1
    if (952 == 582) {
        _ = len(v2)
    } else {
        v1 = (v1 * 3.14159)
        var v5 int = 9223372036854775807
        _ = v5
    }
    _ = len(v2)
    if (42 == 128) {
        var v6 uint16 = 42
        for fi7 := 0; fi7 < 2; fi7++ {
            v6 -= v6
            switch ((v3 + v3) + "hello") {
            case "foo bar":
                for fi8 := 0; fi8 < 16; fi8++ {
                    v9 := GenS1{F0: 2}
                    v6 -= v6
                    var v10 float64 = 0.1
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            case "":
                fmt.Println(v6)
            }
            fmt.Println(v6, v1)
            _ = fi7
        }
        _ = v6
    }
    delete(v2, -2147483648)
    switch 127 {
    case 42:
        v3 = v3
    case 9223372036854775807:
        v11 := GenS1{F0: 2}
        _ = v11
    case 1:
        var v12 int16 = 1000
        _ = v12
    default:
        _ = len(v4)
    }
    if true {
        if (true && false) {
            for fi13 := 0; fi13 < 11; fi13++ {
                delete(v2, (-(fi13)))
                v14 := GenS1{F0: 2147483647}
                _ = fi13
                _ = v14
            }
            switch 9223372036854775807 {
            case -911:
                v15 := GenS1{F0: 663}
                _ = v15
            }
            switch 2147483647 {
            case 1:
                v3 = "foo bar"
            case -9223372036854775808:
                if len(v4) > 0 {
                    v4[0] = 4294967295
                }
            default:
                var v16 float64 = 0.1
                _ = v16
            }
        }
        var v17 int16 = 1000
        _ = len(v4)
        _ = v17
    } else {
        fmt.Println(v1)
        v3 = (v3 + "\n\t")
    }
    v1 = v1
    v18 := GenS1{F0: 218}
    for fi19 := 0; fi19 < 7; fi19++ {
        if false {
            for fi20 := 0; fi20 < 17; fi20++ {
                if (156 <= 42) {
                    fmt.Println(fi19, fi19, v3)
                    v1 = v1
                    if (!(false && (false && ((false || (0 < 115)) && ((v1 - (v1 * (v1 - -2.25))) > ((float32(float64(v1)) - (v1 * 1.5)) + float32((1.5 / 1.0)))))))) {
                        if (!(1 <= 458)) {
                            v21 := GenS2{F0: -128}
                            var v22 int32 = 1
                            var v23 uint16 = 65535
                            _ = v21
                            _ = v22
                            _ = v23
                        } else {
                            var v24 uint64 = 4294967295
                            if true {
                                switch fi19 {
                                case 319:
                                    v25 := GenS2{F0: 100}
                                    _ = v25
                                case -2147483648:
                                    v24 = v24
                                default:
                                    _ = len(v4)
                                }
                            } else {
                                fmt.Println(fi19, v1, v3)
                                switch "foo bar" {
                                case "hello":
                                    if (552 != 2) {
                                        v24 = v24
                                        v3 = v3
                                        switch fi19 {
                                        case 42:
                                            var v26 int64 = 42
                                            _ = v26
                                        }
                                    } else {
                                        var v27 float32 = ((v1 + v1) * ((0.1 + 1e10) - (v1 + 0.1)))
                                        _ = v27
                                    }
                                }
                                var v28 float64 = -2.25
                                v29 := GenS3{F0: 2147483647}
                                _ = v28
                                _ = v29
                            }
                            _ = v24
                        }
                        var v30 float64 = 0.1
                        fmt.Println(fi19, fi19)
                        _ = v30
                    }
                    var v31 float64 = 1.5
                    delete(v2, fi20)
                    _ = v31
                }
                v32 := GenS4{F0: 1, F1: 65535}
                _ = fi20
                _ = v32
            }
        }
        _ = fi19
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
