package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16 }
func genFn0() int16 {
    var v5 float64 = (0.1 / 1.0)
    v6 := []int{42, 42, 0, 128}
    var v7 uint8 = 128
    var v8 int8 = 0
    v6 = append(v6, 2147483647)
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 1000
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]uint32)
    var v3 uint8 = 1
    var v4 uint64 = 42
    var v9 bool = (((-2.25 / (float64(1e10) - 1.5)) / 0.1) <= float64(0.0))
    _ = len(v1)
    v10 := GenS1{F0: -253, F1: ("foo bar" + ("hello" + "hello"))}
    v3 = 127
    for fi11 := 0; fi11 < 12; fi11++ {
        v12 := GenS2{F0: 4294967295}
        for fi13 := 0; fi13 < 14; fi13++ {
            v9 = (!((0 <= 42) && v9))
            _ = fi13
        }
        _ = fi11
        _ = v12
    }
    fmt.Println(v4, v9)
    fmt.Println(v3)
    v9 = ((v4 >= (^(v4))) && (!(v9 && (!v9))))
    switch "a" {
    case "\n\t":
        fmt.Println(v9, v4, v3)
    case "":
        var v14 int16 = 225
        _ = v14
    default:
        if false {
            v4 = (v4 * v4)
            v15 := []int{32768, 127}
            var v16 uint16 = 61
            _ = v15
            _ = v16
        }
    }
    fmt.Println(v3, v4)
    v17 := GenS3{F0: genFn0()}
    if false {
        if (3.14159 < ((1e10 / 1.0) / 0.1)) {
            if (-1 >= -2147483648) {
                v9 = false
                var v18 string = ""
                v4 -= (v4 - 0)
                _ = v18
            } else {
                fmt.Println(v4)
            }
        }
        v9 = v9
        v4 += v4
    } else {
        fmt.Println(v3, v3)
        v9 = v9
        for fi19 := 0; fi19 < 10; fi19++ {
            v9 = (1 >= genFn0())
            _ = fi19
        }
    }
    v4 = 0
    for fi20 := 0; fi20 < 20; fi20++ {
        for fi21 := 0; fi21 < 2; fi21++ {
            fmt.Println(v9, fi21, v9)
            switch (fi21 | (-(fi21))) {
            case 42:
                v22 := GenS1{F0: -32768, F1: (("hello" + ("a" + "a")) + ("" + "\n\t"))}
                _ = v22
            }
            _ = fi21
        }
        var v23 string = ((("a" + "a") + "\n\t") + ("\n\t" + "hello"))
        _ = fi20
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
