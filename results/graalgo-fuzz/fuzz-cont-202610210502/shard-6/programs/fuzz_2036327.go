package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int) int32 {
    var v5 uint64 = 0
    var v6 uint16 = 667
    switch p0 {
    case -855:
        var v7 bool = (127 > 2)
        _ = v7
    default:
        if false {
            v8 := GenS1{F0: 0}
            v9 := GenS1{F0: 0}
            _ = v8
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        v11 := []int{p0, (-(fi10)), (p0 << ((255) & 63)), (fi10 ^ (fi10 >> ((17) & 63)))}
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = v5
    _ = v6
    return -1
}
func genFn1() int32 {
    var v12 uint8 = 1
    v12 -= v12
    _ = v12
    return -54
}

func main() {
    v1 := []string{"", "foo bar", "a"}
    var v2 int16 = 1
    v3 := make(map[string]int)
    v4 := make(map[int]string)
    var v13 uint64 = 7
    var v14 int64 = 1
    var v15 int = -1
    var v16 float64 = 0.1
    v1 = append(v1, "foo bar")
    var v17 int = (v15 * 127)
    for fi18 := 0; fi18 < 5; fi18++ {
        v1 = append(v1, ("" + ((("" + "a") + "a") + ("" + "hello"))))
        var v19 float64 = v16
        _ = fi18
        _ = v19
    }
    v20 := v3[""]
    if ((v14 | (v14 ^ v14)) <= 0) {
        for fi21 := 0; fi21 < 18; fi21++ {
            v22 := GenS1{F0: uint(v13)}
            v23 := GenS1{F0: 42}
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = len(v3)
    }
    v24 := v4[(v17 % ((v20) | 1))]
    v13 = (v13 >> ((667) & 63))
    switch -834 {
    case -2147483648:
        var v25 string = v24
        _ = v25
    case 32768:
        if ((v20 - (v17 >> ((447) & 63))) < 128) {
            v26 := GenS1{F0: 42}
            _ = v26
        }
    case 127:
        var v27 int8 = -1
        _ = v27
    }
    fmt.Println(v17, v16, v24)
    fmt.Println(v13, v17)
    v2 = v2
    var v28 uint8 = 2
    var v29 bool = true
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
