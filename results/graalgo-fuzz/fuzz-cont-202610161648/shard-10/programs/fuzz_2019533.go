package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int64 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint32; F1 int16; F2 int64 }
func genFn0(p0 int16) int {
    if false {
        p0 -= p0
        fmt.Println(p0)
    } else {
        p0 = p0
        var v6 uint32 = 2
        _ = v6
    }
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 int32 = 0
            switch fi7 {
            case -1:
                fmt.Println(fi8)
            case 32768:
                v9 = -190
            case 812:
                if (42 <= -9223372036854775808) {
                    p0 = p0
                    var v10 uint32 = 4294967295
                    if (257 < 65535) {
                        p0 = int16(v9)
                    }
                    _ = v10
                }
            }
            v11 := []int{fi8, (fi7 & (fi7 & fi8)), (fi8 & 128), fi8}
            _ = fi8
            _ = v11
            _ = v9
        }
        _ = fi7
    }
    v12 := GenS1{F0: (("" + ("hello" + "\n\t")) == (("" + "\n\t") + ("" + "hello"))), F1: 0, F2: -312}
    fmt.Println(p0)
    _ = p0
    _ = v12
    return 911
}
func genFn1(p0 int8) int {
    fmt.Println(p0)
    fmt.Println(p0)
    switch "\n\t" {
    case "\n\t":
        for fi13 := 0; fi13 < 1; fi13++ {
            fmt.Println(p0, fi13)
            var v14 string = ""
            _ = fi13
            _ = v14
        }
    case "":
        var v15 uint16 = 255
        _ = v15
    default:
        p0 -= (p0 * -75)
    }
    _ = p0
    return 32767
}

func main() {
    v1 := []int{}
    var v2 string = ("\n\t" + "foo bar")
    var v3 uint64 = 0
    v4 := make(map[uint32]int)
    var v5 int16 = 1000
    v2 = (v2 + (((v2 + "a") + v2) + v2))
    for fi16 := 0; fi16 < 19; fi16++ {
        if len(v1) > 0 {
            v1[0] = fi16
        }
        switch (fi16 / ((fi16) | 1)) {
        case 4294967295:
            v2 = "\n\t"
        case 99:
            if true {
                var v17 int16 = v5
                v1 = append(v1, fi16)
                _ = v17
            }
        }
        _ = fi16
    }
    var v18 float32 = (((((1.5 * 0.1) * (0.0 * -2.25)) - float32(float64(0.1))) / (float32(3.14159) + 3.14159)) * (float32((1.5 * -2.25)) / 3.14159))
    v19 := GenS2{F0: 989, F1: 255}
    v20 := v4[uint32(v3)]
    if (!(42 <= 42)) {
        v21 := GenS3{F0: 4294967295, F1: v5, F2: 684}
        var v22 bool = false
        if true {
            var v23 uint16 = 32767
            _ = v23
        }
        _ = v21
        _ = v22
    }
    v2 = ("\n\t" + "foo bar")
    v24 := GenS3{F0: 1, F1: v5, F2: -1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
