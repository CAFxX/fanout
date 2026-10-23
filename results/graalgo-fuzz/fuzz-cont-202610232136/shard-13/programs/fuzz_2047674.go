package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int32; F1 uint; F2 uint }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16; F1 uint }
func genFn0(p0 uint8, p1 string) int64 {
    if false {
        v6 := GenS1{F0: 65535, F1: 75, F2: 2147483647}
        p1 = (p1 + p1)
        _ = v6
    } else {
        switch p1 {
        case "hello":
            p1 = ((((p1 + ((p1 + "a") + p1)) + ((p1 + p1) + ("hello" + p1))) + ("\n\t" + "a")) + p1)
        case "":
            var v7 bool = true
            _ = v7
        case "a":
            switch 42 {
            case 127:
                v8 := []int{2147483647, int(p0), 64}
                _ = v8
            case 32767:
                p0 -= p0
            case 42:
                v9 := GenS2{F0: 1, F1: 1, F2: 243}
                _ = v9
            default:
                fmt.Println(p0, p0)
            }
        }
        switch 4294967295 {
        case -1:
            p1 = ""
        default:
            for fi10 := 0; fi10 < 16; fi10++ {
                var v11 bool = false
                if (v11 && true) {
                    switch 855 {
                    case 9223372036854775807:
                        v12 := GenS3{F0: v11, F1: p0}
                        _ = v12
                    default:
                        v13 := []int{(fi10 >> ((2) & 63)), fi10}
                        _ = v13
                    }
                }
                p1 = ((p1 + "hello") + p1)
                _ = fi10
                _ = v11
            }
        }
        if true {
            p0 = (p0 - 0)
            p0 -= 255
            var v14 uint32 = 742
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        for fi16 := 0; fi16 < 4; fi16++ {
            p0 += (p0 + 93)
            var v17 int32 = -2147483648
            _ = fi16
            _ = v17
        }
        for fi18 := 0; fi18 < 14; fi18++ {
            p1 = p1
            _ = fi18
        }
        _ = fi15
    }
    _ = p0
    _ = p1
    return -9223372036854775808
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{}
    var v3 uint8 = 2
    v4 := make(map[int]bool)
    v5 := make(map[int]bool)
    switch -2147483648 {
    case -75:
        for fi19 := 0; fi19 < 15; fi19++ {
            v3 = (v3 - v3)
            _ = fi19
        }
    }
    v3 = (v3 / (((^(v3))) | 1))
    var v20 string = "a"
    switch 157 {
    case 42:
        v21 := GenS4{F0: -597}
        _ = v21
    case -1:
        v5[879] = false
    case 32768:
        switch 1 {
        case -1:
            v20 = v20
        default:
            _ = len(v2)
        }
    }
    v3 -= v3
    var v22 int = -981
    switch "a" {
    case "":
        switch v20 {
        case "":
            v2 = append(v2, "")
        }
    case "\n\t":
        var v23 string = v20
        _ = v23
    }
    v22 = 128
    v20 = v20
    v2 = append(v2, (((v20 + v20) + ((v20 + v20) + v20)) + v20))
    v24 := GenS5{F0: 0, F1: 680}
    v25 := GenS5{F0: -1, F1: 2}
    var v26 bool = true
    switch (v22 + v22) {
    case 9223372036854775807:
        v20 = ("" + (v20 + "foo bar"))
    case -67:
        for fi27 := 0; fi27 < 11; fi27++ {
            fmt.Println(v26, v22)
            _ = fi27
        }
    case 1:
        var v28 bool = v26
        _ = v28
    }
    v20 = "a"
    var v29 uint32 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v24
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
