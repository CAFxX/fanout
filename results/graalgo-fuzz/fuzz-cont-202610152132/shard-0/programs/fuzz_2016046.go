package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 string }
type GenS2 struct { F0 uint }
func genFn0(p0 string) int64 {
    v6 := []int{-390, 1}
    for fi7 := 0; fi7 < 17; fi7++ {
        fmt.Println(p0, p0)
        fmt.Println(fi7, p0)
        _ = fi7
    }
    var v8 float32 = (3.14159 - 1.5)
    _ = p0
    _ = v6
    _ = v8
    return 1
}

func main() {
    var v1 uint8 = 82
    var v2 int = 2147483647
    v3 := make(map[int]int)
    v4 := []int{v2, v2, (v2 >> ((0) & 63))}
    var v5 int64 = -136
    v1 = v1
    switch (v2 >> ((65535) & 63)) {
    case 127:
        if ("" <= "a") {
            var v9 uint16 = 2
            if false {
                var v10 uint64 = 18446744073709551615
                if false {
                    for fi11 := 0; fi11 < 7; fi11++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        for fi12 := 0; fi12 < 7; fi12++ {
                            v5 = v5
                            _ = fi12
                        }
                        v10 += (^(v10))
                        _ = fi11
                    }
                    var v13 uint8 = v1
                    var v14 int64 = v5
                    _ = v13
                    _ = v14
                }
                _ = v10
            }
            _ = v9
        }
    default:
        for fi15 := 0; fi15 < 7; fi15++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v16 uint16 = 862
            fmt.Println(fi15, v5)
            _ = fi15
            _ = v16
        }
    }
    _ = len(v4)
    for fi17 := 0; fi17 < 6; fi17++ {
        v18 := GenS1{F0: 533, F1: 42, F2: "hello"}
        v19 := GenS2{F0: 754}
        _ = fi17
        _ = v18
        _ = v19
    }
    v20 := GenS1{F0: 0, F1: -2147483648, F2: ""}
    if (v1 >= (^(v1))) {
        var v21 int32 = -928
        _ = v21
    }
    for fi22 := 0; fi22 < 8; fi22++ {
        var v23 int8 = -1
        v23 = -1
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 2; fi24++ {
        delete(v3, -2147483648)
        var v25 int32 = -1
        _ = fi24
        _ = v25
    }
    var v26 uint32 = 65535
    switch ("foo bar" + "a") {
    case "foo bar":
        for fi27 := 0; fi27 < 16; fi27++ {
            if len(v4) > 0 {
                v4[0] = fi27
            }
            _ = len(v3)
            v5 += v5
            _ = fi27
        }
    }
    v4 = append(v4, v2)
    v5 = (^(v5))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
