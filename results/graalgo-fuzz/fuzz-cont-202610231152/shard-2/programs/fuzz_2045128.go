package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int16 }

func main() {
    var v1 int32 = 0
    var v2 int64 = -916
    v3 := []string{"hello", ("a" + ""), "a"}
    v4 := make(map[int]bool)
    if true {
        v2 = (v2 - 0)
        var v5 int = 2147483647
        _ = v5
    }
    var v6 uint8 = 2
    if (1e10 != float64(3.14159)) {
        _ = len(v3)
        v7 := GenS1{F0: 4294967295, F1: v2, F2: 444}
        _ = v7
    } else {
        switch ("" + ("hello" + ("hello" + "foo bar"))) {
        case "":
            var v8 int = -282
            _ = v8
        case "hello":
            v2 = (v2 >> ((1) & 63))
        case "\n\t":
            var v9 int8 = 1
            _ = v9
        }
    }
    v3 = append(v3, "foo bar")
    for fi10 := 0; fi10 < 5; fi10++ {
        _ = len(v3)
        _ = fi10
    }
    var v11 string = "a"
    for fi12 := 0; fi12 < 19; fi12++ {
        for fi13 := 0; fi13 < 5; fi13++ {
            v1 = v1
            v2 = (v2 | v2)
            _ = fi13
        }
        _ = fi12
    }
    var v14 int64 = v2
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 int8 = 27
        var v17 float64 = 0.1
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 uint32 = 205
    var v19 uint16 = 255
    v1 += -1
    v19 = v19
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
