package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
func genFn0(p0 uint8, p1 int32, p2 int32) int {
    p1 = p2
    p0 = 0
    switch ("\n\t" + ("hello" + (("hello" + "") + ""))) {
    case "":
        p1 += p1
    }
    if (435 >= 1) {
        var v6 float32 = ((-2.25 * -2.25) * (1e10 + 1.5))
        var v7 uint16 = 42
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return 128
}
func genFn1(p0 int16, p1 int16) int16 {
    v8 := []int{174, -9223372036854775808}
    p0 -= (p0 + p0)
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 int16 = p0
        _ = fi9
        _ = v10
    }
    switch "\n\t" {
    case "":
        fmt.Println(p1)
    case "foo bar":
        var v11 int16 = (p1 >> ((18446744073709551615) & 63))
        _ = v11
    default:
        fmt.Println(p1)
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        p1 = (p0 * (p0 / ((p1) | 1)))
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = v8
    return p0
}

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 2
    var v3 float64 = -2.25
    var v4 uint64 = 2
    v5 := make(map[int]uint32)
    if (!(true || ((127 == 0) && true))) {
        for fi13 := 0; fi13 < 6; fi13++ {
            var v14 uint8 = 2
            var v15 uint32 = 65535
            _ = fi13
            _ = v14
            _ = v15
        }
        switch ("" + "a") {
        case "":
            v16 := []int{946}
            _ = v16
        case "\n\t":
            fmt.Println(v4, v3)
        case "hello":
            for fi17 := 0; fi17 < 8; fi17++ {
                if true {
                    fmt.Println(v4)
                    var v18 int8 = -69
                    v19 := GenS1{F0: -677, F1: -2147483648}
                    _ = v18
                    _ = v19
                }
                _ = fi17
            }
        }
        var v20 uint64 = v4
        _ = v20
    }
    var v21 uint8 = 255
    v22 := []int{1, 32768}
    v22 = append(v22, -776)
    var v23 uint64 = v4
    for fi24 := 0; fi24 < 18; fi24++ {
        delete(v5, genFn0((v21 ^ (v21 * (v21 % (((v21 ^ v21)) | 1)))), -514, -430))
        if ((v23 & v4) >= (^(v4))) {
            delete(v5, 1)
            var v25 int8 = int8(v2)
            v26 := v1[int(v25)]
            _ = v25
            _ = v26
        }
        _ = fi24
    }
    fmt.Println(v21, v23, v21)
    v23 = (v4 | (v4 << ((42) & 63)))
    var v27 int = 128
    var v28 uint = 542
    if (42 >= 1000) {
        for fi29 := 0; fi29 < 1; fi29++ {
            for fi30 := 0; fi30 < 7; fi30++ {
                v3 += float64((((1e10 * 0.0) - 0.0) / float32((v3 * v3))))
                if ("a" > "") {
                    fmt.Println(fi29, v2)
                }
                var v31 uint = (v28 / ((v28) | 1))
                _ = fi30
                _ = v31
            }
            v1[genFn0(127, 0, 0)] = -1
            v21 += v21
            _ = fi29
        }
        if false {
            fmt.Println(v23, v2)
            v27 = 128
            for fi32 := 0; fi32 < 14; fi32++ {
                fmt.Println(v23, v2, v28)
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
                _ = fi32
            }
        } else {
            v3 = float64(-2.25)
            for fi33 := 0; fi33 < 19; fi33++ {
                for fi34 := 0; fi34 < 1; fi34++ {
                    var v35 uint64 = (v23 - v23)
                    _ = fi34
                    _ = v35
                }
                _ = fi33
            }
        }
        var v36 float64 = v3
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
