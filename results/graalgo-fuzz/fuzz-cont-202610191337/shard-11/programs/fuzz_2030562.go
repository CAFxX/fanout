package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 uint, p1 uint8) bool {
    p0 -= 18446744073709551615
    p0 = 207
    _ = p0
    _ = p1
    return (((^(p0)) == 564) || false)
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 756
    var v3 string = (((((("\n\t" + "\n\t") + "") + ("foo bar" + (("hello" + "foo bar") + "a"))) + ("hello" + ("foo bar" + "hello"))) + "\n\t") + ("hello" + "foo bar"))
    var v4 int8 = 100
    var v5 int16 = 0
    var v6 uint32 = 268
    if (!("" >= v3)) {
        v5 = v5
    } else {
        v2 += v2
        for fi7 := 0; fi7 < 2; fi7++ {
            delete(v1, v6)
            var v8 int64 = 752
            _ = fi7
            _ = v8
        }
    }
    v3 = v3
    v9 := GenS1{F0: -2147483648}
    var v10 int8 = (v4 >> ((791) & 63))
    if (1.5 > 0.1) {
        var v11 int8 = v4
        v4 += v11
        var v12 string = "a"
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 float32 = 0.0
            switch -197 {
            case -9223372036854775808:
                switch -2147483648 {
                case 1:
                    var v16 bool = false
                    _ = v16
                case 32768:
                    v17 := GenS1{F0: 319}
                    _ = v17
                default:
                    var v18 int32 = -2147483648
                    _ = v18
                }
            case 251:
                var v19 int64 = 0
                _ = v19
            }
            v20 := []int{(-(fi14)), (fi14 * 127)}
            _ = fi14
            _ = v15
            _ = v20
        }
        var v21 int8 = (v10 * 127)
        if (1 == 0) {
            v22 := GenS1{F0: 0}
            _ = v22
        }
        _ = fi13
        _ = v21
    }
    switch 9223372036854775807 {
    case 4294967295:
        v23 := v1[42]
        _ = v23
    case 763:
        fmt.Println(v3, v5)
    case 0:
        v3 = v3
    default:
        if false {
            v5 = 1000
            var v24 int32 = 0
            _ = v24
        } else {
            var v25 float32 = ((((3.14159 + (0.0 - 1e10)) * 0.0) + 1.5) - (float32(((1e10 + float64(1e10)) * 1e10)) / 3.14159))
            v3 = ""
            v10 += 1
            _ = v25
        }
    }
    v10 = (v4 / ((v4) | 1))
    for fi26 := 0; fi26 < 17; fi26++ {
        if true {
            var v27 int32 = 432
            v28 := GenS1{F0: (v27 | (v27 / ((42) | 1)))}
            _ = v27
            _ = v28
        } else {
            v29 := GenS1{F0: -1}
            _ = v29
        }
        v30 := []int{-2147483648, fi26}
        _ = fi26
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
