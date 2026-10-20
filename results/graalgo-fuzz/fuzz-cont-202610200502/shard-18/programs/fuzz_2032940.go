package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
func genFn0(p0 int8, p1 uint32, p2 bool) int {
    if (696 <= uint64(p1)) {
        fmt.Println(p0, p2)
        if false {
            v7 := []int{-2147483648, -9223372036854775808}
            switch "\n\t" {
            case "":
                v8 := GenS1{F0: 100, F1: 413}
                _ = v8
            case "a":
                v9 := GenS1{F0: p0, F1: 9223372036854775807}
                _ = v9
            default:
                var v10 uint16 = 944
                _ = v10
            }
            _ = v7
        }
        v11 := GenS1{F0: p0, F1: -9223372036854775808}
        _ = v11
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        for fi13 := 0; fi13 < 12; fi13++ {
            var v14 uint16 = 209
            var v15 int = fi12
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi12
    }
    p1 = p1
    var v16 int16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 32768
}
func genFn1(p0 uint32, p1 uint64) uint16 {
    p1 -= p1
    if ((2 == 127) && false) {
        p1 -= 422
        var v17 uint8 = 42
        for fi18 := 0; fi18 < 2; fi18++ {
            v17 = (^(v17))
            p1 = p1
            var v19 string = "foo bar"
            _ = fi18
            _ = v19
        }
        _ = v17
    } else {
        switch 128 {
        case 251:
            var v20 uint = 266
            _ = v20
        case 931:
            if (!true) {
                fmt.Println(p0, p1)
                var v21 int64 = 1
                var v22 uint = 136
                _ = v21
                _ = v22
            }
        default:
            for fi23 := 0; fi23 < 16; fi23++ {
                v24 := GenS1{F0: 1, F1: 9223372036854775807}
                switch ("" + "foo bar") {
                case "hello":
                    var v25 uint = 33
                    _ = v25
                case "\n\t":
                    for fi26 := 0; fi26 < 19; fi26++ {
                        v27 := GenS1{F0: 127, F1: 1}
                        var v28 uint16 = 0
                        if ((fi26 <= (fi26 >> ((0) & 63))) || false) {
                            var v29 string = "hello"
                            _ = v29
                        }
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                }
                v30 := GenS1{F0: -1, F1: 0}
                _ = fi23
                _ = v24
                _ = v30
            }
        }
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    return 812
}

func main() {
    var v1 string = "\n\t"
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]bool)
    var v4 int = 127
    v5 := make(map[uint32]uint32)
    v6 := make(map[uint32]uint32)
    var v31 uint8 = 137
    v4 = 757
    for fi32 := 0; fi32 < 12; fi32++ {
        switch v4 {
        case 1:
            v1 = v1
        }
        _ = fi32
    }
    v4 = (v4 << ((2) & 63))
    for fi33 := 0; fi33 < 14; fi33++ {
        var v34 float32 = 0.0
        fmt.Println(v4)
        var v35 int64 = 42
        _ = fi33
        _ = v34
        _ = v35
    }
    for fi36 := 0; fi36 < 7; fi36++ {
        v31 = (v31 - (v31 >> ((uint(v31)) & 63)))
        var v37 int = genFn0(-1, 1, (895 == 762))
        _ = fi36
        _ = v37
    }
    switch (v4 % ((v4) | 1)) {
    case 32768:
        switch 1 {
        case 32768:
            v38 := GenS1{F0: 127, F1: 0}
            _ = v38
        }
    case 1:
        v39 := []int{v4}
        _ = v39
    case -776:
        v4 += (v4 % ((0) | 1))
    }
    fmt.Println(v31, v31)
    fmt.Println(v4)
    var v40 uint32 = 4294967295
    for fi41 := 0; fi41 < 6; fi41++ {
        var v42 float32 = 0.1
        var v43 bool = (v40 >= v40)
        if (370 > 0) {
            v1 = (v1 + "a")
        }
        _ = fi41
        _ = v42
        _ = v43
    }
    var v44 uint64 = 42
    delete(v5, v40)
    var v45 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v40
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
