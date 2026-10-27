package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := []int{-288, -9223372036854775808}
    v2 := make(map[string]int)
    v3 := []string{("\n\t" + "hello"), ("a" + ("\n\t" + (("" + "") + ("hello" + "")))), "", "\n\t"}
    var v4 bool = false
    v5 := []uint32{398}
    v6 := make(map[string]uint32)
    v4 = v4
    fmt.Println(v4)
    v2[""] = 127
    v7 := GenS1{F0: v4}
    fmt.Println(v4)
    v4 = v4
    var v8 int64 = 494
    if v4 {
        fmt.Println(v4, v4)
        v8 -= v8
        v2["foo bar"] = -2147483648
    }
    v8 = v8
    if false {
        if (137 == 1) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            for fi9 := 0; fi9 < 2; fi9++ {
                var v10 string = ("hello" + "a")
                _ = fi9
                _ = v10
            }
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v11 int8 = -21
            _ = v11
        }
        var v12 uint64 = 2
        _ = v12
    }
    var v13 int8 = 21
    for fi14 := 0; fi14 < 7; fi14++ {
        delete(v2, "foo bar")
        v8 += v8
        v13 -= v13
        _ = fi14
    }
    v13 += (-(v13))
    fmt.Println(v13, v8)
    fmt.Println(v13, v4)
    var v15 float32 = (0.0 + -2.25)
    v4 = (268 == 340)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
