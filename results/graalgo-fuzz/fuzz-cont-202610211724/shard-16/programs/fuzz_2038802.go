package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint8, p1 int32) int64 {
    fmt.Println(p0, p0)
    v5 := GenS1{F0: -254}
    _ = p0
    _ = p1
    _ = v5
    return -675
}

func main() {
    var v1 int32 = 86
    var v2 int32 = (v1 | (v1 | (v1 + -624)))
    var v3 uint16 = 42
    var v4 uint16 = (v3 - v3)
    v3 -= (v4 & 0)
    v6 := GenS2{F0: 0}
    switch -1 {
    case -2147483648:
        var v7 float32 = ((1e10 * 0.0) / float32((0.1 * float64(float32(1.5)))))
        _ = v7
    case 32767:
        var v8 int = 2147483647
        _ = v8
    default:
        switch -1 {
        case -95:
            for fi9 := 0; fi9 < 20; fi9++ {
                var v10 uint32 = 662
                _ = fi9
                _ = v10
            }
        }
    }
    var v11 bool = true
    for fi12 := 0; fi12 < 16; fi12++ {
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 uint = 65535
            v15 := GenS1{F0: int64(v3)}
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi12
    }
    if (130 <= 65535) {
        fmt.Println(v4)
        switch -2147483648 {
        case 128:
            v1 = -2147483648
        case -1:
            var v16 int = -9223372036854775808
            _ = v16
        default:
            for fi17 := 0; fi17 < 9; fi17++ {
                var v18 int64 = 24
                v11 = ((v18 % ((v18) | 1)) < (^(v18)))
                _ = fi17
                _ = v18
            }
        }
        v19 := []int{316, 32767, 835}
        _ = v19
    } else {
        v20 := []int{42}
        var v21 uint8 = 127
        _ = v20
        _ = v21
    }
    var v22 bool = (-32768 < 1)
    if (float64(1e10) == float64((0.0 - 0.0))) {
        var v23 uint32 = 954
        switch "" {
        case "a":
            v2 += (v2 >> ((753) & 63))
        case "hello":
            var v24 int8 = 0
            _ = v24
        default:
            v22 = v22
        }
        _ = v23
    }
    var v25 string = ""
    fmt.Println(v2, v3)
    if ((!v22) || true) {
        v4 = (v4 / ((v4) | 1))
        for fi26 := 0; fi26 < 10; fi26++ {
            var v27 uint32 = 1
            _ = fi26
            _ = v27
        }
        v4 -= v3
    } else {
        var v28 uint16 = v4
        v29 := GenS2{F0: 4294967295}
        _ = v28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v6
}
