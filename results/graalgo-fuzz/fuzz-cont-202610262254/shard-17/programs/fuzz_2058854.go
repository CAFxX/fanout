package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint }
func genFn0() int32 {
    switch "" {
    case "":
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 uint16 = 2
            _ = fi7
            _ = v8
        }
    case "foo bar":
        if true {
            for fi9 := 0; fi9 < 20; fi9++ {
                var v10 float32 = 1.5
                var v11 int16 = 1
                for fi12 := 0; fi12 < 19; fi12++ {
                    if (((false && true) || (0 >= 364)) && (!false)) {
                        var v13 int64 = 42
                        _ = v13
                    }
                    _ = fi12
                }
                _ = fi9
                _ = v10
                _ = v11
            }
        } else {
            var v14 string = "a"
            _ = v14
        }
    case "a":
        var v15 uint8 = 135
        _ = v15
    default:
        var v16 uint64 = 48
        _ = v16
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        switch -1 {
        case 32767:
            var v18 float64 = 0.0
            _ = v18
        case -599:
            var v19 string = "\n\t"
            _ = v19
        case 1:
            for fi20 := 0; fi20 < 20; fi20++ {
                if false {
                    v21 := []int{fi17, fi17, fi20}
                    _ = v21
                }
                var v22 int = fi20
                _ = fi20
                _ = v22
            }
        }
        _ = fi17
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        var v24 int32 = int32(fi23)
        v25 := []int{(fi23 >> ((42) & 63))}
        for fi26 := 0; fi26 < 17; fi26++ {
            fmt.Println(fi26)
            _ = fi26
        }
        _ = fi23
        _ = v24
        _ = v25
    }
    var v27 uint64 = 42
    _ = v27
    return 1
}
func genFn1() int {
    v28 := GenS1{F0: 117, F1: "", F2: 42}
    v29 := GenS1{F0: 42, F1: "a", F2: 0}
    v30 := GenS1{F0: 1, F1: "foo bar", F2: 65535}
    _ = v28
    _ = v29
    _ = v30
    return -127
}

func main() {
    var v1 uint8 = 2
    var v2 float64 = 1e10
    v3 := []uint32{68}
    v4 := []int{}
    v5 := []string{}
    v6 := make(map[string]uint32)
    v2 += 1e10
    var v31 bool = false
    var v32 int64 = 1
    v32 = v32
    var v33 float32 = (1.5 + 1.5)
    switch 42 {
    case 127:
        v2 = ((((3.14159 - v2) - v2) + v2) / ((0.1 * (0.0 / float64(v33))) / 1.5))
    default:
        v34 := GenS1{F0: 881, F1: "foo bar", F2: 42}
        _ = v34
    }
    for fi35 := 0; fi35 < 5; fi35++ {
        v36 := GenS1{F0: 65535, F1: ("foo bar" + "\n\t"), F2: 8}
        _ = fi35
        _ = v36
    }
    v37 := GenS1{F0: 42, F1: "hello", F2: 255}
    var v38 uint = 741
    var v39 int8 = -128
    v4 = append(v4, -455)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
