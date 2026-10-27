package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 int8 }
func genFn0() int16 {
    v6 := []int{-9223372036854775808, 802, 32768}
    var v7 int64 = 42
    v6 = append(v6, 9223372036854775807)
    for fi8 := 0; fi8 < 9; fi8++ {
        for fi9 := 0; fi9 < 19; fi9++ {
            if (true && false) {
                var v10 uint16 = 2
                _ = v10
            }
            var v11 float64 = (0.1 * 3.14159)
            _ = fi9
            _ = v11
        }
        _ = fi8
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        for fi13 := 0; fi13 < 15; fi13++ {
            var v14 string = "\n\t"
            var v15 int8 = 101
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi12
    }
    _ = v6
    _ = v7
    return 483
}
func genFn1() uint64 {
    v16 := []int{}
    switch 127 {
    case 169:
        if (false && (1 == 32768)) {
            var v17 string = ""
            _ = v17
        } else {
            if len(v16) > 0 {
                v16[0] = 447
            }
        }
    case 128:
        var v18 uint64 = 0
        _ = v18
    case 9223372036854775807:
        if (((32767 >= 1) || false) || (0 <= 657)) {
            if (4294967295 > 18446744073709551615) {
                _ = len(v16)
            }
            var v19 bool = ((83 >= 2) && false)
            _ = v19
        } else {
            if true {
                var v20 string = (("hello" + "foo bar") + "")
                _ = v20
            } else {
                var v21 float64 = 0.0
                switch "foo bar" {
                case "":
                    v22 := GenS1{F0: 211, F1: 128}
                    _ = v22
                case "foo bar":
                    v23 := GenS2{F0: 23, F1: 32767}
                    _ = v23
                case "hello":
                    fmt.Println(v21)
                }
                var v24 float32 = (((1e10 - 0.0) + (3.14159 / 1.0)) * (1.5 * 1.5))
                _ = v21
                _ = v24
            }
            var v25 uint8 = 42
            v25 += 38
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 8; fi26++ {
        if (false && ((0 > 639) || (42 != 1))) {
            var v27 uint16 = 168
            _ = v27
        } else {
            var v28 float32 = 0.0
            var v29 int8 = 127
            for fi30 := 0; fi30 < 9; fi30++ {
                v28 = float32((1e10 - 1.5))
                fmt.Println(v29, fi30)
                var v31 uint32 = 0
                _ = fi30
                _ = v31
            }
            _ = v28
            _ = v29
        }
        v32 := GenS3{F0: 2}
        var v33 bool = true
        _ = fi26
        _ = v32
        _ = v33
    }
    v16 = append(v16, 844)
    _ = v16
    return 18446744073709551615
}

func main() {
    var v1 bool = (0 >= 4294967295)
    var v2 int16 = 970
    var v3 int = 4294967295
    var v4 uint32 = 82
    v5 := make(map[int]int)
    v34 := GenS3{F0: (v4 + (v4 << ((1) & 63)))}
    var v35 float32 = 0.0
    v1 = true
    var v36 uint32 = v4
    v37 := []int{}
    if v1 {
        if v1 {
            v38 := v5[v3]
            var v39 int16 = v2
            _ = v38
            _ = v39
        } else {
            v35 = 1e10
        }
    }
    var v40 int = v3
    switch "a" {
    case "":
        v3 -= (-(v40))
    case "\n\t":
        var v41 int = v40
        _ = v41
    }
    v35 += v35
    for fi42 := 0; fi42 < 6; fi42++ {
        switch v40 {
        case -9223372036854775808:
            v37 = append(v37, (^(v3)))
        case 0:
            v43 := GenS3{F0: v4}
            _ = v43
        }
        _ = fi42
    }
    v44 := GenS1{F0: 1, F1: 128}
    v45 := GenS3{F0: 2}
    v5[v3] = 190
    if v1 {
        for fi46 := 0; fi46 < 15; fi46++ {
            var v47 bool = false
            _ = fi46
            _ = v47
        }
    }
    if (v1 && true) {
        v36 -= (^(v4))
    }
    if v1 {
        var v48 int32 = 2147483647
        v49 := GenS4{F0: 46, F1: 42}
        _ = len(v5)
        _ = v48
        _ = v49
    }
    v40 = (v3 >> ((0) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v40
    _ = v44
    _ = v45
    _ = v5
}
