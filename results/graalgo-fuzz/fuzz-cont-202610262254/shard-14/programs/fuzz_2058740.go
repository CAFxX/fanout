package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 int; F2 uint }
func genFn0(p0 int16, p1 uint64, p2 uint16) uint32 {
    var v7 uint32 = 451
    v7 = (v7 & 4294967295)
    fmt.Println(v7)
    v8 := []int{int(p0)}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (v7 / (((v7 | v7)) | 1))
}
func genFn1() int32 {
    var v9 bool = true
    v10 := GenS1{F0: 2}
    var v11 uint = 42
    _ = v10
    _ = v11
    _ = v9
    return -902
}

func main() {
    v1 := make(map[string]bool)
    var v2 int = 32767
    v3 := make(map[string]bool)
    var v4 uint = 419
    v5 := make(map[string]int)
    var v6 int32 = -384
    if false {
        var v12 float32 = (0.1 - 0.0)
        switch v2 {
        case 32767:
            for fi13 := 0; fi13 < 6; fi13++ {
                var v14 uint = 255
                v15 := []int{fi13, v2}
                _ = fi13
                _ = v14
                _ = v15
            }
        case 2147483647:
            v2 = (v2 << (((^(v4))) & 63))
        }
        _ = v12
    }
    v4 += v4
    v4 += v4
    var v16 bool = true
    var v17 uint32 = 711
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 bool = v16
        _ = fi18
        _ = v19
    }
    if v16 {
        switch "" {
        case "":
            if ((!true) || ((-2.25 - 1e10) >= 1.5)) {
                for fi20 := 0; fi20 < 2; fi20++ {
                    for fi21 := 0; fi21 < 19; fi21++ {
                        v22 := GenS1{F0: 42}
                        _ = fi21
                        _ = v22
                    }
                    _ = fi20
                }
                for fi23 := 0; fi23 < 18; fi23++ {
                    v4 += 4294967295
                    v24 := v3[("\n\t" + "foo bar")]
                    for fi25 := 0; fi25 < 17; fi25++ {
                        v26 := GenS1{F0: 4294967295}
                        _ = fi25
                        _ = v26
                    }
                    _ = fi23
                    _ = v24
                }
                var v27 int64 = int64(v4)
                _ = v27
            }
        default:
            v28 := GenS1{F0: 1}
            _ = v28
        }
        switch -4 {
        case -106:
            for fi29 := 0; fi29 < 7; fi29++ {
                fmt.Println(v2, v2, v2)
                _ = fi29
            }
        }
        for fi30 := 0; fi30 < 6; fi30++ {
            var v31 float32 = 0.0
            if (v31 < -2.25) {
                v32 := []int{(fi30 / (((fi30 - (v2 + v2))) | 1)), 2147483647, -2147483648}
                v33 := GenS1{F0: 18446744073709551615}
                var v34 uint64 = 2
                _ = v32
                _ = v33
                _ = v34
            }
            v4 = (v4 % (((v4 % ((42) | 1))) | 1))
            _ = fi30
            _ = v31
        }
    } else {
        for fi35 := 0; fi35 < 13; fi35++ {
            var v36 uint32 = 42
            if false {
                fmt.Println(v4, v36, v4)
                var v37 int8 = 42
                _ = v37
            } else {
                var v38 uint = (v4 << (((v4 ^ 65535)) & 63))
                _ = v38
            }
            _ = fi35
            _ = v36
        }
        v6 -= genFn1()
    }
    switch ("hello" + "hello") {
    case "":
        if v16 {
            var v39 int = v2
            _ = v39
        }
    case "hello":
        var v40 float64 = -2.25
        _ = v40
    default:
        var v41 float64 = (0.1 - 1e10)
        _ = v41
    }
    v42 := GenS2{F0: "\n\t", F1: v2, F2: v4}
    var v43 uint8 = 128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
