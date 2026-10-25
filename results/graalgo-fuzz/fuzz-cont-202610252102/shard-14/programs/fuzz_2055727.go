package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string }
type GenS5 struct { F0 string; F1 bool; F2 int8 }
type GenS6 struct { F0 int32 }
func genFn0(p0 uint, p1 uint) int {
    if false {
        v4 := []int{1}
        switch "" {
        case "foo bar":
            var v5 int = -327
            _ = v5
        }
        v6 := GenS1{F0: 128, F1: 1}
        _ = v4
        _ = v6
    } else {
        if false {
            if true {
                var v7 uint32 = 668
                var v8 int64 = 9223372036854775807
                v9 := GenS1{F0: 127, F1: 1}
                _ = v7
                _ = v8
                _ = v9
            } else {
                for fi10 := 0; fi10 < 15; fi10++ {
                    for fi11 := 0; fi11 < 7; fi11++ {
                        if (!false) {
                            v12 := GenS1{F0: 127, F1: 2147483647}
                            _ = v12
                        } else {
                            v13 := []int{319, fi11, fi10, 1}
                            _ = v13
                        }
                        v14 := GenS1{F0: 2, F1: 2147483647}
                        fmt.Println(p0, fi10)
                        _ = fi11
                        _ = v14
                    }
                    _ = fi10
                }
                fmt.Println(p0, p0)
                for fi15 := 0; fi15 < 9; fi15++ {
                    var v16 bool = (false || (false && true))
                    if v16 {
                        v16 = (42 <= 2147483648)
                        v17 := []int{fi15, (fi15 ^ 823), fi15, 32768}
                        v18 := GenS2{F0: 1, F1: 2}
                        _ = v17
                        _ = v18
                    }
                    for fi19 := 0; fi19 < 9; fi19++ {
                        v20 := GenS2{F0: 42, F1: 42}
                        v21 := []int{(fi19 * -2147483648)}
                        var v22 uint32 = 2147483647
                        var v23 float64 = 1.5
                        _ = fi19
                        _ = v20
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    _ = fi15
                    _ = v16
                }
            }
            v24 := GenS3{F0: 65535}
            for fi25 := 0; fi25 < 6; fi25++ {
                for fi26 := 0; fi26 < 18; fi26++ {
                    v27 := GenS2{F0: 2, F1: 1}
                    _ = fi26
                    _ = v27
                }
                _ = fi25
            }
            _ = v24
        }
    }
    for fi28 := 0; fi28 < 19; fi28++ {
        if (!false) {
            var v29 int16 = -1
            p0 += 126
            _ = v29
        }
        switch "hello" {
        case "\n\t":
            v30 := GenS1{F0: 0, F1: 2147483647}
            _ = v30
        default:
            v31 := []int{}
            _ = v31
        }
        _ = fi28
    }
    p0 += 381
    if (0 != -32768) {
        fmt.Println(p1, p1)
        for fi32 := 0; fi32 < 13; fi32++ {
            fmt.Println(p0, fi32)
            var v33 uint8 = 33
            v34 := []int{127}
            _ = fi32
            _ = v33
            _ = v34
        }
        fmt.Println(p0, p0)
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    return 32768
}

func main() {
    var v1 int = 1
    var v2 string = "foo bar"
    var v3 uint8 = 165
    v35 := GenS4{F0: (v2 + v2)}
    fmt.Println(v2, v1, v1)
    if (false || false) {
        var v36 uint64 = 300
        var v37 uint = 255
        v38 := GenS5{F0: v2, F1: (true && false), F2: -125}
        _ = v36
        _ = v37
        _ = v38
    }
    var v39 int32 = 0
    var v40 uint64 = 18446744073709551615
    v2 = ((("hello" + ((v2 + "foo bar") + (v2 + v2))) + v2) + v2)
    if (248 > v3) {
        var v41 int16 = 42
        v40 -= (v40 / ((0) | 1))
        _ = v41
    }
    v1 -= v1
    for fi42 := 0; fi42 < 8; fi42++ {
        v43 := GenS6{F0: (v39 ^ v39)}
        fmt.Println(v2)
        v44 := []int{genFn0(4294967295, 932), -2147483648, (-(fi42))}
        _ = fi42
        _ = v43
        _ = v44
    }
    var v45 uint32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v39
    _ = v40
    _ = v45
}
