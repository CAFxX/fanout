package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 bool; F2 string }
func genFn0(p0 uint8, p1 int, p2 uint16) uint8 {
    p1 -= 2147483647
    for fi4 := 0; fi4 < 15; fi4++ {
        p2 -= p2
        fmt.Println(p1, p0, fi4)
        _ = fi4
    }
    var v5 int32 = 443
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p0
}

func main() {
    var v1 int64 = 176
    var v2 int = -2147483648
    var v3 int32 = -1
    switch 9223372036854775807 {
    case 127:
        v6 := GenS1{F0: "hello"}
        _ = v6
    case -37:
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := []int{fi7, v2}
            v3 += 738
            var v9 int32 = v3
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    v10 := []int{v2}
    v11 := GenS1{F0: "hello"}
    var v12 int = v2
    var v13 float32 = 1.5
    var v14 float32 = (v13 - v13)
    v2 += (-(v2))
    if true {
        if true {
            for fi15 := 0; fi15 < 10; fi15++ {
                _ = len(v10)
                _ = fi15
            }
            var v16 uint = 2
            if (!((true || false) || true)) {
                v12 = (v12 >> ((v16) & 63))
                v2 = v2
                v12 += v2
            }
            _ = v16
        } else {
            v17 := GenS1{F0: "a"}
            if (1 >= v1) {
                if true {
                    v3 = -2147483648
                    var v18 int64 = -9223372036854775808
                    _ = v18
                } else {
                    v3 += v3
                }
            }
            _ = v17
        }
        if true {
            var v19 string = ("a" + ((((("a" + "\n\t") + ("\n\t" + "")) + "") + "") + ("a" + "a")))
            v20 := GenS2{F0: 42, F1: false, F2: ((v19 + ("foo bar" + ("" + v19))) + ("a" + "foo bar"))}
            _ = v19
            _ = v20
        } else {
            var v21 uint64 = 0
            v2 -= v2
            _ = v21
        }
        switch 4294967295 {
        case -516:
            v2 -= (v12 | (-(v12)))
        case -509:
            var v22 int32 = (v3 * v3)
            _ = v22
        case 4294967295:
            var v23 int32 = (v3 - v3)
            _ = v23
        default:
            if len(v10) > 0 {
                v10[0] = 187
            }
        }
    } else {
        fmt.Println(v14)
        v14 += ((v14 + v13) - ((3.14159 - v14) - (v13 / v13)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
