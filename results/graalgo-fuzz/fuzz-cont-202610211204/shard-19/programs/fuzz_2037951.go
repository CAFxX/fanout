package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 string }
type GenS2 struct { F0 int32; F1 uint32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = (1.5 * ((float32(0.0) + 0.0) / (-2.25 + -2.25)))
    var v3 string = ("" + ("foo bar" + ""))
    v4 := GenS1{F0: 1, F1: 50, F2: v3}
    v3 = v3
    for fi5 := 0; fi5 < 10; fi5++ {
        delete(v1, (^(fi5)))
        switch (v3 + v3) {
        case "":
            for fi6 := 0; fi6 < 4; fi6++ {
                fmt.Println(v2, v3)
                _ = fi6
            }
        case "a":
            for fi7 := 0; fi7 < 17; fi7++ {
                v8 := v1[(fi7 ^ fi5)]
                _ = fi7
                _ = v8
            }
        default:
            v9 := []int{(fi5 ^ (fi5 ^ (fi5 * 865)))}
            _ = v9
        }
        v10 := GenS2{F0: 0, F1: 683}
        _ = fi5
        _ = v10
    }
    var v11 uint32 = 392
    v11 = v11
    v1[0] = (^(v11))
    v11 = v11
    if (true && false) {
        v2 = v2
        if (42 < 843) {
            for fi12 := 0; fi12 < 14; fi12++ {
                v3 = v3
                var v13 uint = 0
                _ = fi12
                _ = v13
            }
            var v14 uint = 65535
            _ = v14
        }
    }
    var v15 string = (((("a" + (((((("foo bar" + v3) + "hello") + ("a" + "\n\t")) + ((v3 + "a") + v3)) + (v3 + ("hello" + (v3 + v3)))) + "\n\t")) + v3) + "\n\t") + v3)
    var v16 float32 = ((v2 + 0.0) + v2)
    var v17 string = "\n\t"
    v2 += float32((1.5 / (1e10 - (1.5 / 1.0))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
