package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint64, p1 int32, p2 uint) uint16 {
    switch 42 {
    case -2147483648:
        var v7 int = -318
        _ = v7
    case 286:
        var v8 string = "a"
        _ = v8
    }
    switch "" {
    case "foo bar":
        v9 := []int{int(p0), 42, 99}
        _ = v9
    case "hello":
        v10 := GenS1{F0: false, F1: 100}
        _ = v10
    }
    fmt.Println(p0, p2)
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1(p0 int64, p1 int32) uint64 {
    for fi11 := 0; fi11 < 12; fi11++ {
        p1 -= p1
        var v12 int32 = 1
        _ = fi11
        _ = v12
    }
    if true {
        var v13 uint = 65535
        var v14 int32 = p1
        _ = v13
        _ = v14
    }
    switch int(p1) {
    case 32768:
        p1 -= 0
    case -427:
        v15 := GenS1{F0: ((true || true) || (genFn0(2, p1, 4294967295) == 42)), F1: -128}
        _ = v15
    case -275:
        v16 := []int{}
        _ = v16
    }
    v17 := GenS2{F0: p0, F1: -890}
    _ = p0
    _ = p1
    _ = v17
    return 4294967295
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = -1
    var v3 uint = 32
    v4 := []int{int(v2), 4294967295}
    var v5 int8 = 0
    v6 := make(map[int]bool)
    v18 := GenS1{F0: (v3 >= (v3 - (v3 >> (((v3 | (v3 + (v3 + v3)))) & 63)))), F1: v5}
    delete(v6, 4294967295)
    switch -403 {
    case -422:
        v5 -= 42
    case 4294967295:
        if ("hello" < ("a" + "a")) {
            switch "foo bar" {
            case "\n\t":
                var v19 uint32 = 4294967295
                _ = v19
            case "hello":
                var v20 int8 = 53
                _ = v20
            }
        }
    case -1:
        fmt.Println(v3)
    default:
        for fi21 := 0; fi21 < 13; fi21++ {
            for fi22 := 0; fi22 < 5; fi22++ {
                switch int(v3) {
                case 42:
                    if ((true && (((1e10 / (1e10 / 3.14159)) > float32(float64(1.5))) && (false || (!true)))) || true) {
                        switch (fi22 ^ 32767) {
                        case 1:
                            fmt.Println(v2)
                        case 859:
                            var v23 bool = (((fi21 << (((v3 / ((42) | 1))) & 63)) >= 290) || false)
                            _ = v23
                        case -2147483648:
                            v24 := GenS1{F0: (!true), F1: (v5 + (v5 | v5))}
                            v3 = 18446744073709551615
                            v25 := GenS1{F0: (0 == 2), F1: (v5 * (^(v5)))}
                            _ = v24
                            _ = v25
                        default:
                            v5 = 100
                        }
                        v2 -= v2
                        var v26 int = fi21
                        _ = v26
                    }
                default:
                    var v27 uint = (v3 | (v3 % ((v3) | 1)))
                    _ = v27
                }
                _ = fi22
            }
            v2 -= (v2 + 42)
            _ = fi21
        }
    }
    fmt.Println(v3, v5, v5)
    var v28 float64 = -2.25
    switch ("foo bar" + "hello") {
    case "hello":
        _ = len(v4)
    case "foo bar":
        v2 -= v2
    default:
        var v29 int64 = v2
        _ = v29
    }
    switch -9223372036854775808 {
    case 128:
        fmt.Println(v3, v3, v5)
    default:
        v2 -= v2
    }
    if true {
        v1[65535] = ("" + "a")
    } else {
        if (true && (32767 == 4294967295)) {
            v2 = 0
        } else {
            switch "\n\t" {
            case "foo bar":
                if true {
                    var v30 float32 = 0.0
                    v3 += v3
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v30
                } else {
                    for fi31 := 0; fi31 < 20; fi31++ {
                        fmt.Println(v2)
                        _ = fi31
                    }
                    delete(v6, 1)
                }
            case "":
                fmt.Println(v3, v3)
            case "hello":
                v1[84] = "a"
            }
        }
    }
    if false {
        switch "" {
        case "\n\t":
            var v32 int16 = 0
            _ = v32
        case "foo bar":
            switch "foo bar" {
            case "\n\t":
                var v33 int64 = v2
                _ = v33
            default:
                for fi34 := 0; fi34 < 3; fi34++ {
                    v3 = v3
                    _ = fi34
                }
            }
        case "a":
            v35 := GenS3{F0: 127}
            _ = v35
        default:
            v2 = 1
        }
        var v36 float32 = 3.14159
        _ = v36
    } else {
        if (-2.25 < 0.1) {
            v1[42] = "hello"
            fmt.Println(v3, v3, v28)
        }
    }
    switch "\n\t" {
    case "foo bar":
        v37 := GenS2{F0: v2, F1: 9}
        _ = v37
    case "":
        v1[2] = ("\n\t" + "a")
    case "hello":
        v38 := GenS3{F0: 0}
        _ = v38
    }
    v39 := GenS1{F0: (-1 >= 939), F1: (v5 - 42)}
    var v40 uint = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
