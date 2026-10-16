package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 bool = ((-2.25 != 0.0) && false)
    v2 := make(map[uint32]int)
    var v3 int16 = 0
    fmt.Println(v1, v1)
    v1 = ((-1 >= 9223372036854775807) && (127 < -128))
    for fi4 := 0; fi4 < 7; fi4++ {
        for fi5 := 0; fi5 < 1; fi5++ {
            _ = len(v2)
            _ = fi5
        }
        var v6 uint32 = 943
        switch 32768 {
        case 128:
            if true {
                for fi7 := 0; fi7 < 7; fi7++ {
                    v6 -= v6
                    v8 := GenS1{F0: true}
                    _ = fi7
                    _ = v8
                }
            }
        case 42:
            var v9 uint = 1
            _ = v9
        }
        _ = fi4
        _ = v6
    }
    v1 = false
    switch -9223372036854775808 {
    case -354:
        var v10 int16 = 0
        _ = v10
    default:
        var v11 int = 9223372036854775807
        _ = v11
    }
    fmt.Println(v3, v3)
    if (-361 < -2147483648) {
        v1 = (!false)
    }
    var v12 float32 = -2.25
    v12 -= v12
    v13 := []int{0, int(v3)}
    v3 -= 1
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
