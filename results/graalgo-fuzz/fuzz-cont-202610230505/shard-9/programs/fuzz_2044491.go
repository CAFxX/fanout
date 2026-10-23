package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS3 struct { F0 bool; F1 int8; F2 uint32 }
type GenS4 struct { F0 uint64; F1 int16 }
func genFn0(p0 uint64, p1 uint16, p2 uint16) int32 {
    v5 := []int{-44, -9223372036854775808}
    var v6 int8 = 0
    var v7 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return -204
}

func main() {
    var v1 float64 = (0.1 - (1.5 * 3.14159))
    v2 := []int{}
    v3 := make(map[uint32]int)
    var v4 uint = 255
    delete(v3, 2)
    for fi8 := 0; fi8 < 4; fi8++ {
        if true {
            switch (fi8 | fi8) {
            case -2147483648:
                var v9 uint16 = 0
                _ = v9
            case 2147483647:
                for fi10 := 0; fi10 < 11; fi10++ {
                    v11 := GenS1{F0: 756}
                    var v12 float64 = ((0.1 + -2.25) - v1)
                    var v13 float32 = ((1.5 / 1.5) + ((0.1 * ((0.1 / 1.0) + 0.1)) * 1e10))
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                }
            case 0:
                v14 := GenS1{F0: 1}
                _ = v14
            default:
                _ = len(v3)
            }
            _ = len(v2)
            var v15 uint64 = 4294967295
            _ = v15
        }
        _ = fi8
    }
    for fi16 := 0; fi16 < 19; fi16++ {
        for fi17 := 0; fi17 < 4; fi17++ {
            _ = len(v3)
            var v18 int8 = 42
            v19 := GenS1{F0: 723}
            _ = fi17
            _ = v18
            _ = v19
        }
        v4 = v4
        _ = fi16
    }
    var v20 uint = v4
    var v21 int8 = 0
    if false {
        for fi22 := 0; fi22 < 3; fi22++ {
            v4 -= v20
            _ = fi22
        }
    }
    v1 -= (0.1 * v1)
    if (0.0 <= 0.0) {
        _ = len(v3)
        switch "foo bar" {
        case "hello":
            var v23 int = -1
            _ = v23
        case "foo bar":
            var v24 float64 = 1.5
            _ = v24
        default:
            v21 = 100
        }
        fmt.Println(v4)
    } else {
        var v25 uint8 = 255
        if ("\n\t" <= (("" + (("foo bar" + "a") + "foo bar")) + "")) {
            for fi26 := 0; fi26 < 19; fi26++ {
                if false {
                    var v27 int16 = 32767
                    var v28 uint32 = 307
                    switch (fi26 | int(v28)) {
                    case 127:
                        var v29 uint32 = v28
                        _ = v29
                    }
                    _ = v27
                    _ = v28
                } else {
                    v30 := GenS2{F0: 0, F1: v20, F2: uint16(v25)}
                    _ = v30
                }
                var v31 int = (fi26 >> ((v20) & 63))
                var v32 int8 = v21
                _ = fi26
                _ = v31
                _ = v32
            }
            var v33 int8 = v21
            _ = v33
        }
        _ = v25
    }
    if len(v2) > 0 {
        v2[0] = 204
    }
    switch 9223372036854775807 {
    case 127:
        v34 := GenS1{F0: 184}
        _ = v34
    case 9223372036854775807:
        fmt.Println(v4)
    }
    switch 32768 {
    case -918:
        for fi35 := 0; fi35 < 12; fi35++ {
            fmt.Println(fi35)
            v36 := GenS2{F0: 42, F1: v20, F2: 255}
            v4 = (v20 - v20)
            _ = fi35
            _ = v36
        }
    case -2147483648:
        for fi37 := 0; fi37 < 19; fi37++ {
            v38 := GenS1{F0: 2147483648}
            switch (fi37 << ((1) & 63)) {
            case 138:
                v21 = 36
            case 1:
                var v39 uint = (^(v20))
                _ = v39
            case 0:
                if (false || (!(fi37 > -9223372036854775808))) {
                    var v40 int16 = -145
                    v41 := GenS3{F0: (!(!(true || false))), F1: v21, F2: 4294967295}
                    _ = len(v3)
                    _ = v40
                    _ = v41
                }
            }
            var v42 int64 = 0
            _ = fi37
            _ = v38
            _ = v42
        }
    default:
        v21 += (v21 >> (((^(v4))) & 63))
    }
    var v43 int16 = 32767
    v4 = v4
    v44 := GenS4{F0: 939, F1: (v43 >> (((^(v4))) & 63))}
    v43 += (-(v43))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v43
    _ = v44
}
