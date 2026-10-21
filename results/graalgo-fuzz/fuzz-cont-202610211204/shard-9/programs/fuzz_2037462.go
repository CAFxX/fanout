package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 int8 }
func genFn0(p0 int16, p1 bool) uint64 {
    p1 = p1
    switch 9223372036854775807 {
    case 32767:
        p1 = (!p1)
    }
    v7 := GenS1{F0: 65535, F1: 32767}
    fmt.Println(p1, p0)
    v8 := GenS1{F0: 18446744073709551615, F1: 831}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint64, p1 uint64) int8 {
    fmt.Println(p0, p0)
    v9 := []int{}
    var v10 float32 = (1.5 - (-2.25 + 0.0))
    v11 := GenS2{F0: 2}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return -124
}

func main() {
    var v1 uint32 = 12
    var v2 uint16 = 255
    var v3 uint32 = v1
    var v4 float32 = (0.1 * 3.14159)
    v5 := make(map[string]bool)
    var v6 int16 = 0
    v12 := []int{128, -9223372036854775808, int(v3)}
    v1 = (v1 + v3)
    v13 := GenS3{F0: 18446744073709551615}
    var v14 uint16 = v2
    v14 = v2
    var v15 uint64 = 42
    v1 -= (v3 & (v1 + v1))
    if true {
        v12 = append(v12, -1)
        v15 = (v15 % ((genFn0(0, (!false))) | 1))
    }
    v14 = 2
    v2 -= (v2 & 2)
    if len(v12) > 0 {
        v12[0] = -399
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        switch "hello" {
        case "hello":
            v6 -= v6
        }
        for fi17 := 0; fi17 < 9; fi17++ {
            if (true || true) {
                var v18 float32 = (v4 / 0.1)
                fmt.Println(v1, v15)
                var v19 int64 = -174
                fmt.Println(fi16, v3)
                _ = v18
                _ = v19
            } else {
                if ("\n\t" == "\n\t") {
                    v6 = (-(v6))
                } else {
                    var v20 uint32 = 2147483647
                    _ = v20
                }
            }
            switch fi17 {
            case 162:
                var v21 uint16 = (v2 + v14)
                _ = v21
            case 290:
                v22 := GenS4{F0: (v3 << ((65535) & 63)), F1: 0}
                _ = v22
            }
            _ = fi17
        }
        for fi23 := 0; fi23 < 8; fi23++ {
            if false {
                var v24 string = "\n\t"
                v3 = 2147483647
                var v25 int = fi16
                _ = v24
                _ = v25
            } else {
                var v26 float32 = v4
                v14 -= 1
                var v27 uint16 = (^(v14))
                switch 0 {
                case 127:
                    fmt.Println(v2)
                default:
                    fmt.Println(v2)
                }
                var v28 int = fi16
                _ = v26
                _ = v27
                _ = v28
            }
            v14 = (v2 & 204)
            _ = fi23
        }
        _ = fi16
    }
    switch 1 {
    case -2147483648:
        switch 9223372036854775807 {
        case 0:
            for fi29 := 0; fi29 < 1; fi29++ {
                switch (fi29 + (fi29 >> ((2) & 63))) {
                case -1:
                    v15 = (v15 / (((v15 & v15)) | 1))
                case 0:
                    v6 -= (v6 * 32767)
                }
                v3 = (v3 >> ((4294967295) & 63))
                if len(v12) > 0 {
                    v12[0] = fi29
                }
                _ = fi29
            }
        default:
            var v30 int32 = 1
            _ = v30
        }
    case 32768:
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
    }
    var v31 uint = 4294967295
    if true {
        var v32 int = -9223372036854775808
        _ = v32
    }
    v33 := GenS1{F0: 1, F1: v2}
    var v34 uint64 = (v15 ^ 1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
