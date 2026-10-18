package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint8; F1 bool; F2 uint32 }

func main() {
    v1 := []uint32{}
    v2 := make(map[int]int)
    var v3 string = ""
    var v4 int = 9223372036854775807
    v5 := v2[-242]
    v4 += v4
    switch ("hello" + "\n\t") {
    case "a":
        v4 = (^(v4))
    case "foo bar":
        var v6 uint16 = 42
        _ = v6
    }
    var v7 uint8 = 42
    for fi8 := 0; fi8 < 15; fi8++ {
        v5 += v4
        _ = fi8
    }
    v9 := GenS1{F0: v3, F1: v5}
    v7 += v7
    for fi10 := 0; fi10 < 19; fi10++ {
        v5 = -2147483648
        if (!(2 != 255)) {
            v11 := GenS1{F0: v3, F1: (-(v4))}
            if (true || true) {
                switch v3 {
                case "foo bar":
                    var v12 int64 = 42
                    _ = v12
                case "hello":
                    _ = len(v1)
                }
            }
            _ = v11
        }
        _ = fi10
    }
    v13 := GenS2{F0: (^(v7)), F1: false, F2: 4294967295}
    var v14 string = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
