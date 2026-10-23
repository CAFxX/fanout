package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 string; F1 int16; F2 string }
type GenS4 struct { F0 uint; F1 uint8; F2 bool }
type GenS5 struct { F0 uint32; F1 uint32 }
func genFn0(p0 int64, p1 uint, p2 int64) uint {
    if false {
        v5 := GenS1{F0: "\n\t"}
        if false {
            var v6 int32 = 1
            for fi7 := 0; fi7 < 20; fi7++ {
                var v8 int64 = (p0 ^ 763)
                _ = fi7
                _ = v8
            }
            for fi9 := 0; fi9 < 1; fi9++ {
                p2 -= (p0 / ((-177) | 1))
                _ = fi9
            }
            _ = v6
        }
        v10 := GenS2{F0: 1000, F1: 951}
        _ = v10
        _ = v5
    }
    var v11 uint64 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return p1
}

func main() {
    var v1 int32 = 706
    var v2 uint32 = 4294967295
    var v3 bool = false
    var v4 uint32 = 4294967295
    var v12 int = 42
    v13 := []int{(v12 / ((v12) | 1)), -427, v12, -809}
    if (-91 < 0) {
        var v14 int8 = 42
        fmt.Println(v4, v3)
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        _ = v14
    }
    fmt.Println(v4, v2)
    if false {
        fmt.Println(v4, v2)
        v15 := GenS3{F0: ("a" + ""), F1: -32768, F2: (("\n\t" + "") + ("foo bar" + ""))}
        _ = v15
    } else {
        _ = len(v13)
    }
    v2 = (v2 >> ((350) & 63))
    if v3 {
        var v16 int64 = 225
        v1 = v1
        _ = v16
    }
    switch (v12 >> ((65535) & 63)) {
    case 32768:
        v4 -= v4
    case 61:
        switch (("\n\t" + "hello") + "") {
        case "foo bar":
            if ((v3 || v3) || v3) {
                fmt.Println(v2, v3)
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
                v3 = v3
            }
        default:
            for fi17 := 0; fi17 < 4; fi17++ {
                for fi18 := 0; fi18 < 10; fi18++ {
                    switch "foo bar" {
                    case "a":
                        v1 = v1
                    case "foo bar":
                        if (!true) {
                            v12 += fi18
                            v19 := GenS3{F0: ("" + ""), F1: -1, F2: "hello"}
                            var v20 int64 = 42
                            _ = v19
                            _ = v20
                        } else {
                            var v21 float64 = 1e10
                            v3 = true
                            if len(v13) > 0 {
                                v13[0] = (v12 * v12)
                            }
                            _ = v21
                        }
                    case "\n\t":
                        var v22 int = (v12 + 101)
                        v4 += (^(v4))
                        var v23 uint64 = 18446744073709551615
                        var v24 uint = 484
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    switch "a" {
                    case "hello":
                        var v25 int8 = -1
                        _ = v25
                    case "foo bar":
                        var v26 int64 = 42
                        v27 := GenS4{F0: 42, F1: 1, F2: false}
                        _ = v26
                        _ = v27
                    case "a":
                        if true {
                            switch 9223372036854775807 {
                            case -2147483648:
                                v12 = (v12 << ((42) & 63))
                            case 2147483647:
                                var v28 int = (v12 >> ((genFn0(-672, 255, 0)) & 63))
                                _ = v28
                            case 127:
                                if (!(v3 && (!v3))) {
                                    var v29 int8 = 1
                                    fmt.Println(fi17, v3)
                                    var v30 uint64 = 2
                                    _ = v29
                                    _ = v30
                                } else {
                                    switch (^(v12)) {
                                    case -9223372036854775808:
                                        if (v2 >= (^(v2))) {
                                            v31 := GenS4{F0: 284, F1: 1, F2: true}
                                            _ = v31
                                        } else {
                                            var v32 uint64 = 2
                                            fmt.Println(fi18)
                                            _ = v32
                                        }
                                        var v33 int16 = 1000
                                        _ = v33
                                    case 9223372036854775807:
                                        switch ("" + "") {
                                        case "a":
                                            fmt.Println(v1)
                                        default:
                                            if false {
                                                switch fi17 {
                                                case 32768:
                                                    var v34 float32 = -2.25
                                                    fmt.Println(v1, v34, v3)
                                                    v34 = v34
                                                    _ = v34
                                                case 879:
                                                    fmt.Println(v4, v3, fi17)
                                                case 42:
                                                    v35 := GenS4{F0: 65535, F1: 128, F2: false}
                                                    _ = v35
                                                }
                                                v2 -= (v2 % (((v2 | v4)) | 1))
                                                var v36 int8 = 100
                                                v37 := GenS3{F0: "a", F1: int16(v2), F2: "foo bar"}
                                                var v38 int8 = v36
                                                _ = v36
                                                _ = v37
                                                _ = v38
                                            } else {
                                                var v39 string = "a"
                                                var v40 uint = 0
                                                _ = v39
                                                _ = v40
                                            }
                                        }
                                    case 42:
                                        v1 = v1
                                    default:
                                        fmt.Println(v12, v1)
                                    }
                                    v2 = (v2 & v2)
                                    if v3 {
                                        v41 := GenS4{F0: genFn0(-501, genFn0(1, 162, -565), 0), F1: 127, F2: v3}
                                        var v42 string = "hello"
                                        if (v3 || true) {
                                            v2 = 408
                                            v2 = 2147483648
                                            v43 := GenS1{F0: ((v42 + v42) + v42)}
                                            _ = v43
                                        }
                                        _ = v41
                                        _ = v42
                                    }
                                }
                            default:
                                switch (fi18 * 127) {
                                case 9223372036854775807:
                                    fmt.Println(v2, fi17)
                                case 4294967295:
                                    fmt.Println(v1)
                                    switch (fi17 & 42) {
                                    case -1:
                                        if v3 {
                                            v44 := GenS2{F0: 42, F1: 0}
                                            _ = v44
                                        }
                                    default:
                                        var v45 float64 = ((1.5 * -2.25) / ((1.5 - 0.1) + (float64(0.0) / (0.1 / (3.14159 / -2.25)))))
                                        _ = v45
                                    }
                                    var v46 uint = genFn0(-585, 2, -9223372036854775808)
                                    _ = v46
                                default:
                                    if ((v1 << ((18446744073709551615) & 63)) <= v1) {
                                        v47 := GenS2{F0: 18, F1: uint16(v2)}
                                        _ = v47
                                    }
                                }
                            }
                            v1 = (^(v1))
                        } else {
                            v12 -= 490
                            switch fi17 {
                            case 32768:
                                var v48 int16 = 1
                                _ = v48
                            }
                        }
                    }
                    _ = fi18
                }
                _ = fi17
            }
        }
    }
    var v49 int16 = int16(v2)
    v50 := GenS5{F0: 365, F1: v4}
    v51 := GenS5{F0: 1, F1: v4}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v50
    _ = v51
}
