package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 int64; F2 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32; F1 int8 }
func genFn0(p0 bool) uint16 {
    fmt.Println(p0)
    p0 = p0
    var v4 int8 = 0
    switch 42 {
    case -9223372036854775808:
        v5 := GenS1{F0: v4}
        _ = v5
    case 127:
        v6 := GenS2{F0: v4}
        _ = v6
    case 9223372036854775807:
        v7 := GenS3{F0: 1, F1: -161, F2: 881}
        _ = v7
    }
    _ = p0
    _ = v4
    return 864
}

func main() {
    v1 := make(map[int]string)
    var v2 int64 = 0
    var v3 uint = 65535
    v2 = v2
    v2 -= v2
    v8 := v1[0]
    v8 = (((v8 + v8) + "foo bar") + v8)
    v8 = v8
    v2 = (v2 * (v2 - v2))
    if (float64(-2.25) == 0.0) {
        v2 += v2
    }
    if (!(!true)) {
        for fi9 := 0; fi9 < 18; fi9++ {
            for fi10 := 0; fi10 < 4; fi10++ {
                var v11 uint = v3
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
        v3 -= (v3 / ((v3) | 1))
    }
    var v12 uint8 = 1
    for fi13 := 0; fi13 < 7; fi13++ {
        if false {
            if (false && ((574 >= 1) || (!false))) {
                fmt.Println(fi13, fi13)
                for fi14 := 0; fi14 < 4; fi14++ {
                    var v15 string = v8
                    fmt.Println(v2, fi13, v2)
                    if (!false) {
                        var v16 uint16 = 42
                        v15 = (((v15 + v8) + v8) + v15)
                        v17 := GenS2{F0: 100}
                        _ = v16
                        _ = v17
                    } else {
                        var v18 int32 = -2147483648
                        var v19 int64 = int64(fi13)
                        switch v8 {
                        case "a":
                            if ((1 == 1) || (uint64(v12) != 2)) {
                                v12 = (v12 % (((^(v12))) | 1))
                                v20 := GenS2{F0: 98}
                                v21 := GenS4{F0: 42}
                                _ = v20
                                _ = v21
                            } else {
                                v22 := []int{32768, (fi14 / ((fi14) | 1))}
                                v15 = v8
                                if false {
                                    v8 = ((((("" + (v15 + v15)) + "hello") + ((v8 + v15) + v8)) + ((v15 + (("a" + v8) + v8)) + ("foo bar" + "\n\t"))) + v15)
                                    v15 = "\n\t"
                                }
                                v23 := GenS4{F0: 100}
                                _ = v22
                                _ = v23
                            }
                        case "":
                            v8 = "hello"
                        case "hello":
                            var v24 bool = (v8 > "foo bar")
                            _ = v24
                        }
                        _ = v18
                        _ = v19
                    }
                    _ = fi14
                    _ = v15
                }
            }
        } else {
            v12 -= (^(v12))
            if (v3 >= (v3 >> ((42) & 63))) {
                switch 4294967295 {
                case 32767:
                    var v25 string = "foo bar"
                    _ = v25
                case 691:
                    for fi26 := 0; fi26 < 11; fi26++ {
                        v3 -= (v3 / ((v3) | 1))
                        v3 += (v3 << ((v3) & 63))
                        v3 -= 83
                        _ = fi26
                    }
                case -9223372036854775808:
                    v3 -= v3
                }
            }
            var v27 int = (fi13 + fi13)
            _ = v27
        }
        v28 := GenS5{F0: 162, F1: 0}
        var v29 uint = (v3 - v3)
        _ = fi13
        _ = v28
        _ = v29
    }
    v30 := GenS1{F0: 71}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v30
    _ = v8
}
