package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int64 }

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 0
    var v3 uint8 = (^(v2))
    v4 := make(map[int]int)
    _ = len(v1)
    fmt.Println(v2, v2)
    var v5 uint8 = (v2 >> ((1) & 63))
    v6 := GenS1{F0: 255, F1: "", F2: -9223372036854775808}
    v2 -= v3
    var v7 int32 = 20
    if (false || (!(65535 >= 65535))) {
        v8 := []int{454}
        if (120 >= -1) {
            for fi9 := 0; fi9 < 8; fi9++ {
                v8 = append(v8, fi9)
                _ = fi9
            }
        } else {
            switch 681 {
            case 127:
                if (-134 < 951) {
                    for fi10 := 0; fi10 < 4; fi10++ {
                        v8 = append(v8, fi10)
                        _ = fi10
                    }
                } else {
                    var v11 bool = true
                    v1[2147483647] = 127
                    _ = v11
                }
            case 105:
                for fi12 := 0; fi12 < 14; fi12++ {
                    v5 = 255
                    var v13 int64 = 654
                    var v14 uint16 = 1
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            }
            var v15 uint16 = 1
            _ = v15
        }
        _ = v8
    }
    var v16 uint = 2
    var v17 uint8 = v2
    v17 = 166
    var v18 string = "\n\t"
    fmt.Println(v3, v18, v16)
    v19 := GenS1{F0: 1, F1: (v18 + "foo bar"), F2: 42}
    v20 := GenS1{F0: 266, F1: v18, F2: -257}
    var v21 uint = v16
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
