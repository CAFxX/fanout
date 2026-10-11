package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int32; F1 uint; F2 uint64 }
func genFn0(p0 string) int32 {
    p0 = (p0 + ((p0 + p0) + ("" + p0)))
    p0 = p0
    v7 := []int{}
    p0 = "hello"
    _ = p0
    _ = v7
    return 2147483647
}

func main() {
    v1 := []string{"\n\t", "", ("a" + "")}
    var v2 int16 = 0
    var v3 string = ("hello" + "foo bar")
    v4 := []uint32{2, 4294967295}
    var v5 int = 0
    var v6 uint = 18446744073709551615
    for fi8 := 0; fi8 < 13; fi8++ {
        v3 = "foo bar"
        _ = len(v1)
        var v9 uint32 = 2147483647
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: 4294967295, F1: (0.0 >= (1e10 + -2.25))}
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 uint32 = 601
        var v13 string = "a"
        if false {
            v12 = v12
            v2 += int16(v5)
            if false {
                fmt.Println(v3, fi11)
                if false {
                    var v14 float32 = (((-2.25 - 0.1) - (3.14159 + (((0.1 + 1.5) * ((float32((1.5 - -2.25)) * 3.14159) * 1e10)) * 0.1))) / -2.25)
                    _ = v14
                }
                for fi15 := 0; fi15 < 19; fi15++ {
                    v16 := GenS1{F0: (v12 + v12), F1: (0 < 89)}
                    _ = fi15
                    _ = v16
                }
            } else {
                var v17 uint64 = 4294967295
                switch fi11 {
                case -961:
                    fmt.Println(v6)
                case 1:
                    v1 = append(v1, (v3 + (v3 + v3)))
                case -324:
                    v18 := GenS1{F0: v12, F1: true}
                    _ = v18
                }
                fmt.Println(v5, v13)
                _ = v17
            }
        } else {
            var v19 uint64 = 0
            var v20 uint32 = (v12 * 1)
            var v21 int16 = v2
            _ = v19
            _ = v20
            _ = v21
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        if len(v4) > 0 {
            v4[0] = 720
        }
        _ = fi22
    }
    for fi23 := 0; fi23 < 14; fi23++ {
        if true {
            v2 = v2
            fmt.Println(v3)
            v6 -= v6
        } else {
            fmt.Println(v3, fi23)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        v24 := GenS1{F0: 917, F1: true}
        _ = fi23
        _ = v24
    }
    v25 := GenS1{F0: 42, F1: false}
    v3 = ("a" + "hello")
    for fi26 := 0; fi26 < 20; fi26++ {
        var v27 string = "hello"
        v28 := GenS2{F0: -1, F1: 776, F2: 294}
        v3 = v3
        _ = fi26
        _ = v27
        _ = v28
    }
    v5 -= v5
    var v29 int8 = 100
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
