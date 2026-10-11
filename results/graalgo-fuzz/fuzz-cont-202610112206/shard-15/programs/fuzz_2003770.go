package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 int8; F2 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int8; F1 int64 }
func genFn0(p0 uint16) int8 {
    for fi6 := 0; fi6 < 11; fi6++ {
        p0 += 255
        if false {
            v7 := GenS1{F0: 1}
            v8 := []int{fi6, (fi6 / ((fi6) | 1)), 32768}
            v9 := GenS2{F0: 1, F1: -1, F2: -9223372036854775808}
            _ = v7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    v10 := GenS1{F0: 255}
    _ = p0
    _ = v10
    return 112
}
func genFn1(p0 uint32) uint8 {
    if (198 >= 224) {
        var v11 float32 = 0.1
        fmt.Println(p0, p0)
        if (4294967295 != 2) {
            fmt.Println(p0)
            p0 += (p0 & (p0 & (p0 | 936)))
            p0 += p0
        } else {
            var v12 bool = ((((false || true) && (false || (!(359 < 42)))) || (("\n\t" + "hello") < ("foo bar" + "foo bar"))) || true)
            v13 := []int{2147483647, 127, 32767}
            _ = v12
            _ = v13
        }
        _ = v11
    } else {
        if (32767 < 42) {
            v14 := []int{2147483647}
            switch "foo bar" {
            case "":
                v15 := GenS1{F0: 0}
                _ = v15
            default:
                v16 := GenS3{F0: 32767}
                _ = v16
            }
            p0 = p0
            _ = v14
        } else {
            for fi17 := 0; fi17 < 19; fi17++ {
                for fi18 := 0; fi18 < 8; fi18++ {
                    switch fi18 {
                    case 1:
                        v19 := GenS2{F0: 2147483647, F1: 42, F2: -166}
                        _ = v19
                    case -9223372036854775808:
                        v20 := GenS3{F0: 263}
                        _ = v20
                    case 127:
                        p0 = p0
                    }
                    if (((fi18 * fi17) != (fi18 ^ (fi18 | fi18))) && (!(0 <= -949))) {
                        v21 := GenS1{F0: 2}
                        var v22 string = "foo bar"
                        _ = v21
                        _ = v22
                    } else {
                        v23 := GenS4{F0: fi18}
                        _ = v23
                    }
                    _ = fi18
                }
                _ = fi17
            }
            var v24 int64 = 461
            _ = v24
        }
        v25 := GenS2{F0: 42, F1: -1, F2: 1}
        _ = v25
    }
    v26 := []int{42, 0, -90}
    _ = p0
    _ = v26
    return uint8(p0)
}

func main() {
    var v1 int64 = 352
    var v2 uint8 = 1
    var v3 uint16 = 42
    v4 := make(map[int]string)
    var v5 int64 = v1
    switch int(v2) {
    case -663:
        var v27 int64 = 9223372036854775807
        _ = v27
    case -428:
        var v28 string = (("foo bar" + "foo bar") + "")
        _ = v28
    }
    switch 898 {
    case -935:
        if true {
            v29 := []int{-1, 42}
            v4[2147483647] = "\n\t"
            _ = v29
        }
    default:
        var v30 uint64 = 737
        _ = v30
    }
    _ = len(v4)
    v31 := GenS5{F0: genFn0(665), F1: v5}
    var v32 int8 = genFn0(481)
    fmt.Println(v32, v1)
    var v33 float32 = 1.5
    v1 = -758
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
