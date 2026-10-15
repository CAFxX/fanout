package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint64 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64 }
func genFn0(p0 uint16) uint {
    var v6 int = -2147483648
    var v7 float32 = (1.5 + -2.25)
    var v8 uint32 = 152
    if true {
        var v9 uint = 76
        v10 := []int{v6, v6, (v6 + v6)}
        if len(v10) > 0 {
            v10[0] = 32768
        }
        _ = v10
        _ = v9
    } else {
        if (false && (127 == 127)) {
            v7 = (v7 - v7)
            for fi11 := 0; fi11 < 5; fi11++ {
                v12 := []int{}
                _ = fi11
                _ = v12
            }
            v13 := GenS1{F0: v8, F1: 32767, F2: 754}
            _ = v13
        }
    }
    switch "\n\t" {
    case "\n\t":
        p0 -= (p0 - p0)
    case "hello":
        var v14 float64 = 0.1
        _ = v14
    case "a":
        switch 9223372036854775807 {
        case -9223372036854775808:
            p0 += p0
        default:
            for fi15 := 0; fi15 < 3; fi15++ {
                v16 := []int{(v6 | fi15), v6, 1, int(v8)}
                _ = fi15
                _ = v16
            }
        }
    default:
        switch "" {
        case "hello":
            var v17 uint32 = (v8 >> ((255) & 63))
            _ = v17
        }
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint32, p1 uint64) int {
    if true {
        p0 = p0
        var v18 int8 = -1
        var v19 int64 = -141
        _ = v18
        _ = v19
    } else {
        var v20 float32 = -2.25
        _ = v20
    }
    v21 := GenS2{F0: 0, F1: 87}
    switch "\n\t" {
    case "\n\t":
        p1 += (p1 % ((p1) | 1))
    }
    fmt.Println(p1)
    p1 = (p1 >> ((2) & 63))
    _ = p0
    _ = p1
    _ = v21
    return 2147483647
}

func main() {
    v1 := make(map[int]int)
    v2 := []string{"hello"}
    var v3 string = "hello"
    v4 := make(map[uint32]int)
    var v5 uint16 = 0
    for fi22 := 0; fi22 < 15; fi22++ {
        fmt.Println(v5, v3)
        v3 = "foo bar"
        _ = fi22
    }
    fmt.Println(v5, v3)
    fmt.Println(v3, v5)
    var v23 uint8 = 127
    var v24 float32 = (0.1 * 3.14159)
    var v25 int8 = 42
    v24 -= v24
    switch v3 {
    case "foo bar":
        for fi26 := 0; fi26 < 10; fi26++ {
            v5 = v5
            fmt.Println(v5, v3, v23)
            _ = fi26
        }
    case "a":
        v27 := GenS3{F0: 0, F1: 42}
        _ = v27
    default:
        switch -1 {
        case 9223372036854775807:
            if true {
                for fi28 := 0; fi28 < 20; fi28++ {
                    if false {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v29 uint64 = 42
                        v30 := GenS3{F0: 9223372036854775807, F1: -459}
                        _ = v29
                        _ = v30
                    }
                    _ = fi28
                }
                delete(v1, 42)
            } else {
                for fi31 := 0; fi31 < 15; fi31++ {
                    var v32 uint8 = v23
                    fmt.Println(v5, v25, v24)
                    switch (-(fi31)) {
                    case 9223372036854775807:
                        var v33 uint16 = v5
                        _ = v33
                    case 128:
                        fmt.Println(v3)
                    default:
                        fmt.Println(v32)
                    }
                    _ = fi31
                    _ = v32
                }
                if true {
                    v34 := GenS3{F0: 9223372036854775807, F1: 1}
                    v35 := GenS3{F0: 42, F1: -1}
                    _ = v34
                    _ = v35
                } else {
                    if false {
                        _ = len(v4)
                    }
                    v36 := GenS4{F0: 32768}
                    _ = v36
                }
            }
        case 480:
            for fi37 := 0; fi37 < 1; fi37++ {
                var v38 uint32 = 83
                if (-2147483648 >= 0) {
                    switch v3 {
                    case "foo bar":
                        v38 += (^(v38))
                    case "\n\t":
                        var v39 int64 = -399
                        _ = v39
                    }
                    var v40 string = v3
                    fmt.Println(v5, v24, fi37)
                    _ = v40
                }
                v41 := GenS5{F0: -695}
                _ = fi37
                _ = v38
                _ = v41
            }
        case 478:
            var v42 bool = (!false)
            _ = v42
        default:
            fmt.Println(v25, v23)
        }
    }
    var v43 float32 = (v24 * v24)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v43
    _ = v5
}
