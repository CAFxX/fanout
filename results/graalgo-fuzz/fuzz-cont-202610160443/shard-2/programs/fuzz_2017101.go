package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 uint; F1 uint }
func genFn0(p0 bool) int16 {
    for fi6 := 0; fi6 < 19; fi6++ {
        if true {
            var v7 string = (("" + "") + "hello")
            _ = v7
        }
        _ = fi6
    }
    p0 = p0
    for fi8 := 0; fi8 < 6; fi8++ {
        for fi9 := 0; fi9 < 4; fi9++ {
            fmt.Println(fi9, fi9)
            v10 := GenS1{F0: 42, F1: 42}
            _ = fi9
            _ = v10
        }
        var v11 uint32 = 414
        v12 := []int{fi8}
        _ = fi8
        _ = v11
        _ = v12
    }
    switch -9223372036854775808 {
    case 532:
        var v13 bool = p0
        _ = v13
    case 1:
        v14 := GenS2{F0: 751, F1: 2}
        _ = v14
    case 4294967295:
        v15 := []int{}
        _ = v15
    default:
        v16 := GenS2{F0: 2, F1: 255}
        _ = v16
    }
    p0 = true
    _ = p0
    return 1000
}

func main() {
    v1 := make(map[string]bool)
    var v2 string = "hello"
    v3 := make(map[int]bool)
    v4 := make(map[int]uint32)
    v5 := []string{v2, (v2 + "hello")}
    v17 := GenS2{F0: 852, F1: 255}
    delete(v3, 459)
    var v18 int16 = genFn0(false)
    var v19 float64 = 1e10
    v2 = (v2 + "hello")
    if (42 >= 42) {
        for fi20 := 0; fi20 < 15; fi20++ {
            var v21 string = ""
            _ = len(v3)
            _ = fi20
            _ = v21
        }
        var v22 int16 = v18
        _ = v22
    } else {
        if false {
            if false {
                v23 := GenS3{F0: 42, F1: 1}
                var v24 int64 = 525
                if len(v5) > 0 {
                    v5[0] = v2
                }
                _ = v23
                _ = v24
            }
        } else {
            if false {
                v2 = (v2 + ("foo bar" + v2))
                v25 := GenS2{F0: 514, F1: 127}
                _ = v25
            } else {
                _ = len(v3)
                if false {
                    for fi26 := 0; fi26 < 11; fi26++ {
                        if false {
                            var v27 uint16 = 65535
                            _ = v27
                        }
                        var v28 float32 = 1e10
                        _ = fi26
                        _ = v28
                    }
                } else {
                    fmt.Println(v18)
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
            }
            fmt.Println(v2, v18)
            switch ("foo bar" + v2) {
            case "\n\t":
                var v29 float64 = 1e10
                _ = v29
            case "hello":
                var v30 uint = 2
                _ = v30
            default:
                v19 -= v19
            }
        }
        for fi31 := 0; fi31 < 4; fi31++ {
            fmt.Println(v19)
            v18 = (^(v18))
            _ = fi31
        }
        if true {
            v18 = v18
        }
    }
    v18 = v18
    v32 := GenS3{F0: 18446744073709551615, F1: 895}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
