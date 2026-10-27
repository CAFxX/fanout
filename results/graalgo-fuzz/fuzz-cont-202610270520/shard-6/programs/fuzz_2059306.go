package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int64 }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 197
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]string)
    v6 := GenS1{F0: 1000, F1: ("foo bar" + (("foo bar" + "foo bar") + "a")), F2: 1}
    v2 = v2
    fmt.Println(v2)
    var v7 float64 = (0.0 / 1.5)
    v8 := []int{v2, v2, (v2 & (v2 & (v2 + -9223372036854775808))), -9223372036854775808}
    v2 = v2
    _ = len(v1)
    switch (v2 % (((^(v2))) | 1)) {
    case -160:
        if false {
            if false {
                if (true && false) {
                    switch 334 {
                    case -38:
                        for fi9 := 0; fi9 < 3; fi9++ {
                            switch 32767 {
                            case 0:
                                v10 := GenS1{F0: 1, F1: "hello", F2: -111}
                                _ = v10
                            case -2147483648:
                                fmt.Println(v7)
                            case -1:
                                v2 -= 0
                            }
                            v11 := v4[288]
                            _ = fi9
                            _ = v11
                        }
                    }
                    for fi12 := 0; fi12 < 3; fi12++ {
                        fmt.Println(fi12, v2)
                        _ = fi12
                    }
                } else {
                    if len(v8) > 0 {
                        v8[0] = v2
                    }
                }
                switch 661 {
                case 518:
                    var v13 int = 9223372036854775807
                    _ = v13
                case 32768:
                    v7 -= 3.14159
                case 42:
                    var v14 uint = 237
                    _ = v14
                default:
                    fmt.Println(v2, v2)
                }
                v2 -= 128
            }
            v15 := v5[(^(v2))]
            _ = v15
        }
    case 1:
        v16 := GenS2{F0: 703}
        _ = v16
    }
    v8 = append(v8, 977)
    var v17 float32 = 1.5
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    fmt.Println(v17, v17, v17)
    v2 = 646
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
