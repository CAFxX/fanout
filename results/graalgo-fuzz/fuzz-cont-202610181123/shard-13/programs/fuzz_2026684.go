package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 uint = 18446744073709551615
    var v2 int8 = -121
    v3 := make(map[uint32]int)
    var v4 uint = (v1 * (v1 % (((v1 / ((v1) | 1))) | 1)))
    for fi5 := 0; fi5 < 6; fi5++ {
        fmt.Println(v1, v1)
        var v6 bool = (false || (42 < 2))
        var v7 bool = false
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 int16 = 1000
    v1 -= 18446744073709551615
    v3[288] = -1
    fmt.Println(v2)
    v2 -= (v2 % ((v2) | 1))
    fmt.Println(v2, v2, v4)
    var v9 int32 = -2147483648
    if false {
        v1 = (v1 >> (((^(v4))) & 63))
        v10 := GenS1{F0: 128}
        _ = v10
    }
    fmt.Println(v4, v2, v9)
    var v11 bool = false
    for fi12 := 0; fi12 < 10; fi12++ {
        for fi13 := 0; fi13 < 14; fi13++ {
            delete(v3, 0)
            _ = fi13
        }
        switch ("\n\t" + "foo bar") {
        case "foo bar":
            switch fi12 {
            case 4294967295:
                delete(v3, 2147483648)
            default:
                var v14 string = "a"
                _ = v14
            }
        }
        v3[2147483648] = (fi12 + fi12)
        _ = fi12
    }
    v15 := []int{127, 32768, 0}
    var v16 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
