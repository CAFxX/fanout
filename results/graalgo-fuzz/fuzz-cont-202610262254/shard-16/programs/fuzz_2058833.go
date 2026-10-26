package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
func genFn0() uint64 {
    var v5 int8 = 62
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 int16 = 32767
        v7 = 42
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 841, F1: 42}
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 uint32 = 4294967295
        v11 := GenS1{F0: (^(fi9)), F1: 1}
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 uint32 = 332
    _ = v12
    _ = v5
    _ = v8
    return 729
}

func main() {
    var v1 int8 = -1
    var v2 int8 = (v1 * (v1 | v1))
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 151
    v1 = v1
    fmt.Println(v4, v4)
    v4 -= v4
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 string = "\n\t"
        _ = fi13
        _ = v14
    }
    var v15 string = "foo bar"
    fmt.Println(v1, v4)
    switch v15 {
    case "\n\t":
        v16 := GenS1{F0: 784, F1: 42}
        _ = v16
    case "a":
        var v17 int8 = v2
        _ = v17
    case "hello":
        v15 = v15
    }
    if ((1 <= 42) && ((3.14159 + float32(3.14159)) > 0.0)) {
        if (749 >= 899) {
            var v18 float32 = 0.1
            _ = v18
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
