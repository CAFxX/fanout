package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int }
func genFn0() int32 {
    for fi5 := 0; fi5 < 12; fi5++ {
        var v6 uint32 = 2147483648
        _ = fi5
        _ = v6
    }
    var v7 int8 = -128
    _ = v7
    return 0
}
func genFn1(p0 uint, p1 uint64) uint8 {
    for fi8 := 0; fi8 < 3; fi8++ {
        if (false && ((true && (true || false)) || (!true))) {
            var v9 int32 = 2147483647
            var v10 float64 = 0.0
            _ = v10
            _ = v9
        } else {
            v11 := []int{(-(fi8)), (fi8 % ((fi8) | 1))}
            _ = v11
        }
        _ = fi8
    }
    v12 := []int{}
    _ = p0
    _ = p1
    _ = v12
    return 255
}

func main() {
    var v1 int16 = 42
    v2 := make(map[uint32]int)
    var v3 uint64 = 42
    v4 := make(map[string]bool)
    v13 := GenS1{F0: 2147483648, F1: 9223372036854775807}
    var v14 uint = 368
    for fi15 := 0; fi15 < 8; fi15++ {
        v3 += v3
        var v16 int8 = 72
        v2[42] = (fi15 | fi15)
        _ = fi15
        _ = v16
    }
    switch ("" + "") {
    case "a":
        if false {
            fmt.Println(v14)
            fmt.Println(v3, v1)
        } else {
            v14 = v14
            v1 += v1
            v17 := GenS2{F0: v14}
            _ = v17
        }
    case "foo bar":
        fmt.Println(v14)
    }
    v14 += 786
    if (474 == 65535) {
        var v18 uint8 = genFn1((^(v14)), (^(v3)))
        _ = v18
    } else {
        var v19 float32 = (1.5 + float32(0.1))
        v20 := GenS1{F0: 7, F1: 1}
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        var v22 uint16 = 32767
        v23 := GenS2{F0: v14}
        _ = fi21
        _ = v22
        _ = v23
    }
    v1 -= v1
    for fi24 := 0; fi24 < 17; fi24++ {
        switch fi24 {
        case 128:
            v14 = (v14 ^ v14)
        case -897:
            v25 := v4["a"]
            _ = v25
        case 2147483647:
            switch "\n\t" {
            case "foo bar":
                fmt.Println(v1)
            default:
                for fi26 := 0; fi26 < 9; fi26++ {
                    v14 -= 1
                    v27 := GenS3{F0: 1}
                    var v28 bool = (153 <= 42)
                    v29 := GenS1{F0: 908, F1: -9223372036854775808}
                    fmt.Println(fi26)
                    _ = fi26
                    _ = v27
                    _ = v28
                    _ = v29
                }
            }
        default:
            _ = len(v4)
        }
        for fi30 := 0; fi30 < 20; fi30++ {
            if true {
                var v31 int32 = genFn0()
                v3 = v3
                _ = v31
            }
            v14 -= (v14 / (((^(v14))) | 1))
            var v32 uint16 = 171
            v14 = v14
            _ = fi30
            _ = v32
        }
        v3 -= (v3 | (^(v3)))
        _ = fi24
    }
    for fi33 := 0; fi33 < 5; fi33++ {
        v14 += 42
        _ = fi33
    }
    v34 := GenS4{F0: ("\n\t" != ((("hello" + (("\n\t" + "\n\t") + "hello")) + (("\n\t" + ("hello" + "")) + "")) + ("a" + "a")))}
    if false {
        v35 := []int{1, 32768, 4294967295}
        var v36 string = "a"
        _ = v35
        _ = v36
    }
    v1 += (^(v1))
    for fi37 := 0; fi37 < 19; fi37++ {
        if (true || true) {
            v38 := GenS5{F0: 2147483647}
            v1 = (v1 & -891)
            switch 42 {
            case 796:
                var v39 float64 = float64(-2.25)
                _ = v39
            }
            _ = v38
        } else {
            if true {
                v40 := []int{fi37, (-(fi37)), fi37, fi37}
                _ = v40
            } else {
                for fi41 := 0; fi41 < 15; fi41++ {
                    var v42 int16 = int16(fi41)
                    _ = fi41
                    _ = v42
                }
            }
            var v43 bool = false
            for fi44 := 0; fi44 < 16; fi44++ {
                var v45 uint8 = 128
                switch (fi44 & (fi44 * (fi37 - (fi37 ^ -847)))) {
                case 0:
                    var v46 uint16 = 42
                    _ = v46
                case 2147483647:
                    if v43 {
                        v47 := GenS6{F0: (fi44 % ((int(v45)) | 1))}
                        var v48 string = "hello"
                        v1 = v1
                        _ = v47
                        _ = v48
                    }
                case 4294967295:
                    var v49 uint = 255
                    _ = v49
                }
                _ = fi44
                _ = v45
            }
            _ = v43
        }
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v34
    _ = v4
}
