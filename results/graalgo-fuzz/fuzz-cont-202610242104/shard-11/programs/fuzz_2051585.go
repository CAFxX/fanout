package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint32; F2 int16 }

func main() {
    var v1 int8 = 127
    var v2 int32 = -428
    v3 := []int{32767, -295, 0}
    var v4 float32 = 0.1
    var v5 int64 = 9223372036854775807
    var v6 uint64 = 1
    if (!("foo bar" <= "a")) {
        var v7 int = int(v2)
        _ = v7
    } else {
        v1 = 100
        fmt.Println(v2, v5)
    }
    v5 += (-(v5))
    v3 = append(v3, int(v2))
    if true {
        v8 := GenS1{F0: 1000}
        v4 += (v4 / -2.25)
        var v9 uint64 = v6
        _ = v8
        _ = v9
    }
    switch 2147483647 {
    case 9223372036854775807:
        fmt.Println(v2, v4)
    case 0:
        var v10 int = -1
        _ = v10
    default:
        if (true && (!((-1 <= v2) || false))) {
            var v11 uint8 = 127
            switch 260 {
            case 32768:
                v4 += (v4 + 3.14159)
            }
            if (((false && false) || (true || false)) && (!(!true))) {
                if ((true || true) || (int(v1) == 4294967295)) {
                    if len(v3) > 0 {
                        v3[0] = int(v6)
                    }
                }
                for fi12 := 0; fi12 < 3; fi12++ {
                    v13 := GenS2{F0: 65535, F1: 155, F2: -32768}
                    _ = fi12
                    _ = v13
                }
                var v14 float64 = ((((1.5 + 3.14159) + -2.25) - 3.14159) / (1.5 / -2.25))
                _ = v14
            } else {
                if len(v3) > 0 {
                    v3[0] = -1
                }
                v3 = append(v3, -393)
            }
            _ = v11
        } else {
            if len(v3) > 0 {
                v3[0] = 1
            }
        }
    }
    var v15 int16 = -286
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v16 int32 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
