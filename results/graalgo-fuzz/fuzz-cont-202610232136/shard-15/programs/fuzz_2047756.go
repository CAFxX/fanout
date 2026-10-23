package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 int16, p1 uint64, p2 int64) string {
    var v4 int64 = (p2 % (((p2 << ((18446744073709551615) & 63))) | 1))
    p1 = p1
    for fi5 := 0; fi5 < 15; fi5++ {
        v4 += p2
        v6 := GenS1{F0: (p1 ^ p1)}
        _ = fi5
        _ = v6
    }
    v7 := []int{}
    var v8 uint = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v7
    _ = v8
    return ""
}
func genFn1(p0 int8, p1 uint32) int16 {
    if true {
        v9 := []int{913, 0, -231}
        v9 = append(v9, 127)
        var v10 int16 = -637
        _ = v10
        _ = v9
    }
    var v11 uint64 = 4294967295
    var v12 float32 = 0.1
    v13 := GenS1{F0: (v11 << ((250) & 63))}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    return 514
}

func main() {
    var v1 float32 = (float32(float64(0.1)) / 1.5)
    v2 := make(map[uint32]uint32)
    v3 := []string{"hello", "\n\t"}
    var v14 uint8 = 169
    for fi15 := 0; fi15 < 15; fi15++ {
        _ = len(v3)
        _ = fi15
    }
    if (1 > 0) {
        v16 := GenS1{F0: 42}
        _ = len(v2)
        _ = v16
    } else {
        var v17 uint32 = 6
        switch (("hello" + ("\n\t" + (("\n\t" + "a") + ("hello" + "foo bar")))) + "hello") {
        case "\n\t":
            fmt.Println(v1, v17)
        }
        fmt.Println(v17, v17)
        _ = v17
    }
    v3 = append(v3, "hello")
    v14 = (v14 + 120)
    v18 := GenS1{F0: 2}
    switch ("hello" + "hello") {
    case "a":
        v19 := GenS1{F0: 576}
        _ = v19
    }
    var v20 int32 = -1
    v2[2147483647] = 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v20
    _ = v3
}
