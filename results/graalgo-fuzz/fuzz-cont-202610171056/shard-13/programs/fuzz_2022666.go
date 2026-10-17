package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int8, p1 int8) uint {
    var v6 int32 = -2147483648
    var v7 uint8 = 85
    v8 := []int{0, -2147483648, -9223372036854775808}
    _ = len(v8)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}
func genFn1(p0 int64, p1 uint8, p2 uint32) int8 {
    fmt.Println(p0)
    switch 128 {
    case -694:
        p2 -= (p2 * 492)
    case 42:
        fmt.Println(p0, p2)
    case -9223372036854775808:
        var v9 uint8 = (p1 - (p1 >> ((0) & 63)))
        _ = v9
    default:
        fmt.Println(p1)
    }
    if false {
        var v10 uint16 = 42
        fmt.Println(p0, p0)
        p0 = -9223372036854775808
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    v1 := make(map[string]int)
    var v2 bool = (!((42 == 0) || ((false || false) || (!(false && ((!(true || true)) && (0 > 65535)))))))
    v3 := []string{"foo bar"}
    v4 := make(map[int]string)
    v5 := make(map[int]uint32)
    fmt.Println(v2)
    v4[0] = "foo bar"
    v2 = (750 >= 2147483647)
    v2 = v2
    v11 := GenS1{F0: -353}
    v2 = (42 < 42)
    for fi12 := 0; fi12 < 8; fi12++ {
        _ = len(v3)
        _ = fi12
    }
    v2 = v2
    if false {
        var v13 uint32 = 2147483648
        delete(v5, -263)
        delete(v1, ("" + "foo bar"))
        _ = v13
    } else {
        var v14 float32 = 3.14159
        var v15 float64 = (float64(v14) * float64(float32(-2.25)))
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 20; fi16++ {
        _ = len(v3)
        if false {
            v17 := GenS1{F0: 414}
            _ = v17
        } else {
            switch 636 {
            case 128:
                var v18 float32 = float32(3.14159)
                _ = v18
            }
        }
        _ = fi16
    }
    if (genFn0(127, 127) != 18446744073709551615) {
        delete(v5, 42)
        var v19 uint32 = 2147483648
        _ = v19
    }
    v2 = ("" < "foo bar")
    if (0 > 1) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v20 int64 = 1
        v20 = v20
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
