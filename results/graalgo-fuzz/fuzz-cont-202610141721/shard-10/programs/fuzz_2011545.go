package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 uint; F1 string; F2 int8 }

func main() {
    v1 := []int{32768, 42}
    v2 := []int{623, 0, -883, -100}
    v3 := make(map[int]bool)
    var v4 int64 = -363
    v4 = v4
    var v5 bool = ("" <= "foo bar")
    for fi6 := 0; fi6 < 5; fi6++ {
        fmt.Println(v4, fi6)
        var v7 int16 = 32767
        v7 += v7
        _ = fi6
        _ = v7
    }
    var v8 int16 = 32767
    var v9 uint = 4294967295
    var v10 int32 = -556
    var v11 int16 = v8
    for fi12 := 0; fi12 < 1; fi12++ {
        fmt.Println(v10)
        _ = fi12
    }
    v13 := GenS1{F0: 2, F1: "foo bar"}
    for fi14 := 0; fi14 < 9; fi14++ {
        if (!(true && v5)) {
            fmt.Println(v4, v4, v8)
            switch "" {
            case "\n\t":
                v15 := GenS2{F0: v9, F1: "", F2: 42}
                _ = v15
            case "a":
                v16 := GenS1{F0: 0, F1: ("foo bar" + "a")}
                _ = v16
            case "hello":
                v17 := v3[-9223372036854775808]
                _ = v17
            }
        } else {
            for fi18 := 0; fi18 < 17; fi18++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                fmt.Println(v4, fi14, v8)
                _ = fi18
            }
        }
        _ = fi14
    }
    v19 := GenS1{F0: 2, F1: "foo bar"}
    var v20 uint8 = 217
    for fi21 := 0; fi21 < 13; fi21++ {
        v9 = (v9 >> ((v9) & 63))
        _ = fi21
    }
    for fi22 := 0; fi22 < 12; fi22++ {
        var v23 int8 = 100
        switch "\n\t" {
        case "":
            fmt.Println(fi22)
        default:
            var v24 int64 = (v4 * v4)
            _ = v24
        }
        v23 -= (v23 * -5)
        _ = fi22
        _ = v23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
