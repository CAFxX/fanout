package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint16 }

func main() {
    v1 := make(map[int]bool)
    var v2 int = 128
    var v3 float32 = 1.5
    if (!((32767 > 42) || (32767 >= 255))) {
        for fi4 := 0; fi4 < 4; fi4++ {
            fmt.Println(fi4, fi4)
            var v5 uint = 0
            _ = fi4
            _ = v5
        }
        v6 := GenS1{F0: 972, F1: 124, F2: 1}
        var v7 uint8 = 1
        _ = v6
        _ = v7
    } else {
        delete(v1, (v2 ^ -792))
    }
    switch 2147483647 {
    case -1:
        if (461 != 65535) {
            v8 := GenS1{F0: 4294967295, F1: -1, F2: 255}
            v3 = -2.25
            switch "foo bar" {
            case "":
                v3 = (float32(1.5) / (float32(3.14159) - ((v3 * ((1e10 - v3) / -2.25)) / 1.0)))
            case "foo bar":
                switch (v2 - v2) {
                case -533:
                    for fi9 := 0; fi9 < 14; fi9++ {
                        var v10 int16 = 42
                        _ = fi9
                        _ = v10
                    }
                default:
                    var v11 string = "foo bar"
                    _ = v11
                }
            case "a":
                fmt.Println(v2)
            default:
                fmt.Println(v2, v2)
            }
            _ = v8
        } else {
            v3 += float32((1.5 + (0.1 - ((1e10 - 1.5) + -2.25))))
            v2 = -644
            for fi12 := 0; fi12 < 8; fi12++ {
                var v13 int32 = 42
                _ = fi12
                _ = v13
            }
        }
    case 32768:
        v14 := []int{}
        _ = v14
    }
    v3 -= (0.0 - float32(1.5))
    var v15 bool = false
    v3 += (v3 + (float32(float64(0.0)) * (v3 / (v3 * ((v3 * 1e10) / 3.14159)))))
    for fi16 := 0; fi16 < 1; fi16++ {
        var v17 int64 = -9223372036854775808
        _ = fi16
        _ = v17
    }
    v3 += v3
    var v18 int64 = 9223372036854775807
    if v15 {
        for fi19 := 0; fi19 < 11; fi19++ {
            v18 = (v18 | 766)
            v20 := []int{(fi19 * 812), v2}
            v3 += v3
            _ = fi19
            _ = v20
        }
        var v21 int16 = -32768
        for fi22 := 0; fi22 < 4; fi22++ {
            v23 := []int{(v2 % (((fi22 & (v2 + int(v18)))) | 1))}
            var v24 int = 1
            v24 += (v2 + 32767)
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = v21
    }
    for fi25 := 0; fi25 < 6; fi25++ {
        if (true && (!((!(1e10 != float64((v3 / 1.0)))) || true))) {
            var v26 int8 = 100
            var v27 uint64 = 1
            _ = v26
            _ = v27
        }
        _ = fi25
    }
    var v28 int16 = 32767
    var v29 uint32 = 2147483647
    v30 := []int{v2, v2, (-(v2))}
    var v31 uint = 65535
    var v32 int64 = v18
    v32 -= v32
    var v33 int8 = -15
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
}
