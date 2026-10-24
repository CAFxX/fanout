package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int8; F1 int }
func genFn0() bool {
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 float64 = float64(-2.25)
        _ = fi5
        _ = v6
    }
    if false {
        v7 := GenS1{F0: -1}
        for fi8 := 0; fi8 < 20; fi8++ {
            for fi9 := 0; fi9 < 14; fi9++ {
                var v10 uint32 = 340
                var v11 uint = 1
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = fi8
        }
        _ = v7
    } else {
        v12 := []int{42, 4294967295, -2147483648, 2147483647}
        v12 = append(v12, -2147483648)
        _ = v12
    }
    return true
}
func genFn1(p0 uint64, p1 uint32, p2 int32) uint16 {
    p0 = p0
    if (false && ((float32((float64(-2.25) / 1e10)) / -2.25) < 1e10)) {
        p2 += -2147483648
        fmt.Println(p1, p0, p1)
        p2 = 418
    } else {
        if false {
            fmt.Println(p0, p1, p0)
            v13 := GenS1{F0: int32(p0)}
            _ = v13
        } else {
            var v14 int16 = -1
            _ = v14
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 65535
}

func main() {
    v1 := make(map[string]int)
    v2 := []string{("foo bar" + "\n\t")}
    v3 := []uint32{2, 2, 2147483648}
    var v4 bool = true
    v15 := GenS1{F0: 0}
    v4 = ((20 < 0) && v4)
    v1["foo bar"] = 128
    v4 = true
    var v16 uint8 = 42
    var v17 uint = 42
    for fi18 := 0; fi18 < 10; fi18++ {
        var v19 uint = (^(v17))
        var v20 string = "a"
        v2 = append(v2, ("" + (v20 + (("hello" + v20) + (v20 + v20)))))
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 int8 = -1
    var v22 uint8 = 33
    v23 := GenS2{F0: 1, F1: -898}
    fmt.Println(v21, v17)
    if (1.5 != 1.5) {
        var v24 uint64 = 0
        _ = len(v1)
        _ = v24
    } else {
        v17 = (v17 % (((v17 / (((^(v17))) | 1))) | 1))
    }
    for fi25 := 0; fi25 < 8; fi25++ {
        v26 := GenS3{F0: (v21 >> ((859) & 63)), F1: fi25}
        v27 := GenS2{F0: int32(v17), F1: -602}
        v4 = ((-2.25 - float32(1e10)) <= -2.25)
        _ = fi25
        _ = v26
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
