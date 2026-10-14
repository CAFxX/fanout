package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint; F1 int8 }
func genFn0(p0 string, p1 int, p2 int) uint {
    var v6 uint64 = 2
    p0 = (p0 + p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 816
}

func main() {
    v1 := []string{}
    v2 := make(map[uint32]string)
    var v3 uint = 784
    var v4 int16 = -994
    var v5 uint8 = 255
    var v7 uint32 = 357
    var v8 bool = true
    switch "a" {
    case "a":
        v2[v7] = "\n\t"
    default:
        v3 = (v3 & v3)
    }
    var v9 uint = (^(v3))
    v4 = (v4 % ((v4) | 1))
    var v10 string = ("\n\t" + "foo bar")
    var v11 uint64 = 4294967295
    switch 42 {
    case 1:
        for fi12 := 0; fi12 < 6; fi12++ {
            if (fi12 != fi12) {
                if (!true) {
                    v13 := GenS1{F0: v10, F1: true}
                    _ = v13
                }
                switch ("a" + v10) {
                case "foo bar":
                    var v14 uint64 = v11
                    _ = v14
                case "hello":
                    v15 := GenS2{F0: 1, F1: -81}
                    _ = v15
                case "":
                    v9 -= v9
                }
            }
            _ = fi12
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
