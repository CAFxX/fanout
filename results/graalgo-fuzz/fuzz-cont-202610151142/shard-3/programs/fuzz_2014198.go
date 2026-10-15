package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int64 }
func genFn0(p0 string, p1 int8) int64 {
    v6 := []int{1}
    if false {
        var v7 uint64 = 1
        _ = v7
    }
    _ = len(v6)
    _ = p0
    _ = p1
    _ = v6
    return -561
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]bool)
    var v3 bool = (true || (1000 <= -32768))
    var v4 int16 = 0
    var v5 uint32 = 831
    switch 2147483647 {
    case 127:
        v8 := []int{297, 42, 635, 777}
        _ = v8
    case 32767:
        v9 := GenS1{F0: 150, F1: "a"}
        _ = v9
    case 2147483647:
        var v10 int32 = 0
        _ = v10
    }
    var v11 int64 = -365
    v12 := GenS2{F0: 418}
    _ = len(v2)
    if ((v4 < (-(v4))) && v3) {
        v13 := []int{}
        _ = v13
    } else {
        var v14 bool = ((v11 + (v11 >> ((42) & 63))) == genFn0("foo bar", -128))
        _ = v14
    }
    if (65535 == 255) {
        fmt.Println(v5, v3)
        v5 = 654
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 float64 = float64((float32(1.5) * (3.14159 * -2.25)))
        switch "\n\t" {
        case "":
            if (!true) {
                v17 := GenS2{F0: genFn0("\n\t", 100)}
                for fi18 := 0; fi18 < 6; fi18++ {
                    var v19 float64 = (3.14159 - float64(-2.25))
                    _ = fi18
                    _ = v19
                }
                v4 -= (v4 - v4)
                _ = v17
            } else {
                if true {
                    if v3 {
                        var v20 int8 = -128
                        var v21 float32 = (3.14159 + (1.5 + -2.25))
                        _ = v20
                        _ = v21
                    }
                    var v22 float64 = v16
                    _ = v22
                } else {
                    if (((v16 / (0.0 - -2.25)) / 0.1) == (((v16 * v16) / ((0.1 + (float64(float32(1e10)) * 0.0)) - v16)) / (float64((0.0 * float32(v16))) / (0.1 / v16)))) {
                        var v23 uint64 = 907
                        _ = v23
                    }
                    v3 = false
                    var v24 bool = true
                    _ = v24
                }
                var v25 bool = false
                switch fi15 {
                case 42:
                    var v26 int8 = 1
                    _ = v26
                case 32768:
                    v25 = v3
                case 4294967295:
                    v3 = (291 < 18446744073709551615)
                }
                _ = v25
            }
        case "\n\t":
            v16 = v16
        case "foo bar":
            for fi27 := 0; fi27 < 5; fi27++ {
                var v28 int64 = int64(v5)
                _ = fi27
                _ = v28
            }
        }
        _ = fi15
        _ = v16
    }
    v3 = v3
    var v29 uint64 = uint64(v5)
    v11 += (v11 * (v11 | (-(v11))))
    v30 := v1[v5]
    switch (v30 | 32768) {
    case -948:
        switch (v30 >> ((1) & 63)) {
        case -1:
            v31 := GenS2{F0: (v11 & 42)}
            _ = v31
        case 127:
            var v32 uint = 255
            _ = v32
        default:
            v33 := []int{}
            _ = v33
        }
    default:
        fmt.Println(v29)
    }
    v34 := GenS1{F0: 0, F1: "hello"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v4
    _ = v5
}
