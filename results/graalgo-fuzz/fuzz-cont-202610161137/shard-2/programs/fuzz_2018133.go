package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[int]bool)
    var v3 bool = (!((((0.1 - 3.14159) != float32(1.5)) && true) || (140 >= 25)))
    var v4 int8 = -1
    v5 := make(map[uint32]int)
    var v6 uint32 = 65535
    if false {
        var v7 int64 = -685
        var v8 uint16 = 42
        v9 := GenS1{F0: (v6 + (v6 | (v6 | v6))), F1: 102, F2: v4}
        _ = v7
        _ = v8
        _ = v9
    }
    if (-2147483648 <= -9223372036854775808) {
        var v10 uint64 = v1
        _ = v10
    } else {
        v6 -= 2
    }
    var v11 uint8 = 42
    v12 := GenS1{F0: v6, F1: (v11 | 40), F2: (v4 % (((-(v4))) | 1))}
    var v13 int32 = 25
    _ = len(v5)
    for fi14 := 0; fi14 < 17; fi14++ {
        for fi15 := 0; fi15 < 1; fi15++ {
            var v16 float64 = (1.5 * -2.25)
            if (9223372036854775807 <= 0) {
                if v3 {
                    v17 := []int{}
                    _ = v17
                } else {
                    v2[fi14] = ((v3 || ((0.1 / ((-2.25 + 1.5) - 0.1)) == ((1e10 - 0.0) * -2.25))) && false)
                    var v18 int = fi14
                    _ = v18
                }
                if (fi14 > fi14) {
                    var v19 string = "\n\t"
                    v6 += 1
                    v20 := GenS2{F0: -726}
                    _ = v19
                    _ = v20
                }
            }
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    for fi21 := 0; fi21 < 18; fi21++ {
        switch ("\n\t" + "hello") {
        case "hello":
            v11 -= 43
        }
        _ = fi21
    }
    v22 := GenS1{F0: v6, F1: v11, F2: (^(v4))}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
