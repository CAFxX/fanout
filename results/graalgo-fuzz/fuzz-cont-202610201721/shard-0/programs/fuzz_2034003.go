package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int32; F1 uint }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 863
    var v3 bool = (1 == 0)
    v4 := []int{128, -99, -9223372036854775808}
    var v5 bool = ("" != ("foo bar" + "\n\t"))
    v6 := []string{}
    if v5 {
        if (int(v2) == -1) {
            for fi7 := 0; fi7 < 11; fi7++ {
                _ = len(v1)
                var v8 uint8 = 63
                v3 = true
                _ = fi7
                _ = v8
            }
            v9 := GenS1{F0: 9223372036854775807, F1: 0, F2: 195}
            _ = v9
        }
        if v3 {
            if len(v4) > 0 {
                v4[0] = 252
            }
        } else {
            if (128 <= 42) {
                v1[v2] = v2
                fmt.Println(v2, v2, v5)
            }
            var v10 uint = 18446744073709551615
            for fi11 := 0; fi11 < 2; fi11++ {
                if v5 {
                    v3 = (("a" + "hello") <= "")
                    for fi12 := 0; fi12 < 3; fi12++ {
                        v13 := GenS2{F0: v3, F1: 42}
                        v14 := GenS2{F0: (1.5 != float64(1.5)), F1: 529}
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                } else {
                    v15 := GenS1{F0: 1, F1: 1, F2: v10}
                    _ = v15
                }
                _ = fi11
            }
            _ = v10
        }
    }
    var v16 uint = 4294967295
    var v17 int16 = -1
    fmt.Println(v16, v5, v5)
    fmt.Println(v16)
    if (v17 == v17) {
        _ = len(v1)
        var v18 uint8 = 1
        _ = v18
    } else {
        var v19 uint8 = 255
        v19 += (v19 * 15)
        _ = v19
    }
    v20 := GenS2{F0: v5, F1: (v17 & 1000)}
    v5 = ((0.1 + (((1.5 * -2.25) * -2.25) / float32((1.5 * 3.14159)))) >= -2.25)
    for fi21 := 0; fi21 < 4; fi21++ {
        v3 = v3
        var v22 int8 = -33
        if true {
            switch fi21 {
            case 341:
                fmt.Println(v5, fi21)
            case 32768:
                v1[v2] = v2
            default:
                fmt.Println(fi21, v5)
            }
            var v23 string = ("hello" + "a")
            _ = v23
        }
        _ = fi21
        _ = v22
    }
    switch 128 {
    case -424:
        var v24 int8 = 0
        _ = v24
    default:
        for fi25 := 0; fi25 < 3; fi25++ {
            v3 = v5
            _ = fi25
        }
    }
    var v26 bool = (!(4294967295 > 4294967295))
    for fi27 := 0; fi27 < 20; fi27++ {
        v28 := GenS2{F0: (true || (255 > 42)), F1: v17}
        var v29 int64 = 42
        _ = fi27
        _ = v28
        _ = v29
    }
    fmt.Println(v17, v2)
    v30 := GenS3{F0: 42, F1: 42}
    var v31 uint32 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
