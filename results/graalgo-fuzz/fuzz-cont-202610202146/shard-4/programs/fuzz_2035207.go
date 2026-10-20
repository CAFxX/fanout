package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint32 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint16; F1 int32 }
func genFn0(p0 string, p1 int32, p2 int) int8 {
    fmt.Println(p1)
    v6 := GenS1{F0: p2, F1: p2, F2: 2147483648}
    v7 := GenS2{F0: 2, F1: ("foo bar" + p0)}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -1
}
func genFn1(p0 bool, p1 int64, p2 string) uint64 {
    if p0 {
        var v8 bool = (!true)
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 int32 = -2147483648
            fmt.Println(fi9)
            _ = fi9
            _ = v10
        }
        _ = v8
    } else {
        fmt.Println(p1, p0)
        fmt.Println(p1, p0, p2)
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        v12 := GenS1{F0: (fi11 / (((fi11 ^ fi11)) | 1)), F1: fi11, F2: 239}
        v13 := []int{fi11, (fi11 / ((fi11) | 1))}
        _ = fi11
        _ = v12
        _ = v13
    }
    p1 -= -1
    p1 += (p1 & 38)
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[string]string)
    var v3 uint32 = 761
    v4 := []int{983, -2147483648, -473}
    var v5 uint8 = 127
    v5 -= (v5 / (((v5 ^ (v5 | 37))) | 1))
    v1 = (v1 ^ uint64(v3))
    var v14 int16 = 713
    v5 += v5
    switch "hello" {
    case "":
        var v15 int16 = v14
        _ = v15
    case "a":
        for fi16 := 0; fi16 < 10; fi16++ {
            switch ("a" + "hello") {
            case "hello":
                var v17 int16 = v14
                _ = v17
            }
            _ = fi16
        }
    default:
        v14 = v14
    }
    switch 869 {
    case -2147483648:
        var v18 int32 = 2147483647
        _ = v18
    }
    var v19 bool = (true || ((0.1 + -2.25) >= float64((0.0 + -2.25))))
    var v20 float64 = 0.1
    var v21 uint32 = (v3 + 2147483647)
    switch 544 {
    case 623:
        if (uint16(v5) < uint16(v5)) {
            for fi22 := 0; fi22 < 20; fi22++ {
                var v23 string = "foo bar"
                fmt.Println(v1, fi22)
                var v24 float64 = v20
                _ = fi22
                _ = v23
                _ = v24
            }
            v21 += 2147483647
            switch 4294967295 {
            case 2147483647:
                if v19 {
                    v25 := GenS1{F0: -756, F1: 0, F2: (v3 | (v21 + (v3 << ((42) & 63))))}
                    v1 += 1
                    _ = v25
                } else {
                    var v26 float64 = 1e10
                    _ = v26
                }
            }
        } else {
            var v27 uint8 = v5
            var v28 uint32 = v3
            _ = v27
            _ = v28
        }
    }
    v4 = append(v4, 32768)
    if true {
        var v29 int8 = -1
        v30 := GenS1{F0: 1, F1: 1, F2: 188}
        _ = v29
        _ = v30
    } else {
        v31 := GenS3{F0: 1, F1: -145}
        v32 := GenS2{F0: 65535, F1: "hello"}
        _ = v31
        _ = v32
    }
    var v33 int8 = 127
    fmt.Println(v21, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
