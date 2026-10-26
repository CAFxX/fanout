package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int16 }
type GenS2 struct { F0 uint; F1 int16; F2 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 int64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint; F1 uint32 }
func genFn0(p0 uint64, p1 uint16, p2 string) int16 {
    switch "hello" {
    case "hello":
        fmt.Println(p2, p2, p2)
    case "":
        var v5 uint64 = 372
        _ = v5
    case "foo bar":
        switch 128 {
        case 42:
            v6 := []int{}
            _ = v6
        case 551:
            var v7 int = -410
            _ = v7
        }
    default:
        var v8 uint8 = 127
        _ = v8
    }
    switch -327 {
    case 32767:
        for fi9 := 0; fi9 < 5; fi9++ {
            p2 = "\n\t"
            var v10 uint = 905
            _ = fi9
            _ = v10
        }
    case -2147483648:
        for fi11 := 0; fi11 < 16; fi11++ {
            if false {
                var v12 float64 = (0.1 - -2.25)
                if true {
                    switch (p2 + p2) {
                    case "\n\t":
                        switch fi11 {
                        case 85:
                            p1 = (p1 / ((p1) | 1))
                        case 9223372036854775807:
                            p0 -= (p0 / ((955) | 1))
                        case 32768:
                            v13 := GenS1{F0: p2, F1: 65535, F2: 1000}
                            _ = v13
                        default:
                            var v14 uint = 354
                            _ = v14
                        }
                    }
                }
                _ = v12
            } else {
                var v15 string = p2
                _ = v15
            }
            var v16 string = ("\n\t" + p2)
            _ = fi11
            _ = v16
        }
    case -394:
        var v17 uint = 1
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    return -779
}
func genFn1(p0 uint32, p1 int, p2 int8) uint16 {
    if (!true) {
        switch p1 {
        case 42:
            switch "foo bar" {
            case "":
                p0 = (p0 / ((2147483647) | 1))
            }
        case 562:
            p2 = 42
        case 430:
            switch "" {
            case "":
                for fi18 := 0; fi18 < 2; fi18++ {
                    for fi19 := 0; fi19 < 3; fi19++ {
                        p2 = (p2 + (p2 - (p2 / (((p2 / (((p2 | p2)) | 1))) | 1))))
                        var v20 int8 = p2
                        v21 := GenS2{F0: 1, F1: -32768, F2: -1}
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    _ = fi18
                }
            default:
                p0 += p0
            }
        default:
            v22 := GenS2{F0: 334, F1: 975, F2: -1}
            _ = v22
        }
    }
    v23 := []int{p1}
    if false {
        var v24 bool = (true || true)
        var v25 float64 = 1.5
        _ = v24
        _ = v25
    }
    v26 := GenS3{F0: 543}
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    _ = v26
    return 1
}

func main() {
    v1 := []int{}
    v2 := []string{"a", "\n\t", ("foo bar" + "\n\t"), ("a" + "a")}
    v3 := []int{349, 387}
    var v4 int16 = -32768
    if true {
        v4 = (v4 >> ((950) & 63))
        v1 = append(v1, 4294967295)
    }
    v27 := GenS4{F0: -9223372036854775808, F1: 9223372036854775807}
    v4 += (^(v4))
    var v28 int64 = 611
    if len(v3) > 0 {
        v3[0] = 128
    }
    var v29 float32 = 0.0
    var v30 uint = 65535
    v31 := GenS4{F0: v28, F1: v28}
    switch -768 {
    case 153:
        if true {
            var v32 int8 = 100
            _ = v32
        }
    default:
        v33 := GenS5{F0: 1}
        _ = v33
    }
    v34 := GenS4{F0: v28, F1: v28}
    v30 = v30
    v35 := GenS6{F0: v30, F1: 524}
    v36 := GenS6{F0: v30, F1: 4294967295}
    for fi37 := 0; fi37 < 9; fi37++ {
        var v38 uint8 = 128
        v39 := GenS5{F0: -331}
        _ = fi37
        _ = v38
        _ = v39
    }
    var v40 int = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v40
}
