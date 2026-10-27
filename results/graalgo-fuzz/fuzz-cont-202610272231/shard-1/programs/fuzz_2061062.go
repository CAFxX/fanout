package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 bool, p1 uint8, p2 int16) int32 {
    var v7 int16 = (^(p2))
    fmt.Println(p1, v7)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    var v1 int64 = -1
    var v2 uint16 = 0
    var v3 float32 = 1e10
    var v4 uint64 = 4294967295
    var v5 string = "\n\t"
    var v6 int32 = -1
    var v8 float64 = 1.5
    v9 := []int{42, -2147483648}
    v1 = v1
    _ = len(v9)
    if len(v9) > 0 {
        v9[0] = 127
    }
    v4 -= (v4 | v4)
    var v10 int8 = 22
    var v11 uint16 = v2
    v5 = v5
    if (!(322 >= (v4 / ((687) | 1)))) {
        switch -2147483648 {
        case -28:
            for fi12 := 0; fi12 < 14; fi12++ {
                var v13 int32 = int32(v11)
                for fi14 := 0; fi14 < 9; fi14++ {
                    v2 += 522
                    _ = fi14
                }
                fmt.Println(v6)
                _ = fi12
                _ = v13
            }
        case -2147483648:
            v4 = v4
        case 366:
            var v15 uint16 = (^(v11))
            _ = v15
        }
        for fi16 := 0; fi16 < 8; fi16++ {
            v5 = v5
            _ = fi16
        }
        var v17 uint16 = v11
        _ = v17
    }
    v5 = "foo bar"
    for fi18 := 0; fi18 < 15; fi18++ {
        switch v5 {
        case "":
            for fi19 := 0; fi19 < 17; fi19++ {
                v20 := GenS1{F0: 1}
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                _ = fi19
                _ = v20
            }
        }
        v9 = append(v9, (fi18 >> ((2) & 63)))
        _ = fi18
    }
    var v21 string = (v5 + v5)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
