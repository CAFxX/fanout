package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int }
func genFn0(p0 int64) string {
    v7 := []int{42, 711}
    p0 -= 749
    v8 := GenS1{F0: -944}
    _ = p0
    _ = v7
    _ = v8
    return "a"
}
func genFn1(p0 uint8) bool {
    var v9 float32 = 1.5
    fmt.Println(v9)
    _ = p0
    _ = v9
    return (-194 > 685)
}

func main() {
    v1 := []uint32{42, 2147483648, 876}
    var v2 uint64 = 630
    var v3 int32 = 2147483647
    v4 := []uint32{}
    v5 := []int{1, 32768, 9223372036854775807, -9223372036854775808}
    v6 := make(map[string]uint32)
    switch "a" {
    case "a":
        v6["hello"] = 42
    case "foo bar":
        var v10 int8 = -1
        _ = v10
    default:
        var v11 int8 = 0
        _ = v11
    }
    v12 := v6[""]
    v13 := GenS1{F0: -1}
    if true {
        var v14 int = -473
        for fi15 := 0; fi15 < 5; fi15++ {
            if true {
                v12 = 2147483648
            } else {
                v3 -= (v3 / ((2147483647) | 1))
            }
            var v16 float64 = (1e10 + 3.14159)
            switch -156 {
            case 638:
                fmt.Println(v14, v3, fi15)
            }
            _ = fi15
            _ = v16
        }
        _ = v14
    } else {
        var v17 uint32 = (^(v12))
        for fi18 := 0; fi18 < 6; fi18++ {
            v12 -= (v12 / ((v17) | 1))
            var v19 float32 = (1e10 / -2.25)
            _ = fi18
            _ = v19
        }
        v20 := GenS1{F0: 42}
        _ = v17
        _ = v20
    }
    v21 := GenS1{F0: 1}
    v22 := GenS1{F0: 322}
    v23 := GenS2{F0: false}
    if false {
        v3 = (v3 - (v3 | 1))
        if true {
            fmt.Println(v3, v3)
            _ = len(v4)
            v24 := GenS1{F0: 1}
            _ = v24
        }
        var v25 float32 = (3.14159 * (float32(3.14159) + 0.0))
        _ = v25
    } else {
        var v26 int64 = 42
        for fi27 := 0; fi27 < 8; fi27++ {
            var v28 uint8 = 42
            _ = fi27
            _ = v28
        }
        v3 += v3
        _ = v26
    }
    var v29 float64 = float64(float32(float64(1e10)))
    v29 = 1.5
    v3 += (v3 | 0)
    var v30 float64 = (v29 + -2.25)
    switch "foo bar" {
    case "foo bar":
        if (((v29 - v30) / float64(((3.14159 + 3.14159) + 0.1))) == ((((((0.1 - 1.5) / v29) + v30) + float64(0.1)) - v30) + (v29 * v30))) {
            v3 = (v3 + v3)
            v31 := GenS2{F0: false}
            v3 = v3
            _ = v31
        }
    }
    v32 := GenS3{F0: 0}
    if (32767 == 0) {
        var v33 float32 = 1.5
        var v34 uint8 = 195
        switch -2147483648 {
        case 911:
            v35 := GenS1{F0: -32768}
            _ = v35
        }
        _ = v33
        _ = v34
    } else {
        var v36 uint64 = v2
        _ = v36
    }
    v29 += v29
    v37 := GenS2{F0: false}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
