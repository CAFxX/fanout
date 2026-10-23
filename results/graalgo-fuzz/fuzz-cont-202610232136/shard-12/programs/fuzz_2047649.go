package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = false
    v3 := []int{32767}
    v4 := make(map[int]string)
    var v5 int16 = -458
    v6 := make(map[string]uint32)
    for fi7 := 0; fi7 < 20; fi7++ {
        for fi8 := 0; fi8 < 10; fi8++ {
            _ = len(v4)
            _ = fi8
        }
        fmt.Println(v5, fi7)
        _ = fi7
    }
    fmt.Println(v2)
    v5 = v5
    v9 := GenS1{F0: (v5 | 251)}
    fmt.Println(v5, v5)
    for fi10 := 0; fi10 < 1; fi10++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch fi10 {
        case 127:
            for fi11 := 0; fi11 < 17; fi11++ {
                delete(v6, "a")
                var v12 float32 = (0.0 + -2.25)
                var v13 uint64 = 42
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        _ = fi10
    }
    var v14 bool = (576 > 1)
    v3 = append(v3, -2147483648)
    switch "foo bar" {
    case "a":
        v14 = (2 > 2)
    case "":
        fmt.Println(v5, v14)
    case "foo bar":
        switch 9223372036854775807 {
        case 2147483647:
            for fi15 := 0; fi15 < 10; fi15++ {
                fmt.Println(v14, fi15, v2)
                var v16 float64 = float64((1e10 - 0.1))
                _ = fi15
                _ = v16
            }
        case 1:
            for fi17 := 0; fi17 < 4; fi17++ {
                if (!v2) {
                    var v18 uint64 = 4294967295
                    if (1 > 1) {
                        for fi19 := 0; fi19 < 16; fi19++ {
                            fmt.Println(v14)
                            var v20 uint32 = 65535
                            v1[(v20 * v20)] = fi19
                            _ = fi19
                            _ = v20
                        }
                        var v21 uint16 = 0
                        _ = v21
                    }
                    var v22 float64 = (float64(((-2.25 - 3.14159) * (3.14159 - float32(float64(0.0))))) + 3.14159)
                    _ = v18
                    _ = v22
                }
                _ = fi17
            }
        }
    default:
        fmt.Println(v5)
    }
    var v23 float32 = 0.1
    var v24 int16 = (v5 | 1)
    v14 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
