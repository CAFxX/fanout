package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
func genFn0() int32 {
    var v4 int32 = 42
    var v5 uint = 255
    switch "foo bar" {
    case "\n\t":
        var v6 int32 = v4
        _ = v6
    default:
        var v7 int64 = 1
        _ = v7
    }
    _ = v4
    _ = v5
    return (v4 % ((0) | 1))
}
func genFn1(p0 string, p1 uint64, p2 uint) uint8 {
    fmt.Println(p0, p0)
    var v8 int8 = 100
    for fi9 := 0; fi9 < 19; fi9++ {
        for fi10 := 0; fi10 < 14; fi10++ {
            var v11 uint64 = p1
            v12 := []int{fi10, fi9}
            v8 = (v8 >> (((p2 / (((^(p2))) | 1))) & 63))
            _ = fi10
            _ = v11
            _ = v12
        }
        for fi13 := 0; fi13 < 3; fi13++ {
            switch fi9 {
            case 127:
                var v14 uint64 = (p1 * (p1 * p1))
                _ = v14
            default:
                v15 := []int{fi13, (^(fi13))}
                _ = v15
            }
            p0 = p0
            var v16 int8 = v8
            _ = fi13
            _ = v16
        }
        var v17 float32 = 0.0
        _ = fi9
        _ = v17
    }
    if true {
        p1 = p1
    } else {
        fmt.Println(p2, p2)
        if (!false) {
            for fi18 := 0; fi18 < 2; fi18++ {
                v19 := []int{(fi18 / ((fi18) | 1)), (fi18 * 32767), (fi18 * -679), fi18}
                for fi20 := 0; fi20 < 16; fi20++ {
                    var v21 float64 = 3.14159
                    var v22 int8 = (v8 + (^(v8)))
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                for fi23 := 0; fi23 < 2; fi23++ {
                    p2 = p2
                    v24 := GenS1{F0: 42, F1: -1}
                    _ = fi23
                    _ = v24
                }
                _ = fi18
                _ = v19
            }
            if (false && false) {
                fmt.Println(p2, p0)
                v25 := []int{int(p1), -1}
                _ = v25
            }
        }
    }
    p2 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 128
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]string)
    var v3 uint16 = 904
    var v26 int64 = 42
    v26 += (v26 - (v26 % ((v26) | 1)))
    switch 128 {
    case 32768:
        for fi27 := 0; fi27 < 9; fi27++ {
            var v28 uint32 = 930
            delete(v1, (fi27 % ((-320) | 1)))
            _ = fi27
            _ = v28
        }
    case -2147483648:
        var v29 int64 = (v26 / (((v26 << ((2) & 63))) | 1))
        _ = v29
    case -705:
        for fi30 := 0; fi30 < 4; fi30++ {
            v2[65535] = ("a" + (((("hello" + "hello") + "hello") + ((("foo bar" + "\n\t") + "foo bar") + ("\n\t" + "foo bar"))) + "hello"))
            _ = fi30
        }
    }
    v31 := []int{}
    v32 := GenS1{F0: -364, F1: -1}
    fmt.Println(v3, v3)
    switch "\n\t" {
    case "foo bar":
        for fi33 := 0; fi33 < 5; fi33++ {
            v26 = v26
            v26 = int64(v3)
            _ = fi33
        }
    default:
        var v34 float32 = 0.0
        _ = v34
    }
    var v35 uint64 = 0
    var v36 uint8 = 128
    switch "hello" {
    case "a":
        fmt.Println(v35)
    }
    switch int(v26) {
    case -9223372036854775808:
        fmt.Println(v26, v35)
    }
    delete(v2, 2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v35
    _ = v36
}
