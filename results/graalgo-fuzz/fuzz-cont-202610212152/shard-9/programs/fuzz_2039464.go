package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int32; F1 int16; F2 int16 }
func genFn0(p0 int16) int {
    p0 = (p0 + (p0 % ((p0) | 1)))
    var v5 int = 9223372036854775807
    _ = p0
    _ = v5
    return (v5 + 127)
}
func genFn1(p0 uint32, p1 uint64) int {
    switch "a" {
    case "hello":
        v6 := []int{128, -448}
        _ = v6
    case "":
        var v7 int8 = 67
        _ = v7
    }
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return 32767
}

func main() {
    var v1 uint16 = 772
    var v2 int32 = 0
    v3 := make(map[uint32]string)
    var v4 int = 4294967295
    var v8 float64 = -2.25
    v1 += v1
    if true {
        if false {
            var v9 string = (("" + "") + ("\n\t" + "\n\t"))
            for fi10 := 0; fi10 < 17; fi10++ {
                var v11 float32 = (0.1 - (0.0 / 1.0))
                switch -9223372036854775808 {
                case 42:
                    for fi12 := 0; fi12 < 11; fi12++ {
                        var v13 uint8 = 127
                        v14 := GenS1{F0: "\n\t", F1: uint(v13)}
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                case 407:
                    var v15 uint8 = 2
                    _ = v15
                default:
                    switch 748 {
                    case 0:
                        var v16 uint32 = 2147483647
                        _ = v16
                    case 32767:
                        var v17 int8 = 127
                        _ = v17
                    default:
                        var v18 int8 = 127
                        _ = v18
                    }
                }
                var v19 uint = 915
                _ = fi10
                _ = v11
                _ = v19
            }
            _ = v9
        } else {
            v4 = (v4 & v4)
        }
    }
    v2 += 80
    fmt.Println(v2, v2)
    switch v4 {
    case 0:
        v20 := GenS2{F0: (^(v2)), F1: 42}
        _ = v20
    case -161:
        var v21 uint16 = v1
        _ = v21
    default:
        v22 := GenS3{F0: v2, F1: -1, F2: -32768}
        _ = v22
    }
    for fi23 := 0; fi23 < 14; fi23++ {
        for fi24 := 0; fi24 < 5; fi24++ {
            var v25 uint8 = 127
            _ = fi24
            _ = v25
        }
        v26 := []int{}
        var v27 float32 = (1.5 + (-2.25 * 0.0))
        _ = fi23
        _ = v26
        _ = v27
    }
    var v28 int16 = 32767
    for fi29 := 0; fi29 < 10; fi29++ {
        v4 -= 32767
        _ = fi29
    }
    v30 := GenS1{F0: ("foo bar" + ""), F1: 431}
    if ((1.5 / (1.5 + (1.5 - -2.25))) > 3.14159) {
        for fi31 := 0; fi31 < 14; fi31++ {
            switch genFn1(2147483647, 607) {
            case 128:
                switch fi31 {
                case 259:
                    for fi32 := 0; fi32 < 12; fi32++ {
                        var v33 int64 = -498
                        var v34 int8 = 87
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                case 127:
                    var v35 uint32 = 419
                    _ = v35
                default:
                    var v36 string = "hello"
                    _ = v36
                }
            case -9223372036854775808:
                var v37 int = fi31
                _ = v37
            case 600:
                if ((fi31 >> ((364) & 63)) == (v4 % ((32768) | 1))) {
                    fmt.Println(fi31, v4)
                }
            default:
                if false {
                    v1 += (v1 << ((4294967295) & 63))
                    v3[2147483647] = ((("a" + "hello") + "hello") + (("" + "foo bar") + "\n\t"))
                } else {
                    if (genFn1(uint32(v1), 18446744073709551615) > (v4 ^ (v4 / (((^(v4))) | 1)))) {
                        var v38 bool = false
                        var v39 string = "\n\t"
                        _ = v38
                        _ = v39
                    }
                    var v40 bool = (98 > 2)
                    _ = v40
                }
            }
            var v41 string = (("a" + (("hello" + "a") + "\n\t")) + "\n\t")
            _ = fi31
            _ = v41
        }
        var v42 int = v4
        _ = v42
    }
    v43 := GenS3{F0: v2, F1: (v28 / ((v28) | 1)), F2: (v28 & v28)}
    switch "" {
    case "hello":
        var v44 bool = (v2 <= v2)
        _ = v44
    case "\n\t":
        _ = len(v3)
    default:
        fmt.Println(v8, v4, v2)
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v43
    _ = v8
}
