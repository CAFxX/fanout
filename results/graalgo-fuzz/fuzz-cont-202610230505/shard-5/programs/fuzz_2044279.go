package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int16 }
type GenS2 struct { F0 uint8; F1 int64 }
func genFn0(p0 uint32, p1 int8) int64 {
    switch "" {
    case "":
        v6 := GenS1{F0: 9223372036854775807, F1: (p1 >> ((42) & 63)), F2: -1}
        _ = v6
    case "foo bar":
        var v7 uint32 = p0
        _ = v7
    default:
        for fi8 := 0; fi8 < 19; fi8++ {
            v9 := []int{-9223372036854775808, (fi8 << ((1) & 63))}
            _ = len(v9)
            p0 += p0
            _ = fi8
            _ = v9
        }
    }
    v10 := GenS1{F0: -977, F1: (p1 + 1), F2: 719}
    v11 := []int{-9223372036854775808, -2147483648}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 345
}

func main() {
    var v1 int32 = 0
    v2 := make(map[string]uint32)
    v3 := []uint32{2}
    var v4 uint32 = 1
    var v5 int64 = 461
    var v12 int = 128
    v5 = (v5 - 9223372036854775807)
    if ((42 < 0) || (4294967295 == 126)) {
        if false {
            v13 := GenS1{F0: (v5 & genFn0(65535, 42)), F1: 0, F2: int16(v1)}
            var v14 bool = true
            delete(v2, "")
            _ = v13
            _ = v14
        } else {
            v1 = (^(v1))
            for fi15 := 0; fi15 < 15; fi15++ {
                switch 32768 {
                case 80:
                    v1 = (v1 * (v1 + -884))
                case 795:
                    switch (v12 << ((2) & 63)) {
                    case -1:
                        for fi16 := 0; fi16 < 5; fi16++ {
                            v12 += (fi16 >> ((1) & 63))
                            if false {
                                switch "a" {
                                case "foo bar":
                                    fmt.Println(v12)
                                }
                                v1 = v1
                            } else {
                                v17 := GenS1{F0: -9223372036854775808, F1: 68, F2: 32767}
                                _ = v17
                            }
                            v1 -= (^(v1))
                            _ = fi16
                        }
                    }
                case 9223372036854775807:
                    v18 := GenS1{F0: (v5 - (v5 | v5)), F1: 91, F2: 42}
                    _ = v18
                default:
                    for fi19 := 0; fi19 < 10; fi19++ {
                        if false {
                            var v20 uint = 876
                            _ = v20
                        }
                        var v21 string = "foo bar"
                        _ = fi19
                        _ = v21
                    }
                }
                v5 -= 9223372036854775807
                _ = fi15
            }
        }
        for fi22 := 0; fi22 < 10; fi22++ {
            v23 := GenS2{F0: 1, F1: v5}
            var v24 uint = 4294967295
            fmt.Println(v12, v12)
            _ = fi22
            _ = v23
            _ = v24
        }
        var v25 string = "a"
        _ = v25
    } else {
        for fi26 := 0; fi26 < 4; fi26++ {
            if false {
                v1 += (-(v1))
                var v27 bool = (v12 != (v12 | fi26))
                _ = v27
            }
            _ = fi26
        }
    }
    var v28 bool = ((232 > 127) && false)
    switch v12 {
    case 32767:
        for fi29 := 0; fi29 < 6; fi29++ {
            for fi30 := 0; fi30 < 17; fi30++ {
                v28 = v28
                if (42 != 1) {
                    fmt.Println(v4)
                }
                v12 = (v12 | 128)
                _ = fi30
            }
            v31 := GenS2{F0: 127, F1: v5}
            v4 -= v4
            _ = fi29
            _ = v31
        }
    case -9223372036854775808:
        v28 = ((("" + "foo bar") + "foo bar") != "foo bar")
    }
    v32 := v2["a"]
    var v33 uint16 = 886
    switch v12 {
    case 891:
        v34 := GenS1{F0: (v5 >> ((18446744073709551615) & 63)), F1: -33, F2: 688}
        _ = v34
    case 2147483647:
        var v35 uint64 = 42
        _ = v35
    case 0:
        v12 -= -9223372036854775808
    }
    switch (-(v12)) {
    case 523:
        var v36 int16 = -424
        _ = v36
    case -429:
        for fi37 := 0; fi37 < 11; fi37++ {
            if v28 {
                _ = len(v3)
                for fi38 := 0; fi38 < 14; fi38++ {
                    v39 := v2["hello"]
                    _ = fi38
                    _ = v39
                }
            } else {
                for fi40 := 0; fi40 < 18; fi40++ {
                    v33 -= (v33 >> ((2) & 63))
                    _ = fi40
                }
            }
            _ = fi37
        }
    case 2147483647:
        v1 = int32(v4)
    }
    var v41 int64 = 328
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v41
    _ = v5
}
