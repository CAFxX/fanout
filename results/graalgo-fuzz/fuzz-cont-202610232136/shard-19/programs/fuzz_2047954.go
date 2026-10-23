package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint32 = 4294967295
    v3 := []uint32{}
    v4 := make(map[string]uint32)
    v5 := []int{}
    fmt.Println(v2)
    _ = len(v4)
    var v6 uint = 1
    delete(v4, "hello")
    var v7 string = "a"
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 float32 = 0.1
        _ = fi8
        _ = v9
    }
    _ = len(v1)
    fmt.Println(v6)
    v7 = v7
    if true {
        for fi10 := 0; fi10 < 8; fi10++ {
            switch (fi10 >> ((60) & 63)) {
            case 0:
                for fi11 := 0; fi11 < 9; fi11++ {
                    var v12 int64 = -1
                    var v13 int64 = v12
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            default:
                var v14 int16 = 1
                _ = v14
            }
            v6 -= v6
            _ = fi10
        }
        if (1.5 < (float32((-2.25 / 1.5)) / -2.25)) {
            v2 -= v2
            var v15 int8 = 100
            var v16 bool = false
            _ = v15
            _ = v16
        } else {
            for fi17 := 0; fi17 < 7; fi17++ {
                v7 = "hello"
                v18 := GenS1{F0: 1, F1: 357, F2: -523}
                _ = fi17
                _ = v18
            }
        }
        switch (v7 + v7) {
        case "hello":
            var v19 int32 = 2147483647
            _ = v19
        }
    }
    fmt.Println(v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
