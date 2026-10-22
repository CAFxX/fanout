package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0() int {
    v6 := GenS1{F0: 42}
    var v7 uint8 = 128
    _ = v6
    _ = v7
    return 128
}
func genFn1(p0 uint16, p1 int64) uint64 {
    p0 = (p0 - 561)
    for fi8 := 0; fi8 < 19; fi8++ {
        p0 += p0
        if (!(-213 > fi8)) {
            v9 := GenS1{F0: 1}
            for fi10 := 0; fi10 < 15; fi10++ {
                switch int(p0) {
                case 42:
                    switch fi10 {
                    case 95:
                        var v11 uint16 = p0
                        _ = v11
                    default:
                        p0 += p0
                    }
                default:
                    p0 = p0
                }
                var v12 uint16 = 42
                _ = fi10
                _ = v12
            }
            _ = v9
        }
        _ = fi8
    }
    if false {
        if ((p1 / ((785) | 1)) >= p1) {
            var v13 uint64 = 908
            var v14 uint = 530
            _ = v13
            _ = v14
        } else {
            var v15 int8 = 1
            var v16 int16 = -32768
            fmt.Println(v15, v16, p0)
            _ = v15
            _ = v16
        }
    } else {
        switch -171 {
        case 9223372036854775807:
            v17 := []int{-1}
            _ = v17
        }
        var v18 float32 = ((1.5 - float32(1.5)) - 1e10)
        var v19 uint64 = 4294967295
        _ = v18
        _ = v19
    }
    p0 -= (p0 - p0)
    fmt.Println(p0)
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := make(map[string]string)
    var v2 int8 = 100
    var v3 int32 = -2147483648
    var v4 int = 42
    var v5 string = ("\n\t" + "foo bar")
    var v20 int64 = 9223372036854775807
    v2 = (v2 << ((0) & 63))
    _ = len(v1)
    v20 = (v20 * 9223372036854775807)
    var v21 int8 = v2
    v22 := v1[(((("\n\t" + (v5 + (v5 + ((((("hello" + (v5 + (v5 + ("\n\t" + v5)))) + ("a" + v5)) + v5) + v5) + ("" + ""))))) + v5) + v5) + "foo bar")]
    var v23 string = "hello"
    v2 += (v2 | (v2 - (^(v2))))
    switch v5 {
    case "hello":
        v24 := []int{(^(v4))}
        _ = v24
    default:
        v23 = v5
    }
    v25 := GenS1{F0: 481}
    var v26 int32 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
