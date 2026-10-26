package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int8 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 uint }
func genFn0(p0 uint64, p1 string) uint {
    for fi4 := 0; fi4 < 1; fi4++ {
        p0 -= p0
        if false {
            p1 = ("a" + (p1 + "hello"))
            if false {
                v5 := []int{(fi4 + fi4), (fi4 - 32768)}
                if false {
                    fmt.Println(p0, p1)
                    if true {
                        v6 := GenS1{F0: 1, F1: 65535, F2: 0}
                        p1 = "a"
                        fmt.Println(fi4, fi4, p1)
                        _ = v6
                    }
                    var v7 float32 = 1.5
                    _ = v7
                }
                var v8 float64 = float64((0.0 * float32((float64(3.14159) + -2.25))))
                _ = v5
                _ = v8
            } else {
                for fi9 := 0; fi9 < 12; fi9++ {
                    var v10 float64 = 3.14159
                    _ = fi9
                    _ = v10
                }
            }
            v11 := GenS1{F0: -1, F1: 255, F2: 127}
            _ = v11
        }
        p0 -= p0
        _ = fi4
    }
    p1 = "foo bar"
    var v12 int = 1
    fmt.Println(p1, v12, p1)
    _ = p0
    _ = p1
    _ = v12
    return 255
}
func genFn1(p0 uint) uint8 {
    p0 -= (p0 * p0)
    p0 = 528
    var v13 uint16 = 1
    _ = p0
    _ = v13
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{("\n\t" + ("hello" + "a"))}
    v3 := make(map[int]int)
    if false {
        if ((((("" + "\n\t") + "foo bar") + ((("a" + "foo bar") + ("" + "\n\t")) + "foo bar")) + ("foo bar" + ("a" + ""))) <= ("" + "hello")) {
            v14 := GenS2{F0: 1, F1: true}
            _ = v14
        }
        var v15 float32 = -2.25
        _ = v15
    } else {
        var v16 int32 = -848
        v16 -= (v16 - v16)
        _ = v16
    }
    v17 := GenS1{F0: 1000, F1: 65535, F2: 1}
    for fi18 := 0; fi18 < 7; fi18++ {
        for fi19 := 0; fi19 < 1; fi19++ {
            var v20 string = "a"
            v20 = v20
            _ = fi19
            _ = v20
        }
        var v21 int32 = 1
        _ = fi18
        _ = v21
    }
    var v22 uint32 = 2
    if false {
        for fi23 := 0; fi23 < 7; fi23++ {
            var v24 uint32 = v22
            var v25 float32 = 0.0
            var v26 int64 = -333
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
    } else {
        switch "\n\t" {
        case "\n\t":
            _ = len(v1)
        }
        v2 = append(v2, "a")
    }
    v22 = 65535
    var v27 int16 = -572
    if true {
        var v28 uint64 = 1
        _ = v28
    }
    v29 := GenS3{F0: genFn0(4294967295, "hello")}
    v30 := v1[(v22 | 665)]
    v31 := GenS3{F0: 42}
    switch "\n\t" {
    case "a":
        delete(v3, 284)
    case "hello":
        v27 += (v27 - v27)
    case "":
        v32 := GenS3{F0: 349}
        _ = v32
    default:
        v27 = v27
    }
    v27 = v27
    for fi33 := 0; fi33 < 5; fi33++ {
        v27 -= v27
        for fi34 := 0; fi34 < 5; fi34++ {
            v27 -= v27
            _ = fi34
        }
        _ = fi33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
