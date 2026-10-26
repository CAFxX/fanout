package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint }
func genFn0(p0 int64, p1 uint16) int16 {
    var v6 uint16 = 1
    p0 = -257
    _ = p0
    _ = p1
    _ = v6
    return 32767
}

func main() {
    var v1 int8 = 127
    v2 := make(map[int]string)
    v3 := []string{"", "", (("\n\t" + "a") + "\n\t"), "\n\t"}
    var v4 int8 = v1
    v5 := []int{9223372036854775807, -264}
    v7 := v2[644]
    fmt.Println(v4, v1)
    for fi8 := 0; fi8 < 4; fi8++ {
        v4 -= v1
        _ = fi8
    }
    v7 = (("\n\t" + v7) + v7)
    v7 = ((v7 + v7) + "\n\t")
    v2[128] = v7
    var v9 float64 = 1e10
    if (904 > 808) {
        var v10 int = -386
        _ = v10
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    fmt.Println(v9, v7)
    for fi11 := 0; fi11 < 4; fi11++ {
        v3 = append(v3, "hello")
        _ = fi11
    }
    var v12 int64 = 9223372036854775807
    fmt.Println(v1)
    var v13 bool = true
    if false {
        switch -9223372036854775808 {
        case -9223372036854775808:
            v14 := GenS1{F0: 2, F1: 32767}
            _ = v14
        case 42:
            var v15 uint16 = 200
            _ = v15
        case -262:
            v4 += (-(v1))
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch "a" {
        case "":
            v16 := GenS2{F0: 1}
            _ = v16
        default:
            _ = len(v2)
        }
    } else {
        v17 := GenS2{F0: 18446744073709551615}
        _ = v17
    }
    var v18 int32 = -1
    v1 = v4
    fmt.Println(v7, v12, v13)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
