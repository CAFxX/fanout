package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 int; F2 bool }
func genFn0(p0 int, p1 uint8) uint8 {
    for fi6 := 0; fi6 < 2; fi6++ {
        p1 = 2
        p0 = 247
        var v7 int32 = 2147483647
        _ = fi6
        _ = v7
    }
    if false {
        p1 -= (^(p1))
        p0 += p0
        if (0.1 == (-2.25 * (0.1 * float32(-2.25)))) {
            p1 -= (p1 & 106)
            fmt.Println(p1)
            fmt.Println(p0)
        }
    } else {
        switch ((("foo bar" + "foo bar") + "\n\t") + "") {
        case "foo bar":
            switch (^(p0)) {
            case 128:
                for fi8 := 0; fi8 < 5; fi8++ {
                    var v9 float32 = 3.14159
                    _ = fi8
                    _ = v9
                }
            default:
                fmt.Println(p0, p1)
            }
        case "hello":
            if false {
                var v10 int8 = 42
                _ = v10
            }
        case "":
            if false {
                var v11 uint32 = 2147483648
                _ = v11
            } else {
                fmt.Println(p1, p0)
                p0 = 32768
                var v12 bool = ((1 < 102) || true)
                _ = v12
            }
        }
    }
    p1 = p1
    var v13 int16 = 32767
    for fi14 := 0; fi14 < 4; fi14++ {
        v15 := []int{}
        _ = fi14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v13
    return (p1 % ((p1) | 1))
}
func genFn1() uint8 {
    for fi16 := 0; fi16 < 17; fi16++ {
        var v17 int64 = -677
        _ = fi16
        _ = v17
    }
    var v18 uint32 = 0
    var v19 uint32 = v18
    for fi20 := 0; fi20 < 19; fi20++ {
        v19 += (v18 << ((0) & 63))
        _ = fi20
    }
    _ = v18
    _ = v19
    return 2
}

func main() {
    var v1 int16 = -54
    var v2 int64 = 52
    v3 := make(map[uint32]int)
    v4 := make(map[int]bool)
    var v5 uint8 = 255
    v21 := GenS1{F0: genFn0(-862, v5), F1: 622}
    v2 -= v2
    v22 := GenS2{F0: 32767}
    v5 = 24
    for fi23 := 0; fi23 < 13; fi23++ {
        var v24 uint32 = 2
        _ = fi23
        _ = v24
    }
    v25 := []int{42}
    v26 := GenS2{F0: -92}
    switch 42 {
    case -1:
        delete(v4, 32767)
    case -289:
        for fi27 := 0; fi27 < 1; fi27++ {
            var v28 int16 = (^(v1))
            switch "" {
            case "hello":
                if true {
                    v29 := GenS2{F0: fi27}
                    v1 = v1
                    _ = v29
                } else {
                    var v30 int = -2147483648
                    for fi31 := 0; fi31 < 17; fi31++ {
                        fmt.Println(v5, fi31)
                        v30 += (-(v30))
                        _ = fi31
                    }
                    _ = v30
                }
            }
            _ = fi27
            _ = v28
        }
    }
    v3[2147483648] = -667
    var v32 uint32 = 785
    if len(v25) > 0 {
        v25[0] = 32767
    }
    if (42 == -9223372036854775808) {
        fmt.Println(v2, v1, v1)
        v33 := GenS3{F0: 18446744073709551615, F1: 43, F2: false}
        _ = v33
    }
    fmt.Println(v2, v32)
    if len(v25) > 0 {
        v25[0] = 2147483647
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
