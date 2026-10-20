package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 uint8 = 128
    v2 := make(map[int]int)
    var v3 int = 0
    var v4 string = "a"
    v5 := GenS1{F0: 0}
    v6 := GenS1{F0: 32767}
    if false {
        for fi7 := 0; fi7 < 19; fi7++ {
            for fi8 := 0; fi8 < 7; fi8++ {
                v9 := []int{fi7, fi8}
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
        _ = len(v2)
    } else {
        for fi10 := 0; fi10 < 5; fi10++ {
            v11 := []int{884}
            _ = fi10
            _ = v11
        }
        v1 = (v1 & v1)
    }
    var v12 uint32 = 42
    if (!false) {
        if ((((1.5 / 1.0) - (3.14159 / 1.0)) + -2.25) > 0.0) {
            if (!true) {
                v4 = "a"
                switch (v3 & v3) {
                case -9223372036854775808:
                    v3 -= 32768
                case 4294967295:
                    v3 = (v3 ^ (v3 >> ((18446744073709551615) & 63)))
                case 0:
                    _ = len(v2)
                default:
                    v13 := v2[-2147483648]
                    _ = v13
                }
            } else {
                if (!true) {
                    v14 := GenS2{F0: 985}
                    _ = v14
                }
                v3 += 42
            }
            var v15 float64 = -2.25
            v4 = "a"
            _ = v15
        } else {
            v16 := GenS1{F0: 255}
            v2[v3] = (^(v3))
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        v1 = (v1 | 137)
        fmt.Println(v3)
        v12 = v12
        _ = fi17
    }
    var v18 uint64 = 1
    if true {
        fmt.Println(v18)
    } else {
        fmt.Println(v3, v12)
        delete(v2, (v3 << ((65535) & 63)))
    }
    v1 -= (v1 | 250)
    var v19 string = v4
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
