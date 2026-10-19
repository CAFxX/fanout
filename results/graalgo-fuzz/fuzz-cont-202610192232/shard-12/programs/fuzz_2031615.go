package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }

func main() {
    var v1 int16 = 689
    var v2 bool = false
    v3 := []uint32{1}
    v1 -= (v1 << ((18446744073709551615) & 63))
    v4 := GenS1{F0: -782, F1: -1}
    if len(v3) > 0 {
        v3[0] = 919
    }
    if (-1 != 2147483647) {
        var v5 uint = 726
        switch -398 {
        case -416:
            v6 := GenS1{F0: 1, F1: 47}
            _ = v6
        case -2147483648:
            v2 = (244 < 0)
        default:
            if v2 {
                if len(v3) > 0 {
                    v3[0] = 2
                }
            } else {
                if false {
                    var v7 bool = (0.0 < -2.25)
                    _ = v7
                }
            }
        }
        _ = v5
    } else {
        if false {
            var v8 uint8 = 199
            var v9 uint32 = 1
            _ = v8
            _ = v9
        }
    }
    v2 = v2
    v10 := GenS1{F0: v1, F1: 42}
    for fi11 := 0; fi11 < 13; fi11++ {
        fmt.Println(v1, v1, v1)
        _ = fi11
    }
    switch 1 {
    case 2147483647:
        if (!(42 >= 0)) {
            v12 := GenS1{F0: (v1 >> ((879) & 63)), F1: int64(v1)}
            _ = v12
        } else {
            var v13 int32 = 1
            v2 = v2
            switch "a" {
            case "":
                var v14 int64 = -863
                _ = v14
            case "\n\t":
                var v15 string = ""
                _ = v15
            }
            _ = v13
        }
    default:
        if false {
            v1 = 0
        } else {
            for fi16 := 0; fi16 < 1; fi16++ {
                var v17 int8 = 127
                v2 = true
                if v2 {
                    _ = len(v3)
                    if (!(33 > 0)) {
                        switch fi16 {
                        case 32768:
                            var v18 uint16 = 2
                            _ = v18
                        case 42:
                            if v2 {
                                for fi19 := 0; fi19 < 4; fi19++ {
                                    v17 = (-(v17))
                                    _ = fi19
                                }
                            } else {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v20 := GenS1{F0: (v1 % ((0) | 1)), F1: 1}
                                _ = v20
                            }
                        case -931:
                            switch "" {
                            case "":
                                var v21 uint = 18446744073709551615
                                _ = v21
                            case "\n\t":
                                for fi22 := 0; fi22 < 12; fi22++ {
                                    fmt.Println(v2)
                                    fmt.Println(fi22, v17)
                                    _ = fi22
                                }
                            }
                        default:
                            for fi23 := 0; fi23 < 19; fi23++ {
                                v3 = append(v3, 65535)
                                v17 += -119
                                fmt.Println(v2, v17, fi23)
                                _ = fi23
                            }
                        }
                        v1 = 1
                    }
                    for fi24 := 0; fi24 < 16; fi24++ {
                        if (v2 || ((^(v17)) < v17)) {
                            var v25 int64 = -1
                            v25 = v25
                            if (v25 < (^(v25))) {
                                var v26 uint64 = 800
                                fmt.Println(fi24)
                                _ = v26
                            }
                            _ = v25
                        }
                        var v27 int16 = 42
                        _ = fi24
                        _ = v27
                    }
                }
                _ = fi16
                _ = v17
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v28 int32 = 1
            _ = v28
        }
    }
    if (-110 >= 117) {
        v1 = (-(v1))
    } else {
        if len(v3) > 0 {
            v3[0] = 2147483647
        }
        v29 := GenS1{F0: (^(v1)), F1: 407}
        v30 := GenS1{F0: (v1 >> ((255) & 63)), F1: -9223372036854775808}
        _ = v29
        _ = v30
    }
    var v31 float64 = float64(-2.25)
    v32 := GenS1{F0: v1, F1: 1}
    v2 = (9223372036854775807 <= -495)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
}
