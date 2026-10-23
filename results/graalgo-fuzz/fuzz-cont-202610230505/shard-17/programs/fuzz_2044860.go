package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int32 }
func genFn0(p0 int64) uint {
    var v5 float32 = 1.5
    var v6 string = ("hello" + "a")
    v7 := GenS1{F0: p0, F1: 675, F2: -215}
    var v8 int8 = 1
    fmt.Println(v8)
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 482
}

func main() {
    v1 := make(map[string]int)
    var v2 string = ("\n\t" + "a")
    var v3 float32 = 0.0
    var v4 uint16 = 255
    v1[(v2 + v2)] = 32768
    if false {
        v9 := []int{32767, -2147483648, -934, 9223372036854775807}
        _ = v9
    }
    if (false && (!false)) {
        v3 += ((1.5 + float32(-2.25)) + v3)
        for fi10 := 0; fi10 < 20; fi10++ {
            v11 := v1[v2]
            var v12 uint = 4294967295
            _ = fi10
            _ = v11
            _ = v12
        }
        fmt.Println(v3)
    } else {
        for fi13 := 0; fi13 < 6; fi13++ {
            for fi14 := 0; fi14 < 3; fi14++ {
                v2 = ((v2 + "foo bar") + v2)
                fmt.Println(v4)
                delete(v1, (("a" + "") + "\n\t"))
                _ = fi14
            }
            var v15 string = v2
            for fi16 := 0; fi16 < 9; fi16++ {
                fmt.Println(fi13, fi16)
                _ = fi16
            }
            _ = fi13
            _ = v15
        }
    }
    switch "foo bar" {
    case "\n\t":
        v17 := GenS1{F0: -1, F1: -9223372036854775808, F2: 2147483647}
        _ = v17
    case "a":
        fmt.Println(v2)
    }
    v1["foo bar"] = -1
    v4 -= (v4 * 63)
    if true {
        fmt.Println(v2)
        var v18 uint8 = 202
        _ = v18
    } else {
        v2 = (("" + (((("a" + (v2 + "foo bar")) + v2) + v2) + v2)) + (v2 + (v2 + ("\n\t" + "foo bar"))))
        v2 = (v2 + (v2 + "\n\t"))
    }
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
