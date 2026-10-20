package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 int32 = 0
    v2 := []int{1, -9223372036854775808, -801}
    v3 := make(map[uint32]bool)
    v2 = append(v2, 4294967295)
    var v4 int64 = -783
    var v5 float32 = 1.5
    var v6 int = -434
    fmt.Println(v1, v5)
    var v7 uint64 = 227
    v1 -= 1
    var v8 uint32 = 2147483648
    v6 = -913
    v2 = append(v2, v6)
    for fi9 := 0; fi9 < 18; fi9++ {
        switch (-(v6)) {
        case 9223372036854775807:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case 0:
            switch "a" {
            case "":
                if (!(1000 == 1000)) {
                    for fi10 := 0; fi10 < 12; fi10++ {
                        if len(v2) > 0 {
                            v2[0] = fi9
                        }
                        _ = fi10
                    }
                } else {
                    v1 += v1
                }
            default:
                var v11 uint = 0
                _ = v11
            }
        }
        v3[v8] = ((v8 % ((2) | 1)) >= v8)
        _ = fi9
    }
    v5 = (0.1 + 0.1)
    for fi12 := 0; fi12 < 14; fi12++ {
        if (((65535 <= 250) || (v6 >= 4294967295)) || false) {
            delete(v3, v8)
            for fi13 := 0; fi13 < 3; fi13++ {
                v14 := GenS1{F0: (^(v8))}
                _ = fi13
                _ = v14
            }
        }
        _ = fi12
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
