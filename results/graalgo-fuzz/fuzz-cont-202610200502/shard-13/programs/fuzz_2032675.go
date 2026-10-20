package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int8 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 int16; F1 uint64 }
func genFn0(p0 int32, p1 int32, p2 int16) uint8 {
    v4 := []int{284}
    if (!("foo bar" != "")) {
        var v5 string = "a"
        v6 := GenS1{F0: 629, F1: (p0 | (p1 ^ (p1 & p0))), F2: 1}
        _ = v5
        _ = v6
    } else {
        p2 = p2
        var v7 float32 = float32((1.5 * 1.5))
        _ = v7
    }
    var v8 uint64 = 4294967295
    for fi9 := 0; fi9 < 2; fi9++ {
        fmt.Println(v8, p1)
        if (98 >= -12) {
            v10 := GenS2{F0: "foo bar", F1: 255}
            p2 = (p2 | (p2 - (p2 | (p2 << ((2) & 63)))))
            fmt.Println(v8, v8)
            _ = v10
        } else {
            fmt.Println(v8, fi9)
            var v11 float64 = (3.14159 / float64((0.0 + 1.5)))
            v12 := GenS1{F0: 65535, F1: (p1 >> ((4294967295) & 63)), F2: 65}
            _ = v11
            _ = v12
        }
        v13 := GenS1{F0: 65535, F1: p1, F2: 42}
        _ = fi9
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v8
    return 221
}

func main() {
    var v1 int8 = 100
    v2 := make(map[string]string)
    var v3 bool = (0 == 1)
    var v14 uint = 65535
    v3 = true
    switch "hello" {
    case "a":
        v1 = 0
    default:
        for fi15 := 0; fi15 < 5; fi15++ {
            fmt.Println(v14, v3)
            var v16 bool = (167 > 42)
            if (255 <= 37) {
                if v16 {
                    var v17 uint32 = 4294967295
                    _ = v17
                }
                fmt.Println(v16, v3, v1)
            } else {
                v16 = v3
                var v18 int8 = (v1 ^ v1)
                var v19 bool = (!v3)
                _ = v18
                _ = v19
            }
            _ = fi15
            _ = v16
        }
    }
    switch "\n\t" {
    case "hello":
        if (2 == 850) {
            var v20 bool = v3
            _ = len(v2)
            fmt.Println(v3, v3)
            _ = v20
        } else {
            v1 = 112
            v1 -= (v1 >> (((v14 | (v14 | 2))) & 63))
        }
    default:
        var v21 string = "\n\t"
        _ = v21
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        switch fi22 {
        case 2147483647:
            _ = len(v2)
        case 315:
            v23 := GenS2{F0: (("\n\t" + "foo bar") + ("foo bar" + "foo bar")), F1: 18446744073709551615}
            _ = v23
        }
        _ = len(v2)
        delete(v2, ((("foo bar" + ("hello" + "hello")) + ("\n\t" + "a")) + "foo bar"))
        _ = fi22
    }
    v24 := GenS3{F0: v14, F1: 18446744073709551615}
    if false {
        for fi25 := 0; fi25 < 3; fi25++ {
            v2[("" + "foo bar")] = ("hello" + "foo bar")
            _ = fi25
        }
        v3 = false
        var v26 uint16 = 510
        _ = v26
    } else {
        v27 := []int{}
        _ = v27
    }
    v28 := GenS4{F0: 42, F1: 4294967295}
    if (4294967295 > 18446744073709551615) {
        v3 = (!v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v24
    _ = v28
    _ = v3
}
