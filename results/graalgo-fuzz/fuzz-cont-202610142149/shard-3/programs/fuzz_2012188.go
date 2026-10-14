package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int }
type GenS2 struct { F0 uint8; F1 int }
func genFn0(p0 int32, p1 uint) int16 {
    p1 -= (p1 & p1)
    p0 += p0
    if (0 < 680) {
        fmt.Println(p1, p0)
        var v4 bool = ((224 < 1000) && false)
        _ = v4
    }
    var v5 uint8 = 127
    _ = p0
    _ = p1
    _ = v5
    return 213
}

func main() {
    var v1 float32 = -2.25
    var v2 string = ""
    v3 := make(map[int]int)
    if false {
        v3[-830] = 2147483647
        for fi6 := 0; fi6 < 7; fi6++ {
            for fi7 := 0; fi7 < 1; fi7++ {
                var v8 uint64 = 4294967295
                _ = len(v3)
                v9 := GenS1{F0: fi7, F1: v2, F2: (^(fi6))}
                _ = fi7
                _ = v8
                _ = v9
            }
            if (!(2147483647 != 0)) {
                v2 = v2
                var v10 int8 = 0
                var v11 int = 4294967295
                _ = v10
                _ = v11
            } else {
                if false {
                    if false {
                        for fi12 := 0; fi12 < 15; fi12++ {
                            switch (((v2 + "a") + "hello") + ("hello" + v2)) {
                            case "":
                                var v13 uint = 58
                                _ = v13
                            }
                            var v14 bool = (42 <= 88)
                            v15 := GenS2{F0: 0, F1: fi6}
                            _ = fi12
                            _ = v14
                            _ = v15
                        }
                    }
                }
                v16 := GenS2{F0: 2, F1: (fi6 * 32767)}
                v2 = v2
                _ = v16
            }
            _ = fi6
        }
        switch (((v2 + (v2 + (v2 + ((v2 + "hello") + v2)))) + v2) + ((v2 + v2) + v2)) {
        case "":
            v1 = v1
        default:
            v1 += (v1 - (1.5 / v1))
        }
    }
    fmt.Println(v2)
    var v17 bool = true
    v17 = false
    var v18 uint8 = 45
    for fi19 := 0; fi19 < 9; fi19++ {
        var v20 uint16 = 42
        _ = fi19
        _ = v20
    }
    v21 := GenS1{F0: 127, F1: "\n\t", F2: -750}
    var v22 bool = (float32(1e10) == (1.5 * (-2.25 * v1)))
    fmt.Println(v22, v18)
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
