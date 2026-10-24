package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 int8; F2 bool }
func genFn0(p0 int32) bool {
    for fi5 := 0; fi5 < 6; fi5++ {
        if (!((^(fi5)) == fi5)) {
            v6 := []int{fi5, fi5, (fi5 + fi5)}
            v6 = append(v6, 32768)
            _ = v6
        }
        _ = fi5
    }
    v7 := GenS1{F0: 1000}
    for fi8 := 0; fi8 < 19; fi8++ {
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 int32 = (p0 * (p0 << ((18446744073709551615) & 63)))
            v10 += v10
            v11 := GenS1{F0: -686}
            var v12 uint32 = 65535
            var v13 bool = (true && (65535 < 418))
            fmt.Println(v12, v10, v13)
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
            _ = v13
        }
        _ = fi8
    }
    p0 += (p0 + (p0 >> ((4294967295) & 63)))
    _ = p0
    _ = v7
    return true
}

func main() {
    v1 := []string{(("hello" + "\n\t") + "a"), ("\n\t" + "foo bar")}
    v2 := make(map[uint32]bool)
    var v3 string = ""
    v4 := make(map[int]string)
    v3 = v3
    if false {
        v3 = ("a" + v3)
        for fi14 := 0; fi14 < 11; fi14++ {
            var v15 int16 = 621
            if false {
                v15 = (v15 >> ((2) & 63))
            }
            v15 = v15
            _ = fi14
            _ = v15
        }
    }
    v3 = v3
    v16 := GenS1{F0: -1}
    v3 = (v3 + (v3 + v3))
    var v17 uint8 = 104
    v17 = (^(v17))
    v18 := GenS2{F0: "\n\t", F1: 100, F2: (4294967295 != 0)}
    if false {
        fmt.Println(v17, v17)
        v19 := GenS2{F0: (("" + v3) + v3), F1: 100, F2: true}
        _ = v19
    }
    v17 = v17
    if (!true) {
        if (65535 >= 0) {
            var v20 int16 = 214
            _ = v20
        } else {
            var v21 uint = 612
            _ = v21
        }
        var v22 int = 32768
        v23 := GenS1{F0: -69}
        _ = v22
        _ = v23
    }
    var v24 float64 = (1.5 - 0.0)
    v17 += 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
