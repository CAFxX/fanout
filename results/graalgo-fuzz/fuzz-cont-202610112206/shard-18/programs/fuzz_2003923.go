package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 bool; F2 bool }
type GenS3 struct { F0 uint64; F1 uint16 }
func genFn0(p0 uint64, p1 int16, p2 uint8) uint64 {
    var v4 int32 = -2147483648
    v5 := GenS1{F0: 65535}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 4294967295
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int = 128
    var v3 float32 = float32((1.5 / 1e10))
    for fi6 := 0; fi6 < 8; fi6++ {
        for fi7 := 0; fi7 < 8; fi7++ {
            fmt.Println(v3, v3, v3)
            _ = fi7
        }
        for fi8 := 0; fi8 < 10; fi8++ {
            v9 := GenS2{F0: (true && (-9223372036854775808 == 42)), F1: (!true), F2: true}
            var v10 bool = (!false)
            fmt.Println(fi8)
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi6
    }
    v11 := GenS2{F0: (3.14159 >= float64(((v3 / v3) / v3))), F1: true, F2: (-2147483648 != -1)}
    var v12 bool = ((!true) && (1 >= -803))
    fmt.Println(v3, v3, v2)
    if true {
        var v13 int16 = 1000
        switch 0 {
        case -9223372036854775808:
            var v14 uint32 = 2147483648
            _ = v14
        case 32767:
            fmt.Println(v3, v12)
        }
        _ = v13
    }
    if v12 {
        _ = len(v1)
        var v15 float32 = (v3 - v3)
        _ = v15
    } else {
        v12 = false
        v1[("a" + ("" + "\n\t"))] = 512
        v12 = (1000 >= 840)
    }
    v16 := GenS2{F0: (314 >= (v2 & (v2 + v2))), F1: v12, F2: (100 > 100)}
    if v12 {
        var v17 int8 = 33
        for fi18 := 0; fi18 < 4; fi18++ {
            if true {
                v19 := GenS1{F0: 1}
                var v20 uint32 = 167
                _ = v19
                _ = v20
            } else {
                if v12 {
                    v21 := GenS2{F0: (false && v12), F1: ((v17 >= (-(v17))) || (v12 && true)), F2: v12}
                    var v22 uint = 4294967295
                    v3 += v3
                    _ = v21
                    _ = v22
                }
                v23 := []int{(fi18 - -56)}
                v17 = v17
                _ = v23
            }
            v3 += v3
            _ = fi18
        }
        v12 = ((933 != 65535) && true)
        _ = v17
    } else {
        v2 = v2
        fmt.Println(v12)
    }
    if false {
        delete(v1, "\n\t")
        v24 := GenS3{F0: 2, F1: 1}
        _ = v24
    } else {
        fmt.Println(v2, v3)
    }
    v2 = (v2 | (v2 * v2))
    var v25 int8 = 127
    for fi26 := 0; fi26 < 7; fi26++ {
        if v12 {
            v27 := []int{}
            _ = v27
        } else {
            v28 := v1["hello"]
            var v29 int = (fi26 - 32767)
            for fi30 := 0; fi30 < 14; fi30++ {
                fmt.Println(fi30, v29, fi26)
                v3 -= -2.25
                _ = fi30
            }
            _ = v28
            _ = v29
        }
        var v31 int = (v2 & fi26)
        _ = fi26
        _ = v31
    }
    var v32 string = ("foo bar" + (("\n\t" + ("foo bar" + "")) + (("a" + "a") + "foo bar")))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v25
    _ = v3
    _ = v32
}
