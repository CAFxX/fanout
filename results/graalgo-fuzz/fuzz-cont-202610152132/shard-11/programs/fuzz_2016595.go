package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 string; F2 uint8 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 uint16; F1 int32; F2 uint16 }
func genFn0() uint64 {
    var v7 uint = 42
    for fi8 := 0; fi8 < 2; fi8++ {
        for fi9 := 0; fi9 < 2; fi9++ {
            v10 := GenS1{F0: 0}
            _ = fi9
            _ = v10
        }
        if (!((v7 & (v7 * (v7 ^ (v7 | (v7 / (((^(v7))) | 1)))))) != v7)) {
            fmt.Println(fi8, fi8)
            switch fi8 {
            case 551:
                var v11 uint8 = 135
                _ = v11
            case -713:
                for fi12 := 0; fi12 < 9; fi12++ {
                    if false {
                        v13 := []int{fi8, fi12, fi12}
                        _ = v13
                    }
                    _ = fi12
                }
            }
            v14 := GenS2{F0: 65535, F1: "foo bar", F2: 91}
            _ = v14
        }
        _ = fi8
    }
    v7 = v7
    v7 += 4294967295
    _ = v7
    return 1
}
func genFn1(p0 int8) int64 {
    p0 -= (p0 / ((p0) | 1))
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return int64(p0)
}

func main() {
    var v1 uint16 = 42
    v2 := make(map[string]int)
    v3 := []uint32{2147483647, 2, 346}
    v4 := []string{"foo bar", "foo bar"}
    var v5 uint = 0
    var v6 float32 = 3.14159
    var v15 int32 = -879
    switch (("\n\t" + "hello") + "\n\t") {
    case "\n\t":
        if false {
            fmt.Println(v15)
            v6 = 1.5
        }
    case "":
        for fi16 := 0; fi16 < 2; fi16++ {
            var v17 int32 = -916
            if false {
                _ = len(v3)
                v1 -= (v1 * v1)
                fmt.Println(v17, v6)
            } else {
                for fi18 := 0; fi18 < 19; fi18++ {
                    v1 = 65535
                    _ = fi18
                }
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi16
            _ = v17
        }
    default:
        v19 := GenS1{F0: -2147483648}
        _ = v19
    }
    if false {
        var v20 int32 = -2147483648
        fmt.Println(v1, v1)
        v21 := v2["foo bar"]
        _ = v20
        _ = v21
    }
    var v22 uint16 = v1
    var v23 float32 = v6
    v24 := GenS3{F0: 0, F1: -32768}
    var v25 uint = (v5 | 18446744073709551615)
    var v26 bool = (1 > 4294967295)
    var v27 int8 = 112
    if true {
        var v28 uint32 = 0
        var v29 uint32 = v28
        _ = v28
        _ = v29
    }
    v30 := GenS4{F0: (^(v1)), F1: 2147483647, F2: (v1 | 1)}
    if ((v1 << ((154) & 63)) == 32767) {
        var v31 int32 = -2147483648
        for fi32 := 0; fi32 < 11; fi32++ {
            var v33 int = (fi32 >> ((v5) & 63))
            _ = fi32
            _ = v33
        }
        if v26 {
            v22 -= (^(v22))
        }
        _ = v31
    } else {
        delete(v2, "a")
    }
    v2["\n\t"] = 9223372036854775807
    var v34 int8 = v27
    v27 = 81
    v35 := GenS1{F0: (v15 | 0)}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
