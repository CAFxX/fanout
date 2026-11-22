package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int16) uint32 {
    for fi4 := 0; fi4 < 14; fi4++ {
        v5 := []int{(fi4 - (fi4 & (fi4 | fi4))), (fi4 * -1), (-(fi4))}
        p0 -= (p0 & -32768)
        _ = fi4
        _ = v5
    }
    fmt.Println(p0)
    v6 := GenS1{F0: (-521 > -1), F1: 1}
    p0 = (p0 + p0)
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 int64 = -893
        fmt.Println(fi7, v8, fi7)
        if false {
            if true {
                for fi9 := 0; fi9 < 7; fi9++ {
                    var v10 int32 = -1
                    _ = fi9
                    _ = v10
                }
            }
            var v11 uint = 4294967295
            _ = v11
        } else {
            fmt.Println(p0, fi7)
            v12 := GenS2{F0: (v8 % ((v8) | 1))}
            if true {
                fmt.Println(p0, p0, p0)
                var v13 int = -2147483648
                var v14 int = (-(v13))
                _ = v13
                _ = v14
            }
            _ = v12
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = v6
    return 42
}
func genFn1(p0 int32, p1 uint32, p2 int) uint8 {
    v15 := GenS3{F0: 42}
    var v16 bool = true
    switch "\n\t" {
    case "a":
        var v17 int16 = 364
        _ = v17
    case "hello":
        v18 := GenS2{F0: 9223372036854775807}
        _ = v18
    default:
        for fi19 := 0; fi19 < 12; fi19++ {
            fmt.Println(v16)
            var v20 uint8 = 42
            _ = fi19
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        v22 := []int{p2, (fi21 - (p2 + (p2 & p2))), (p2 & -209)}
        if false {
            _ = len(v22)
            var v23 uint = 65535
            fmt.Println(p2, p0)
            _ = v23
        } else {
            if false {
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
            }
            v16 = true
        }
        p0 = -984
        _ = fi21
        _ = v22
    }
    var v24 int = (p2 << ((225) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    _ = v24
    return 0
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]int)
    v3 := []uint32{0, 1, 180}
    var v25 int8 = 127
    if (4294967295 > 4294967295) {
        for fi26 := 0; fi26 < 9; fi26++ {
            var v27 float64 = 1e10
            v3 = append(v3, 2)
            switch "\n\t" {
            case "foo bar":
                var v28 string = ""
                _ = v28
            case "a":
                v29 := v1["hello"]
                _ = v29
            case "hello":
                delete(v1, ("hello" + "\n\t"))
            default:
                var v30 int8 = v25
                _ = v30
            }
            _ = fi26
            _ = v27
        }
    } else {
        fmt.Println(v25)
    }
    var v31 string = ("foo bar" + "hello")
    var v32 int16 = 32767
    var v33 int64 = 0
    if ((!false) || false) {
        for fi34 := 0; fi34 < 1; fi34++ {
            fmt.Println(v31)
            var v35 uint64 = 2
            _ = fi34
            _ = v35
        }
        var v36 bool = (1e10 < (-2.25 + (((0.0 + 1.5) / 1.5) * -2.25)))
        _ = len(v3)
        _ = v36
    }
    switch ("a" + (v31 + "\n\t")) {
    case "foo bar":
        var v37 int = 128
        _ = v37
    }
    v25 -= v25
    delete(v2, -238)
    v32 += v32
    v25 = 127
    v33 = 417
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v33
}
