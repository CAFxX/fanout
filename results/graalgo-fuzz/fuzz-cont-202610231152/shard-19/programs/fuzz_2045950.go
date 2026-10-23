package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64 }
func genFn0(p0 uint32) int64 {
    if true {
        var v5 int64 = 1
        for fi6 := 0; fi6 < 8; fi6++ {
            if (1 >= 127) {
                p0 += (^(p0))
                v7 := []int{(fi6 ^ 127), fi6}
                _ = v7
            } else {
                for fi8 := 0; fi8 < 14; fi8++ {
                    if (true && true) {
                        v9 := []int{(^(fi8))}
                        v10 := GenS1{F0: true, F1: p0, F2: 2}
                        switch "a" {
                        case "foo bar":
                            v5 += v5
                        case "a":
                            v5 = (v5 >> ((340) & 63))
                        }
                        var v11 float64 = -2.25
                        v5 -= (v5 ^ (v5 ^ -9223372036854775808))
                        var v12 int16 = 787
                        _ = v10
                        _ = v11
                        _ = v12
                        _ = v9
                    } else {
                        fmt.Println(fi6, p0)
                    }
                    v13 := []int{(fi6 + (fi6 / (((fi6 / ((fi6) | 1))) | 1)))}
                    _ = fi8
                    _ = v13
                }
                var v14 bool = (39 != 54)
                v15 := GenS1{F0: ((3.14159 == float64(float32(0.1))) && (true && true)), F1: (p0 * (p0 ^ (p0 * 2147483647))), F2: 42}
                _ = v14
                _ = v15
            }
            fmt.Println(p0, fi6, fi6)
            _ = fi6
        }
        v16 := GenS2{F0: 513}
        _ = v16
        _ = v5
    } else {
        for fi17 := 0; fi17 < 16; fi17++ {
            var v18 uint16 = 1
            _ = fi17
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        fmt.Println(p0, p0)
        var v20 int64 = 0
        _ = fi19
        _ = v20
    }
    p0 -= 2
    fmt.Println(p0)
    v21 := []int{}
    _ = p0
    _ = v21
    return 42
}

func main() {
    var v1 uint8 = 42
    v2 := []int{}
    v3 := []string{("foo bar" + "a"), "hello", (("foo bar" + "hello") + "hello")}
    var v4 uint = 0
    if (-128 >= 42) {
        var v22 uint32 = 1
        _ = v22
    } else {
        for fi23 := 0; fi23 < 3; fi23++ {
            for fi24 := 0; fi24 < 6; fi24++ {
                switch fi24 {
                case 0:
                    var v25 int8 = int8(v4)
                    _ = v25
                case -1:
                    v4 -= 65535
                default:
                    v3 = append(v3, "foo bar")
                }
                switch fi23 {
                case -2147483648:
                    if (9223372036854775807 > genFn0(4294967295)) {
                        var v26 uint16 = 505
                        if true {
                            fmt.Println(fi23, fi23)
                        }
                        v27 := GenS1{F0: (v4 >= (v4 - v4)), F1: 242, F2: (v1 & (^(v1)))}
                        _ = v26
                        _ = v27
                    }
                case -20:
                    v1 -= 45
                default:
                    var v28 int = (fi24 * fi23)
                    _ = v28
                }
                _ = fi24
            }
            _ = fi23
        }
        fmt.Println(v4)
    }
    v1 = (v1 << ((v4) & 63))
    switch -234 {
    case 4294967295:
        v4 = (v4 | v4)
    case 0:
        v29 := GenS3{F0: 42}
        _ = v29
    case 42:
        fmt.Println(v1, v1)
    default:
        fmt.Println(v4)
    }
    v1 = v1
    v1 += v1
    var v30 float64 = float64(float32(0.1))
    for fi31 := 0; fi31 < 10; fi31++ {
        v32 := GenS4{F0: int64(fi31)}
        _ = fi31
        _ = v32
    }
    for fi33 := 0; fi33 < 11; fi33++ {
        var v34 uint64 = 18446744073709551615
        if ((true || ((65535 != 412) && false)) || true) {
            var v35 uint64 = (v34 - (v34 * 42))
            _ = v35
        } else {
            if (false && ((!(-32768 >= 1)) || (false && (0 != -65)))) {
                var v36 bool = false
                _ = v36
            } else {
                v4 -= v4
                var v37 int = (fi33 * (fi33 + (fi33 << ((18446744073709551615) & 63))))
                _ = v37
            }
        }
        v30 -= ((float64(0.0) - float64(0.1)) - v30)
        _ = fi33
        _ = v34
    }
    var v38 int16 = 808
    v3 = append(v3, "")
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v39 string = ("a" + (("foo bar" + "\n\t") + "\n\t"))
    fmt.Println(v39, v4, v38)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v38
    _ = v39
    _ = v4
}
