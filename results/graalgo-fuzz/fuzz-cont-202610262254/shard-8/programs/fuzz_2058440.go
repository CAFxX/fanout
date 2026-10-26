package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int32, p1 int8, p2 string) int8 {
    p0 = p0
    if false {
        var v5 int16 = 32767
        _ = v5
    }
    p1 -= p1
    v6 := []int{0, -698}
    if false {
        p0 += 42
        if false {
            var v7 int64 = 9223372036854775807
            _ = v7
        }
        p1 += p1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 42
}

func main() {
    v1 := []string{"hello", "foo bar", "hello"}
    var v2 uint64 = 626
    var v3 string = "foo bar"
    v4 := []string{("foo bar" + ""), ((("hello" + ("hello" + v3)) + v3) + v3), "\n\t"}
    fmt.Println(v2, v2)
    for fi8 := 0; fi8 < 16; fi8++ {
        v3 = v3
        v2 -= 2
        _ = fi8
    }
    v3 = ""
    var v9 bool = (v3 <= v3)
    var v10 float64 = ((0.1 + 1.5) * 1e10)
    var v11 uint32 = 42
    var v12 int = 32768
    v13 := GenS1{F0: 744}
    var v14 uint32 = v11
    switch v3 {
    case "hello":
        var v15 int = (v12 / ((-1) | 1))
        _ = v15
    case "foo bar":
        v12 = (v12 & (v12 & 32767))
    case "\n\t":
        var v16 uint64 = 425
        _ = v16
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
