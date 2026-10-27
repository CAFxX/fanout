package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int16 }
func genFn0(p0 uint32, p1 uint, p2 string) int16 {
    if false {
        for fi4 := 0; fi4 < 7; fi4++ {
            var v5 int16 = 32767
            _ = fi4
            _ = v5
        }
        if false {
            var v6 float32 = 0.1
            p0 -= p0
            var v7 int16 = -1
            _ = v6
            _ = v7
        } else {
            var v8 float32 = ((-2.25 + 1.5) + ((3.14159 + ((-2.25 / 1.5) * 1.5)) / -2.25))
            p0 -= 979
            _ = v8
        }
        for fi9 := 0; fi9 < 13; fi9++ {
            var v10 int32 = -1
            fmt.Println(p2, p2)
            fmt.Println(fi9)
            _ = fi9
            _ = v10
        }
    }
    fmt.Println(p1, p0)
    p1 += 2
    _ = p0
    _ = p1
    _ = p2
    return 1000
}
func genFn1() int8 {
    v11 := []int{2147483647, 701, 4294967295}
    if (!("hello" <= "hello")) {
        var v12 bool = false
        if false {
            fmt.Println(v12)
            var v13 float64 = (3.14159 + (-2.25 * -2.25))
            _ = v13
        }
        var v14 uint64 = 686
        _ = v12
        _ = v14
    }
    var v15 uint = 2
    v16 := GenS1{F0: 0, F1: 795, F2: -32768}
    _ = v11
    _ = v15
    _ = v16
    return -1
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[int]int)
    var v3 float32 = 0.0
    v17 := []int{2147483647, -9223372036854775808}
    v1 = v1
    v17 = append(v17, 32768)
    v2[-2147483648] = 9223372036854775807
    v18 := GenS1{F0: 1, F1: 2147483648, F2: -228}
    v19 := GenS1{F0: 32767, F1: 828, F2: -298}
    v1 = v1
    switch (((("" + "a") + (("\n\t" + "hello") + ("hello" + ("" + "\n\t")))) + "") + "\n\t") {
    case "":
        var v20 float64 = 3.14159
        _ = v20
    }
    v3 += (v3 - (v3 * (((v3 / 1.0) / v3) * (v3 + 1.5))))
    switch 128 {
    case -1:
        v3 += 3.14159
    }
    var v21 int64 = 9223372036854775807
    fmt.Println(v1, v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
}
