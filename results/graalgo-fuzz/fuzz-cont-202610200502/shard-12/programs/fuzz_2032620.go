package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 int }
func genFn0(p0 int8, p1 int64) int16 {
    switch -1 {
    case 32768:
        fmt.Println(p1, p1)
    default:
        switch 9223372036854775807 {
        case 2147483647:
            v5 := []int{}
            _ = v5
        }
    }
    var v6 uint64 = 2
    if true {
        var v7 uint16 = 955
        v8 := []int{-2147483648}
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    return 1
}

func main() {
    var v1 uint16 = 244
    v2 := []string{((("foo bar" + (("hello" + "a") + "foo bar")) + "foo bar") + ""), (("" + "") + "foo bar"), "\n\t", "a"}
    var v3 int = 1
    var v4 uint64 = 131
    for fi9 := 0; fi9 < 16; fi9++ {
        switch 1 {
        case 848:
            v10 := GenS1{F0: -1}
            _ = v10
        case 0:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            v3 = -9223372036854775808
        }
        _ = fi9
    }
    v4 -= v4
    fmt.Println(v1, v3)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v11 int64 = 1
    fmt.Println(v4)
    switch "\n\t" {
    case "a":
        for fi12 := 0; fi12 < 1; fi12++ {
            var v13 float32 = 1.5
            var v14 uint = 426
            v15 := GenS2{F0: "\n\t"}
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 int = 4294967295
        if (!true) {
            if ((genFn0(0, (v11 + v11)) >= -32768) && true) {
                fmt.Println(v3)
                v4 -= 0
                var v18 uint64 = (^(v4))
                _ = v18
            } else {
                if false {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    fmt.Println(v17, v3)
                }
                var v19 string = "\n\t"
                for fi20 := 0; fi20 < 18; fi20++ {
                    fmt.Println(v1, v11, v1)
                    _ = fi20
                }
                _ = v19
            }
            v21 := GenS3{F0: genFn0(1, (v11 << ((4294967295) & 63))), F1: (v3 >> ((0) & 63))}
            _ = v21
        }
        _ = fi16
        _ = v17
    }
    v3 = (-(v3))
    fmt.Println(v1, v3, v11)
    fmt.Println(v11, v1)
    for fi22 := 0; fi22 < 5; fi22++ {
        for fi23 := 0; fi23 < 12; fi23++ {
            v2 = append(v2, "hello")
            _ = fi23
        }
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
