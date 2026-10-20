package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 string }
func genFn0(p0 uint16, p1 uint) string {
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := []int{}
        var v8 float32 = float32(1.5)
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        p0 -= p0
        v10 := GenS1{F0: 65535, F1: false, F2: ("\n\t" + "\n\t")}
        _ = fi9
        _ = v10
    }
    var v11 int64 = 0
    _ = p0
    _ = p1
    _ = v11
    return "foo bar"
}

func main() {
    var v1 uint32 = 4294967295
    var v2 int = -1
    v3 := make(map[int]string)
    var v4 int8 = 0
    var v5 float64 = (float64((1.5 + 0.1)) / 0.1)
    if false {
        if (0 <= 42) {
            var v12 int64 = 374
            v5 = float64((1.5 - 1.5))
            _ = v12
        }
        v13 := GenS1{F0: 65535, F1: true, F2: ""}
        _ = v13
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        v4 = (v4 + v4)
        v1 = v1
        _ = fi14
    }
    if (true && (!(!false))) {
        fmt.Println(v5, v1, v1)
        var v15 uint = 42
        var v16 int = (v2 ^ -141)
        _ = v15
        _ = v16
    } else {
        for fi17 := 0; fi17 < 11; fi17++ {
            var v18 int8 = v4
            v2 += (v2 | fi17)
            _ = fi17
            _ = v18
        }
        switch (v2 * (^(v2))) {
        case -9223372036854775808:
            var v19 float64 = 3.14159
            _ = v19
        default:
            var v20 uint64 = uint64(v1)
            _ = v20
        }
        v21 := GenS1{F0: 2, F1: ((v4 & -1) < 127), F2: "\n\t"}
        _ = v21
    }
    var v22 float32 = 3.14159
    fmt.Println(v2, v22)
    var v23 float32 = (0.0 - 1.5)
    var v24 string = "foo bar"
    for fi25 := 0; fi25 < 15; fi25++ {
        v24 = "foo bar"
        _ = fi25
    }
    var v26 float32 = -2.25
    var v27 int8 = 74
    for fi28 := 0; fi28 < 12; fi28++ {
        switch v24 {
        case "a":
            v29 := GenS1{F0: 327, F1: (int8(v1) >= v4), F2: (((v24 + v24) + v24) + "")}
            _ = v29
        default:
            if false {
                v1 = (v1 + v1)
            } else {
                if true {
                    for fi30 := 0; fi30 < 6; fi30++ {
                        v1 += 2147483647
                        v26 -= float32(float64(0.1))
                        _ = fi30
                    }
                }
                v5 = (v5 - (v5 - 1.5))
                var v31 int16 = -32768
                _ = v31
            }
        }
        v32 := []int{127, (fi28 / ((v2) | 1)), (v2 & 1), (v2 ^ (fi28 % ((fi28) | 1)))}
        _ = fi28
        _ = v32
    }
    for fi33 := 0; fi33 < 16; fi33++ {
        var v34 int16 = 42
        v24 = ("hello" + v24)
        v35 := GenS1{F0: 42, F1: (885 <= 61), F2: ("hello" + (v24 + v24))}
        _ = fi33
        _ = v34
        _ = v35
    }
    var v36 uint16 = 949
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
