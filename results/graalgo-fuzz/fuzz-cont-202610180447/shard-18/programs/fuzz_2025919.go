package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 string; F2 int64 }
type GenS3 struct { F0 int64; F1 int64; F2 int8 }
type GenS4 struct { F0 uint16; F1 uint16 }
func genFn0(p0 int16, p1 uint16) uint16 {
    v4 := []int{42, 104}
    for fi5 := 0; fi5 < 17; fi5++ {
        v6 := GenS1{F0: (p1 < p1)}
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    return p1
}

func main() {
    var v1 float64 = (1.5 + 1.5)
    var v2 int8 = -3
    var v3 int = -2147483648
    v7 := GenS2{F0: (v3 << ((65535) & 63)), F1: (("\n\t" + "\n\t") + "a"), F2: 9223372036854775807}
    switch v3 {
    case 9223372036854775807:
        var v8 uint8 = 225
        _ = v8
    }
    if true {
        for fi9 := 0; fi9 < 2; fi9++ {
            if false {
                v1 -= v1
                if (1 < 155) {
                    v10 := []int{262, (fi9 / (((v3 & fi9)) | 1)), v3, -2147483648}
                    _ = v10
                }
                for fi11 := 0; fi11 < 5; fi11++ {
                    v2 = (v2 | 127)
                    _ = fi11
                }
            }
            _ = fi9
        }
        if false {
            v12 := GenS3{F0: 42, F1: -9223372036854775808, F2: (^(v2))}
            _ = v12
        } else {
            v13 := []int{v3}
            _ = v13
        }
    }
    v14 := GenS2{F0: 0, F1: ("a" + "hello"), F2: 1}
    switch 42 {
    case 2147483647:
        fmt.Println(v3, v1, v2)
    case 0:
        v15 := []int{(v3 + v3), 130, v3, v3}
        _ = v15
    case 127:
        fmt.Println(v3, v1)
    default:
        switch "" {
        case "foo bar":
            var v16 int16 = 32767
            _ = v16
        default:
            v17 := GenS1{F0: true}
            _ = v17
        }
    }
    var v18 int16 = -437
    for fi19 := 0; fi19 < 16; fi19++ {
        v20 := GenS4{F0: 173, F1: 613}
        _ = fi19
        _ = v20
    }
    if ((2 < 4294967295) && true) {
        v18 += v18
        var v21 int32 = -2147483648
        fmt.Println(v1, v2, v2)
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v7
}
