package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
func genFn0(p0 int) string {
    p0 = (p0 * p0)
    for fi5 := 0; fi5 < 19; fi5++ {
        p0 = (p0 / ((p0) | 1))
        var v6 uint = 18446744073709551615
        var v7 uint32 = 2
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 bool = true
    v9 := GenS1{F0: 722, F1: 255}
    v10 := GenS1{F0: 255, F1: 0}
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return "hello"
}
func genFn1(p0 uint64, p1 int32, p2 int64) bool {
    if false {
        v11 := []int{32767, -524, 128, 465}
        switch ("foo bar" + "\n\t") {
        case "":
            var v12 int32 = p1
            _ = v12
        case "\n\t":
            var v13 uint64 = 865
            _ = v13
        default:
            if false {
                p0 -= (p0 - p0)
                p1 += p1
            } else {
                var v14 uint8 = 70
                _ = v14
            }
        }
        var v15 float64 = 3.14159
        _ = v11
        _ = v15
    }
    var v16 uint32 = 1
    p0 = (p0 - (p0 % ((18446744073709551615) | 1)))
    var v17 string = (("a" + ("foo bar" + (("" + "\n\t") + ""))) + ("\n\t" + "foo bar"))
    fmt.Println(p0, v17)
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return true
}

func main() {
    var v1 uint16 = 2
    var v2 float32 = 0.1
    v3 := make(map[uint32]bool)
    var v4 int64 = -1
    for fi18 := 0; fi18 < 11; fi18++ {
        var v19 int32 = -1
        _ = fi18
        _ = v19
    }
    var v20 float32 = v2
    for fi21 := 0; fi21 < 12; fi21++ {
        fmt.Println(v2, v2)
        _ = fi21
    }
    var v22 uint = 65535
    v2 = v20
    v23 := GenS1{F0: (v22 >> ((v22) & 63)), F1: v1}
    fmt.Println(v4, v1)
    for fi24 := 0; fi24 < 6; fi24++ {
        for fi25 := 0; fi25 < 14; fi25++ {
            v26 := []int{(-(fi24)), fi24}
            _ = fi25
            _ = v26
        }
        delete(v3, 42)
        var v27 bool = true
        _ = fi24
        _ = v27
    }
    if false {
        var v28 int32 = -832
        switch ("hello" + "\n\t") {
        case "foo bar":
            fmt.Println(v28, v4)
        case "":
            if false {
                fmt.Println(v1, v22, v2)
                fmt.Println(v28)
                v29 := GenS1{F0: (v22 / (((v22 % (((v22 - 451)) | 1))) | 1)), F1: 597}
                _ = v29
            }
        case "a":
            v30 := GenS1{F0: (v22 + 847), F1: (v1 | (v1 | v1))}
            _ = v30
        default:
            var v31 uint = 154
            _ = v31
        }
        _ = v28
    } else {
        if false {
            var v32 bool = false
            var v33 float64 = 1.5
            v4 = (v4 >> ((v22) & 63))
            _ = v32
            _ = v33
        } else {
            var v34 uint16 = v1
            v35 := []int{37, 678}
            v36 := GenS1{F0: (v22 + 255), F1: (^(v34))}
            _ = v34
            _ = v35
            _ = v36
        }
        var v37 float64 = float64(1.5)
        _ = v37
    }
    var v38 uint16 = (v1 * (v1 + (^(v1))))
    var v39 bool = false
    fmt.Println(v22, v4, v22)
    v1 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v38
    _ = v39
    _ = v4
}
