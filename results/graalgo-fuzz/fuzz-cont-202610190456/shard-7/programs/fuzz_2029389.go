package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int8 }
type GenS2 struct { F0 int64; F1 uint32 }
func genFn0(p0 uint8, p1 int64) string {
    var v7 float32 = float32((-2.25 + 0.0))
    if (!((!true) && ((((float32(float64((v7 * float32(3.14159)))) <= v7) && true) || false) || (false || true)))) {
        var v8 uint32 = 0
        var v9 float32 = (((v7 / v7) - ((v7 - v7) * (v7 / float32(3.14159)))) / 3.14159)
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 2; fi10++ {
            for fi11 := 0; fi11 < 7; fi11++ {
                var v12 int8 = 100
                _ = fi11
                _ = v12
            }
            var v13 uint = 1
            for fi14 := 0; fi14 < 10; fi14++ {
                v15 := []int{fi14, (fi14 >> ((v13) & 63)), fi10}
                v13 -= v13
                _ = fi14
                _ = v15
            }
            _ = fi10
            _ = v13
        }
        v16 := []int{0, 657, 128}
        _ = v16
    }
    v17 := []int{740, 1, -1, 42}
    var v18 int32 = 783
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    _ = v7
    return ""
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = (488 < 42)
    v3 := make(map[int]string)
    var v4 uint64 = 42
    v5 := []int{953, -1, -9223372036854775808, -490}
    v6 := make(map[uint32]bool)
    var v19 int8 = -60
    if (!((float32(-2.25) + -2.25) != 3.14159)) {
        v20 := GenS1{F0: 0, F1: 426, F2: v19}
        for fi21 := 0; fi21 < 3; fi21++ {
            v19 = v19
            _ = fi21
        }
        fmt.Println(v2, v19, v2)
        _ = v20
    } else {
        v2 = v2
    }
    fmt.Println(v19, v2)
    v19 -= (v19 - v19)
    v2 = v2
    switch ("" + "\n\t") {
    case "":
        if v2 {
            for fi22 := 0; fi22 < 2; fi22++ {
                var v23 int32 = -2147483648
                _ = fi22
                _ = v23
            }
        } else {
            var v24 float64 = 0.0
            _ = v24
        }
    case "hello":
        v25 := GenS2{F0: 890, F1: 600}
        _ = v25
    default:
        v2 = true
    }
    var v26 int32 = -214
    var v27 int32 = v26
    v19 = v19
    for fi28 := 0; fi28 < 20; fi28++ {
        var v29 int = (-(fi28))
        if (true && (0.0 == (1e10 / 3.14159))) {
            var v30 float64 = float64((((0.1 - ((1e10 + 3.14159) / 1e10)) / 1.0) / 3.14159))
            switch (^(fi28)) {
            case 570:
                fmt.Println(v29)
            }
            _ = v30
        }
        _ = fi28
        _ = v29
    }
    v31 := GenS1{F0: 1, F1: 68, F2: 100}
    if (1 <= 127) {
        _ = len(v1)
        for fi32 := 0; fi32 < 11; fi32++ {
            fmt.Println(v4, v4, v19)
            _ = fi32
        }
    }
    fmt.Println(v26, v2, v26)
    var v33 string = "\n\t"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
