package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 string }

func main() {
    var v1 uint16 = 65535
    var v2 int8 = 0
    var v3 bool = false
    var v4 uint32 = 1
    var v5 string = "foo bar"
    var v6 int32 = int32(v2)
    for fi7 := 0; fi7 < 11; fi7++ {
        fmt.Println(v5, v3, v5)
        if v3 {
            v4 -= v4
        }
        _ = fi7
    }
    v1 = (v1 >> ((730) & 63))
    var v8 int8 = (^(v2))
    v9 := []int{}
    for fi10 := 0; fi10 < 10; fi10++ {
        for fi11 := 0; fi11 < 5; fi11++ {
            var v12 int16 = -781
            fmt.Println(v6, v3)
            v2 += v8
            _ = fi11
            _ = v12
        }
        if ((!v3) || v3) {
            v6 = v6
            switch "foo bar" {
            case "hello":
                for fi13 := 0; fi13 < 20; fi13++ {
                    fmt.Println(fi10, v4, v1)
                    _ = fi13
                }
            default:
                v14 := GenS1{F0: 2, F1: -18}
                _ = v14
            }
        }
        v15 := GenS1{F0: 18446744073709551615, F1: (v8 % ((v2) | 1))}
        _ = fi10
        _ = v15
    }
    v16 := GenS1{F0: 0, F1: -77}
    v8 = (-(v2))
    v6 = (v6 / ((369) | 1))
    v1 = (v1 ^ (v1 % ((v1) | 1)))
    var v17 float64 = 3.14159
    fmt.Println(v2, v5, v3)
    switch 9223372036854775807 {
    case 1:
        switch "hello" {
        case "foo bar":
            v5 = "\n\t"
        }
    default:
        if false {
            v18 := GenS2{F0: ""}
            if true {
                for fi19 := 0; fi19 < 10; fi19++ {
                    for fi20 := 0; fi20 < 16; fi20++ {
                        var v21 uint = 42
                        _ = fi20
                        _ = v21
                    }
                    _ = fi19
                }
                v22 := GenS2{F0: "a"}
                v9 = append(v9, 0)
                _ = v22
            } else {
                v17 = float64(1e10)
                fmt.Println(v4, v3, v8)
            }
            v1 += 2
            _ = v18
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
