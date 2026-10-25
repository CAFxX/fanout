package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 bool }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 uint16; F1 int64 }
func genFn0(p0 uint8) uint16 {
    p0 = 19
    switch "\n\t" {
    case "\n\t":
        if (!false) {
            var v7 bool = ((false && (-9223372036854775808 < -1)) || true)
            v7 = v7
            for fi8 := 0; fi8 < 5; fi8++ {
                var v9 int = (fi8 % ((int(p0)) | 1))
                var v10 uint16 = 65535
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = v7
        } else {
            for fi11 := 0; fi11 < 7; fi11++ {
                p0 = (p0 & (p0 << ((1) & 63)))
                p0 += 2
                switch 32768 {
                case 9223372036854775807:
                    var v12 int = fi11
                    _ = v12
                default:
                    var v13 uint32 = 206
                    _ = v13
                }
                _ = fi11
            }
            p0 = p0
            p0 = (p0 << ((7) & 63))
        }
    }
    v14 := []int{}
    v15 := GenS1{F0: (p0 >> ((255) & 63)), F1: -450, F2: false}
    _ = p0
    _ = v14
    _ = v15
    return 61
}

func main() {
    var v1 uint32 = 2147483648
    var v2 uint8 = 64
    v3 := make(map[uint32]string)
    var v4 uint32 = v1
    v5 := make(map[string]uint32)
    var v6 uint8 = (v2 - (v2 + v2))
    if false {
        v4 = (^(v4))
        var v16 bool = (65535 < uint16(v2))
        _ = v16
    } else {
        if (((("foo bar" + "foo bar") + "") + "foo bar") < "foo bar") {
            for fi17 := 0; fi17 < 9; fi17++ {
                var v18 int64 = -9223372036854775808
                if true {
                    v2 = v2
                }
                var v19 string = "foo bar"
                _ = fi17
                _ = v18
                _ = v19
            }
        }
    }
    var v20 uint32 = (^(v1))
    for fi21 := 0; fi21 < 2; fi21++ {
        v22 := GenS1{F0: 127, F1: (fi21 & fi21), F2: false}
        _ = fi21
        _ = v22
    }
    v23 := GenS2{F0: (v20 / ((v4) | 1)), F1: 1}
    v4 = v4
    var v24 int8 = 46
    fmt.Println(v6, v6, v2)
    v25 := []int{312, -649}
    v26 := GenS1{F0: 1, F1: -9223372036854775808, F2: ("a" >= (("foo bar" + "") + ""))}
    for fi27 := 0; fi27 < 12; fi27++ {
        v28 := GenS3{F0: 0, F1: 42}
        for fi29 := 0; fi29 < 13; fi29++ {
            _ = len(v25)
            _ = fi29
        }
        for fi30 := 0; fi30 < 13; fi30++ {
            v31 := v3[v4]
            _ = fi30
            _ = v31
        }
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
