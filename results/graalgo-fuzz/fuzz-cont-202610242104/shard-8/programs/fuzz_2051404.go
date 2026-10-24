package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint32; F1 int32; F2 int32 }
func genFn0(p0 int64) int64 {
    if false {
        fmt.Println(p0)
        for fi6 := 0; fi6 < 3; fi6++ {
            var v7 uint8 = 255
            _ = fi6
            _ = v7
        }
        fmt.Println(p0)
    }
    if false {
        var v8 uint8 = 42
        for fi9 := 0; fi9 < 20; fi9++ {
            v8 = (v8 << ((4294967295) & 63))
            switch -596 {
            case -316:
                var v10 int = 674
                _ = v10
            }
            _ = fi9
        }
        if false {
            var v11 string = "foo bar"
            _ = v11
        }
        _ = v8
    }
    switch 992 {
    case 526:
        for fi12 := 0; fi12 < 14; fi12++ {
            var v13 int = (fi12 & fi12)
            v14 := GenS1{F0: 2147483647}
            _ = fi12
            _ = v13
            _ = v14
        }
    case 9223372036854775807:
        var v15 int32 = 1
        _ = v15
    }
    p0 = p0
    p0 -= 42
    _ = p0
    return (p0 % ((p0) | 1))
}
func genFn1() int8 {
    var v16 uint64 = 4294967295
    if false {
        v17 := []int{-2147483648, 2147483647}
        switch "\n\t" {
        case "hello":
            fmt.Println(v16)
        case "":
            v18 := GenS1{F0: -1}
            _ = v18
        case "foo bar":
            v16 -= (v16 / ((v16) | 1))
        default:
            var v19 bool = true
            _ = v19
        }
        v16 -= (v16 >> ((350) & 63))
        _ = v17
    }
    _ = v16
    return 108
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = 628
    var v3 int32 = (v2 / ((42) | 1))
    var v4 string = ("a" + "foo bar")
    v5 := make(map[int]string)
    v20 := GenS2{F0: 9223372036854775807, F1: 42, F2: 32767}
    v21 := GenS1{F0: (^(v3))}
    var v22 uint64 = 392
    for fi23 := 0; fi23 < 10; fi23++ {
        if (v2 < (v3 ^ (v2 - 432))) {
            v24 := v5[(fi23 >> ((42) & 63))]
            fmt.Println(v3, v4)
            for fi25 := 0; fi25 < 10; fi25++ {
                fmt.Println(v22, v3)
                _ = fi25
            }
            _ = v24
        } else {
            var v26 float64 = 0.1
            var v27 int = fi23
            _ = v26
            _ = v27
        }
        _ = fi23
    }
    v28 := []int{819, -519}
    if (4294967295 <= 2) {
        switch v4 {
        case "hello":
            v3 -= 0
        default:
            for fi29 := 0; fi29 < 18; fi29++ {
                var v30 float64 = 1e10
                var v31 uint64 = v22
                _ = fi29
                _ = v30
                _ = v31
            }
        }
        if (!(0 < 491)) {
            for fi32 := 0; fi32 < 13; fi32++ {
                v22 -= (v22 / (((^(v22))) | 1))
                fmt.Println(v22)
                var v33 uint16 = 2
                _ = fi32
                _ = v33
            }
            if false {
                if false {
                    switch -302 {
                    case 32767:
                        switch 32768 {
                        case 32767:
                            fmt.Println(v3, v3)
                        }
                    case 0:
                        v34 := GenS1{F0: v3}
                        _ = v34
                    default:
                        var v35 int8 = -119
                        _ = v35
                    }
                    if len(v28) >= 2 {
                        v28 = v28[1:len(v28)]
                    }
                }
            }
        }
    }
    v36 := GenS3{F0: 0, F1: v2, F2: -1}
    if true {
        fmt.Println(v22, v3, v3)
        for fi37 := 0; fi37 < 16; fi37++ {
            var v38 bool = ((1e10 > float32(1.5)) || true)
            var v39 int8 = -1
            _ = fi37
            _ = v38
            _ = v39
        }
    }
    var v40 bool = (!false)
    v41 := GenS1{F0: (^(v3))}
    for fi42 := 0; fi42 < 14; fi42++ {
        if v40 {
            var v43 uint64 = v22
            _ = v43
        }
        fmt.Println(v2, fi42)
        _ = fi42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v28
    _ = v3
    _ = v36
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
