package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
func genFn0() int64 {
    v7 := []int{1, -1, 9223372036854775807, 507}
    if false {
        var v8 int64 = 42
        _ = v8
    } else {
        var v9 uint = 256
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 int16 = 0
            if true {
                var v12 float64 = float64(3.14159)
                fmt.Println(fi10)
                _ = v12
            } else {
                fmt.Println(fi10, v9, fi10)
                if (42 != -2147483648) {
                    v11 = v11
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    _ = len(v7)
                }
            }
            _ = fi10
            _ = v11
        }
        _ = v9
    }
    v13 := GenS1{F0: -829, F1: 0}
    v7 = append(v7, -1)
    var v14 uint32 = 0
    _ = v13
    _ = v14
    _ = v7
    return -9223372036854775808
}

func main() {
    var v1 bool = ("a" >= "hello")
    v2 := make(map[string]int)
    var v3 int16 = 0
    var v4 int8 = 0
    v5 := make(map[int]int)
    v6 := make(map[uint32]bool)
    fmt.Println(v1, v3, v4)
    v1 = false
    for fi15 := 0; fi15 < 13; fi15++ {
        var v16 int8 = v4
        _ = fi15
        _ = v16
    }
    var v17 uint32 = 542
    v18 := GenS1{F0: -251, F1: 1}
    v4 = v4
    if v1 {
        var v19 uint32 = 340
        var v20 float32 = 0.0
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        var v22 bool = (v1 && ((v4 & (v4 % (((v4 >> ((65535) & 63))) | 1))) > 0))
        switch 42 {
        case 358:
            if (1e10 != 1.5) {
                var v23 string = ""
                v22 = v22
                _ = v23
            } else {
                var v24 uint64 = 4294967295
                var v25 int32 = 341
                var v26 uint64 = 777
                _ = v24
                _ = v25
                _ = v26
            }
        case 2147483647:
            for fi27 := 0; fi27 < 9; fi27++ {
                var v28 float64 = (-2.25 + 0.1)
                _ = fi27
                _ = v28
            }
        }
        v3 = 617
        _ = fi21
        _ = v22
    }
    var v29 string = "a"
    var v30 int8 = v4
    if (0.0 > 0.0) {
        fmt.Println(v29, v30, v29)
    }
    v29 = ((v29 + v29) + (v29 + (("foo bar" + "foo bar") + ("foo bar" + "\n\t"))))
    var v31 uint = 65535
    if true {
        var v32 int8 = v30
        v1 = v1
        var v33 uint16 = 784
        _ = v32
        _ = v33
    } else {
        for fi34 := 0; fi34 < 17; fi34++ {
            fmt.Println(v3, v31)
            var v35 string = "a"
            _ = fi34
            _ = v35
        }
        var v36 int8 = (v30 * 127)
        var v37 float32 = ((0.1 / 1.0) * 1.5)
        _ = v36
        _ = v37
    }
    v38 := GenS1{F0: int64(v17), F1: 1}
    switch v29 {
    case "hello":
        fmt.Println(v4, v31, v3)
    case "\n\t":
        v39 := []int{128, 2147483647, 4294967295, -2147483648}
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
