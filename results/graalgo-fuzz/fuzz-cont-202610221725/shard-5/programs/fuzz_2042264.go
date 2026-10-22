package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 int64, p1 int64) int8 {
    var v6 uint = 1
    v7 := []int{-851, -1}
    if true {
        _ = len(v7)
        if true {
            var v8 uint = 929
            v9 := GenS1{F0: 380}
            _ = v8
            _ = v9
        }
        for fi10 := 0; fi10 < 1; fi10++ {
            if false {
                p1 += 579
                var v11 uint = (v6 * v6)
                for fi12 := 0; fi12 < 3; fi12++ {
                    var v13 float32 = 1e10
                    v6 -= (^(v11))
                    _ = fi12
                    _ = v13
                }
                _ = v11
            }
            fmt.Println(p0, p1)
            fmt.Println(fi10)
            _ = fi10
        }
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 79
}
func genFn1() uint32 {
    var v14 uint64 = 1
    var v15 uint = 65535
    _ = v14
    _ = v15
    return 964
}

func main() {
    var v1 uint32 = 2147483648
    var v2 uint64 = 4294967295
    var v3 int = 128
    v4 := make(map[string]int)
    var v5 float64 = -2.25
    v3 = (^(v3))
    v2 += 4294967295
    if false {
        v16 := []int{(v3 - v3), (v3 * (v3 / ((v3) | 1)))}
        switch 4294967295 {
        case 9223372036854775807:
            delete(v4, "hello")
        case 42:
            var v17 int32 = -1
            _ = v17
        case -534:
            for fi18 := 0; fi18 < 5; fi18++ {
                v19 := GenS1{F0: (v2 * 18446744073709551615)}
                var v20 int8 = 1
                v1 = 5
                _ = fi18
                _ = v19
                _ = v20
            }
        }
        v16 = append(v16, v3)
        _ = v16
    }
    var v21 int16 = -32768
    if false {
        if (v5 < (v5 - (v5 - float64((1e10 - 0.0))))) {
            delete(v4, "foo bar")
        }
        if (!true) {
            v22 := []int{}
            _ = v22
        }
    } else {
        if false {
            var v23 bool = (128 < 52)
            var v24 int16 = 32767
            _ = v23
            _ = v24
        }
    }
    switch int(v2) {
    case 128:
        v3 += v3
    case 32768:
        v25 := GenS1{F0: (v2 * 820)}
        _ = v25
    default:
        v5 = 1.5
    }
    var v26 int = 42
    switch "hello" {
    case "foo bar":
        fmt.Println(v1, v21, v2)
    default:
        v4[("\n\t" + "foo bar")] = v3
    }
    for fi27 := 0; fi27 < 4; fi27++ {
        var v28 float64 = 0.0
        fmt.Println(v1)
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
