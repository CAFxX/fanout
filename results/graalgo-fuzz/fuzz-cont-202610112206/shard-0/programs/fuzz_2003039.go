package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int64 }
func genFn0() uint {
    v7 := []int{32767}
    switch 32768 {
    case -808:
        for fi8 := 0; fi8 < 20; fi8++ {
            fmt.Println(fi8)
            _ = fi8
        }
    case -2147483648:
        if false {
            var v9 uint64 = 31
            var v10 string = "\n\t"
            _ = v10
            _ = v9
        } else {
            v7 = append(v7, -1)
            for fi11 := 0; fi11 < 10; fi11++ {
                fmt.Println(fi11)
                var v12 int16 = 32767
                _ = fi11
                _ = v12
            }
            switch 32767 {
            case 2147483647:
                for fi13 := 0; fi13 < 3; fi13++ {
                    switch 2147483647 {
                    case 42:
                        var v14 int64 = 943
                        _ = v14
                    case 0:
                        v7 = append(v7, -9223372036854775808)
                    }
                    _ = fi13
                }
            case 9223372036854775807:
                var v15 uint64 = 1
                _ = v15
            case -730:
                var v16 uint8 = 116
                _ = v16
            }
        }
    case -1:
        var v17 float32 = ((1e10 / 0.1) - float32(3.14159))
        _ = v17
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        var v19 float32 = 0.0
        v20 := GenS1{F0: 0, F1: 42}
        if ((60 == 127) || false) {
            for fi21 := 0; fi21 < 16; fi21++ {
                v22 := GenS1{F0: 2, F1: 2147483648}
                _ = len(v7)
                v7 = append(v7, 371)
                _ = fi21
                _ = v22
            }
            if ("" > (("a" + "") + "\n\t")) {
                v19 -= v19
            } else {
                v19 += ((float32(0.1) * ((1.5 / v19) * float32((3.14159 * (float64(-2.25) / 0.1))))) / (((v19 * v19) * float32(1e10)) + 1e10))
                if (!("a" <= "")) {
                    v23 := GenS1{F0: 255, F1: 209}
                    var v24 uint16 = 1
                    _ = v23
                    _ = v24
                }
                if (false || false) {
                    var v25 uint8 = 128
                    _ = v25
                }
            }
        } else {
            v19 -= float32(0.0)
            var v26 uint8 = 255
            _ = v26
        }
        _ = fi18
        _ = v19
        _ = v20
    }
    var v27 int64 = 0
    _ = v27
    _ = v7
    return 4294967295
}

func main() {
    var v1 int8 = 42
    var v2 uint64 = 42
    v3 := []int{2147483647, 42, 970}
    var v4 int64 = -393
    v5 := []int{}
    var v6 float32 = 0.1
    var v28 uint16 = 2
    switch "foo bar" {
    case "\n\t":
        switch "a" {
        case "foo bar":
            switch "\n\t" {
            case "a":
                _ = len(v3)
            }
        case "":
            if (1 > 42) {
                for fi29 := 0; fi29 < 16; fi29++ {
                    var v30 int32 = 1
                    v31 := GenS2{F0: (v30 ^ (v30 << ((genFn0()) & 63))), F1: 42}
                    v32 := GenS1{F0: 2, F1: 0}
                    _ = fi29
                    _ = v30
                    _ = v31
                    _ = v32
                }
            } else {
                v2 = v2
                var v33 float64 = 1.5
                _ = v33
            }
        default:
            v1 += v1
        }
    case "":
        fmt.Println(v1)
    case "foo bar":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        v2 -= 2
    }
    var v34 int16 = 136
    v2 -= v2
    v35 := GenS3{F0: v4}
    v2 += 250
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v5 = append(v5, 809)
    for fi36 := 0; fi36 < 6; fi36++ {
        for fi37 := 0; fi37 < 20; fi37++ {
            if (((^(v2)) != 4294967295) || true) {
                if len(v5) > 0 {
                    v5[0] = (^(fi36))
                }
                if (!true) {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                } else {
                    if false {
                        _ = len(v3)
                        v34 -= (v34 + (v34 % ((v34) | 1)))
                    }
                    v5 = append(v5, (fi36 | (fi37 % (((fi37 + fi37)) | 1))))
                }
                var v38 int = (fi36 * 2147483647)
                _ = v38
            } else {
                fmt.Println(v4, v4, v4)
                var v39 int64 = v4
                _ = v39
            }
            _ = fi37
        }
        if false {
            switch int(v28) {
            case -2147483648:
                v28 += (v28 % (((v28 & 505)) | 1))
            }
            switch "\n\t" {
            case "":
                _ = len(v3)
            case "hello":
                var v40 int16 = int16(v28)
                _ = v40
            default:
                if (false || false) {
                    var v41 float64 = -2.25
                    switch "a" {
                    case "hello":
                        fmt.Println(v6)
                    }
                    v5 = append(v5, (-(fi36)))
                    _ = v41
                }
            }
        }
        _ = fi36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
