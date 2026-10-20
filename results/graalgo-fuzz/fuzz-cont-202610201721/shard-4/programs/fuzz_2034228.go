package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 int16; F2 string }
func genFn0() string {
    var v4 bool = false
    if v4 {
        v4 = (255 == 962)
    }
    v5 := []int{342, -1}
    _ = v4
    _ = v5
    return ("hello" + ("\n\t" + "\n\t"))
}

func main() {
    v1 := make(map[string]int)
    var v2 int64 = 386
    var v3 string = (("a" + "") + "")
    fmt.Println(v3)
    v2 -= v2
    if true {
        fmt.Println(v2, v3)
        for fi6 := 0; fi6 < 12; fi6++ {
            fmt.Println(fi6)
            v7 := GenS1{F0: 2147483648, F1: 1000, F2: 2}
            v3 = "a"
            _ = fi6
            _ = v7
        }
        var v8 int = 32768
        _ = v8
    } else {
        v9 := GenS2{F0: 127}
        for fi10 := 0; fi10 < 19; fi10++ {
            _ = len(v1)
            v11 := GenS2{F0: fi10}
            var v12 float32 = 3.14159
            _ = fi10
            _ = v11
            _ = v12
        }
        var v13 uint = 1
        _ = v13
        _ = v9
    }
    v14 := GenS2{F0: 156}
    switch "a" {
    case "hello":
        v3 = v3
    case "\n\t":
        var v15 int32 = -796
        _ = v15
    case "foo bar":
        v16 := []int{-9223372036854775808, -910}
        _ = v16
    default:
        v2 = (v2 | (v2 >> ((1) & 63)))
    }
    v17 := GenS1{F0: 4294967295, F1: 0, F2: 65535}
    v2 += v2
    v2 += v2
    var v18 int16 = -408
    _ = len(v1)
    v19 := GenS3{F0: 127, F1: v18, F2: v3}
    v18 += int16(v2)
    fmt.Println(v18, v18)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
