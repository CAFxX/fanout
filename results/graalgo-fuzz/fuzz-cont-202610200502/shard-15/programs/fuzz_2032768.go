package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 int64; F2 int }
type GenS4 struct { F0 uint; F1 int32; F2 uint8 }
type GenS5 struct { F0 int16; F1 int16 }
func genFn0(p0 uint8, p1 int64) int8 {
    v5 := GenS1{F0: 2}
    var v6 uint = 1
    v6 += (v6 << ((20) & 63))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -115
}
func genFn1() uint8 {
    var v7 int32 = 2147483647
    switch 127 {
    case 198:
        v7 = (-(v7))
    case 1:
        v8 := GenS1{F0: 255}
        _ = v8
    case 4294967295:
        var v9 int8 = 0
        _ = v9
    }
    _ = v7
    return 82
}

func main() {
    var v1 float32 = 1e10
    v2 := make(map[uint32]int)
    var v3 int = 4294967295
    v4 := []int{42}
    var v10 string = ("" + "")
    var v11 uint64 = 42
    for fi12 := 0; fi12 < 7; fi12++ {
        var v13 uint16 = 583
        _ = fi12
        _ = v13
    }
    switch (v3 * v3) {
    case 127:
        v14 := GenS2{F0: (v3 + 725)}
        _ = v14
    default:
        switch ((v10 + (v10 + v10)) + v10) {
        case "foo bar":
            v4 = append(v4, v3)
        case "a":
            v15 := GenS2{F0: v3}
            _ = v15
        case "\n\t":
            v1 -= -2.25
        default:
            var v16 int = v3
            _ = v16
        }
    }
    v17 := GenS3{F0: 786, F1: 1, F2: (v3 - 1)}
    if len(v4) > 0 {
        v4[0] = v3
    }
    fmt.Println(v10)
    v18 := GenS4{F0: 42, F1: 42, F2: 127}
    if false {
        var v19 float32 = 3.14159
        for fi20 := 0; fi20 < 9; fi20++ {
            fmt.Println(v11, fi20)
            _ = fi20
        }
        _ = v19
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        var v22 bool = false
        if v22 {
            fmt.Println(v3, v10)
        }
        fmt.Println(v1)
        _ = fi21
        _ = v22
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if false {
        v3 = -1
        for fi23 := 0; fi23 < 2; fi23++ {
            delete(v2, 65535)
            fmt.Println(v11, v1, v1)
            v24 := GenS5{F0: -896, F1: 0}
            _ = fi23
            _ = v24
        }
        var v25 bool = true
        _ = v25
    } else {
        switch "" {
        case "a":
            var v26 int16 = 32767
            _ = v26
        case "foo bar":
            v11 = (v11 / ((200) | 1))
        default:
            v3 -= v3
        }
        if (0 <= -32768) {
            v27 := GenS5{F0: 1000, F1: 42}
            _ = v27
        } else {
            if false {
                fmt.Println(v10)
            }
            if (v10 == "\n\t") {
                for fi28 := 0; fi28 < 15; fi28++ {
                    for fi29 := 0; fi29 < 2; fi29++ {
                        fmt.Println(v10, fi28)
                        fmt.Println(v3)
                        if false {
                            var v30 string = "\n\t"
                            var v31 uint16 = 255
                            _ = v30
                            _ = v31
                        }
                        var v32 uint64 = (v11 & 513)
                        _ = fi29
                        _ = v32
                    }
                    var v33 uint32 = 2147483648
                    _ = fi28
                    _ = v33
                }
                _ = len(v2)
                v1 += (v1 + float32((0.1 - 0.1)))
            }
        }
        if (592 >= 65535) {
            var v34 string = "hello"
            var v35 uint32 = 779
            if ((0.0 == 3.14159) && false) {
                v34 = v34
            } else {
                if (v11 == (v11 | 535)) {
                    v36 := GenS4{F0: 4294967295, F1: -1, F2: 1}
                    v4 = append(v4, (v3 & (v3 - (v3 + (-(v3))))))
                    _ = v36
                } else {
                    v2[2] = -2147483648
                    for fi37 := 0; fi37 < 13; fi37++ {
                        delete(v2, (v35 / ((v35) | 1)))
                        var v38 uint32 = (^(v35))
                        _ = fi37
                        _ = v38
                    }
                    var v39 float64 = 3.14159
                    _ = v39
                }
            }
            _ = v34
            _ = v35
        }
    }
    _ = len(v4)
    for fi40 := 0; fi40 < 5; fi40++ {
        v10 = v10
        var v41 int32 = 42
        _ = fi40
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
