package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 string; F1 uint }
func genFn0() bool {
    var v7 uint16 = 0
    v7 += (v7 / (((^(v7))) | 1))
    v7 += (v7 - v7)
    v8 := []int{0, 4294967295, 128}
    _ = v7
    _ = v8
    return false
}

func main() {
    var v1 uint16 = 65535
    v2 := make(map[uint32]bool)
    var v3 int = 45
    v4 := []uint32{2, 356, 0}
    v5 := make(map[uint32]uint32)
    v6 := []uint32{4294967295, 586, 2147483648}
    if true {
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 uint8 = 1
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 17; fi11++ {
            if (255 >= 68) {
                var v12 int32 = 1
                _ = v12
            } else {
                v13 := GenS1{F0: 0, F1: (v3 >> ((2) & 63)), F2: (fi11 % ((v3) | 1))}
                _ = v13
            }
            v1 -= (^(v1))
            v3 = v3
            _ = fi11
        }
        if len(v4) > 0 {
            v4[0] = 767
        }
    } else {
        var v14 uint8 = uint8(v1)
        var v15 float32 = -2.25
        _ = v14
        _ = v15
    }
    if len(v6) > 0 {
        v6[0] = 0
    }
    _ = len(v4)
    var v16 uint = 683
    for fi17 := 0; fi17 < 9; fi17++ {
        switch (v3 - v3) {
        case 1:
            if false {
                var v18 uint32 = 0
                var v19 float32 = ((1e10 - ((3.14159 + 0.1) + ((0.0 - 3.14159) * 1e10))) * 0.0)
                _ = v18
                _ = v19
            } else {
                _ = len(v4)
            }
        case 4294967295:
            v3 += (v3 + 1)
        default:
            var v20 int = (fi17 + (v3 - v3))
            _ = v20
        }
        if true {
            for fi21 := 0; fi21 < 13; fi21++ {
                fmt.Println(v1, v1, fi17)
                _ = fi21
            }
            v3 += (v3 + v3)
            if ((^(v1)) > 32767) {
                var v22 int64 = -801
                if true {
                    v3 = (v3 - v3)
                } else {
                    v3 = 0
                }
                v22 -= 0
                _ = v22
            } else {
                v23 := GenS2{F0: 255}
                var v24 bool = false
                _ = v23
                _ = v24
            }
        } else {
            if ((2 < 36) || true) {
                for fi25 := 0; fi25 < 5; fi25++ {
                    switch (("" + (("a" + "") + (("hello" + "foo bar") + "foo bar"))) + (("foo bar" + "\n\t") + ((("\n\t" + "a") + "") + ""))) {
                    case "hello":
                        var v26 uint32 = 2
                        _ = v26
                    case "":
                        var v27 int32 = -720
                        _ = v27
                    }
                    _ = fi25
                }
            } else {
                switch "hello" {
                case "foo bar":
                    v28 := GenS2{F0: 42}
                    _ = v28
                case "hello":
                    var v29 int16 = -1
                    _ = v29
                case "a":
                    var v30 uint16 = (v1 % (((^(v1))) | 1))
                    _ = v30
                default:
                    var v31 uint = (v16 - v16)
                    _ = v31
                }
                var v32 bool = ((v3 * v3) >= v3)
                fmt.Println(v1, v3, fi17)
                _ = v32
            }
            var v33 int = -924
            _ = v33
        }
        _ = fi17
    }
    var v34 uint16 = v1
    switch v3 {
    case 32767:
        for fi35 := 0; fi35 < 9; fi35++ {
            switch v3 {
            case 1:
                var v36 uint8 = 20
                _ = v36
            }
            v16 += (v16 ^ (v16 & (v16 / ((0) | 1))))
            _ = fi35
        }
    default:
        v1 = v1
    }
    if (v34 == v1) {
        switch v3 {
        case 32767:
            switch 9223372036854775807 {
            case -9223372036854775808:
                var v37 uint32 = 2
                _ = v37
            case 9223372036854775807:
                fmt.Println(v3, v1)
            default:
                v38 := GenS2{F0: 255}
                _ = v38
            }
        case 4294967295:
            v39 := GenS1{F0: 42, F1: 4294967295, F2: v3}
            _ = v39
        case -2147483648:
            fmt.Println(v16, v3, v1)
        }
        v1 = 1
    }
    v40 := GenS3{F0: -1, F1: true}
    var v41 int8 = 100
    var v42 int64 = -9223372036854775808
    var v43 int32 = 2147483647
    v44 := GenS4{F0: ("" + "hello"), F1: v16}
    v45 := GenS3{F0: (v42 << ((v16) & 63)), F1: ((-(v43)) > v43)}
    v46 := GenS1{F0: 2, F1: (v3 % (((^(v3))) | 1)), F2: (v3 / ((v3) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
