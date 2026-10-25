package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 uint8, p1 bool, p2 int64) int16 {
    p1 = (2 >= 255)
    switch "hello" {
    case "\n\t":
        p1 = p1
    case "hello":
        for fi5 := 0; fi5 < 16; fi5++ {
            p2 += p2
            var v6 bool = p1
            _ = fi5
            _ = v6
        }
    case "a":
        if ((1 <= 2) || true) {
            v7 := []int{-654, -9223372036854775808, -425, 32767}
            _ = v7
        }
    default:
        var v8 int32 = 166
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return 32767
}

func main() {
    var v1 uint64 = 42
    var v2 int = -951
    v3 := []uint32{}
    var v4 uint = 42
    fmt.Println(v4, v2)
    fmt.Println(v1)
    for fi9 := 0; fi9 < 10; fi9++ {
        if len(v3) > 0 {
            v3[0] = 2147483648
        }
        v4 -= 0
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi9
    }
    var v10 string = ("foo bar" + "foo bar")
    var v11 uint = 255
    v12 := GenS1{F0: 42}
    v3 = append(v3, 2147483648)
    if len(v3) > 0 {
        v3[0] = 2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
