package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 string; F2 int16 }
type GenS4 struct { F0 string }
func genFn0(p0 int, p1 int64, p2 int8) uint64 {
    var v4 int8 = p2
    v5 := GenS1{F0: -32768, F1: false}
    v4 += (v4 * p2)
    for fi6 := 0; fi6 < 9; fi6++ {
        v4 = p2
        v7 := GenS1{F0: 42, F1: (4294967295 >= 2)}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        if false {
            v9 := []int{fi8, (p0 + p0), 42}
            var v10 float64 = 0.0
            _ = v10
            _ = v9
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 4294967295
}
func genFn1() int8 {
    var v11 int = -818
    v11 = v11
    if (true && (0.0 <= 3.14159)) {
        for fi12 := 0; fi12 < 4; fi12++ {
            v13 := []int{-1, -1}
            fmt.Println(fi12, v11)
            _ = fi12
            _ = v13
        }
        v11 = 32767
        if (-1 >= -32768) {
            var v14 int32 = -230
            v15 := GenS1{F0: -32768, F1: (0 != 65535)}
            _ = v14
            _ = v15
        } else {
            v16 := GenS1{F0: 392, F1: false}
            _ = v16
        }
    } else {
        v17 := GenS2{F0: -794}
        _ = v17
    }
    _ = v11
    return 1
}

func main() {
    v1 := []uint32{0, 42}
    var v2 int8 = 100
    var v3 int8 = (v2 >> ((42) & 63))
    v18 := GenS1{F0: -590, F1: (false || false)}
    var v19 int32 = -422
    var v20 int8 = (v2 ^ (v2 & v3))
    if false {
        var v21 int32 = (v19 - v19)
        v22 := GenS3{F0: (v3 * v2), F1: "\n\t", F2: 0}
        v23 := GenS4{F0: "hello"}
        _ = v21
        _ = v22
        _ = v23
    }
    if ((-32768 > 1) || true) {
        v24 := GenS1{F0: -313, F1: false}
        v3 -= -91
        _ = v24
    } else {
        for fi25 := 0; fi25 < 3; fi25++ {
            if (0 == 0) {
                var v26 uint8 = 230
                var v27 int64 = 0
                _ = v26
                _ = v27
            } else {
                v20 += (v20 << ((255) & 63))
                v19 += v19
                var v28 uint = 52
                _ = v28
            }
            var v29 int32 = v19
            _ = fi25
            _ = v29
        }
        v30 := GenS3{F0: v2, F1: "a", F2: 1000}
        v1 = append(v1, 2)
        _ = v30
    }
    switch 9223372036854775807 {
    case 2147483647:
        fmt.Println(v20, v20)
    case 4294967295:
        fmt.Println(v20, v3, v3)
    }
    var v31 uint64 = 0
    if (!false) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi32 := 0; fi32 < 9; fi32++ {
            switch (^(fi32)) {
            case -9223372036854775808:
                var v33 uint16 = 255
                _ = v33
            case 1:
                var v34 uint16 = 42
                _ = v34
            default:
                if (true && true) {
                    switch ("" + "\n\t") {
                    case "a":
                        v3 = (v3 << ((4294967295) & 63))
                    }
                    if (!true) {
                        v3 = v3
                    }
                } else {
                    v35 := GenS2{F0: 1000}
                    for fi36 := 0; fi36 < 4; fi36++ {
                        v2 = v3
                        fmt.Println(v2, v19)
                        _ = fi36
                    }
                    _ = v35
                }
            }
            v1 = append(v1, uint32(v31))
            _ = fi32
        }
        var v37 int8 = (v20 / ((v2) | 1))
        _ = v37
    }
    v31 -= v31
    var v38 float64 = -2.25
    for fi39 := 0; fi39 < 11; fi39++ {
        v2 += 0
        fmt.Println(fi39)
        fmt.Println(v38, v20, v19)
        _ = fi39
    }
    v2 = (v2 - (v3 & int8(v31)))
    if (genFn0(-888, 280, v2) > (^(v31))) {
        v40 := GenS4{F0: ("" + "")}
        _ = v40
    }
    var v41 uint64 = v31
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v38
    _ = v41
}
