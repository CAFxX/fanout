package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 int32; F2 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint; F1 uint8; F2 uint8 }
type GenS6 struct { F0 int32; F1 string }
func genFn0() uint {
    if (!false) {
        for fi7 := 0; fi7 < 14; fi7++ {
            if (3.14159 <= 3.14159) {
                if (false || (1000 >= -445)) {
                    v8 := []int{fi7, (fi7 + (fi7 - fi7)), (fi7 % ((fi7) | 1))}
                    if false {
                        var v9 int8 = 0
                        _ = v9
                    }
                    var v10 bool = true
                    _ = v10
                    _ = v8
                } else {
                    var v11 int64 = 1
                    var v12 string = ""
                    switch v12 {
                    case "foo bar":
                        for fi13 := 0; fi13 < 4; fi13++ {
                            v14 := []int{238}
                            var v15 uint64 = 1
                            fmt.Println(v12, v15, v12)
                            var v16 int64 = (^(v11))
                            _ = fi13
                            _ = v14
                            _ = v15
                            _ = v16
                        }
                    }
                    _ = v11
                    _ = v12
                }
                var v17 int32 = 0
                var v18 float64 = 3.14159
                _ = v17
                _ = v18
            } else {
                v19 := []int{(fi7 << ((457) & 63)), 32768}
                v20 := GenS1{F0: 1, F1: "hello"}
                _ = v19
                _ = v20
            }
            v21 := GenS2{F0: 4294967295}
            var v22 int16 = 1
            _ = fi7
            _ = v21
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 9; fi23++ {
        for fi24 := 0; fi24 < 18; fi24++ {
            var v25 uint16 = 410
            v26 := GenS2{F0: uint(v25)}
            _ = fi24
            _ = v25
            _ = v26
        }
        switch (fi23 - fi23) {
        case -702:
            var v27 bool = (2 <= 4294967295)
            _ = v27
        }
        switch "\n\t" {
        case "":
            v28 := GenS3{F0: -1, F1: 1, F2: -1}
            _ = v28
        case "a":
            var v29 float64 = 1e10
            _ = v29
        case "foo bar":
            fmt.Println(fi23)
        }
        _ = fi23
    }
    v30 := GenS4{F0: ("foo bar" + "hello")}
    _ = v30
    return 18446744073709551615
}
func genFn1(p0 int16, p1 uint8, p2 uint) uint16 {
    for fi31 := 0; fi31 < 5; fi31++ {
        p2 = 1
        switch -9223372036854775808 {
        case -10:
            p0 += 0
        case -1:
            p2 -= p2
        case 0:
            v32 := GenS5{F0: (^(p2)), F1: 128, F2: p1}
            _ = v32
        }
        p0 -= -32768
        _ = fi31
    }
    var v33 int64 = 848
    fmt.Println(p2, p1, p2)
    var v34 int32 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v33
    _ = v34
    return 0
}

func main() {
    var v1 uint8 = 2
    var v2 uint32 = 49
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]int)
    var v5 float64 = ((3.14159 - -2.25) - 1e10)
    v6 := []uint32{(v2 / (((v2 >> ((18446744073709551615) & 63))) | 1)), (v2 << ((0) & 63)), v2, uint32(v1)}
    fmt.Println(v5)
    var v35 uint32 = 612
    if ("a" == "") {
        switch "\n\t" {
        case "foo bar":
            v36 := GenS4{F0: ("a" + "foo bar")}
            _ = v36
        case "hello":
            var v37 float32 = (float32(float64(0.1)) - ((((float32(float64(1.5)) * 1.5) * 1.5) + (1.5 * 1.5)) * 0.1))
            _ = v37
        }
        v38 := GenS2{F0: 65535}
        fmt.Println(v2, v5)
        _ = v38
    } else {
        v35 = uint32(v1)
    }
    v39 := GenS6{F0: 42, F1: "hello"}
    var v40 int32 = 809
    if (4294967295 >= 298) {
        if len(v6) > 0 {
            v6[0] = v35
        }
        var v41 int16 = 1000
        _ = v41
    }
    v1 += v1
    fmt.Println(v2)
    if true {
        var v42 bool = false
        v43 := GenS4{F0: ""}
        v44 := v4[4294967295]
        _ = v42
        _ = v43
        _ = v44
    } else {
        var v45 int16 = 42
        var v46 uint8 = v1
        switch -923 {
        case -9223372036854775808:
            for fi47 := 0; fi47 < 13; fi47++ {
                _ = len(v4)
                _ = fi47
            }
        case -724:
            v45 -= (v45 << ((42) & 63))
        case -1:
            if (0 <= -9223372036854775808) {
                v48 := GenS6{F0: (-(v40)), F1: ""}
                for fi49 := 0; fi49 < 10; fi49++ {
                    v35 += (v2 - v35)
                    _ = fi49
                }
                _ = v48
            }
        default:
            fmt.Println(v5, v1)
        }
        _ = v45
        _ = v46
    }
    v2 -= 174
    switch "a" {
    case "foo bar":
        v50 := v4[42]
        _ = v50
    case "hello":
        var v51 int64 = 9223372036854775807
        _ = v51
    case "\n\t":
        v40 = (v40 + 42)
    default:
        v3[(^(v35))] = (v35 ^ v2)
    }
    var v52 uint64 = uint64(v35)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v52
    _ = v6
}
