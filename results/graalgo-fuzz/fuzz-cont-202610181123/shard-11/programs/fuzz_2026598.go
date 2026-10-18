package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint, p1 int16, p2 int16) uint8 {
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: 127, F1: 42, F2: 455}
        p2 += 0
        _ = fi7
        _ = v8
    }
    if false {
        var v9 uint64 = 42
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1() uint16 {
    var v10 uint64 = 98
    for fi11 := 0; fi11 < 7; fi11++ {
        v10 = 95
        if true {
            if true {
                v12 := []int{fi11}
                _ = v12
            }
            v10 += (v10 & (v10 / (((^(v10))) | 1)))
        }
        _ = fi11
    }
    var v13 int32 = 0
    v13 = (v13 % ((1) | 1))
    _ = v10
    _ = v13
    return 65535
}

func main() {
    var v1 uint16 = 1
    v2 := []string{}
    var v3 uint = 2
    v4 := make(map[string]int)
    var v5 string = "hello"
    v6 := []int{401, -1, 0, -1}
    switch ("hello" + ("" + v5)) {
    case "a":
        delete(v4, "hello")
    case "":
        fmt.Println(v1, v1, v1)
    case "foo bar":
        var v14 uint8 = 1
        _ = v14
    }
    v15 := GenS1{F0: 32767, F1: 502, F2: 42}
    for fi16 := 0; fi16 < 1; fi16++ {
        var v17 int64 = 0
        for fi18 := 0; fi18 < 8; fi18++ {
            v17 = v17
            _ = fi18
        }
        _ = fi16
        _ = v17
    }
    var v19 float64 = (((0.0 + ((float64((3.14159 * (0.1 * (((-2.25 / (1e10 + -2.25)) * (1e10 - ((3.14159 * (((1.5 * 3.14159) + 0.1) / (0.1 + 0.1))) * -2.25))) + 0.0)))) - 1.5) + 1.5)) + float64((1.5 * (1.5 + 0.1)))) * 0.1)
    var v20 string = v5
    if true {
        v20 = ("" + ((((v20 + (v5 + "\n\t")) + "hello") + v5) + "hello"))
        v5 = "a"
        for fi21 := 0; fi21 < 20; fi21++ {
            if (2147483647 >= 437) {
                for fi22 := 0; fi22 < 14; fi22++ {
                    v4[(v20 + "hello")] = fi22
                    var v23 uint32 = 0
                    _ = fi22
                    _ = v23
                }
                v5 = ((v20 + v5) + "a")
                _ = len(v2)
            } else {
                var v24 int16 = 0
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
                v25 := GenS2{F0: 894}
                _ = v24
                _ = v25
            }
            fmt.Println(v1)
            for fi26 := 0; fi26 < 12; fi26++ {
                v20 = (("foo bar" + (v5 + v5)) + (((v20 + (("a" + v20) + "hello")) + "hello") + "a"))
                _ = fi26
            }
            _ = fi21
        }
    }
    v3 += v3
    v27 := GenS1{F0: 1, F1: -1, F2: 1}
    if true {
        fmt.Println(v3, v3)
        var v28 string = ((("a" + "hello") + v5) + v5)
        _ = v28
    } else {
        v20 = ("" + v20)
    }
    v5 = "foo bar"
    v29 := GenS3{F0: 127}
    v20 = ""
    for fi30 := 0; fi30 < 6; fi30++ {
        v31 := GenS3{F0: 128}
        for fi32 := 0; fi32 < 12; fi32++ {
            v33 := v4[v20]
            fmt.Println(fi32, v19, v20)
            _ = fi32
            _ = v33
        }
        _ = fi30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
