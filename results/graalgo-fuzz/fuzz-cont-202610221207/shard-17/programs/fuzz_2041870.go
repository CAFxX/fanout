package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 uint64; F1 bool; F2 bool }
func genFn0(p0 uint64, p1 bool) uint32 {
    p1 = true
    var v5 int8 = 42
    var v6 bool = p1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 2147483648
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint32 = 2147483647
    var v3 float32 = (0.1 / 1.0)
    var v4 uint16 = 65535
    var v7 int = 574
    v2 = uint32(v4)
    var v8 int16 = -32768
    var v9 uint8 = 255
    v2 = uint32(v4)
    v10 := []int{-9223372036854775808}
    for fi11 := 0; fi11 < 15; fi11++ {
        if true {
            for fi12 := 0; fi12 < 2; fi12++ {
                _ = len(v1)
                _ = fi12
            }
        } else {
            var v13 uint = 4294967295
            _ = v13
        }
        _ = fi11
    }
    v3 = float32(1.5)
    var v14 uint32 = v2
    var v15 float64 = (-2.25 / 1.0)
    for fi16 := 0; fi16 < 19; fi16++ {
        fmt.Println(v15, v2, v3)
        switch ("foo bar" + ("foo bar" + "")) {
        case "a":
            if (!(!true)) {
                var v17 float32 = float32((float64(0.1) * (1.5 / float64(1.5))))
                _ = v17
            }
        case "\n\t":
            v2 -= (^(v14))
        case "":
            if true {
                v9 += v9
                fmt.Println(v4, v7)
                switch "foo bar" {
                case "a":
                    if ("" >= "a") {
                        v15 += (float64(1e10) - 3.14159)
                        if true {
                            v4 -= 79
                        } else {
                            var v18 int64 = -85
                            v8 = (v8 >> ((386) & 63))
                            _ = v18
                        }
                        var v19 int = (v7 & fi16)
                        _ = v19
                    } else {
                        for fi20 := 0; fi20 < 17; fi20++ {
                            v21 := GenS1{F0: (^(v9)), F1: ("a" + "")}
                            _ = fi20
                            _ = v21
                        }
                    }
                }
            } else {
                for fi22 := 0; fi22 < 3; fi22++ {
                    v23 := v1[v14]
                    fmt.Println(v14)
                    fmt.Println(v14)
                    _ = fi22
                    _ = v23
                }
                if (42 == 727) {
                    for fi24 := 0; fi24 < 11; fi24++ {
                        var v25 float64 = 0.1
                        _ = fi24
                        _ = v25
                    }
                }
            }
        }
        v26 := GenS1{F0: v9, F1: ("foo bar" + (("" + "") + ("" + "hello")))}
        _ = fi16
        _ = v26
    }
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    v1[(v14 ^ v2)] = (-(v7))
    v27 := GenS2{F0: 952, F1: true, F2: (v4 != v4)}
    v3 -= (float32((((float64(v3) * v15) * float64(0.0)) * v15)) - v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
