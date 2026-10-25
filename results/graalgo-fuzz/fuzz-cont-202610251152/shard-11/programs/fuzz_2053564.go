package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS5 struct { F0 uint16; F1 uint8; F2 int8 }
func genFn0(p0 int16) uint32 {
    v4 := GenS1{F0: 32768, F1: -504}
    var v5 uint16 = 1
    v6 := []int{1, 0}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 397
}
func genFn1(p0 uint64, p1 string, p2 bool) int32 {
    v7 := []int{}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    if false {
        if (!true) {
            v8 := GenS1{F0: int(p0), F1: 102}
            var v9 uint64 = p0
            _ = v8
            _ = v9
        }
        if (p0 > (p0 & p0)) {
            var v10 uint16 = 282
            v10 -= 0
            switch -40 {
            case 406:
                v10 += (v10 ^ 1)
            case -1:
                if p2 {
                    var v11 float64 = 1.5
                    _ = v11
                }
            case -9223372036854775808:
                var v12 uint16 = 0
                _ = v12
            }
            _ = v10
        } else {
            for fi13 := 0; fi13 < 19; fi13++ {
                p1 = "\n\t"
                _ = fi13
            }
        }
        for fi14 := 0; fi14 < 10; fi14++ {
            p2 = (127 >= 255)
            _ = fi14
        }
    }
    var v15 int8 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v7
    return 2147483647
}

func main() {
    var v1 float32 = 3.14159
    var v2 uint8 = 127
    var v3 int64 = 0
    v16 := GenS2{F0: (!((-113 != -1) && (!(!true))))}
    v3 += (v3 % ((v3) | 1))
    for fi17 := 0; fi17 < 16; fi17++ {
        if ((!true) || false) {
            v2 = (v2 | v2)
            v18 := GenS3{F0: -1, F1: 221}
            _ = v18
        }
        for fi19 := 0; fi19 < 11; fi19++ {
            var v20 int8 = -39
            _ = fi19
            _ = v20
        }
        var v21 float64 = float64(((v1 * v1) - ((0.0 - float32((3.14159 - ((1e10 + 1e10) + (0.1 / -2.25))))) * v1)))
        _ = fi17
        _ = v21
    }
    fmt.Println(v1, v2, v1)
    v22 := []int{-2147483648, 2147483647}
    for fi23 := 0; fi23 < 13; fi23++ {
        var v24 int16 = -755
        _ = fi23
        _ = v24
    }
    if true {
        for fi25 := 0; fi25 < 18; fi25++ {
            var v26 uint8 = (v2 ^ 42)
            _ = fi25
            _ = v26
        }
    }
    switch 42 {
    case -489:
        v27 := GenS2{F0: false}
        _ = v27
    }
    if ((-41 == 1) && (true && false)) {
        for fi28 := 0; fi28 < 10; fi28++ {
            fmt.Println(fi28)
            for fi29 := 0; fi29 < 10; fi29++ {
                var v30 bool = (!false)
                if false {
                    v2 -= 0
                }
                var v31 bool = false
                _ = fi29
                _ = v30
                _ = v31
            }
            _ = fi28
        }
        for fi32 := 0; fi32 < 12; fi32++ {
            for fi33 := 0; fi33 < 5; fi33++ {
                var v34 uint32 = 2147483648
                var v35 float64 = ((-2.25 + (((0.1 - 0.1) / 1e10) * 1.5)) - 1.5)
                if true {
                    _ = len(v22)
                }
                v1 += 3.14159
                if true {
                    v36 := GenS4{F0: genFn0(-32768), F1: 684, F2: 42}
                    fmt.Println(v2)
                    var v37 uint8 = 191
                    _ = v36
                    _ = v37
                } else {
                    var v38 int64 = 0
                    v34 = (^(v34))
                    v39 := GenS5{F0: 0, F1: v2, F2: 71}
                    _ = v38
                    _ = v39
                }
                _ = fi33
                _ = v34
                _ = v35
            }
            fmt.Println(v2, fi32, v1)
            var v40 int64 = (v3 | v3)
            _ = fi32
            _ = v40
        }
    } else {
        var v41 int16 = -1
        fmt.Println(v41, v41)
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v3
}
