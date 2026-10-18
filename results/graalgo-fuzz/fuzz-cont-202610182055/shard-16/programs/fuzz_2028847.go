package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 bool = (true || true)
    var v2 uint16 = 32767
    var v3 float64 = -2.25
    var v4 uint16 = (v2 / ((v2) | 1))
    v5 := make(map[string]int)
    v6 := make(map[string]string)
    var v7 bool = v1
    v8 := GenS1{F0: 0, F1: 1, F2: 97}
    v1 = v7
    var v9 int8 = 1
    var v10 uint64 = 0
    v11 := GenS2{F0: (^(v10))}
    if (false && (((^(v10)) != (v10 - 929)) || v7)) {
        for fi12 := 0; fi12 < 13; fi12++ {
            var v13 int = fi12
            for fi14 := 0; fi14 < 8; fi14++ {
                switch fi12 {
                case 127:
                    fmt.Println(v4, v9)
                }
                var v15 int16 = 32767
                _ = fi14
                _ = v15
            }
            v4 += v2
            _ = fi12
            _ = v13
        }
        v5["foo bar"] = 138
        delete(v6, (("foo bar" + "hello") + ("hello" + ("hello" + ""))))
    } else {
        var v16 uint8 = 105
        var v17 uint8 = (v16 + v16)
        _ = v16
        _ = v17
    }
    var v18 bool = true
    switch "\n\t" {
    case "":
        var v19 uint8 = 0
        _ = v19
    }
    for fi20 := 0; fi20 < 19; fi20++ {
        v21 := GenS1{F0: 0, F1: 2147483647, F2: 1}
        _ = fi20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
