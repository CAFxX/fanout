package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint64 }
type GenS2 struct { F0 uint8; F1 int32 }

func main() {
    var v1 uint16 = 42
    var v2 uint32 = 2147483647
    v3 := make(map[string]string)
    v4 := make(map[int]string)
    if (!(!true)) {
        v1 = 65535
        for fi5 := 0; fi5 < 2; fi5++ {
            if false {
                for fi6 := 0; fi6 < 4; fi6++ {
                    var v7 int64 = 1
                    _ = fi6
                    _ = v7
                }
                switch "foo bar" {
                case "":
                    fmt.Println(v1, v1)
                default:
                    v8 := v3["foo bar"]
                    _ = v8
                }
            }
            var v9 string = ("hello" + ("a" + ("a" + "a")))
            _ = fi5
            _ = v9
        }
    }
    v10 := v3[("foo bar" + ("a" + "\n\t"))]
    if (false || (-30 <= -39)) {
        var v11 float32 = float32(0.1)
        var v12 float32 = v11
        var v13 uint8 = 187
        _ = v11
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: 0, F1: -128, F2: 18446744073709551615}
    var v15 uint64 = 2
    fmt.Println(v10, v10, v15)
    for fi16 := 0; fi16 < 17; fi16++ {
        v1 = v1
        v17 := GenS1{F0: 32767, F1: 127, F2: (^(v15))}
        v1 = (v1 - (^(v1)))
        _ = fi16
        _ = v17
    }
    v10 = v10
    var v18 int16 = 639
    v19 := GenS2{F0: 108, F1: -2147483648}
    for fi20 := 0; fi20 < 6; fi20++ {
        for fi21 := 0; fi21 < 15; fi21++ {
            v22 := []int{fi21, (fi21 >> ((841) & 63)), fi21, (fi20 & 2147483647)}
            _ = fi21
            _ = v22
        }
        v23 := GenS1{F0: v18, F1: 89, F2: v15}
        _ = fi20
        _ = v23
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
