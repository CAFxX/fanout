package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int8 }
func genFn0(p0 int) int16 {
    if true {
        p0 = p0
        v5 := []int{}
        v6 := GenS1{F0: 18446744073709551615, F1: 0, F2: 42}
        _ = v5
        _ = v6
    } else {
        var v7 float32 = (1.5 + 1e10)
        _ = v7
    }
    if true {
        var v8 int16 = -32768
        fmt.Println(p0, v8)
        if (!false) {
            var v9 int64 = -1
            _ = v9
        }
        _ = v8
    }
    v10 := []int{p0}
    for fi11 := 0; fi11 < 11; fi11++ {
        p0 -= (fi11 >> ((4294967295) & 63))
        _ = fi11
    }
    _ = p0
    _ = v10
    return 481
}

func main() {
    v1 := []string{"hello", "", "\n\t"}
    var v2 string = ""
    var v3 int16 = 0
    var v4 int8 = 42
    v12 := GenS1{F0: 42, F1: (v3 ^ v3), F2: v4}
    if false {
        if (0 <= -405) {
            v3 = (v3 >> ((346) & 63))
        }
    }
    v13 := GenS1{F0: 961, F1: (v3 + (v3 ^ 42)), F2: (v4 << ((18446744073709551615) & 63))}
    if true {
        for fi14 := 0; fi14 < 20; fi14++ {
            v15 := GenS1{F0: 4294967295, F1: 32767, F2: (v4 & v4)}
            var v16 bool = true
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    if len(v1) > 0 {
        v1[0] = v2
    }
    fmt.Println(v2)
    if (!false) {
        for fi17 := 0; fi17 < 7; fi17++ {
            for fi18 := 0; fi18 < 20; fi18++ {
                var v19 int32 = 929
                fmt.Println(v4)
                v1 = append(v1, v2)
                _ = fi18
                _ = v19
            }
            v2 = ""
            fmt.Println(v2)
            _ = fi17
        }
        v2 = ""
    }
    var v20 bool = true
    fmt.Println(v20, v20)
    for fi21 := 0; fi21 < 2; fi21++ {
        v3 = (v3 >> ((42) & 63))
        fmt.Println(v3, v4, v20)
        _ = fi21
    }
    for fi22 := 0; fi22 < 11; fi22++ {
        fmt.Println(v3)
        v4 += v4
        var v23 int16 = genFn0((fi22 | -2147483648))
        _ = fi22
        _ = v23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
