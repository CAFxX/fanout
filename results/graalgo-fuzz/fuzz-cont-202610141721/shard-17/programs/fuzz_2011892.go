package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 uint64; F2 uint32 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 int; F1 string }
type GenS5 struct { F0 bool; F1 int32 }
func genFn0(p0 int16, p1 uint8) int {
    var v7 uint64 = 4294967295
    if true {
        var v8 uint8 = 0
        for fi9 := 0; fi9 < 17; fi9++ {
            v10 := GenS1{F0: 230}
            _ = fi9
            _ = v10
        }
        _ = v8
    } else {
        if true {
            for fi11 := 0; fi11 < 13; fi11++ {
                var v12 int32 = 0
                v13 := GenS2{F0: (fi11 >> ((255) & 63)), F1: (^(v7)), F2: 1}
                p0 -= p0
                _ = fi11
                _ = v12
                _ = v13
            }
            v14 := GenS2{F0: 4294967295, F1: v7, F2: 2147483648}
            v7 = (v7 << ((65535) & 63))
            _ = v14
        } else {
            fmt.Println(p0)
            var v15 uint32 = 2147483647
            v16 := []int{42, 0, -9223372036854775808}
            _ = v15
            _ = v16
        }
        switch -264 {
        case 42:
            var v17 uint = 253
            _ = v17
        case 4294967295:
            v18 := []int{}
            _ = v18
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return -772
}

func main() {
    var v1 int = -132
    v2 := make(map[string]bool)
    var v3 int = (^(v1))
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]uint32)
    var v6 uint64 = 42
    v19 := GenS3{F0: 1000, F1: v6}
    v4[65535] = (0 >= 128)
    var v20 float64 = float64(-2.25)
    for fi21 := 0; fi21 < 14; fi21++ {
        var v22 uint8 = 42
        v23 := []int{v1, 985, 2147483647}
        v20 = (v20 - (v20 - (1e10 / (v20 / 0.1))))
        _ = fi21
        _ = v22
        _ = v23
    }
    _ = len(v4)
    v3 -= (v1 / ((128) | 1))
    if (true && false) {
        v24 := []int{32768, (v1 * v3)}
        var v25 uint8 = 234
        _ = v24
        _ = v25
    }
    fmt.Println(v1, v1, v3)
    var v26 string = "hello"
    for fi27 := 0; fi27 < 7; fi27++ {
        switch v1 {
        case -943:
            v28 := GenS3{F0: 983, F1: 606}
            _ = v28
        case -1:
            v29 := GenS4{F0: (-(v3)), F1: v26}
            _ = v29
        case 1:
            v26 = v26
        }
        for fi30 := 0; fi30 < 5; fi30++ {
            v20 += float64(-2.25)
            fmt.Println(v20)
            var v31 int = 32767
            _ = fi30
            _ = v31
        }
        _ = fi27
    }
    var v32 int64 = 1
    v33 := v2[v26]
    if v33 {
        for fi34 := 0; fi34 < 20; fi34++ {
            for fi35 := 0; fi35 < 4; fi35++ {
                var v36 int64 = 305
                delete(v2, (v26 + ""))
                v37 := []int{42, (fi35 << ((2) & 63))}
                _ = fi35
                _ = v36
                _ = v37
            }
            _ = fi34
        }
        switch (v1 >> ((4294967295) & 63)) {
        case 32767:
            fmt.Println(v26)
        case 1:
            for fi38 := 0; fi38 < 13; fi38++ {
                for fi39 := 0; fi39 < 2; fi39++ {
                    var v40 int64 = 42
                    var v41 float64 = v20
                    v20 = (((1e10 * -2.25) * 1.5) / (v20 / 1.0))
                    _ = fi39
                    _ = v40
                    _ = v41
                }
                if false {
                    v42 := []int{780, v1, (v1 ^ genFn0(42, 46)), (v1 | v3)}
                    var v43 uint = 42
                    _ = v42
                    _ = v43
                } else {
                    switch v26 {
                    case "":
                        v44 := GenS1{F0: 394}
                        _ = v44
                    }
                    v45 := GenS3{F0: 32767, F1: 42}
                    _ = v45
                }
                v33 = v33
                _ = fi38
            }
        case 739:
            v46 := GenS5{F0: (v33 || v33), F1: 42}
            _ = v46
        default:
            if v33 {
                v1 -= v1
                v26 = v26
                var v47 uint32 = 65535
                _ = v47
            }
        }
        var v48 int32 = 692
        _ = v48
    }
    var v49 uint16 = 32767
    v26 = v26
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v49
    _ = v5
    _ = v6
}
