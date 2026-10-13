package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8 }
func genFn0(p0 int8, p1 string) int16 {
    p1 = p1
    var v6 string = "a"
    for fi7 := 0; fi7 < 1; fi7++ {
        switch p1 {
        case "hello":
            var v8 int32 = -1
            _ = v8
        }
        switch "hello" {
        case "a":
            p0 = (p0 ^ (p0 - (p0 << ((65535) & 63))))
        case "":
            if (293 == 2) {
                for fi9 := 0; fi9 < 15; fi9++ {
                    if (0.0 >= 1.5) {
                        var v10 string = p1
                        var v11 float32 = (0.1 / 3.14159)
                        v12 := []int{-488}
                        _ = v10
                        _ = v11
                        _ = v12
                    }
                    _ = fi9
                }
                var v13 uint32 = 1
                switch (p1 + ("foo bar" + "hello")) {
                case "\n\t":
                    v14 := []int{}
                    _ = v14
                case "":
                    for fi15 := 0; fi15 < 19; fi15++ {
                        p0 = (p0 % (((p0 | p0)) | 1))
                        v16 := GenS1{F0: 2147483648}
                        fmt.Println(v6)
                        if (false && (!(!false))) {
                            fmt.Println(p1, v6)
                        }
                        _ = fi15
                        _ = v16
                    }
                case "foo bar":
                    for fi17 := 0; fi17 < 16; fi17++ {
                        fmt.Println(fi17)
                        p0 += p0
                        var v18 float32 = (float32((float64((-2.25 * (3.14159 / ((3.14159 * (1.5 * 0.1)) + ((float32(1.5) - 3.14159) * 1e10))))) + ((1.5 + (1e10 + (0.1 / 1.5))) + (0.0 + 0.1)))) / (((3.14159 * -2.25) + 1.5) * ((1e10 / float32(0.1)) - 1.5)))
                        _ = fi17
                        _ = v18
                    }
                default:
                    v19 := []int{-9223372036854775808}
                    _ = v19
                }
                _ = v13
            }
        }
        var v20 uint16 = 1
        _ = fi7
        _ = v20
    }
    switch 127 {
    case 127:
        var v21 uint = 1
        _ = v21
    case -2147483648:
        fmt.Println(p1, p0, p0)
    }
    _ = p0
    _ = p1
    _ = v6
    return 42
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[uint32]int)
    var v3 int8 = -34
    var v4 int64 = v1
    var v5 int16 = 1
    v4 -= int64(v5)
    v2[4294967295] = -2147483648
    for fi22 := 0; fi22 < 2; fi22++ {
        var v23 uint = 42
        _ = fi22
        _ = v23
    }
    var v24 int = -332
    v1 += (v4 << ((392) & 63))
    v2[2] = (v24 ^ 329)
    var v25 int32 = 25
    v26 := GenS1{F0: 0}
    var v27 uint16 = 42
    v28 := GenS1{F0: 42}
    v29 := GenS2{F0: 120}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
