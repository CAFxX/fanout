package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8; F1 uint32 }
type GenS5 struct { F0 int32; F1 uint16 }
func genFn0(p0 string, p1 string, p2 uint16) uint64 {
    if true {
        fmt.Println(p1, p1, p2)
        v4 := GenS1{F0: 2147483648, F1: 130, F2: 42}
        _ = v4
    } else {
        fmt.Println(p0, p1)
        if ((-2.25 < (1e10 + 1e10)) || ((1e10 + 3.14159) == 1e10)) {
            p0 = p1
            switch -322 {
            case 32768:
                for fi5 := 0; fi5 < 1; fi5++ {
                    v6 := []int{(fi5 >> ((42) & 63)), (fi5 >> ((2) & 63)), (fi5 + (fi5 % (((-(fi5))) | 1)))}
                    _ = fi5
                    _ = v6
                }
            }
            if true {
                v7 := GenS2{F0: -699}
                for fi8 := 0; fi8 < 5; fi8++ {
                    var v9 int64 = 654
                    var v10 float32 = 1.5
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                var v11 int = 0
                _ = v11
                _ = v7
            }
        } else {
            var v12 string = p1
            _ = v12
        }
    }
    v13 := []int{32768, -57, 4294967295, 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 4294967295
}

func main() {
    var v1 int16 = 0
    var v2 string = "\n\t"
    var v3 int8 = 1
    var v14 string = (v2 + (("a" + (v2 + ((("\n\t" + "hello") + "hello") + v2))) + v2))
    v1 = v1
    switch 2147483647 {
    case 389:
        if false {
            var v15 int64 = -9223372036854775808
            v2 = v2
            fmt.Println(v14, v2, v3)
            _ = v15
        } else {
            v3 += 42
            v16 := GenS3{F0: v2}
            _ = v16
        }
    case 0:
        for fi17 := 0; fi17 < 18; fi17++ {
            switch v14 {
            case "":
                v14 = v2
            default:
                v18 := GenS4{F0: (v3 ^ 127), F1: 0}
                _ = v18
            }
            _ = fi17
        }
    }
    v19 := GenS2{F0: -32768}
    for fi20 := 0; fi20 < 7; fi20++ {
        for fi21 := 0; fi21 < 15; fi21++ {
            v22 := GenS5{F0: -1, F1: 0}
            v1 = v1
            v2 = v14
            _ = fi21
            _ = v22
        }
        var v23 uint64 = genFn0("a", v14, 0)
        v23 -= (v23 << ((68) & 63))
        _ = fi20
        _ = v23
    }
    var v24 int = 1
    fmt.Println(v14, v14)
    fmt.Println(v1, v1, v2)
    v25 := GenS3{F0: (v14 + v14)}
    for fi26 := 0; fi26 < 13; fi26++ {
        if false {
            var v27 int16 = (v1 / (((v1 - v1)) | 1))
            _ = v27
        }
        var v28 int16 = (-(v1))
        var v29 int8 = (v3 - (v3 ^ (^(v3))))
        _ = fi26
        _ = v28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v3
}
