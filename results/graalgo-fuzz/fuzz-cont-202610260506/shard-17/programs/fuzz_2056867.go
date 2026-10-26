package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 string; F2 int32 }
func genFn0() uint16 {
    v5 := []int{32768, 32767, -483, 42}
    var v6 uint16 = 32767
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = v5
    _ = v6
    return 65535
}
func genFn1(p0 uint8) int8 {
    v7 := GenS1{F0: ("hello" + (("a" + "foo bar") + "a"))}
    if (true && (128 < 127)) {
        var v8 int32 = -1
        for fi9 := 0; fi9 < 7; fi9++ {
            for fi10 := 0; fi10 < 16; fi10++ {
                v11 := GenS1{F0: "foo bar"}
                _ = fi10
                _ = v11
            }
            v8 += (v8 % (((^(v8))) | 1))
            fmt.Println(fi9, v8, v8)
            _ = fi9
        }
        _ = v8
    } else {
        var v12 float32 = 1e10
        _ = v12
    }
    _ = p0
    _ = v7
    return 42
}

func main() {
    var v1 uint64 = 218
    var v2 uint64 = (v1 & 31)
    var v3 int16 = -64
    var v4 string = "a"
    var v13 uint = 103
    for fi14 := 0; fi14 < 15; fi14++ {
        v15 := []int{32768, fi14, (fi14 | (fi14 << ((18446744073709551615) & 63)))}
        _ = fi14
        _ = v15
    }
    var v16 bool = (false && true)
    v4 = v4
    fmt.Println(v4, v13, v3)
    v3 = (v3 / (((^(v3))) | 1))
    if ((0.1 + -2.25) == 0.1) {
        if false {
            for fi17 := 0; fi17 < 2; fi17++ {
                v1 += v2
                _ = fi17
            }
        } else {
            var v18 uint = v13
            _ = v18
        }
        v19 := GenS1{F0: ""}
        _ = v19
    }
    switch "hello" {
    case "foo bar":
        var v20 float32 = (3.14159 * 0.0)
        _ = v20
    default:
        v2 = 652
    }
    v21 := GenS2{F0: 0, F1: ("" + (v4 + v4)), F2: -1}
    for fi22 := 0; fi22 < 20; fi22++ {
        for fi23 := 0; fi23 < 11; fi23++ {
            v24 := GenS2{F0: 255, F1: (("a" + (v4 + v4)) + "hello"), F2: -1}
            _ = fi23
            _ = v24
        }
        v25 := []int{fi22, (fi22 / ((fi22) | 1))}
        var v26 uint16 = 32767
        _ = fi22
        _ = v25
        _ = v26
    }
    var v27 uint8 = 2
    var v28 uint64 = v1
    switch 128 {
    case 0:
        v29 := GenS1{F0: (("hello" + v4) + "hello")}
        _ = v29
    case 32768:
        fmt.Println(v28, v3)
    default:
        for fi30 := 0; fi30 < 18; fi30++ {
            fmt.Println(v13, v4, fi30)
            _ = fi30
        }
    }
    for fi31 := 0; fi31 < 17; fi31++ {
        var v32 int64 = -9223372036854775808
        var v33 int16 = v3
        var v34 int16 = v3
        _ = fi31
        _ = v32
        _ = v33
        _ = v34
    }
    if (v16 || (!false)) {
        v35 := []int{}
        _ = v35
    } else {
        v28 = v28
        v36 := GenS2{F0: v27, F1: v4, F2: -815}
        _ = v36
    }
    for fi37 := 0; fi37 < 5; fi37++ {
        var v38 uint64 = (^(v2))
        for fi39 := 0; fi39 < 13; fi39++ {
            var v40 float64 = (1.5 + 1.5)
            v41 := []int{(fi39 + fi39), fi37}
            if len(v41) >= 2 {
                v41 = v41[1:len(v41)]
            }
            _ = fi39
            _ = v40
            _ = v41
        }
        _ = fi37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v21
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
