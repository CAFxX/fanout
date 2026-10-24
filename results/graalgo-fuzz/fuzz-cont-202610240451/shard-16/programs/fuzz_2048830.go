package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 string; F1 uint; F2 int16 }
func genFn0() uint {
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 int = 128
        _ = fi7
        _ = v8
    }
    var v9 float32 = 1e10
    for fi10 := 0; fi10 < 9; fi10++ {
        v11 := []int{282, (fi10 - fi10), fi10}
        _ = fi10
        _ = v11
    }
    fmt.Println(v9)
    _ = v9
    return 255
}

func main() {
    var v1 uint64 = 2
    var v2 int = 9223372036854775807
    v3 := []uint32{}
    var v4 int32 = 1
    v5 := make(map[string]bool)
    var v6 uint32 = 2
    _ = len(v5)
    v1 = v1
    v2 = int(v1)
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 float32 = (((-2.25 + 1e10) - float32(0.1)) - 3.14159)
        v14 := GenS1{F0: (("hello" + "\n\t") + ""), F1: 107}
        for fi15 := 0; fi15 < 20; fi15++ {
            var v16 int8 = 100
            _ = fi15
            _ = v16
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    switch (("hello" + ("" + "foo bar")) + "a") {
    case "hello":
        v3 = append(v3, v6)
    case "":
        v17 := GenS1{F0: "\n\t", F1: 128}
        _ = v17
    default:
        switch (^(v2)) {
        case -979:
            if false {
                if false {
                    var v18 float32 = 1e10
                    fmt.Println(v18, v1)
                    _ = v18
                }
            } else {
                fmt.Println(v6, v1)
                v1 -= v1
                v6 = v6
            }
        case -206:
            v2 = v2
        default:
            var v19 uint32 = 2
            _ = v19
        }
    }
    var v20 bool = true
    fmt.Println(v20, v20)
    var v21 float64 = 0.1
    for fi22 := 0; fi22 < 5; fi22++ {
        v21 = -2.25
        _ = fi22
    }
    v1 -= 642
    v1 += 19
    _ = len(v5)
    if v20 {
        v20 = (1 > 127)
        if v20 {
            switch (^(v2)) {
            case 0:
                switch (v2 / ((9223372036854775807) | 1)) {
                case 0:
                    if (("foo bar" + "") >= "\n\t") {
                        var v23 uint8 = 2
                        _ = len(v5)
                        fmt.Println(v20, v23, v2)
                        _ = v23
                    } else {
                        for fi24 := 0; fi24 < 1; fi24++ {
                            var v25 uint8 = 255
                            _ = fi24
                            _ = v25
                        }
                    }
                case 4294967295:
                    v21 -= 0.1
                }
            default:
                _ = len(v3)
            }
            v2 -= v2
            var v26 uint16 = 0
            _ = v26
        }
        v20 = ((true || (42 >= -233)) || (!v20))
    } else {
        v27 := GenS2{F0: "\n\t", F1: 255, F2: 0}
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
