package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 bool }
type GenS2 struct { F0 uint32; F1 bool; F2 uint32 }
type GenS3 struct { F0 uint8; F1 string }
func genFn0(p0 string, p1 int16) uint64 {
    fmt.Println(p0)
    var v4 float64 = (-2.25 + 1e10)
    switch (p0 + "") {
    case "hello":
        var v5 bool = (((false || (250 == 42)) || false) || false)
        _ = v5
    case "a":
        for fi6 := 0; fi6 < 4; fi6++ {
            v7 := []int{(fi6 * fi6), (-(fi6)), fi6, (fi6 + (-(fi6)))}
            _ = fi6
            _ = v7
        }
    case "":
        var v8 int8 = -1
        _ = v8
    default:
        var v9 float64 = float64(((1.5 + float32((v4 / v4))) / 1e10))
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v4
    return 42
}

func main() {
    var v1 uint16 = 32767
    v2 := []uint32{0, 2147483647, 1}
    v3 := make(map[uint32]string)
    v1 += 65535
    for fi10 := 0; fi10 < 10; fi10++ {
        if len(v2) > 0 {
            v2[0] = 65535
        }
        _ = fi10
    }
    var v11 uint64 = genFn0(("" + "\n\t"), -119)
    var v12 string = "a"
    v13 := GenS1{F0: 46, F1: 2, F2: false}
    switch v12 {
    case "a":
        for fi14 := 0; fi14 < 20; fi14++ {
            v1 = 255
            _ = fi14
        }
    case "":
        var v15 uint32 = 0
        _ = v15
    }
    v12 = ""
    fmt.Println(v12)
    var v16 int = 832
    v17 := GenS2{F0: 65535, F1: false, F2: 4294967295}
    if ((true || true) && true) {
        v1 -= v1
        v12 = (("a" + (((v12 + "hello") + v12) + v12)) + v12)
    } else {
        var v18 int8 = 1
        if ((0.0 / (1e10 - (0.0 / (((1.5 - 1e10) - 0.0) + (-2.25 / 1.0))))) <= 0.0) {
            v3[363] = (("a" + (((v12 + v12) + "") + ("foo bar" + v12))) + ("" + ("foo bar" + "foo bar")))
            switch "foo bar" {
            case "a":
                fmt.Println(v12, v18)
            case "\n\t":
                var v19 float32 = float32(-2.25)
                _ = v19
            default:
                if (87 == v18) {
                    _ = len(v3)
                }
            }
            v2 = append(v2, 42)
        }
        _ = v18
    }
    switch "\n\t" {
    case "":
        for fi20 := 0; fi20 < 13; fi20++ {
            var v21 string = "hello"
            if false {
                if true {
                    var v22 float32 = 0.0
                    fmt.Println(v12, v12)
                    _ = v22
                } else {
                    v1 = (v1 ^ (v1 << ((1) & 63)))
                    v21 = (v21 + "")
                    v16 = fi20
                }
                v23 := v3[2]
                _ = v23
            } else {
                for fi24 := 0; fi24 < 5; fi24++ {
                    fmt.Println(v21, fi20)
                    _ = len(v3)
                    switch (-(fi20)) {
                    case -9223372036854775808:
                        v16 = 32768
                    case 9223372036854775807:
                        v1 = v1
                    }
                    v11 -= v11
                    var v25 int8 = 0
                    _ = fi24
                    _ = v25
                }
                var v26 int8 = 127
                _ = v26
            }
            _ = fi20
            _ = v21
        }
    case "a":
        var v27 float32 = 0.1
        _ = v27
    case "hello":
        v1 -= 255
    }
    if ((v1 << ((42) & 63)) > v1) {
        v3[1] = v12
        var v28 uint16 = (^(v1))
        switch "hello" {
        case "":
            v11 = genFn0(((v12 + v12) + (("hello" + "a") + "foo bar")), 32767)
        case "a":
            v29 := GenS1{F0: 127, F1: 230, F2: (-1 <= 0)}
            _ = v29
        default:
            for fi30 := 0; fi30 < 3; fi30++ {
                fmt.Println(v16)
                if true {
                    v2 = append(v2, 60)
                } else {
                    _ = len(v2)
                    var v31 int8 = -128
                    v32 := GenS3{F0: uint8(v11), F1: v12}
                    _ = v31
                    _ = v32
                }
                v11 -= v11
                _ = fi30
            }
        }
        _ = v28
    } else {
        var v33 uint64 = (v11 - genFn0(("\n\t" + v12), 42))
        for fi34 := 0; fi34 < 13; fi34++ {
            var v35 uint8 = 116
            v36 := GenS3{F0: (v35 ^ uint8(v1)), F1: "a"}
            _ = fi34
            _ = v35
            _ = v36
        }
        v1 = (v1 + 0)
        _ = v33
    }
    switch v12 {
    case "\n\t":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "a":
        if (-9223372036854775808 >= -9223372036854775808) {
            _ = len(v2)
        }
    default:
        fmt.Println(v1)
    }
    var v37 uint8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v37
}
