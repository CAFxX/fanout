package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 int64; F2 string }
func genFn0(p0 uint32, p1 string, p2 string) string {
    for fi7 := 0; fi7 < 8; fi7++ {
        if (true || (!false)) {
            v8 := GenS1{F0: -78, F1: 726}
            _ = v8
        }
        _ = fi7
    }
    var v9 uint16 = 929
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return "foo bar"
}

func main() {
    var v1 uint16 = 241
    var v2 int8 = -1
    var v3 uint32 = 2147483648
    v4 := make(map[string]int)
    v5 := []int{32768, 42, -199, 42}
    var v6 int64 = -1
    v6 += v6
    switch -349 {
    case 275:
        if (!(18446744073709551615 == 4294967295)) {
            v2 += (-(v2))
        } else {
            v10 := GenS1{F0: 1, F1: 534}
            _ = v10
        }
    case 2147483647:
        v1 = v1
    case 373:
        var v11 float64 = 3.14159
        _ = v11
    }
    v12 := GenS2{F0: 2}
    var v13 uint32 = uint32(v1)
    v14 := GenS3{F0: ((v13 * v3) > 820), F1: v6, F2: ("hello" + ("" + (("\n\t" + (("a" + ("" + "a")) + ((("foo bar" + "\n\t") + "\n\t") + ("" + "hello")))) + "a")))}
    v3 -= v3
    var v15 float32 = 3.14159
    v6 = (^(v6))
    v16 := GenS2{F0: 255}
    _ = len(v5)
    var v17 string = (("\n\t" + "") + "hello")
    var v18 uint16 = v1
    var v19 int = 103
    v20 := v4[(v17 + v17)]
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
