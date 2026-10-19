package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint16 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool; F1 int }
type GenS6 struct { F0 int64 }
func genFn0(p0 int, p1 uint8) int16 {
    p1 = (p1 | p1)
    v6 := GenS1{F0: "\n\t", F1: ("hello" + "a"), F2: 1}
    v7 := GenS1{F0: ("a" + ""), F1: ("a" + "a"), F2: 0}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 int32, p1 int64) uint16 {
    p1 += 999
    if (1 >= p0) {
        p0 -= p0
        for fi8 := 0; fi8 < 2; fi8++ {
            if (!((0 > 1) || true)) {
                var v9 uint16 = 32767
                _ = v9
            } else {
                var v10 float32 = 1.5
                _ = v10
            }
            _ = fi8
        }
        var v11 float32 = 0.1
        _ = v11
    }
    switch "hello" {
    case "hello":
        for fi12 := 0; fi12 < 14; fi12++ {
            for fi13 := 0; fi13 < 17; fi13++ {
                v14 := GenS1{F0: "foo bar", F1: "", F2: 704}
                switch fi12 {
                case 9223372036854775807:
                    p1 += (p1 >> ((42) & 63))
                case 42:
                    p0 = p0
                }
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
    case "foo bar":
        if false {
            v15 := GenS2{F0: 18446744073709551615, F1: p1}
            _ = v15
        } else {
            var v16 bool = ("hello" == "")
            p1 = 775
            _ = v16
        }
    case "\n\t":
        fmt.Println(p1)
    }
    p0 = p0
    switch -850 {
    case 9223372036854775807:
        for fi17 := 0; fi17 < 1; fi17++ {
            var v18 int = (fi17 + (-(fi17)))
            p1 += p1
            fmt.Println(p0)
            _ = fi17
            _ = v18
        }
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{392}
    v3 := make(map[string]string)
    var v4 float32 = 1.5
    v5 := make(map[uint32]uint32)
    v4 = v4
    var v19 uint16 = 0
    for fi20 := 0; fi20 < 3; fi20++ {
        fmt.Println(fi20, v4)
        v21 := GenS2{F0: 520, F1: 60}
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        fmt.Println(fi22, fi22, v4)
        _ = fi22
    }
    if (true || (8 == 255)) {
        v23 := GenS2{F0: 65535, F1: 9223372036854775807}
        switch "foo bar" {
        case "a":
            for fi24 := 0; fi24 < 15; fi24++ {
                v25 := GenS1{F0: "foo bar", F1: "", F2: (v19 ^ v19)}
                v4 += (float32(1.5) - 1.5)
                _ = fi24
                _ = v25
            }
        }
        _ = v23
    }
    switch "foo bar" {
    case "a":
        if false {
            if true {
                v26 := GenS3{F0: -905, F1: 34}
                _ = v26
            } else {
                var v27 int16 = 535
                if (42 > 541) {
                    v27 = v27
                    fmt.Println(v27, v27)
                    var v28 uint = 2
                    _ = v28
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if (18446744073709551615 <= 42) {
                        if len(v2) > 0 {
                            v2[0] = -9223372036854775808
                        }
                        fmt.Println(v19, v27)
                        v4 += v4
                    } else {
                        v29 := GenS3{F0: (v27 | v27), F1: 403}
                        var v30 uint16 = (v19 >> ((302) & 63))
                        _ = v29
                        _ = v30
                    }
                    _ = len(v3)
                }
                _ = len(v2)
                _ = v27
            }
            v31 := GenS4{F0: "\n\t"}
            _ = v31
        } else {
            v32 := GenS5{F0: false, F1: -9223372036854775808}
            _ = len(v5)
            _ = v32
        }
    case "":
        var v33 uint = 255
        _ = v33
    case "foo bar":
        for fi34 := 0; fi34 < 3; fi34++ {
            var v35 bool = (true && (!true))
            _ = fi34
            _ = v35
        }
    default:
        var v36 int32 = -648
        _ = v36
    }
    v37 := GenS6{F0: 133}
    v38 := GenS6{F0: 9223372036854775807}
    var v39 float64 = 3.14159
    if ((1.5 - 3.14159) == float32(0.1)) {
        if len(v2) > 0 {
            v2[0] = 0
        }
    } else {
        for fi40 := 0; fi40 < 16; fi40++ {
            v19 += 32767
            _ = fi40
        }
        var v41 string = ("\n\t" + "")
        _ = v41
    }
    if (!false) {
        var v42 int = 582
        if false {
            var v43 int8 = 26
            v44 := GenS2{F0: 255, F1: 42}
            _ = v43
            _ = v44
        }
        _ = v42
    }
    var v45 int = -9223372036854775808
    var v46 int8 = 33
    _ = len(v3)
    var v47 uint64 = 391
    v48 := GenS2{F0: 0, F1: -9223372036854775808}
    v39 = -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
}
