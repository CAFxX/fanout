package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 string }
type GenS2 struct { F0 int64; F1 uint; F2 string }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int64, p1 uint32) uint16 {
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := []int{}
        v8 := GenS1{F0: 4294967295, F1: false, F2: ("foo bar" + "a")}
        switch 127 {
        case 1:
            p0 = (p0 + p0)
        case 0:
            fmt.Println(p1, p1, p0)
        default:
            if (1 <= 32767) {
                for fi9 := 0; fi9 < 12; fi9++ {
                    _ = len(v7)
                    _ = fi9
                }
                switch (fi6 * (fi6 << ((2) & 63))) {
                case 128:
                    for fi10 := 0; fi10 < 16; fi10++ {
                        _ = len(v7)
                        v11 := GenS2{F0: (p0 >> ((4294967295) & 63)), F1: 18446744073709551615, F2: "a"}
                        if (!false) {
                            var v12 uint16 = 367
                            if (0.0 > 1e10) {
                                v12 += (v12 >> ((1) & 63))
                                var v13 uint16 = v12
                                _ = v13
                            }
                            v14 := GenS3{F0: 127, F1: 2, F2: (p1 + (p1 - p1))}
                            var v15 float32 = -2.25
                            _ = v12
                            _ = v14
                            _ = v15
                        }
                        _ = fi10
                        _ = v11
                    }
                default:
                    var v16 uint8 = 108
                    _ = v16
                }
            } else {
                p1 -= p1
                for fi17 := 0; fi17 < 1; fi17++ {
                    p1 = (^(p1))
                    _ = fi17
                }
                fmt.Println(fi6, p0, fi6)
            }
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    var v18 int32 = -749
    for fi19 := 0; fi19 < 11; fi19++ {
        v18 = (v18 / ((1) | 1))
        if true {
            v18 = v18
        }
        var v20 int64 = p0
        _ = fi19
        _ = v20
    }
    p1 = 42
    _ = p0
    _ = p1
    _ = v18
    return 0
}
func genFn1() int {
    switch ("a" + "foo bar") {
    case "a":
        var v21 int = -52
        _ = v21
    case "hello":
        for fi22 := 0; fi22 < 9; fi22++ {
            var v23 int8 = 0
            _ = fi22
            _ = v23
        }
    case "foo bar":
        v24 := []int{}
        _ = v24
    }
    var v25 bool = false
    for fi26 := 0; fi26 < 8; fi26++ {
        v25 = v25
        switch fi26 {
        case -2147483648:
            switch (fi26 / ((fi26) | 1)) {
            case 42:
                var v27 uint64 = 0
                _ = v27
            }
        case -9223372036854775808:
            for fi28 := 0; fi28 < 6; fi28++ {
                if true {
                    v29 := []int{fi28}
                    var v30 uint64 = 0
                    _ = v29
                    _ = v30
                } else {
                    v31 := GenS4{F0: 127}
                    _ = v31
                }
                var v32 uint32 = 383
                var v33 float32 = 0.0
                _ = fi28
                _ = v32
                _ = v33
            }
        default:
            var v34 uint8 = 255
            _ = v34
        }
        _ = fi26
    }
    switch (("foo bar" + "a") + ("hello" + "\n\t")) {
    case "foo bar":
        v35 := GenS4{F0: 0}
        _ = v35
    }
    _ = v25
    return 32768
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []int{9223372036854775807, 92}
    var v3 float64 = 0.0
    var v4 bool = false
    var v5 bool = (251 != 255)
    v36 := GenS3{F0: uint8(v1), F1: 42, F2: 2147483647}
    var v37 int = 1
    _ = len(v2)
    var v38 int16 = 1
    for fi39 := 0; fi39 < 8; fi39++ {
        fmt.Println(v3)
        v1 = 2
        fmt.Println(v38, fi39, v4)
        _ = fi39
    }
    v1 = (v1 * v1)
    var v40 float64 = float64(0.1)
    var v41 uint64 = v1
    var v42 int = (-(v37))
    fmt.Println(v38)
    fmt.Println(v1, v42, v40)
    v43 := GenS1{F0: 186, F1: v4, F2: "\n\t"}
    if true {
        if (!false) {
            var v44 uint32 = 0
            v40 = 1e10
            _ = v44
        } else {
            var v45 int8 = 127
            if len(v2) > 0 {
                v2[0] = (v42 * (v42 + 0))
            }
            _ = v45
        }
        var v46 int16 = 575
        _ = v46
    }
    _ = len(v2)
    fmt.Println(v38, v38, v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
