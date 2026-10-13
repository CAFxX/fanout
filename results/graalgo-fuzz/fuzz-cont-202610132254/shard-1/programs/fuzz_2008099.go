package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int64 }
type GenS2 struct { F0 bool }

func main() {
    var v1 float64 = (0.0 * float64((((1.5 - float32(float64(0.0))) - -2.25) / float32(3.14159))))
    v2 := make(map[string]string)
    var v3 int8 = 42
    v4 := make(map[int]string)
    var v5 int64 = 0
    var v6 string = "\n\t"
    var v7 int64 = (v5 / (((v5 | 1)) | 1))
    fmt.Println(v1, v5)
    v8 := GenS1{F0: v3, F1: v7, F2: (v5 * v5)}
    v9 := v4[42]
    v9 = ((v6 + "a") + v9)
    fmt.Println(v3, v9)
    v4[42] = ("\n\t" + "\n\t")
    v3 = v3
    fmt.Println(v1, v6)
    v10 := GenS1{F0: (v3 % (((v3 + (^(v3)))) | 1)), F1: v5, F2: v7}
    var v11 float64 = 1.5
    for fi12 := 0; fi12 < 19; fi12++ {
        for fi13 := 0; fi13 < 20; fi13++ {
            v11 -= v1
            if true {
                switch int(v5) {
                case 128:
                    v9 = ("\n\t" + v6)
                case -953:
                    var v14 float64 = 1e10
                    _ = v14
                default:
                    v15 := v4[int(v5)]
                    _ = v15
                }
            }
            v11 -= v1
            _ = fi13
        }
        v16 := GenS1{F0: 100, F1: (v5 / (((v7 << ((42) & 63))) | 1)), F2: (v7 + 42)}
        switch -240 {
        case -592:
            v1 = -2.25
        case 32767:
            var v17 float32 = (3.14159 * float32(v1))
            _ = v17
        default:
            if false {
                var v18 int16 = 0
                _ = v18
            } else {
                v19 := []int{fi12, fi12}
                _ = v19
            }
        }
        _ = fi12
        _ = v16
    }
    v20 := GenS1{F0: 83, F1: -1, F2: v7}
    v21 := GenS2{F0: (((v5 + int64(v3)) <= 42) && (-438 > 0))}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
