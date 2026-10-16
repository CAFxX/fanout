package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int32 }
type GenS2 struct { F0 int32; F1 int8; F2 int32 }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 uint; F1 int16 }
func genFn0(p0 int16, p1 uint, p2 string) string {
    var v5 int64 = 644
    if (-21 >= -1) {
        var v6 string = p2
        v7 := []int{845, 127, 128}
        _ = v6
        _ = v7
    } else {
        var v8 string = p2
        if ((v5 << ((p1) & 63)) < v5) {
            switch "hello" {
            case "a":
                var v9 uint16 = 65535
                _ = v9
            case "hello":
                fmt.Println(v8, p1)
            default:
                fmt.Println(v5)
            }
            if (-958 > -119) {
                p2 = v8
                if true {
                    v10 := []int{127, 32767}
                    _ = v10
                }
                if ((-128 > -128) || false) {
                    if true {
                        for fi11 := 0; fi11 < 10; fi11++ {
                            for fi12 := 0; fi12 < 12; fi12++ {
                                v13 := GenS1{F0: -120, F1: -128, F2: 2147483647}
                                p1 = (p1 & p1)
                                _ = fi12
                                _ = v13
                            }
                            if true {
                                var v14 uint32 = 177
                                var v15 int32 = 58
                                _ = v14
                                _ = v15
                            }
                            _ = fi11
                        }
                    }
                    for fi16 := 0; fi16 < 7; fi16++ {
                        p0 += p0
                        v17 := GenS2{F0: -2147483648, F1: -1, F2: -1}
                        _ = fi16
                        _ = v17
                    }
                    v18 := []int{218}
                    _ = v18
                }
            } else {
                var v19 int = -903
                _ = v19
            }
        } else {
            var v20 uint8 = 79
            for fi21 := 0; fi21 < 2; fi21++ {
                if (18446744073709551615 >= (p1 << ((759) & 63))) {
                    var v22 int32 = -1
                    v23 := GenS3{F0: (("\n\t" + "a") != ("foo bar" + v8)), F1: 2147483647}
                    _ = v22
                    _ = v23
                }
                var v24 int16 = p0
                _ = fi21
                _ = v24
            }
            v25 := []int{32768, 4294967295, -909}
            _ = v20
            _ = v25
        }
        var v26 float64 = -2.25
        _ = v26
        _ = v8
    }
    v27 := []int{}
    var v28 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    _ = v28
    _ = v5
    return "hello"
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int16 = 396
    v3 := []int{9223372036854775807, 772, 9223372036854775807, 499}
    var v4 int32 = 884
    _ = len(v1)
    v4 = 637
    _ = len(v3)
    v29 := GenS1{F0: 42, F1: 0, F2: v4}
    var v30 int32 = (v4 - 2147483647)
    var v31 uint = 813
    var v32 uint8 = 128
    switch 42 {
    case 32768:
        var v33 uint16 = 0
        _ = v33
    case 2147483647:
        var v34 uint8 = 128
        _ = v34
    default:
        _ = len(v3)
    }
    if true {
        var v35 int64 = -1
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v36 uint16 = 255
        _ = v35
        _ = v36
    } else {
        v37 := GenS1{F0: -128, F1: -1, F2: int32(v32)}
        _ = v37
    }
    v38 := GenS1{F0: 42, F1: 0, F2: v4}
    v1[2147483647] = 65535
    for fi39 := 0; fi39 < 18; fi39++ {
        var v40 uint16 = 0
        fmt.Println(v30, v2, v31)
        _ = fi39
        _ = v40
    }
    fmt.Println(v32, v31)
    for fi41 := 0; fi41 < 6; fi41++ {
        for fi42 := 0; fi42 < 7; fi42++ {
            v2 = (v2 - v2)
            switch (("hello" + (("a" + (("\n\t" + ("" + ("hello" + "\n\t"))) + (("\n\t" + "a") + "foo bar"))) + ("foo bar" + "a"))) + "a") {
            case "hello":
                v43 := GenS4{F0: v31, F1: v2}
                var v44 int64 = 684
                v2 += v2
                switch "a" {
                case "a":
                    var v45 uint16 = 1
                    _ = v45
                default:
                    v32 = v32
                }
                _ = v43
                _ = v44
            case "foo bar":
                var v46 int16 = (v2 - int16(v31))
                _ = v46
            }
            switch "" {
            case "":
                switch fi41 {
                case 754:
                    v31 = 4294967295
                case 2147483647:
                    v47 := GenS4{F0: (v31 * v31), F1: v2}
                    _ = v47
                case -9223372036854775808:
                    delete(v1, 388)
                default:
                    var v48 int64 = -9223372036854775808
                    _ = v48
                }
            case "hello":
                v32 = (v32 << ((42) & 63))
            case "a":
                v4 += int32(fi41)
            default:
                var v49 uint8 = (v32 >> (((^(v31))) & 63))
                _ = v49
            }
            _ = fi42
        }
        _ = fi41
    }
    fmt.Println(v30)
    fmt.Println(v4)
    v4 += (v4 & v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v38
    _ = v4
}
