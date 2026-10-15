package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16 }

func main() {
    v1 := []string{("hello" + "hello"), "hello", "a"}
    v2 := []string{"hello"}
    v3 := make(map[uint32]string)
    var v4 uint8 = 85
    _ = len(v2)
    v4 = (^(v4))
    var v5 uint8 = v4
    var v6 float32 = (0.0 + (((1.5 + 3.14159) - 1e10) + -2.25))
    for fi7 := 0; fi7 < 15; fi7++ {
        if (v6 <= float32(0.1)) {
            for fi8 := 0; fi8 < 7; fi8++ {
                _ = len(v2)
                v9 := GenS1{F0: -316}
                v5 -= (^(v5))
                _ = fi8
                _ = v9
            }
            v5 += (v5 * (^(v4)))
            fmt.Println(v5, v6)
        } else {
            v5 -= 42
            if false {
                for fi10 := 0; fi10 < 5; fi10++ {
                    v11 := GenS1{F0: fi10}
                    var v12 float32 = v6
                    if ((!true) && (-2147483648 < 1)) {
                        v5 = v4
                    }
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                v4 += v4
            }
            var v13 float64 = (-2.25 - 0.1)
            _ = v13
        }
        var v14 uint64 = 18446744073709551615
        if false {
            v5 += (v4 * (v5 ^ (v5 + v4)))
            v15 := GenS1{F0: (fi7 << ((0) & 63))}
            _ = v15
        }
        _ = fi7
        _ = v14
    }
    v16 := GenS1{F0: -2147483648}
    if false {
        if (true && (-119 > 547)) {
            var v17 int = 4294967295
            v17 += -691
            _ = v17
        } else {
            if true {
                v5 -= 128
            } else {
                v18 := GenS2{F0: 42}
                _ = v18
            }
        }
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        var v20 float64 = (3.14159 / -2.25)
        _ = fi19
        _ = v20
    }
    v5 -= v5
    v6 = v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
