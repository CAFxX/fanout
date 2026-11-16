package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int64, p1 uint16) int16 {
    p0 = (^(p0))
    v7 := GenS1{F0: 198}
    _ = p0
    _ = p1
    _ = v7
    return -15
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 62
    var v3 int32 = -1
    var v4 int64 = -24
    v5 := make(map[uint32]bool)
    v6 := []string{"a", ("hello" + "a")}
    var v8 int = 508
    var v9 uint = 448
    if (127 <= 2) {
        if len(v6) > 0 {
            v6[0] = "a"
        }
        for fi10 := 0; fi10 < 14; fi10++ {
            if len(v6) > 0 {
                v6[0] = ("\n\t" + ((("a" + (("" + "\n\t") + "hello")) + (("a" + "hello") + "hello")) + ("\n\t" + (("foo bar" + "") + "hello"))))
            }
            _ = fi10
        }
    } else {
        for fi11 := 0; fi11 < 18; fi11++ {
            var v12 bool = true
            if ((true || true) || (!(v3 != 1))) {
                for fi13 := 0; fi13 < 6; fi13++ {
                    v4 = 0
                    _ = fi13
                }
                v8 = -2147483648
            }
            _ = fi11
            _ = v12
        }
        v2 -= v9
    }
    fmt.Println(v8)
    if false {
        switch v8 {
        case -700:
            v14 := GenS1{F0: (v9 * v2)}
            _ = v14
        case 140:
            v9 = v2
        }
    } else {
        v15 := GenS2{F0: v8, F1: 1}
        var v16 int = (v8 << (((v9 >> ((255) & 63))) & 63))
        for fi17 := 0; fi17 < 4; fi17++ {
            switch fi17 {
            case 128:
                switch (fi17 - v8) {
                case 4294967295:
                    v9 = 4294967295
                case 128:
                    v2 += (v9 / ((v2) | 1))
                }
            case -1:
                var v18 float32 = ((0.0 * (-2.25 + -2.25)) / (float32(1.5) / ((float32(((float64(1e10) + 1.5) / ((0.0 / 0.1) - (1e10 / 1.0)))) - (0.0 + -2.25)) / (-2.25 / float32(0.1)))))
                _ = v18
            default:
                var v19 uint64 = 18446744073709551615
                _ = v19
            }
            fmt.Println(v8, v16, v3)
            var v20 uint8 = 13
            _ = fi17
            _ = v20
        }
        _ = v15
        _ = v16
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        v22 := GenS3{F0: 42}
        _ = fi21
        _ = v22
    }
    v3 = v3
    v2 -= (v2 << ((v9) & 63))
    fmt.Println(v4, v2)
    var v23 uint64 = 42
    v3 = (v3 | v3)
    if false {
        v2 += v2
        v24 := GenS1{F0: 4294967295}
        if ((0.0 - 0.0) < 1.5) {
            if true {
                var v25 string = ""
                v26 := GenS2{F0: (^(v8)), F1: v23}
                _ = v25
                _ = v26
            }
        } else {
            for fi27 := 0; fi27 < 19; fi27++ {
                if ((fi27 | 32767) <= fi27) {
                    v28 := GenS1{F0: (^(v2))}
                    _ = v28
                } else {
                    if true {
                        var v29 uint16 = 539
                        _ = v29
                    } else {
                        fmt.Println(v3)
                        v2 += (v2 * v9)
                        fmt.Println(v3)
                    }
                }
                v1["hello"] = v8
                var v30 uint64 = (^(v23))
                _ = fi27
                _ = v30
            }
            v31 := GenS1{F0: (v9 / ((18446744073709551615) | 1))}
            _ = v31
        }
        _ = v24
    }
    var v32 uint16 = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
