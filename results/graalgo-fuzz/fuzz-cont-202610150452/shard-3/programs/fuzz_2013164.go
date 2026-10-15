package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 int32; F2 int8 }
func genFn0(p0 int8, p1 uint8) uint64 {
    switch 9223372036854775807 {
    case -1:
        p0 = (p0 * (p0 << ((691) & 63)))
    case 739:
        fmt.Println(p0)
    case 1:
        v6 := GenS1{F0: true, F1: 0, F2: 943}
        _ = v6
    default:
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 uint32 = 84
            _ = fi7
            _ = v8
        }
    }
    p1 = p1
    _ = p0
    _ = p1
    return 2
}
func genFn1(p0 string, p1 uint32, p2 int8) int16 {
    for fi9 := 0; fi9 < 17; fi9++ {
        v10 := GenS1{F0: true, F1: uint(p1), F2: 255}
        for fi11 := 0; fi11 < 19; fi11++ {
            switch -658 {
            case -835:
                fmt.Println(p0, p2)
            case 0:
                var v12 uint32 = 2
                _ = v12
            case 9223372036854775807:
                p0 = ("hello" + "")
            }
            v13 := []int{2147483647}
            if false {
                p1 -= p1
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
            }
            _ = fi11
            _ = v13
        }
        _ = fi9
        _ = v10
    }
    switch p0 {
    case "\n\t":
        fmt.Println(p1, p0, p0)
    case "foo bar":
        if true {
            var v14 int8 = 100
            var v15 bool = false
            _ = v14
            _ = v15
        }
    }
    if (true && true) {
        if true {
            var v16 int8 = (p2 - (p2 / ((p2) | 1)))
            var v17 uint64 = 840
            var v18 uint64 = (v17 * (^(v17)))
            _ = v16
            _ = v17
            _ = v18
        }
        var v19 uint8 = 42
        var v20 int = -1
        _ = v19
        _ = v20
    }
    if true {
        var v21 int32 = 0
        _ = v21
    } else {
        var v22 uint64 = 915
        var v23 float32 = ((-2.25 / 1.5) - float32(float64(float32(1e10))))
        _ = v22
        _ = v23
    }
    _ = p0
    _ = p1
    _ = p2
    return 32767
}

func main() {
    v1 := []uint32{65535, 65535, 521, 42}
    v2 := []uint32{2147483647, 999, 2147483648, 1}
    var v3 int16 = -1
    v4 := make(map[uint32]int)
    v5 := []uint32{4294967295, 69}
    var v24 int64 = 944
    switch "hello" {
    case "foo bar":
        fmt.Println(v3, v3)
    case "a":
        var v25 uint32 = 728
        _ = v25
    default:
        v24 += v24
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        for fi27 := 0; fi27 < 8; fi27++ {
            var v28 int8 = int8(fi26)
            fmt.Println(v28)
            _ = fi27
            _ = v28
        }
        v29 := GenS1{F0: false, F1: 2, F2: 155}
        if (!false) {
            v24 -= 42
        }
        _ = fi26
        _ = v29
    }
    var v30 float64 = float64(-2.25)
    var v31 uint = 624
    var v32 int32 = 42
    for fi33 := 0; fi33 < 7; fi33++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v34 int = (fi33 >> (((v31 / (((v31 << (((v31 / ((1) | 1))) & 63))) | 1))) & 63))
        _ = fi33
        _ = v34
    }
    switch "" {
    case "hello":
        if true {
            for fi35 := 0; fi35 < 18; fi35++ {
                var v36 uint32 = 65535
                var v37 uint32 = 773
                _ = fi35
                _ = v36
                _ = v37
            }
            v31 = 339
        } else {
            var v38 float64 = (((1.5 - v30) + v30) * v30)
            var v39 int64 = v24
            var v40 int64 = -9223372036854775808
            _ = v38
            _ = v39
            _ = v40
        }
    default:
        switch 2147483647 {
        case 2147483647:
            v32 = v32
        default:
            v41 := GenS2{F0: 776}
            _ = v41
        }
    }
    v42 := GenS3{F0: 42, F1: (v32 - v32), F2: -41}
    v43 := GenS1{F0: false, F1: v31, F2: (v31 + (^(v31)))}
    if len(v1) > 0 {
        v1[0] = 2147483648
    }
    v44 := GenS1{F0: (true && false), F1: 0, F2: 2}
    var v45 bool = false
    var v46 int64 = v24
    var v47 string = ("hello" + "a")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
}
