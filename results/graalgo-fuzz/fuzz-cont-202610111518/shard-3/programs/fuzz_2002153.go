package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 uint) int32 {
    p0 -= p0
    var v4 int = 42
    _ = p0
    _ = v4
    return -890
}
func genFn1() int16 {
    v5 := []int{}
    switch -2147483648 {
    case 32768:
        _ = len(v5)
    case 9223372036854775807:
        var v6 uint8 = 135
        _ = v6
    case 919:
        var v7 int = -2147483648
        _ = v7
    }
    if true {
        for fi8 := 0; fi8 < 19; fi8++ {
            v9 := GenS1{F0: -699, F1: 771, F2: -1}
            _ = fi8
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        v11 := GenS2{F0: -380}
        var v12 uint16 = 65535
        fmt.Println(v12)
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 string = "\n\t"
    _ = v13
    _ = v5
    return -32768
}

func main() {
    v1 := []uint32{562, 4294967295, 965, 2}
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]uint32)
    var v14 float64 = 1e10
    if true {
        v14 = 1e10
        v14 = 1.5
        var v15 int16 = genFn1()
        _ = v15
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        for fi17 := 0; fi17 < 16; fi17++ {
            fmt.Println(v14)
            v14 += v14
            v14 = (v14 * float64(0.0))
            v18 := GenS3{F0: 812}
            v19 := v3[4294967295]
            v3[v19] = (v19 << ((4294967295) & 63))
            fmt.Println(fi16)
            _ = fi17
            _ = v18
            _ = v19
        }
        for fi20 := 0; fi20 < 9; fi20++ {
            var v21 uint64 = 151
            v21 = v21
            _ = fi20
            _ = v21
        }
        _ = fi16
    }
    v14 = float64(0.0)
    v14 = -2.25
    v22 := GenS3{F0: 869}
    switch 902 {
    case -1:
        if (-1 != -128) {
            v23 := GenS1{F0: -1, F1: 1000, F2: -126}
            _ = v23
        }
    case 127:
        var v24 bool = (!((v14 + 0.0) > (v14 / (v14 - 0.0))))
        _ = v24
    }
    switch ("\n\t" + "hello") {
    case "\n\t":
        var v25 uint32 = 0
        _ = v25
    case "":
        v14 = 0.0
    }
    if (-772 >= 32767) {
        v26 := GenS3{F0: 2147483648}
        _ = v26
    }
    v2[65535] = "\n\t"
    v2[611] = ("\n\t" + ("a" + "foo bar"))
    if false {
        fmt.Println(v14)
        var v27 float32 = ((((-2.25 - ((0.0 + float32(0.1)) * 1e10)) + (0.0 * 1.5)) * (float32(-2.25) / 1.5)) * 0.0)
        _ = v27
    } else {
        if false {
            v14 = 1e10
        } else {
            v28 := GenS4{F0: 1}
            v29 := GenS3{F0: 1}
            _ = v28
            _ = v29
        }
    }
    v14 = ((v14 / v14) + (float64(1.5) - ((v14 + v14) * (v14 / v14))))
    v30 := v3[4294967295]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v30
}
