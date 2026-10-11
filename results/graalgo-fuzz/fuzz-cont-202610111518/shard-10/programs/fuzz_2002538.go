package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 int32; F1 uint }
func genFn0(p0 uint, p1 uint8, p2 int32) int16 {
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 uint16 = 65535
        var v7 uint = 2
        _ = fi5
        _ = v6
        _ = v7
    }
    p0 = 0
    _ = p0
    _ = p1
    _ = p2
    return -32768
}
func genFn1(p0 uint) bool {
    for fi8 := 0; fi8 < 19; fi8++ {
        v9 := GenS1{F0: 32767}
        _ = fi8
        _ = v9
    }
    if ((589 == -2147483648) || (!false)) {
        for fi10 := 0; fi10 < 10; fi10++ {
            v11 := GenS1{F0: 1}
            _ = fi10
            _ = v11
        }
        p0 = (^(p0))
    }
    v12 := []int{128, 103, 42, 127}
    if len(v12) > 0 {
        v12[0] = 32767
    }
    p0 += 42
    _ = p0
    _ = v12
    return false
}

func main() {
    var v1 int32 = 0
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]int)
    v4 := []int{-1, 32767, 522, 2147483647}
    for fi13 := 0; fi13 < 19; fi13++ {
        for fi14 := 0; fi14 < 10; fi14++ {
            if (true || true) {
                var v15 uint8 = 42
                switch (("foo bar" + "foo bar") + "hello") {
                case "a":
                    var v16 int32 = 191
                    _ = v16
                }
                _ = v15
            } else {
                var v17 float64 = 1.5
                _ = v17
            }
            v18 := v3[42]
            if false {
                var v19 uint32 = 4294967295
                if (false && (1 == (-(v1)))) {
                    v20 := GenS2{F0: -1}
                    v1 = (v1 << ((255) & 63))
                    _ = v20
                } else {
                    fmt.Println(v1, v19, v1)
                }
                _ = v19
            }
            v21 := GenS3{F0: 32767, F1: (("\n\t" + "foo bar") + "\n\t")}
            _ = fi14
            _ = v18
            _ = v21
        }
        for fi22 := 0; fi22 < 13; fi22++ {
            v23 := GenS3{F0: 32767, F1: "hello"}
            v24 := GenS4{F0: 659, F1: 1}
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = fi13
    }
    v1 = (v1 ^ v1)
    fmt.Println(v1)
    v25 := v2[0]
    v26 := GenS4{F0: 32767, F1: 0}
    v25 = v25
    if len(v4) > 0 {
        v4[0] = -9223372036854775808
    }
    for fi27 := 0; fi27 < 18; fi27++ {
        var v28 bool = true
        for fi29 := 0; fi29 < 10; fi29++ {
            v30 := GenS4{F0: 0, F1: -57}
            v1 -= v1
            _ = fi29
            _ = v30
        }
        for fi31 := 0; fi31 < 19; fi31++ {
            if ("" >= "") {
                if false {
                    switch fi31 {
                    case -469:
                        switch ("" + "") {
                        case "\n\t":
                            var v32 uint32 = v25
                            _ = v32
                        }
                    case 127:
                        var v33 float32 = ((-2.25 / float32(0.1)) + (0.1 / float32(1e10)))
                        _ = v33
                    default:
                        v1 = v1
                    }
                    switch (-(fi31)) {
                    case 32768:
                        var v34 string = "\n\t"
                        _ = v34
                    case 32767:
                        _ = len(v2)
                    default:
                        v35 := v3[v25]
                        _ = v35
                    }
                } else {
                    var v36 uint64 = 698
                    var v37 int64 = -85
                    _ = v36
                    _ = v37
                }
                v28 = (true || (282 == 9223372036854775807))
                switch ("hello" + "hello") {
                case "a":
                    if v28 {
                        var v38 uint32 = (v25 ^ v25)
                        if (true && (!(42 < int64(v25)))) {
                            var v39 uint8 = 2
                            v40 := GenS5{F0: (v1 - (-(v1))), F1: uint(v38)}
                            _ = v39
                            _ = v40
                        } else {
                            delete(v3, v25)
                        }
                        _ = v38
                    }
                case "hello":
                    v1 = (v1 + 1)
                }
            } else {
                var v41 uint16 = 32767
                _ = v41
            }
            var v42 int64 = -374
            _ = fi31
            _ = v42
        }
        _ = fi27
        _ = v28
    }
    if (!false) {
        var v43 uint32 = v25
        var v44 int16 = -336
        switch 4294967295 {
        case 32767:
            v44 += v44
        }
        _ = v43
        _ = v44
    }
    switch 867 {
    case 691:
        switch ("foo bar" + "foo bar") {
        case "a":
            v25 = 42
        case "hello":
            switch 2147483647 {
            case 4294967295:
                _ = len(v4)
            default:
                var v45 bool = (v25 > (v25 >> ((4294967295) & 63)))
                _ = v45
            }
        case "":
            v46 := GenS1{F0: 1}
            _ = v46
        default:
            if false {
                v1 = 2147483647
                switch int(v25) {
                case 32768:
                    v47 := GenS1{F0: 1000}
                    _ = v47
                case 1:
                    fmt.Println(v25, v1)
                }
            } else {
                if true {
                    var v48 int8 = 1
                    for fi49 := 0; fi49 < 16; fi49++ {
                        v48 = 42
                        var v50 uint16 = 65535
                        _ = fi49
                        _ = v50
                    }
                    _ = v48
                } else {
                    if (false && true) {
                        var v51 int8 = -128
                        _ = v51
                    }
                }
            }
        }
    case 4294967295:
        var v52 string = ("hello" + "a")
        _ = v52
    default:
        v1 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
