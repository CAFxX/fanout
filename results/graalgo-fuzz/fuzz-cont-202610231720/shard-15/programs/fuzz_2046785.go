package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int64 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 int = 0
    var v2 float32 = -2.25
    v3 := make(map[string]int)
    v4 := make(map[string]uint32)
    v1 = (v1 + v1)
    if false {
        switch -2147483648 {
        case -647:
            var v5 float64 = 1e10
            _ = v5
        case 1:
            v6 := GenS1{F0: 0, F1: 0, F2: -1}
            _ = v6
        }
        if true {
            v2 = 3.14159
            if (true || (0 > 427)) {
                var v7 uint8 = 2
                _ = v7
            } else {
                for fi8 := 0; fi8 < 10; fi8++ {
                    var v9 uint16 = 1
                    if true {
                        fmt.Println(fi8)
                        v10 := GenS2{F0: -759, F1: 2147483647}
                        _ = v10
                    }
                    _ = fi8
                    _ = v9
                }
                v1 -= v1
            }
        } else {
            v11 := []int{v1, v1, (v1 - (v1 * v1)), (v1 - v1)}
            fmt.Println(v1)
            v11 = append(v11, v1)
            _ = v11
        }
        var v12 int8 = 42
        _ = v12
    } else {
        for fi13 := 0; fi13 < 3; fi13++ {
            var v14 uint16 = 0
            v1 -= v1
            v15 := []int{(v1 & (v1 >> ((546) & 63))), v1, (-(v1))}
            _ = fi13
            _ = v14
            _ = v15
        }
        var v16 uint64 = 927
        if true {
            v17 := GenS1{F0: 539, F1: (v16 & 42), F2: 9223372036854775807}
            v1 = v1
            v18 := GenS2{F0: -856, F1: -1}
            _ = v17
            _ = v18
        }
        _ = v16
    }
    var v19 float64 = 3.14159
    v2 -= 1.5
    switch "a" {
    case "hello":
        v20 := GenS3{F0: 127}
        _ = v20
    }
    v21 := GenS1{F0: 42, F1: 911, F2: -652}
    var v22 uint8 = 42
    for fi23 := 0; fi23 < 12; fi23++ {
        var v24 int = -2147483648
        var v25 uint = 0
        var v26 string = ""
        _ = fi23
        _ = v24
        _ = v25
        _ = v26
    }
    if true {
        switch (v1 * (v1 + v1)) {
        case 32768:
            fmt.Println(v22, v2, v2)
        case -181:
            v27 := GenS1{F0: 65535, F1: 42, F2: 9223372036854775807}
            _ = v27
        default:
            if false {
                v19 -= (v19 + v19)
                var v28 int8 = 100
                _ = v28
            } else {
                _ = len(v4)
            }
        }
        var v29 uint8 = 42
        _ = v29
    }
    switch (v1 >> ((780) & 63)) {
    case -2147483648:
        var v30 uint16 = 255
        _ = v30
    case 127:
        v31 := GenS3{F0: 0}
        _ = v31
    case 4294967295:
        v2 += 1e10
    }
    var v32 int16 = 32767
    fmt.Println(v19, v2, v19)
    fmt.Println(v19)
    v33 := GenS1{F0: uint16(v22), F1: 759, F2: 0}
    v32 = (v32 | v32)
    v32 += 42
    v34 := []int{int(v32), 4294967295, (-(v1)), (-(v1))}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
