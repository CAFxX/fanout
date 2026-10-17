package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint; F1 bool }
func genFn0(p0 int8, p1 int, p2 uint64) int8 {
    fmt.Println(p0)
    var v4 float32 = 0.1
    switch (p1 << ((1) & 63)) {
    case 0:
        p1 -= (p1 - (p1 & p1))
    default:
        for fi5 := 0; fi5 < 4; fi5++ {
            for fi6 := 0; fi6 < 13; fi6++ {
                v7 := GenS1{F0: 2, F1: 0}
                p2 -= 943
                p0 -= (-(p0))
                v8 := []int{(-(fi6)), (p1 ^ fi6), 42, (p1 >> ((1) & 63))}
                _ = fi6
                _ = v7
                _ = v8
            }
            _ = fi5
        }
    }
    v9 := []int{p1, (p1 & (^(p1))), (p1 | p1)}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v9
    return (p0 % (((p0 + p0)) | 1))
}

func main() {
    var v1 int16 = 1
    v2 := []string{("foo bar" + ("foo bar" + "foo bar")), "hello"}
    var v3 int64 = 712
    if (("\n\t" + "\n\t") >= ("foo bar" + ("hello" + "a"))) {
        switch "hello" {
        case "\n\t":
            var v10 uint64 = 18446744073709551615
            _ = v10
        }
        var v11 int16 = v1
        _ = v11
    }
    v3 = v3
    var v12 int64 = v3
    fmt.Println(v12)
    v1 = -32768
    fmt.Println(v12, v12, v1)
    for fi13 := 0; fi13 < 3; fi13++ {
        if false {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        v14 := GenS2{F0: 761, F1: 127}
        _ = fi13
        _ = v14
    }
    var v15 uint16 = 1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        v16 := GenS1{F0: 37, F1: v15}
        for fi17 := 0; fi17 < 11; fi17++ {
            var v18 int16 = 213
            var v19 uint8 = 2
            _ = len(v2)
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = v16
    } else {
        for fi20 := 0; fi20 < 18; fi20++ {
            switch -1 {
            case 9223372036854775807:
                if ((("hello" + "foo bar") + "foo bar") == "\n\t") {
                    for fi21 := 0; fi21 < 12; fi21++ {
                        var v22 uint = 4294967295
                        _ = fi21
                        _ = v22
                    }
                }
            case 881:
                for fi23 := 0; fi23 < 11; fi23++ {
                    v24 := GenS3{F0: (fi23 & fi23)}
                    if (51 != 2) {
                        v25 := GenS4{F0: 65535, F1: false}
                        v15 += v15
                        fmt.Println(v15, v15)
                        _ = v25
                    } else {
                        v2 = append(v2, (("hello" + "foo bar") + ("" + "a")))
                    }
                    if false {
                        var v26 uint32 = 175
                        _ = v26
                    }
                    _ = fi23
                    _ = v24
                }
            }
            switch "foo bar" {
            case "hello":
                v1 -= (v1 + v1)
            case "":
                v12 = v3
            default:
                for fi27 := 0; fi27 < 4; fi27++ {
                    fmt.Println(v15)
                    var v28 int = (fi20 << ((733) & 63))
                    _ = fi27
                    _ = v28
                }
            }
            v15 = 1
            _ = fi20
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
}
