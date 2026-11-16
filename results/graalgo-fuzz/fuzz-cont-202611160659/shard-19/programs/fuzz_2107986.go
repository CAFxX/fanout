package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 string; F1 uint }
type GenS6 struct { F0 uint64; F1 uint; F2 uint }
func genFn0(p0 int8, p1 int, p2 int8) uint64 {
    var v4 uint8 = 42
    for fi5 := 0; fi5 < 19; fi5++ {
        switch (("a" + "") + "\n\t") {
        case "":
            var v6 int32 = 256
            _ = v6
        case "a":
            switch "\n\t" {
            case "hello":
                var v7 uint16 = 1
                _ = v7
            case "":
                fmt.Println(p0)
            }
        case "hello":
            p2 += p0
        default:
            switch -2147483648 {
            case 239:
                p0 = (p2 >> ((4294967295) & 63))
            case 32768:
                v8 := GenS1{F0: 742}
                _ = v8
            default:
                p1 += -2147483648
            }
        }
        _ = fi5
    }
    if false {
        switch "" {
        case "":
            if (42 < 255) {
                var v9 float64 = -2.25
                v10 := GenS2{F0: 42, F1: (-(p0))}
                switch "foo bar" {
                case "":
                    p2 -= p0
                case "hello":
                    v11 := []int{p1}
                    _ = v11
                default:
                    v12 := GenS3{F0: -1, F1: 802}
                    _ = v12
                }
                _ = v10
                _ = v9
            } else {
                p1 += 127
                var v13 uint16 = 65535
                if false {
                    for fi14 := 0; fi14 < 10; fi14++ {
                        fmt.Println(p2, v13, p2)
                        _ = fi14
                    }
                    v4 = (v4 << ((2) & 63))
                } else {
                    var v15 uint = 42
                    _ = v15
                }
                _ = v13
            }
        case "foo bar":
            v16 := GenS1{F0: 2147483648}
            _ = v16
        case "hello":
            var v17 int8 = (p0 + p2)
            _ = v17
        default:
            if ((^(p1)) >= p1) {
                var v18 float64 = 3.14159
                _ = v18
            }
        }
        for fi19 := 0; fi19 < 13; fi19++ {
            v4 = v4
            fmt.Println(p1, fi19)
            _ = fi19
        }
    }
    v20 := GenS4{F0: 127, F1: (-(p0))}
    var v21 uint8 = (v4 << ((0) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    _ = v4
    return 1
}
func genFn1(p0 string) int32 {
    p0 = ("" + "")
    v22 := GenS5{F0: p0, F1: 0}
    if false {
        v23 := []int{-198, 42}
        p0 = (p0 + p0)
        if ((181 >= 926) || false) {
            v24 := GenS5{F0: p0, F1: 18446744073709551615}
            switch 0 {
            case 127:
                for fi25 := 0; fi25 < 5; fi25++ {
                    fmt.Println(p0, p0)
                    v23 = append(v23, (fi25 % ((567) | 1)))
                    _ = fi25
                }
            case -1:
                for fi26 := 0; fi26 < 13; fi26++ {
                    fmt.Println(fi26)
                    if (!true) {
                        var v27 uint8 = 0
                        _ = v27
                    }
                    var v28 int16 = 1
                    _ = fi26
                    _ = v28
                }
            default:
                if len(v23) > 0 {
                    v23[0] = 42
                }
            }
            fmt.Println(p0)
            _ = v24
        }
        _ = v23
    }
    v29 := []int{-1}
    _ = p0
    _ = v22
    _ = v29
    return 1
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]bool)
    var v3 int = 9223372036854775807
    if true {
        var v30 bool = ((false && true) || false)
        var v31 int32 = genFn1(("foo bar" + ("foo bar" + "a")))
        var v32 uint16 = 32767
        _ = v30
        _ = v31
        _ = v32
    }
    switch v3 {
    case 634:
        v33 := []int{v3, 127}
        _ = v33
    case -9223372036854775808:
        for fi34 := 0; fi34 < 20; fi34++ {
            for fi35 := 0; fi35 < 2; fi35++ {
                switch ("\n\t" + ("a" + ("\n\t" + "\n\t"))) {
                case "":
                    switch "\n\t" {
                    case "hello":
                        var v36 int32 = 2147483647
                        _ = v36
                    case "foo bar":
                        v37 := GenS1{F0: 1}
                        _ = v37
                    }
                }
                switch 4294967295 {
                case 4294967295:
                    if false {
                        var v38 uint32 = 0
                        v38 += v38
                        _ = v38
                    }
                case 1:
                    v39 := GenS3{F0: 0, F1: 42}
                    _ = v39
                }
                _ = fi35
            }
            _ = fi34
        }
    case 9223372036854775807:
        for fi40 := 0; fi40 < 4; fi40++ {
            v41 := GenS5{F0: "\n\t", F1: 0}
            _ = fi40
            _ = v41
        }
    }
    v42 := []int{v3, v3, (^(v3))}
    if false {
        var v43 uint64 = 1
        _ = v43
    }
    v44 := GenS4{F0: (v3 | v3), F1: 127}
    v45 := GenS6{F0: 2, F1: 42, F2: 4294967295}
    for fi46 := 0; fi46 < 4; fi46++ {
        var v47 int64 = -723
        _ = len(v42)
        _ = fi46
        _ = v47
    }
    fmt.Println(v3)
    for fi48 := 0; fi48 < 5; fi48++ {
        fmt.Println(fi48)
        var v49 uint64 = 18446744073709551615
        var v50 int32 = -1
        _ = fi48
        _ = v49
        _ = v50
    }
    v3 = v3
    var v51 int16 = 32767
    for fi52 := 0; fi52 < 8; fi52++ {
        v42 = append(v42, fi52)
        for fi53 := 0; fi53 < 6; fi53++ {
            var v54 uint32 = 2147483647
            _ = fi53
            _ = v54
        }
        _ = fi52
    }
    var v55 int32 = -521
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v51:", v51)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v44
    _ = v45
    _ = v51
    _ = v55
}
