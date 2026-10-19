package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 int8; F1 uint8; F2 uint16 }
func genFn0(p0 uint, p1 int8) uint16 {
    switch "\n\t" {
    case "\n\t":
        var v4 bool = (!(1 >= -1))
        _ = v4
    case "a":
        if true {
            var v5 int64 = 9223372036854775807
            v5 += v5
            if false {
                v6 := []int{}
                _ = v6
            }
            _ = v5
        } else {
            p0 = p0
            switch "foo bar" {
            case "foo bar":
                if false {
                    var v7 bool = (2 <= 600)
                    if v7 {
                        for fi8 := 0; fi8 < 3; fi8++ {
                            p0 = (p0 & p0)
                            v9 := []int{fi8, (fi8 - -779), -925}
                            v10 := GenS1{F0: (fi8 << ((p0) & 63))}
                            _ = fi8
                            _ = v10
                            _ = v9
                        }
                    } else {
                        var v11 int = 286
                        var v12 bool = (2 >= 65535)
                        _ = v11
                        _ = v12
                    }
                    _ = v7
                }
            }
            v13 := []int{-434}
            _ = v13
        }
    }
    var v14 uint16 = 42
    for fi15 := 0; fi15 < 8; fi15++ {
        v16 := GenS1{F0: (^(fi15))}
        for fi17 := 0; fi17 < 11; fi17++ {
            v18 := GenS1{F0: (fi15 | (fi15 ^ fi17))}
            v19 := []int{(fi15 % ((fi15) | 1)), fi15}
            var v20 uint = (^(p0))
            v21 := GenS1{F0: (fi17 | -1)}
            fmt.Println(v20)
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
            _ = v21
        }
        _ = fi15
        _ = v16
    }
    if (!(1 > 32767)) {
        switch "a" {
        case "foo bar":
            v22 := GenS1{F0: 42}
            _ = v22
        case "":
            var v23 int = 803
            _ = v23
        case "\n\t":
            if (42 > 198) {
                if true {
                    v24 := []int{}
                    if ((v14 % ((v14) | 1)) < 1) {
                        var v25 int16 = 461
                        var v26 bool = (3.14159 > (float32(0.0) - 3.14159))
                        if len(v24) > 0 {
                            v24[0] = -944
                        }
                        _ = v25
                        _ = v26
                    }
                    fmt.Println(p1, p0, p1)
                    _ = v24
                } else {
                    var v27 uint64 = 4294967295
                    _ = v27
                }
                p0 -= p0
            }
        default:
            v28 := []int{759, -2147483648, 2147483647, 9223372036854775807}
            _ = v28
        }
        v29 := []int{-745, 686, 1, 9223372036854775807}
        _ = v29
    }
    _ = p0
    _ = p1
    _ = v14
    return v14
}
func genFn1() bool {
    var v30 uint64 = 4294967295
    v31 := []int{-780, 127}
    fmt.Println(v30)
    fmt.Println(v30)
    if true {
        fmt.Println(v30)
    } else {
        v30 = v30
        v32 := GenS2{F0: 2, F1: 1}
        var v33 string = "\n\t"
        _ = v32
        _ = v33
    }
    _ = v30
    _ = v31
    return (1e10 == -2.25)
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 0
    var v3 int8 = 42
    v3 -= v3
    if false {
        var v34 int = 32768
        delete(v1, (("a" + (("" + "\n\t") + "foo bar")) + "foo bar"))
        _ = v34
    } else {
        var v35 uint16 = 255
        if (true && ((-1 < -2147483648) || (v3 == (v3 / ((v3) | 1))))) {
            v35 -= 234
        }
        _ = v35
    }
    if (!(1 == 17)) {
        v2 = v2
    }
    var v36 uint64 = 18446744073709551615
    if (-776 != 42) {
        switch -9223372036854775808 {
        case 2147483647:
            var v37 int32 = 2147483647
            _ = v37
        case 32767:
            v2 += v2
        default:
            v3 += (v3 >> ((v2) & 63))
        }
        var v38 uint8 = 128
        v39 := []int{-2147483648, 499, 32768, 524}
        _ = v38
        _ = v39
    }
    switch "hello" {
    case "hello":
        v3 = 93
    case "foo bar":
        v3 -= (v3 | 127)
    case "":
        v40 := []int{}
        _ = v40
    default:
        v41 := GenS1{F0: 9223372036854775807}
        _ = v41
    }
    var v42 uint64 = (v36 ^ (v36 | v36))
    v43 := GenS3{F0: (v3 - (^(v3))), F1: 127, F2: genFn0((v2 * v2), int8(v2))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v42
    _ = v43
}
