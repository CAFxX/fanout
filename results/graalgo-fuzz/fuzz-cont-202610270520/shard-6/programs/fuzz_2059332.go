package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int32 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 uint16; F1 bool; F2 int16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint32; F1 uint64 }
func genFn0(p0 int16, p1 uint16, p2 int32) string {
    var v7 string = "\n\t"
    var v8 uint8 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return "foo bar"
}
func genFn1(p0 string) int8 {
    switch -2147483648 {
    case -30:
        v9 := GenS1{F0: 42, F1: 255, F2: 0}
        _ = v9
    default:
        for fi10 := 0; fi10 < 9; fi10++ {
            var v11 int64 = -797
            var v12 int16 = -1
            v13 := []int{(fi10 | (fi10 + fi10)), (fi10 - 4294967295)}
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
    }
    v14 := []int{2147483647}
    switch 32767 {
    case 409:
        switch 1 {
        case 0:
            var v15 uint64 = 4294967295
            _ = v15
        case -268:
            p0 = (p0 + (p0 + ((p0 + p0) + (((p0 + "foo bar") + p0) + "a"))))
        case 4294967295:
            fmt.Println(p0)
        default:
            for fi16 := 0; fi16 < 4; fi16++ {
                switch 844 {
                case 1:
                    var v17 uint8 = 195
                    _ = v17
                }
                var v18 uint16 = 7
                var v19 uint8 = 127
                _ = fi16
                _ = v18
                _ = v19
            }
        }
    }
    v20 := GenS2{F0: 0, F1: 42}
    _ = p0
    _ = v14
    _ = v20
    return 1
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = 3.14159
    v3 := []string{((("foo bar" + "foo bar") + "\n\t") + (("foo bar" + ("a" + (("\n\t" + "foo bar") + "a"))) + "\n\t"))}
    var v4 int64 = -418
    v5 := make(map[uint32]uint32)
    var v6 int8 = -11
    v21 := GenS3{F0: 1, F1: true, F2: 0}
    v2 = 0.1
    for fi22 := 0; fi22 < 20; fi22++ {
        for fi23 := 0; fi23 < 1; fi23++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi23
        }
        v6 += v6
        v24 := GenS1{F0: (v6 + (v6 + 0)), F1: 32767, F2: -708}
        _ = fi22
        _ = v24
    }
    if (2147483647 < -945) {
        if false {
            switch "a" {
            case "\n\t":
                fmt.Println(v4)
            case "a":
                var v25 uint16 = 750
                _ = v25
            default:
                var v26 uint8 = 2
                _ = v26
            }
        } else {
            v2 = (v2 + 1e10)
        }
    } else {
        fmt.Println(v4)
        v6 += (v6 * v6)
        var v27 string = ("foo bar" + "hello")
        _ = v27
    }
    v4 -= -9223372036854775808
    v2 = 0.1
    _ = len(v1)
    if (!false) {
        v4 += v4
        v3 = append(v3, "hello")
    } else {
        var v28 float64 = (0.0 - (-2.25 + (0.1 - 0.1)))
        switch ("a" + ("foo bar" + ("hello" + ""))) {
        case "\n\t":
            v29 := v1[356]
            _ = v29
        }
        var v30 bool = (2147483647 != 438)
        _ = v28
        _ = v30
    }
    switch "\n\t" {
    case "foo bar":
        for fi31 := 0; fi31 < 8; fi31++ {
            v32 := GenS4{F0: 0}
            _ = fi31
            _ = v32
        }
    case "a":
        v33 := GenS4{F0: 537}
        _ = v33
    default:
        v34 := GenS3{F0: 65535, F1: false, F2: 42}
        _ = v34
    }
    switch -2147483648 {
    case 42:
        if true {
            v35 := GenS4{F0: 65535}
            _ = v35
        } else {
            var v36 uint64 = 42
            _ = v36
        }
    case 4294967295:
        var v37 int = -2147483648
        _ = v37
    case 127:
        var v38 int16 = 351
        _ = v38
    }
    if true {
        fmt.Println(v2, v6, v2)
    } else {
        v4 += 0
        v4 += v4
        if (v6 == (v6 / ((v6) | 1))) {
            switch -1 {
            case 420:
                fmt.Println(v4, v6)
            case -9223372036854775808:
                var v39 uint = 42
                _ = v39
            default:
                for fi40 := 0; fi40 < 13; fi40++ {
                    if (0 <= 874) {
                        v41 := v5[2147483647]
                        var v42 uint16 = 32767
                        for fi43 := 0; fi43 < 11; fi43++ {
                            var v44 uint32 = v41
                            v2 -= float32(0.0)
                            v3 = append(v3, "\n\t")
                            _ = fi43
                            _ = v44
                        }
                        _ = v41
                        _ = v42
                    }
                    var v45 uint32 = 4294967295
                    _ = fi40
                    _ = v45
                }
            }
            v46 := GenS5{F0: 4294967295, F1: 909}
            _ = v46
        }
    }
    if true {
        v1[35] = 32768
    }
    v2 += 1.5
    if true {
        v4 = 0
        v4 = 0
    }
    v6 = (v6 | v6)
    var v47 float32 = 3.14159
    v1[450] = -196
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
