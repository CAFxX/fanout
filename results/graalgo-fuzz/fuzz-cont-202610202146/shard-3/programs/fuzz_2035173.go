package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int16; F1 int8; F2 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 int64) uint32 {
    if true {
        switch 858 {
        case 32767:
            v4 := GenS1{F0: 65535, F1: 78}
            _ = v4
        }
        if true {
            p0 = p0
            v5 := GenS1{F0: 65535, F1: 42}
            _ = v5
        }
        v6 := []int{310}
        _ = v6
    }
    p0 = p0
    v7 := GenS1{F0: 1, F1: 902}
    switch 813 {
    case 127:
        p0 = 9223372036854775807
    case -9223372036854775808:
        switch "" {
        case "hello":
            fmt.Println(p0)
        case "foo bar":
            if (42 == (p0 >> ((42) & 63))) {
                var v8 string = ("foo bar" + "foo bar")
                v8 = v8
                var v9 int = 657
                _ = v8
                _ = v9
            } else {
                if false {
                    v10 := []int{-2147483648}
                    _ = v10
                } else {
                    v11 := GenS1{F0: 65535, F1: 883}
                    var v12 float64 = (3.14159 - 1.5)
                    _ = v11
                    _ = v12
                }
                v13 := GenS1{F0: 65535, F1: 18446744073709551615}
                _ = v13
            }
        case "a":
            var v14 uint64 = 2
            _ = v14
        default:
            if true {
                p0 = p0
            }
        }
    case 2147483647:
        v15 := GenS1{F0: 42, F1: 33}
        _ = v15
    }
    _ = p0
    _ = v7
    return 895
}
func genFn1() int8 {
    v16 := []int{-9223372036854775808, -601}
    var v17 uint64 = 42
    for fi18 := 0; fi18 < 4; fi18++ {
        v19 := GenS2{F0: 1000, F1: 111, F2: (!(128 < 127))}
        for fi20 := 0; fi20 < 20; fi20++ {
            switch ("hello" + ("a" + "\n\t")) {
            case "foo bar":
                var v21 float32 = (1e10 / (0.0 - (0.1 + 0.0)))
                _ = v21
            case "a":
                fmt.Println(fi18, fi18, fi20)
            }
            _ = fi20
        }
        _ = fi18
        _ = v19
    }
    var v22 bool = (-2147483648 != -1)
    v23 := GenS3{F0: (v17 & (v17 | 2))}
    _ = v16
    _ = v17
    _ = v22
    _ = v23
    return 100
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 13
    v3 := make(map[string]bool)
    v24 := []int{690}
    if true {
        if true {
            var v25 uint16 = 811
            _ = v25
        } else {
            v2 -= (v2 & (^(v2)))
            if len(v24) >= 2 {
                v24 = v24[1:len(v24)]
            }
        }
        v24 = append(v24, 32767)
        v2 -= (v2 >> ((1) & 63))
    }
    if (2 >= 1) {
        v2 = (^(v2))
        v2 += (v2 & (^(v2)))
        if false {
            for fi26 := 0; fi26 < 8; fi26++ {
                switch (fi26 ^ 9223372036854775807) {
                case 4294967295:
                    var v27 string = "foo bar"
                    _ = v27
                case -439:
                    v28 := GenS3{F0: 2}
                    _ = v28
                default:
                    switch (fi26 / ((-1) | 1)) {
                    case -9223372036854775808:
                        v24 = append(v24, (-(fi26)))
                    default:
                        if true {
                            v29 := GenS4{F0: 1, F1: 1}
                            _ = v29
                        }
                    }
                }
                if true {
                    v30 := GenS4{F0: 184, F1: 174}
                    var v31 int32 = -2147483648
                    _ = v30
                    _ = v31
                } else {
                    v24 = append(v24, fi26)
                    var v32 float32 = 1e10
                    v33 := GenS5{F0: 683}
                    _ = v32
                    _ = v33
                }
                _ = fi26
            }
        }
    } else {
        v2 += (v2 ^ (v2 - 42))
        v2 -= (v2 & (^(v2)))
    }
    var v34 int8 = -1
    v2 = (^(v2))
    switch 42 {
    case -739:
        v35 := GenS5{F0: 18446744073709551615}
        _ = v35
    default:
        if (0.1 != -2.25) {
            for fi36 := 0; fi36 < 18; fi36++ {
                v2 -= (^(v2))
                _ = fi36
            }
            var v37 uint32 = genFn0(42)
            for fi38 := 0; fi38 < 3; fi38++ {
                v34 = (v34 ^ -128)
                _ = fi38
            }
            _ = v37
        }
    }
    var v39 int8 = 1
    v34 += 1
    for fi40 := 0; fi40 < 8; fi40++ {
        v41 := GenS1{F0: 1, F1: 0}
        if (true && false) {
            if len(v24) > 0 {
                v24[0] = (^(fi40))
            }
            _ = len(v24)
            for fi42 := 0; fi42 < 1; fi42++ {
                switch fi42 {
                case 4294967295:
                    v34 = (-(v39))
                case 1:
                    if false {
                        v2 -= genFn0(-1)
                    }
                }
                _ = fi42
            }
        }
        if (57 > 70) {
            var v43 uint64 = 409
            v34 = genFn1()
            _ = v43
        }
        _ = fi40
        _ = v41
    }
    v39 = -63
    for fi44 := 0; fi44 < 15; fi44++ {
        var v45 int = fi44
        v39 = (v39 | v34)
        fmt.Println(fi44, v39)
        _ = fi44
        _ = v45
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v34
    _ = v39
}
