package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int }

func main() {
    var v1 uint16 = 32767
    v2 := make(map[string]int)
    var v3 int8 = 42
    var v4 int8 = v3
    v5 := make(map[string]string)
    var v6 uint32 = 322
    for fi7 := 0; fi7 < 14; fi7++ {
        if ((false || false) || false) {
            v8 := []int{(^(fi7)), fi7, (fi7 << ((313) & 63))}
            for fi9 := 0; fi9 < 17; fi9++ {
                v1 = 579
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                _ = fi9
            }
            _ = v8
        }
        v10 := v5["\n\t"]
        _ = fi7
        _ = v10
    }
    if (-32768 >= -32768) {
        var v11 float64 = 1.5
        _ = v11
    }
    if (true && false) {
        v12 := GenS1{F0: (-302 != -1), F1: 4294967295, F2: 15}
        var v13 int16 = -1
        _ = v12
        _ = v13
    } else {
        var v14 uint16 = (v1 << ((1) & 63))
        switch 436 {
        case -491:
            if (v3 == v3) {
                v3 = v3
                v14 = v1
                _ = len(v2)
            }
        case 32768:
            v14 -= (^(v14))
        default:
            if true {
                v15 := GenS2{F0: "foo bar"}
                _ = v15
            }
        }
        _ = v14
    }
    v4 -= (v3 * 1)
    v3 = v4
    for fi16 := 0; fi16 < 5; fi16++ {
        v17 := GenS3{F0: (fi16 ^ (fi16 << ((0) & 63)))}
        v18 := []int{32768, 128, fi16}
        _ = fi16
        _ = v17
        _ = v18
    }
    if (false && true) {
        var v19 string = "a"
        delete(v5, (v19 + (("a" + ("a" + v19)) + (v19 + "\n\t"))))
        _ = v19
    } else {
        for fi20 := 0; fi20 < 20; fi20++ {
            var v21 int8 = -41
            _ = fi20
            _ = v21
        }
    }
    if ("hello" >= "hello") {
        var v22 uint = 1
        for fi23 := 0; fi23 < 3; fi23++ {
            v4 -= v4
            fmt.Println(v6, fi23, v6)
            _ = fi23
        }
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
