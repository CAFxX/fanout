package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string }

func main() {
    var v1 int8 = 42
    v2 := make(map[int]int)
    v3 := make(map[int]int)
    v4 := make(map[int]bool)
    v5 := make(map[uint32]uint32)
    switch "" {
    case "hello":
        v6 := GenS1{F0: 100}
        _ = v6
    default:
        if true {
            fmt.Println(v1)
        }
    }
    if (1.5 > (0.0 * -2.25)) {
        var v7 int = -9223372036854775808
        v8 := GenS2{F0: ("foo bar" + "a")}
        _ = v7
        _ = v8
    }
    v9 := GenS2{F0: ""}
    for fi10 := 0; fi10 < 14; fi10++ {
        v1 += v1
        v1 += (v1 | (v1 - (v1 + 0)))
        v11 := []int{fi10, (fi10 | 0)}
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        switch ("" + "a") {
        case "hello":
            v3[fi12] = (fi12 + -704)
        case "\n\t":
            v1 = (v1 | v1)
        default:
            switch fi12 {
            case -1:
                if false {
                    v1 = v1
                    var v13 string = "a"
                    var v14 int16 = -570
                    _ = v13
                    _ = v14
                } else {
                    fmt.Println(fi12, fi12)
                }
            default:
                for fi15 := 0; fi15 < 15; fi15++ {
                    v16 := []int{}
                    v17 := v3[-9223372036854775808]
                    var v18 int = 4294967295
                    _ = fi15
                    _ = v16
                    _ = v17
                    _ = v18
                }
            }
        }
        _ = fi12
    }
    v1 = 42
    var v19 int = -790
    for fi20 := 0; fi20 < 2; fi20++ {
        var v21 uint16 = 758
        v19 = fi20
        _ = fi20
        _ = v21
    }
    v1 += 1
    fmt.Println(v19, v19)
    v1 = (v1 | v1)
    v19 = (v19 >> ((42) & 63))
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
