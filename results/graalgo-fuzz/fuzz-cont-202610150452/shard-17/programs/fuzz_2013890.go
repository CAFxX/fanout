package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 uint16 = 0
    var v2 uint64 = 1
    var v3 float32 = 0.0
    var v4 float32 = v3
    v1 = uint16(v2)
    v5 := GenS1{F0: -1}
    var v6 uint8 = 42
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            if false {
                if false {
                    switch fi7 {
                    case 717:
                        if (true || (-25 != 221)) {
                            var v9 float64 = float64(((0.1 * v3) + v4))
                            _ = v9
                        }
                    }
                    var v10 uint64 = v2
                    _ = v10
                }
            } else {
                var v11 bool = false
                fmt.Println(v3)
                _ = v11
            }
            _ = fi8
        }
        _ = fi7
    }
    var v12 int64 = -9223372036854775808
    v3 = v3
    var v13 uint16 = v1
    fmt.Println(v2, v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
