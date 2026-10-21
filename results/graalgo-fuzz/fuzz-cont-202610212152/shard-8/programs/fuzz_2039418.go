package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 int8; F2 uint8 }
func genFn0(p0 bool, p1 uint8) uint64 {
    v4 := []int{32767}
    v5 := GenS1{F0: 849}
    var v6 float32 = 0.1
    if len(v4) > 0 {
        v4[0] = 810
    }
    if (p1 != (p1 >> ((65535) & 63))) {
        p1 -= 255
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 4294967295
}
func genFn1(p0 uint8) int16 {
    switch 1 {
    case 32768:
        for fi7 := 0; fi7 < 16; fi7++ {
            v8 := GenS2{F0: (32767 != 680)}
            _ = fi7
            _ = v8
        }
    case 9223372036854775807:
        var v9 string = "hello"
        _ = v9
    }
    if (42 != 1000) {
        if (!(0.1 <= (0.0 * float32((3.14159 / 1.0))))) {
            var v10 uint16 = 255
            if (1e10 != 1.5) {
                for fi11 := 0; fi11 < 19; fi11++ {
                    v12 := []int{-2147483648}
                    for fi13 := 0; fi13 < 12; fi13++ {
                        fmt.Println(p0, fi13, fi11)
                        switch 4294967295 {
                        case -2147483648:
                            var v14 int8 = 127
                            _ = v14
                        }
                        _ = len(v12)
                        _ = fi13
                    }
                    _ = fi11
                    _ = v12
                }
                v15 := GenS1{F0: 288}
                fmt.Println(p0, v10)
                _ = v15
            } else {
                v16 := []int{9223372036854775807, 102}
                _ = v16
            }
            p0 += (p0 >> ((0) & 63))
            _ = v10
        }
        for fi17 := 0; fi17 < 16; fi17++ {
            v18 := []int{32768, (fi17 & (fi17 << ((255) & 63))), -2147483648}
            _ = fi17
            _ = v18
        }
    }
    var v19 int64 = -9223372036854775808
    _ = p0
    _ = v19
    return 924
}

func main() {
    var v1 string = "hello"
    var v2 uint16 = 255
    v3 := make(map[string]int)
    v1 = "hello"
    fmt.Println(v1, v1)
    var v20 bool = true
    for fi21 := 0; fi21 < 17; fi21++ {
        v22 := GenS1{F0: 32767}
        _ = fi21
        _ = v22
    }
    v1 = v1
    for fi23 := 0; fi23 < 7; fi23++ {
        var v24 int64 = 460
        v25 := GenS1{F0: 0}
        _ = fi23
        _ = v24
        _ = v25
    }
    var v26 uint32 = 65535
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        var v27 float64 = (1.5 + float64(float32(1.5)))
        _ = v27
    case "hello":
        fmt.Println(v20, v1, v2)
    case "\n\t":
        v28 := GenS3{F0: 64, F1: 127, F2: 128}
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v20
    _ = v26
    _ = v3
}
