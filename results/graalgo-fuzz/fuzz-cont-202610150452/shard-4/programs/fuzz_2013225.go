package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int16, p1 int32) bool {
    var v4 string = ((("a" + "\n\t") + ("foo bar" + "")) + "foo bar")
    for fi5 := 0; fi5 < 17; fi5++ {
        v6 := GenS1{F0: p0}
        p1 += -960
        if (57 <= -128) {
            v7 := []int{(fi5 * fi5)}
            var v8 float64 = float64(0.1)
            for fi9 := 0; fi9 < 2; fi9++ {
                v10 := GenS1{F0: p0}
                _ = fi9
                _ = v10
            }
            _ = v7
            _ = v8
        } else {
            v4 = ""
            var v11 string = (("hello" + v4) + (v4 + ((("" + v4) + v4) + (v4 + (((v4 + (v4 + (v4 + ("hello" + ("" + (v4 + (v4 + v4))))))) + (v4 + (v4 + v4))) + "\n\t")))))
            var v12 float64 = 0.0
            _ = v11
            _ = v12
        }
        _ = fi5
        _ = v6
    }
    var v13 int64 = 988
    var v14 uint64 = 1
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v4
    return (!false)
}

func main() {
    var v1 int32 = 2147483647
    var v2 uint64 = 4294967295
    var v3 uint64 = (v2 - v2)
    if ((1 == -388) || (true || (!(902 > 761)))) {
        v3 = v2
        v15 := GenS1{F0: 246}
        v1 -= (v1 >> ((65535) & 63))
        _ = v15
    }
    v16 := []int{}
    var v17 bool = (!(-2147483648 <= -1))
    var v18 uint8 = 0
    if (100 >= 67) {
        var v19 int32 = (v1 & (v1 * (v1 + v1)))
        var v20 float64 = (1.5 * 1.5)
        var v21 string = "\n\t"
        _ = v19
        _ = v20
        _ = v21
    }
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    for fi22 := 0; fi22 < 17; fi22++ {
        v16 = append(v16, 32767)
        var v23 bool = ((v1 + (v1 | (v1 ^ (v1 * (-(v1)))))) == int32(fi22))
        _ = fi22
        _ = v23
    }
    var v24 uint8 = v18
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v3
}
