package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint16 }

func main() {
    var v1 uint = 65535
    v2 := make(map[string]int)
    var v3 uint = v1
    for fi4 := 0; fi4 < 2; fi4++ {
        var v5 uint16 = 255
        if true {
            _ = len(v2)
        }
        v2["hello"] = 2147483647
        _ = fi4
        _ = v5
    }
    fmt.Println(v1, v1)
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 int32 = -601
        _ = fi6
        _ = v7
    }
    v1 = 606
    v3 = v1
    if (true || true) {
        var v8 int8 = 0
        v3 = 0
        _ = v8
    } else {
        v9 := GenS1{F0: 2, F1: 1, F2: 203}
        _ = v9
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 int = 128
        for fi12 := 0; fi12 < 10; fi12++ {
            v3 = 42
            v13 := GenS1{F0: 255, F1: 656, F2: 2}
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 18; fi14++ {
            var v15 float64 = ((-2.25 * ((-2.25 * 1.5) - 3.14159)) * 0.0)
            var v16 int16 = 1
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi10
        _ = v11
    }
    fmt.Println(v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}
