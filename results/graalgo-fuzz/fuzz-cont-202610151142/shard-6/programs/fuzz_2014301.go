package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 bool }

func main() {
    var v1 float64 = 1e10
    var v2 int64 = 42
    v3 := make(map[string]int)
    var v4 uint64 = 4294967295
    v5 := make(map[string]int)
    var v6 string = ("a" + ("" + "hello"))
    v1 -= float64(3.14159)
    v2 += (v2 / (((v2 * (v2 ^ (v2 >> ((2) & 63))))) | 1))
    var v7 uint32 = 0
    var v8 int = 2147483647
    switch (v8 & (v8 << ((0) & 63))) {
    case 32767:
        switch (v6 + ("\n\t" + v6)) {
        case "\n\t":
            fmt.Println(v7, v4)
        case "foo bar":
            for fi9 := 0; fi9 < 16; fi9++ {
                if false {
                    fmt.Println(v4)
                }
                _ = fi9
            }
        case "":
            v10 := GenS1{F0: 1, F1: 18446744073709551615}
            _ = v10
        }
    case -1:
        if (84 >= 0) {
            if false {
                switch (v8 * (v8 << ((159) & 63))) {
                case 873:
                    for fi11 := 0; fi11 < 11; fi11++ {
                        var v12 int = fi11
                        _ = fi11
                        _ = v12
                    }
                case 42:
                    var v13 string = v6
                    _ = v13
                case -9223372036854775808:
                    for fi14 := 0; fi14 < 1; fi14++ {
                        v7 -= (v7 - v7)
                        v15 := GenS1{F0: -128, F1: 4294967295}
                        _ = fi14
                        _ = v15
                    }
                default:
                    v16 := []int{v8}
                    _ = v16
                }
            }
            var v17 uint32 = v7
            _ = v17
        }
    default:
        v18 := GenS1{F0: -24, F1: 4294967295}
        _ = v18
    }
    v7 += (v7 * (^(v7)))
    var v19 uint64 = (^(v4))
    if true {
        var v20 int64 = v2
        _ = v20
    } else {
        _ = len(v5)
        fmt.Println(v1, v6, v4)
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        if true {
            v4 -= v19
            var v22 bool = false
            v23 := []int{(fi21 & fi21), fi21, v8, (fi21 >> ((18446744073709551615) & 63))}
            _ = v22
            _ = v23
        } else {
            v24 := GenS1{F0: 0, F1: 42}
            var v25 int32 = -169
            v1 = (((float64(1.5) / (float64(((float32(v1) + 1.5) + 1e10)) / v1)) / (v1 * ((float64(-2.25) + v1) * (3.14159 / v1)))) + (v1 / 1.0))
            _ = v24
            _ = v25
        }
        v4 += (v4 << ((476) & 63))
        _ = fi21
    }
    v2 -= 1
    if (false && false) {
        v4 = (v19 - v4)
        var v26 bool = true
        _ = v26
    }
    v1 = 1.5
    fmt.Println(v4)
    v1 = v1
    v27 := GenS1{F0: 1, F1: 4294967295}
    v8 = 4294967295
    v28 := GenS1{F0: 100, F1: 4294967295}
    for fi29 := 0; fi29 < 17; fi29++ {
        for fi30 := 0; fi30 < 17; fi30++ {
            if (!false) {
                switch ((v6 + ("foo bar" + "")) + (v6 + v6)) {
                case "a":
                    v6 = ("a" + "")
                case "hello":
                    v8 = (-(fi30))
                }
            }
            var v31 uint32 = 42
            _ = fi30
            _ = v31
        }
        v7 -= v7
        switch -573 {
        case 127:
            var v32 int64 = (^(v2))
            _ = v32
        case 2147483647:
            for fi33 := 0; fi33 < 5; fi33++ {
                v34 := []int{(fi29 | fi33), v8}
                var v35 uint8 = 42
                v36 := GenS1{F0: 70, F1: 42}
                _ = fi33
                _ = v34
                _ = v35
                _ = v36
            }
        case 0:
            v37 := GenS2{F0: (-1 < 32767)}
            _ = v37
        default:
            v4 -= v4
        }
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
