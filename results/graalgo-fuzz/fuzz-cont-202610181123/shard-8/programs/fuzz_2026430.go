package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 int64; F1 uint8; F2 uint }
type GenS6 struct { F0 uint64 }
func genFn0(p0 string) int64 {
    switch "hello" {
    case "hello":
        var v7 bool = ("hello" <= "\n\t")
        _ = v7
    case "foo bar":
        v8 := []int{0, -9223372036854775808}
        _ = v8
    }
    var v9 int16 = 0
    var v10 uint16 = 42
    _ = p0
    _ = v10
    _ = v9
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]bool)
    var v4 uint64 = 0
    var v5 int8 = 1
    v6 := make(map[int]uint32)
    for fi11 := 0; fi11 < 6; fi11++ {
        v12 := GenS1{F0: (v5 + -48), F1: fi11}
        _ = fi11
        _ = v12
    }
    fmt.Println(v5, v4)
    var v13 uint8 = 1
    if (0 <= 0) {
        if (127 == v5) {
            var v14 uint64 = 1
            var v15 int32 = -2147483648
            for fi16 := 0; fi16 < 2; fi16++ {
                for fi17 := 0; fi17 < 9; fi17++ {
                    var v18 uint = 65535
                    v15 += (v15 | (v15 % ((v15) | 1)))
                    _ = fi17
                    _ = v18
                }
                _ = fi16
            }
            _ = v14
            _ = v15
        } else {
            var v19 bool = false
            _ = v19
        }
    }
    v13 -= (v13 & 128)
    v20 := GenS2{F0: int16(v13)}
    if true {
        var v21 string = ("\n\t" + "hello")
        v22 := GenS3{F0: 138}
        var v23 float64 = (1.5 * (((0.0 + 1e10) * float64(3.14159)) * (float64((1e10 - ((float32(((1e10 * 1.5) * 0.0)) / float32(3.14159)) * 1.5))) + 0.0)))
        _ = v21
        _ = v22
        _ = v23
    }
    switch ("a" + "a") {
    case "\n\t":
        if ((("a" + "foo bar") + "\n\t") != "a") {
            if true {
                v24 := []int{}
                v5 = v5
                _ = v24
            } else {
                v25 := GenS4{F0: 4294967295, F1: 42}
                _ = v25
            }
            for fi26 := 0; fi26 < 4; fi26++ {
                var v27 uint64 = (^(v4))
                switch 99 {
                case 2147483647:
                    v28 := GenS4{F0: 32768, F1: 1}
                    _ = v28
                }
                for fi29 := 0; fi29 < 6; fi29++ {
                    var v30 string = ""
                    fmt.Println(v30, v4)
                    var v31 int = (fi29 / (((fi26 % (((fi26 - (fi29 & (fi29 - fi29)))) | 1))) | 1))
                    _ = len(v3)
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                _ = fi26
                _ = v27
            }
            if false {
                fmt.Println(v4)
                var v32 int8 = (v5 | v5)
                _ = v32
            }
        }
    case "a":
        var v33 bool = ((42 != 2) && (false || (!false)))
        _ = v33
    case "":
        var v34 uint = 0
        _ = v34
    default:
        switch "" {
        case "\n\t":
            for fi35 := 0; fi35 < 20; fi35++ {
                v13 = 42
                v36 := []int{}
                _ = fi35
                _ = v36
            }
        case "":
            v13 -= 42
        case "hello":
            for fi37 := 0; fi37 < 12; fi37++ {
                delete(v3, 264)
                v4 -= (v4 - v4)
                _ = fi37
            }
        }
    }
    if true {
        fmt.Println(v5, v5)
        v38 := GenS5{F0: 339, F1: 128, F2: 42}
        var v39 uint32 = 0
        _ = v38
        _ = v39
    } else {
        for fi40 := 0; fi40 < 11; fi40++ {
            var v41 uint16 = 32767
            var v42 int = fi40
            _ = fi40
            _ = v41
            _ = v42
        }
        if true {
            v5 += (v5 % (((v5 << ((1) & 63))) | 1))
            var v43 int64 = -1
            var v44 int64 = v43
            _ = v43
            _ = v44
        }
    }
    switch -2147483648 {
    case -9223372036854775808:
        var v45 float32 = (((((0.0 - 1e10) * 3.14159) / (0.1 * (((((0.1 * 3.14159) * 0.1) + float32(((-2.25 + 1e10) + 1.5))) + ((3.14159 - 1.5) * ((0.0 + 1.5) / 0.1))) / 1.5))) - ((0.0 / float32((1e10 * 3.14159))) * 0.1)) - 1.5)
        _ = v45
    case 42:
        fmt.Println(v13, v13, v5)
    case 4294967295:
        v5 = (v5 << ((2) & 63))
    default:
        v13 += v13
    }
    switch "" {
    case "hello":
        fmt.Println(v13, v4, v4)
    case "foo bar":
        v46 := GenS5{F0: genFn0(("" + "a")), F1: (v13 % ((v13) | 1)), F2: 4294967295}
        _ = v46
    case "":
        _ = len(v3)
    default:
        v47 := v3[0]
        _ = v47
    }
    v48 := GenS1{F0: 0, F1: 9223372036854775807}
    fmt.Println(v4, v13, v13)
    v49 := GenS6{F0: (^(v4))}
    v4 = (v4 + (v4 - 2))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
