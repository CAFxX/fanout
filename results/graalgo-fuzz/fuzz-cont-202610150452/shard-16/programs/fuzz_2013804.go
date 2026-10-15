package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int }
type GenS2 struct { F0 int32; F1 int32; F2 int32 }
func genFn0(p0 int, p1 uint, p2 string) int16 {
    if (127 == 65) {
        p0 = p0
    } else {
        p2 = p2
        p2 = p2
        for fi4 := 0; fi4 < 4; fi4++ {
            var v5 uint64 = 18446744073709551615
            for fi6 := 0; fi6 < 19; fi6++ {
                v7 := []int{(fi6 << ((p1) & 63))}
                if len(v7) > 0 {
                    v7[0] = (fi4 - 602)
                }
                _ = fi6
                _ = v7
            }
            _ = fi4
            _ = v5
        }
    }
    v8 := []int{p0}
    switch ("" + ((p2 + (p2 + p2)) + ("\n\t" + p2))) {
    case "foo bar":
        fmt.Println(p0, p0)
    case "":
        for fi9 := 0; fi9 < 13; fi9++ {
            var v10 int32 = -359
            if len(v8) > 0 {
                v8[0] = (p0 % (((fi9 << ((1) & 63))) | 1))
            }
            var v11 int = p0
            _ = fi9
            _ = v10
            _ = v11
        }
    case "hello":
        var v12 uint8 = 42
        _ = v12
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 870
}

func main() {
    var v1 uint64 = 500
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]int)
    var v13 int32 = 671
    v1 = (v1 % (((^(v1))) | 1))
    switch 4294967295 {
    case 0:
        v14 := GenS1{F0: 1, F1: 666, F2: int(v13)}
        _ = v14
    case -828:
        var v15 int8 = 42
        _ = v15
    }
    switch 2147483647 {
    case -761:
        v13 -= (v13 / ((482) | 1))
    }
    fmt.Println(v1, v13)
    var v16 string = ""
    if false {
        fmt.Println(v1)
    } else {
        v13 = v13
    }
    v16 = (v16 + v16)
    fmt.Println(v13, v16)
    v16 = v16
    if (0.0 < (float64(1e10) * 3.14159)) {
        v17 := []int{}
        fmt.Println(v13)
        _ = v17
    } else {
        v16 = "a"
        v13 -= -1
        if true {
            for fi18 := 0; fi18 < 12; fi18++ {
                if true {
                    v19 := []int{(^(fi18))}
                    _ = v19
                }
                var v20 float32 = -2.25
                _ = fi18
                _ = v20
            }
            if (127 < 1) {
                v21 := GenS2{F0: (-(v13)), F1: v13, F2: 2147483647}
                _ = v21
            } else {
                v16 = (((((("foo bar" + v16) + (v16 + v16)) + v16) + v16) + "") + "hello")
                var v22 int16 = -773
                _ = v22
            }
        }
    }
    for fi23 := 0; fi23 < 1; fi23++ {
        var v24 uint = 2
        _ = fi23
        _ = v24
    }
    v13 = v13
    v3[0] = 32768
    v1 += v1
    if (!((((v16 + "foo bar") + (((v16 + "a") + v16) + (v16 + ((((v16 + ((v16 + "\n\t") + v16)) + (v16 + v16)) + (v16 + ("" + "foo bar"))) + ("foo bar" + v16))))) + "foo bar") < "\n\t")) {
        var v25 int32 = v13
        _ = v25
    }
    v26 := GenS1{F0: 109, F1: 65535, F2: 417}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v26
    _ = v3
}
