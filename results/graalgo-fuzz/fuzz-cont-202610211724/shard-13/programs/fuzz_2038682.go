package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int8 }
func genFn0() int32 {
    var v5 float64 = 1e10
    var v6 string = ""
    var v7 string = v6
    var v8 int16 = 1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = 42
    var v3 float64 = ((0.0 * 0.1) * (0.1 * 1.5))
    var v4 int8 = -128
    var v9 float32 = (0.0 - -2.25)
    v10 := v1[4294967295]
    var v11 int64 = 376
    if v10 {
        var v12 string = (("a" + (("hello" + "foo bar") + "\n\t")) + "foo bar")
        var v13 float64 = v3
        _ = v12
        _ = v13
    } else {
        var v14 int8 = (-(v4))
        fmt.Println(v11, v10)
        v15 := []int{}
        _ = v14
        _ = v15
    }
    if ((2 > 210) || false) {
        fmt.Println(v11)
        switch 9223372036854775807 {
        case -9223372036854775808:
            for fi16 := 0; fi16 < 20; fi16++ {
                v17 := []int{}
                fmt.Println(v3, v9, v4)
                _ = fi16
                _ = v17
            }
        case -1:
            fmt.Println(v3, v9)
        }
        if false {
            if (255 >= 128) {
                if (float64(v9) >= 0.0) {
                    var v18 uint8 = 42
                    _ = v18
                }
                v19 := []int{4294967295, 9223372036854775807}
                v11 = v11
                _ = v19
            }
            switch 4294967295 {
            case -659:
                for fi20 := 0; fi20 < 19; fi20++ {
                    fmt.Println(fi20)
                    _ = fi20
                }
            case -521:
                var v21 int64 = v2
                _ = v21
            }
        }
    } else {
        fmt.Println(v3, v3, v4)
        v22 := GenS1{F0: (v2 | (^(v11))), F1: v4, F2: (v4 * -97)}
        _ = v22
    }
    switch "\n\t" {
    case "":
        var v23 float32 = ((-2.25 - float32(((v3 / 1.0) + float64(1e10)))) + v9)
        _ = v23
    case "a":
        v24 := []int{2147483647}
        _ = v24
    case "hello":
        var v25 int = 32767
        _ = v25
    }
    var v26 uint64 = 274
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v9
}
