package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 uint, p1 int, p2 uint) uint {
    if (-1 != -1) {
        v7 := []int{}
        _ = v7
    }
    v8 := []int{32767, p1}
    var v9 uint64 = 17
    v8 = append(v8, -684)
    if (p1 != 2147483647) {
        var v10 int8 = 42
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return (^(p2))
}

func main() {
    var v1 uint8 = 218
    v2 := []uint32{42, 42, 2}
    var v3 uint = 4294967295
    var v4 int64 = 1
    v5 := make(map[string]string)
    v6 := make(map[string]string)
    switch "foo bar" {
    case "foo bar":
        _ = len(v2)
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 float64 = 3.14159
        v1 += 90
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        fmt.Println(fi13, v4, v4)
        var v14 int64 = v4
        _ = fi13
        _ = v14
    }
    v1 = v1
    if false {
        v4 -= v4
    } else {
        fmt.Println(v4, v1)
        switch "" {
        case "\n\t":
            fmt.Println(v1, v1)
        case "hello":
            fmt.Println(v4)
        default:
            var v15 uint16 = 677
            _ = v15
        }
    }
    var v16 string = (("" + ("" + ("\n\t" + "a"))) + "\n\t")
    v16 = v16
    if true {
        for fi17 := 0; fi17 < 15; fi17++ {
            for fi18 := 0; fi18 < 13; fi18++ {
                v2 = append(v2, 2147483647)
                v19 := GenS1{F0: uint32(v1)}
                _ = fi18
                _ = v19
            }
            v4 += v4
            switch "a" {
            case "a":
                var v20 uint8 = 1
                _ = v20
            case "hello":
                for fi21 := 0; fi21 < 9; fi21++ {
                    var v22 uint = (^(v3))
                    _ = len(v6)
                    var v23 int32 = 500
                    var v24 int64 = v4
                    var v25 string = v16
                    _ = fi21
                    _ = v22
                    _ = v23
                    _ = v24
                    _ = v25
                }
            case "\n\t":
                for fi26 := 0; fi26 < 6; fi26++ {
                    fmt.Println(v16, v16)
                    _ = fi26
                }
            default:
                v1 -= (v1 >> (((v3 + v3)) & 63))
            }
            _ = fi17
        }
    }
    v27 := v5["a"]
    v27 = v27
    v4 = v4
    var v28 int32 = -973
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
