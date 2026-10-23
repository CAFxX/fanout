package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint64, p1 int32) int64 {
    if false {
        switch "\n\t" {
        case "hello":
            var v7 bool = false
            _ = v7
        case "foo bar":
            p0 = (p0 >> ((255) & 63))
        default:
            v8 := []int{127, 9223372036854775807, -649}
            _ = v8
        }
        switch 0 {
        case 0:
            switch 699 {
            case 385:
                p1 = p1
            }
        default:
            v9 := []int{-1}
            _ = v9
        }
        fmt.Println(p1, p1)
    }
    var v10 uint32 = 2
    if true {
        p0 = p0
        v11 := []int{int(p1)}
        _ = v11
    } else {
        for fi12 := 0; fi12 < 5; fi12++ {
            fmt.Println(p0)
            v10 = v10
            v13 := []int{(^(fi12)), 924, (-(fi12))}
            _ = fi12
            _ = v13
        }
        var v14 int64 = -708
        _ = v14
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 float32 = (float32(1e10) / -2.25)
        p0 = 1
        var v17 int32 = p1
        _ = fi15
        _ = v16
        _ = v17
    }
    switch ("hello" + (("" + "") + "a")) {
    case "hello":
        if (-128 < 1) {
            var v18 float32 = float32(1e10)
            _ = v18
        }
    default:
        for fi19 := 0; fi19 < 2; fi19++ {
            p1 = p1
            v20 := GenS1{F0: v10, F1: 1000, F2: -2147483648}
            v21 := []int{fi19, fi19, fi19, fi19}
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    _ = p0
    _ = p1
    _ = v10
    return 6
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]string)
    var v3 uint = 2
    v4 := make(map[uint32]bool)
    v5 := make(map[string]string)
    v6 := []string{"\n\t", "hello", ("\n\t" + (((("" + "foo bar") + "foo bar") + "\n\t") + "foo bar"))}
    var v22 bool = true
    if len(v6) > 0 {
        v6[0] = ("" + ("" + (("" + (("a" + "foo bar") + ("hello" + "a"))) + "\n\t")))
    }
    var v23 int16 = 0
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v24 uint64 = 1
    fmt.Println(v3, v22)
    v25 := GenS1{F0: 2147483648, F1: 1, F2: 0}
    fmt.Println(v24)
    fmt.Println(v23)
    v26 := GenS2{F0: 835}
    var v27 float32 = 1.5
    for fi28 := 0; fi28 < 8; fi28++ {
        for fi29 := 0; fi29 < 12; fi29++ {
            fmt.Println(v27, v3)
            _ = fi29
        }
        v3 += (v3 % ((v3) | 1))
        _ = fi28
    }
    v30 := GenS3{F0: -2147483648}
    v31 := GenS3{F0: 677}
    delete(v2, 654)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
