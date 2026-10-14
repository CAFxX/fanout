package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }

func main() {
    var v1 bool = (((3.14159 * float64(1.5)) * -2.25) == (1.5 * 0.1))
    v2 := make(map[string]int)
    var v3 int = 0
    var v4 int32 = -8
    for fi5 := 0; fi5 < 8; fi5++ {
        v1 = (1.5 < (((0.1 - (0.0 + (((3.14159 / 1.0) / 0.1) * float32(-2.25)))) * 3.14159) / -2.25))
        _ = fi5
    }
    var v6 bool = v1
    fmt.Println(v6, v4)
    _ = len(v2)
    fmt.Println(v3)
    var v7 int32 = v4
    v4 = 0
    var v8 uint64 = 285
    v9 := GenS1{F0: 2, F1: 1}
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 int32 = (v4 / (((v7 % ((0) | 1))) | 1))
        v7 = 945
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: 42, F1: 568}
    for fi13 := 0; fi13 < 20; fi13++ {
        if (2 < 2) {
            var v14 uint16 = 65535
            if true {
                for fi15 := 0; fi15 < 10; fi15++ {
                    v8 = 84
                    v3 -= (v3 & 1)
                    _ = fi15
                }
                if (42 == 537) {
                    var v16 uint32 = 2147483648
                    _ = v16
                }
                v17 := []int{(v3 + (v3 | fi13)), fi13, fi13, v3}
                _ = v17
            } else {
                v4 = v7
                v18 := []int{fi13, fi13}
                for fi19 := 0; fi19 < 19; fi19++ {
                    v1 = (!(!true))
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                    v14 = v14
                    _ = fi19
                }
                _ = v18
            }
            var v20 int16 = 366
            _ = v14
            _ = v20
        } else {
            v1 = v6
        }
        _ = fi13
    }
    fmt.Println(v7, v6, v4)
    var v21 uint8 = 42
    var v22 int = v3
    var v23 uint = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
