package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 int16; F1 int; F2 int32 }
type GenS4 struct { F0 int64; F1 int8; F2 int64 }
func genFn0(p0 uint64, p1 uint16) int64 {
    v4 := []int{-2147483648, 1, -544}
    p0 -= 42
    v5 := GenS1{F0: p0, F1: 42, F2: 489}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 536
}
func genFn1(p0 uint64, p1 uint32, p2 bool) bool {
    p0 = 0
    switch "hello" {
    case "a":
        p1 += (p1 | (p1 | p1))
    case "\n\t":
        v6 := GenS2{F0: p1, F1: p2}
        _ = v6
    case "hello":
        var v7 bool = (genFn0((p0 + 477), 42) != -9223372036854775808)
        _ = v7
    default:
        var v8 uint64 = (^(p0))
        _ = v8
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        for fi10 := 0; fi10 < 7; fi10++ {
            switch fi9 {
            case 4294967295:
                var v11 bool = (p2 && (122 >= 148))
                _ = v11
            case 128:
                if p2 {
                    fmt.Println(p2, fi9)
                    var v12 uint8 = 2
                    v13 := GenS1{F0: p0, F1: 1, F2: 1}
                    _ = v12
                    _ = v13
                }
            case 32768:
                switch fi9 {
                case -1:
                    v14 := []int{0}
                    _ = v14
                case 565:
                    v15 := []int{}
                    _ = v15
                default:
                    p0 += p0
                }
            default:
                p1 = 0
            }
            v16 := GenS3{F0: -615, F1: 94, F2: 553}
            _ = fi10
            _ = v16
        }
        _ = fi9
    }
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    return (!(1 > 162))
}

func main() {
    var v1 string = ("hello" + "")
    v2 := make(map[string]uint32)
    var v3 string = v1
    if (-783 != 1000) {
        var v17 float32 = 0.0
        var v18 int8 = 42
        _ = v17
        _ = v18
    }
    switch "\n\t" {
    case "\n\t":
        v3 = ""
    case "a":
        v1 = "a"
    }
    v1 = "\n\t"
    delete(v2, "")
    v3 = ((v1 + "\n\t") + "\n\t")
    if (false || (!("hello" > v1))) {
        for fi19 := 0; fi19 < 19; fi19++ {
            var v20 float32 = 1.5
            v21 := GenS3{F0: 1, F1: fi19, F2: 1}
            _ = fi19
            _ = v20
            _ = v21
        }
        var v22 uint = 65535
        v23 := GenS4{F0: genFn0(2, 2), F1: -1, F2: 443}
        _ = v22
        _ = v23
    } else {
        v24 := []int{}
        _ = v24
    }
    v1 = (v3 + v3)
    fmt.Println(v1)
    for fi25 := 0; fi25 < 4; fi25++ {
        if (485 >= 854) {
            v1 = v1
            v1 = ("\n\t" + ("\n\t" + ""))
        } else {
            fmt.Println(v3, v1)
            var v26 uint = 2
            fmt.Println(v3, v1, v1)
            _ = v26
        }
        var v27 uint8 = 128
        _ = fi25
        _ = v27
    }
    var v28 int64 = 42
    fmt.Println(v1, v1, v28)
    fmt.Println(v3, v3, v28)
    for fi29 := 0; fi29 < 14; fi29++ {
        v30 := []int{(fi29 + (^(fi29)))}
        _ = fi29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
}
