package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint64, p1 int16, p2 int32) int32 {
    var v4 uint = 0
    var v5 uint16 = 385
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p2
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]string)
    v3 := make(map[int]string)
    v2["foo bar"] = (("a" + ("foo bar" + "a")) + ((("foo bar" + ("a" + ("\n\t" + ("\n\t" + "")))) + "a") + "\n\t"))
    var v6 uint = 42
    var v7 float32 = 0.1
    for fi8 := 0; fi8 < 16; fi8++ {
        switch "\n\t" {
        case "hello":
            fmt.Println(v6, v7)
        case "foo bar":
            for fi9 := 0; fi9 < 13; fi9++ {
                v6 = (v6 - (v6 << (((v6 + v6)) & 63)))
                delete(v1, -259)
                v10 := []int{fi9, (fi9 * (fi8 << (((v6 % ((v6) | 1))) & 63)))}
                _ = fi9
                _ = v10
            }
        }
        for fi11 := 0; fi11 < 9; fi11++ {
            v6 = v6
            var v12 string = "foo bar"
            _ = fi11
            _ = v12
        }
        var v13 bool = true
        _ = fi8
        _ = v13
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        v15 := GenS1{F0: "foo bar", F1: 0, F2: 1}
        v7 += 1.5
        _ = fi14
        _ = v15
    }
    v6 = v6
    switch 32768 {
    case 1:
        for fi16 := 0; fi16 < 5; fi16++ {
            v7 = (v7 / (1.5 / 0.1))
            var v17 string = (((("hello" + "\n\t") + "\n\t") + ("hello" + "hello")) + ((("hello" + ("a" + "hello")) + (("" + "hello") + "\n\t")) + "a"))
            v18 := GenS2{F0: 700}
            _ = fi16
            _ = v17
            _ = v18
        }
    case -822:
        v7 += -2.25
    case 892:
        switch "" {
        case "":
            fmt.Println(v7, v7)
        default:
            v6 -= v6
        }
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        switch fi19 {
        case 327:
            var v20 float64 = (1e10 / 1e10)
            _ = v20
        case 285:
            delete(v2, (("\n\t" + "foo bar") + "a"))
        }
        v21 := GenS3{F0: 229}
        fmt.Println(v7)
        _ = fi19
        _ = v21
    }
    switch 0 {
    case 389:
        var v22 string = ("" + ("a" + "a"))
        _ = v22
    case -806:
        v6 = v6
    case 4294967295:
        v1[-844] = "\n\t"
    }
    switch 4294967295 {
    case 351:
        v23 := GenS2{F0: 1}
        _ = v23
    case 156:
        v7 -= 1.5
    }
    var v24 int8 = 7
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v6
    _ = v7
}
