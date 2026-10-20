package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 string; F1 uint; F2 int16 }
func genFn0() int64 {
    v4 := []int{}
    switch ("foo bar" + "hello") {
    case "foo bar":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "hello":
        var v5 uint = 2
        _ = v5
    case "a":
        v6 := GenS1{F0: 0, F1: -1}
        _ = v6
    default:
        var v7 int64 = 1
        _ = v7
    }
    _ = v4
    return -859
}

func main() {
    var v1 float64 = -2.25
    v2 := make(map[string]string)
    v3 := make(map[uint32]bool)
    v8 := GenS2{F0: "a", F1: 42, F2: 0}
    fmt.Println(v1)
    v1 += 1.5
    var v9 int = -2147483648
    var v10 float32 = 1.5
    if (!(!false)) {
        v9 += 1
    }
    v10 += v10
    delete(v3, 424)
    if (-32768 == -134) {
        switch 4294967295 {
        case -9223372036854775808:
            v11 := GenS2{F0: "hello", F1: 0, F2: 1000}
            _ = v11
        case 964:
            switch ("hello" + (("foo bar" + (("a" + ("a" + "hello")) + "foo bar")) + ("foo bar" + ""))) {
            case "\n\t":
                var v12 uint32 = 1
                _ = v12
            default:
                if (1 >= 127) {
                    switch (v9 & (-(v9))) {
                    case 874:
                        delete(v3, 2)
                    }
                    v13 := GenS1{F0: 2147483648, F1: 0}
                    _ = v13
                }
            }
        default:
            for fi14 := 0; fi14 < 16; fi14++ {
                var v15 int32 = 796
                for fi16 := 0; fi16 < 18; fi16++ {
                    if true {
                        var v17 int32 = v15
                        v18 := []int{}
                        v19 := GenS1{F0: 1, F1: 100}
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    _ = fi16
                }
                fmt.Println(v10, v1, v9)
                _ = fi14
                _ = v15
            }
        }
    }
    v1 = (v1 - v1)
    if true {
        var v20 uint16 = 2
        _ = v20
    } else {
        fmt.Println(v10, v9)
    }
    switch "foo bar" {
    case "\n\t":
        v10 = 1e10
    case "foo bar":
        v21 := []int{v9, -2147483648, (-(v9))}
        _ = v21
    case "":
        var v22 uint64 = 728
        _ = v22
    default:
        var v23 float64 = (3.14159 - 0.0)
        _ = v23
    }
    switch (v9 << ((296) & 63)) {
    case 9223372036854775807:
        v10 = 3.14159
    case 4294967295:
        v9 = v9
    default:
        var v24 float64 = v1
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
