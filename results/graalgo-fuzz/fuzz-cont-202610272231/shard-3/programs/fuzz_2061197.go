package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint64; F1 uint8 }
func genFn0(p0 int64, p1 int8) int64 {
    if false {
        var v4 uint64 = 4294967295
        if true {
            v5 := []int{-9223372036854775808}
            switch int(v4) {
            case -132:
                v6 := GenS1{F0: 42, F1: 498}
                _ = v6
            default:
                fmt.Println(p1, p0, p1)
            }
            _ = v5
        }
        v4 -= (^(v4))
        _ = v4
    }
    p1 = (^(p1))
    _ = p0
    _ = p1
    return 20
}

func main() {
    var v1 uint = 676
    v2 := make(map[string]bool)
    var v3 bool = ((42 < 340) && false)
    v1 += v1
    v1 += 437
    var v7 float64 = -2.25
    v3 = v3
    switch "" {
    case "":
        for fi8 := 0; fi8 < 5; fi8++ {
            fmt.Println(fi8, v3, fi8)
            v9 := GenS2{F0: 4294967295, F1: 11}
            if v3 {
                for fi10 := 0; fi10 < 13; fi10++ {
                    var v11 int8 = 100
                    var v12 uint32 = 244
                    v12 = (v12 / (((v12 - v12)) | 1))
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                v13 := GenS1{F0: -713, F1: 2}
                _ = v13
            }
            _ = fi8
            _ = v9
        }
    default:
        if false {
            if (82 != 0) {
                var v14 int16 = 32767
                _ = v14
            } else {
                v1 -= v1
            }
        }
    }
    var v15 uint16 = 65535
    v1 += 1
    fmt.Println(v1, v3, v15)
    v1 += (v1 ^ 768)
    var v16 uint8 = 127
    for fi17 := 0; fi17 < 2; fi17++ {
        v18 := []int{32767, (fi17 % ((fi17) | 1)), 0}
        _ = fi17
        _ = v18
    }
    v1 -= 4294967295
    if true {
        for fi19 := 0; fi19 < 3; fi19++ {
            for fi20 := 0; fi20 < 18; fi20++ {
                var v21 bool = ((v16 << ((v1) & 63)) == 189)
                v3 = (42 > 2147483647)
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
        v22 := GenS2{F0: 0, F1: v16}
        for fi23 := 0; fi23 < 16; fi23++ {
            switch (fi23 << ((v1) & 63)) {
            case -2147483648:
                v24 := []int{fi23, fi23, 149}
                _ = v24
            default:
                switch "hello" {
                case "\n\t":
                    var v25 float32 = float32(v7)
                    _ = v25
                }
            }
            var v26 int32 = -2147483648
            for fi27 := 0; fi27 < 4; fi27++ {
                var v28 uint16 = v15
                _ = fi27
                _ = v28
            }
            _ = fi23
            _ = v26
        }
        _ = v22
    }
    v29 := v2[(((("\n\t" + ("\n\t" + (("" + (("hello" + "foo bar") + "")) + "foo bar"))) + "") + ("" + "\n\t")) + ("\n\t" + "foo bar"))]
    switch 1 {
    case 2147483647:
        v29 = (false && (v7 > 3.14159))
    case -298:
        v16 -= (v16 | 255)
    case 299:
        v15 = (v15 * v15)
    }
    v2["\n\t"] = v3
    v30 := GenS2{F0: uint64(v15), F1: v16}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v7
}
