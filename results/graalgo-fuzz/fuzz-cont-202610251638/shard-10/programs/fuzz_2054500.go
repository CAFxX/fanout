package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint16, p1 int8, p2 int8) int {
    v4 := []int{128, 127}
    v5 := GenS1{F0: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -9223372036854775808
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = -2147483648
    var v3 uint16 = 615
    fmt.Println(v3, v3)
    v1[488] = 65535
    for fi6 := 0; fi6 < 4; fi6++ {
        switch "" {
        case "hello":
            v1[0] = 0
        }
        v3 = v3
        v7 := GenS1{F0: int64(v3)}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        fmt.Println(v2, v2, fi8)
        v9 := GenS1{F0: -1}
        _ = fi8
        _ = v9
    }
    if (!(false || false)) {
        v10 := []int{}
        _ = v10
    }
    v2 = v2
    v3 = (v3 / ((v3) | 1))
    var v11 int8 = 110
    fmt.Println(v3, v11)
    for fi12 := 0; fi12 < 14; fi12++ {
        v2 += -710
        _ = fi12
    }
    var v13 uint8 = 127
    var v14 uint32 = 0
    var v15 uint16 = (v3 >> ((65535) & 63))
    switch "foo bar" {
    case "hello":
        v16 := []int{9223372036854775807, 2147483647, 127, -995}
        _ = v16
    case "":
        v17 := GenS1{F0: 1}
        _ = v17
    case "a":
        var v18 uint32 = 459
        _ = v18
    }
    switch "foo bar" {
    case "foo bar":
        v15 -= v15
    case "":
        v13 -= (v13 | uint8(v14))
    case "\n\t":
        v14 = (v14 / (((v14 % ((440) | 1))) | 1))
    default:
        v1[(^(v14))] = (v14 - (v14 + (v14 / ((601) | 1))))
    }
    switch "" {
    case "hello":
        switch 0 {
        case 0:
            var v19 string = "a"
            _ = v19
        }
    }
    var v20 string = ("\n\t" + "")
    switch 127 {
    case 324:
        fmt.Println(v2, v3)
    case 807:
        var v21 int8 = (v11 * v11)
        _ = v21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
}
