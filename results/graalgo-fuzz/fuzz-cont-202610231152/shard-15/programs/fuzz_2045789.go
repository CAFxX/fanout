package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 bool; F1 int; F2 bool }
type GenS4 struct { F0 uint64; F1 bool }
func genFn0(p0 uint32) uint64 {
    fmt.Println(p0)
    for fi5 := 0; fi5 < 16; fi5++ {
        switch "" {
        case "hello":
            var v6 int32 = 1
            _ = v6
        case "a":
            switch ("foo bar" + ("foo bar" + (("\n\t" + "foo bar") + "hello"))) {
            case "hello":
                switch ("hello" + "\n\t") {
                case "":
                    p0 = p0
                case "hello":
                    p0 = (^(p0))
                }
            default:
                var v7 int = (fi5 * fi5)
                _ = v7
            }
        case "foo bar":
            var v8 float64 = 1e10
            _ = v8
        default:
            if (18446744073709551615 != 18446744073709551615) {
                var v9 float64 = 3.14159
                _ = v9
            }
        }
        var v10 uint8 = 128
        _ = fi5
        _ = v10
    }
    var v11 int8 = -109
    v11 = -128
    v12 := GenS1{F0: 42, F1: "a"}
    _ = p0
    _ = v11
    _ = v12
    return 18446744073709551615
}
func genFn1() int16 {
    v13 := []int{-112, -780, 1}
    var v14 string = ""
    var v15 string = ""
    _ = v13
    _ = v14
    _ = v15
    return 710
}

func main() {
    var v1 uint16 = 0
    var v2 int64 = 42
    v3 := make(map[uint32]bool)
    var v4 uint16 = v1
    if ((v2 / ((1) | 1)) != 0) {
        v4 -= v4
        v2 -= -270
        v16 := []int{0, -1, 0, 717}
        _ = v16
    } else {
        fmt.Println(v2, v1, v1)
        if false {
            fmt.Println(v1)
            v17 := []int{32767, -315}
            _ = v17
        } else {
            v1 += 0
        }
        v18 := GenS1{F0: 128, F1: ("foo bar" + "\n\t")}
        _ = v18
    }
    if false {
        v19 := []int{-2147483648, 42}
        v20 := GenS1{F0: 42, F1: "foo bar"}
        v21 := GenS2{F0: 516, F1: 1}
        _ = v19
        _ = v20
        _ = v21
    } else {
        v22 := GenS2{F0: 18446744073709551615, F1: 2}
        v2 = (v2 >> ((2) & 63))
        _ = v22
    }
    var v23 string = (("hello" + ("foo bar" + "hello")) + ("foo bar" + "hello"))
    var v24 int32 = -1
    var v25 int32 = (v24 << ((750) & 63))
    for fi26 := 0; fi26 < 4; fi26++ {
        if (3.14159 <= (float32(1e10) / ((3.14159 * 3.14159) - -2.25))) {
            v27 := v3[0]
            var v28 int8 = int8(v25)
            v29 := GenS2{F0: 42, F1: 42}
            _ = v27
            _ = v28
            _ = v29
        }
        _ = fi26
    }
    for fi30 := 0; fi30 < 20; fi30++ {
        var v31 float64 = 3.14159
        _ = fi30
        _ = v31
    }
    v25 = v25
    var v32 uint64 = 4294967295
    if (0 <= (^(v1))) {
        switch v23 {
        case "":
            v32 -= (v32 / ((0) | 1))
        case "hello":
            var v33 int = -2147483648
            _ = v33
        case "foo bar":
            if (!(true || (false && (127 <= 85)))) {
                v25 = v24
            } else {
                for fi34 := 0; fi34 < 6; fi34++ {
                    v2 = (v2 - (v2 ^ v2))
                    var v35 string = ("a" + (v23 + ((v23 + (v23 + (v23 + v23))) + "")))
                    var v36 uint32 = 2147483648
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            }
        }
        switch v23 {
        case "hello":
            var v37 int64 = (v2 - (v2 % (((v2 / ((645) | 1))) | 1)))
            _ = v37
        case "a":
            if true {
                if true {
                    v38 := GenS1{F0: 128, F1: v23}
                    v39 := GenS1{F0: -1, F1: "a"}
                    v40 := []int{1, 32767}
                    _ = v38
                    _ = v39
                    _ = v40
                } else {
                    v41 := GenS3{F0: (2147483647 < -1), F1: -135, F2: false}
                    v2 += (v2 & v2)
                    _ = v41
                }
                var v42 uint64 = genFn0(42)
                v43 := []int{129}
                _ = v42
                _ = v43
            } else {
                v2 = -928
                var v44 uint16 = (v4 >> ((1) & 63))
                _ = v44
            }
        default:
            for fi45 := 0; fi45 < 19; fi45++ {
                v25 += 339
                fmt.Println(v32, v32, v2)
                switch fi45 {
                case 32768:
                    var v46 bool = false
                    _ = v46
                }
                _ = fi45
            }
        }
        v47 := v3[1]
        _ = v47
    } else {
        var v48 int8 = -34
        for fi49 := 0; fi49 < 3; fi49++ {
            var v50 int = (fi49 & (fi49 & int(v24)))
            v51 := GenS2{F0: 1, F1: 1}
            for fi52 := 0; fi52 < 3; fi52++ {
                var v53 int8 = -1
                v54 := []int{(v50 & fi49)}
                _ = fi52
                _ = v53
                _ = v54
            }
            _ = fi49
            _ = v50
            _ = v51
        }
        _ = v48
    }
    if (genFn1() >= 1) {
        for fi55 := 0; fi55 < 11; fi55++ {
            v56 := GenS2{F0: 848, F1: 65535}
            _ = fi55
            _ = v56
        }
        v4 = v4
    }
    switch v23 {
    case "\n\t":
        for fi57 := 0; fi57 < 10; fi57++ {
            v24 -= v24
            var v58 float64 = 3.14159
            var v59 int = (fi57 + (-(fi57)))
            _ = fi57
            _ = v58
            _ = v59
        }
    case "":
        v25 += (v24 & v25)
    default:
        v60 := GenS4{F0: 0, F1: (!(!(true && (!true))))}
        _ = v60
    }
    v23 = (((((v23 + "") + v23) + "a") + v23) + v23)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v32
    _ = v4
}
