package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int32 }

func main() {
    var v1 int64 = 42
    var v2 int16 = 0
    v3 := make(map[string]bool)
    var v4 bool = (((0.0 + 1e10) - 1.5) <= -2.25)
    var v5 uint8 = 0
    v6 := GenS1{F0: 904, F1: uint64(v5), F2: 1}
    v3["foo bar"] = true
    for fi7 := 0; fi7 < 20; fi7++ {
        switch ("hello" + (("hello" + "") + "")) {
        case "foo bar":
            var v8 int = (fi7 % ((154) | 1))
            _ = v8
        case "hello":
            v5 = v5
        }
        var v9 string = "\n\t"
        for fi10 := 0; fi10 < 6; fi10++ {
            fmt.Println(v4, v4)
            v4 = (!(v5 < 177))
            v11 := []int{fi10, fi7}
            _ = fi10
            _ = v11
        }
        _ = fi7
        _ = v9
    }
    if (635 >= 255) {
        var v12 int32 = -98
        _ = v12
    } else {
        switch 9223372036854775807 {
        case -726:
            v2 = int16(v1)
        case -194:
            var v13 int8 = 8
            _ = v13
        }
        for fi14 := 0; fi14 < 9; fi14++ {
            var v15 uint = 255
            _ = fi14
            _ = v15
        }
        var v16 int16 = v2
        _ = v16
    }
    v2 = (v2 * v2)
    v1 += (^(v1))
    if (118 <= 216) {
        v1 += (v1 | v1)
    } else {
        var v17 int8 = -1
        _ = v17
    }
    for fi18 := 0; fi18 < 2; fi18++ {
        var v19 float64 = -2.25
        v4 = (("" != "hello") && v4)
        _ = fi18
        _ = v19
    }
    delete(v3, "a")
    fmt.Println(v2, v2)
    fmt.Println(v2, v2, v4)
    var v20 int16 = (v2 % ((-1) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
