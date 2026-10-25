package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 int32; F1 uint; F2 int8 }
type GenS4 struct { F0 int16; F1 int64 }
func genFn0(p0 uint32, p1 uint8, p2 uint16) uint64 {
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := []int{fi7, (fi7 & fi7), (^(fi7))}
        switch "a" {
        case "foo bar":
            for fi9 := 0; fi9 < 1; fi9++ {
                p1 = uint8(p0)
                _ = fi9
            }
        case "hello":
            _ = len(v8)
        case "":
            p2 -= 0
        }
        var v10 int = fi7
        _ = fi7
        _ = v10
        _ = v8
    }
    v11 := GenS1{F0: p2, F1: -332}
    fmt.Println(p1, p1)
    var v12 int = 42
    for fi13 := 0; fi13 < 3; fi13++ {
        p0 = p0
        fmt.Println(v12, p1, p0)
        _ = fi13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return uint64(p0)
}

func main() {
    var v1 float64 = (float64((((0.0 * (3.14159 + float32((((0.1 - 3.14159) + 0.1) * float64(0.0))))) + 1.5) / -2.25)) / 1.0)
    v2 := []string{("a" + "\n\t"), "a", "a", ("a" + "hello")}
    v3 := make(map[uint32]int)
    v4 := []int{}
    v5 := make(map[string]bool)
    v6 := make(map[int]uint32)
    v1 = v1
    var v14 int64 = 190
    var v15 float32 = (0.0 - 3.14159)
    var v16 float32 = (v15 - 1e10)
    var v17 uint8 = 26
    fmt.Println(v14)
    v1 = 1e10
    for fi18 := 0; fi18 < 18; fi18++ {
        if (true && (642 <= 0)) {
            switch fi18 {
            case 2147483647:
                v3[2] = 32768
            case 127:
                var v19 int16 = 42
                _ = v19
            case 32767:
                v16 = 0.1
            }
            for fi20 := 0; fi20 < 10; fi20++ {
                if false {
                    switch (fi18 | fi20) {
                    case 0:
                        var v21 float32 = v15
                        _ = v21
                    }
                    if (255 >= 255) {
                        v22 := GenS2{F0: 1, F1: ("foo bar" + "\n\t")}
                        _ = len(v2)
                        v23 := GenS1{F0: 1, F1: v14}
                        _ = v22
                        _ = v23
                    }
                    v24 := GenS1{F0: 32767, F1: (^(v14))}
                    _ = v24
                } else {
                    v25 := GenS1{F0: 65535, F1: v14}
                    _ = v25
                }
                var v26 int16 = 0
                _ = fi20
                _ = v26
            }
        }
        var v27 uint64 = 4294967295
        _ = fi18
        _ = v27
    }
    v28 := GenS3{F0: 810, F1: uint(v17), F2: 127}
    v14 -= (^(v14))
    for fi29 := 0; fi29 < 6; fi29++ {
        for fi30 := 0; fi30 < 4; fi30++ {
            v3[0] = 376
            _ = fi30
        }
        v31 := GenS2{F0: 666, F1: "hello"}
        _ = fi29
        _ = v31
    }
    var v32 uint32 = 1
    for fi33 := 0; fi33 < 9; fi33++ {
        v34 := v5[""]
        v6[1] = 192
        v16 = (1e10 / ((v15 / v15) + v15))
        _ = fi33
        _ = v34
    }
    v15 += v16
    if ("foo bar" <= (("foo bar" + "") + ("" + "a"))) {
        var v35 int16 = -1
        if false {
            v2 = append(v2, ("foo bar" + "hello"))
            if (!((!false) && (v16 != v15))) {
                v16 += v16
                if ((v35 & v35) <= (-(v35))) {
                    switch ("hello" + "hello") {
                    case "a":
                        switch -309 {
                        case 9223372036854775807:
                            fmt.Println(v16, v16)
                        }
                    case "foo bar":
                        var v36 uint = 4294967295
                        _ = v36
                    case "\n\t":
                        for fi37 := 0; fi37 < 6; fi37++ {
                            switch "hello" {
                            case "\n\t":
                                v38 := GenS1{F0: 32767, F1: -1}
                                _ = v38
                            case "foo bar":
                                var v39 uint8 = (v17 ^ (v17 >> ((65535) & 63)))
                                _ = v39
                            }
                            _ = fi37
                        }
                    }
                    for fi40 := 0; fi40 < 6; fi40++ {
                        v15 = v15
                        _ = fi40
                    }
                } else {
                    var v41 int32 = 42
                    _ = v41
                }
            } else {
                fmt.Println(v32, v1)
                v42 := GenS4{F0: v35, F1: 9223372036854775807}
                if true {
                    var v43 float32 = ((v15 + v16) + v16)
                    var v44 uint8 = 127
                    for fi45 := 0; fi45 < 11; fi45++ {
                        v32 = (^(v32))
                        switch fi45 {
                        case 32767:
                            v16 += ((v15 / (v16 * (v43 * -2.25))) - v43)
                        case 517:
                            fmt.Println(v43, v16)
                        case -1:
                            var v46 uint16 = 65535
                            _ = v46
                        default:
                            v47 := GenS2{F0: 2, F1: ("foo bar" + "")}
                            _ = v47
                        }
                        v43 += (v16 / v16)
                        _ = fi45
                    }
                    _ = v43
                    _ = v44
                }
                _ = v42
            }
            var v48 float32 = ((-2.25 + 1e10) * (v16 * v16))
            _ = v48
        } else {
            var v49 uint16 = 32767
            var v50 int16 = (v35 >> ((2) & 63))
            _ = v49
            _ = v50
        }
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
