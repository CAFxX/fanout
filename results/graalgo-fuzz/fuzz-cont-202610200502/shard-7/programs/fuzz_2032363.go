package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int64; F2 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 uint64, p1 bool) uint16 {
    fmt.Println(p0, p0)
    if p1 {
        v6 := []int{128}
        _ = v6
    } else {
        var v7 uint32 = 42
        v8 := []int{int(p0), 4294967295}
        var v9 float64 = (0.1 - 0.1)
        _ = v7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: p0}
    _ = p0
    _ = p1
    _ = v10
    return 2
}

func main() {
    var v1 int16 = 0
    var v2 uint = 255
    var v3 float32 = 0.0
    var v4 int8 = -1
    var v5 string = (("a" + "foo bar") + "foo bar")
    fmt.Println(v5, v2)
    switch v5 {
    case "\n\t":
        var v11 uint16 = 1
        _ = v11
    }
    var v12 string = v5
    v13 := GenS1{F0: 117}
    v14 := GenS2{F0: 286, F1: 42, F2: v1}
    v5 = v12
    v2 -= (v2 * 42)
    var v15 uint64 = 453
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 uint16 = genFn0(v15, true)
        v15 -= (v15 << (((v2 | v2)) & 63))
        _ = fi16
        _ = v17
    }
    if (!(1e10 != 1e10)) {
        var v18 int32 = 75
        _ = v18
    }
    switch -2147483648 {
    case 32767:
        if true {
            switch "" {
            case "":
                v19 := GenS1{F0: (v15 ^ 2)}
                _ = v19
            case "\n\t":
                if false {
                    var v20 uint = v2
                    var v21 uint16 = 65535
                    v3 += v3
                    _ = v20
                    _ = v21
                }
            default:
                fmt.Println(v1, v12)
            }
        }
    default:
        for fi22 := 0; fi22 < 2; fi22++ {
            switch fi22 {
            case 4294967295:
                v23 := []int{fi22, (fi22 - fi22), fi22, 32767}
                _ = v23
            default:
                v24 := GenS3{F0: 121}
                _ = v24
            }
            v3 = float32(0.1)
            for fi25 := 0; fi25 < 12; fi25++ {
                v26 := []int{(fi22 * fi25), fi25, (fi25 + (fi22 & (fi22 >> ((v2) & 63))))}
                var v27 int64 = -9223372036854775808
                if false {
                    if ((v4 >> ((42) & 63)) != 1) {
                        var v28 uint32 = 569
                        var v29 int = 2147483647
                        if true {
                            v30 := GenS2{F0: (v27 | -1), F1: -1, F2: (v1 >> ((846) & 63))}
                            var v31 uint8 = 113
                            _ = v30
                            _ = v31
                        } else {
                            v32 := GenS2{F0: v27, F1: v27, F2: v1}
                            _ = v32
                        }
                        _ = v28
                        _ = v29
                    } else {
                        var v33 int16 = v1
                        var v34 uint32 = 4294967295
                        switch (fi25 - fi25) {
                        case -1:
                            v5 = v12
                            var v35 float32 = v3
                            _ = v35
                        }
                        _ = v33
                        _ = v34
                    }
                }
                _ = fi25
                _ = v26
                _ = v27
            }
            _ = fi22
        }
    }
    if false {
        v12 = (v12 + v12)
        v36 := GenS2{F0: -9223372036854775808, F1: 9223372036854775807, F2: int16(v2)}
        _ = v36
    }
    v1 += (v1 + (v1 & 626))
    for fi37 := 0; fi37 < 3; fi37++ {
        switch ("a" + v12) {
        case "hello":
            v12 = v5
        case "\n\t":
            v4 = 119
        case "":
            if false {
                v38 := []int{(fi37 & fi37), int(v4)}
                v39 := GenS4{F0: 0}
                _ = v38
                _ = v39
            }
        default:
            var v40 string = v5
            _ = v40
        }
        _ = fi37
    }
    v1 += 1
    if (1 <= -1) {
        var v41 float64 = (float64(v3) / 1.5)
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
