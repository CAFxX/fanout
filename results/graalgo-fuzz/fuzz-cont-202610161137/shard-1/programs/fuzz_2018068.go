package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint; F1 uint64 }
func genFn0() int8 {
    for fi4 := 0; fi4 < 6; fi4++ {
        var v5 string = "a"
        var v6 float32 = (-2.25 - 0.1)
        v6 -= 0.0
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := []int{42, 690}
    switch 1 {
    case -2147483648:
        var v8 int8 = -1
        _ = v8
    }
    _ = v7
    return -128
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]bool)
    var v3 int = 4294967295
    for fi9 := 0; fi9 < 7; fi9++ {
        if (63 == 95) {
            var v10 bool = (42 != -9223372036854775808)
            var v11 bool = ((!true) || true)
            _ = v10
            _ = v11
        } else {
            if true {
                if (-1 < -1) {
                    fmt.Println(v3, fi9)
                    fmt.Println(v3)
                }
            }
            v3 = (^(v3))
        }
        var v12 float32 = 0.0
        v13 := GenS1{F0: int16(fi9), F1: 126}
        _ = fi9
        _ = v12
        _ = v13
    }
    v3 -= 32767
    v14 := GenS2{F0: 1, F1: ""}
    var v15 float64 = (-2.25 - ((float64(1e10) + -2.25) - 0.0))
    v16 := GenS3{F0: 65535, F1: 42}
    fmt.Println(v3)
    var v17 int16 = 525
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v15)
    default:
        var v18 uint = 0
        _ = v18
    }
    if false {
        var v19 int32 = -712
        v20 := []int{v3, int(v17), -271, v3}
        for fi21 := 0; fi21 < 18; fi21++ {
            var v22 float64 = 1.5
            _ = fi21
            _ = v22
        }
        _ = v19
        _ = v20
    }
    var v23 int16 = v17
    for fi24 := 0; fi24 < 6; fi24++ {
        v25 := GenS3{F0: 336, F1: 768}
        var v26 uint8 = 127
        for fi27 := 0; fi27 < 2; fi27++ {
            var v28 float32 = 1.5
            v29 := GenS3{F0: 2, F1: 42}
            switch (fi24 % ((fi24) | 1)) {
            case 2147483647:
                var v30 int16 = 1000
                _ = v30
            }
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v3
}
