package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 bool }
func genFn0(p0 int8, p1 uint) bool {
    for fi4 := 0; fi4 < 9; fi4++ {
        p0 -= (-(p0))
        _ = fi4
    }
    if false {
        for fi5 := 0; fi5 < 10; fi5++ {
            switch (fi5 - 0) {
            case 128:
                var v6 int16 = 218
                _ = v6
            default:
                p1 -= p1
            }
            if false {
                p0 += p0
            } else {
                for fi7 := 0; fi7 < 8; fi7++ {
                    p0 = p0
                    v8 := GenS1{F0: fi5}
                    _ = fi7
                    _ = v8
                }
                fmt.Println(p0, p1)
                var v9 bool = false
                _ = v9
            }
            fmt.Println(fi5, p1, p0)
            _ = fi5
        }
        var v10 int8 = p0
        fmt.Println(v10)
        _ = v10
    } else {
        switch 32768 {
        case 4294967295:
            p1 = p1
        case 0:
            p0 = p0
        default:
            v11 := GenS1{F0: 127}
            _ = v11
        }
    }
    _ = p0
    _ = p1
    return (!false)
}
func genFn1() int8 {
    v12 := []int{-189, 32767, 4294967295}
    if false {
        v13 := GenS1{F0: 1}
        _ = v13
    } else {
        var v14 int = 4294967295
        v15 := GenS1{F0: 127}
        if len(v12) > 0 {
            v12[0] = (v14 / (((v14 * -1)) | 1))
        }
        _ = v14
        _ = v15
    }
    if false {
        v16 := GenS2{F0: "hello", F1: 0}
        var v17 int64 = 42
        _ = v16
        _ = v17
    } else {
        var v18 bool = (1.5 > 0.1)
        var v19 int8 = 43
        for fi20 := 0; fi20 < 2; fi20++ {
            if false {
                v21 := GenS2{F0: "hello", F1: 0}
                _ = v21
            }
            var v22 float32 = float32(0.0)
            var v23 float64 = -2.25
            _ = fi20
            _ = v22
            _ = v23
        }
        _ = v18
        _ = v19
    }
    if false {
        var v24 uint64 = 42
        _ = v24
    } else {
        switch 127 {
        case -466:
            for fi25 := 0; fi25 < 11; fi25++ {
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                _ = fi25
            }
        case 874:
            var v26 int32 = 42
            _ = v26
        case 32767:
            if true {
                var v27 float32 = (0.0 * 0.1)
                var v28 string = "a"
                v29 := GenS1{F0: 32767}
                _ = v27
                _ = v28
                _ = v29
            }
        default:
            for fi30 := 0; fi30 < 9; fi30++ {
                var v31 uint32 = 65535
                v32 := GenS2{F0: ("" + "hello"), F1: v31}
                v33 := GenS3{F0: true}
                _ = fi30
                _ = v31
                _ = v32
                _ = v33
            }
        }
        switch 0 {
        case 32768:
            var v34 int32 = -727
            _ = v34
        default:
            var v35 uint32 = 65535
            _ = v35
        }
    }
    _ = v12
    return 42
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = -933
    var v3 int8 = 42
    var v36 float32 = 0.1
    fmt.Println(v2)
    fmt.Println(v2, v2)
    v36 += (v36 / (3.14159 / ((v36 * v36) * (((v36 + v36) + float32((1e10 * float64(1.5)))) - v36))))
    for fi37 := 0; fi37 < 13; fi37++ {
        v38 := []int{(-(fi37)), fi37, (^(fi37))}
        v2 = v2
        v39 := GenS1{F0: (-(fi37))}
        _ = fi37
        _ = v38
        _ = v39
    }
    v36 = v36
    v40 := []int{}
    var v41 float64 = (0.0 + 1e10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v40
    _ = v41
}
