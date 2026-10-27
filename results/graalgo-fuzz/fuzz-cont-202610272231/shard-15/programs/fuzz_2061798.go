package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64; F1 uint8 }
type GenS6 struct { F0 int16; F1 uint16 }
func genFn0(p0 int64, p1 uint8, p2 int8) int8 {
    for fi4 := 0; fi4 < 14; fi4++ {
        p0 = (p0 << ((42) & 63))
        v5 := GenS1{F0: p0}
        v6 := []int{fi4, 1, fi4, (-(fi4))}
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := GenS2{F0: 4294967295, F1: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p2 + 80)
}
func genFn1(p0 uint16, p1 bool) int16 {
    switch ("foo bar" + "\n\t") {
    case "hello":
        fmt.Println(p0, p1)
    case "\n\t":
        v8 := []int{-9223372036854775808, 83}
        _ = v8
    default:
        if (p1 || (!p1)) {
            v9 := GenS3{F0: 140, F1: 0}
            _ = v9
        } else {
            for fi10 := 0; fi10 < 7; fi10++ {
                v11 := GenS1{F0: 9223372036854775807}
                p1 = true
                _ = fi10
                _ = v11
            }
            fmt.Println(p0, p1)
            for fi12 := 0; fi12 < 5; fi12++ {
                var v13 uint16 = (^(p0))
                var v14 int32 = 0
                fmt.Println(v13, v14, p0)
                _ = fi12
                _ = v13
                _ = v14
            }
        }
    }
    p0 += (^(p0))
    p1 = (float32((0.0 + float64(float32(float64(0.0))))) > ((-2.25 + (0.1 + 3.14159)) + 3.14159))
    fmt.Println(p1, p1)
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return 32767
}

func main() {
    var v1 int8 = 92
    v2 := make(map[int]bool)
    v3 := []int{884, 944, -2147483648, 0}
    switch 39 {
    case 973:
        for fi15 := 0; fi15 < 7; fi15++ {
            for fi16 := 0; fi16 < 2; fi16++ {
                v1 = genFn0(-474, 127, -119)
                _ = fi16
            }
            _ = fi15
        }
    default:
        fmt.Println(v1)
    }
    v1 += -47
    var v17 int8 = v1
    v18 := GenS4{F0: ((1.5 / 1.0) > 1e10)}
    v19 := GenS1{F0: 9223372036854775807}
    if (false || false) {
        fmt.Println(v1)
        var v20 int = 128
        _ = v20
    }
    v17 = (v17 + genFn0(-593, 127, v1))
    switch 9223372036854775807 {
    case 127:
        var v21 uint32 = 301
        _ = v21
    case 2147483647:
        switch -585 {
        case -2147483648:
            v17 -= 0
        case -900:
            v22 := GenS3{F0: 865, F1: 42}
            _ = v22
        case 128:
            v17 = (v17 | v17)
        }
    case -319:
        v17 += genFn0(-1, 42, v17)
    }
    if false {
        if true {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            v3 = append(v3, 128)
        }
    } else {
        if (164 < 42) {
            var v23 string = ""
            var v24 float64 = float64((1e10 + (1.5 - (0.0 + 0.0))))
            _ = v23
            _ = v24
        } else {
            var v25 uint = 18446744073709551615
            _ = v25
        }
    }
    _ = len(v3)
    switch -9223372036854775808 {
    case -90:
        var v26 int8 = genFn0(1, 0, (v1 & v1))
        _ = v26
    case 128:
        delete(v2, 9223372036854775807)
    case -198:
        for fi27 := 0; fi27 < 19; fi27++ {
            if len(v3) > 0 {
                v3[0] = (fi27 >> ((418) & 63))
            }
            for fi28 := 0; fi28 < 6; fi28++ {
                v1 += 1
                _ = fi28
            }
            for fi29 := 0; fi29 < 17; fi29++ {
                fmt.Println(fi27)
                v1 = v1
                v30 := GenS5{F0: 42, F1: 42}
                v31 := GenS4{F0: false}
                v32 := GenS6{F0: 1000, F1: 32767}
                v1 = (v17 >> ((1) & 63))
                _ = fi29
                _ = v30
                _ = v31
                _ = v32
            }
            _ = fi27
        }
    default:
        var v33 string = "a"
        _ = v33
    }
    var v34 float32 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v34
}
