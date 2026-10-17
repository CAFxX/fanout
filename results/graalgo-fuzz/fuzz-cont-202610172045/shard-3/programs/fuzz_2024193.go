package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8 }
func genFn0(p0 string, p1 uint) uint8 {
    p0 = p0
    if (!true) {
        switch 4294967295 {
        case 403:
            p1 = 18446744073709551615
        }
        var v7 uint64 = 18446744073709551615
        switch (p0 + p0) {
        case "":
            p1 = 526
        case "hello":
            v7 -= (^(v7))
        case "foo bar":
            v7 -= v7
        }
        _ = v7
    } else {
        var v8 uint = p1
        _ = v8
    }
    var v9 uint8 = 128
    v9 -= 1
    _ = p0
    _ = p1
    _ = v9
    return (v9 ^ v9)
}

func main() {
    var v1 int = 32768
    var v2 bool = (0 == int32(v1))
    v3 := make(map[int]uint32)
    var v4 int32 = -1
    v5 := []string{"", "a", "", "\n\t"}
    v6 := make(map[uint32]uint32)
    v2 = v2
    fmt.Println(v1, v4)
    v1 -= (v1 - v1)
    _ = len(v6)
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 float64 = 3.14159
        for fi12 := 0; fi12 < 17; fi12++ {
            var v13 uint = 1
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    var v14 int8 = -128
    if v2 {
        fmt.Println(v4, v2)
        v4 = v4
    }
    var v15 uint64 = 973
    fmt.Println(v14, v2)
    if ((v14 % ((v14) | 1)) == v14) {
        var v16 uint = 255
        for fi17 := 0; fi17 < 18; fi17++ {
            var v18 int8 = v14
            v5 = append(v5, "hello")
            _ = fi17
            _ = v18
        }
        switch "hello" {
        case "\n\t":
            v19 := GenS1{F0: 9223372036854775807, F1: 2, F2: uint8(v16)}
            _ = v19
        case "":
            if ((v15 / ((v15) | 1)) <= (v15 | 1)) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            } else {
                var v20 uint8 = 0
                v20 = 255
                _ = v20
            }
        default:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        _ = v16
    }
    switch "foo bar" {
    case "":
        if v2 {
            v21 := GenS1{F0: 784, F1: 831, F2: 127}
            fmt.Println(v15, v2, v14)
            for fi22 := 0; fi22 < 20; fi22++ {
                switch "foo bar" {
                case "\n\t":
                    if v2 {
                        var v23 int64 = 42
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        for fi24 := 0; fi24 < 15; fi24++ {
                            fmt.Println(v15, v2)
                            fmt.Println(v1)
                            _ = fi24
                        }
                        _ = v23
                    } else {
                        v25 := GenS2{F0: ((("\n\t" + "a") + "foo bar") + ("a" + ""))}
                        _ = v25
                    }
                case "foo bar":
                    for fi26 := 0; fi26 < 4; fi26++ {
                        if v2 {
                            fmt.Println(v15, fi26, v2)
                        }
                        _ = fi26
                    }
                case "":
                    v14 -= v14
                }
                var v27 int64 = int64(v4)
                v28 := GenS3{F0: 100}
                _ = fi22
                _ = v27
                _ = v28
            }
            _ = v21
        }
    case "a":
        if false {
            var v29 uint8 = genFn0("a", 18446744073709551615)
            v4 = (v4 | (^(v4)))
            _ = v29
        } else {
            var v30 uint8 = 1
            v30 -= genFn0("foo bar", 1)
            var v31 uint64 = v15
            _ = v30
            _ = v31
        }
    }
    var v32 int16 = 32767
    if false {
        v5 = append(v5, "\n\t")
        var v33 uint16 = 2
        _ = v33
    }
    for fi34 := 0; fi34 < 11; fi34++ {
        var v35 int32 = (v4 | 1)
        fmt.Println(v4, v35)
        switch (-(v1)) {
        case 32768:
            switch fi34 {
            case 32767:
                fmt.Println(v2, v14)
            case 419:
                switch ("foo bar" + "foo bar") {
                case "foo bar":
                    var v36 int8 = -78
                    _ = v36
                default:
                    v37 := GenS2{F0: "hello"}
                    _ = v37
                }
            default:
                var v38 int8 = (v14 >> ((18446744073709551615) & 63))
                _ = v38
            }
        default:
            v4 += (v4 & v35)
        }
        _ = fi34
        _ = v35
    }
    v32 -= v32
    for fi39 := 0; fi39 < 4; fi39++ {
        var v40 float32 = 0.0
        v32 += (v32 - v32)
        v15 = (v15 + (^(v15)))
        _ = fi39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
