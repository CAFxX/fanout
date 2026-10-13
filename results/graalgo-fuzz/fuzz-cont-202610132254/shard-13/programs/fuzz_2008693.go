package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0() uint8 {
    switch (("\n\t" + "a") + "a") {
    case "\n\t":
        for fi7 := 0; fi7 < 18; fi7++ {
            var v8 int32 = 2147483647
            _ = fi7
            _ = v8
        }
    case "a":
        var v9 string = (("" + ((("\n\t" + "a") + "\n\t") + "a")) + "\n\t")
        _ = v9
    }
    v10 := []int{-356, 4294967295, 0}
    switch ("hello" + ("hello" + "")) {
    case "foo bar":
        var v11 uint64 = 2
        _ = v11
    default:
        var v12 int16 = 0
        _ = v12
    }
    var v13 bool = false
    _ = v10
    _ = v13
    return 42
}

func main() {
    var v1 int16 = 32767
    var v2 int8 = 1
    v3 := make(map[string]uint32)
    var v4 float32 = float32(3.14159)
    v5 := make(map[uint32]bool)
    var v6 int = -251
    v14 := v5[2]
    v15 := v3["foo bar"]
    v6 = (-(v6))
    for fi16 := 0; fi16 < 7; fi16++ {
        v2 -= 127
        _ = fi16
    }
    v2 = v2
    v5[(v15 & (v15 >> ((53) & 63)))] = (!v14)
    v17 := GenS1{F0: 42}
    var v18 float64 = (float64(v4) / float64(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
