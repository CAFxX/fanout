package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int; F1 uint8 }
func genFn0() int8 {
    var v5 float32 = ((float32(1.5) - float32(1e10)) / ((float32(-2.25) - 1.5) * 1e10))
    v5 -= (v5 / ((v5 - (((0.0 / v5) * v5) - 0.0)) - 1.5))
    var v6 uint32 = 0
    var v7 uint = 4294967295
    _ = v5
    _ = v6
    _ = v7
    return 100
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = -914
    var v3 int16 = v2
    v4 := make(map[string]string)
    v8 := []int{}
    fmt.Println(v2, v3)
    v2 = (v2 >> ((0) & 63))
    v3 = v2
    if len(v8) > 0 {
        v8[0] = 236
    }
    v9 := GenS1{F0: "\n\t", F1: 42}
    var v10 uint32 = 652
    for fi11 := 0; fi11 < 6; fi11++ {
        if false {
            for fi12 := 0; fi12 < 19; fi12++ {
                if (!true) {
                    if true {
                        v10 = v10
                    }
                    v13 := GenS1{F0: "a", F1: -1}
                    _ = v13
                } else {
                    if false {
                        v14 := GenS1{F0: "hello", F1: 980}
                        v10 += (v10 >> ((65535) & 63))
                        _ = len(v4)
                        v2 = v2
                        switch "\n\t" {
                        case "a":
                            switch fi11 {
                            case 794:
                                v15 := v1[(fi12 >> ((743) & 63))]
                                if true {
                                    fmt.Println(v2)
                                    fmt.Println(v3)
                                    v16 := GenS1{F0: ("" + "a"), F1: -683}
                                    _ = v16
                                }
                                _ = v15
                            }
                        case "":
                            fmt.Println(v10)
                        }
                        _ = v14
                    } else {
                        var v17 float32 = 1e10
                        v8 = append(v8, (-(fi12)))
                        _ = v17
                    }
                }
                _ = fi12
            }
            v3 -= (v3 % ((v2) | 1))
        }
        _ = fi11
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        v19 := v1[-190]
        _ = fi18
        _ = v19
    }
    v20 := GenS1{F0: ("hello" + "hello"), F1: 0}
    for fi21 := 0; fi21 < 10; fi21++ {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = len(v4)
        switch fi21 {
        case 32767:
            var v22 uint = 723
            _ = v22
        case 2147483647:
            v2 = v2
        case 1:
            var v23 float32 = (-2.25 - float32(0.0))
            _ = v23
        }
        _ = fi21
    }
    for fi24 := 0; fi24 < 14; fi24++ {
        v10 = 1
        v10 = v10
        _ = fi24
    }
    v3 -= 32767
    v25 := GenS2{F0: -205, F1: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
