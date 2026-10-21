package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 int64; F2 bool }
type GenS4 struct { F0 uint32; F1 uint8; F2 uint8 }
func genFn0(p0 uint16, p1 uint16) int16 {
    if ((1.5 - 0.1) <= float64(0.0)) {
        for fi7 := 0; fi7 < 17; fi7++ {
            if false {
                switch "" {
                case "foo bar":
                    fmt.Println(fi7, p1)
                case "\n\t":
                    p0 = (p0 >> ((0) & 63))
                default:
                    p0 += p0
                }
                v8 := []int{0, fi7, -938, -2147483648}
                _ = v8
            }
            _ = fi7
        }
        p1 += 139
    }
    var v9 uint32 = 2
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := []int{569}
        for fi12 := 0; fi12 < 3; fi12++ {
            v13 := GenS1{F0: (fi10 | fi10), F1: 32767}
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v9
    return 549
}
func genFn1() int8 {
    var v14 int64 = -9223372036854775808
    v15 := GenS1{F0: 42, F1: 255}
    _ = v14
    _ = v15
    return 47
}

func main() {
    var v1 string = "\n\t"
    var v2 uint16 = 2
    var v3 int64 = 625
    var v4 int16 = -532
    var v5 uint32 = 2147483648
    v6 := make(map[int]uint32)
    var v16 uint8 = 255
    if (128 == (v16 & (v16 % (((v16 - (^(v16)))) | 1)))) {
        v2 += 65535
    }
    if true {
        v2 -= v2
        var v17 int64 = (^(v3))
        _ = v17
    }
    var v18 float64 = ((0.1 - ((1e10 + (3.14159 / (1e10 - float64((0.0 - (1.5 * 1e10)))))) / ((3.14159 / ((0.1 + float64(float32(-2.25))) - -2.25)) - float64(3.14159)))) / 1e10)
    if (v16 >= (v16 + (v16 * v16))) {
        for fi19 := 0; fi19 < 14; fi19++ {
            var v20 int16 = (v4 & (-(v4)))
            v21 := GenS1{F0: (fi19 << ((42) & 63)), F1: v2}
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    switch -476 {
    case 2147483647:
        switch "hello" {
        case "foo bar":
            if true {
                var v22 uint32 = (v5 >> ((582) & 63))
                var v23 int8 = 1
                _ = v22
                _ = v23
            } else {
                v24 := GenS2{F0: -9223372036854775808}
                _ = v24
            }
        case "hello":
            fmt.Println(v2, v16, v16)
        }
    case 0:
        if (-2147483648 >= -427) {
            var v25 float64 = ((float64(1e10) + ((v18 / 3.14159) * (v18 / 1.0))) + 0.0)
            _ = v25
        }
    case 831:
        v3 = v3
    }
    var v26 float64 = (v18 * 0.1)
    v27 := GenS3{F0: 477, F1: (^(v3)), F2: ((-1 == v3) || (42 >= 385))}
    for fi28 := 0; fi28 < 2; fi28++ {
        if (((^(v2)) < v2) || (v18 >= (1.5 - ((v26 * v18) * v26)))) {
            v3 += v3
            fmt.Println(v4, v2)
            fmt.Println(v3, v26)
        }
        switch (fi28 << ((81) & 63)) {
        case -273:
            fmt.Println(fi28)
        case 9223372036854775807:
            v4 = v4
        default:
            v29 := []int{(fi28 % (((fi28 / (((fi28 - int(v4))) | 1))) | 1))}
            _ = v29
        }
        _ = fi28
    }
    for fi30 := 0; fi30 < 14; fi30++ {
        for fi31 := 0; fi31 < 16; fi31++ {
            v32 := GenS4{F0: 1, F1: 1, F2: v16}
            switch v1 {
            case "foo bar":
                var v33 uint16 = (^(v2))
                _ = v33
            }
            _ = fi31
            _ = v32
        }
        var v34 float32 = float32(v18)
        _ = fi30
        _ = v34
    }
    fmt.Println(v3)
    v3 = v3
    var v35 uint64 = 0
    v36 := []int{1, -2147483648, 32767}
    v16 = (v16 >> ((437) & 63))
    var v37 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
