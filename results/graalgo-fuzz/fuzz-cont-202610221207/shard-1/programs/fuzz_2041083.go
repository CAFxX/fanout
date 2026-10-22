package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }

func main() {
    var v1 uint32 = 65535
    var v2 uint64 = 18446744073709551615
    var v3 uint16 = 2
    var v4 bool = (v3 <= 42)
    var v5 int = 1
    v6 := make(map[int]uint32)
    fmt.Println(v1, v3, v4)
    v7 := []int{}
    for fi8 := 0; fi8 < 19; fi8++ {
        for fi9 := 0; fi9 < 1; fi9++ {
            v1 += 42
            _ = fi9
        }
        _ = fi8
    }
    var v10 uint = 354
    v7 = append(v7, v5)
    for fi11 := 0; fi11 < 7; fi11++ {
        fmt.Println(v4, v1)
        delete(v6, -390)
        v12 := GenS1{F0: -1, F1: (v2 - (v2 + (^(v2))))}
        _ = fi11
        _ = v12
    }
    if (((1e10 / float64((1.5 / 1.5))) - -2.25) < 0.1) {
        fmt.Println(v3)
        fmt.Println(v5)
    }
    if ((v10 < (v10 + (v10 % ((v10) | 1)))) && ("a" <= "a")) {
        v13 := GenS1{F0: 1, F1: (v2 & uint64(v10))}
        _ = v13
    }
    var v14 int32 = 42
    if false {
        switch int(v1) {
        case -1:
            if v4 {
                var v15 float32 = 1.5
                _ = v15
            } else {
                switch ("foo bar" + ((((("" + "foo bar") + "\n\t") + "a") + "") + ("" + "a"))) {
                case "a":
                    fmt.Println(v2, v4)
                case "hello":
                    for fi16 := 0; fi16 < 3; fi16++ {
                        for fi17 := 0; fi17 < 7; fi17++ {
                            var v18 uint8 = 0
                            _ = fi17
                            _ = v18
                        }
                        _ = fi16
                    }
                case "foo bar":
                    var v19 float64 = float64(((-2.25 - float32(((1.5 + float64(0.0)) - 0.0))) - 0.1))
                    _ = v19
                }
                v3 = v3
            }
        case 32767:
            _ = len(v6)
        case 724:
            v20 := v6[(v5 & int(v2))]
            _ = v20
        default:
            var v21 uint16 = 901
            _ = v21
        }
    }
    if v4 {
        v5 = v5
    }
    v2 = v2
    fmt.Println(v5, v1, v4)
    v22 := GenS1{F0: -763, F1: 4294967295}
    var v23 int = v5
    switch -9223372036854775808 {
    case 4294967295:
        fmt.Println(v14, v3, v1)
    default:
        v24 := GenS1{F0: 32767, F1: v2}
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
