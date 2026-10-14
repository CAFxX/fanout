package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 bool; F1 string; F2 bool }
type GenS3 struct { F0 bool; F1 string; F2 uint32 }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int; F1 int64 }
type GenS7 struct { F0 uint64; F1 uint; F2 bool }
type GenS8 struct { F0 int64; F1 bool }
type GenS9 struct { F0 uint32; F1 string; F2 uint32 }
func genFn0(p0 int8, p1 uint32) uint {
    if (1000 == 42) {
        v6 := GenS1{F0: 255, F1: p1}
        _ = v6
    } else {
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 int32 = 201
            p0 = p0
            var v9 float32 = 1.5
            _ = fi7
            _ = v8
            _ = v9
        }
        switch 4294967295 {
        case 9223372036854775807:
            var v10 int8 = -1
            _ = v10
        case -509:
            var v11 bool = true
            _ = v11
        }
        p0 = 0
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        p1 = (^(p1))
        var v13 float32 = float32((float64(-2.25) - 1e10))
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 bool = true
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: 128, F1: 0}
    _ = p0
    _ = p1
    _ = v16
    return 65535
}

func main() {
    v1 := []string{}
    v2 := []uint32{842, 2147483648}
    var v3 uint16 = 65535
    var v4 int32 = 2147483647
    var v5 int8 = 42
    if true {
        v17 := GenS2{F0: true, F1: "foo bar", F2: ((1.5 != 1.5) || true)}
        _ = v17
    }
    v5 += v5
    v3 += v3
    if true {
        v18 := GenS3{F0: true, F1: "\n\t", F2: 848}
        v4 -= v4
        _ = v18
    } else {
        if ("" >= "") {
            v4 += -1
        }
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        var v20 float32 = 0.0
        v21 := GenS3{F0: true, F1: ("\n\t" + "a"), F2: 2}
        _ = fi19
        _ = v20
        _ = v21
    }
    var v22 int32 = (v4 - (v4 << ((1) & 63)))
    v23 := GenS4{F0: true, F1: -32768}
    v3 = (v3 >> ((uint(v3)) & 63))
    if (v22 > v4) {
        var v24 uint32 = 2147483647
        if true {
            switch ("" + "hello") {
            case "foo bar":
                for fi25 := 0; fi25 < 14; fi25++ {
                    fmt.Println(v5)
                    var v26 int8 = (v5 % (((v5 & (v5 / ((v5) | 1)))) | 1))
                    _ = fi25
                    _ = v26
                }
            default:
                var v27 int8 = (v5 | v5)
                _ = v27
            }
        } else {
            v28 := GenS4{F0: false, F1: 936}
            fmt.Println(v3, v24)
            _ = v28
        }
        v4 = (v22 >> ((uint(v3)) & 63))
        _ = v24
    } else {
        var v29 string = ""
        if (65535 > 18446744073709551615) {
            switch 42 {
            case 42:
                fmt.Println(v22, v4)
            case 485:
                if len(v1) > 0 {
                    v1[0] = ""
                }
            default:
                v30 := GenS5{F0: 124}
                _ = v30
            }
            if (!(1 == 922)) {
                if false {
                    v31 := GenS6{F0: 32768, F1: -9223372036854775808}
                    for fi32 := 0; fi32 < 16; fi32++ {
                        var v33 int = (fi32 - (fi32 % (((fi32 | 42)) | 1)))
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        for fi34 := 0; fi34 < 16; fi34++ {
                            v29 = (((v29 + v29) + v29) + v29)
                            v35 := GenS7{F0: 4294967295, F1: 0, F2: true}
                            _ = fi34
                            _ = v35
                        }
                        _ = fi32
                        _ = v33
                    }
                    _ = v31
                } else {
                    v36 := GenS8{F0: 0, F1: false}
                    _ = v36
                }
                _ = len(v1)
            } else {
                for fi37 := 0; fi37 < 8; fi37++ {
                    v38 := GenS9{F0: 2, F1: "hello", F2: 2}
                    var v39 uint16 = v3
                    _ = fi37
                    _ = v38
                    _ = v39
                }
                for fi40 := 0; fi40 < 3; fi40++ {
                    for fi41 := 0; fi41 < 12; fi41++ {
                        var v42 int64 = 0
                        v22 += 0
                        _ = fi41
                        _ = v42
                    }
                    var v43 uint64 = 138
                    _ = fi40
                    _ = v43
                }
            }
            var v44 uint8 = 42
            _ = v44
        }
        v45 := GenS8{F0: 9223372036854775807, F1: true}
        _ = v29
        _ = v45
    }
    switch 128 {
    case 4294967295:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case 9223372036854775807:
        var v46 uint8 = 2
        _ = v46
    }
    v22 = (v4 << ((18446744073709551615) & 63))
    v3 += v3
    v47 := GenS1{F0: 14, F1: 161}
    var v48 float64 = 0.1
    v49 := GenS4{F0: (42 < int64(v3)), F1: -1}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
}
