package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 uint32 = 2147483647
    v2 := []uint32{}
    v3 := []int{958, 32767}
    v4 := make(map[uint32]string)
    v5 := []int{}
    var v6 int32 = -677
    var v7 float32 = -2.25
    if ("a" == "\n\t") {
        v6 = v6
    }
    switch -9223372036854775808 {
    case 128:
        v7 = v7
    case 42:
        var v8 uint64 = 2
        _ = v8
    default:
        var v9 int = 189
        _ = v9
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 bool = false
        v7 = ((1.5 * v7) / v7)
        _ = fi10
        _ = v11
    }
    fmt.Println(v6, v1)
    for fi12 := 0; fi12 < 15; fi12++ {
        v13 := v4[(v1 * v1)]
        var v14 uint16 = 65535
        var v15 string = v13
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    delete(v4, 0)
    fmt.Println(v7, v7)
    var v16 string = "a"
    fmt.Println(v16)
    var v17 float64 = 0.1
    fmt.Println(v16)
    var v18 int16 = -1
    v18 -= -1
    delete(v4, v1)
    if true {
        if false {
            v19 := GenS1{F0: 65535}
            if true {
                var v20 int16 = -676
                fmt.Println(v20, v18, v18)
                _ = v20
            } else {
                v6 += (v6 - 433)
                var v21 uint = 255
                v1 = (v1 / ((2147483647) | 1))
                _ = v21
            }
            fmt.Println(v18, v7, v16)
            _ = v19
        } else {
            delete(v4, v1)
        }
        switch v16 {
        case "foo bar":
            if len(v2) > 0 {
                v2[0] = 0
            }
        }
        for fi22 := 0; fi22 < 4; fi22++ {
            v1 -= v1
            switch "a" {
            case "\n\t":
                var v23 bool = (fi22 < 128)
                _ = v23
            case "hello":
                v18 += (v18 >> ((672) & 63))
            case "a":
                var v24 int32 = -2147483648
                _ = v24
            }
            var v25 float32 = -2.25
            _ = fi22
            _ = v25
        }
    } else {
        v7 -= v7
        var v26 uint8 = 219
        var v27 bool = false
        _ = v26
        _ = v27
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
