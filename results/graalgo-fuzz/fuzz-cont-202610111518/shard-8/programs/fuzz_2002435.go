package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint }
func genFn0(p0 string, p1 uint32, p2 uint64) bool {
    if true {
        v6 := []int{-863, 127, 166, 9223372036854775807}
        _ = v6
    }
    for fi7 := 0; fi7 < 11; fi7++ {
        if (-32768 == 1000) {
            fmt.Println(fi7, p1)
            var v8 uint32 = 42
            _ = v8
        } else {
            switch (((("" + "hello") + (("" + "hello") + (("" + ("\n\t" + p0)) + p0))) + p0) + p0) {
            case "\n\t":
                v9 := []int{-267}
                _ = v9
            default:
                var v10 int32 = 1
                _ = v10
            }
            v11 := GenS1{F0: 1, F1: (fi7 ^ fi7), F2: 2147483647}
            _ = v11
        }
        if (255 > uint16(p1)) {
            p2 += p2
            var v12 string = "\n\t"
            _ = v12
        } else {
            v13 := GenS1{F0: 42, F1: fi7, F2: p1}
            _ = v13
        }
        v14 := GenS2{F0: 42, F1: 431}
        _ = fi7
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    return (100 > -13)
}
func genFn1(p0 bool, p1 string, p2 uint64) int16 {
    if p0 {
        if (!false) {
            var v15 string = (((p1 + "foo bar") + ("\n\t" + (((p1 + ((("\n\t" + (((("a" + (p1 + p1)) + (("foo bar" + p1) + "hello")) + p1) + p1)) + p1) + "")) + (p1 + p1)) + (("a" + p1) + p1)))) + (p1 + ""))
            p2 = 2
            v16 := []int{-803}
            _ = v15
            _ = v16
        }
        p1 = ""
        for fi17 := 0; fi17 < 6; fi17++ {
            fmt.Println(p1)
            _ = fi17
        }
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        v19 := GenS1{F0: 705, F1: (-(fi18)), F2: uint32(p2)}
        if p0 {
            var v20 int64 = 1
            _ = v20
        }
        _ = fi18
        _ = v19
    }
    if p0 {
        fmt.Println(p1)
        for fi21 := 0; fi21 < 19; fi21++ {
            v22 := GenS2{F0: 42, F1: 42}
            p2 += 1
            v23 := GenS2{F0: 1, F1: 18446744073709551615}
            _ = fi21
            _ = v22
            _ = v23
        }
    } else {
        p1 = p1
        for fi24 := 0; fi24 < 14; fi24++ {
            for fi25 := 0; fi25 < 7; fi25++ {
                var v26 int16 = 768
                p0 = (!p0)
                _ = fi25
                _ = v26
            }
            if (-128 == 100) {
                var v27 uint = 2
                for fi28 := 0; fi28 < 4; fi28++ {
                    var v29 float32 = 3.14159
                    fmt.Println(p2)
                    if true {
                        p0 = true
                    } else {
                        if p0 {
                            p0 = (255 != 1)
                        }
                        var v30 uint = v27
                        v31 := []int{-2147483648, -9223372036854775808}
                        _ = v30
                        _ = v31
                    }
                    _ = fi28
                    _ = v29
                }
                _ = v27
            }
            p2 = 556
            _ = fi24
        }
    }
    v32 := []int{-2147483648, 1, 1, -2147483648}
    for fi33 := 0; fi33 < 6; fi33++ {
        p1 = p1
        _ = fi33
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v32
    return 32767
}

func main() {
    var v1 int64 = 1
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]bool)
    v4 := make(map[int]int)
    var v5 uint64 = 1
    v34 := GenS1{F0: 1, F1: -2147483648, F2: 65535}
    v1 -= 42
    v35 := GenS2{F0: -1, F1: 1}
    v1 -= v1
    var v36 uint64 = v5
    delete(v3, 16)
    fmt.Println(v36, v1)
    switch "\n\t" {
    case "foo bar":
        v5 += 18446744073709551615
    case "\n\t":
        v5 = (^(v5))
    default:
        var v37 uint64 = 2
        _ = v37
    }
    v5 = v5
    v38 := GenS2{F0: 1000, F1: 2}
    fmt.Println(v5, v1, v5)
    var v39 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
