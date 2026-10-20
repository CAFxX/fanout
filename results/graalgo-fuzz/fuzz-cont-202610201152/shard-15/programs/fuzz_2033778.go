package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int32 }
type GenS2 struct { F0 string; F1 uint32; F2 uint64 }
type GenS3 struct { F0 int32; F1 uint32; F2 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8; F1 int16; F2 uint64 }
func genFn0() int16 {
    if (!false) {
        var v5 int16 = -974
        _ = v5
    }
    if (42 != -583) {
        if (!false) {
            v6 := GenS1{F0: 42, F1: -50, F2: -1}
            _ = v6
        } else {
            var v7 string = "foo bar"
            if false {
                var v8 uint32 = 2147483647
                v9 := []int{42, 772, 1}
                _ = v8
                _ = v9
            }
            v10 := GenS1{F0: -9223372036854775808, F1: -119, F2: -2147483648}
            _ = v10
            _ = v7
        }
    } else {
        var v11 int16 = 0
        _ = v11
    }
    var v12 string = "hello"
    switch 263 {
    case -1:
        v12 = ((v12 + (v12 + "hello")) + "foo bar")
    case 2147483647:
        v12 = v12
    default:
        v12 = v12
    }
    _ = v12
    return 1000
}
func genFn1(p0 string, p1 int8, p2 uint64) bool {
    switch p0 {
    case "hello":
        p1 -= (p1 % ((p1) | 1))
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        switch (fi13 * (fi13 >> ((18446744073709551615) & 63))) {
        case 2147483647:
            v14 := GenS1{F0: 0, F1: 127, F2: 479}
            _ = v14
        case 325:
            p0 = "a"
        case 4294967295:
            switch (fi13 % ((-2147483648) | 1)) {
            case 32768:
                for fi15 := 0; fi15 < 4; fi15++ {
                    v16 := GenS2{F0: "hello", F1: 2147483647, F2: 1}
                    if (!(!(18446744073709551615 < 18446744073709551615))) {
                        var v17 int32 = -2147483648
                        switch p0 {
                        case "hello":
                            p1 -= (^(p1))
                        case "foo bar":
                            v18 := GenS2{F0: "", F1: uint32(p2), F2: p2}
                            _ = v18
                        }
                        v19 := []int{(fi15 / ((fi13) | 1))}
                        _ = v17
                        _ = v19
                    }
                    fmt.Println(p2, p0, fi13)
                    _ = fi15
                    _ = v16
                }
            default:
                var v20 uint32 = 234
                _ = v20
            }
        }
        if (1 < 306) {
            fmt.Println(p1)
            p1 = p1
        }
        _ = fi13
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        p2 = p2
        for fi22 := 0; fi22 < 19; fi22++ {
            v23 := GenS3{F0: 0, F1: 42, F2: "a"}
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    var v24 int16 = 0
    for fi25 := 0; fi25 < 6; fi25++ {
        p1 -= 102
        for fi26 := 0; fi26 < 13; fi26++ {
            fmt.Println(v24, p2, v24)
            p0 = p0
            fmt.Println(v24, v24)
            _ = fi26
        }
        _ = fi25
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    return false
}

func main() {
    var v1 int = -9223372036854775808
    v2 := make(map[string]string)
    var v3 uint16 = 0
    var v4 uint64 = 1
    v1 -= v1
    for fi27 := 0; fi27 < 12; fi27++ {
        v28 := GenS2{F0: "", F1: 2147483648, F2: (v4 << ((255) & 63))}
        var v29 uint32 = 4294967295
        v1 = fi27
        _ = fi27
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 5; fi30++ {
        var v31 uint8 = 2
        _ = fi30
        _ = v31
    }
    switch (v1 / (((v1 / ((v1) | 1))) | 1)) {
    case 1:
        v32 := GenS4{F0: 918}
        _ = v32
    case 32767:
        v4 += 1
    }
    if (v3 < 1) {
        for fi33 := 0; fi33 < 10; fi33++ {
            fmt.Println(fi33, v4)
            if (4294967295 > v4) {
                v3 -= 467
                var v34 uint32 = 705
                v34 -= v34
                _ = v34
            }
            _ = fi33
        }
        var v35 int = v1
        _ = v35
    }
    var v36 int32 = -2147483648
    v4 -= 1
    if true {
        delete(v2, "")
        var v37 uint32 = 261
        v1 -= v1
        _ = v37
    }
    switch v1 {
    case 2147483647:
        if (((v4 ^ (v4 % ((v4) | 1))) == 42) && (586 == 65535)) {
            var v38 uint32 = 2147483647
            v3 = (v3 << ((4294967295) & 63))
            _ = v38
        }
    case 32767:
        var v39 int16 = 42
        _ = v39
    case -820:
        v40 := GenS2{F0: ("hello" + ""), F1: 27, F2: (v4 ^ v4)}
        _ = v40
    default:
        var v41 int64 = -1
        _ = v41
    }
    v42 := GenS2{F0: "foo bar", F1: 2147483647, F2: 0}
    fmt.Println(v4, v3)
    if true {
        var v43 int = v1
        delete(v2, "a")
        v44 := GenS5{F0: 0, F1: -1, F2: (^(v4))}
        _ = v43
        _ = v44
    } else {
        v1 += v1
        v4 += 4294967295
    }
    v45 := GenS1{F0: int64(v4), F1: -128, F2: (v36 | int32(v4))}
    v46 := []int{v1, v1}
    fmt.Println(v36, v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v46:", len(v46), v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v42
    _ = v45
    _ = v46
}
