package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 string }
type GenS2 struct { F0 string; F1 string; F2 uint8 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int64, p1 uint64) int16 {
    if (true || ("\n\t" != ("a" + "\n\t"))) {
        switch 127 {
        case -1:
            v6 := []int{0}
            _ = v6
        case 954:
            p1 -= 213
        }
    } else {
        for fi7 := 0; fi7 < 8; fi7++ {
            p1 = (p1 & p1)
            _ = fi7
        }
    }
    v8 := []int{830}
    p0 -= -147
    var v9 int = 679
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 32767
}
func genFn1() int8 {
    switch 4294967295 {
    case 9223372036854775807:
        var v10 int = -202
        _ = v10
    case 12:
        v11 := []int{27}
        _ = v11
    }
    var v12 int8 = -128
    _ = v12
    return v12
}

func main() {
    var v1 int64 = 0
    v2 := []int{0, -638, -1, 1}
    var v3 bool = true
    v4 := []int{32767, -9223372036854775808, 572}
    var v5 uint32 = 42
    v3 = v3
    fmt.Println(v5, v3)
    fmt.Println(v1, v3, v1)
    if (!(1 >= 2)) {
        if false {
            v1 = (v1 << ((745) & 63))
            if (189 < 2) {
                var v13 int32 = 2147483647
                if false {
                    fmt.Println(v3, v3)
                    v3 = v3
                    var v14 uint = 1
                    _ = v14
                } else {
                    v2 = append(v2, -1)
                    fmt.Println(v1)
                    v15 := GenS1{F0: (v5 << ((255) & 63)), F1: 65535, F2: "foo bar"}
                    _ = v15
                }
                _ = v13
            }
        }
    }
    fmt.Println(v5)
    if true {
        var v16 float64 = (0.0 / 3.14159)
        var v17 int8 = 127
        _ = v16
        _ = v17
    }
    fmt.Println(v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v18 := GenS2{F0: "a", F1: "hello", F2: 42}
    if true {
        v3 = ((42 > uint16(v5)) && v3)
        v1 = (v1 | (v1 ^ v1))
        var v19 string = "\n\t"
        _ = v19
    }
    _ = len(v4)
    v1 = (-(v1))
    switch 810 {
    case 4294967295:
        for fi20 := 0; fi20 < 11; fi20++ {
            switch 32768 {
            case 2147483647:
                for fi21 := 0; fi21 < 8; fi21++ {
                    v4 = append(v4, fi20)
                    var v22 float32 = 3.14159
                    var v23 int64 = 9223372036854775807
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            case -527:
                v24 := GenS3{F0: (v5 | 0)}
                _ = v24
            default:
                var v25 int = 4294967295
                _ = v25
            }
            v1 += v1
            _ = fi20
        }
    case -9223372036854775808:
        var v26 int = -9223372036854775808
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
