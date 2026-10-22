package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 uint16; F2 int32 }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 uint16; F1 uint16 }
type GenS6 struct { F0 uint; F1 int32; F2 int64 }
func genFn0(p0 string, p1 uint8, p2 int8) uint8 {
    v7 := []int{}
    p1 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p1
}

func main() {
    var v1 int16 = -60
    var v2 uint8 = 128
    v3 := make(map[int]bool)
    var v4 uint32 = 0
    var v5 string = ""
    v6 := make(map[uint32]bool)
    v8 := GenS1{F0: true, F1: 65535}
    var v9 int = -339
    if (!false) {
        switch 0 {
        case -971:
            v10 := GenS2{F0: v2}
            _ = v10
        default:
            var v11 int64 = 9223372036854775807
            _ = v11
        }
        v12 := GenS3{F0: -443, F1: 943, F2: 833}
        if (269 < uint16(v2)) {
            v13 := GenS4{F0: 936, F1: v4}
            if ((v4 / ((291) | 1)) <= 903) {
                for fi14 := 0; fi14 < 9; fi14++ {
                    delete(v6, v4)
                    for fi15 := 0; fi15 < 5; fi15++ {
                        var v16 uint64 = 42
                        _ = fi15
                        _ = v16
                    }
                    _ = fi14
                }
                v1 += v1
                v9 += v9
            }
            _ = v13
        }
        _ = v12
    }
    v4 -= 700
    v17 := GenS3{F0: 694, F1: 620, F2: 2147483647}
    for fi18 := 0; fi18 < 16; fi18++ {
        fmt.Println(v9, v9, v9)
        _ = fi18
    }
    v5 = ((v5 + "\n\t") + ("foo bar" + v5))
    if true {
        v1 = (v1 ^ v1)
        v19 := GenS1{F0: ((((v9 < (v9 * v9)) && ((true || false) || false)) && (!((v9 * v9) <= (v9 & (-(v9)))))) && (!true)), F1: 255}
        v4 -= (^(v4))
        _ = v19
    }
    for fi20 := 0; fi20 < 11; fi20++ {
        v21 := GenS5{F0: 653, F1: 0}
        if ("foo bar" != v5) {
            v3[v9] = false
            fmt.Println(v1, fi20, v1)
        }
        _ = fi20
        _ = v21
    }
    if false {
        v9 -= (v9 >> ((42) & 63))
        var v22 uint8 = v2
        v2 += (v22 & (^(v2)))
        _ = v22
    }
    if true {
        var v23 uint64 = 955
        switch (v9 + -1) {
        case 2147483647:
            v24 := []int{}
            _ = v24
        case 42:
            v25 := GenS5{F0: 2, F1: 65535}
            _ = v25
        default:
            var v26 int32 = 42
            _ = v26
        }
        _ = v23
    } else {
        v27 := []int{}
        v2 = (v2 ^ (v2 * 1))
        v5 = ("foo bar" + (v5 + v5))
        _ = v27
    }
    var v28 int8 = int8(v4)
    v5 = v5
    if (!(true || true)) {
        switch (v9 / ((42) | 1)) {
        case -362:
            fmt.Println(v9, v5)
        default:
            v29 := GenS6{F0: 255, F1: 2147483647, F2: 0}
            _ = v29
        }
        v2 -= v2
        v30 := v3[(v9 + (v9 >> ((4294967295) & 63)))]
        _ = v30
    } else {
        for fi31 := 0; fi31 < 9; fi31++ {
            v32 := []int{(v9 - v9), fi31, (fi31 + (fi31 * v9)), 4294967295}
            fmt.Println(v2, v2, v28)
            fmt.Println(v28, v2)
            _ = fi31
            _ = v32
        }
        var v33 uint32 = (v4 >> ((42) & 63))
        _ = v33
    }
    var v34 int = v9
    fmt.Println(v4, v9, v9)
    switch (v9 * (-(v9))) {
    case 128:
        for fi35 := 0; fi35 < 14; fi35++ {
            v9 = v9
            _ = fi35
        }
    case 9223372036854775807:
        switch v5 {
        case "\n\t":
            v3[int(v1)] = (true && true)
        case "a":
            v9 += v34
        default:
            switch ((v5 + ("a" + "foo bar")) + (v5 + "foo bar")) {
            case "":
                v28 = v28
            }
        }
    case -881:
        v2 += genFn0("", v2, v28)
    default:
        switch (v34 | (v34 / (((v9 / ((v9) | 1))) | 1))) {
        case 32768:
            var v36 uint8 = (v2 + 114)
            _ = v36
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
