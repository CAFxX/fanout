package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
func genFn0() int32 {
    for fi5 := 0; fi5 < 4; fi5++ {
        for fi6 := 0; fi6 < 14; fi6++ {
            if true {
                if (!true) {
                    var v7 uint64 = 1
                    _ = v7
                }
                v8 := GenS1{F0: 1, F1: false}
                _ = v8
            }
            if false {
                var v9 uint64 = 2
                fmt.Println(fi5, fi5, fi6)
                _ = v9
            } else {
                var v10 int8 = 100
                v11 := GenS1{F0: int16(v10), F1: ("\n\t" <= (("foo bar" + "hello") + "hello"))}
                v12 := []int{}
                _ = v10
                _ = v11
                _ = v12
            }
            fmt.Println(fi5)
            _ = fi6
        }
        var v13 uint64 = 0
        _ = fi5
        _ = v13
    }
    var v14 int16 = 584
    v14 -= v14
    var v15 float32 = 1.5
    if false {
        var v16 uint32 = 61
        _ = v16
    }
    _ = v14
    _ = v15
    return 26
}
func genFn1(p0 uint32, p1 int64, p2 int32) uint {
    v17 := GenS1{F0: 987, F1: true}
    if (((("a" + "\n\t") + "\n\t") + "") >= (((("\n\t" + "hello") + ("foo bar" + (("a" + "hello") + ""))) + ("\n\t" + "foo bar")) + ("a" + "\n\t"))) {
        p1 -= p1
        if (-537 >= 32767) {
            if true {
                p1 -= p1
                var v18 string = "\n\t"
                _ = v18
            }
        }
        var v19 int64 = (p1 % ((155) | 1))
        _ = v19
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        v21 := []int{fi20, fi20, (^(fi20)), -1}
        _ = fi20
        _ = v21
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return 576
}

func main() {
    var v1 string = ("foo bar" + "")
    var v2 int32 = -135
    var v3 bool = false
    var v4 int8 = 100
    if (1 > 197) {
        var v22 int16 = -1
        _ = v22
    }
    v3 = v3
    var v23 bool = (!v3)
    var v24 uint8 = 128
    if (0 >= 0) {
        v25 := []int{4294967295, 42, 127, 0}
        _ = len(v25)
        switch 392 {
        case 42:
            v3 = v23
        case 968:
            if len(v25) > 0 {
                v25[0] = -2147483648
            }
        default:
            switch 127 {
            case 341:
                _ = len(v25)
            case 127:
                fmt.Println(v1, v24, v23)
            default:
                var v26 uint8 = (v24 >> ((42) & 63))
                _ = v26
            }
        }
        _ = v25
    }
    v4 += -128
    var v27 int32 = (^(v2))
    switch v1 {
    case "":
        for fi28 := 0; fi28 < 4; fi28++ {
            if (false || (!v3)) {
                v24 = v24
            }
            _ = fi28
        }
    case "hello":
        v3 = v3
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v4
}
