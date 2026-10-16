package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint }
func genFn0(p0 uint16, p1 uint64, p2 int32) string {
    p0 = (^(p0))
    if (-119 <= 32767) {
        if false {
            for fi5 := 0; fi5 < 10; fi5++ {
                if (!true) {
                    for fi6 := 0; fi6 < 14; fi6++ {
                        fmt.Println(p2, fi6)
                        _ = fi6
                    }
                }
                _ = fi5
            }
            switch int(p0) {
            case 9223372036854775807:
                var v7 int32 = (-(p2))
                _ = v7
            case -9223372036854775808:
                for fi8 := 0; fi8 < 1; fi8++ {
                    fmt.Println(fi8, p2)
                    p2 = (^(p2))
                    _ = fi8
                }
            }
        } else {
            p2 -= (^(p2))
        }
    } else {
        p2 += (-(p2))
        var v9 uint16 = 0
        v10 := []int{}
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return "\n\t"
}
func genFn1() uint32 {
    switch "a" {
    case "foo bar":
        var v11 int = 350
        _ = v11
    case "\n\t":
        v12 := GenS1{F0: -383, F1: 2147483647}
        _ = v12
    }
    v13 := GenS2{F0: "hello"}
    _ = v13
    return 42
}

func main() {
    var v1 float64 = (0.0 * 1e10)
    var v2 float32 = float32(1.5)
    var v3 uint32 = 65535
    v4 := make(map[uint32]int)
    var v14 uint64 = 1
    var v15 int8 = 104
    v16 := GenS1{F0: 0, F1: -2147483648}
    if (!((32768 <= -1) && false)) {
        v17 := []int{-1, 527, int(v14)}
        if (!false) {
            if true {
                var v18 int = 4294967295
                var v19 int = 4294967295
                v20 := GenS3{F0: 667}
                _ = v18
                _ = v19
                _ = v20
            }
            v15 += (v15 / (((v15 | 0)) | 1))
            switch ("a" + "hello") {
            case "":
                switch 32767 {
                case 128:
                    for fi21 := 0; fi21 < 15; fi21++ {
                        fmt.Println(v15)
                        var v22 int32 = 338
                        if len(v17) > 0 {
                            v17[0] = fi21
                        }
                        _ = fi21
                        _ = v22
                    }
                case -111:
                    v23 := GenS2{F0: "foo bar"}
                    _ = v23
                case 2147483647:
                    v2 = v2
                }
            }
        }
        v3 -= (^(v3))
        _ = v17
    } else {
        fmt.Println(v3, v15)
        var v24 float32 = v2
        _ = len(v4)
        _ = v24
    }
    v4[v3] = 128
    v15 -= (^(v15))
    v15 -= v15
    v3 -= (v3 / (((v3 + 65535)) | 1))
    fmt.Println(v14, v1, v1)
    if (true || false) {
        v25 := []int{}
        var v26 float32 = 1.5
        var v27 float64 = 1.5
        _ = v25
        _ = v26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
