package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int }

func main() {
    v1 := make(map[int]string)
    var v2 float64 = (3.14159 * float64(0.1))
    var v3 int = 23
    var v4 float64 = (((1.5 * (float64((float32(0.0) / 1e10)) - float64(1e10))) + float64(((3.14159 * -2.25) * 3.14159))) * (v2 - (v2 - v2)))
    fmt.Println(v2, v2)
    var v5 uint = 4294967295
    var v6 int16 = 1
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 uint16 = 32767
        var v9 int32 = 0
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 float64 = float64(0.0)
    var v11 bool = ((true || false) && (false || (v10 != float64((-2.25 * (0.0 * float32((((3.14159 / (v2 - (((float64(((((0.1 * 1.5) - -2.25) * (3.14159 / (0.1 - -2.25))) + 1.5)) / ((v4 * v4) / 1.0)) + (3.14159 / (0.0 + (1.5 * (float64((1e10 + ((float32(((float64(1e10) / float64(0.1)) - (v2 / v2))) + float32((1.5 / 1.5))) + 3.14159))) / 3.14159))))) * (v2 * 1e10)))) - v4) + v2))))))))
    fmt.Println(v2, v4)
    switch "foo bar" {
    case "a":
        for fi12 := 0; fi12 < 10; fi12++ {
            if (true && true) {
                fmt.Println(v5, v4)
            }
            if (!((fi12 & 4294967295) != (v3 / (((v3 & fi12)) | 1)))) {
                for fi13 := 0; fi13 < 13; fi13++ {
                    var v14 string = ("\n\t" + ("\n\t" + "a"))
                    var v15 uint8 = 128
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                v5 += v5
            }
            var v16 int16 = v6
            _ = fi12
            _ = v16
        }
    default:
        var v17 uint16 = 0
        _ = v17
    }
    v10 -= 0.1
    var v18 uint8 = 1
    var v19 uint64 = 4294967295
    for fi20 := 0; fi20 < 13; fi20++ {
        var v21 bool = (!v11)
        v22 := GenS1{F0: 100, F1: 1, F2: (fi20 | (fi20 % (((v3 & 32767)) | 1)))}
        _ = fi20
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 14; fi23++ {
        v11 = true
        _ = fi23
    }
    _ = len(v1)
    if true {
        v5 -= 2
        v24 := GenS1{F0: 124, F1: -1, F2: 112}
        if v11 {
            var v25 uint = v5
            _ = v25
        } else {
            for fi26 := 0; fi26 < 18; fi26++ {
                var v27 float32 = 1.5
                _ = fi26
                _ = v27
            }
            if v11 {
                _ = len(v1)
                var v28 bool = (v4 > ((0.1 + 1e10) - (v4 * v10)))
                _ = v28
            } else {
                var v29 uint = v5
                _ = v29
            }
            v30 := []int{(v3 | (v3 + int(v6)))}
            _ = v30
        }
        _ = v24
    } else {
        var v31 uint16 = 1
        for fi32 := 0; fi32 < 17; fi32++ {
            var v33 uint8 = (v18 + (v18 & v18))
            v11 = (152 == 2147483648)
            var v34 int64 = -303
            _ = fi32
            _ = v33
            _ = v34
        }
        v35 := GenS1{F0: -128, F1: -128, F2: (v3 ^ (v3 >> (((v5 & (v5 | (v5 % ((v5) | 1))))) & 63)))}
        _ = v31
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
