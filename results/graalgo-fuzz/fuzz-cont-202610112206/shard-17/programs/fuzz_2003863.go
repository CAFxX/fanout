package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 int16 = 32767
    var v2 float32 = float32(1.5)
    v3 := []uint32{65535}
    v4 := make(map[uint32]string)
    var v5 uint16 = 19
    var v6 bool = false
    for fi7 := 0; fi7 < 8; fi7++ {
        v2 = 0.1
        _ = fi7
    }
    for fi8 := 0; fi8 < 10; fi8++ {
        switch fi8 {
        case 32767:
            delete(v4, 4294967295)
        }
        for fi9 := 0; fi9 < 2; fi9++ {
            delete(v4, 65535)
            if (-455 >= -603) {
                v6 = ((fi9 / ((fi9) | 1)) == -9223372036854775808)
                var v10 float64 = (3.14159 * 1e10)
                if (!(1 != 856)) {
                    v11 := v4[4294967295]
                    _ = v11
                } else {
                    var v12 float32 = float32((v10 * v10))
                    _ = v12
                }
                _ = len(v4)
                _ = v10
            }
            if ((v1 * -32768) > -411) {
                var v13 int = (fi8 ^ 128)
                _ = v13
            }
            _ = fi9
        }
        v14 := GenS1{F0: 42}
        _ = fi8
        _ = v14
    }
    switch "foo bar" {
    case "hello":
        var v15 uint = 65535
        _ = v15
    case "foo bar":
        fmt.Println(v5, v6)
    case "":
        var v16 int8 = 5
        _ = v16
    }
    var v17 int64 = -9223372036854775808
    var v18 uint64 = 2
    v3 = append(v3, 2)
    v4[65535] = ("a" + "foo bar")
    fmt.Println(v6, v18, v17)
    v5 += (v5 - v5)
    for fi19 := 0; fi19 < 7; fi19++ {
        var v20 uint64 = (v18 + v18)
        _ = fi19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
