package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 int; F1 int64; F2 int8 }
type GenS4 struct { F0 bool; F1 int8 }
func genFn0(p0 uint32, p1 uint16, p2 int64) int64 {
    if (1.5 != 0.0) {
        var v6 uint32 = (^(p0))
        _ = v6
    }
    v7 := GenS1{F0: 12, F1: 42, F2: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -1
}
func genFn1() uint32 {
    v8 := []int{}
    v9 := GenS2{F0: 4294967295, F1: 539}
    if true {
        if true {
            var v10 uint16 = 2
            if false {
                var v11 uint64 = 2
                _ = v11
            }
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            _ = v10
        }
        v12 := GenS2{F0: 0, F1: 0}
        var v13 int32 = -1
        _ = v12
        _ = v13
    }
    switch "hello" {
    case "foo bar":
        _ = len(v8)
    }
    _ = v8
    _ = v9
    return 2147483647
}

func main() {
    var v1 uint8 = 132
    var v2 uint64 = 2
    var v3 int64 = 0
    var v4 int64 = 42
    var v5 uint8 = v1
    if false {
        v14 := []int{32767, 1}
        _ = v14
    } else {
        v5 -= v5
        var v15 uint8 = v5
        _ = v15
    }
    switch -480 {
    case 42:
        v16 := []int{4294967295, 128, int(v1)}
        _ = v16
    }
    var v17 int = 9223372036854775807
    v18 := GenS1{F0: 65535, F1: v2, F2: 605}
    if true {
        v2 = v2
    }
    v17 -= -743
    var v19 bool = (1 < 255)
    var v20 int = 32768
    var v21 int = (v20 * (-(v17)))
    var v22 int64 = -9223372036854775808
    v23 := []int{v20, (^(v21)), 127}
    if (916 > 65535) {
        switch (-(v20)) {
        case 2147483647:
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
        case 110:
            for fi24 := 0; fi24 < 1; fi24++ {
                var v25 float32 = float32((0.1 * 1.5))
                v23 = append(v23, v20)
                if true {
                    var v26 uint = 466
                    _ = v26
                }
                _ = fi24
                _ = v25
            }
        }
        for fi27 := 0; fi27 < 10; fi27++ {
            if (((v4 & v22) >= 9223372036854775807) && (false && v19)) {
                v28 := GenS2{F0: v2, F1: 42}
                v22 = (-(v3))
                _ = v28
            }
            for fi29 := 0; fi29 < 9; fi29++ {
                v3 = (-(v4))
                v21 = v17
                _ = fi29
            }
            _ = fi27
        }
        fmt.Println(v5, v5)
    }
    var v30 int64 = genFn0(336, 503, (v3 - v22))
    v3 -= v4
    v31 := GenS2{F0: (v2 % ((v2) | 1)), F1: v2}
    if ((1 != 1) || false) {
        var v32 float64 = ((float64(0.1) * 0.1) * 1.5)
        var v33 uint64 = v2
        if len(v23) > 0 {
            v23[0] = int(v33)
        }
        _ = v32
        _ = v33
    }
    var v34 uint16 = 483
    switch "a" {
    case "hello":
        switch (("a" + "a") + ("" + ("\n\t" + "\n\t"))) {
        case "foo bar":
            for fi35 := 0; fi35 < 12; fi35++ {
                v22 = genFn0(uint32(v2), 197, v22)
                if len(v23) >= 2 {
                    v23 = v23[1:len(v23)]
                }
                _ = fi35
            }
        default:
            var v36 int32 = 2147483647
            _ = v36
        }
    case "a":
        switch (v20 << ((0) & 63)) {
        case 128:
            v37 := GenS2{F0: v2, F1: (v2 / ((v2) | 1))}
            _ = v37
        }
    case "":
        if (1 >= 0) {
            v34 = (^(v34))
        } else {
            for fi38 := 0; fi38 < 14; fi38++ {
                v39 := GenS3{F0: (v20 & (fi38 | v20)), F1: -563, F2: 1}
                v40 := GenS1{F0: v34, F1: (v2 ^ (v2 - (^(v2)))), F2: 32767}
                for fi41 := 0; fi41 < 19; fi41++ {
                    var v42 uint8 = (v5 >> ((238) & 63))
                    var v43 uint32 = 2147483647
                    v44 := GenS4{F0: (127 >= 127), F1: 127}
                    _ = fi41
                    _ = v42
                    _ = v43
                    _ = v44
                }
                _ = fi38
                _ = v39
                _ = v40
            }
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
            v45 := GenS4{F0: false, F1: 127}
            _ = v45
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v4
    _ = v5
}
