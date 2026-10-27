package main

import (
    "fmt"
)

func genFn0() uint32 {
    var v5 uint32 = 65535
    switch ("a" + "") {
    case "foo bar":
        v5 = (v5 << ((255) & 63))
    case "":
        var v6 uint8 = 1
        _ = v6
    }
    v5 = (v5 >> ((42) & 63))
    _ = v5
    return v5
}
func genFn1(p0 int64, p1 uint64) uint8 {
    p0 -= (p0 - (p0 + (p0 / ((9223372036854775807) | 1))))
    var v7 int8 = 54
    v8 := []int{-9223372036854775808, 32767}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 int16 = 0
    var v2 int64 = 0
    var v3 int = 127
    v4 := []string{"foo bar"}
    var v9 uint64 = 18446744073709551615
    fmt.Println(v2)
    var v10 uint = 0
    for fi11 := 0; fi11 < 1; fi11++ {
        v3 += fi11
        var v12 float32 = 0.0
        v4 = append(v4, (("hello" + "a") + ("hello" + "a")))
        _ = fi11
        _ = v12
    }
    fmt.Println(v10)
    switch (v3 + v3) {
    case 284:
        var v13 uint8 = 75
        _ = v13
    case 4294967295:
        var v14 int = (v3 - -841)
        _ = v14
    default:
        var v15 int64 = (^(v2))
        _ = v15
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        v2 -= v2
        v10 += (^(v10))
        if (!false) {
            var v17 uint64 = v9
            _ = v17
        } else {
            v9 -= v9
        }
        _ = fi16
    }
    if false {
        fmt.Println(v10, v10)
        var v18 uint32 = 2147483648
        _ = v18
    }
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        fmt.Println(v1, v10, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
