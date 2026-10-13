package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint32, p1 int16) uint16 {
    switch "foo bar" {
    case "hello":
        v5 := []int{-453}
        _ = v5
    case "foo bar":
        p0 -= 0
    }
    var v6 int = -811
    for fi7 := 0; fi7 < 19; fi7++ {
        p1 = (p1 | (p1 + (p1 & p1)))
        var v8 uint64 = 2
        var v9 int16 = p1
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    return 255
}
func genFn1() int32 {
    v10 := []int{-9223372036854775808, 124, 2147483647, -1}
    var v11 uint64 = 792
    v12 := GenS1{F0: 32767}
    var v13 string = ("\n\t" + "foo bar")
    v11 += (v11 | (v11 * (v11 << ((4294967295) & 63))))
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return -330
}

func main() {
    var v1 int = 101
    var v2 uint64 = 2
    v3 := make(map[string]string)
    var v4 int16 = 1
    for fi14 := 0; fi14 < 17; fi14++ {
        if (177 > 0) {
            fmt.Println(v1, fi14, fi14)
        } else {
            if false {
                var v15 int64 = 1
                fmt.Println(v2, fi14)
                for fi16 := 0; fi16 < 8; fi16++ {
                    var v17 uint64 = v2
                    _ = fi16
                    _ = v17
                }
                _ = v15
            } else {
                v18 := GenS2{F0: 604}
                _ = v18
            }
            v19 := []int{fi14, (^(fi14)), -1, 2147483647}
            _ = v19
        }
        var v20 float32 = -2.25
        for fi21 := 0; fi21 < 3; fi21++ {
            v20 -= v20
            v2 = 0
            fmt.Println(v4, v4, v1)
            var v22 uint32 = 4294967295
            _ = fi21
            _ = v22
        }
        _ = fi14
        _ = v20
    }
    v23 := GenS1{F0: genFn0(1, v4)}
    switch -231 {
    case -362:
        v24 := GenS1{F0: 76}
        _ = v24
    case 42:
        v3[((((("" + ("" + "\n\t")) + "") + "\n\t") + "") + ("foo bar" + ""))] = ("\n\t" + ("a" + "hello"))
    case 0:
        var v25 int16 = (v4 + 0)
        _ = v25
    default:
        v2 = (v2 / (((v2 / ((v2) | 1))) | 1))
    }
    v26 := []int{1, (v1 | (v1 ^ v1))}
    v27 := GenS1{F0: 255}
    v4 += v4
    var v28 string = "a"
    if true {
        v2 -= v2
        v29 := v3["\n\t"]
        switch 0 {
        case 9223372036854775807:
            fmt.Println(v4)
        case 552:
            var v30 string = "a"
            _ = v30
        case -797:
            fmt.Println(v2, v2)
        }
        _ = v29
    }
    var v31 int64 = -666
    fmt.Println(v28, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
}
