package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 uint64 = 42
    v2 := make(map[string]uint32)
    var v3 float32 = -2.25
    var v4 int64 = -960
    var v5 string = "\n\t"
    v3 -= ((1e10 * v3) - v3)
    delete(v2, v5)
    fmt.Println(v1, v4)
    var v6 int16 = 431
    var v7 float64 = (0.1 - ((3.14159 * float64(float32((3.14159 + -2.25)))) - (((((-2.25 + (0.1 + 0.1)) - -2.25) + 1.5) + 1e10) - 1e10)))
    var v8 int = -9223372036854775808
    var v9 uint8 = 255
    fmt.Println(v4, v7)
    v10 := GenS1{F0: (v9 >> ((18446744073709551615) & 63))}
    fmt.Println(v9, v3, v9)
    if (874 >= 0) {
        var v11 bool = false
        _ = v11
    }
    if (true && false) {
        for fi12 := 0; fi12 < 7; fi12++ {
            if false {
                fmt.Println(v5, v7, v7)
            } else {
                var v13 bool = (!true)
                for fi14 := 0; fi14 < 4; fi14++ {
                    v15 := GenS1{F0: v9}
                    _ = fi14
                    _ = v15
                }
                _ = v13
            }
            v1 = 42
            _ = fi12
        }
    }
    if (1 < -1) {
        for fi16 := 0; fi16 < 9; fi16++ {
            var v17 uint16 = 0
            _ = fi16
            _ = v17
        }
        if (true && (-1 <= -1)) {
            for fi18 := 0; fi18 < 7; fi18++ {
                fmt.Println(v6, v3, v8)
                _ = fi18
            }
            for fi19 := 0; fi19 < 7; fi19++ {
                var v20 string = v5
                var v21 int = v8
                v22 := []int{(fi19 % ((503) | 1))}
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
        } else {
            v7 = v7
        }
        var v23 uint8 = v9
        _ = v23
    }
    v3 = 3.14159
    v24 := v2["a"]
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
