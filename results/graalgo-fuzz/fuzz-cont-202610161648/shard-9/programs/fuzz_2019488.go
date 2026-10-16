package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 int8; F1 int }
func genFn0(p0 uint8) uint16 {
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 string = "\n\t"
        _ = fi7
        _ = v8
    }
    fmt.Println(p0)
    fmt.Println(p0)
    for fi9 := 0; fi9 < 20; fi9++ {
        if true {
            v10 := []int{9223372036854775807}
            _ = v10
        } else {
            fmt.Println(fi9, p0)
            var v11 int = (^(fi9))
            _ = v11
        }
        fmt.Println(p0, p0)
        v12 := []int{fi9, (-(fi9)), (fi9 | (fi9 << ((0) & 63))), (fi9 << ((2) & 63))}
        _ = fi9
        _ = v12
    }
    v13 := GenS1{F0: 512, F1: (57 < 42)}
    _ = p0
    _ = v13
    return 42
}

func main() {
    var v1 int8 = 100
    v2 := make(map[string]int)
    var v3 uint16 = 0
    var v4 uint32 = 2
    var v5 float32 = 0.0
    var v6 int8 = v1
    v3 += 32767
    var v14 int8 = v1
    _ = len(v2)
    if (v5 >= 1e10) {
        fmt.Println(v3)
        if (18446744073709551615 <= 65535) {
            v6 += int8(v4)
            var v15 int16 = 42
            _ = v15
        } else {
            switch "" {
            case "hello":
                v5 = 0.0
            case "foo bar":
                var v16 int16 = 32767
                _ = v16
            case "":
                v2[("\n\t" + "\n\t")] = -681
            }
            v4 = (v4 >> ((2) & 63))
            for fi17 := 0; fi17 < 9; fi17++ {
                var v18 string = (("a" + ("hello" + "foo bar")) + "hello")
                v5 = 1e10
                v5 = (float32(float64((v5 + 1e10))) / 1.5)
                _ = fi17
                _ = v18
            }
        }
        v19 := GenS1{F0: 42, F1: true}
        _ = v19
    } else {
        fmt.Println(v5, v6)
    }
    v4 -= v4
    if (2147483647 > 42) {
        v3 = (v3 | v3)
    }
    v20 := GenS1{F0: 4294967295, F1: (!(v1 > (v1 & v6)))}
    _ = len(v2)
    v21 := GenS2{F0: -627, F1: 4294967295}
    v3 -= v3
    v22 := []int{}
    fmt.Println(v6)
    if (true || false) {
        if false {
            if len(v22) > 0 {
                v22[0] = -758
            }
        }
    }
    v5 = 1.5
    for fi23 := 0; fi23 < 17; fi23++ {
        var v24 uint64 = 2
        for fi25 := 0; fi25 < 9; fi25++ {
            v14 = (v1 / (((v1 | 125)) | 1))
            v22 = append(v22, 0)
            _ = fi25
        }
        delete(v2, "\n\t")
        _ = fi23
        _ = v24
    }
    delete(v2, ("" + "foo bar"))
    v26 := GenS3{F0: (v1 | v1), F1: -1}
    var v27 uint64 = 998
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
