package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32 }
func genFn0() uint64 {
    var v6 string = "\n\t"
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            v9 := []int{(fi8 * (fi8 - (fi8 / ((-2147483648) | 1)))), fi7}
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    var v10 float64 = (-2.25 * 1.5)
    fmt.Println(v10, v10)
    if (!false) {
        switch ("" + "") {
        case "hello":
            v11 := []int{}
            _ = v11
        case "foo bar":
            v12 := GenS1{F0: 9223372036854775807, F1: 703, F2: 1}
            _ = v12
        case "a":
            fmt.Println(v6)
        }
        v13 := GenS2{F0: 2}
        _ = v13
    } else {
        switch -713 {
        case -9223372036854775808:
            v6 = ((v6 + "\n\t") + ((("foo bar" + ("" + v6)) + v6) + v6))
        case -2147483648:
            for fi14 := 0; fi14 < 11; fi14++ {
                v6 = v6
                _ = fi14
            }
        }
        fmt.Println(v6)
    }
    _ = v10
    _ = v6
    return 2
}
func genFn1(p0 int8) int16 {
    p0 -= p0
    v15 := GenS1{F0: 834, F1: -476, F2: 32767}
    for fi16 := 0; fi16 < 18; fi16++ {
        var v17 uint64 = 496
        v17 += 829
        var v18 int16 = 42
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 int = 127
    _ = p0
    _ = v15
    _ = v19
    return 0
}

func main() {
    var v1 float64 = 1.5
    var v2 uint = 2
    var v3 float32 = (float32(v1) * (3.14159 + (-2.25 * 0.0)))
    var v4 string = "\n\t"
    v5 := make(map[string]int)
    var v20 float32 = 1.5
    if (!(!false)) {
        var v21 string = v4
        _ = v21
    } else {
        v22 := []int{127, 921}
        _ = v22
    }
    v23 := GenS1{F0: -573, F1: 42, F2: 255}
    fmt.Println(v2, v20)
    if ((!true) && true) {
        if false {
            var v24 uint64 = 198
            delete(v5, "\n\t")
            if (!false) {
                var v25 uint64 = v24
                v26 := []int{32768, 0, 9223372036854775807}
                fmt.Println(v25)
                _ = v25
                _ = v26
            }
            _ = v24
        }
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        v28 := GenS1{F0: 9223372036854775807, F1: 595, F2: 2}
        v29 := GenS1{F0: -256, F1: 9223372036854775807, F2: 42}
        _ = fi27
        _ = v28
        _ = v29
    }
    if false {
        v2 = 42
        v2 = v2
        v3 = float32((v1 - v1))
    } else {
        var v30 int8 = 50
        fmt.Println(v1)
        _ = v30
    }
    var v31 int = -226
    v32 := GenS3{F0: 4294967295}
    v33 := GenS2{F0: 829}
    if ((v2 << ((580) & 63)) == v2) {
        var v34 float64 = float64((v3 - v20))
        var v35 string = (v4 + (v4 + v4))
        _ = v34
        _ = v35
    }
    var v36 float64 = v1
    v3 += ((v20 - v20) / float32((v1 - v1)))
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v36
    _ = v4
    _ = v5
}
