package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0() int16 {
    v4 := GenS1{F0: 0}
    var v5 bool = true
    v5 = v5
    _ = v4
    _ = v5
    return 282
}
func genFn1(p0 int64, p1 int16) int64 {
    v6 := []int{4294967295, 2147483647}
    if true {
        var v7 int8 = -58
        _ = v7
    }
    p1 = (p1 ^ p1)
    var v8 uint32 = 4294967295
    _ = p0
    _ = p1
    _ = v6
    _ = v8
    return p0
}

func main() {
    var v1 int = -1
    v2 := make(map[uint32]bool)
    var v3 float64 = (0.0 / 1.0)
    for fi9 := 0; fi9 < 9; fi9++ {
        _ = len(v2)
        _ = fi9
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        v3 -= float64(float32((v3 + v3)))
        var v11 int8 = 104
        _ = fi10
        _ = v11
    }
    if (!(3.14159 <= 1e10)) {
        var v12 int8 = 40
        _ = len(v2)
        _ = v12
    } else {
        var v13 string = "hello"
        v14 := GenS1{F0: 32767}
        if (1 == 2) {
            for fi15 := 0; fi15 < 20; fi15++ {
                var v16 uint16 = 32767
                var v17 int16 = 42
                _ = fi15
                _ = v16
                _ = v17
            }
        } else {
            var v18 uint64 = 42
            switch v13 {
            case "":
                var v19 uint8 = 2
                _ = v19
            }
            v1 = v1
            _ = v18
        }
        _ = v13
        _ = v14
    }
    fmt.Println(v1, v1)
    for fi20 := 0; fi20 < 10; fi20++ {
        v3 = 1.5
        v21 := GenS1{F0: -752}
        _ = fi20
        _ = v21
    }
    v1 += 32768
    var v22 string = "a"
    var v23 int64 = 9223372036854775807
    var v24 uint = 2
    v1 = v1
    switch v1 {
    case -2147483648:
        v24 += 2
    case 32767:
        v22 = v22
    case 42:
        fmt.Println(v1, v1, v23)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
