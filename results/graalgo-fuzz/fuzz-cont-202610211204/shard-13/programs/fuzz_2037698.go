package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int = 1
    var v3 uint32 = 65535
    var v4 int32 = 0
    var v5 int = (v2 << ((255) & 63))
    v1[-632] = v3
    var v6 string = ("foo bar" + ("\n\t" + "hello"))
    switch v6 {
    case "foo bar":
        if false {
            v7 := GenS1{F0: -837}
            _ = v7
        } else {
            fmt.Println(v4, v6, v5)
        }
    default:
        switch "" {
        case "":
            var v8 uint16 = 591
            _ = v8
        case "foo bar":
            var v9 uint = 689
            _ = v9
        case "hello":
            if false {
                for fi10 := 0; fi10 < 17; fi10++ {
                    v11 := []int{}
                    _ = fi10
                    _ = v11
                }
            } else {
                v2 -= 1
            }
        }
    }
    var v12 uint8 = 0
    var v13 int8 = 1
    var v14 float32 = 0.1
    _ = len(v1)
    var v15 bool = false
    if false {
        fmt.Println(v6, v12, v3)
        v16 := []int{v5}
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
