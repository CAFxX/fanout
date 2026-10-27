package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string; F1 uint32; F2 uint8 }
type GenS6 struct { F0 uint32; F1 uint }
type GenS7 struct { F0 uint; F1 int8; F2 bool }
func genFn0(p0 int16, p1 int64, p2 uint16) uint {
    for fi5 := 0; fi5 < 13; fi5++ {
        v6 := []int{(fi5 << ((4294967295) & 63)), (fi5 | (fi5 % ((fi5) | 1))), 42, (-(fi5))}
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := GenS1{F0: 107}
        var v9 string = ""
        _ = fi7
        _ = v8
        _ = v9
    }
    switch 32767 {
    case 42:
        if (2 > 0) {
            p1 = (p1 * p1)
            v10 := GenS2{F0: 18446744073709551615}
            if (18446744073709551615 < 322) {
                v11 := GenS2{F0: 255}
                if true {
                    for fi12 := 0; fi12 < 20; fi12++ {
                        p2 = (p2 + 0)
                        var v13 float32 = 1.5
                        p2 = 32767
                        _ = fi12
                        _ = v13
                    }
                    var v14 uint8 = 128
                    _ = v14
                }
                _ = v11
            }
            _ = v10
        }
    case 522:
        if false {
            p2 = 42
        }
    case -306:
        v15 := GenS3{F0: ("hello" + "foo bar"), F1: 57}
        _ = v15
    default:
        var v16 uint32 = 909
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    return 65535
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = 345
    v3 := make(map[uint32]string)
    var v4 int8 = 127
    if true {
        var v17 string = "a"
        v18 := GenS1{F0: -1}
        _ = v17
        _ = v18
    }
    var v19 uint32 = 4294967295
    var v20 uint8 = 42
    v21 := []int{32767, 1, 2147483647}
    delete(v1, (v19 ^ v19))
    for fi22 := 0; fi22 < 19; fi22++ {
        if true {
            var v23 uint64 = 2
            if (!(int8(v23) > (v4 & 0))) {
                var v24 uint16 = uint16(v19)
                _ = v24
            }
            var v25 uint64 = v23
            _ = v23
            _ = v25
        } else {
            var v26 float32 = (3.14159 + 0.1)
            var v27 int64 = -9223372036854775808
            v28 := GenS3{F0: ("\n\t" + "a"), F1: (v4 ^ (v4 << ((65535) & 63)))}
            _ = v26
            _ = v27
            _ = v28
        }
        v20 = (v20 | (v20 - v20))
        if (false || false) {
            var v29 int16 = (v2 << ((4294967295) & 63))
            if (470 == 1) {
                if true {
                    v29 += v29
                }
            } else {
                delete(v3, (v19 >> ((0) & 63)))
                var v30 float32 = -2.25
                if (!false) {
                    v29 = v29
                    v20 += v20
                } else {
                    v2 = -206
                    fmt.Println(v4, fi22)
                    for fi31 := 0; fi31 < 17; fi31++ {
                        v30 -= (float32(0.1) / v30)
                        if (!false) {
                            v30 = 3.14159
                            v19 -= v19
                        }
                        v32 := GenS4{F0: (^(v19))}
                        _ = fi31
                        _ = v32
                    }
                }
                _ = v30
            }
            _ = v29
        }
        _ = fi22
    }
    fmt.Println(v4)
    if (-1 <= 2147483647) {
        for fi33 := 0; fi33 < 1; fi33++ {
            fmt.Println(v2, v19, v20)
            if (0.1 == (-2.25 - -2.25)) {
                var v34 float32 = (1.5 / 3.14159)
                _ = v34
            }
            _ = fi33
        }
        if (976 == 9223372036854775807) {
            v35 := GenS5{F0: (("\n\t" + "") + "a"), F1: v19, F2: 0}
            v36 := GenS6{F0: 1, F1: 1}
            v37 := GenS7{F0: 18446744073709551615, F1: v4, F2: true}
            _ = v35
            _ = v36
            _ = v37
        }
        var v38 int64 = 690
        _ = v38
    }
    fmt.Println(v19, v2)
    _ = len(v1)
    for fi39 := 0; fi39 < 17; fi39++ {
        for fi40 := 0; fi40 < 6; fi40++ {
            var v41 bool = (32767 == 430)
            _ = fi40
            _ = v41
        }
        _ = fi39
    }
    switch ("foo bar" + ("hello" + ("foo bar" + ("hello" + "a")))) {
    case "hello":
        v42 := v3[v19]
        _ = v42
    }
    var v43 float64 = ((0.0 - ((0.0 * (3.14159 / 3.14159)) + 3.14159)) - 1e10)
    v20 += (v20 ^ 1)
    if len(v21) > 0 {
        v21[0] = 128
    }
    v2 = (v2 | (v2 * (^(v2))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v43
}
