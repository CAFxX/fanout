package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint16 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int8) uint {
    switch 9223372036854775807 {
    case -9223372036854775808:
        for fi6 := 0; fi6 < 15; fi6++ {
            switch "a" {
            case "hello":
                v7 := []int{fi6}
                _ = v7
            default:
                var v8 uint64 = 2
                _ = v8
            }
            if false {
                p0 += p0
            } else {
                var v9 int8 = p0
                if (!true) {
                    v9 -= v9
                    var v10 uint32 = 65535
                    _ = v10
                }
                var v11 float32 = 0.1
                _ = v11
                _ = v9
            }
            p0 += p0
            _ = fi6
        }
    case 2147483647:
        p0 += (p0 & p0)
    case 32768:
        var v12 uint64 = 42
        _ = v12
    }
    v13 := GenS1{F0: (("foo bar" + ("" + "")) + ("\n\t" + ("foo bar" + "a"))), F1: "\n\t", F2: 32767}
    _ = p0
    _ = v13
    return 1
}

func main() {
    var v1 int = 4294967295
    v2 := []int{}
    v3 := make(map[int]int)
    v4 := []uint32{2147483647}
    v5 := make(map[string]string)
    v1 = 4294967295
    v14 := GenS1{F0: "hello", F1: "", F2: 1}
    v1 += (v1 << ((65535) & 63))
    for fi15 := 0; fi15 < 13; fi15++ {
        if true {
            if len(v2) > 0 {
                v2[0] = 4294967295
            }
        } else {
            for fi16 := 0; fi16 < 9; fi16++ {
                v1 = 775
                v1 += fi16
                _ = fi16
            }
        }
        v17 := GenS2{F0: 17}
        var v18 uint8 = 1
        _ = fi15
        _ = v17
        _ = v18
    }
    if true {
        var v19 int32 = 42
        var v20 uint = 468
        var v21 string = "\n\t"
        _ = v19
        _ = v20
        _ = v21
    } else {
        var v22 float32 = (1.5 / (float32(-2.25) * 1.5))
        if len(v4) > 0 {
            v4[0] = 1
        }
        for fi23 := 0; fi23 < 8; fi23++ {
            v24 := GenS2{F0: 2}
            _ = fi23
            _ = v24
        }
        _ = v22
    }
    var v25 uint8 = 128
    v25 -= (v25 & v25)
    if (!true) {
        v26 := GenS2{F0: (v25 << ((255) & 63))}
        _ = v26
    } else {
        var v27 int32 = -2147483648
        _ = v27
    }
    v5[(((("\n\t" + "\n\t") + "foo bar") + "") + "a")] = "a"
    fmt.Println(v1, v25)
    v25 = (v25 % ((v25) | 1))
    var v28 int = 32768
    for fi29 := 0; fi29 < 9; fi29++ {
        for fi30 := 0; fi30 < 17; fi30++ {
            var v31 int8 = -40
            _ = fi30
            _ = v31
        }
        v28 = v28
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
