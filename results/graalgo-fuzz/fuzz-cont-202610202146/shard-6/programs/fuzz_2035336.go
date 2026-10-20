package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int }
type GenS2 struct { F0 int64; F1 string }
func genFn0() uint64 {
    switch 562 {
    case 75:
        if (!(0.0 <= (-2.25 - 3.14159))) {
            for fi4 := 0; fi4 < 20; fi4++ {
                v5 := GenS1{F0: 42, F1: 42, F2: fi4}
                _ = fi4
                _ = v5
            }
            var v6 uint32 = 4294967295
            v6 = (^(v6))
            _ = v6
        }
    }
    var v7 float64 = float64((0.0 + (1e10 + (-2.25 * 0.0))))
    v7 += v7
    if true {
        v7 -= v7
        var v8 string = "hello"
        _ = v8
    } else {
        var v9 int32 = -1
        v9 -= v9
        for fi10 := 0; fi10 < 20; fi10++ {
            if true {
                if true {
                    var v11 string = (((("a" + "a") + "a") + "a") + "\n\t")
                    v11 = v11
                    v11 = ((v11 + "a") + v11)
                    _ = v11
                }
                v7 += (float64(1.5) * v7)
            }
            if ((307 < int16(v9)) && false) {
                var v12 float32 = 0.1
                _ = v12
            }
            var v13 uint16 = 42
            _ = fi10
            _ = v13
        }
        _ = v9
    }
    _ = v7
    return 203
}

func main() {
    v1 := []int{-1}
    v2 := make(map[uint32]string)
    v3 := make(map[int]int)
    var v14 float32 = 1.5
    _ = len(v3)
    fmt.Println(v14)
    _ = len(v3)
    v14 = (((float32(((1e10 + 0.1) - 1e10)) + float32((float64(0.0) / ((((0.0 - 0.0) + ((1.5 * 0.1) + (1.5 * (1.5 + -2.25)))) * (float64(v14) - -2.25)) - (0.1 + 1e10))))) / v14) / 1e10)
    var v15 uint64 = 4294967295
    v14 -= (3.14159 - 1.5)
    for fi16 := 0; fi16 < 1; fi16++ {
        var v17 int = (fi16 * 1)
        _ = fi16
        _ = v17
    }
    v14 = (3.14159 / ((v14 + v14) - v14))
    fmt.Println(v14)
    for fi18 := 0; fi18 < 19; fi18++ {
        for fi19 := 0; fi19 < 12; fi19++ {
            if true {
                v20 := GenS2{F0: -1, F1: (("a" + ("foo bar" + "hello")) + "a")}
                _ = v20
            }
            fmt.Println(v14)
            _ = fi19
        }
        v14 -= float32((-2.25 - 1.5))
        var v21 uint = 186
        _ = fi18
        _ = v21
    }
    v15 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
