package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 string; F1 uint }
func genFn0(p0 int16, p1 string) uint32 {
    v4 := []int{17}
    if true {
        var v5 float64 = 1.5
        fmt.Println(v5, v5)
        _ = v5
    } else {
        fmt.Println(p0, p1)
    }
    var v6 float32 = 3.14159
    v6 += (((3.14159 - float32(1e10)) + 1.5) * 0.0)
    _ = p0
    _ = p1
    _ = v4
    _ = v6
    return 633
}
func genFn1(p0 uint, p1 uint32) uint8 {
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            switch fi7 {
            case -297:
                v9 := []int{fi8, (fi8 >> (((p0 >> ((p0) & 63))) & 63)), int(p1), 870}
                _ = v9
            }
            var v10 int16 = -20
            _ = fi8
            _ = v10
        }
        if (!false) {
            if false {
                fmt.Println(fi7, p0)
            } else {
                for fi11 := 0; fi11 < 15; fi11++ {
                    var v12 bool = false
                    _ = fi11
                    _ = v12
                }
            }
            for fi13 := 0; fi13 < 16; fi13++ {
                v14 := GenS1{F0: fi7}
                _ = fi13
                _ = v14
            }
            p1 = p1
        }
        _ = fi7
    }
    v15 := []int{}
    if (true || (!(2 > 207))) {
        fmt.Println(p0)
        if len(v15) > 0 {
            v15[0] = 32767
        }
        p0 += 18446744073709551615
    }
    var v16 float64 = 0.0
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return 128
}

func main() {
    var v1 uint64 = 955
    var v2 int32 = -1
    var v3 int64 = 42
    var v17 uint = 255
    var v18 string = (("hello" + "") + ("a" + ""))
    v19 := []int{4294967295, -1, int(v1), -503}
    var v20 uint = (v17 % (((v17 & v17)) | 1))
    v3 = v3
    switch v18 {
    case "":
        if true {
            v17 += (v20 - 1)
        } else {
            v20 = 967
            v17 = v20
            if false {
                v18 = ("foo bar" + "hello")
            }
        }
    case "a":
        for fi21 := 0; fi21 < 8; fi21++ {
            var v22 int64 = (^(v3))
            switch fi21 {
            case 127:
                v20 = v20
            default:
                switch 32768 {
                case 685:
                    v23 := GenS2{F0: 242, F1: 42}
                    _ = v23
                }
            }
            _ = fi21
            _ = v22
        }
    case "\n\t":
        v24 := GenS3{F0: v18, F1: (v20 << (((v17 - 0)) & 63))}
        _ = v24
    default:
        var v25 uint16 = 32767
        _ = v25
    }
    fmt.Println(v2, v3)
    v18 = (v18 + "")
    switch 855 {
    case 128:
        var v26 int = -976
        _ = v26
    case 785:
        var v27 string = v18
        _ = v27
    }
    for fi28 := 0; fi28 < 18; fi28++ {
        _ = len(v19)
        for fi29 := 0; fi29 < 3; fi29++ {
            var v30 uint16 = 496
            if len(v19) > 0 {
                v19[0] = 128
            }
            v19 = append(v19, (-(fi28)))
            _ = fi29
            _ = v30
        }
        _ = fi28
    }
    v18 = ("" + "a")
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
