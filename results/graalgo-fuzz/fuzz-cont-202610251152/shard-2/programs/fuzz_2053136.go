package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 int; F2 int16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int32 }
func genFn0(p0 uint32, p1 bool) uint16 {
    if ((((((1e10 * -2.25) + 1.5) == 0.1) && (p0 < 2)) && false) && (1 != 1)) {
        for fi7 := 0; fi7 < 9; fi7++ {
            var v8 int8 = 42
            fmt.Println(p1, fi7)
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 uint = 18446744073709551615
            fmt.Println(v10, p0)
            _ = fi9
            _ = v10
        }
        fmt.Println(p0)
    } else {
        var v11 int = -2147483648
        var v12 bool = p1
        v13 := GenS1{F0: -1, F1: 1, F2: 1}
        _ = v11
        _ = v12
        _ = v13
    }
    v14 := GenS2{F0: (1e10 > float64(1.5))}
    var v15 float32 = ((0.1 + 3.14159) - 1e10)
    var v16 int32 = -909
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    _ = v16
    return 703
}

func main() {
    var v1 string = ("" + "hello")
    v2 := make(map[uint32]int)
    var v3 int16 = -552
    var v4 int32 = 42
    var v5 uint = 4294967295
    var v6 int32 = (v4 / ((42) | 1))
    fmt.Println(v3, v1, v6)
    switch 437 {
    case 0:
        v4 += (^(v6))
    case -2147483648:
        var v17 int8 = 0
        _ = v17
    case 9223372036854775807:
        if (!false) {
            for fi18 := 0; fi18 < 9; fi18++ {
                v19 := GenS3{F0: 127, F1: fi18, F2: v3}
                var v20 bool = false
                if (42 != 24) {
                    fmt.Println(fi18)
                    if (127 == -128) {
                        var v21 int = (^(fi18))
                        fmt.Println(fi18)
                        _ = v21
                    }
                }
                _ = fi18
                _ = v19
                _ = v20
            }
        }
    }
    v5 -= v5
    var v22 uint32 = 2147483648
    for fi23 := 0; fi23 < 4; fi23++ {
        for fi24 := 0; fi24 < 2; fi24++ {
            v25 := []int{-1}
            _ = fi24
            _ = v25
        }
        if (v22 == 2) {
            var v26 int32 = v6
            _ = v26
        }
        _ = fi23
    }
    if false {
        v3 = v3
        v2[(v22 - (v22 / ((1) | 1)))] = 32768
        for fi27 := 0; fi27 < 6; fi27++ {
            var v28 float64 = 1.5
            for fi29 := 0; fi29 < 14; fi29++ {
                var v30 uint = (^(v5))
                switch "foo bar" {
                case "foo bar":
                    v22 = v22
                case "":
                    v31 := []int{fi27, fi27, fi29}
                    _ = v31
                case "a":
                    v28 += float64(-2.25)
                }
                if (false && (!true)) {
                    v1 = v1
                }
                _ = fi29
                _ = v30
            }
            v2[v22] = (fi27 | fi27)
            _ = fi27
            _ = v28
        }
    } else {
        v32 := GenS4{F0: -1}
        var v33 bool = (!true)
        var v34 int16 = v3
        _ = v32
        _ = v33
        _ = v34
    }
    v35 := []int{32768, 128, -689, 32768}
    var v36 int64 = -98
    for fi37 := 0; fi37 < 14; fi37++ {
        for fi38 := 0; fi38 < 16; fi38++ {
            _ = len(v2)
            if true {
                switch ((("hello" + (v1 + v1)) + (((((v1 + v1) + "") + ("\n\t" + "")) + v1) + v1)) + ((v1 + v1) + ((("hello" + "a") + (v1 + "")) + v1))) {
                case "":
                    _ = len(v2)
                }
                v39 := GenS5{F0: 1}
                v3 += (v3 | (v3 >> ((1) & 63)))
                _ = v39
            }
            var v40 uint = (^(v5))
            _ = fi38
            _ = v40
        }
        for fi41 := 0; fi41 < 8; fi41++ {
            var v42 float32 = (3.14159 + 1e10)
            _ = fi41
            _ = v42
        }
        _ = fi37
    }
    v4 = v6
    var v43 uint8 = 1
    v44 := GenS6{F0: (v4 / (((v4 * 1)) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
