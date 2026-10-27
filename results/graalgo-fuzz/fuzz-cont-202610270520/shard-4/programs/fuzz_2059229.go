package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int16; F1 int; F2 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16; F1 uint32 }
type GenS6 struct { F0 bool; F1 string; F2 int8 }
func genFn0(p0 uint16) bool {
    var v4 uint = 42
    if false {
        p0 -= 0
    } else {
        if (231 == 309) {
            var v5 uint16 = p0
            _ = v5
        } else {
            var v6 int32 = -2147483648
            p0 -= p0
            _ = v6
        }
        for fi7 := 0; fi7 < 10; fi7++ {
            v8 := GenS1{F0: (true || (!true)), F1: 234, F2: -439}
            switch fi7 {
            case -9223372036854775808:
                v9 := []int{fi7, fi7}
                _ = v9
            case 42:
                v4 -= v4
            case 0:
                for fi10 := 0; fi10 < 12; fi10++ {
                    v11 := GenS2{F0: 0, F1: -87}
                    _ = fi10
                    _ = v11
                }
            }
            _ = fi7
            _ = v8
        }
    }
    _ = p0
    _ = v4
    return (-464 == -1)
}
func genFn1(p0 uint) string {
    v12 := []int{}
    switch 32767 {
    case -1:
        p0 -= 660
    default:
        v13 := GenS3{F0: -32768, F1: -2147483648, F2: 137}
        _ = v13
    }
    switch "\n\t" {
    case "":
        switch -582 {
        case -540:
            p0 = 4294967295
        case 576:
            v14 := GenS4{F0: 42}
            _ = v14
        case -9223372036854775808:
            var v15 uint8 = 154
            _ = v15
        }
    case "\n\t":
        for fi16 := 0; fi16 < 2; fi16++ {
            v17 := GenS4{F0: 235}
            var v18 uint = (p0 & 2)
            _ = fi16
            _ = v17
            _ = v18
        }
    case "foo bar":
        fmt.Println(p0)
    }
    _ = p0
    _ = v12
    return "\n\t"
}

func main() {
    var v1 int = -9223372036854775808
    v2 := make(map[string]uint32)
    var v3 uint32 = 2147483647
    v1 = 473
    fmt.Println(v1)
    for fi19 := 0; fi19 < 13; fi19++ {
        v1 = fi19
        var v20 string = ("foo bar" + "\n\t")
        switch v1 {
        case 150:
            for fi21 := 0; fi21 < 3; fi21++ {
                v22 := GenS5{F0: 159, F1: 450}
                _ = fi21
                _ = v22
            }
        case 32767:
            v1 = (v1 - int(v3))
        }
        _ = fi19
        _ = v20
    }
    if false {
        fmt.Println(v1, v3)
        v23 := GenS6{F0: (-1 >= 373), F1: "foo bar", F2: -35}
        _ = v23
    } else {
        v3 += v3
        v3 -= (v3 * v3)
        switch v1 {
        case 2147483647:
            var v24 bool = true
            _ = v24
        }
    }
    switch v1 {
    case 2147483647:
        v25 := GenS6{F0: (0 > int32(v3)), F1: "\n\t", F2: int8(v1)}
        _ = v25
    case 0:
        var v26 uint8 = 42
        _ = v26
    default:
        var v27 int32 = 451
        _ = v27
    }
    v3 = 0
    if false {
        fmt.Println(v3, v1)
    }
    var v28 uint32 = (v3 * v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
}
