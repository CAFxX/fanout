package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int; F1 uint; F2 bool }
type GenS3 struct { F0 bool; F1 uint32 }
func genFn0(p0 string, p1 int32, p2 int16) uint {
    for fi5 := 0; fi5 < 16; fi5++ {
        p0 = p0
        v6 := GenS1{F0: 734, F1: (("" + ("hello" + "hello")) + ((p0 + ((((p0 + "a") + "foo bar") + p0) + p0)) + "a"))}
        v7 := []int{(-(fi5)), -1}
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 uint16 = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 1
}

func main() {
    var v1 int8 = 42
    var v2 int32 = 2147483647
    var v3 bool = (!false)
    var v4 float32 = (0.0 * 1.5)
    v2 = v2
    for fi9 := 0; fi9 < 16; fi9++ {
        for fi10 := 0; fi10 < 2; fi10++ {
            v11 := []int{(-(fi10))}
            _ = fi10
            _ = v11
        }
        v12 := GenS2{F0: fi9, F1: 65535, F2: v3}
        _ = fi9
        _ = v12
    }
    if (127 < 9223372036854775807) {
        var v13 int = 2147483647
        _ = v13
    }
    v14 := []int{32767, 127, 128, -2147483648}
    if false {
        var v15 uint = 42
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        _ = v15
    } else {
        switch 42 {
        case -1:
            for fi16 := 0; fi16 < 15; fi16++ {
                if (!(!(-528 < -32768))) {
                    v17 := GenS3{F0: v3, F1: 4294967295}
                    var v18 float32 = v4
                    switch "hello" {
                    case "":
                        var v19 int32 = (-(v2))
                        _ = v19
                    case "a":
                        var v20 uint64 = 39
                        _ = v20
                    }
                    _ = v17
                    _ = v18
                }
                _ = len(v14)
                v1 -= 100
                _ = fi16
            }
        default:
            if true {
                v1 -= v1
                _ = len(v14)
                var v21 int16 = 185
                _ = v21
            } else {
                v2 += (^(v2))
            }
        }
        v2 += v2
    }
    var v22 uint = 1
    switch "a" {
    case "\n\t":
        v1 -= 0
    case "hello":
        var v23 uint8 = 255
        _ = v23
    case "a":
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
    }
    v4 = (v4 + -2.25)
    var v24 int = 387
    if true {
        var v25 int32 = (v2 >> (((v22 - 18446744073709551615)) & 63))
        for fi26 := 0; fi26 < 5; fi26++ {
            var v27 uint32 = 2
            var v28 uint8 = 79
            _ = fi26
            _ = v27
            _ = v28
        }
        for fi29 := 0; fi29 < 16; fi29++ {
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            v1 = 0
            var v30 int64 = 61
            _ = fi29
            _ = v30
        }
        _ = v25
    }
    var v31 float64 = (0.0 * float64(float32(float64(v4))))
    v24 += v24
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v24
    _ = v3
    _ = v31
    _ = v4
}
