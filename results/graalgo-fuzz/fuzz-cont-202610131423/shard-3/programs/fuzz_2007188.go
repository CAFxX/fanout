package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := []string{("" + "a"), "hello", "hello"}
    var v2 float64 = float64(3.14159)
    var v3 int16 = 95
    for fi4 := 0; fi4 < 11; fi4++ {
        v1 = append(v1, "a")
        fmt.Println(v2)
        var v5 uint32 = 1
        _ = fi4
        _ = v5
    }
    var v6 string = "\n\t"
    for fi7 := 0; fi7 < 5; fi7++ {
        if (551 >= 18446744073709551615) {
            var v8 uint8 = 255
            var v9 bool = (51 == 0)
            _ = v8
            _ = v9
        } else {
            for fi10 := 0; fi10 < 7; fi10++ {
                v11 := GenS1{F0: (2 > 128)}
                _ = fi10
                _ = v11
            }
        }
        v3 -= v3
        _ = fi7
    }
    var v12 uint8 = 1
    switch -600 {
    case 127:
        switch 32768 {
        case 128:
            if len(v1) > 0 {
                v1[0] = v6
            }
        case 491:
            v2 = (v2 + (v2 / float64((-2.25 * 1.5))))
        default:
            v6 = "a"
        }
    case 1:
        v13 := GenS1{F0: (!true)}
        _ = v13
    default:
        v12 -= 2
    }
    var v14 int32 = -332
    for fi15 := 0; fi15 < 4; fi15++ {
        v16 := GenS1{F0: true}
        fmt.Println(v12, v14, fi15)
        if false {
            if true {
                _ = len(v1)
            } else {
                var v17 float32 = -2.25
                var v18 float32 = v17
                _ = v17
                _ = v18
            }
        }
        _ = fi15
        _ = v16
    }
    v2 += (float64(float32(1e10)) / ((v2 * 1.5) + (v2 * float64((1.5 / ((0.1 / ((3.14159 - 0.0) / 3.14159)) - (-2.25 - float32(float64(3.14159)))))))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v6
}
