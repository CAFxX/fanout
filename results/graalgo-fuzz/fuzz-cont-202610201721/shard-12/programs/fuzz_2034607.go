package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 int64 = 0
    var v2 bool = (!true)
    v3 := make(map[string]int)
    v4 := make(map[uint32]int)
    var v5 int64 = v1
    if v2 {
        v5 += (v1 - (v5 & v1))
    } else {
        fmt.Println(v2)
        v1 -= v5
        for fi6 := 0; fi6 < 7; fi6++ {
            v7 := GenS1{F0: 320}
            fmt.Println(v1)
            var v8 int = fi6
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 int8 = 0
        if (18446744073709551615 == 42) {
            v11 := GenS1{F0: 42}
            v12 := GenS1{F0: 32767}
            for fi13 := 0; fi13 < 10; fi13++ {
                v14 := GenS1{F0: 309}
                var v15 int16 = -1
                v2 = v2
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = v11
            _ = v12
        }
        v10 += v10
        _ = fi9
        _ = v10
    }
    switch (("\n\t" + "a") + ("a" + "")) {
    case "a":
        var v16 float64 = (3.14159 + 0.0)
        _ = v16
    case "\n\t":
        var v17 string = "hello"
        _ = v17
    default:
        for fi18 := 0; fi18 < 19; fi18++ {
            var v19 uint16 = 65535
            if v2 {
                var v20 int8 = -55
                _ = v20
            }
            _ = fi18
            _ = v19
        }
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        v5 -= v5
        _ = fi21
    }
    var v22 int = 32768
    for fi23 := 0; fi23 < 4; fi23++ {
        v24 := []int{32767}
        _ = fi23
        _ = v24
    }
    var v25 uint8 = 0
    v25 += 255
    if (42 < 1) {
        var v26 float64 = 0.0
        _ = v26
    }
    fmt.Println(v1, v1)
    if v2 {
        var v27 uint64 = 2
        _ = v27
    }
    if v2 {
        v5 = v1
    }
    v28 := []int{v22}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
