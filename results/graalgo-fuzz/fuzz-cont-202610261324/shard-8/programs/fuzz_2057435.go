package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 bool = (1 == 817)
    v2 := []int{}
    v3 := make(map[uint32]int)
    v4 := make(map[int]bool)
    v5 := []int{1, 9223372036854775807}
    var v6 int8 = -40
    v7 := GenS1{F0: 128}
    var v8 string = ("a" + ("a" + "a"))
    var v9 uint16 = 255
    fmt.Println(v1, v6)
    _ = len(v4)
    var v10 bool = true
    for fi11 := 0; fi11 < 6; fi11++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v6 = (v6 | (^(v6)))
        fmt.Println(v9, v9, v9)
        _ = fi11
    }
    switch 1 {
    case -1:
        for fi12 := 0; fi12 < 10; fi12++ {
            if false {
                fmt.Println(v1, v10)
                v13 := GenS1{F0: 0}
                if v10 {
                    switch "a" {
                    case "a":
                        v8 = "a"
                    }
                }
                _ = v13
            } else {
                fmt.Println(fi12)
                for fi14 := 0; fi14 < 7; fi14++ {
                    v8 = (v8 + "hello")
                    var v15 int = (fi14 >> ((65535) & 63))
                    _ = fi14
                    _ = v15
                }
            }
            v16 := GenS1{F0: 2}
            _ = fi12
            _ = v16
        }
    default:
        var v17 int32 = 732
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
