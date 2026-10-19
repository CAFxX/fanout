package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 bool }
func genFn0() uint {
    var v4 bool = ((-9223372036854775808 >= 32768) && ((!false) && ((0 > 9223372036854775807) || ("a" >= ("a" + "foo bar")))))
    switch ("foo bar" + "hello") {
    case "foo bar":
        fmt.Println(v4)
    case "a":
        for fi5 := 0; fi5 < 20; fi5++ {
            var v6 float64 = (0.1 - 1e10)
            _ = fi5
            _ = v6
        }
    case "\n\t":
        if (1 >= -750) {
            for fi7 := 0; fi7 < 4; fi7++ {
                v4 = true
                var v8 int = fi7
                for fi9 := 0; fi9 < 2; fi9++ {
                    switch 579 {
                    case 42:
                        v10 := GenS1{F0: (42 != 9223372036854775807), F1: 2}
                        _ = v10
                    case 9223372036854775807:
                        fmt.Println(v8)
                    default:
                        v11 := GenS2{F0: 0, F1: 2}
                        _ = v11
                    }
                    var v12 uint64 = 18446744073709551615
                    v8 = v8
                    _ = fi9
                    _ = v12
                }
                _ = fi7
                _ = v8
            }
            for fi13 := 0; fi13 < 4; fi13++ {
                switch (fi13 | fi13) {
                case 1:
                    var v14 float32 = 0.1
                    _ = v14
                case 32767:
                    if v4 {
                        var v15 int32 = 0
                        v16 := []int{}
                        var v17 int8 = -72
                        _ = v15
                        _ = v16
                        _ = v17
                    } else {
                        fmt.Println(fi13)
                    }
                case -205:
                    for fi18 := 0; fi18 < 19; fi18++ {
                        v19 := GenS2{F0: 127, F1: 37}
                        _ = fi18
                        _ = v19
                    }
                }
                _ = fi13
            }
        }
    }
    var v20 float32 = ((3.14159 - -2.25) - (1e10 - (float32(-2.25) * 1.5)))
    for fi21 := 0; fi21 < 5; fi21++ {
        v4 = (2 > 2)
        switch (^(fi21)) {
        case 4294967295:
            v22 := GenS1{F0: true, F1: 255}
            _ = v22
        case -805:
            if v4 {
                v4 = false
                if v4 {
                    fmt.Println(v20, v4, v4)
                    var v23 int = fi21
                    _ = v23
                }
                v24 := []int{}
                _ = v24
            }
        default:
            if v4 {
                v25 := GenS1{F0: v4, F1: 139}
                _ = v25
            } else {
                fmt.Println(v4, v20, fi21)
            }
        }
        var v26 uint16 = 1
        _ = fi21
        _ = v26
    }
    _ = v20
    _ = v4
    return 797
}

func main() {
    v1 := []int{-423, 786}
    v2 := make(map[int]string)
    v3 := []string{}
    var v27 float32 = float32(-2.25)
    v27 = float32(-2.25)
    var v28 int32 = 42
    v29 := GenS3{F0: true}
    if false {
        if ((!(4294967295 >= 207)) && true) {
            var v30 int64 = -20
            var v31 string = "\n\t"
            _ = v30
            _ = v31
        }
        v27 = 1e10
    }
    v32 := v2[-1]
    _ = len(v3)
    v32 = v32
    v32 = v32
    var v33 int8 = 100
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
}
