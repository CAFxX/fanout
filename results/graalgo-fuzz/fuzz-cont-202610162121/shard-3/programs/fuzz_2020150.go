package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int32; F1 uint; F2 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64; F1 int32; F2 int16 }
func genFn0() bool {
    var v7 uint8 = 128
    switch -839 {
    case 1:
        switch (("foo bar" + "") + ("foo bar" + ("\n\t" + "hello"))) {
        case "hello":
            switch "\n\t" {
            case "a":
                switch -1 {
                case -28:
                    var v8 int32 = -2147483648
                    _ = v8
                default:
                    v9 := []int{-9223372036854775808, -516}
                    _ = v9
                }
            default:
                v10 := []int{32767, -801, 9223372036854775807, 2147483647}
                _ = v10
            }
        case "":
            v11 := []int{490, 9223372036854775807}
            _ = v11
        case "\n\t":
            var v12 int = 1
            _ = v12
        default:
            for fi13 := 0; fi13 < 17; fi13++ {
                for fi14 := 0; fi14 < 11; fi14++ {
                    fmt.Println(fi14)
                    v15 := GenS1{F0: 1, F1: v7, F2: 65535}
                    fmt.Println(fi13)
                    _ = fi14
                    _ = v15
                }
                switch "\n\t" {
                case "\n\t":
                    v16 := []int{fi13, 911}
                    _ = v16
                default:
                    for fi17 := 0; fi17 < 14; fi17++ {
                        fmt.Println(fi13, fi17)
                        _ = fi17
                    }
                }
                _ = fi13
            }
        }
    default:
        v18 := []int{42, 886, 32768}
        _ = v18
    }
    var v19 int32 = -2147483648
    fmt.Println(v19, v19)
    _ = v19
    _ = v7
    return false
}
func genFn1(p0 int32) uint32 {
    var v20 uint = 414
    v21 := GenS2{F0: -623, F1: v20, F2: 255}
    _ = p0
    _ = v20
    _ = v21
    return 70
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]bool)
    var v3 int = -136
    v4 := []string{}
    var v5 uint8 = 66
    v6 := []string{"", ("foo bar" + "hello"), "\n\t", "hello"}
    var v22 uint8 = (v5 >> ((11) & 63))
    _ = len(v2)
    if true {
        if (42 > -79) {
            v5 -= (v22 / (((v5 & 1)) | 1))
        } else {
            v22 += v5
            fmt.Println(v22, v3, v3)
            var v23 int16 = 1
            _ = v23
        }
        var v24 bool = false
        var v25 bool = true
        _ = v24
        _ = v25
    }
    if (((255 != 4294967295) && false) || false) {
        fmt.Println(v3, v3)
    }
    fmt.Println(v5, v5, v5)
    v5 = (v22 / ((v22) | 1))
    for fi26 := 0; fi26 < 6; fi26++ {
        for fi27 := 0; fi27 < 12; fi27++ {
            v28 := GenS3{F0: (v5 << ((1) & 63))}
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    var v29 uint16 = 698
    var v30 uint32 = 2147483648
    fmt.Println(v29, v22, v3)
    v22 -= (v5 % (((v22 / (((^(v5))) | 1))) | 1))
    if false {
        var v31 uint64 = 675
        _ = v31
    } else {
        v32 := GenS4{F0: 0, F1: 2147483647, F2: -1}
        _ = v32
    }
    v3 = v3
    if true {
        fmt.Println(v29)
    } else {
        v30 = (v30 & genFn1(0))
        var v33 int8 = -1
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
