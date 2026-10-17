package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 int; F1 int8; F2 uint16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 bool; F1 uint }
type GenS6 struct { F0 uint64; F1 uint64; F2 uint64 }
func genFn0(p0 int, p1 uint8) int32 {
    for fi5 := 0; fi5 < 10; fi5++ {
        v6 := GenS1{F0: 2147483647}
        _ = fi5
        _ = v6
    }
    switch (("hello" + "hello") + "\n\t") {
    case "":
        for fi7 := 0; fi7 < 16; fi7++ {
            p0 = (p0 % ((p0) | 1))
            if (true && (float64((3.14159 * 3.14159)) >= -2.25)) {
                if true {
                    fmt.Println(fi7)
                    v8 := []int{32767, -588, p0}
                    for fi9 := 0; fi9 < 2; fi9++ {
                        p1 = (^(p1))
                        p1 = (p1 | (p1 % (((p1 % ((p1) | 1))) | 1)))
                        _ = fi9
                    }
                    _ = v8
                } else {
                    v10 := GenS2{F0: (p0 ^ fi7), F1: 1}
                    _ = v10
                }
            }
            _ = fi7
        }
    }
    _ = p0
    _ = p1
    return -2147483648
}

func main() {
    v1 := []int{}
    var v2 uint16 = 0
    v3 := []int{42, -9223372036854775808, -9223372036854775808, -1}
    var v4 float32 = -2.25
    v11 := GenS2{F0: 127, F1: -561}
    fmt.Println(v4)
    fmt.Println(v2, v2)
    if ((4294967295 == 2) && (622 > 563)) {
        v12 := GenS2{F0: 128, F1: 42}
        v13 := GenS2{F0: -855, F1: 134}
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        var v15 uint64 = 0
        v16 := GenS3{F0: -533, F1: 127, F2: 101}
        _ = fi14
        _ = v15
        _ = v16
    }
    if true {
        v4 = (-2.25 * 0.1)
        switch "hello" {
        case "":
            for fi17 := 0; fi17 < 9; fi17++ {
                if false {
                    var v18 uint64 = 215
                    _ = v18
                }
                for fi19 := 0; fi19 < 10; fi19++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v20 int16 = 42
                    _ = fi19
                    _ = v20
                }
                _ = fi17
            }
        case "\n\t":
            var v21 int16 = 961
            _ = v21
        default:
            var v22 int16 = 993
            _ = v22
        }
        for fi23 := 0; fi23 < 8; fi23++ {
            if (706 >= 410) {
                if len(v3) > 0 {
                    v3[0] = fi23
                }
                fmt.Println(v2, v4)
                v24 := GenS1{F0: 252}
                _ = v24
            }
            var v25 int32 = 517
            _ = fi23
            _ = v25
        }
    } else {
        var v26 uint8 = 255
        v27 := GenS4{F0: 32767}
        var v28 uint16 = 693
        _ = v26
        _ = v27
        _ = v28
    }
    v29 := GenS5{F0: (127 >= 127), F1: 371}
    v30 := GenS6{F0: 1, F1: 168, F2: 4294967295}
    fmt.Println(v2, v4)
    v2 = (^(v2))
    v2 += (v2 << ((18446744073709551615) & 63))
    fmt.Println(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
