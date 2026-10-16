package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
func genFn0(p0 uint8, p1 int) bool {
    fmt.Println(p0, p0)
    if (((772 != 32767) && (0 < -32768)) && true) {
        p0 += p0
        for fi6 := 0; fi6 < 13; fi6++ {
            v7 := GenS1{F0: 65535, F1: 2}
            var v8 uint = 4294967295
            p0 = p0
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    _ = p0
    _ = p1
    return (true || (true || ((1e10 < float64(3.14159)) || true)))
}
func genFn1() int16 {
    var v9 uint64 = 699
    fmt.Println(v9)
    v10 := []int{}
    v9 = 2
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 int8 = 42
        _ = fi11
        _ = v12
    }
    _ = v10
    _ = v9
    return 805
}

func main() {
    var v1 float32 = 1.5
    var v2 bool = false
    var v3 float64 = 3.14159
    v4 := make(map[string]string)
    v5 := make(map[string]bool)
    switch "a" {
    case "hello":
        v13 := GenS1{F0: 42, F1: 899}
        _ = v13
    }
    v2 = (4294967295 >= 65535)
    v2 = v2
    fmt.Println(v1)
    var v14 uint8 = 127
    fmt.Println(v14, v2)
    for fi15 := 0; fi15 < 4; fi15++ {
        var v16 float32 = v1
        var v17 int8 = 0
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := GenS1{F0: 1, F1: 201}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
