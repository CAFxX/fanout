package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64 }

func main() {
    v1 := make(map[string]int)
    var v2 bool = true
    v3 := make(map[int]uint32)
    var v4 int32 = 963
    v5 := make(map[uint32]string)
    var v6 uint32 = 662
    fmt.Println(v4, v4)
    for fi7 := 0; fi7 < 9; fi7++ {
        v2 = (false && ((!(100 >= 100)) || v2))
        switch 9223372036854775807 {
        case 732:
            var v8 int16 = 42
            _ = v8
        case 127:
            var v9 string = "hello"
            _ = v9
        case 1:
            var v10 uint = 2
            _ = v10
        default:
            if (0.1 < float64(3.14159)) {
                var v11 int64 = -1
                var v12 uint = 1
                var v13 uint = (v12 & (v12 << ((255) & 63)))
                _ = v11
                _ = v12
                _ = v13
            }
        }
        _ = fi7
    }
    var v14 int32 = (v4 ^ v4)
    if true {
        var v15 int32 = (v14 ^ v14)
        v6 = (v6 % ((v6) | 1))
        switch 128 {
        case 146:
            v16 := []int{}
            _ = v16
        case 1:
            v17 := GenS1{F0: false, F1: (v14 | (v15 / ((-1) | 1)))}
            _ = v17
        case 32768:
            v18 := []int{-2147483648, int(v6)}
            _ = v18
        default:
            v19 := GenS2{F0: (v6 / ((v6) | 1))}
            _ = v19
        }
        _ = v15
    }
    v2 = (32767 != 793)
    var v20 int16 = 1000
    delete(v1, (("foo bar" + ("foo bar" + "")) + "hello"))
    v21 := GenS3{F0: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
