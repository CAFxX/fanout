package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 int16; F2 int32 }
func genFn0(p0 int64, p1 int64) int32 {
    var v5 bool = true
    if true {
        v6 := GenS1{F0: 1}
        for fi7 := 0; fi7 < 11; fi7++ {
            for fi8 := 0; fi8 < 6; fi8++ {
                v5 = (!(65535 <= 4294967295))
                if v5 {
                    var v9 string = "hello"
                    v10 := []int{(fi8 | (fi8 & (fi8 >> ((42) & 63)))), (fi7 | fi7), (fi7 >> ((65535) & 63)), -2147483648}
                    switch v9 {
                    case "\n\t":
                        var v11 bool = false
                        _ = v11
                    case "foo bar":
                        p1 += p0
                    }
                    _ = v10
                    _ = v9
                } else {
                    fmt.Println(p1, fi8, fi8)
                    v12 := []int{(fi7 >> ((255) & 63))}
                    var v13 uint64 = 2
                    _ = v12
                    _ = v13
                }
                p0 -= p1
                _ = fi8
            }
            _ = fi7
        }
        p0 += p0
        _ = v6
    } else {
        var v14 int32 = -1
        var v15 uint8 = 127
        var v16 uint32 = 601
        _ = v14
        _ = v15
        _ = v16
    }
    if v5 {
        if (1000 >= 674) {
            v17 := GenS1{F0: -1}
            switch ("a" + "\n\t") {
            case "\n\t":
                fmt.Println(p0)
            case "a":
                for fi18 := 0; fi18 < 1; fi18++ {
                    v19 := []int{(fi18 << ((2) & 63))}
                    _ = fi18
                    _ = v19
                }
            case "":
                fmt.Println(p0, p1, v5)
            default:
                var v20 uint16 = 666
                _ = v20
            }
            var v21 int8 = 127
            _ = v17
            _ = v21
        }
    }
    if (v5 && false) {
        fmt.Println(p1, p0, v5)
        v5 = v5
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v5
    return 42
}

func main() {
    var v1 uint = 42
    var v2 bool = true
    var v3 uint16 = 2
    v4 := make(map[uint32]int)
    fmt.Println(v3)
    for fi22 := 0; fi22 < 1; fi22++ {
        switch "hello" {
        case "":
            if v2 {
                switch ("" + "a") {
                case "":
                    fmt.Println(fi22, v3)
                }
                if (-1 < -9223372036854775808) {
                    v23 := GenS1{F0: 32767}
                    v24 := GenS2{F0: 150}
                    var v25 uint16 = (^(v3))
                    _ = v23
                    _ = v24
                    _ = v25
                } else {
                    v1 -= (v1 * v1)
                    var v26 uint8 = 255
                    if (!false) {
                        v2 = (!v2)
                    }
                    _ = v26
                }
            }
        case "\n\t":
            var v27 uint64 = 4294967295
            _ = v27
        case "hello":
            v28 := []int{fi22, (fi22 / ((fi22) | 1)), fi22, fi22}
            _ = v28
        }
        fmt.Println(v2)
        var v29 int = 523
        _ = fi22
        _ = v29
    }
    if (1e10 > 0.0) {
        if v2 {
            var v30 float32 = 1.5
            fmt.Println(v1, v30, v2)
            _ = v30
        }
        fmt.Println(v3, v1, v3)
        v31 := GenS3{F0: 877, F1: 0, F2: 42}
        _ = v31
    } else {
        v32 := GenS3{F0: 2147483647, F1: 207, F2: -2147483648}
        var v33 float32 = float32(0.0)
        var v34 string = "\n\t"
        _ = v32
        _ = v33
        _ = v34
    }
    v4[746] = 127
    _ = len(v4)
    v3 -= (v3 | v3)
    v3 = (v3 - 255)
    var v35 bool = false
    for fi36 := 0; fi36 < 8; fi36++ {
        var v37 bool = false
        _ = fi36
        _ = v37
    }
    switch 9223372036854775807 {
    case 4294967295:
        if ((!false) && (!v2)) {
            v38 := []int{2147483647, -2147483648}
            v35 = v35
            _ = v38
        } else {
            fmt.Println(v2, v2)
        }
    case 42:
        for fi39 := 0; fi39 < 3; fi39++ {
            v1 = (v1 + (v1 / (((v1 - (v1 ^ 397))) | 1)))
            for fi40 := 0; fi40 < 3; fi40++ {
                v3 -= v3
                _ = fi40
            }
            _ = fi39
        }
    case 567:
        for fi41 := 0; fi41 < 16; fi41++ {
            v42 := GenS1{F0: 680}
            var v43 uint32 = 2147483647
            _ = fi41
            _ = v42
            _ = v43
        }
    default:
        var v44 bool = ((!(!v2)) && v2)
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
}
