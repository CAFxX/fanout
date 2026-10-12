package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 uint8 = 0
    var v2 int8 = 96
    var v3 uint = 0
    var v4 int16 = -32768
    var v5 uint32 = 1
    var v6 uint8 = v1
    v2 = 42
    var v7 string = ("" + "\n\t")
    for fi8 := 0; fi8 < 3; fi8++ {
        if true {
            v9 := []int{}
            _ = v9
        } else {
            v10 := GenS1{F0: 0}
            _ = v10
        }
        v2 = v2
        _ = fi8
    }
    var v11 uint32 = (v5 ^ (v5 | v5))
    v4 = (v4 << (((^(v3))) & 63))
    v5 += v5
    for fi12 := 0; fi12 < 19; fi12++ {
        switch fi12 {
        case -346:
            switch 128 {
            case -2147483648:
                var v13 string = v7
                _ = v13
            case -1:
                var v14 float32 = 1.5
                _ = v14
            }
        case -660:
            v4 += (v4 | (-(v4)))
        case 127:
            v15 := GenS1{F0: -2147483648}
            _ = v15
        default:
            if true {
                switch fi12 {
                case 42:
                    var v16 uint32 = (v5 - v11)
                    _ = v16
                case -91:
                    v17 := []int{128}
                    _ = v17
                case 128:
                    switch fi12 {
                    case 32768:
                        fmt.Println(v11, fi12, v1)
                    case 2147483647:
                        var v18 uint32 = 662
                        _ = v18
                    }
                }
                v4 -= (v4 ^ (-(v4)))
            }
        }
        _ = fi12
    }
    var v19 bool = (339 == 870)
    fmt.Println(v5, v11)
    var v20 int8 = v2
    v21 := []int{476, 32768, 32768}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
