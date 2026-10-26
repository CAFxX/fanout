package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]bool)
    v3 := make(map[int]int)
    var v4 bool = ((!false) || true)
    var v5 int32 = -2147483648
    var v6 float64 = -2.25
    for fi7 := 0; fi7 < 15; fi7++ {
        v1[fi7] = fi7
        v8 := GenS1{F0: 65535}
        v9 := GenS1{F0: 42}
        _ = fi7
        _ = v8
        _ = v9
    }
    if (v4 && false) {
        for fi10 := 0; fi10 < 11; fi10++ {
            switch (-(fi10)) {
            case -2147483648:
                for fi11 := 0; fi11 < 14; fi11++ {
                    v12 := GenS2{F0: (v4 || false)}
                    _ = fi11
                    _ = v12
                }
            case 1:
                switch ("foo bar" + "hello") {
                case "a":
                    v13 := GenS3{F0: -1}
                    _ = v13
                default:
                    var v14 uint = 42
                    _ = v14
                }
            default:
                for fi15 := 0; fi15 < 4; fi15++ {
                    delete(v3, (fi15 | 648))
                    fmt.Println(v6)
                    v16 := GenS4{F0: 211}
                    _ = fi15
                    _ = v16
                }
            }
            _ = fi10
        }
        v17 := v2[2]
        _ = v17
    }
    var v18 bool = (((((v6 * v6) - v6) + v6) <= (v6 / 0.1)) || v4)
    if (!(42 == -217)) {
        switch 1 {
        case 32767:
            var v19 uint8 = 127
            _ = v19
        default:
            v5 -= (v5 | v5)
        }
        fmt.Println(v4, v6, v6)
        for fi20 := 0; fi20 < 12; fi20++ {
            if v4 {
                var v21 uint64 = 830
                _ = v21
            }
            _ = fi20
        }
    }
    fmt.Println(v18, v5)
    v18 = (39 < -1)
    v22 := GenS5{F0: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
