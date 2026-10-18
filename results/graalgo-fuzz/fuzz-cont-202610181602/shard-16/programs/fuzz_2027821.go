package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 bool }

func main() {
    var v1 string = ""
    var v2 uint8 = 127
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]int)
    var v5 int32 = -2147483648
    var v6 string = v1
    if true {
        var v7 int8 = 42
        _ = v7
    }
    fmt.Println(v5, v5)
    v6 = (v6 + v6)
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 int = fi8
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: v2}
    if false {
        if true {
            v5 -= (v5 >> ((724) & 63))
        }
        v11 := GenS2{F0: false, F1: 42}
        _ = v11
    }
    v12 := GenS2{F0: (3.14159 != 1.5), F1: 1}
    v13 := GenS3{F0: true}
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 float64 = (1.5 * (3.14159 * (-2.25 + -2.25)))
        _ = fi14
        _ = v15
    }
    var v16 string = ((((v6 + ("a" + "foo bar")) + v1) + ("" + "a")) + ("\n\t" + v6))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
