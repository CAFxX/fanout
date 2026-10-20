package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 bool = false
    v2 := make(map[string]int)
    var v3 int32 = -2147483648
    var v4 uint16 = 0
    var v5 string = (("\n\t" + "\n\t") + "a")
    var v6 uint32 = 393
    var v7 string = "\n\t"
    if (!v1) {
        var v8 int64 = -9223372036854775808
        _ = v8
    }
    v9 := GenS1{F0: v3}
    v10 := []int{-533, 32768}
    v1 = true
    for fi11 := 0; fi11 < 4; fi11++ {
        fmt.Println(v4, v4)
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 float32 = -2.25
            v3 -= (v3 / (((-(v3))) | 1))
            var v14 float64 = (float64((v13 + v13)) - 1e10)
            _ = fi12
            _ = v13
            _ = v14
        }
        v15 := GenS1{F0: 42}
        _ = fi11
        _ = v15
    }
    var v16 string = "a"
    for fi17 := 0; fi17 < 20; fi17++ {
        v2[("foo bar" + v16)] = int(v6)
        switch "foo bar" {
        case "hello":
            v18 := GenS1{F0: 0}
            _ = v18
        case "foo bar":
            var v19 bool = (!(784 > 42))
            _ = v19
        }
        _ = fi17
    }
    var v20 int = 9223372036854775807
    v21 := GenS1{F0: 42}
    v2["\n\t"] = v20
    var v22 uint = 823
    if v1 {
        var v23 int32 = v3
        v24 := v2[v5]
        _ = v23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
