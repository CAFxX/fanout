package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint64; F1 int32; F2 int16 }
func genFn0(p0 int) uint {
    p0 = (p0 << ((1) & 63))
    var v7 int8 = 0
    fmt.Println(p0, p0)
    _ = p0
    _ = v7
    return 396
}
func genFn1(p0 uint, p1 int64, p2 uint32) uint32 {
    for fi8 := 0; fi8 < 14; fi8++ {
        for fi9 := 0; fi9 < 1; fi9++ {
            v10 := []int{470, 258, fi8}
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 15; fi11++ {
            v12 := []int{fi11, -340, 1, (fi11 << (((p0 >> ((65535) & 63))) & 63))}
            fmt.Println(fi8, fi11, p2)
            switch "hello" {
            case "a":
                var v13 int = fi11
                _ = v13
            default:
                p2 = (^(p2))
            }
            var v14 string = ("\n\t" + ((("hello" + ("" + "a")) + "\n\t") + "\n\t"))
            _ = fi11
            _ = v12
            _ = v14
        }
        _ = fi8
    }
    if (!(100 == -47)) {
        var v15 uint64 = 717
        fmt.Println(v15, p2, p1)
        v16 := GenS1{F0: -1}
        _ = v15
        _ = v16
    }
    v17 := []int{1, int(p2), 127, -878}
    if (("foo bar" + ("\n\t" + "")) <= "") {
        for fi18 := 0; fi18 < 13; fi18++ {
            var v19 uint64 = 4294967295
            v19 = v19
            if (!true) {
                var v20 float64 = -2.25
                switch fi18 {
                case 2147483647:
                    p2 = 4294967295
                case 32767:
                    v21 := GenS1{F0: -2147483648}
                    _ = v21
                default:
                    p0 += (p0 + genFn0((fi18 | 128)))
                }
                p1 -= p1
                _ = v20
            } else {
                v22 := GenS2{F0: -597}
                var v23 uint = p0
                v17 = append(v17, 1)
                _ = v22
                _ = v23
            }
            _ = fi18
            _ = v19
        }
    } else {
        var v24 int32 = 1
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
        _ = v24
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return (p2 >> ((genFn0(128)) & 63))
}

func main() {
    v1 := make(map[int]int)
    var v2 bool = (false && true)
    var v3 uint = 496
    var v4 uint16 = 988
    var v5 string = (("a" + "\n\t") + (("\n\t" + "foo bar") + "\n\t"))
    var v6 float64 = 3.14159
    if (!((!(!v2)) && v2)) {
        switch ((v5 + "a") + ("a" + "hello")) {
        case "\n\t":
            fmt.Println(v2, v5, v4)
        case "foo bar":
            v25 := []int{656}
            _ = v25
        case "hello":
            if (!v2) {
                for fi26 := 0; fi26 < 1; fi26++ {
                    v5 = (v5 + (("foo bar" + v5) + "hello"))
                    v27 := GenS1{F0: -2147483648}
                    v6 += v6
                    _ = fi26
                    _ = v27
                }
                var v28 int16 = 32767
                for fi29 := 0; fi29 < 9; fi29++ {
                    fmt.Println(v2)
                    _ = fi29
                }
                _ = v28
            } else {
                var v30 int16 = -907
                _ = v30
            }
        default:
            v31 := GenS2{F0: -1}
            _ = v31
        }
        var v32 uint64 = 618
        v33 := GenS3{F0: "\n\t"}
        _ = v32
        _ = v33
    }
    v3 = 858
    for fi34 := 0; fi34 < 1; fi34++ {
        v5 = "foo bar"
        v3 += v3
        v1[-9223372036854775808] = fi34
        _ = fi34
    }
    for fi35 := 0; fi35 < 13; fi35++ {
        if (((1 != 1) || false) || ((((v6 - v6) + v6) < -2.25) || ((1 > 42) || (0 == -694)))) {
            var v36 float64 = v6
            _ = v36
        }
        _ = len(v1)
        v37 := GenS4{F0: 944, F1: ((((2147483648 == genFn1((v3 << ((v3) & 63)), 102, 2)) || true) || v2) && v2)}
        _ = fi35
        _ = v37
    }
    if true {
        v38 := GenS1{F0: -2147483648}
        if true {
            v2 = true
        }
        _ = v38
    }
    v6 += (v6 + -2.25)
    var v39 int16 = 42
    switch 267 {
    case 127:
        switch 907 {
        case 42:
            var v40 float64 = v6
            _ = v40
        case -606:
            var v41 uint64 = 706
            _ = v41
        case -425:
            switch "a" {
            case "hello":
                if false {
                    fmt.Println(v6, v3)
                }
            case "foo bar":
                if (!(127 < 196)) {
                    var v42 uint = (v3 - genFn0(32767))
                    v43 := GenS5{F0: 4294967295}
                    _ = v42
                    _ = v43
                }
            }
        }
    case 2147483647:
        fmt.Println(v3, v5)
    case 582:
        v44 := GenS3{F0: (v5 + "hello")}
        _ = v44
    default:
        for fi45 := 0; fi45 < 14; fi45++ {
            fmt.Println(v5, v3)
            _ = fi45
        }
    }
    v46 := GenS6{F0: 4294967295, F1: -572, F2: v39}
    for fi47 := 0; fi47 < 8; fi47++ {
        var v48 int64 = -838
        var v49 string = v5
        _ = fi47
        _ = v48
        _ = v49
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
