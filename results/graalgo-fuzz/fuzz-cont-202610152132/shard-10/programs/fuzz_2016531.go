package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint16; F1 int8; F2 int32 }
func genFn0(p0 int8, p1 uint, p2 string) bool {
    v6 := []int{2147483647, -3, -9223372036854775808, 0}
    v6 = append(v6, 539)
    if true {
        var v7 int = 2147483647
        fmt.Println(p1, p2, p0)
        _ = v7
    } else {
        fmt.Println(p0)
        var v8 uint64 = 1
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return true
}

func main() {
    v1 := []string{"", "", "\n\t", ""}
    var v2 int = 42
    v3 := make(map[string]string)
    v4 := make(map[string]bool)
    v5 := []int{(v2 >> ((4294967295) & 63)), (v2 + (v2 << ((524) & 63)))}
    var v9 bool = (!((1.5 * float64(0.0)) == -2.25))
    v10 := GenS1{F0: -859, F1: ((v2 + v2) > 127), F2: 255}
    for fi11 := 0; fi11 < 12; fi11++ {
        v1 = append(v1, (("hello" + "") + "a"))
        _ = fi11
    }
    fmt.Println(v9, v2)
    v12 := GenS1{F0: -9223372036854775808, F1: v9, F2: 2}
    var v13 uint = 836
    if ((!false) && (v9 && v9)) {
        v13 = 720
        delete(v4, ("a" + ("a" + (("hello" + "a") + ("\n\t" + "hello")))))
        v13 += (v13 / ((v13) | 1))
    } else {
        v13 += 42
        v2 += -2147483648
        if true {
            switch ("foo bar" + "a") {
            case "":
                v2 = v2
            case "a":
                v14 := GenS2{F0: 304, F1: 127, F2: 2147483647}
                _ = v14
            default:
                v2 = -1
            }
        }
    }
    fmt.Println(v13, v2)
    var v15 bool = (false || v9)
    v16 := GenS2{F0: 42, F1: 75, F2: -588}
    v17 := GenS2{F0: 65535, F1: 77, F2: 42}
    v18 := GenS2{F0: 2, F1: 42, F2: 367}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (!(9223372036854775807 > -397)) {
        for fi19 := 0; fi19 < 4; fi19++ {
            var v20 int = v2
            _ = fi19
            _ = v20
        }
        v9 = (!((-339 < 9223372036854775807) || (1.5 <= -2.25)))
        for fi21 := 0; fi21 < 15; fi21++ {
            fmt.Println(v13, v9, v9)
            _ = len(v4)
            _ = fi21
        }
    } else {
        for fi22 := 0; fi22 < 3; fi22++ {
            var v23 int = (fi22 % (((^(v2))) | 1))
            _ = fi22
            _ = v23
        }
        v2 += v2
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
