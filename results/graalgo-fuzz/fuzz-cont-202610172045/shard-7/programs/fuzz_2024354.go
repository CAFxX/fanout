package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint8 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 string }
func genFn0(p0 uint64) uint8 {
    p0 += (p0 % ((1) | 1))
    v4 := []int{}
    if true {
        fmt.Println(p0)
        v5 := GenS1{F0: 0, F1: 800, F2: 10}
        _ = v5
    }
    var v6 int64 = 9223372036854775807
    _ = p0
    _ = v4
    _ = v6
    return 255
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]int)
    v3 := make(map[uint32]string)
    var v7 uint8 = genFn0(2)
    if (false && false) {
        v8 := GenS1{F0: 0, F1: -1, F2: genFn0(1)}
        v9 := GenS2{F0: (20 <= 685), F1: 2}
        if true {
            for fi10 := 0; fi10 < 16; fi10++ {
                if (("\n\t" + ((("\n\t" + "a") + "") + ("foo bar" + ""))) <= "") {
                    switch (fi10 / ((fi10) | 1)) {
                    case -156:
                        var v11 uint8 = v7
                        _ = v11
                    default:
                        for fi12 := 0; fi12 < 15; fi12++ {
                            var v13 float32 = 1.5
                            _ = fi12
                            _ = v13
                        }
                    }
                }
                v14 := GenS2{F0: false, F1: 0}
                v7 += (v7 ^ (v7 << ((564) & 63)))
                _ = fi10
                _ = v14
            }
            fmt.Println(v7)
            if false {
                _ = len(v3)
            }
        } else {
            var v15 int16 = 414
            var v16 float64 = 3.14159
            fmt.Println(v15, v7, v16)
            _ = v15
            _ = v16
        }
        _ = v8
        _ = v9
    } else {
        v7 += v7
    }
    if false {
        v17 := GenS3{F0: "a"}
        _ = v17
    }
    var v18 bool = (((0 != 917) && (float32(0.0) <= (0.1 / ((1.5 - (0.1 * 0.0)) - ((3.14159 - 3.14159) - 3.14159))))) && (4294967295 > 18446744073709551615))
    v18 = (v18 && v18)
    if (!v18) {
        for fi19 := 0; fi19 < 5; fi19++ {
            fmt.Println(v7, fi19, v7)
            fmt.Println(v18)
            _ = fi19
        }
    }
    v20 := v1[463]
    for fi21 := 0; fi21 < 16; fi21++ {
        if (445 <= -1) {
            v22 := GenS2{F0: (-223 <= fi21), F1: 32}
            var v23 int8 = 124
            _ = v22
            _ = v23
        }
        var v24 int64 = -1
        _ = fi21
        _ = v24
    }
    if false {
        if (1 != 1) {
            fmt.Println(v7)
            v18 = false
            v18 = false
        } else {
            v25 := GenS1{F0: -1, F1: -710, F2: v7}
            _ = v25
        }
        v26 := GenS3{F0: "\n\t"}
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v7
}
