package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int64 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 int; F1 uint16 }
func genFn0(p0 int64, p1 bool, p2 int32) string {
    v7 := GenS1{F0: p0, F1: 1, F2: p0}
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return "\n\t"
}

func main() {
    var v1 bool = true
    v2 := []string{((("hello" + "\n\t") + ("" + "a")) + "\n\t")}
    var v3 int64 = 0
    v4 := []int{-466, 32768}
    v5 := []int{}
    var v6 float64 = 3.14159
    var v8 int32 = -948
    for fi9 := 0; fi9 < 20; fi9++ {
        v10 := GenS1{F0: (v3 | v3), F1: 4294967295, F2: (v3 | (^(v3)))}
        if (false && (true || (42 == 32767))) {
            fmt.Println(v8, v6)
            v11 := GenS1{F0: 912, F1: 4294967295, F2: v3}
            _ = v11
        } else {
            var v12 uint32 = 2
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    var v13 bool = (-1 != -1)
    for fi14 := 0; fi14 < 5; fi14++ {
        for fi15 := 0; fi15 < 17; fi15++ {
            v3 = (v3 / ((-450) | 1))
            v8 = v8
            var v16 uint8 = 168
            _ = fi15
            _ = v16
        }
        v17 := GenS2{F0: -1, F1: "a"}
        _ = fi14
        _ = v17
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        var v19 int8 = 0
        _ = fi18
        _ = v19
    }
    switch 1 {
    case 1:
        v5 = append(v5, 854)
    case 42:
        if (9223372036854775807 >= v3) {
            v20 := GenS1{F0: (v3 % ((int64(v8)) | 1)), F1: 2, F2: 1}
            switch 128 {
            case 664:
                fmt.Println(v13)
            case 987:
                v8 = v8
            default:
                for fi21 := 0; fi21 < 12; fi21++ {
                    for fi22 := 0; fi22 < 19; fi22++ {
                        if ((347 != (fi21 / ((fi21) | 1))) || false) {
                            v8 -= v8
                            v23 := GenS1{F0: (^(v3)), F1: 807, F2: (^(v3))}
                            v6 -= v6
                            _ = v23
                        } else {
                            fmt.Println(v13)
                            v24 := GenS1{F0: 286, F1: 246, F2: (v3 * 0)}
                            _ = v24
                        }
                        var v25 int32 = v8
                        v5 = append(v5, fi21)
                        _ = fi22
                        _ = v25
                    }
                    _ = fi21
                }
            }
            _ = v20
        } else {
            v4 = append(v4, 2147483647)
            switch (("\n\t" + ("a" + "hello")) + "\n\t") {
            case "":
                var v26 int8 = int8(v8)
                _ = v26
            default:
                switch 4294967295 {
                case 1:
                    var v27 uint16 = 668
                    _ = v27
                case 32768:
                    v3 -= v3
                case 283:
                    switch 0 {
                    case 2147483647:
                        v28 := GenS2{F0: 2147483647, F1: "\n\t"}
                        _ = v28
                    default:
                        if v1 {
                            for fi29 := 0; fi29 < 2; fi29++ {
                                if v1 {
                                    v30 := GenS1{F0: (v3 / (((^(v3))) | 1)), F1: 526, F2: v3}
                                    _ = v30
                                } else {
                                    v5 = append(v5, 498)
                                    v31 := GenS3{F0: 127, F1: 0}
                                    _ = v31
                                }
                                _ = fi29
                            }
                        }
                    }
                }
            }
            fmt.Println(v3)
        }
    case -802:
        if len(v4) > 0 {
            v4[0] = 2147483647
        }
    }
    var v32 int64 = v3
    for fi33 := 0; fi33 < 11; fi33++ {
        _ = len(v5)
        _ = fi33
    }
    v13 = ((-1 <= 42) || v1)
    v8 += v8
    var v34 bool = v13
    var v35 uint64 = 1
    for fi36 := 0; fi36 < 6; fi36++ {
        v1 = (496 > 2)
        v2 = append(v2, "a")
        _ = fi36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
