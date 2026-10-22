package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 uint; F2 bool }
type GenS4 struct { F0 bool; F1 int16 }
func genFn0(p0 bool) uint32 {
    v6 := GenS1{F0: "\n\t", F1: 124, F2: 675}
    v7 := GenS1{F0: ("" + "\n\t"), F1: 2, F2: 0}
    for fi8 := 0; fi8 < 3; fi8++ {
        p0 = false
        _ = fi8
    }
    if true {
        for fi9 := 0; fi9 < 13; fi9++ {
            for fi10 := 0; fi10 < 16; fi10++ {
                p0 = (!p0)
                if (-2.25 >= 0.0) {
                    p0 = (false && true)
                    v11 := []int{fi10, fi10, (^(fi10)), fi9}
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    p0 = (!false)
                    fmt.Println(p0, fi10, fi9)
                    _ = v11
                }
                _ = fi10
            }
            fmt.Println(p0, fi9)
            _ = fi9
        }
        p0 = (!p0)
        v12 := []int{}
        _ = v12
    } else {
        v13 := GenS2{F0: 1}
        _ = v13
    }
    var v14 int32 = 0
    _ = p0
    _ = v14
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"foo bar"}
    v3 := []int{}
    var v4 string = "a"
    var v5 float64 = -2.25
    for fi15 := 0; fi15 < 1; fi15++ {
        v2 = append(v2, "a")
        for fi16 := 0; fi16 < 14; fi16++ {
            v2 = append(v2, (v4 + v4))
            _ = fi16
        }
        _ = fi15
    }
    var v17 int8 = 44
    if false {
        fmt.Println(v5)
    }
    v18 := GenS3{F0: (!true), F1: 4294967295, F2: (!(v4 < (("foo bar" + v4) + "\n\t")))}
    v19 := v1["hello"]
    if true {
        v4 = v4
    }
    if len(v3) > 0 {
        v3[0] = -1
    }
    _ = len(v3)
    var v20 int16 = 0
    v20 = (v20 / ((v20) | 1))
    var v21 int16 = -95
    _ = len(v3)
    var v22 uint16 = 1
    for fi23 := 0; fi23 < 18; fi23++ {
        v24 := v1[(((v4 + (v4 + (v4 + "hello"))) + v4) + (v4 + (v4 + v4)))]
        if (!true) {
            fmt.Println(v21, v22)
        } else {
            fmt.Println(v4, v24, v5)
            v25 := GenS4{F0: (true || true), F1: (-(v21))}
            switch fi23 {
            case -305:
                if true {
                    var v26 int8 = v17
                    var v27 string = v4
                    _ = v26
                    _ = v27
                }
            case 128:
                v1["hello"] = v19
            case 9223372036854775807:
                fmt.Println(v20)
            }
            _ = v25
        }
        _ = fi23
        _ = v24
    }
    v28 := GenS3{F0: (v19 <= 1), F1: 2, F2: true}
    var v29 uint16 = v22
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
