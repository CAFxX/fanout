package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 uint16) bool {
    if (uint64(p0) >= 18446744073709551615) {
        for fi4 := 0; fi4 < 8; fi4++ {
            p0 = 46
            var v5 float64 = 1e10
            var v6 float64 = ((3.14159 - ((((((v5 * v5) * 0.0) + v5) / v5) - v5) + 3.14159)) + ((3.14159 * (0.1 + ((((1.5 / v5) + v5) * float64(0.1)) * 1.5))) - (-2.25 / v5)))
            _ = fi4
            _ = v5
            _ = v6
        }
    }
    if ((!(-1 >= -66)) || true) {
        p0 = p0
    } else {
        p0 = 126
        var v7 uint = 0
        if false {
            switch -2147483648 {
            case 32768:
                var v8 float32 = ((1e10 - (3.14159 - ((3.14159 * 1.5) - (3.14159 + ((1.5 / 0.1) * -2.25))))) - (((((0.1 + (3.14159 / 1e10)) * ((1.5 + float32(1e10)) / -2.25)) - float32(0.0)) + (1.5 - 0.1)) - ((1.5 - (float32((3.14159 / float64((float32(((((float64((((0.1 - (-2.25 / float32((0.1 - (1e10 / -2.25))))) - 0.1) / 3.14159)) + (0.0 + 0.0)) * 1.5) / float64(1.5)) * 3.14159)) / -2.25)))) + 3.14159)) * (-2.25 * (1.5 - 0.0)))))
                _ = v8
            default:
                v9 := GenS1{F0: (!(32767 < -1))}
                _ = v9
            }
            var v10 int64 = 514
            _ = v10
        }
        _ = v7
    }
    _ = p0
    return (1e10 >= 0.0)
}

func main() {
    var v1 string = "foo bar"
    var v2 uint32 = 1
    v3 := make(map[string]bool)
    v2 = v2
    var v11 int8 = -126
    for fi12 := 0; fi12 < 14; fi12++ {
        v2 += (v2 ^ v2)
        _ = fi12
    }
    fmt.Println(v2, v11, v11)
    v13 := GenS1{F0: true}
    v3[("foo bar" + v1)] = true
    var v14 float32 = ((1e10 / 0.1) - -2.25)
    var v15 uint = 18446744073709551615
    if (1.5 <= 3.14159) {
        _ = len(v3)
    }
    var v16 int = 2147483647
    v15 = v15
    var v17 uint8 = 128
    var v18 bool = (1 <= -1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
