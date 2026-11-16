package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() uint8 {
    var v6 int8 = 37
    var v7 float32 = 1e10
    var v8 int = 1
    _ = v6
    _ = v7
    _ = v8
    return 2
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]int)
    v3 := []int{127}
    var v4 uint8 = 128
    v5 := make(map[uint32]uint32)
    v4 -= (v4 - v4)
    fmt.Println(v4)
    v4 = (^(v4))
    switch ("" + "a") {
    case "hello":
        v9 := GenS1{F0: 3}
        _ = v9
    case "a":
        var v10 uint8 = v4
        _ = v10
    default:
        delete(v2, 2)
    }
    v4 += 128
    v4 += v4
    if true {
        var v11 int8 = -128
        fmt.Println(v4, v4)
        _ = v11
    }
    if true {
        switch 479 {
        case 998:
            var v12 uint16 = 65535
            _ = v12
        case -377:
            v4 -= v4
        default:
            v13 := GenS1{F0: 100}
            _ = v13
        }
    } else {
        var v14 uint = 648
        var v15 uint64 = 444
        _ = v14
        _ = v15
    }
    v4 = 255
    v4 = v4
    fmt.Println(v4)
    v4 -= v4
    v4 += genFn0()
    if false {
        fmt.Println(v4)
    } else {
        _ = len(v2)
        v4 += (v4 >> ((4294967295) & 63))
    }
    v4 += v4
    if (v4 <= 255) {
        var v16 uint64 = 0
        v16 -= v16
        switch 32768 {
        case 127:
            v4 += v4
        default:
            if false {
                delete(v5, 65535)
            }
        }
        _ = v16
    }
    var v17 uint32 = 2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
