package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 bool }
func genFn0(p0 string, p1 bool, p2 uint16) int32 {
    p2 = p2
    for fi6 := 0; fi6 < 16; fi6++ {
        switch p0 {
        case "a":
            fmt.Println(p0)
        case "hello":
            if p1 {
                var v7 int = 2147483647
                var v8 float64 = (-2.25 - 1e10)
                _ = v7
                _ = v8
            } else {
                p0 = ("foo bar" + "hello")
            }
        default:
            if (true || (0.0 != (((-2.25 - (1e10 / (0.0 + 3.14159))) + (float32((1.5 + 1e10)) / 1.0)) + (0.0 + ((float32(-2.25) * (float32(-2.25) / 1e10)) + ((((3.14159 + (3.14159 + 1e10)) - 1.5) - ((1.5 + (float32((0.1 * -2.25)) + 3.14159)) / 1.0)) + -2.25)))))) {
                p1 = (((1.5 + 0.1) - float64(float32(1.5))) != 1.5)
                var v9 int64 = -9223372036854775808
                _ = v9
            }
        }
        p1 = false
        v10 := GenS1{F0: p2, F1: p1}
        _ = fi6
        _ = v10
    }
    p1 = (463 != -2147483648)
    var v11 uint8 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return int32(v11)
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = 1
    v3 := make(map[int]int)
    var v4 string = "foo bar"
    v5 := make(map[int]string)
    if true {
        for fi12 := 0; fi12 < 10; fi12++ {
            if ((4294967295 > 0) || false) {
                _ = len(v3)
                fmt.Println(v4)
            }
            _ = fi12
        }
        var v13 string = v4
        v4 = v13
        _ = v13
    }
    v4 = v4
    switch -9223372036854775808 {
    case 1:
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 uint32 = 408
            var v16 uint16 = 0
            _ = fi14
            _ = v15
            _ = v16
        }
    case 4294967295:
        v4 = "a"
    default:
        v17 := GenS1{F0: 649, F1: true}
        _ = v17
    }
    _ = len(v3)
    for fi18 := 0; fi18 < 19; fi18++ {
        var v19 float64 = -2.25
        switch (("hello" + "") + ("hello" + (v4 + ((v4 + v4) + v4)))) {
        case "\n\t":
            v20 := v1[2147483647]
            _ = v20
        default:
            if true {
                v21 := []int{fi18, 42}
                switch ((v4 + v4) + (v4 + v4)) {
                case "hello":
                    if true {
                        var v22 float32 = (0.0 / 0.1)
                        _ = v22
                    } else {
                        var v23 uint8 = 130
                        v19 = v19
                        v24 := GenS1{F0: 65535, F1: true}
                        _ = v23
                        _ = v24
                    }
                case "foo bar":
                    v2 -= 1
                case "a":
                    if true {
                        if (false && true) {
                            var v25 uint = 0
                            _ = v25
                        } else {
                            v2 = v2
                            v26 := GenS1{F0: 65535, F1: (564 > 1)}
                            _ = v26
                        }
                    }
                }
                _ = v21
            } else {
                v1[671] = v4
                switch fi18 {
                case 4294967295:
                    if true {
                        v27 := GenS2{F0: true}
                        v19 = 1.5
                        _ = v27
                    }
                }
                v19 -= ((v19 + float64(3.14159)) + v19)
            }
        }
        _ = fi18
        _ = v19
    }
    v28 := v3[32768]
    v3[v28] = (v28 ^ int(v2))
    for fi29 := 0; fi29 < 1; fi29++ {
        var v30 bool = (-491 > -1)
        _ = fi29
        _ = v30
    }
    var v31 float64 = 1.5
    var v32 uint16 = 2
    var v33 float64 = v31
    v34 := []int{(-(v28))}
    delete(v5, v28)
    var v35 uint16 = 96
    if (!true) {
        var v36 uint64 = 676
        fmt.Println(v31)
        v2 = -759
        _ = v36
    } else {
        fmt.Println(v4)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
