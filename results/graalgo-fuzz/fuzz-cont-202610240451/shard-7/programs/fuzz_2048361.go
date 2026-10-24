package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 float64 = 0.0
    v2 := make(map[int]uint32)
    v3 := make(map[string]bool)
    var v4 int32 = 0
    v5 := make(map[int]int)
    v6 := make(map[string]uint32)
    if true {
        switch 2147483647 {
        case 127:
            fmt.Println(v4, v4)
        }
        v1 = 0.1
    }
    var v7 float32 = float32(((((v1 - ((((((v1 - v1) - v1) * ((float64(0.0) - v1) + 0.1)) * v1) - 3.14159) / v1)) / -2.25) / ((((float64(0.0) / v1) / v1) - 0.1) * v1)) + v1))
    if ((-32768 != -1) || (1 >= 43)) {
        var v8 uint32 = 1
        if true {
            if false {
                v1 -= v1
            }
            var v9 float32 = v7
            _ = v9
        } else {
            for fi10 := 0; fi10 < 20; fi10++ {
                var v11 uint16 = 28
                _ = fi10
                _ = v11
            }
            v7 -= (v7 / v7)
        }
        _ = v8
    } else {
        v12 := GenS1{F0: -2147483648}
        var v13 string = "foo bar"
        _ = v12
        _ = v13
    }
    var v14 bool = (32767 == 42)
    fmt.Println(v7)
    var v15 uint8 = 255
    v16 := GenS1{F0: 128}
    switch "foo bar" {
    case "hello":
        for fi17 := 0; fi17 < 4; fi17++ {
            for fi18 := 0; fi18 < 3; fi18++ {
                if (18446744073709551615 >= 42) {
                    v19 := []int{(fi18 % ((9223372036854775807) | 1)), 32768}
                    var v20 uint64 = 4294967295
                    if (false && false) {
                        fmt.Println(v4, v1)
                        v14 = v14
                    } else {
                        switch "a" {
                        case "":
                            v20 = 1
                        case "a":
                            fmt.Println(v7)
                        }
                    }
                    _ = v19
                    _ = v20
                } else {
                    var v21 float32 = (v7 * v7)
                    fmt.Println(v14)
                    _ = v21
                }
                var v22 int32 = 1
                _ = fi18
                _ = v22
            }
            _ = fi17
        }
    case "foo bar":
        v1 -= v1
    }
    v15 += v15
    for fi23 := 0; fi23 < 9; fi23++ {
        var v24 string = "\n\t"
        for fi25 := 0; fi25 < 6; fi25++ {
            v7 = 1e10
            _ = fi25
        }
        var v26 int8 = 100
        _ = fi23
        _ = v24
        _ = v26
    }
    for fi27 := 0; fi27 < 7; fi27++ {
        v14 = true
        _ = fi27
    }
    var v28 float64 = (v1 + v1)
    v15 += v15
    v4 = 398
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
