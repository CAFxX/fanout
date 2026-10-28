package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 string; F1 bool }
func genFn0(p0 uint64) int {
    var v6 string = "a"
    v6 = v6
    _ = p0
    _ = v6
    return 127
}
func genFn1() bool {
    v7 := []int{127, 127, 2147483647}
    var v8 int8 = 0
    _ = v7
    _ = v8
    return (!false)
}

func main() {
    var v1 uint32 = 953
    v2 := []uint32{2147483648, v1, v1}
    v3 := []int{782, 2147483647, 32767}
    var v4 uint = 4294967295
    v5 := []int{0, -2147483648, 32767}
    if (false && true) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v9 := GenS1{F0: 0, F1: v1}
        switch -701 {
        case 127:
            fmt.Println(v4, v4)
        case 32768:
            var v10 uint8 = 127
            _ = v10
        default:
            switch 4294967295 {
            case 127:
                if (976 != genFn0(18446744073709551615)) {
                    var v11 uint16 = 0
                    _ = v11
                } else {
                    var v12 uint16 = 443
                    v13 := GenS1{F0: (v1 - 2), F1: (v1 % (((v1 + uint32(v12))) | 1))}
                    _ = v12
                    _ = v13
                }
            case 42:
                for fi14 := 0; fi14 < 16; fi14++ {
                    v4 = v4
                    _ = fi14
                }
            }
        }
        _ = v9
    }
    v15 := GenS1{F0: v1, F1: v1}
    var v16 uint = (v4 * (v4 >> ((24) & 63)))
    v17 := GenS2{F0: ("" + ("" + "a")), F1: (true && false)}
    for fi18 := 0; fi18 < 18; fi18++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi18
    }
    v4 = (v4 + 4294967295)
    v16 = 18446744073709551615
    v1 -= (v1 - (v1 >> ((v16) & 63)))
    for fi19 := 0; fi19 < 9; fi19++ {
        var v20 uint64 = 4294967295
        v2 = append(v2, (v1 * v1))
        _ = fi19
        _ = v20
    }
    var v21 int64 = -1
    v1 += 641
    var v22 uint64 = 18446744073709551615
    if len(v5) > 0 {
        v5[0] = 1
    }
    fmt.Println(v16, v21)
    v22 = v22
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
