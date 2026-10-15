package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 string }
type GenS2 struct { F0 int32; F1 int8; F2 int8 }
type GenS3 struct { F0 uint16; F1 int64; F2 int8 }
func genFn0(p0 string) uint64 {
    var v4 string = (p0 + p0)
    if (0 != 100) {
        v4 = "hello"
        p0 = p0
    } else {
        v5 := GenS1{F0: 925, F1: 18446744073709551615, F2: ((("\n\t" + (p0 + v4)) + (v4 + ("foo bar" + v4))) + v4)}
        for fi6 := 0; fi6 < 12; fi6++ {
            fmt.Println(p0, v4, fi6)
            _ = fi6
        }
        for fi7 := 0; fi7 < 18; fi7++ {
            v8 := GenS1{F0: -216, F1: 1, F2: v4}
            v9 := GenS2{F0: -469, F1: 42, F2: -128}
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = v5
    }
    v10 := []int{-2147483648, -258}
    var v11 uint32 = 2147483648
    _ = len(v10)
    _ = p0
    _ = v10
    _ = v11
    _ = v4
    return 2
}
func genFn1(p0 string, p1 uint64) int16 {
    for fi12 := 0; fi12 < 1; fi12++ {
        v13 := GenS2{F0: -1, F1: -24, F2: 109}
        if true {
            p1 = genFn0((p0 + p0))
        }
        p1 += (p1 / ((p1) | 1))
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        if false {
            switch (((p0 + (p0 + p0)) + "hello") + "foo bar") {
            case "":
                if false {
                    switch 122 {
                    case 42:
                        fmt.Println(p0, fi14)
                    default:
                        fmt.Println(p1, p0, fi14)
                    }
                    for fi15 := 0; fi15 < 12; fi15++ {
                        var v16 float32 = -2.25
                        _ = fi15
                        _ = v16
                    }
                }
            default:
                p1 += (p1 * (p1 * 293))
            }
            p0 = (p0 + p0)
            var v17 uint16 = 918
            _ = v17
        }
        p0 = p0
        var v18 string = (p0 + p0)
        _ = fi14
        _ = v18
    }
    var v19 uint64 = p1
    p1 += genFn0(((p0 + "\n\t") + ("foo bar" + p0)))
    fmt.Println(v19, p0)
    _ = p0
    _ = p1
    _ = v19
    return 1
}

func main() {
    var v1 float64 = 1.5
    var v2 int = 2147483647
    v3 := make(map[int]int)
    var v20 int64 = 42
    for fi21 := 0; fi21 < 2; fi21++ {
        delete(v3, fi21)
        _ = fi21
    }
    v1 += v1
    if true {
        v22 := GenS1{F0: -909, F1: 18446744073709551615, F2: "hello"}
        v23 := []int{682, (v2 % ((0) | 1)), v2, v2}
        if (!false) {
            v2 = 0
            fmt.Println(v2)
        } else {
            v24 := GenS1{F0: -1, F1: 18446744073709551615, F2: "a"}
            for fi25 := 0; fi25 < 2; fi25++ {
                v3[(v2 << ((65535) & 63))] = v2
                var v26 int8 = 73
                _ = fi25
                _ = v26
            }
            if false {
                if (0.1 != 0.0) {
                    var v27 uint64 = 330
                    if (v1 != v1) {
                        v27 = 1
                    } else {
                        v28 := GenS3{F0: 255, F1: (-(v20)), F2: -44}
                        _ = v28
                    }
                    v29 := GenS3{F0: 895, F1: (v20 & v20), F2: 42}
                    _ = v27
                    _ = v29
                } else {
                    for fi30 := 0; fi30 < 13; fi30++ {
                        var v31 float64 = v1
                        if true {
                            fmt.Println(v2)
                        }
                        fmt.Println(v2, v20, fi30)
                        _ = fi30
                        _ = v31
                    }
                    v20 += v20
                    fmt.Println(v20, v2)
                }
                var v32 int64 = 1
                _ = v32
            }
            _ = v24
        }
        _ = v22
        _ = v23
    }
    v20 = v20
    delete(v3, (v2 | 0))
    for fi33 := 0; fi33 < 4; fi33++ {
        v1 = v1
        _ = fi33
    }
    v2 -= v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v20
    _ = v3
}
