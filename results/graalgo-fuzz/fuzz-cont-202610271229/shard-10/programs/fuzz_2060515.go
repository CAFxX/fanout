package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint8; F1 uint8; F2 int16 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int64) uint32 {
    if ((-2.25 + (0.0 * (-2.25 - -2.25))) >= 0.1) {
        v7 := []int{}
        _ = v7
    } else {
        if false {
            if false {
                v8 := GenS1{F0: 271, F1: (-(p0)), F2: 127}
                p0 = (p0 % ((p0) | 1))
                _ = v8
            } else {
                p0 -= 0
            }
            fmt.Println(p0)
            var v9 float32 = 1.5
            _ = v9
        } else {
            var v10 int16 = -1
            v11 := []int{4294967295, 612}
            _ = len(v11)
            _ = v10
            _ = v11
        }
    }
    var v12 float64 = float64(float32(3.14159))
    p0 = p0
    var v13 int16 = 0
    fmt.Println(v13)
    _ = p0
    _ = v12
    _ = v13
    return 100
}

func main() {
    var v1 float64 = 3.14159
    var v2 uint = 0
    v3 := make(map[uint32]string)
    var v4 int64 = 9223372036854775807
    var v5 uint64 = 0
    var v6 float32 = (((-2.25 / ((0.1 / 1e10) * (1e10 + 1e10))) * ((1.5 - 0.1) + 1e10)) / ((float32(v1) / 3.14159) + (1.5 * (float32(3.14159) * (((3.14159 / 1e10) / -2.25) + float32((-2.25 / float64((((1e10 + (float32(v1) + ((-2.25 - 1.5) - 1e10))) + 0.0) + float32(v1))))))))))
    for fi14 := 0; fi14 < 2; fi14++ {
        var v15 int = fi14
        var v16 int16 = -581
        var v17 string = "hello"
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    fmt.Println(v2, v2)
    for fi18 := 0; fi18 < 11; fi18++ {
        var v19 uint = v2
        fmt.Println(v2, v6, v4)
        _ = fi18
        _ = v19
    }
    v20 := v3[2147483648]
    var v21 uint32 = uint32(v5)
    for fi22 := 0; fi22 < 14; fi22++ {
        switch 128 {
        case 4294967295:
            v23 := GenS1{F0: (v2 % (((v2 - 2)) | 1)), F1: -1, F2: 221}
            _ = v23
        default:
            if false {
                var v24 int32 = 0
                _ = v24
            }
        }
        v25 := []int{(fi22 >> (((v2 % (((v2 * 0)) | 1))) & 63)), fi22}
        _ = fi22
        _ = v25
    }
    v5 += (v5 * (v5 / ((1) | 1)))
    var v26 uint32 = v21
    for fi27 := 0; fi27 < 5; fi27++ {
        for fi28 := 0; fi28 < 18; fi28++ {
            v6 = (1.5 + -2.25)
            v2 -= (^(v2))
            v2 = 18446744073709551615
            _ = fi28
        }
        fmt.Println(v1, fi27)
        _ = fi27
    }
    for fi29 := 0; fi29 < 17; fi29++ {
        for fi30 := 0; fi30 < 20; fi30++ {
            switch ((("hello" + ("\n\t" + (("\n\t" + (v20 + "hello")) + "hello"))) + v20) + "hello") {
            case "a":
                v31 := GenS1{F0: (^(v2)), F1: 524, F2: 246}
                var v32 int8 = int8(v4)
                v2 += 433
                _ = v31
                _ = v32
            case "\n\t":
                v33 := GenS1{F0: 761, F1: (v4 ^ (v4 ^ (v4 + (^(v4))))), F2: 128}
                _ = v33
            }
            v26 = (v21 % (((^(v21))) | 1))
            var v34 int8 = 1
            if true {
                fmt.Println(v20)
            }
            v1 = -2.25
            _ = fi30
            _ = v34
        }
        fmt.Println(v5, v5)
        switch (fi29 | -177) {
        case -2147483648:
            for fi35 := 0; fi35 < 4; fi35++ {
                var v36 string = v20
                v26 += (v26 + genFn0(v4))
                _ = fi35
                _ = v36
            }
        }
        _ = fi29
    }
    var v37 bool = false
    fmt.Println(v26, v21)
    v38 := GenS2{F0: 2, F1: 0, F2: 42}
    v39 := []int{}
    v40 := GenS3{F0: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
