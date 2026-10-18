package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 int8; F1 int }

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = 1e10
    var v3 bool = ("hello" != "a")
    v4 := []string{}
    var v5 bool = (-20 < 100)
    if (2 > 1) {
        var v6 float32 = (v2 - v2)
        _ = v6
    } else {
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := GenS1{F0: ((3.14159 * -2.25) > 1e10), F1: 97, F2: (fi7 * (fi7 >> ((4294967295) & 63)))}
            v5 = (-2.25 <= (float64(v2) / 3.14159))
            var v9 int16 = 56
            _ = fi7
            _ = v8
            _ = v9
        }
        var v10 uint8 = 255
        _ = v10
    }
    v11 := GenS2{F0: 0, F1: 1}
    fmt.Println(v3)
    v12 := GenS3{F0: 1, F1: -1}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v13 uint8 = 2
    var v14 int8 = -64
    for fi15 := 0; fi15 < 2; fi15++ {
        var v16 bool = (!(1 == 42))
        _ = fi15
        _ = v16
    }
    var v17 string = "foo bar"
    v13 = v13
    v2 = ((float32((((float64(3.14159) / (0.0 + 3.14159)) / -2.25) - (0.0 / (3.14159 / 1.5)))) - 0.1) * float32(3.14159))
    v17 = v17
    fmt.Println(v2, v3)
    v1[((("foo bar" + ((v17 + v17) + (v17 + (v17 + (v17 + "hello"))))) + "hello") + "\n\t")] = v3
    for fi18 := 0; fi18 < 20; fi18++ {
        if len(v4) > 0 {
            v4[0] = (v17 + ("a" + "foo bar"))
        }
        v5 = v3
        _ = fi18
    }
    v2 += ((v2 - 0.0) * (1.5 + v2))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
