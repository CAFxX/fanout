package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 uint32; F1 int64; F2 int32 }

func main() {
    var v1 float64 = -2.25
    var v2 uint64 = 190
    v3 := make(map[int]bool)
    var v4 bool = ((false && true) && true)
    var v5 uint8 = 14
    var v6 uint8 = (v5 + 2)
    v7 := GenS1{F0: 2, F1: 347}
    for fi8 := 0; fi8 < 10; fi8++ {
        v9 := []int{}
        _ = fi8
        _ = v9
    }
    switch "" {
    case "":
        if (false && v4) {
            v6 += v5
            v10 := v3[-240]
            v11 := GenS1{F0: 65535, F1: 0}
            _ = v10
            _ = v11
        }
    case "a":
        v12 := []int{}
        _ = v12
    case "\n\t":
        switch 9223372036854775807 {
        case 286:
            var v13 uint16 = 0
            _ = v13
        case 2147483647:
            v14 := v3[127]
            _ = v14
        case -9223372036854775808:
            fmt.Println(v4)
        }
    default:
        var v15 bool = v4
        _ = v15
    }
    var v16 bool = ((4294967295 < 234) || v4)
    var v17 uint64 = (v2 + uint64(v5))
    var v18 int32 = -2147483648
    var v19 string = ("foo bar" + "a")
    v20 := []int{-328, 127, -949, 4294967295}
    if (-585 < 0) {
        v1 = 0.0
        switch int(v18) {
        case 9223372036854775807:
            if len(v20) > 0 {
                v20[0] = -9223372036854775808
            }
        default:
            v1 = v1
        }
        if (v16 || v16) {
            var v21 float32 = -2.25
            _ = v21
        } else {
            v17 -= v2
            v5 = (v6 | v6)
            fmt.Println(v6, v4)
        }
    }
    if ((v16 && true) || v16) {
        var v22 uint32 = 2
        v23 := GenS2{F0: 4294967295, F1: 1}
        v3[32767] = v16
        _ = v22
        _ = v23
    } else {
        if (-1 > -1) {
            v19 = (v19 + "hello")
            var v24 float32 = 1e10
            for fi25 := 0; fi25 < 12; fi25++ {
                if v16 {
                    v26 := v3[(fi25 | (-(fi25)))]
                    for fi27 := 0; fi27 < 15; fi27++ {
                        var v28 uint64 = (v2 << ((uint(v6)) & 63))
                        v26 = (!v4)
                        v29 := GenS2{F0: 1, F1: 945}
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                    _ = v26
                }
                var v30 uint16 = 42
                _ = fi25
                _ = v30
            }
            _ = v24
        }
        v31 := GenS1{F0: uint16(v5), F1: 42}
        _ = v31
    }
    if ((0 < 973) && ((v18 % ((0) | 1)) > v18)) {
        for fi32 := 0; fi32 < 14; fi32++ {
            _ = len(v20)
            _ = len(v3)
            _ = fi32
        }
        if v4 {
            var v33 float32 = -2.25
            var v34 int8 = 1
            _ = v33
            _ = v34
        } else {
            switch ((v19 + "\n\t") + (v19 + v19)) {
            case "":
                fmt.Println(v6, v16)
            case "foo bar":
                if (436 != 42) {
                    var v35 int = 9223372036854775807
                    v18 = int32(v35)
                    switch v19 {
                    case "\n\t":
                        switch (v35 % (((v35 / (((v35 & v35)) | 1))) | 1)) {
                        case 9223372036854775807:
                            var v36 bool = (true || v16)
                            _ = v36
                        }
                    case "hello":
                        var v37 uint64 = (v2 >> ((2) & 63))
                        _ = v37
                    case "foo bar":
                        var v38 int8 = int8(v2)
                        _ = v38
                    }
                    _ = v35
                }
            }
        }
        v20 = append(v20, 32767)
    }
    switch v19 {
    case "\n\t":
        var v39 bool = (true && (!false))
        _ = v39
    case "a":
        v40 := GenS3{F0: 42, F1: 9223372036854775807, F2: (^(v18))}
        _ = v40
    case "":
        if true {
            delete(v3, -2147483648)
        }
    default:
        v20 = append(v20, 212)
    }
    for fi41 := 0; fi41 < 4; fi41++ {
        switch "hello" {
        case "":
            for fi42 := 0; fi42 < 10; fi42++ {
                v20 = append(v20, fi41)
                fmt.Println(v16)
                v1 += v1
                _ = fi42
            }
        default:
            var v43 int64 = 564
            _ = v43
        }
        var v44 uint64 = v2
        _ = fi41
        _ = v44
    }
    v45 := GenS3{F0: 1, F1: -9223372036854775808, F2: (v18 * (v18 / (((v18 + (v18 & -1))) | 1)))}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
    _ = v7
}
