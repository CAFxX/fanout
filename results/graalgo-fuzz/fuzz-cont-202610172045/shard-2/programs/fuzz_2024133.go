package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 bool }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 int8; F1 bool; F2 int64 }
type GenS5 struct { F0 int; F1 uint16 }
func genFn0(p0 bool, p1 int8, p2 uint) int {
    if p0 {
        var v5 int64 = 1
        _ = v5
    }
    v6 := []int{0, -114, 9, int(p1)}
    var v7 uint16 = uint16(p2)
    var v8 int = 453
    _ = len(v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return v8
}
func genFn1(p0 int, p1 uint8) uint32 {
    var v9 int64 = -161
    v10 := []int{122, (p0 | 979), (p0 << ((2) & 63))}
    v11 := GenS1{F0: 1, F1: 65535, F2: (true && (true || (42 < 65535)))}
    switch p0 {
    case 2147483647:
        var v12 uint16 = 0
        _ = v12
    case 4294967295:
        if true {
            if ((p0 & (p0 % ((p0) | 1))) <= 4294967295) {
                for fi13 := 0; fi13 < 17; fi13++ {
                    v9 -= (v9 + v9)
                    p0 = (fi13 - p0)
                    _ = len(v10)
                    _ = fi13
                }
            }
            var v14 int16 = -478
            _ = v14
        } else {
            switch "foo bar" {
            case "foo bar":
                var v15 int = 579
                _ = v15
            case "":
                var v16 int64 = (v9 << ((255) & 63))
                _ = v16
            }
        }
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        switch (fi17 & (-(p0))) {
        case 701:
            v18 := GenS2{F0: 18446744073709551615, F1: 65535}
            _ = v18
        case -940:
            p0 += (p0 - (fi17 - (p0 >> ((1000) & 63))))
        case -172:
            v9 = (-(v9))
        }
        _ = fi17
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return 590
}

func main() {
    var v1 uint32 = 2147483648
    var v2 int = -133
    var v3 int8 = -128
    var v4 int16 = 234
    v19 := GenS3{F0: 127, F1: 0}
    var v20 uint64 = 2
    var v21 uint16 = 0
    v3 = 42
    var v22 int16 = (v4 >> ((270) & 63))
    var v23 int16 = (v22 & 802)
    fmt.Println(v20, v1)
    v24 := GenS3{F0: 176, F1: 0}
    v25 := GenS2{F0: 1, F1: 199}
    v26 := []int{(-(v2)), v2}
    switch 4294967295 {
    case 32767:
        for fi27 := 0; fi27 < 6; fi27++ {
            var v28 uint8 = 15
            v29 := GenS4{F0: v3, F1: false, F2: 42}
            _ = fi27
            _ = v28
            _ = v29
        }
    case -126:
        v30 := GenS1{F0: (v21 & (^(v21))), F1: 42, F2: true}
        _ = v30
    }
    for fi31 := 0; fi31 < 8; fi31++ {
        if true {
            v32 := GenS1{F0: (v21 - v21), F1: v21, F2: ((3.14159 * 0.0) != 0.1)}
            for fi33 := 0; fi33 < 20; fi33++ {
                v34 := GenS3{F0: 158, F1: 279}
                _ = fi33
                _ = v34
            }
            _ = v32
        } else {
            v35 := GenS5{F0: fi31, F1: 636}
            _ = v35
        }
        _ = fi31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
