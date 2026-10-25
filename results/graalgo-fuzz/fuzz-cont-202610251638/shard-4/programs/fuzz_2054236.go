package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS2 struct { F0 bool; F1 uint8; F2 bool }
type GenS3 struct { F0 uint; F1 int16 }

func main() {
    var v1 string = ("a" + "\n\t")
    var v2 uint64 = 4294967295
    var v3 int16 = -576
    v4 := GenS1{F0: 42, F1: 82, F2: 2147483647}
    var v5 int = -1
    v6 := GenS2{F0: false, F1: 176, F2: (true || false)}
    switch v1 {
    case "hello":
        for fi7 := 0; fi7 < 19; fi7++ {
            v3 -= (v3 | (^(v3)))
            v1 = ("" + "foo bar")
            v1 = ("\n\t" + "hello")
            _ = fi7
        }
    }
    fmt.Println(v5, v1, v1)
    for fi8 := 0; fi8 < 15; fi8++ {
        v5 -= (fi8 % (((fi8 % (((fi8 | (v5 * v5))) | 1))) | 1))
        var v9 uint32 = 2
        var v10 int = (-(v5))
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := []int{(v5 + 834), (v5 >> ((4294967295) & 63)), (v5 << ((402) & 63)), v5}
    for fi12 := 0; fi12 < 17; fi12++ {
        if (((false || (false || ((1.5 - 0.1) < (1e10 + 0.0)))) && (!(true || true))) && false) {
            var v13 uint64 = v2
            _ = v13
        } else {
            v11 = append(v11, 127)
        }
        v11 = append(v11, 42)
        if true {
            var v14 uint32 = 1
            if (!false) {
                var v15 uint = 18446744073709551615
                var v16 int32 = -2147483648
                _ = len(v11)
                _ = v15
                _ = v16
            } else {
                for fi17 := 0; fi17 < 9; fi17++ {
                    v18 := GenS3{F0: 30, F1: v3}
                    var v19 float64 = 0.1
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            }
            _ = v14
        }
        _ = fi12
    }
    v2 -= v2
    v3 = (^(v3))
    var v20 uint = 264
    v3 += (v3 << ((v20) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
