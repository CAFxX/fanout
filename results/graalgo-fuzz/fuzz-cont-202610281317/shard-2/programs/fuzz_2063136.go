package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint; F1 int64; F2 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS5 struct { F0 int64 }
func genFn0(p0 int32) int8 {
    v4 := GenS1{F0: 2, F1: p0}
    for fi5 := 0; fi5 < 14; fi5++ {
        v6 := GenS1{F0: 255, F1: p0}
        var v7 float64 = (1.5 / (-2.25 + 1.5))
        _ = fi5
        _ = v6
        _ = v7
    }
    p0 = p0
    _ = p0
    _ = v4
    return 45
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]bool)
    var v3 int16 = 0
    var v8 int64 = -462
    if false {
        var v9 string = "hello"
        delete(v2, -507)
        _ = v9
    } else {
        var v10 string = ("\n\t" + "")
        v11 := []int{1, 32767, 127}
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 int = (fi12 | (fi12 ^ 42))
            var v14 uint32 = 2147483647
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = v10
        _ = v11
    }
    fmt.Println(v3)
    v3 = v3
    var v15 bool = ((false || true) || (-916 <= (v8 % ((-1) | 1))))
    if ((-60 == genFn0(-2147483648)) || true) {
        var v16 float64 = 3.14159
        fmt.Println(v16)
        v17 := []int{1, 4294967295, 32768, -844}
        _ = v16
        _ = v17
    } else {
        if (!v15) {
            for fi18 := 0; fi18 < 10; fi18++ {
                var v19 bool = (1.5 >= (float32(((0.1 / 0.1) - 1e10)) + 1.5))
                for fi20 := 0; fi20 < 4; fi20++ {
                    var v21 float32 = 1.5
                    if v19 {
                        v22 := GenS1{F0: 4294967295, F1: 2147483647}
                        _ = v22
                    } else {
                        v23 := GenS2{F0: 516, F1: 0, F2: fi20}
                        if true {
                            v19 = v19
                            var v24 uint = 1
                            _ = v24
                        }
                        if (128 <= (fi20 + (-(fi18)))) {
                            v25 := GenS1{F0: 2, F1: -456}
                            if v19 {
                                var v26 int16 = (v3 % (((v3 | (v3 % ((v3) | 1)))) | 1))
                                v8 = v8
                                _ = v26
                            }
                            _ = v25
                        } else {
                            v27 := GenS1{F0: 2, F1: 0}
                            v28 := GenS3{F0: (((42 != 127) || false) || (1.5 <= 0.1))}
                            v15 = v15
                            var v29 int = fi18
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                        _ = v23
                    }
                    switch "hello" {
                    case "hello":
                        v8 -= v8
                    case "":
                        fmt.Println(v3)
                    case "\n\t":
                        fmt.Println(v3, fi18)
                    }
                    _ = fi20
                    _ = v21
                }
                _ = fi18
                _ = v19
            }
            v15 = ((float32(1.5) / (3.14159 + ((float32(((-2.25 / float64(-2.25)) / (1e10 - (float64((0.0 + 3.14159)) - ((1e10 - ((1e10 * (3.14159 - -2.25)) * (0.1 + 3.14159))) * (float64((float32(0.0) + -2.25)) - ((((1.5 - 0.1) / 1e10) / 1e10) * -2.25))))))) + 0.0) + -2.25))) != 1e10)
        }
    }
    if false {
        v3 = v3
        v30 := GenS2{F0: 833, F1: v8, F2: 0}
        _ = v30
    }
    fmt.Println(v15, v3, v3)
    if (!("a" >= "foo bar")) {
        if v15 {
            if true {
                var v31 int = 2147483647
                _ = v31
            }
            switch 659 {
            case -2147483648:
                v15 = v15
            }
        } else {
            v3 = (v3 % ((v3) | 1))
        }
        v3 += (-(v3))
        for fi32 := 0; fi32 < 19; fi32++ {
            v33 := GenS4{F0: 65535, F1: (112 <= 2), F2: 18446744073709551615}
            v2[fi32] = (2 <= 42)
            _ = fi32
            _ = v33
        }
    }
    v34 := GenS5{F0: 42}
    var v35 int16 = (v3 >> ((136) & 63))
    var v36 int64 = (v8 & 9223372036854775807)
    var v37 int8 = 100
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v8
}
