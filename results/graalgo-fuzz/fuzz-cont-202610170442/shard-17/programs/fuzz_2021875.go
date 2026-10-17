package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int32 }
type GenS2 struct { F0 uint32; F1 uint64; F2 int8 }
type GenS3 struct { F0 int8; F1 string; F2 int32 }
func genFn0() bool {
    if true {
        v5 := GenS1{F0: "hello", F1: "foo bar", F2: -430}
        var v6 uint32 = 42
        for fi7 := 0; fi7 < 13; fi7++ {
            v8 := []int{fi7}
            if (100 <= -128) {
                if (2 != 968) {
                    var v9 int = (fi7 ^ fi7)
                    var v10 uint8 = 127
                    _ = v10
                    _ = v9
                } else {
                    for fi11 := 0; fi11 < 11; fi11++ {
                        switch (fi11 - (fi11 >> ((2) & 63))) {
                        case 32767:
                            v6 += 1
                            fmt.Println(fi7)
                        case 108:
                            v6 += (v6 ^ v6)
                        default:
                            fmt.Println(fi11)
                        }
                        if true {
                            v6 = 42
                            var v12 string = ("" + ("\n\t" + "foo bar"))
                            v8 = append(v8, (fi11 - (fi7 | 615)))
                            _ = v12
                        } else {
                            v6 -= (v6 | v6)
                        }
                        _ = len(v8)
                        _ = fi11
                    }
                    v6 = v6
                    var v13 float32 = 0.1
                    _ = v13
                }
                v6 -= 42
            }
            _ = fi7
            _ = v8
        }
        _ = v5
        _ = v6
    }
    var v14 uint8 = 128
    fmt.Println(v14)
    _ = v14
    return false
}
func genFn1(p0 uint8) int8 {
    var v15 float64 = -2.25
    var v16 float64 = float64(3.14159)
    p0 = (p0 << ((uint(p0)) & 63))
    var v17 uint16 = 42
    var v18 uint8 = 0
    _ = p0
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    return 42
}

func main() {
    var v1 uint = 4294967295
    var v2 uint32 = 0
    var v3 int64 = -1
    var v4 uint64 = 613
    var v19 string = ""
    for fi20 := 0; fi20 < 11; fi20++ {
        v21 := []int{fi20, fi20}
        var v22 float32 = 1e10
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 int = 4294967295
    var v24 uint32 = 2147483647
    var v25 float32 = 0.0
    for fi26 := 0; fi26 < 1; fi26++ {
        v19 = ("hello" + "foo bar")
        fmt.Println(v23, v2, v24)
        _ = fi26
    }
    if (-2.25 != 0.1) {
        v27 := GenS2{F0: v24, F1: v4, F2: 105}
        if (v2 == (^(v2))) {
            v24 = uint32(v1)
        } else {
            v23 += v23
        }
        for fi28 := 0; fi28 < 11; fi28++ {
            for fi29 := 0; fi29 < 18; fi29++ {
                if false {
                    var v30 int16 = -299
                    var v31 uint64 = v4
                    _ = v30
                    _ = v31
                }
                _ = fi29
            }
            v24 += (v24 >> ((v1) & 63))
            var v32 int64 = (v3 / (((v3 ^ v3)) | 1))
            _ = fi28
            _ = v32
        }
        _ = v27
    } else {
        for fi33 := 0; fi33 < 2; fi33++ {
            var v34 int64 = v3
            _ = fi33
            _ = v34
        }
    }
    var v35 int8 = 67
    var v36 uint64 = 733
    for fi37 := 0; fi37 < 5; fi37++ {
        v38 := GenS2{F0: (v2 >> ((v1) & 63)), F1: v4, F2: (v35 << (((^(v1))) & 63))}
        v39 := []int{}
        _ = fi37
        _ = v38
        _ = v39
    }
    if (v19 < (("a" + v19) + v19)) {
        switch v23 {
        case 0:
            v40 := GenS2{F0: 889, F1: v36, F2: (v35 * 127)}
            _ = v40
        }
    }
    switch (-(v23)) {
    case -2147483648:
        v25 += ((v25 * v25) / v25)
    case 2147483647:
        for fi41 := 0; fi41 < 20; fi41++ {
            v42 := GenS3{F0: (v35 & v35), F1: (v19 + ("\n\t" + "\n\t")), F2: 42}
            fmt.Println(v19, v2, v24)
            _ = fi41
            _ = v42
        }
    case -706:
        v43 := GenS1{F0: "", F1: (v19 + v19), F2: 36}
        _ = v43
    default:
        var v44 uint16 = 219
        _ = v44
    }
    var v45 uint32 = (^(v24))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v45
}
