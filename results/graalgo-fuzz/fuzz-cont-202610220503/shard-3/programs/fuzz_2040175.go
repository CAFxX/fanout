package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 int32; F2 uint }

func main() {
    var v1 int16 = -929
    v2 := make(map[uint32]int)
    var v3 int = 650
    var v4 int32 = 1
    switch (^(v3)) {
    case 412:
        delete(v2, 42)
    default:
        var v5 int8 = -128
        _ = v5
    }
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := GenS1{F0: 98}
        v1 = v1
        switch (fi6 ^ v3) {
        case 32767:
            v4 = v4
        }
        _ = fi6
        _ = v7
    }
    v8 := GenS2{F0: 127, F1: v4, F2: 173}
    for fi9 := 0; fi9 < 18; fi9++ {
        if (!false) {
            if true {
                v4 = (v4 | 0)
                v10 := []int{}
                _ = v10
            }
        } else {
            var v11 int64 = 1
            _ = len(v2)
            _ = v11
        }
        if (0 <= 127) {
            var v12 int8 = 0
            _ = v12
        } else {
            var v13 uint32 = 150
            for fi14 := 0; fi14 < 8; fi14++ {
                switch ("foo bar" + "\n\t") {
                case "a":
                    v13 -= v13
                case "\n\t":
                    v15 := []int{42, v3, fi9}
                    _ = v15
                case "":
                    v16 := GenS2{F0: 127, F1: (v4 >> ((255) & 63)), F2: 919}
                    _ = v16
                }
                switch v3 {
                case -845:
                    v13 -= v13
                case 710:
                    var v17 int64 = 1
                    _ = v17
                case -1:
                    v18 := []int{4294967295, fi9, fi9}
                    _ = v18
                }
                v19 := GenS1{F0: (v13 * (v13 ^ (v13 - (^(v13)))))}
                _ = fi14
                _ = v19
            }
            for fi20 := 0; fi20 < 3; fi20++ {
                v3 -= v3
                fmt.Println(v3, v1)
                v1 -= (v1 - -1)
                _ = fi20
            }
            _ = v13
        }
        _ = fi9
    }
    fmt.Println(v4)
    if (0 <= v1) {
        v2[938] = 1
    }
    var v21 uint64 = 2
    if (v1 >= (v1 | v1)) {
        delete(v2, 2147483647)
        fmt.Println(v21, v21)
        if false {
            for fi22 := 0; fi22 < 15; fi22++ {
                for fi23 := 0; fi23 < 8; fi23++ {
                    var v24 int8 = 127
                    _ = fi23
                    _ = v24
                }
                fmt.Println(v21, fi22)
                fmt.Println(fi22, fi22)
                _ = fi22
            }
        } else {
            v1 += -32768
            v25 := []int{(v3 & int(v4)), (v3 / ((4294967295) | 1)), v3}
            if (!(true && true)) {
                for fi26 := 0; fi26 < 8; fi26++ {
                    var v27 int16 = v1
                    delete(v2, 65535)
                    var v28 uint64 = 1
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                fmt.Println(v3, v3, v4)
                v29 := GenS2{F0: 0, F1: (v4 - (-(v4))), F2: 0}
                _ = v29
            } else {
                v3 -= (v3 & int(v4))
                var v30 int8 = -120
                _ = v30
            }
            _ = v25
        }
    }
    v31 := GenS1{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v31
    _ = v4
    _ = v8
}
