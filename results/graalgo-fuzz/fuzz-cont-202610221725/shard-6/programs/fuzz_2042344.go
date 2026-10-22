package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 uint32, p1 uint32, p2 bool) int8 {
    for fi7 := 0; fi7 < 18; fi7++ {
        p0 = 2
        _ = fi7
    }
    p0 -= (^(p0))
    fmt.Println(p2)
    if false {
        fmt.Println(p1, p2, p1)
    }
    if p2 {
        v8 := []int{-2147483648, 128}
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1(p0 int32, p1 uint) bool {
    var v9 int32 = p0
    v9 += p0
    v10 := []int{}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return (!false)
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = 0.1
    v3 := make(map[string]bool)
    var v4 float64 = v2
    var v5 uint64 = 4294967295
    v6 := make(map[uint32]int)
    var v11 int8 = 127
    var v12 uint = 345
    var v13 uint64 = 2
    v13 = 164
    switch (("" + "") + "\n\t") {
    case "hello":
        fmt.Println(v13, v12, v12)
    case "\n\t":
        v14 := GenS1{F0: (v13 / ((4294967295) | 1))}
        _ = v14
    }
    var v15 string = "foo bar"
    v16 := GenS1{F0: v5}
    v17 := GenS1{F0: 1}
    v15 = (v15 + v15)
    var v18 uint64 = (v5 | (v5 >> (((^(v12))) & 63)))
    delete(v1, 1)
    if (2 >= 944) {
        var v19 bool = (genFn0(65535, 2, (1 >= 42)) <= v11)
        _ = len(v3)
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
