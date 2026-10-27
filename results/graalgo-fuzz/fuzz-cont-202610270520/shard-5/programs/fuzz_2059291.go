package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 int; F2 bool }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 uint32 }
func genFn0() int32 {
    switch -414 {
    case -572:
        var v5 bool = (43 == 38)
        _ = v5
    case 572:
        v6 := GenS1{F0: 1}
        _ = v6
    case -990:
        v7 := GenS2{F0: 2, F1: 9223372036854775807, F2: (false && true)}
        _ = v7
    default:
        for fi8 := 0; fi8 < 14; fi8++ {
            v9 := []int{fi8, 127, fi8}
            var v10 bool = true
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    var v11 bool = false
    _ = v11
    return 2147483647
}
func genFn1() uint8 {
    var v12 bool = ((!(1e10 > 1e10)) && false)
    for fi13 := 0; fi13 < 4; fi13++ {
        fmt.Println(fi13, v12)
        _ = fi13
    }
    v12 = (!(3.14159 == 0.1))
    var v14 int16 = 866
    v14 = (v14 % (((v14 - (v14 >> ((534) & 63)))) | 1))
    _ = v12
    _ = v14
    return 2
}

func main() {
    v1 := []int{-281, 42}
    var v2 float32 = float32(1e10)
    var v3 uint8 = 128
    v4 := []string{((("foo bar" + "hello") + "hello") + ((("foo bar" + ("a" + "foo bar")) + (("foo bar" + "hello") + ("a" + ("hello" + "")))) + ("foo bar" + "hello"))), "a", "", ""}
    if ("hello" > ("a" + ("a" + "a"))) {
        if false {
            if false {
                for fi15 := 0; fi15 < 5; fi15++ {
                    var v16 uint64 = 472
                    for fi17 := 0; fi17 < 15; fi17++ {
                        var v18 uint16 = 255
                        v19 := GenS2{F0: 42, F1: -425, F2: true}
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    _ = fi15
                    _ = v16
                }
                v4 = append(v4, ((("" + "\n\t") + ("a" + "hello")) + "a"))
                v4 = append(v4, "foo bar")
            }
        } else {
            var v20 int = 1
            var v21 uint32 = 660
            v22 := GenS1{F0: 146}
            _ = v20
            _ = v21
            _ = v22
        }
    } else {
        v2 += 1.5
        if (221 < 1) {
            v3 += (v3 / (((v3 >> ((255) & 63))) | 1))
            switch -594 {
            case -14:
                if (!true) {
                    var v23 int16 = 837
                    var v24 int16 = (v23 | 1000)
                    _ = v23
                    _ = v24
                } else {
                    fmt.Println(v3, v2)
                    v3 -= (^(v3))
                }
            default:
                v3 = (v3 / ((v3) | 1))
            }
        } else {
            var v25 uint64 = 35
            for fi26 := 0; fi26 < 16; fi26++ {
                v27 := GenS3{F0: int64(fi26), F1: (v3 / (((v3 & (v3 - (v3 << ((349) & 63))))) | 1))}
                v28 := GenS4{F0: 2147483648}
                _ = fi26
                _ = v27
                _ = v28
            }
            fmt.Println(v3)
            _ = v25
        }
        if (v2 != v2) {
            if len(v4) > 0 {
                v4[0] = ""
            }
        } else {
            _ = len(v4)
        }
    }
    v1 = append(v1, 1)
    v29 := GenS4{F0: 0}
    v4 = append(v4, ("\n\t" + "hello"))
    v3 += v3
    v30 := GenS1{F0: 351}
    _ = len(v4)
    v31 := GenS3{F0: 213, F1: v3}
    var v32 uint8 = v3
    var v33 string = ""
    for fi34 := 0; fi34 < 3; fi34++ {
        if false {
            for fi35 := 0; fi35 < 3; fi35++ {
                v3 = (v32 / (((v32 / (((v3 - (v32 / ((v3) | 1)))) | 1))) | 1))
                if true {
                    v36 := GenS1{F0: 18446744073709551615}
                    var v37 int16 = -32768
                    v32 += v3
                    _ = v36
                    _ = v37
                }
                var v38 int64 = -9223372036854775808
                _ = fi35
                _ = v38
            }
            fmt.Println(v3)
        }
        _ = fi34
    }
    v39 := GenS1{F0: 546}
    var v40 uint8 = (^(v32))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v39
    _ = v4
    _ = v40
}
