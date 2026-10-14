package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
func genFn0() int8 {
    var v6 bool = false
    fmt.Println(v6)
    _ = v6
    return 127
}
func genFn1() int {
    var v7 uint16 = 865
    var v8 int = 0
    _ = v7
    _ = v8
    return v8
}

func main() {
    var v1 uint8 = 127
    v2 := []uint32{300}
    var v3 uint = 42
    var v4 int16 = 495
    var v5 int16 = (-(v4))
    v9 := GenS1{F0: "hello", F1: -128}
    v5 += -1
    var v10 int32 = 0
    switch 42 {
    case 567:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        switch 85 {
        case -1:
            _ = len(v2)
        case -9223372036854775808:
            v11 := GenS1{F0: "hello", F1: 43}
            _ = v11
        case 388:
            if len(v2) > 0 {
                v2[0] = 42
            }
        default:
            v3 += 810
        }
    }
    v10 += (v10 - 874)
    if (4294967295 < 875) {
        for fi12 := 0; fi12 < 10; fi12++ {
            v13 := GenS1{F0: ("" + "foo bar"), F1: 10}
            var v14 int64 = 798
            var v15 int16 = (^(v5))
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
    }
    v3 = (v3 & v3)
    if (9223372036854775807 <= 42) {
        var v16 int16 = 0
        v17 := GenS1{F0: "a", F1: genFn0()}
        _ = v16
        _ = v17
    }
    var v18 uint8 = (v1 & (v1 ^ 255))
    var v19 uint16 = 424
    v2 = append(v2, 2)
    for fi20 := 0; fi20 < 15; fi20++ {
        for fi21 := 0; fi21 < 3; fi21++ {
            v1 = v1
            if len(v2) > 0 {
                v2[0] = uint32(v19)
            }
            if true {
                var v22 uint32 = 0
                if true {
                    fmt.Println(v19, fi20, v19)
                }
                _ = v22
            }
            _ = fi21
        }
        switch fi20 {
        case 387:
            var v23 bool = ((!false) && (89 < 42))
            _ = v23
        case 4294967295:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = fi20
    }
    v24 := GenS1{F0: "foo bar", F1: -57}
    v19 = (v19 / ((v19) | 1))
    fmt.Println(v10)
    v2 = append(v2, 301)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
