package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 uint16, p1 uint32, p2 int32) uint16 {
    var v4 uint8 = 128
    if false {
        v5 := GenS1{F0: false}
        _ = v5
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (p0 & p0)
}

func main() {
    var v1 uint64 = 0
    var v2 uint32 = 2147483648
    var v3 uint16 = 32767
    var v6 uint16 = (v3 ^ (v3 % ((255) | 1)))
    var v7 uint32 = (v2 & 0)
    v3 += 65535
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := GenS1{F0: false}
        _ = fi8
        _ = v9
    }
    v10 := []int{830, 2147483647}
    if (!(int64(v7) == 153)) {
        switch "" {
        case "hello":
            v11 := GenS1{F0: (("hello" + "hello") == "a")}
            _ = v11
        default:
            v1 -= (v1 % ((v1) | 1))
        }
        for fi12 := 0; fi12 < 10; fi12++ {
            switch (("hello" + ((("foo bar" + ("hello" + "foo bar")) + "foo bar") + "\n\t")) + "foo bar") {
            case "":
                v2 += (v2 & (v7 ^ v7))
            default:
                v1 = (v1 ^ v1)
            }
            fmt.Println(v6, fi12, v7)
            _ = fi12
        }
        var v13 int32 = -621
        _ = v13
    }
    fmt.Println(v7, v6)
    v1 += (^(v1))
    var v14 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
