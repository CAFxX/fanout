package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int8; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint8; F1 uint }
func genFn0(p0 int8, p1 bool, p2 int32) string {
    switch 587 {
    case 2147483647:
        var v5 int64 = 9223372036854775807
        _ = v5
    case 32768:
        var v6 int = 128
        _ = v6
    case 42:
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 uint16 = 2
            p1 = (2147483647 == 1)
            _ = fi7
            _ = v8
        }
    default:
        var v9 string = "a"
        _ = v9
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        p2 += (p2 & p2)
        for fi11 := 0; fi11 < 17; fi11++ {
            fmt.Println(p2, p2)
            if p1 {
                v12 := GenS1{F0: 596, F1: p0}
                p1 = (p1 || true)
                if (42 > 640) {
                    if (!true) {
                        p0 = p0
                        v13 := []int{fi10, (fi11 >> ((42) & 63))}
                        v14 := GenS2{F0: (^(p0)), F1: 2, F2: 0}
                        _ = v13
                        _ = v14
                    } else {
                        var v15 uint = 1
                        var v16 int64 = 42
                        _ = v15
                        _ = v16
                    }
                    fmt.Println(p1, fi11)
                } else {
                    if true {
                        var v17 uint64 = 837
                        v18 := GenS1{F0: 42, F1: p0}
                        _ = v17
                        _ = v18
                    } else {
                        if p1 {
                            var v19 int64 = -1
                            _ = v19
                        } else {
                            switch "" {
                            case "a":
                                p0 = p0
                            case "\n\t":
                                switch fi10 {
                                case -2147483648:
                                    var v20 uint32 = 1
                                    var v21 float32 = 1.5
                                    var v22 uint32 = v20
                                    _ = v20
                                    _ = v21
                                    _ = v22
                                case 32767:
                                    p0 -= 1
                                case -1:
                                    var v23 uint64 = 4294967295
                                    _ = v23
                                }
                            case "":
                                fmt.Println(fi10)
                            }
                        }
                        p0 = 1
                    }
                }
                _ = v12
            }
            _ = fi11
        }
        _ = fi10
    }
    var v24 float32 = (-2.25 - (3.14159 + 0.1))
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    return "\n\t"
}

func main() {
    v1 := []int{32767, 0, 9223372036854775807}
    var v2 uint = 944
    v3 := make(map[int]string)
    var v4 int32 = 0
    v25 := v3[9223372036854775807]
    _ = len(v1)
    for fi26 := 0; fi26 < 15; fi26++ {
        if (0 != 182) {
            for fi27 := 0; fi27 < 7; fi27++ {
                if true {
                    var v28 string = v25
                    var v29 uint8 = 14
                    switch v25 {
                    case "\n\t":
                        var v30 int16 = 0
                        _ = v30
                    case "a":
                        fmt.Println(v25)
                    }
                    _ = v28
                    _ = v29
                }
                fmt.Println(fi26, fi26, v25)
                v25 = (v25 + v25)
                _ = fi27
            }
            var v31 string = v25
            _ = v31
        } else {
            v1 = append(v1, (fi26 >> ((v2) & 63)))
        }
        _ = fi26
    }
    switch -1 {
    case 0:
        if false {
            v32 := GenS1{F0: 1000, F1: 0}
            v2 -= v2
            v25 = (("a" + v25) + v25)
            _ = v32
        }
    case -518:
        var v33 float64 = float64(0.1)
        _ = v33
    }
    v3[4294967295] = v25
    for fi34 := 0; fi34 < 19; fi34++ {
        for fi35 := 0; fi35 < 4; fi35++ {
            v4 -= (v4 * (v4 / (((-(v4))) | 1)))
            v4 = -2147483648
            _ = fi35
        }
        v3[(fi34 / ((fi34) | 1))] = v25
        _ = fi34
    }
    if (false && false) {
        switch ("hello" + "foo bar") {
        case "hello":
            v36 := GenS3{F0: 0, F1: v2}
            _ = v36
        case "":
            fmt.Println(v4)
        default:
            if (v2 == v2) {
                var v37 uint = (v2 ^ 65535)
                _ = v37
            }
        }
        switch "foo bar" {
        case "":
            fmt.Println(v4, v25)
        case "foo bar":
            v2 = v2
        }
        switch -643 {
        case -665:
            var v38 int32 = v4
            _ = v38
        case -1:
            var v39 string = v25
            _ = v39
        case 42:
            v40 := GenS1{F0: -478, F1: -1}
            _ = v40
        }
    }
    v2 = (v2 | v2)
    for fi41 := 0; fi41 < 7; fi41++ {
        var v42 uint8 = 127
        _ = fi41
        _ = v42
    }
    fmt.Println(v2, v4, v4)
    v4 = v4
    var v43 float32 = (1.5 / 1.0)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v43
}
