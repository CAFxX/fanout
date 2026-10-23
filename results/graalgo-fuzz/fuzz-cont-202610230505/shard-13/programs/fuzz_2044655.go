package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]int)
    var v3 int8 = 100
    var v4 int64 = -1
    delete(v1, ("a" + ("a" + "hello")))
    v4 -= (v4 >> ((42) & 63))
    if false {
        for fi5 := 0; fi5 < 8; fi5++ {
            if true {
                v6 := GenS1{F0: (true || true)}
                _ = v6
            } else {
                var v7 int8 = (v3 - v3)
                var v8 uint = 132
                if false {
                    fmt.Println(v7, v3, fi5)
                } else {
                    delete(v1, "hello")
                    var v9 int16 = 1000
                    var v10 int16 = (v9 << ((v8) & 63))
                    _ = v10
                    _ = v9
                }
                _ = v7
                _ = v8
            }
            v2[""] = (fi5 >> ((0) & 63))
            _ = fi5
        }
    }
    if true {
        v3 = (-(v3))
        if (2 > 18446744073709551615) {
            var v11 string = "foo bar"
            v4 -= v4
            _ = v11
        }
    }
    var v12 uint = 2
    if true {
        v13 := GenS1{F0: (true && (-793 < 82))}
        _ = v13
    }
    for fi14 := 0; fi14 < 10; fi14++ {
        var v15 int8 = v3
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: true}
    v17 := []int{-9223372036854775808, -812}
    v18 := GenS1{F0: true}
    for fi19 := 0; fi19 < 2; fi19++ {
        v3 += 1
        var v20 uint32 = 0
        v20 += (v20 * v20)
        _ = fi19
        _ = v20
    }
    v3 += v3
    _ = len(v2)
    if true {
        var v21 int32 = int32(v12)
        _ = v21
    } else {
        v12 -= (v12 / ((2) | 1))
        var v22 uint = v12
        v12 = (v22 * (v22 ^ (v12 >> ((v12) & 63))))
        _ = v22
    }
    v23 := GenS1{F0: false}
    var v24 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
