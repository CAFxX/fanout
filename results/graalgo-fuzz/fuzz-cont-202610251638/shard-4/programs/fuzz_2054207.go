package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int }
type GenS2 struct { F0 bool }
func genFn0(p0 int, p1 int32) int {
    var v4 string = "\n\t"
    v5 := GenS1{F0: 2, F1: 2, F2: p0}
    fmt.Println(v4, p0, p0)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 2147483647
}
func genFn1() int32 {
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := GenS1{F0: 0, F1: 65535, F2: fi6}
        var v8 int32 = -2147483648
        switch 4294967295 {
        case -2147483648:
            v9 := GenS1{F0: 900, F1: 788, F2: fi6}
            _ = v9
        case -173:
            var v10 int64 = 42
            _ = v10
        case 32768:
            var v11 uint64 = 42
            _ = v11
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    v12 := []int{4294967295, -153, 42}
    var v13 string = "a"
    _ = v12
    _ = v13
    return -1
}

func main() {
    v1 := []string{"", ("hello" + ("hello" + "a")), ("hello" + "foo bar"), "\n\t"}
    var v2 float32 = (3.14159 + 1.5)
    var v3 int32 = -1
    for fi14 := 0; fi14 < 6; fi14++ {
        if false {
            if true {
                v1 = append(v1, "hello")
                var v15 int8 = 0
                _ = v15
            }
        }
        var v16 int64 = -9223372036854775808
        if true {
            if len(v1) > 0 {
                v1[0] = (("" + ("\n\t" + "foo bar")) + "hello")
            }
            if len(v1) > 0 {
                v1[0] = "a"
            }
        }
        _ = fi14
        _ = v16
    }
    var v17 float64 = (0.0 / (0.1 * (1e10 / float64((1e10 / v2)))))
    v3 = (v3 & genFn1())
    var v18 float32 = -2.25
    var v19 float32 = 1.5
    v19 -= (3.14159 + v18)
    v2 += 1.5
    if true {
        for fi20 := 0; fi20 < 2; fi20++ {
            fmt.Println(v19)
            _ = fi20
        }
        for fi21 := 0; fi21 < 7; fi21++ {
            fmt.Println(v17)
            for fi22 := 0; fi22 < 17; fi22++ {
                v1 = append(v1, "foo bar")
                if false {
                    v18 = float32(v17)
                    var v23 uint16 = 0
                    if len(v1) > 0 {
                        v1[0] = ("" + "a")
                    }
                    _ = v23
                } else {
                    if true {
                        var v24 int64 = -461
                        _ = v24
                    }
                    var v25 float32 = ((v18 * (float32(v17) / v18)) + v19)
                    _ = v25
                }
                _ = fi22
            }
            for fi26 := 0; fi26 < 13; fi26++ {
                fmt.Println(fi21, fi21)
                _ = fi26
            }
            _ = fi21
        }
    }
    var v27 int8 = -75
    fmt.Println(v19)
    v28 := GenS2{F0: true}
    for fi29 := 0; fi29 < 11; fi29++ {
        v30 := GenS1{F0: 32767, F1: 83, F2: int(v3)}
        _ = fi29
        _ = v30
    }
    var v31 uint64 = 18446744073709551615
    v32 := GenS2{F0: true}
    if len(v1) > 0 {
        v1[0] = (("hello" + ("a" + "foo bar")) + ("\n\t" + "hello"))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v32
}
