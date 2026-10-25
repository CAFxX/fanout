package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 int, p1 int32) bool {
    var v7 int = (p0 & 32768)
    p1 = p1
    _ = p0
    _ = p1
    _ = v7
    return false
}
func genFn1(p0 int64) uint8 {
    switch 887 {
    case 128:
        if false {
            p0 += p0
            fmt.Println(p0)
        } else {
            v8 := []int{415, -851, -714}
            _ = v8
        }
    case 42:
        switch 4294967295 {
        case 181:
            var v9 int = 348
            _ = v9
        }
    case 660:
        var v10 uint64 = 4294967295
        _ = v10
    }
    var v11 int64 = p0
    fmt.Println(p0, v11)
    for fi12 := 0; fi12 < 17; fi12++ {
        v11 += -1
        _ = fi12
    }
    switch ("" + "a") {
    case "\n\t":
        var v13 bool = (502 != 388)
        _ = v13
    }
    _ = p0
    _ = v11
    return 255
}

func main() {
    v1 := make(map[int]int)
    var v2 uint64 = 42
    var v3 float32 = -2.25
    v4 := []string{("foo bar" + ("hello" + "a")), "", ("hello" + "\n\t"), ("\n\t" + "hello")}
    var v5 float32 = float32(1e10)
    var v6 float64 = ((float64((float32(-2.25) + v3)) - 1e10) + (1e10 * (((float64(v3) - (1.5 - 0.1)) + (float64(v3) * (float64(v3) * (1e10 + float64(0.0))))) / 1.0)))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v14 float32 = (0.0 / 1.5)
    var v15 uint16 = 773
    if false {
        _ = len(v1)
        _ = len(v1)
    }
    v16 := GenS1{F0: 127}
    v5 -= v5
    v17 := GenS1{F0: -1}
    _ = len(v4)
    v6 -= (v6 * float64(v5))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
