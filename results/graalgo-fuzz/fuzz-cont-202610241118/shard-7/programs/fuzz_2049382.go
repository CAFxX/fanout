package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 int16 = 0
    var v2 int32 = -85
    v3 := make(map[uint32]uint32)
    var v4 bool = false
    var v5 int64 = 9223372036854775807
    v5 = int64(v1)
    switch 42 {
    case 324:
        if (v4 || v4) {
            for fi6 := 0; fi6 < 14; fi6++ {
                if (65535 != 755) {
                    v7 := []int{}
                    v2 = v2
                    for fi8 := 0; fi8 < 7; fi8++ {
                        fmt.Println(fi8)
                        v2 += (v2 << ((1) & 63))
                        _ = fi8
                    }
                    _ = v7
                } else {
                    var v9 float64 = float64(0.1)
                    _ = v9
                }
                _ = fi6
            }
            var v10 int64 = -172
            _ = v10
        }
    case -725:
        if (!(0.1 < 0.1)) {
            var v11 int16 = (-(v1))
            _ = v11
        } else {
            if (v4 && ((v1 - v1) <= (^(v1)))) {
                var v12 int = 9223372036854775807
                _ = v12
            }
            if v4 {
                for fi13 := 0; fi13 < 19; fi13++ {
                    v5 += v5
                    var v14 int32 = (v2 >> ((630) & 63))
                    _ = fi13
                    _ = v14
                }
                var v15 uint = 18446744073709551615
                _ = v15
            }
        }
    }
    fmt.Println(v2, v5, v5)
    v2 = v2
    v1 = (^(v1))
    var v16 uint = 276
    v17 := GenS1{F0: ("\n\t" + "\n\t"), F1: 0}
    var v18 float64 = 3.14159
    for fi19 := 0; fi19 < 2; fi19++ {
        v4 = (2147483647 != 0)
        if (v1 <= v1) {
            v2 += v2
            switch ("\n\t" + ("a" + "foo bar")) {
            case "":
                var v20 int32 = (v2 & v2)
                _ = v20
            case "a":
                if (v4 || v4) {
                    fmt.Println(v16, v16)
                    v16 += v16
                    var v21 uint32 = 42
                    _ = v21
                }
            case "foo bar":
                v22 := GenS1{F0: (("a" + "foo bar") + "hello"), F1: 1}
                _ = v22
            default:
                v5 += 42
            }
            for fi23 := 0; fi23 < 8; fi23++ {
                var v24 string = "\n\t"
                var v25 uint16 = 65535
                _ = fi23
                _ = v24
                _ = v25
            }
        } else {
            v26 := v3[392]
            _ = v26
        }
        _ = fi19
    }
    if false {
        v1 -= (v1 >> ((731) & 63))
        if v4 {
            if v4 {
                v27 := []int{9223372036854775807, 127}
                if v4 {
                    switch -24 {
                    case -845:
                        if ((v2 + (v2 ^ -2147483648)) == v2) {
                            v2 -= (v2 & v2)
                            var v28 uint64 = 423
                            _ = v28
                        }
                    case 443:
                        switch 42 {
                        case 0:
                            var v29 int8 = 0
                            _ = v29
                        case 542:
                            var v30 int64 = (v5 / ((v5) | 1))
                            _ = v30
                        }
                    case -2147483648:
                        for fi31 := 0; fi31 < 1; fi31++ {
                            v3[2147483647] = 459
                            v1 -= (v1 >> ((v16) & 63))
                            var v32 int16 = v1
                            _ = fi31
                            _ = v32
                        }
                    }
                    var v33 int64 = v5
                    _ = v33
                } else {
                    v34 := GenS1{F0: "hello", F1: 1}
                    _ = v34
                }
                _ = v27
            } else {
                var v35 uint32 = 873
                _ = v35
            }
            v4 = ((v5 | (v5 & v5)) != v5)
        } else {
            v36 := GenS2{F0: 0}
            _ = len(v3)
            v3[65535] = 2147483647
            _ = v36
        }
    } else {
        for fi37 := 0; fi37 < 20; fi37++ {
            if false {
                v38 := GenS1{F0: "foo bar", F1: -116}
                _ = v38
            }
            fmt.Println(fi37, v2, v5)
            v39 := v3[42]
            _ = fi37
            _ = v39
        }
        v2 = 42
    }
    v4 = v4
    v40 := []int{-858}
    v40 = append(v40, -749)
    v5 -= v5
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
}
