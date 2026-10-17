package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 string; F1 uint; F2 int32 }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 uint8; F1 int16; F2 uint32 }
func genFn0(p0 bool, p1 int, p2 uint16) int {
    for fi4 := 0; fi4 < 18; fi4++ {
        p2 = 1
        var v5 uint32 = 42
        if p0 {
            v6 := []int{32767, fi4, (^(fi4))}
            _ = v6
        } else {
            v5 = v5
        }
        _ = fi4
        _ = v5
    }
    if (true && (65535 >= 0)) {
        v7 := GenS1{F0: 353, F1: 0}
        _ = v7
    }
    v8 := GenS1{F0: 65535, F1: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 42
}
func genFn1(p0 uint32, p1 uint64, p2 bool) int32 {
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 string = ("\n\t" + ("a" + "hello"))
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        for fi12 := 0; fi12 < 16; fi12++ {
            v13 := []int{}
            if p2 {
                v14 := GenS2{F0: "foo bar", F1: 560, F2: -2147483648}
                _ = v14
            }
            fmt.Println(p2, p0)
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        var v16 uint32 = p0
        var v17 int = (fi15 >> ((36) & 63))
        _ = fi15
        _ = v16
        _ = v17
    }
    fmt.Println(p2)
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = 0
    var v3 int64 = 1
    var v18 int16 = 42
    var v19 bool = ("" >= ("foo bar" + "hello"))
    switch -2147483648 {
    case 1:
        for fi20 := 0; fi20 < 3; fi20++ {
            switch "a" {
            case "":
                if v19 {
                    v21 := GenS1{F0: 4294967295, F1: (v2 ^ (v2 << ((1) & 63)))}
                    var v22 int64 = (v3 * v3)
                    _ = v21
                    _ = v22
                }
            case "a":
                v19 = true
            case "\n\t":
                v23 := GenS1{F0: 65535, F1: genFn1(4294967295, 4294967295, v19)}
                _ = v23
            }
            _ = fi20
        }
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        v25 := []int{fi24}
        _ = fi24
        _ = v25
    }
    if (1 == 872) {
        for fi26 := 0; fi26 < 3; fi26++ {
            var v27 uint64 = 0
            _ = fi26
            _ = v27
        }
        if v19 {
            v2 = v2
            fmt.Println(v2)
            fmt.Println(v19)
        }
        v28 := []int{1, 9223372036854775807, 4294967295, genFn0(((127 == 127) && v19), -9223372036854775808, 1)}
        _ = v28
    }
    if (v19 && (3 < -71)) {
        v29 := GenS3{F0: 42, F1: (^(v3))}
        v30 := []int{genFn0((v3 == -1), genFn0((-2.25 > -2.25), 4294967295, 42), 42), -1, -9223372036854775808, 2147483647}
        for fi31 := 0; fi31 < 20; fi31++ {
            v2 = v2
            var v32 uint16 = 32767
            _ = fi31
            _ = v32
        }
        _ = v29
        _ = v30
    } else {
        fmt.Println(v2, v19)
    }
    var v33 float32 = float32(3.14159)
    for fi34 := 0; fi34 < 12; fi34++ {
        for fi35 := 0; fi35 < 7; fi35++ {
            v3 -= v3
            if v19 {
                v18 = v18
                fmt.Println(v18)
            } else {
                switch "\n\t" {
                case "\n\t":
                    v36 := []int{fi34, genFn0((1e10 > 1e10), (fi35 - -947), 42)}
                    _ = v36
                case "":
                    v37 := v1[32767]
                    _ = v37
                default:
                    if (65535 >= 181) {
                        v3 = -166
                        v1[-1] = fi34
                    } else {
                        var v38 float32 = (1e10 - (3.14159 + (v33 / v33)))
                        var v39 uint = 1
                        v19 = (100 <= -30)
                        _ = v38
                        _ = v39
                    }
                }
                if (v18 == v18) {
                    if false {
                        var v40 string = "a"
                        fmt.Println(v19)
                        _ = v40
                    }
                    v33 = (v33 + (((v33 / ((0.0 + (float32(1.5) * v33)) + v33)) - 0.1) / 1.0))
                    v41 := v1[42]
                    _ = v41
                }
                v3 += 0
            }
            v18 -= (v18 / ((v18) | 1))
            v2 = (v2 | (v2 * (v2 - (v2 * v2))))
            _ = fi35
        }
        _ = fi34
    }
    var v42 uint = 4294967295
    v43 := GenS4{F0: 128, F1: -1, F2: 2147483648}
    switch -9223372036854775808 {
    case -9223372036854775808:
        v44 := GenS3{F0: 0, F1: (-(v3))}
        _ = v44
    case -1:
        v2 = v2
    }
    var v45 uint8 = 29
    switch ("\n\t" + "hello") {
    case "foo bar":
        delete(v1, 128)
    default:
        for fi46 := 0; fi46 < 13; fi46++ {
            v33 += (v33 / 1.5)
            _ = fi46
        }
    }
    v33 += v33
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v42
    _ = v43
    _ = v45
}
