package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 uint }
type GenS4 struct { F0 uint; F1 uint32 }
func genFn0(p0 uint16, p1 uint32) uint8 {
    v7 := GenS1{F0: -1}
    v8 := []int{32768, 9223372036854775807, 1}
    var v9 float64 = (1.5 * 0.0)
    switch "a" {
    case "a":
        switch "hello" {
        case "\n\t":
            p1 -= p1
        case "hello":
            var v10 string = ((("" + "foo bar") + ("a" + "")) + "hello")
            _ = v10
        default:
            switch "a" {
            case "\n\t":
                var v11 int = -9223372036854775808
                _ = v11
            case "a":
                v12 := GenS1{F0: 0}
                _ = v12
            case "foo bar":
                p0 = p0
            default:
                fmt.Println(v9, p1)
            }
        }
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]string)
    var v3 uint64 = 987
    var v4 uint = 42
    v5 := make(map[string]int)
    var v6 float32 = (float32(-2.25) + 1.5)
    if false {
        for fi13 := 0; fi13 < 14; fi13++ {
            var v14 uint16 = 32767
            _ = fi13
            _ = v14
        }
        var v15 int = 42
        _ = v15
    } else {
        var v16 float64 = 3.14159
        _ = v16
    }
    v17 := GenS1{F0: 1}
    if (4294967295 != 2147483647) {
        v18 := GenS2{F0: ("" + ("a" + ""))}
        _ = v18
    } else {
        var v19 int = 32768
        var v20 int16 = -32768
        v3 = 2
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        v4 += 403
        v22 := []int{(fi21 + (fi21 % ((-1) | 1))), -221}
        v2[0] = ("a" + "foo bar")
        _ = fi21
        _ = v22
    }
    v23 := GenS2{F0: ("hello" + "")}
    switch ("a" + "\n\t") {
    case "hello":
        for fi24 := 0; fi24 < 16; fi24++ {
            var v25 int = fi24
            v26 := v2[1]
            _ = fi24
            _ = v25
            _ = v26
        }
    }
    var v27 uint8 = uint8(v3)
    v28 := GenS2{F0: "a"}
    if (-392 > 42) {
        if false {
            if (v3 <= (v3 >> (((v4 / (((v4 ^ (v4 + v4))) | 1))) & 63))) {
                var v29 string = "hello"
                _ = v29
            }
        } else {
            for fi30 := 0; fi30 < 18; fi30++ {
                v6 += 1.5
                _ = fi30
            }
            switch 0 {
            case 127:
                if (-1 <= 9223372036854775807) {
                    v31 := []int{-118, 4294967295}
                    _ = v31
                } else {
                    v32 := GenS3{F0: (!(0 > v4)), F1: (v4 ^ v4)}
                    v4 -= v4
                    _ = v32
                }
            case 128:
                if ((9223372036854775807 == -1) || true) {
                    v33 := GenS4{F0: v4, F1: uint32(v4)}
                    if (255 <= (v27 - v27)) {
                        var v34 uint16 = 1
                        v27 = 2
                        v3 = v3
                        _ = v34
                    }
                    _ = v33
                }
            default:
                var v35 int32 = -1
                _ = v35
            }
        }
        var v36 uint8 = v27
        _ = v36
    }
    var v37 float32 = v6
    var v38 float32 = 1e10
    fmt.Println(v27, v37, v27)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
