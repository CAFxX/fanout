package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 string = ("foo bar" + ("hello" + ("" + (("hello" + ((("foo bar" + "") + "hello") + ("a" + "hello"))) + ((("foo bar" + (("hello" + ("\n\t" + "\n\t")) + "\n\t")) + (("foo bar" + ("" + "\n\t")) + "hello")) + "\n\t")))))
    var v2 float64 = (-2.25 + 0.0)
    v3 := make(map[string]bool)
    fmt.Println(v1, v1)
    var v4 float64 = 0.1
    if true {
        for fi5 := 0; fi5 < 19; fi5++ {
            v6 := GenS1{F0: 18446744073709551615, F1: 2}
            _ = fi5
            _ = v6
        }
    } else {
        var v7 int32 = -2147483648
        fmt.Println(v4, v4)
        for fi8 := 0; fi8 < 12; fi8++ {
            switch -9223372036854775808 {
            case -46:
                v9 := GenS2{F0: 532}
                _ = v9
            case -9223372036854775808:
                v2 = (0.1 - float64((1e10 + (float32(3.14159) - ((3.14159 + 0.0) * 3.14159)))))
            default:
                if (true && true) {
                    switch 32767 {
                    case 9223372036854775807:
                        v10 := []int{(fi8 / ((1) | 1)), 32767, fi8}
                        _ = v10
                    }
                    v11 := v3[v1]
                    _ = v11
                }
            }
            _ = fi8
        }
        _ = v7
    }
    var v12 uint16 = 141
    fmt.Println(v2, v1, v2)
    v13 := GenS1{F0: 1, F1: 171}
    v4 += (v2 + v2)
    if false {
        switch 1 {
        case 2147483647:
            switch 42 {
            case 2147483647:
                if false {
                    var v14 uint = 65535
                    var v15 int8 = 100
                    _ = v14
                    _ = v15
                }
            }
        case -1:
            v16 := []int{}
            _ = v16
        }
        var v17 int64 = 42
        if (-2.25 >= (float32(v2) + 1e10)) {
            fmt.Println(v2)
            v18 := v3[(v1 + v1)]
            var v19 int8 = int8(v12)
            _ = v18
            _ = v19
        }
        _ = v17
    }
    v4 += v4
    v20 := v3[v1]
    v21 := GenS1{F0: 141, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
