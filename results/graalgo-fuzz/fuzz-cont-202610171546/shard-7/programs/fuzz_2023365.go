package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int8 }
func genFn0(p0 int, p1 int16, p2 uint64) uint16 {
    p0 += p0
    switch p0 {
    case 32768:
        if ((p0 >> ((1) & 63)) > p0) {
            v7 := GenS1{F0: p2, F1: 1}
            v8 := []int{(p0 ^ p0), (p0 << ((136) & 63)), (p0 ^ p0), (p0 ^ p0)}
            _ = v7
            _ = v8
        } else {
            for fi9 := 0; fi9 < 19; fi9++ {
                if (249 != 1) {
                    if ((((42 >= 20) && (!false)) && false) || (1.5 == 3.14159)) {
                        var v10 uint8 = 127
                        _ = v10
                    }
                    var v11 uint8 = 129
                    var v12 int64 = 393
                    _ = v11
                    _ = v12
                } else {
                    var v13 string = ""
                    v14 := GenS1{F0: 18446744073709551615, F1: 1}
                    p0 -= (^(p0))
                    _ = v13
                    _ = v14
                }
                _ = fi9
            }
            p1 += (p1 | -838)
        }
    }
    v15 := GenS2{F0: -33}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 33
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]uint32)
    var v3 int32 = 0
    v4 := make(map[string]int)
    var v5 bool = (243 == 102)
    v6 := []uint32{4294967295}
    fmt.Println(v5, v3)
    v16 := v4["\n\t"]
    v5 = (v5 || false)
    var v17 float64 = (float64(1.5) * -2.25)
    fmt.Println(v3, v5)
    var v18 string = (("foo bar" + ("foo bar" + "\n\t")) + "\n\t")
    fmt.Println(v16)
    v16 = (v16 & v16)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
