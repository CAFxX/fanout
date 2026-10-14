package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int64; F1 string; F2 string }
func genFn0(p0 int64, p1 uint8, p2 int) int {
    if (-789 == 1) {
        v5 := []int{}
        p0 = 42
        switch p2 {
        case 127:
            var v6 uint8 = p1
            _ = v6
        case 32767:
            fmt.Println(p1, p2, p2)
        default:
            var v7 uint = 0
            _ = v7
        }
        _ = v5
    }
    var v8 int64 = int64(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return (p2 >> ((255) & 63))
}
func genFn1(p0 uint, p1 int16) string {
    var v9 int = 4294967295
    if false {
        switch (v9 ^ genFn0(-756, 42, (v9 | 520))) {
        case -1:
            fmt.Println(p1, p0, p1)
        default:
            v9 -= (-(v9))
        }
        var v10 int8 = -1
        v10 += v10
        _ = v10
    } else {
        var v11 uint16 = 65535
        var v12 int16 = (p1 + p1)
        p0 += (^(p0))
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        v14 := []int{v9, (v9 | (^(fi13))), v9, -1}
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        for fi15 := 0; fi15 < 2; fi15++ {
            v14 = append(v14, (fi13 - 628))
            _ = fi15
        }
        _ = fi13
        _ = v14
    }
    v16 := []int{(^(v9)), (v9 << ((p0) & 63)), genFn0(-111, 217, int(p1)), (v9 * 509)}
    for fi17 := 0; fi17 < 15; fi17++ {
        for fi18 := 0; fi18 < 17; fi18++ {
            switch (v9 % ((v9) | 1)) {
            case -9223372036854775808:
                p0 = p0
            case -2147483648:
                if false {
                    var v19 uint32 = 1
                    v9 = v9
                    v9 -= (fi18 / ((fi18) | 1))
                    _ = v19
                }
            }
            _ = len(v16)
            v20 := GenS1{F0: 0}
            _ = fi18
            _ = v20
        }
        p0 = p0
        _ = fi17
    }
    _ = p0
    _ = p1
    _ = v16
    _ = v9
    return ("foo bar" + "hello")
}

func main() {
    var v1 int16 = 1000
    v2 := []int{131, 32768, 9223372036854775807, -9223372036854775808}
    var v3 uint64 = 671
    v4 := make(map[int]uint32)
    fmt.Println(v1, v3)
    v21 := GenS2{F0: 587, F1: 18446744073709551615}
    for fi22 := 0; fi22 < 12; fi22++ {
        var v23 bool = false
        var v24 uint8 = 229
        _ = fi22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 14; fi25++ {
        switch "hello" {
        case "\n\t":
            var v26 bool = (true || true)
            _ = v26
        case "":
            fmt.Println(fi25, fi25)
        default:
            var v27 int = (fi25 * (fi25 + fi25))
            _ = v27
        }
        _ = fi25
    }
    var v28 int32 = 0
    var v29 int16 = -745
    for fi30 := 0; fi30 < 2; fi30++ {
        v28 = (^(v28))
        var v31 uint = 255
        _ = fi30
        _ = v31
    }
    var v32 bool = (v29 >= (-(v1)))
    v33 := GenS1{F0: 128}
    if v32 {
        v34 := GenS2{F0: uint16(v3), F1: 255}
        var v35 int32 = v28
        var v36 string = "hello"
        _ = v34
        _ = v35
        _ = v36
    } else {
        var v37 uint8 = 154
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
        _ = v37
    }
    switch 1 {
    case 1:
        _ = len(v2)
    case 686:
        for fi38 := 0; fi38 < 17; fi38++ {
            fmt.Println(v29, v29)
            delete(v4, (fi38 | fi38))
            _ = fi38
        }
    }
    for fi39 := 0; fi39 < 5; fi39++ {
        var v40 uint = 255
        if true {
            switch "\n\t" {
            case "a":
                fmt.Println(v29, v28, fi39)
            case "hello":
                var v41 int8 = -1
                _ = v41
            default:
                v29 = v1
            }
            v42 := GenS3{F0: -1, F1: "\n\t", F2: (("a" + "a") + "a")}
            var v43 uint32 = 901
            _ = v42
            _ = v43
        } else {
            v32 = v32
            if ((!(0 < int64(v28))) || v32) {
                var v44 float64 = float64(float32(float64((float32(1.5) - float32(float64((-2.25 - 0.0)))))))
                _ = v44
            }
        }
        _ = fi39
        _ = v40
    }
    v3 -= v3
    fmt.Println(v1, v28, v28)
    v45 := GenS1{F0: 1}
    v28 = v28
    fmt.Println(v1, v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v45
}
