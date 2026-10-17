package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
func genFn0() uint64 {
    var v5 int8 = 127
    v6 := GenS1{F0: 42}
    _ = v5
    _ = v6
    return 812
}
func genFn1() uint32 {
    var v7 float64 = 0.1
    v8 := []int{}
    v9 := GenS1{F0: 1}
    _ = len(v8)
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 226
    var v3 int = -103
    var v4 uint = uint(v2)
    v2 = 18446744073709551615
    v10 := []int{9223372036854775807}
    var v11 uint8 = 147
    v11 = v11
    v4 -= v4
    var v12 bool = false
    switch ("\n\t" + "\n\t") {
    case "a":
        var v13 bool = v12
        _ = v13
    case "foo bar":
        var v14 bool = (241 == 32767)
        _ = v14
    }
    var v15 int = 0
    v10 = append(v10, -9223372036854775808)
    var v16 bool = (!false)
    fmt.Println(v12, v4)
    v15 = -1
    for fi17 := 0; fi17 < 20; fi17++ {
        var v18 int = (^(v3))
        fmt.Println(v16, v3, v16)
        v1[2] = (!(float64(3.14159) <= -2.25))
        _ = fi17
        _ = v18
    }
    v19 := GenS2{F0: -366}
    switch (v3 / ((v3) | 1)) {
    case 1:
        var v20 uint16 = 0
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
