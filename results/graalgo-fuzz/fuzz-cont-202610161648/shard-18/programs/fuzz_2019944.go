package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint; F1 int16; F2 uint32 }
type GenS4 struct { F0 uint64; F1 uint; F2 int32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int16 }
func genFn0() uint8 {
    v7 := GenS1{F0: 42, F1: 0}
    if (2 > 1) {
        v8 := []int{127, 9223372036854775807, 127, -861}
        if false {
            for fi9 := 0; fi9 < 19; fi9++ {
                _ = len(v8)
                _ = fi9
            }
        } else {
            switch "" {
            case "a":
                var v10 uint64 = 477
                _ = v10
            }
            if (2 <= 1) {
                var v11 string = ("hello" + "foo bar")
                v11 = v11
                v12 := GenS1{F0: -970, F1: 0}
                _ = v11
                _ = v12
            }
            var v13 uint64 = 42
            _ = v13
        }
        _ = v8
    } else {
        var v14 uint8 = 115
        if false {
            var v15 uint8 = (v14 >> ((255) & 63))
            var v16 int8 = 42
            _ = v15
            _ = v16
        }
        for fi17 := 0; fi17 < 5; fi17++ {
            if false {
                var v18 uint = 167
                var v19 uint16 = 563
                _ = v18
                _ = v19
            }
            _ = fi17
        }
        _ = v14
    }
    _ = v7
    return 0
}
func genFn1(p0 uint32, p1 bool, p2 bool) int64 {
    v20 := GenS1{F0: -1, F1: 0}
    p0 += 2
    var v21 int = -737
    v22 := GenS2{F0: -125, F1: "hello"}
    var v23 uint = 529
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    return 1
}

func main() {
    var v1 uint64 = 1
    var v2 int32 = -1
    v3 := make(map[uint32]string)
    var v4 float64 = (3.14159 / -2.25)
    v5 := make(map[string]int)
    v6 := make(map[int]uint32)
    _ = len(v6)
    switch 42 {
    case 32767:
        if (-1 <= -1) {
            if (!(177 <= 4294967295)) {
                v1 = (v1 / (((v1 << ((1) & 63))) | 1))
            } else {
                v24 := GenS1{F0: 783, F1: v1}
                v25 := GenS3{F0: 434, F1: 1, F2: 850}
                v26 := []int{935}
                _ = v24
                _ = v25
                _ = v26
            }
            v27 := GenS1{F0: -499, F1: 145}
            _ = v27
        }
    case -472:
        fmt.Println(v2, v4)
    case -1:
        fmt.Println(v1, v1, v2)
    }
    for fi28 := 0; fi28 < 1; fi28++ {
        var v29 int64 = 1
        var v30 int16 = -32768
        _ = fi28
        _ = v29
        _ = v30
    }
    v1 = (v1 - (v1 ^ (v1 * 2)))
    var v31 bool = true
    fmt.Println(v31, v31)
    if v31 {
        var v32 uint8 = genFn0()
        v33 := GenS2{F0: v2, F1: "foo bar"}
        var v34 uint32 = 2
        _ = v32
        _ = v33
        _ = v34
    }
    v1 = (v1 % ((v1) | 1))
    var v35 uint = 893
    var v36 uint8 = 43
    v35 = 1
    v37 := GenS1{F0: -1, F1: uint64(v35)}
    if v31 {
        if v31 {
            fmt.Println(v2)
            var v38 float64 = v4
            _ = v38
        } else {
            switch 663 {
            case 1000:
                for fi39 := 0; fi39 < 14; fi39++ {
                    for fi40 := 0; fi40 < 14; fi40++ {
                        var v41 uint = 0
                        fmt.Println(fi40)
                        v2 = -2147483648
                        v4 = v4
                        _ = fi40
                        _ = v41
                    }
                    v42 := GenS1{F0: -136, F1: (v1 / ((v1) | 1))}
                    _ = fi39
                    _ = v42
                }
            case 32768:
                for fi43 := 0; fi43 < 20; fi43++ {
                    switch -26 {
                    case 361:
                        var v44 uint64 = 2
                        _ = v44
                    default:
                        fmt.Println(v4, v1, v35)
                    }
                    _ = fi43
                }
            }
            var v45 int64 = -1
            _ = v45
        }
        v46 := GenS4{F0: (v1 * (v1 & 2)), F1: v35, F2: (v2 ^ v2)}
        _ = v46
    } else {
        var v47 uint64 = (v1 * 1)
        var v48 uint16 = 317
        v49 := []int{128, 2147483647, 127}
        _ = v47
        _ = v48
        _ = v49
    }
    var v50 uint16 = uint16(v35)
    v51 := GenS5{F0: v35}
    v52 := GenS6{F0: -415}
    v53 := []int{42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v53:", len(v53), v53)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v6
}
