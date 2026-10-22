package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint32 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 bool; F1 uint32; F2 int64 }
type GenS4 struct { F0 string; F1 uint8 }
func genFn0() uint {
    var v6 uint = 516
    var v7 int32 = -2147483648
    _ = v6
    _ = v7
    return v6
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 83
    v3 := []string{("\n\t" + (("hello" + ("" + ("a" + ""))) + "hello")), ("" + "\n\t"), ("\n\t" + ("hello" + ("\n\t" + "hello")))}
    v4 := []int{32768}
    var v5 uint32 = 65535
    var v8 int16 = -32768
    var v9 int8 = 42
    for fi10 := 0; fi10 < 10; fi10++ {
        switch "foo bar" {
        case "hello":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "foo bar":
            v11 := GenS1{F0: genFn0(), F1: 642, F2: (v5 + 2)}
            _ = v11
        default:
            v3 = append(v3, "a")
        }
        v5 += (v5 / (((v5 & v5)) | 1))
        _ = fi10
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        v13 := GenS1{F0: 0, F1: 2147483647, F2: v5}
        fmt.Println(v9, fi12)
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        var v15 uint = 1
        v16 := GenS2{F0: true, F1: 9223372036854775807}
        switch (-(fi14)) {
        case 42:
            if true {
                v1[fi14] = true
                v3 = append(v3, "a")
            } else {
                for fi17 := 0; fi17 < 12; fi17++ {
                    v5 -= (v5 - 2147483648)
                    v18 := GenS3{F0: (9223372036854775807 == -1), F1: 0, F2: -54}
                    _ = fi17
                    _ = v18
                }
                var v19 uint16 = 1
                _ = v19
            }
        case 4294967295:
            v15 = (v15 * genFn0())
        default:
            switch "hello" {
            case "foo bar":
                var v20 int64 = 0
                _ = v20
            default:
                v8 -= 855
            }
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    v8 -= v8
    v5 = (v5 | 30)
    v2 = (v2 % ((v2) | 1))
    switch -217 {
    case 225:
        var v21 int16 = v8
        _ = v21
    case 32767:
        for fi22 := 0; fi22 < 4; fi22++ {
            var v23 uint8 = 1
            _ = fi22
            _ = v23
        }
    case -1:
        for fi24 := 0; fi24 < 2; fi24++ {
            for fi25 := 0; fi25 < 12; fi25++ {
                v9 -= (v9 ^ v9)
                _ = fi25
            }
            v9 -= (v9 & v9)
            _ = fi24
        }
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        v5 -= (v5 / (((^(v5))) | 1))
        _ = fi26
    }
    var v27 int8 = v9
    var v28 uint8 = 121
    v5 += (v5 >> ((18446744073709551615) & 63))
    v28 = v28
    v29 := GenS1{F0: 144, F1: 658, F2: v5}
    v30 := GenS4{F0: "", F1: 128}
    if (false || false) {
        for fi31 := 0; fi31 < 7; fi31++ {
            var v32 float32 = ((1.5 + 3.14159) * float32((0.1 + -2.25)))
            v9 -= int8(v5)
            _ = fi31
            _ = v32
        }
        var v33 int64 = 42
        switch 374 {
        case -9223372036854775808:
            v27 += v9
        case 32767:
            switch "foo bar" {
            case "":
                var v34 uint16 = 65535
                _ = v34
            }
        default:
            v1[42] = true
        }
        _ = v33
    } else {
        _ = len(v4)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
