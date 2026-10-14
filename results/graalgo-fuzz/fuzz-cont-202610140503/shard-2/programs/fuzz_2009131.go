package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int8; F1 string; F2 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16; F1 uint8; F2 uint8 }
func genFn0() int {
    var v7 float64 = 1e10
    v7 += v7
    v7 += (float64(1e10) * v7)
    _ = v7
    return -2147483648
}
func genFn1(p0 uint16) uint8 {
    fmt.Println(p0)
    v8 := []int{127}
    for fi9 := 0; fi9 < 5; fi9++ {
        if len(v8) > 0 {
            v8[0] = (-(fi9))
        }
        for fi10 := 0; fi10 < 3; fi10++ {
            switch ("a" + ("a" + "a")) {
            case "hello":
                if (((p0 & 32767) <= p0) || true) {
                    var v11 uint64 = 1
                    _ = len(v8)
                    switch genFn0() {
                    case 32767:
                        var v12 float64 = 3.14159
                        _ = v12
                    case -272:
                        p0 = 855
                    case -348:
                        v11 = 488
                    default:
                        if len(v8) > 0 {
                            v8[0] = (fi9 ^ (fi9 - (fi10 / (((fi10 | fi10)) | 1))))
                        }
                    }
                    v11 = v11
                    fmt.Println(p0, v11, v11)
                    _ = v11
                }
            case "\n\t":
                var v13 string = "hello"
                _ = v13
            default:
                if true {
                    var v14 uint32 = 42
                    _ = v14
                } else {
                    var v15 string = "\n\t"
                    switch (((v15 + v15) + "hello") + "") {
                    case "\n\t":
                        var v16 uint64 = 452
                        _ = v16
                    case "a":
                        if false {
                            v15 = "\n\t"
                        }
                    case "hello":
                        p0 += (p0 & p0)
                    }
                    _ = v15
                }
            }
            switch (fi9 << ((0) & 63)) {
            case 9223372036854775807:
                var v17 string = ("a" + "foo bar")
                _ = v17
            case 32767:
                var v18 uint8 = 63
                _ = v18
            }
            _ = fi10
        }
        _ = fi9
    }
    v19 := GenS1{F0: 1, F1: 4294967295}
    _ = p0
    _ = v19
    _ = v8
    return 42
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]string)
    v3 := []int{9223372036854775807}
    v4 := []uint32{}
    v5 := []uint32{4294967295}
    var v6 int = 4294967295
    for fi20 := 0; fi20 < 8; fi20++ {
        if (505 >= 65535) {
            v21 := v1[(v6 | v6)]
            v22 := GenS2{F0: 0, F1: "hello", F2: 65535}
            _ = v21
            _ = v22
        }
        _ = fi20
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        _ = len(v1)
        v24 := GenS3{F0: -730}
        v6 = fi23
        _ = fi23
        _ = v24
    }
    v3 = append(v3, (-(v6)))
    if (0 > 42) {
        v6 = (v6 >> ((18446744073709551615) & 63))
    } else {
        for fi25 := 0; fi25 < 14; fi25++ {
            fmt.Println(fi25)
            _ = fi25
        }
    }
    var v26 float32 = 1.5
    var v27 uint64 = 477
    for fi28 := 0; fi28 < 1; fi28++ {
        fmt.Println(v26)
        for fi29 := 0; fi29 < 3; fi29++ {
            fmt.Println(fi29, v26)
            switch (fi29 ^ v6) {
            case -301:
                switch -1 {
                case 42:
                    v30 := GenS2{F0: -17, F1: "foo bar", F2: 4294967295}
                    _ = v30
                }
            case 32768:
                fmt.Println(v27, fi28, v26)
            }
            _ = fi29
        }
        _ = fi28
    }
    v27 = 18446744073709551615
    if false {
        switch genFn0() {
        case -2147483648:
            v26 -= 1e10
        case 32768:
            var v31 uint32 = 256
            _ = v31
        case -9223372036854775808:
            v26 = 1e10
        default:
            var v32 uint = 2
            _ = v32
        }
        var v33 int8 = int8(v27)
        _ = v33
    } else {
        v34 := GenS4{F0: 400}
        _ = v34
    }
    v26 -= v26
    switch "\n\t" {
    case "a":
        v35 := GenS5{F0: 505, F1: 42, F2: genFn1(0)}
        _ = v35
    case "":
        var v36 int = genFn0()
        _ = v36
    case "hello":
        fmt.Println(v26)
    }
    if len(v5) > 0 {
        v5[0] = 2
    }
    fmt.Println(v26)
    for fi37 := 0; fi37 < 11; fi37++ {
        v38 := GenS2{F0: -128, F1: "foo bar", F2: 65535}
        v26 = 0.1
        _ = fi37
        _ = v38
    }
    for fi39 := 0; fi39 < 17; fi39++ {
        v40 := GenS5{F0: 42, F1: uint8(v27), F2: uint8(v27)}
        _ = fi39
        _ = v40
    }
    switch ("hello" + "hello") {
    case "a":
        v2[v6] = "\n\t"
    case "foo bar":
        v41 := GenS5{F0: 255, F1: 106, F2: 151}
        _ = v41
    case "hello":
        fmt.Println(v27, v27, v27)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
