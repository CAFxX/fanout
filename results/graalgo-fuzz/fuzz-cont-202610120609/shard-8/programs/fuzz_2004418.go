package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 int16; F1 int16; F2 int16 }
type GenS4 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS5 struct { F0 string; F1 int16; F2 bool }
func genFn0(p0 int, p1 uint64, p2 uint) uint32 {
    var v4 int16 = -1
    v5 := GenS1{F0: p2}
    v6 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return uint32(p1)
}
func genFn1() int16 {
    v7 := GenS1{F0: 655}
    if (true || (!(3.14159 > (1e10 + 1.5)))) {
        v8 := GenS1{F0: 4294967295}
        switch 105 {
        case 127:
            v9 := GenS2{F0: -9223372036854775808, F1: "a"}
            _ = v9
        default:
            if ("a" >= (("\n\t" + "foo bar") + ("foo bar" + ""))) {
                var v10 uint16 = 686
                _ = v10
            } else {
                var v11 uint8 = 2
                v12 := []int{1}
                v13 := GenS3{F0: 32767, F1: 1, F2: -1}
                _ = v11
                _ = v12
                _ = v13
            }
        }
        _ = v8
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        var v15 string = "\n\t"
        _ = fi14
        _ = v15
    }
    _ = v7
    return -32768
}

func main() {
    var v1 int32 = 225
    var v2 string = "a"
    var v3 int64 = 9223372036854775807
    v16 := GenS2{F0: 1, F1: v2}
    for fi17 := 0; fi17 < 10; fi17++ {
        if (true || ("\n\t" <= (v2 + v2))) {
            fmt.Println(v1)
            fmt.Println(v1, v2, v1)
        }
        _ = fi17
    }
    v18 := GenS4{F0: 42, F1: 2, F2: 2}
    fmt.Println(v1, v1)
    fmt.Println(v1)
    if true {
        for fi19 := 0; fi19 < 17; fi19++ {
            fmt.Println(v1)
            v1 = (v1 >> ((997) & 63))
            switch "foo bar" {
            case "\n\t":
                var v20 float64 = 3.14159
                _ = v20
            }
            _ = fi19
        }
    }
    v1 = v1
    v21 := GenS1{F0: 18446744073709551615}
    switch 4294967295 {
    case 766:
        v22 := []int{32767, 127, 42, 995}
        _ = v22
    case 2147483647:
        if (!((v3 % (((v3 - (v3 & (v3 & -9223372036854775808)))) | 1)) < 42)) {
            switch "\n\t" {
            case "a":
                v23 := GenS5{F0: v2, F1: genFn1(), F2: false}
                _ = v23
            }
            var v24 int64 = (v3 >> ((65535) & 63))
            _ = v24
        }
    default:
        fmt.Println(v1, v1)
    }
    v1 = (v1 | (v1 + v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v21
    _ = v3
}
