package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 uint64; F2 int8 }
func genFn0() uint64 {
    for fi7 := 0; fi7 < 20; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            switch "foo bar" {
            case "a":
                fmt.Println(fi8)
            case "hello":
                var v9 int32 = 1
                _ = v9
            default:
                v10 := []int{42, -108}
                _ = v10
            }
            var v11 int16 = -1
            _ = fi8
            _ = v11
        }
        _ = fi7
    }
    var v12 uint32 = 4294967295
    v13 := []int{}
    _ = v12
    _ = v13
    return 2
}
func genFn1(p0 int64, p1 uint64, p2 int) int64 {
    if false {
        p1 += (p1 & genFn0())
        for fi14 := 0; fi14 < 9; fi14++ {
            v15 := GenS1{F0: 0, F1: p1, F2: 127}
            _ = fi14
            _ = v15
        }
        p2 = p2
    }
    for fi16 := 0; fi16 < 20; fi16++ {
        var v17 int = fi16
        v17 = (v17 << ((2) & 63))
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 20; fi18++ {
        var v19 float32 = 0.1
        v20 := GenS2{F0: 925}
        if false {
            if (!(v19 > (0.0 * (v19 + v19)))) {
                v21 := GenS2{F0: 42}
                var v22 int16 = 141
                _ = v21
                _ = v22
            }
        } else {
            v23 := []int{32767}
            fmt.Println(p0)
            _ = v23
        }
        _ = fi18
        _ = v19
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    return 430
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]uint32)
    v3 := []uint32{2147483648, 4294967295}
    var v4 int = -9223372036854775808
    var v5 float64 = float64((3.14159 / 3.14159))
    v6 := make(map[int]string)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v24 uint = 255
    for fi25 := 0; fi25 < 16; fi25++ {
        fmt.Println(v24, v4, v24)
        _ = fi25
    }
    var v26 int = -1
    fmt.Println(v26, v26, v24)
    if (true && (4294967295 != 758)) {
        if (!(!(!true))) {
            switch ("\n\t" + ("" + "a")) {
            case "\n\t":
                var v27 int = (v26 >> ((v24) & 63))
                _ = v27
            case "foo bar":
                var v28 float32 = (3.14159 / 1e10)
                _ = v28
            case "hello":
                v29 := GenS3{F0: v24, F1: 218, F2: 92}
                _ = v29
            default:
                switch (v4 << (((v24 | (v24 >> ((v24) & 63)))) & 63)) {
                case 127:
                    var v30 bool = (444 != -2147483648)
                    _ = v30
                }
            }
            var v31 uint8 = 175
            _ = v31
        } else {
            v24 = v24
        }
    } else {
        v32 := GenS3{F0: v24, F1: 42, F2: 127}
        _ = v32
    }
    v26 = v4
    v33 := GenS1{F0: 42, F1: genFn0(), F2: v4}
    if false {
        v5 = v5
    } else {
        v4 = v26
    }
    var v34 uint8 = 151
    fmt.Println(v4, v24)
    var v35 uint16 = 32767
    switch (v26 ^ (-(v26))) {
    case 128:
        delete(v6, 9223372036854775807)
    default:
        var v36 uint32 = 0
        _ = v36
    }
    for fi37 := 0; fi37 < 5; fi37++ {
        var v38 int32 = int32(v34)
        _ = fi37
        _ = v38
    }
    if false {
        var v39 int64 = -1
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
