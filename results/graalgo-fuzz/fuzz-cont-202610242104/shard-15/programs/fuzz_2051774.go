package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16; F1 int64; F2 uint16 }

func main() {
    v1 := make(map[string]int)
    var v2 float64 = (-2.25 * ((-2.25 * 1.5) - (0.1 * (0.0 * (1.5 + (float64(1.5) * -2.25))))))
    var v3 bool = true
    if (!false) {
        if v3 {
            var v4 uint64 = 1
            if v3 {
                v5 := GenS1{F0: -1}
                _ = v5
            }
            _ = v4
        } else {
            var v6 uint = 0
            v2 -= ((v2 / (v2 * (3.14159 - v2))) * v2)
            _ = v6
        }
        if (v3 || (1000 >= -652)) {
            fmt.Println(v3, v2)
            for fi7 := 0; fi7 < 16; fi7++ {
                v8 := GenS1{F0: 127}
                _ = fi7
                _ = v8
            }
        } else {
            if v3 {
                v9 := GenS2{F0: 18446744073709551615, F1: 42}
                v10 := GenS3{F0: 1}
                fmt.Println(v2, v2)
                _ = v10
                _ = v9
            }
            for fi11 := 0; fi11 < 2; fi11++ {
                var v12 float32 = 0.0
                _ = len(v1)
                _ = fi11
                _ = v12
            }
            v13 := []int{148, -2147483648}
            _ = v13
        }
        v3 = false
    } else {
        var v14 int8 = 98
        _ = v14
    }
    if false {
        for fi15 := 0; fi15 < 6; fi15++ {
            v16 := GenS3{F0: 1}
            v17 := GenS2{F0: 2, F1: 0}
            _ = fi15
            _ = v16
            _ = v17
        }
        var v18 float32 = float32(1e10)
        switch "a" {
        case "hello":
            for fi19 := 0; fi19 < 4; fi19++ {
                fmt.Println(fi19, v3, v18)
                _ = fi19
            }
        case "\n\t":
            v18 = 0.1
        }
        _ = v18
    } else {
        fmt.Println(v3)
        var v20 int8 = 108
        _ = v20
    }
    if v3 {
        v21 := []int{123, 127, 465}
        fmt.Println(v2, v2)
        _ = v21
    } else {
        _ = len(v1)
    }
    for fi22 := 0; fi22 < 4; fi22++ {
        v3 = (2 > 274)
        _ = fi22
    }
    v2 = float64((1.5 - -2.25))
    var v23 uint = 65535
    var v24 uint64 = 668
    v25 := v1[""]
    if (2 > (v23 - (v23 - v23))) {
        switch v25 {
        case -9223372036854775808:
            v24 = (v24 / (((v24 << (((^(v23))) & 63))) | 1))
        default:
            var v26 int32 = 1
            _ = v26
        }
        for fi27 := 0; fi27 < 11; fi27++ {
            for fi28 := 0; fi28 < 4; fi28++ {
                v29 := []int{}
                fmt.Println(v25, fi28, v25)
                v30 := GenS2{F0: (v24 * (^(v24))), F1: 246}
                fmt.Println(fi27, v25, v3)
                var v31 uint64 = (v24 % (((v24 | (v24 % ((v24) | 1)))) | 1))
                if false {
                    v32 := GenS2{F0: 4294967295, F1: 127}
                    v33 := GenS2{F0: (^(v24)), F1: 20}
                    _ = v32
                    _ = v33
                } else {
                    _ = len(v1)
                }
                _ = fi28
                _ = v29
                _ = v30
                _ = v31
            }
            fmt.Println(v23, v3, v2)
            v25 = v25
            _ = fi27
        }
    }
    delete(v1, (("hello" + "foo bar") + "hello"))
    var v34 float32 = (float32(v2) - 1.5)
    for fi35 := 0; fi35 < 10; fi35++ {
        fmt.Println(fi35, fi35)
        v36 := GenS4{F0: 32767, F1: -1, F2: 534}
        var v37 uint64 = (v24 & v24)
        _ = fi35
        _ = v36
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v34
}
