package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int8 }
type GenS2 struct { F0 int64; F1 int32; F2 string }
func genFn0(p0 uint8, p1 int32, p2 uint16) int32 {
    p0 = p0
    var v7 string = "hello"
    if (false || (v7 != (("a" + (v7 + v7)) + v7))) {
        var v8 uint32 = 4294967295
        if true {
            if true {
                v9 := GenS1{F0: p0, F1: -1, F2: 42}
                var v10 uint32 = 4294967295
                _ = v10
                _ = v9
            } else {
                for fi11 := 0; fi11 < 10; fi11++ {
                    switch ((v7 + "hello") + "a") {
                    case "\n\t":
                        for fi12 := 0; fi12 < 6; fi12++ {
                            v13 := []int{fi11}
                            _ = fi12
                            _ = v13
                        }
                    case "hello":
                        v14 := []int{-36, fi11, -343}
                        _ = v14
                    case "foo bar":
                        if true {
                            var v15 int32 = (-(p1))
                            v16 := []int{-2147483648}
                            _ = v15
                            _ = v16
                        } else {
                            fmt.Println(p2, p2, p2)
                            fmt.Println(fi11)
                            p2 -= (p2 - 1)
                        }
                    }
                    _ = fi11
                }
                v17 := GenS1{F0: 21, F1: 42, F2: -1}
                _ = v17
            }
            var v18 int16 = 0
            _ = v18
        } else {
            v19 := GenS1{F0: p0, F1: -32768, F2: 42}
            _ = v19
        }
        v8 += v8
        _ = v8
    } else {
        v20 := []int{}
        for fi21 := 0; fi21 < 17; fi21++ {
            p0 = (p0 >> ((65535) & 63))
            _ = fi21
        }
        v22 := GenS1{F0: 152, F1: -184, F2: 89}
        _ = v20
        _ = v22
    }
    v23 := GenS1{F0: p0, F1: -1, F2: 113}
    v24 := GenS1{F0: p0, F1: 42, F2: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    _ = v24
    _ = v7
    return int32(p2)
}
func genFn1(p0 string) int8 {
    if true {
        v25 := []int{}
        var v26 int8 = -128
        _ = v25
        _ = v26
    }
    if true {
        fmt.Println(p0)
    }
    _ = p0
    return -124
}

func main() {
    var v1 int32 = 0
    var v2 float64 = 3.14159
    var v3 uint16 = 42
    var v4 uint32 = 32
    v5 := make(map[uint32]int)
    v6 := []int{4294967295, 2147483647}
    fmt.Println(v3, v3, v3)
    for fi27 := 0; fi27 < 1; fi27++ {
        v2 += ((3.14159 - v2) / 1e10)
        _ = fi27
    }
    for fi28 := 0; fi28 < 5; fi28++ {
        for fi29 := 0; fi29 < 18; fi29++ {
            v1 = (v1 | 2147483647)
            v1 = v1
            _ = fi29
        }
        var v30 uint8 = 127
        v31 := GenS2{F0: 9223372036854775807, F1: (v1 / ((2147483647) | 1)), F2: ("a" + "a")}
        _ = fi28
        _ = v30
        _ = v31
    }
    var v32 uint16 = v3
    v1 += genFn0(2, 1, v3)
    v4 += v4
    if true {
        var v33 string = "a"
        _ = v33
    }
    var v34 string = "a"
    var v35 uint8 = 255
    v3 = v3
    var v36 float64 = float64((3.14159 + -2.25))
    for fi37 := 0; fi37 < 14; fi37++ {
        v36 += v36
        _ = fi37
    }
    delete(v5, v4)
    if true {
        fmt.Println(v2, v34)
    }
    fmt.Println(v3, v2, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
