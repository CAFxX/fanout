package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 int16; F2 int16 }
func genFn0(p0 uint32, p1 int) string {
    switch -315 {
    case 9223372036854775807:
        if ((p1 - p1) != p1) {
            v7 := []int{}
            _ = v7
        }
    default:
        v8 := []int{(p1 - (p1 ^ (p1 * p1))), p1, (p1 & p1), 32768}
        _ = v8
    }
    p0 += (p0 ^ (^(p0)))
    v9 := []int{p1, (p1 + p1), -2147483648, 42}
    if (!true) {
        switch 0 {
        case 32767:
            var v10 uint8 = 42
            _ = v10
        case -2147483648:
            if (false || (9223372036854775807 != 1)) {
                for fi11 := 0; fi11 < 4; fi11++ {
                    p0 -= p0
                    _ = fi11
                }
            }
        default:
            p0 = 2147483648
        }
        if true {
            p1 -= 4294967295
        }
        for fi12 := 0; fi12 < 1; fi12++ {
            p0 = (p0 % ((p0) | 1))
            _ = fi12
        }
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v9
    return "foo bar"
}
func genFn1(p0 uint64) uint64 {
    v13 := []int{165, -783}
    var v14 int16 = -1
    v15 := GenS1{F0: 128}
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    _ = p0
    _ = v13
    _ = v14
    _ = v15
    return (p0 << ((4294967295) & 63))
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = (0.1 * float64(1.5))
    var v3 uint16 = 2
    v4 := []string{}
    var v5 uint64 = 968
    var v6 uint8 = 128
    if (false || (1 <= 0)) {
        for fi16 := 0; fi16 < 15; fi16++ {
            v17 := v1[2147483647]
            fmt.Println(v17, v3, v2)
            _ = fi16
            _ = v17
        }
    } else {
        v2 += v2
    }
    v6 = (^(v6))
    for fi18 := 0; fi18 < 11; fi18++ {
        for fi19 := 0; fi19 < 20; fi19++ {
            if true {
                v3 -= 2
            } else {
                v2 = 0.1
                fmt.Println(fi19, fi18)
            }
            if (float64(-2.25) < (((v2 / v2) + (0.0 - v2)) - v2)) {
                _ = len(v4)
                _ = len(v4)
                v20 := GenS1{F0: (v6 << ((0) & 63))}
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if (65535 > 0) {
                    var v21 int8 = 1
                    _ = v21
                } else {
                    var v22 int64 = 0
                    if ((v6 + (v6 ^ (^(v6)))) >= v6) {
                        _ = len(v1)
                        var v23 float64 = v2
                        switch "" {
                        case "foo bar":
                            if false {
                                var v24 int32 = 42
                                _ = v24
                            }
                        default:
                            var v25 uint = 0
                            _ = v25
                        }
                        _ = v23
                    } else {
                        v22 = v22
                        v6 += v6
                    }
                    v5 -= (v5 << ((1) & 63))
                    _ = v22
                }
                _ = v20
            }
            v26 := GenS2{F0: 1}
            _ = fi19
            _ = v26
        }
        v27 := GenS3{F0: false, F1: 1000, F2: 1}
        _ = fi18
        _ = v27
    }
    v3 -= v3
    _ = len(v4)
    switch "hello" {
    case "\n\t":
        var v28 uint32 = 65535
        _ = v28
    }
    fmt.Println(v2)
    v5 += v5
    v2 -= (float64(0.0) - 3.14159)
    if (42 > v3) {
        for fi29 := 0; fi29 < 15; fi29++ {
            if (!(false || true)) {
                if true {
                    for fi30 := 0; fi30 < 12; fi30++ {
                        v31 := GenS3{F0: ((0.0 <= (0.1 * (((float32(0.1) + (1e10 - 3.14159)) / -2.25) + 0.1))) || true), F1: 0, F2: 902}
                        _ = fi30
                        _ = v31
                    }
                    v32 := GenS2{F0: 32767}
                    _ = v32
                }
            }
            for fi33 := 0; fi33 < 8; fi33++ {
                var v34 int8 = 68
                var v35 int64 = -1
                _ = fi33
                _ = v34
                _ = v35
            }
            _ = fi29
        }
        if ((((v5 & (^(v5))) <= v5) && true) && (v5 == genFn1(4294967295))) {
            var v36 bool = (2147483647 != 4294967295)
            _ = v36
        }
    } else {
        if len(v4) > 0 {
            v4[0] = "hello"
        }
        v5 = (v5 & (v5 ^ (^(v5))))
    }
    if (true || (-1 > 9223372036854775807)) {
        if true {
            var v37 int8 = -41
            v37 += v37
            _ = len(v1)
            _ = v37
        }
    } else {
        v3 = (v3 >> ((580) & 63))
        v2 = (v2 - v2)
        v2 -= v2
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
