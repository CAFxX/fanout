package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := make(map[string]int)
    var v2 float64 = (-2.25 + 0.1)
    v3 := make(map[uint32]string)
    v4 := []uint32{65535, 464, 65535}
    var v5 uint64 = 905
    var v6 uint16 = uint16(v5)
    switch -9223372036854775808 {
    case -9223372036854775808:
        var v7 int64 = 42
        _ = v7
    case 1:
        var v8 string = ("\n\t" + ("foo bar" + "foo bar"))
        _ = v8
    default:
        for fi9 := 0; fi9 < 16; fi9++ {
            v2 += v2
            v10 := GenS1{F0: 1}
            _ = fi9
            _ = v10
        }
    }
    fmt.Println(v5, v6)
    v11 := GenS1{F0: 42}
    v12 := GenS1{F0: 9223372036854775807}
    for fi13 := 0; fi13 < 6; fi13++ {
        v2 += (((((((((((float64((float32(v2) * (float32(1e10) * (1e10 + (-2.25 - -2.25))))) - v2) * v2) / (1e10 - v2)) / 3.14159) + float64(0.0)) + v2) - v2) - 0.1) / 1.0) + (v2 + v2)) - float64(-2.25))
        var v14 int64 = 42
        switch int(v5) {
        case -2147483648:
            var v15 float32 = (float32((v2 - -2.25)) / 1.5)
            _ = v15
        }
        _ = fi13
        _ = v14
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        if (0 > 32767) {
            if false {
                switch "\n\t" {
                case "foo bar":
                    v6 = v6
                case "a":
                    _ = len(v4)
                default:
                    for fi17 := 0; fi17 < 10; fi17++ {
                        var v18 int = -125
                        if false {
                            fmt.Println(fi17, fi16)
                        }
                        _ = fi17
                        _ = v18
                    }
                }
            }
        }
        v5 -= 42
        v4 = append(v4, 153)
        _ = fi16
    }
    var v19 uint64 = v5
    var v20 float64 = v2
    var v21 uint = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
