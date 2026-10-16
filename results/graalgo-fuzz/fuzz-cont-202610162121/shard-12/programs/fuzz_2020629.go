package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 int16; F1 uint16; F2 uint32 }
func genFn0() uint32 {
    v5 := []int{-9223372036854775808, 42, 128}
    var v6 uint16 = 42
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 int64, p1 int, p2 bool) int64 {
    var v7 uint16 = 0
    p0 = 1
    if (true && true) {
        switch (p1 << ((743) & 63)) {
        case 128:
            if (!false) {
                p1 += p1
                v8 := []int{1, (-(p1)), (p1 / ((p1) | 1)), p1}
                p0 += p0
                _ = v8
            } else {
                v9 := GenS1{F0: ("" + ""), F1: -359, F2: 0}
                _ = v9
            }
        case -9223372036854775808:
            v10 := []int{32768, (p1 - 9223372036854775807)}
            _ = v10
        }
        p0 = p0
    } else {
        var v11 uint8 = 112
        var v12 bool = (-1 == 2147483647)
        _ = v11
        _ = v12
    }
    if (65535 == 190) {
        var v13 uint32 = 65535
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 int16 = 1000
            _ = fi14
            _ = v15
        }
        _ = v13
    } else {
        var v16 uint64 = 42
        for fi17 := 0; fi17 < 17; fi17++ {
            v18 := []int{p1, int(v7)}
            for fi19 := 0; fi19 < 1; fi19++ {
                fmt.Println(v16, fi17, p1)
                p0 = (p0 - (p0 & 709))
                v20 := GenS2{F0: (p1 & 338)}
                _ = fi19
                _ = v20
            }
            _ = fi17
            _ = v18
        }
        var v21 uint8 = 2
        _ = v16
        _ = v21
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p0
}

func main() {
    var v1 uint16 = 42
    var v2 uint32 = 0
    var v3 int32 = 42
    var v4 int16 = 1
    var v22 int8 = 1
    v23 := GenS1{F0: "hello", F1: (v4 << ((258) & 63)), F2: v1}
    v24 := GenS1{F0: "hello", F1: (v4 & (v4 - (v4 * -32768))), F2: (^(v1))}
    v25 := []int{42, -2147483648, -1}
    for fi26 := 0; fi26 < 7; fi26++ {
        var v27 int64 = 463
        v28 := GenS2{F0: -698}
        v27 -= (v27 & v27)
        _ = fi26
        _ = v27
        _ = v28
    }
    v29 := GenS3{F0: (true && true), F1: v2}
    var v30 int64 = genFn1(1, 42, false)
    fmt.Println(v3, v1, v1)
    if (!(!(2 > 208))) {
        v2 -= (^(v2))
        var v31 int64 = (v30 | v30)
        _ = v31
    }
    switch 127 {
    case 0:
        var v32 int64 = (v30 & v30)
        _ = v32
    case 9223372036854775807:
        v4 -= (v4 >> ((2) & 63))
    case 61:
        fmt.Println(v1)
    }
    v2 += (v2 & 2147483647)
    if true {
        var v33 bool = false
        var v34 int8 = v22
        v30 -= (v30 & 42)
        _ = v33
        _ = v34
    }
    if len(v25) >= 2 {
        v25 = v25[1:len(v25)]
    }
    for fi35 := 0; fi35 < 17; fi35++ {
        if len(v25) >= 2 {
            v25 = v25[1:len(v25)]
        }
        var v36 int8 = v22
        for fi37 := 0; fi37 < 3; fi37++ {
            _ = len(v25)
            v30 = v30
            _ = fi37
        }
        _ = fi35
        _ = v36
    }
    v38 := GenS2{F0: 4294967295}
    if false {
        var v39 int8 = (v22 + (v22 - -73))
        var v40 int8 = v22
        _ = v39
        _ = v40
    }
    v41 := GenS4{F0: (v4 | v4), F1: (v1 | (v1 >> ((1) & 63))), F2: (v2 * v2)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v38
    _ = v4
    _ = v41
}
