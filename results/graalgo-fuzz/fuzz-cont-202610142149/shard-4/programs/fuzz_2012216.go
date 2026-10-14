package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint8 }

func main() {
    var v1 uint16 = 459
    var v2 bool = true
    var v3 bool = v2
    v4 := []int{2147483647, -9223372036854775808}
    v4 = append(v4, 2147483647)
    fmt.Println(v1, v1, v2)
    v3 = ((((-2.25 + 1e10) * 3.14159) + -2.25) != ((0.0 * (float64(1.5) / 1.0)) - (0.0 - 1e10)))
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 string = "a"
        if (v1 > 32767) {
            if len(v4) > 0 {
                v4[0] = fi5
            }
            v3 = v3
        }
        _ = fi5
        _ = v6
    }
    switch "hello" {
    case "":
        var v7 uint16 = (v1 >> ((65535) & 63))
        _ = v7
    }
    v8 := GenS1{F0: (1000 != 1), F1: 0, F2: 255}
    for fi9 := 0; fi9 < 16; fi9++ {
        v2 = v2
        v4 = append(v4, 127)
        for fi10 := 0; fi10 < 10; fi10++ {
            v11 := GenS1{F0: v3, F1: 18446744073709551615, F2: 2}
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    v12 := GenS1{F0: true, F1: 42, F2: 255}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v3)
    v13 := GenS1{F0: (!((float64(1e10) * ((-2.25 / (float64((1e10 * (0.1 + (-2.25 - -2.25)))) - 0.0)) + float64(1e10))) >= ((1e10 / (0.1 - 1e10)) * ((((-2.25 - float64((((float32(1.5) + (1.5 * 3.14159)) - 3.14159) / 1e10))) + (1.5 - float64(0.1))) / 0.1) - 0.0)))), F1: 18446744073709551615, F2: 255}
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
