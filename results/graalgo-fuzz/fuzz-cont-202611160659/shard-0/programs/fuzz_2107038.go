package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{}
    var v3 uint = 65535
    v4 := make(map[int]bool)
    v5 := make(map[uint32]int)
    v6 := make(map[int]int)
    var v7 int16 = 32767
    var v8 bool = (true || (((!(1 == 127)) || (!false)) && (0 != 127)))
    for fi9 := 0; fi9 < 9; fi9++ {
        switch "foo bar" {
        case "a":
            var v10 uint32 = 181
            _ = v10
        case "\n\t":
            var v11 uint64 = 696
            _ = v11
        case "":
            v12 := GenS1{F0: 2, F1: 212}
            _ = v12
        }
        _ = fi9
    }
    var v13 int32 = 0
    var v14 int32 = (v13 ^ v13)
    if (65 >= -1) {
        v15 := v6[214]
        switch (v15 % (((v15 ^ (v15 << (((v3 >> ((255) & 63))) & 63)))) | 1)) {
        case 42:
            for fi16 := 0; fi16 < 19; fi16++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v3 += (v3 % ((v3) | 1))
                v8 = v8
                _ = fi16
            }
        default:
            if false {
                for fi17 := 0; fi17 < 15; fi17++ {
                    if len(v2) > 0 {
                        v2[0] = "foo bar"
                    }
                    var v18 int8 = 127
                    _ = fi17
                    _ = v18
                }
                var v19 int = v15
                _ = v19
            } else {
                for fi20 := 0; fi20 < 7; fi20++ {
                    var v21 uint64 = 91
                    switch "hello" {
                    case "":
                        for fi22 := 0; fi22 < 9; fi22++ {
                            var v23 int64 = -9223372036854775808
                            v21 = v21
                            _ = fi22
                            _ = v23
                        }
                    case "a":
                        if true {
                            v24 := GenS1{F0: 61, F1: 2}
                            _ = v24
                        }
                    default:
                        for fi25 := 0; fi25 < 12; fi25++ {
                            v21 -= 0
                            _ = fi25
                        }
                    }
                    _ = fi20
                    _ = v21
                }
                v8 = (563 != (v3 | v3))
                for fi26 := 0; fi26 < 10; fi26++ {
                    for fi27 := 0; fi27 < 9; fi27++ {
                        switch 128 {
                        case 128:
                            fmt.Println(v7, v15, v8)
                        default:
                            switch 2147483647 {
                            case 2147483647:
                                var v28 uint64 = 18446744073709551615
                                _ = v28
                            default:
                                if (v8 || v8) {
                                    var v29 float64 = ((float64(1e10) - ((float64(1e10) * (0.0 * 0.0)) / 0.1)) - 3.14159)
                                    v8 = false
                                    _ = v29
                                } else {
                                    switch (fi27 | int(v7)) {
                                    case -841:
                                        var v30 uint32 = 4294967295
                                        _ = v30
                                    case 42:
                                        v15 -= (v15 | -2147483648)
                                    case 9223372036854775807:
                                        v8 = ((0.0 * float64(0.0)) > (1.5 - (3.14159 + ((0.1 - 1.5) + (0.0 / 1.0)))))
                                    default:
                                        v3 -= v3
                                    }
                                    v31 := GenS1{F0: 146, F1: 255}
                                    _ = v31
                                }
                            }
                        }
                        if ((0.1 * 1.5) < (0.1 * (0.0 - 1.5))) {
                            v32 := GenS1{F0: 255, F1: 42}
                            _ = v32
                        }
                        delete(v1, 340)
                        var v33 int16 = -1
                        _ = fi27
                        _ = v33
                    }
                    v15 -= int(v3)
                    _ = len(v1)
                    _ = fi26
                }
            }
        }
        var v34 int16 = 1000
        _ = v15
        _ = v34
    }
    for fi35 := 0; fi35 < 15; fi35++ {
        var v36 float64 = ((1e10 - -2.25) / (1.5 / -2.25))
        var v37 uint32 = 1
        _ = fi35
        _ = v36
        _ = v37
    }
    v38 := GenS1{F0: 226, F1: 74}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
