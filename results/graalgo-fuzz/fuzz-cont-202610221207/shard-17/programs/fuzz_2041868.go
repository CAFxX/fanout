package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint16; F1 int64; F2 uint }
type GenS3 struct { F0 int8; F1 uint32 }
func genFn0(p0 uint32, p1 int, p2 uint64) uint8 {
    p0 -= p0
    var v5 uint8 = 128
    v6 := []int{p1, -2147483648}
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 uint16 = 65535
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return v5
}

func main() {
    var v1 int32 = 581
    v2 := make(map[uint32]bool)
    var v3 float32 = 1.5
    v4 := []string{((("\n\t" + ("foo bar" + "foo bar")) + "") + "hello")}
    var v9 float32 = 1.5
    var v10 float32 = v9
    fmt.Println(v3, v10)
    v11 := GenS1{F0: 255, F1: v1, F2: genFn0(2, 42, 519)}
    v3 += 0.1
    for fi12 := 0; fi12 < 11; fi12++ {
        v13 := v2[2147483648]
        v10 = float32(1.5)
        _ = fi12
        _ = v13
    }
    var v14 uint32 = 65535
    switch -740 {
    case 433:
        if true {
            v15 := GenS2{F0: 251, F1: 42, F2: 4294967295}
            var v16 int64 = -9223372036854775808
            switch ("" + ((("foo bar" + "a") + "a") + "")) {
            case "\n\t":
                fmt.Println(v1, v16)
            default:
                v17 := GenS3{F0: 42, F1: v14}
                _ = v17
            }
            _ = v15
            _ = v16
        } else {
            var v18 uint16 = 255
            v4 = append(v4, "\n\t")
            _ = v18
        }
    case -2147483648:
        for fi19 := 0; fi19 < 17; fi19++ {
            var v20 int64 = 0
            for fi21 := 0; fi21 < 1; fi21++ {
                delete(v2, v14)
                switch "a" {
                case "hello":
                    v10 = (v10 * 3.14159)
                default:
                    var v22 int16 = 629
                    _ = v22
                }
                _ = fi21
            }
            var v23 int16 = 32767
            _ = fi19
            _ = v20
            _ = v23
        }
    case 127:
        fmt.Println(v14)
    default:
        var v24 int = -339
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
