package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint }

func main() {
    var v1 int64 = 42
    var v2 uint8 = 127
    var v3 float64 = 3.14159
    var v4 string = (("a" + "hello") + ("\n\t" + ("\n\t" + "hello")))
    v5 := make(map[uint32]uint32)
    v2 = 42
    for fi6 := 0; fi6 < 20; fi6++ {
        for fi7 := 0; fi7 < 16; fi7++ {
            var v8 int32 = -734
            _ = fi7
            _ = v8
        }
        if (false && true) {
            for fi9 := 0; fi9 < 17; fi9++ {
                var v10 uint32 = 922
                v3 += v3
                switch v4 {
                case "a":
                    if (false || false) {
                        v1 -= v1
                    } else {
                        v11 := GenS1{F0: -1, F1: (^(v10))}
                        fmt.Println(v1, v4)
                        _ = v11
                    }
                case "hello":
                    v12 := GenS1{F0: 42, F1: (v10 / ((65535) | 1))}
                    v13 := GenS1{F0: -1, F1: (v10 | (v10 >> ((1) & 63)))}
                    v4 = v4
                    v14 := []int{(fi6 << ((2) & 63))}
                    _ = v12
                    _ = v13
                    _ = v14
                case "foo bar":
                    v1 -= v1
                default:
                    v15 := GenS1{F0: 42, F1: (v10 / (((v10 | (v10 - (v10 ^ 512)))) | 1))}
                    _ = v15
                }
                if (false || false) {
                    var v16 string = v4
                    _ = v16
                }
                _ = fi9
                _ = v10
            }
        } else {
            v3 = (1e10 / v3)
        }
        _ = fi6
    }
    v17 := GenS2{F0: -32768, F1: (!(42 <= 1))}
    v18 := GenS1{F0: 931, F1: 65535}
    if false {
        var v19 uint16 = 530
        for fi20 := 0; fi20 < 1; fi20++ {
            if true {
                v21 := []int{fi20, (fi20 >> ((0) & 63)), (fi20 ^ -9223372036854775808), (fi20 << ((255) & 63))}
                _ = v21
            } else {
                var v22 string = v4
                var v23 int64 = (v1 % ((v1) | 1))
                _ = v22
                _ = v23
            }
            switch v4 {
            case "a":
                v1 += v1
            case "":
                v19 = v19
            default:
                v24 := []int{127}
                _ = v24
            }
            _ = fi20
        }
        var v25 uint32 = 4294967295
        _ = v19
        _ = v25
    }
    for fi26 := 0; fi26 < 11; fi26++ {
        var v27 int16 = 32767
        _ = fi26
        _ = v27
    }
    v1 = v1
    if (!(1 < 0)) {
        for fi28 := 0; fi28 < 11; fi28++ {
            v2 -= (^(v2))
            fmt.Println(v1)
            _ = fi28
        }
        v29 := []int{1, -9223372036854775808}
        _ = v29
    }
    v4 = v4
    delete(v5, 2147483648)
    var v30 float64 = 1.5
    v31 := GenS1{F0: 1, F1: 178}
    if false {
        for fi32 := 0; fi32 < 18; fi32++ {
            if false {
                var v33 string = (v4 + "hello")
                var v34 int32 = int32(v1)
                _ = v33
                _ = v34
            } else {
                v35 := GenS1{F0: 2147483647, F1: 0}
                for fi36 := 0; fi36 < 2; fi36++ {
                    var v37 int16 = 1
                    _ = fi36
                    _ = v37
                }
                _ = v35
            }
            v38 := GenS3{F0: 2}
            switch (fi32 ^ (^(fi32))) {
            case 32768:
                fmt.Println(v30, fi32, v30)
            case 0:
                v2 += 0
            }
            _ = fi32
            _ = v38
        }
        v39 := GenS4{F0: 65535}
        _ = v39
    } else {
        fmt.Println(v2, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
