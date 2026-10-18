package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 26
    var v3 bool = true
    var v4 int16 = 976
    var v5 uint8 = 42
    for fi6 := 0; fi6 < 15; fi6++ {
        switch (fi6 / ((fi6) | 1)) {
        case 128:
            v7 := GenS1{F0: 18446744073709551615}
            _ = v7
        default:
            if (!false) {
                var v8 float32 = 0.1
                switch ("foo bar" + "\n\t") {
                case "\n\t":
                    fmt.Println(v2, v4)
                case "a":
                    var v9 uint64 = 0
                    _ = v9
                }
                _ = v8
            } else {
                for fi10 := 0; fi10 < 12; fi10++ {
                    v11 := []int{}
                    _ = fi10
                    _ = v11
                }
            }
        }
        _ = fi6
    }
    var v12 float64 = 1.5
    v13 := GenS1{F0: 2}
    switch 535 {
    case -659:
        v5 -= v2
    }
    var v14 int = 32767
    v15 := GenS1{F0: 42}
    v16 := GenS1{F0: 18446744073709551615}
    if (!v3) {
        fmt.Println(v3, v3, v2)
        var v17 uint16 = 773
        var v18 int16 = (v4 - 1)
        _ = v17
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
