package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint8 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint; F1 uint16 }
func genFn0(p0 int16, p1 int16, p2 uint) uint {
    var v7 float32 = 0.1
    v8 := []int{32767, -2147483648}
    v9 := GenS1{F0: 2, F1: p0, F2: 118}
    var v10 uint8 = 35
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return (p2 << ((p2) & 63))
}

func main() {
    var v1 int8 = 1
    v2 := make(map[string]int)
    v3 := make(map[string]uint32)
    var v4 uint16 = 32767
    v5 := make(map[int]int)
    var v6 uint16 = (v4 - v4)
    v11 := GenS2{F0: 32768, F1: -32768}
    _ = len(v2)
    var v12 float64 = 1e10
    v13 := GenS2{F0: 457, F1: 1}
    if (1e10 <= 3.14159) {
        for fi14 := 0; fi14 < 8; fi14++ {
            v6 += v6
            _ = fi14
        }
        fmt.Println(v4)
        v4 += (^(v6))
    }
    switch "hello" {
    case "\n\t":
        v12 = v12
    default:
        for fi15 := 0; fi15 < 19; fi15++ {
            for fi16 := 0; fi16 < 14; fi16++ {
                if (!((!true) || (false && (100 >= 42)))) {
                    var v17 float64 = v12
                    v18 := GenS3{F0: "\n\t", F1: "\n\t"}
                    delete(v3, ("\n\t" + ("a" + ("foo bar" + ""))))
                    _ = v17
                    _ = v18
                } else {
                    var v19 int64 = 0
                    var v20 float32 = float32(float64(float32(float64(1.5))))
                    v21 := GenS4{F0: 794}
                    var v22 int16 = 646
                    fmt.Println(v12, v19, v22)
                    _ = v19
                    _ = v20
                    _ = v21
                    _ = v22
                }
                v23 := GenS5{F0: 4294967295, F1: v4}
                if ((-2.25 < (0.1 * 0.1)) && true) {
                    var v24 uint8 = 128
                    fmt.Println(v1, v4)
                    _ = v24
                }
                _ = fi16
                _ = v23
            }
            v25 := []int{(-(fi15)), (^(fi15))}
            _ = fi15
            _ = v25
        }
    }
    var v26 string = (("foo bar" + "a") + ("\n\t" + "foo bar"))
    v27 := GenS2{F0: int(v1), F1: 111}
    v28 := GenS4{F0: -366}
    for fi29 := 0; fi29 < 10; fi29++ {
        v26 = (v26 + "hello")
        _ = fi29
    }
    var v30 float32 = (-2.25 / -2.25)
    if false {
        switch "\n\t" {
        case "\n\t":
            v31 := v5[-632]
            _ = v31
        }
        switch 128 {
        case -50:
            for fi32 := 0; fi32 < 7; fi32++ {
                v33 := []int{-307, 1}
                var v34 int16 = -32768
                _ = fi32
                _ = v33
                _ = v34
            }
        default:
            fmt.Println(v30)
        }
        var v35 int64 = 976
        _ = v35
    } else {
        switch 127 {
        case -9223372036854775808:
            v1 = v1
        case 127:
            var v36 uint16 = 42
            _ = v36
        case -1:
            v30 = 0.1
        }
        for fi37 := 0; fi37 < 2; fi37++ {
            var v38 int32 = -2147483648
            if true {
                v38 -= (-(v38))
                var v39 uint = 199
                v40 := GenS1{F0: v4, F1: -1, F2: 1}
                _ = v39
                _ = v40
            } else {
                v41 := []int{}
                if ((!(2 < 144)) || (v38 < (-(v38)))) {
                    var v42 bool = (-31 <= 0)
                    var v43 float32 = v30
                    _ = v42
                    _ = v43
                } else {
                    fmt.Println(v4)
                    _ = len(v2)
                    var v44 float32 = (v30 / 1.5)
                    _ = v44
                }
                _ = v41
            }
            fmt.Println(v6)
            _ = fi37
            _ = v38
        }
        if true {
            v45 := []int{263, 32768, -1, 822}
            _ = v45
        }
    }
    var v46 uint32 = 13
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
