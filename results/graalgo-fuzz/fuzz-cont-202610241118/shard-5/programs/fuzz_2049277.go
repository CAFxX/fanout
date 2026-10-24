package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int32 }
type GenS2 struct { F0 int64; F1 uint; F2 int16 }
type GenS3 struct { F0 uint32; F1 int; F2 uint64 }
func genFn0(p0 string) int8 {
    var v4 uint8 = 122
    if false {
        switch -9223372036854775808 {
        case 32767:
            for fi5 := 0; fi5 < 12; fi5++ {
                v4 = (v4 >> ((65535) & 63))
                fmt.Println(fi5)
                _ = fi5
            }
        case 2147483647:
            switch "a" {
            case "hello":
                p0 = (p0 + p0)
            case "":
                p0 = ("" + p0)
            }
        case 637:
            switch 9223372036854775807 {
            case 32767:
                v6 := []int{}
                _ = v6
            }
        }
        v7 := []int{-9223372036854775808, 128, -815}
        v4 += 0
        _ = v7
    }
    v4 += 0
    _ = p0
    _ = v4
    return -1
}
func genFn1() uint {
    var v8 int8 = 107
    var v9 float64 = 0.0
    for fi10 := 0; fi10 < 16; fi10++ {
        fmt.Println(v9, v9)
        _ = fi10
    }
    v9 -= float64(0.0)
    var v11 uint64 = 42
    _ = v11
    _ = v8
    _ = v9
    return 137
}

func main() {
    var v1 uint64 = 42
    var v2 int32 = -353
    v3 := make(map[int]string)
    if (false || (9223372036854775807 == -711)) {
        v1 += (v1 * 18446744073709551615)
    }
    var v12 uint32 = 2147483648
    v13 := v3[4294967295]
    if ((!true) && true) {
        fmt.Println(v13)
        if true {
            for fi14 := 0; fi14 < 14; fi14++ {
                delete(v3, 9223372036854775807)
                switch fi14 {
                case -195:
                    v12 = (v12 ^ v12)
                case -935:
                    for fi15 := 0; fi15 < 3; fi15++ {
                        v2 += (^(v2))
                        v3[fi14] = "a"
                        v16 := GenS1{F0: fi15, F1: (v13 + ((v13 + v13) + v13)), F2: v2}
                        _ = fi15
                        _ = v16
                    }
                default:
                    if (!(9223372036854775807 != 9223372036854775807)) {
                        var v17 float32 = 0.0
                        v2 = (v2 ^ 42)
                        _ = v17
                    }
                }
                _ = fi14
            }
            v18 := GenS1{F0: 1, F1: ("" + (v13 + ((((v13 + "\n\t") + "a") + (("\n\t" + (v13 + ("\n\t" + (("" + v13) + (v13 + "\n\t"))))) + v13)) + (v13 + (v13 + v13))))), F2: 1}
            v13 = v13
            _ = v18
        }
        for fi19 := 0; fi19 < 12; fi19++ {
            if ((v1 + 18446744073709551615) != v1) {
                v20 := GenS1{F0: fi19, F1: (v13 + "foo bar"), F2: v2}
                _ = v20
            }
            _ = fi19
        }
    }
    for fi21 := 0; fi21 < 16; fi21++ {
        for fi22 := 0; fi22 < 14; fi22++ {
            var v23 int8 = 127
            v23 = -128
            v24 := GenS2{F0: -824, F1: genFn1(), F2: -824}
            var v25 bool = (!(v13 != v13))
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
        }
        v1 -= 729
        if (((!true) && false) || false) {
            v1 -= 18446744073709551615
            v1 -= 1
            var v26 string = v13
            _ = v26
        } else {
            v13 = (v13 + (((v13 + "\n\t") + ("\n\t" + "")) + v13))
            v27 := GenS3{F0: (v12 ^ v12), F1: fi21, F2: (v1 | v1)}
            fmt.Println(v1, v12)
            _ = v27
        }
        _ = fi21
    }
    var v28 float32 = (0.0 * 1e10)
    delete(v3, 0)
    if (!(!true)) {
        fmt.Println(v28)
        v2 = (v2 * (v2 % (((-(v2))) | 1)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v28
    _ = v3
}
