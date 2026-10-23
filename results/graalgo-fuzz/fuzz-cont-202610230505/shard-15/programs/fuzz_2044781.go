package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 string; F1 bool; F2 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 uint8; F2 int64 }
func genFn0(p0 int8, p1 uint64) int16 {
    for fi6 := 0; fi6 < 12; fi6++ {
        if false {
            var v7 uint32 = 65535
            var v8 int16 = -32768
            v9 := GenS1{F0: 230, F1: p1}
            _ = v7
            _ = v8
            _ = v9
        } else {
            for fi10 := 0; fi10 < 14; fi10++ {
                v11 := GenS2{F0: "a", F1: false, F2: true}
                _ = fi10
                _ = v11
            }
        }
        v12 := []int{fi6, fi6, 515}
        _ = fi6
        _ = v12
    }
    switch 32767 {
    case -9223372036854775808:
        p1 += 42
    case 32768:
        for fi13 := 0; fi13 < 12; fi13++ {
            p1 -= (p1 | (p1 / (((p1 & 4294967295)) | 1)))
            _ = fi13
        }
    }
    p1 -= p1
    for fi14 := 0; fi14 < 12; fi14++ {
        var v15 int32 = 2147483647
        _ = fi14
        _ = v15
    }
    var v16 bool = true
    _ = p0
    _ = p1
    _ = v16
    return 42
}
func genFn1(p0 uint8, p1 bool) uint32 {
    if p1 {
        switch "\n\t" {
        case "a":
            var v17 bool = (100 == 57)
            _ = v17
        case "":
            var v18 uint32 = 0
            _ = v18
        case "foo bar":
            for fi19 := 0; fi19 < 16; fi19++ {
                v20 := []int{fi19, fi19}
                var v21 int16 = -32768
                v22 := GenS3{F0: (v21 - v21)}
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
        }
        fmt.Println(p0, p0)
    } else {
        var v23 int32 = 342
        _ = v23
    }
    p0 = p0
    _ = p0
    _ = p1
    return 2147483648
}

func main() {
    var v1 uint = 961
    var v2 float32 = 3.14159
    var v3 uint32 = 826
    var v4 int8 = 42
    v5 := make(map[int]uint32)
    switch 32768 {
    case 2147483647:
        switch -2147483648 {
        case 685:
            var v24 int32 = -2147483648
            _ = v24
        case 128:
            if (0 > -164) {
                var v25 int8 = (v4 + (v4 / (((v4 * 100)) | 1)))
                v26 := []int{-233, -2147483648, 91, 9223372036854775807}
                v27 := GenS4{F0: -1, F1: 42, F2: -956}
                _ = v25
                _ = v26
                _ = v27
            }
        case 3:
            v2 = (3.14159 * (v2 - v2))
        default:
            fmt.Println(v2, v2)
        }
    case 0:
        v28 := []int{}
        _ = v28
    case -1:
        switch "" {
        case "a":
            fmt.Println(v4, v2, v2)
        }
    }
    var v29 bool = ((42 >= 9223372036854775807) && ((-1 < -2147483648) || false))
    v30 := GenS3{F0: genFn0(v4, 592)}
    var v31 int32 = 42
    v31 += (v31 & 2147483647)
    for fi32 := 0; fi32 < 6; fi32++ {
        v33 := v5[(fi32 % ((int(v4)) | 1))]
        _ = fi32
        _ = v33
    }
    var v34 string = ("hello" + "hello")
    switch "foo bar" {
    case "":
        if true {
            v31 = (v31 + (-(v31)))
            _ = len(v5)
        } else {
            v31 = -2147483648
        }
    case "hello":
        if true {
            if true {
                var v35 string = (((v34 + v34) + v34) + "foo bar")
                _ = v35
            } else {
                var v36 uint32 = v3
                _ = v36
            }
        } else {
            for fi37 := 0; fi37 < 19; fi37++ {
                var v38 int8 = v4
                _ = fi37
                _ = v38
            }
        }
    case "\n\t":
        for fi39 := 0; fi39 < 17; fi39++ {
            var v40 uint8 = 127
            v31 += -2147483648
            _ = fi39
            _ = v40
        }
    }
    var v41 int16 = -650
    if v29 {
        v42 := GenS3{F0: (v41 & (v41 - v41))}
        for fi43 := 0; fi43 < 2; fi43++ {
            v4 = 127
            _ = fi43
        }
        _ = v42
    }
    fmt.Println(v41)
    v44 := GenS3{F0: (v41 << ((v1) & 63))}
    delete(v5, int(v41))
    switch v34 {
    case "foo bar":
        v3 = 2
    case "a":
        if (v31 < (v31 >> (((v1 & v1)) & 63))) {
            v41 = int16(v4)
        }
    default:
        switch "foo bar" {
        case "":
            fmt.Println(v4)
        }
    }
    v45 := GenS3{F0: (v41 * (v41 << ((v1) & 63)))}
    v46 := GenS1{F0: v3, F1: 42}
    v47 := GenS4{F0: 0, F1: 255, F2: -267}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v4
    _ = v41
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
}
