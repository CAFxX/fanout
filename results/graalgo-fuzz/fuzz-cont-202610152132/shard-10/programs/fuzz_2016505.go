package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 int16, p1 uint8) int {
    var v4 string = ""
    if false {
        var v5 int32 = -1
        _ = v5
    } else {
        p1 += (^(p1))
        p0 += 0
    }
    _ = p0
    _ = p1
    _ = v4
    return 32768
}
func genFn1() int16 {
    switch "foo bar" {
    case "\n\t":
        var v6 int64 = 1
        _ = v6
    case "hello":
        var v7 int16 = 32767
        _ = v7
    default:
        var v8 int64 = 1
        _ = v8
    }
    var v9 int16 = 1000
    _ = v9
    return (v9 % ((v9) | 1))
}

func main() {
    v1 := make(map[string]string)
    var v2 int = -9223372036854775808
    v3 := make(map[int]bool)
    fmt.Println(v2)
    v10 := GenS1{F0: -362}
    v2 = v2
    for fi11 := 0; fi11 < 8; fi11++ {
        if (-128 > -103) {
            fmt.Println(fi11, fi11)
            var v12 int8 = 100
            _ = v12
        } else {
            switch (-(fi11)) {
            case 4294967295:
                v2 = (v2 + (^(v2)))
            }
        }
        v13 := v3[v2]
        _ = fi11
        _ = v13
    }
    v2 = v2
    v14 := GenS1{F0: 1}
    if true {
        var v15 int8 = 127
        _ = v15
    } else {
        var v16 float32 = 3.14159
        var v17 uint = 1
        v18 := []int{v2, v2}
        _ = v16
        _ = v17
        _ = v18
    }
    v19 := GenS1{F0: 42}
    if (float32(3.14159) > -2.25) {
        v2 = (v2 | 32767)
    } else {
        v2 -= genFn0(42, 0)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v14
    _ = v19
    _ = v2
    _ = v3
}
