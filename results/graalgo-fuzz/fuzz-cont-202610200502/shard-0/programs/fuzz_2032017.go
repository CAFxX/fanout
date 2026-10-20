package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 uint64; F2 int64 }
func genFn0(p0 bool) bool {
    fmt.Println(p0)
    var v5 int64 = 613
    v6 := GenS1{F0: 0}
    var v7 int32 = -893
    fmt.Println(p0, p0)
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return false
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 18446744073709551615
    var v3 float32 = 0.1
    v4 := make(map[int]string)
    switch "a" {
    case "":
        var v8 int64 = -466
        _ = v8
    case "foo bar":
        for fi9 := 0; fi9 < 5; fi9++ {
            var v10 int8 = -1
            _ = fi9
            _ = v10
        }
    case "a":
        switch "foo bar" {
        case "\n\t":
            v2 = v2
        default:
            for fi11 := 0; fi11 < 11; fi11++ {
                var v12 int16 = -288
                _ = fi11
                _ = v12
            }
        }
    default:
        switch ("foo bar" + "") {
        case "foo bar":
            v13 := GenS1{F0: 0}
            _ = v13
        case "a":
            var v14 uint32 = 42
            _ = v14
        }
    }
    v2 = v2
    var v15 uint32 = 2147483647
    v15 += v15
    v16 := GenS1{F0: 1}
    var v17 uint8 = 128
    if (!true) {
        for fi18 := 0; fi18 < 16; fi18++ {
            v1[(v15 | v15)] = fi18
            if (0 != 65) {
                var v19 uint = 674
                _ = v19
            } else {
                v20 := []int{(^(fi18)), (fi18 << ((65535) & 63)), (fi18 + fi18)}
                for fi21 := 0; fi21 < 2; fi21++ {
                    v3 = (3.14159 * v3)
                    if true {
                        if len(v20) > 0 {
                            v20[0] = fi18
                        }
                        v22 := GenS2{F0: (^(v2)), F1: v2, F2: 9223372036854775807}
                        switch (("" + "hello") + (("\n\t" + "hello") + "foo bar")) {
                        case "foo bar":
                            v2 -= v2
                        default:
                            var v23 uint32 = 121
                            _ = v23
                        }
                        _ = v22
                    }
                    _ = fi21
                }
                _ = v20
            }
            var v24 int8 = 127
            _ = fi18
            _ = v24
        }
        v25 := []int{}
        _ = v25
    }
    var v26 uint32 = (^(v15))
    v27 := v4[42]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
