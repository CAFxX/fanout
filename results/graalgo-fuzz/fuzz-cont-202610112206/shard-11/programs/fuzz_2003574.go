package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 bool; F1 int8; F2 bool }
type GenS6 struct { F0 uint64 }
func genFn0() uint64 {
    var v5 int = -935
    v5 += v5
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 uint64 = 4294967295
        for fi8 := 0; fi8 < 2; fi8++ {
            fmt.Println(v7, fi6, fi6)
            if false {
                switch fi8 {
                case 128:
                    v9 := GenS1{F0: -544}
                    _ = v9
                case -9223372036854775808:
                    v10 := GenS1{F0: -789}
                    _ = v10
                case 9223372036854775807:
                    v7 = (v7 << ((255) & 63))
                    v11 := []int{v5}
                    fmt.Println(v5, v5)
                    _ = v11
                }
                v12 := []int{42, (fi6 >> ((1) & 63))}
                var v13 uint16 = 2
                _ = v12
                _ = v13
            } else {
                if (float32(0.0) != 1e10) {
                    switch "\n\t" {
                    case "\n\t":
                        if true {
                            v7 -= (^(v7))
                        }
                    case "a":
                        var v14 uint16 = 65535
                        _ = v14
                    case "hello":
                        var v15 uint = 2
                        _ = v15
                    }
                    v16 := GenS1{F0: 0}
                    _ = v16
                } else {
                    var v17 bool = false
                    v17 = true
                    switch (("foo bar" + "hello") + (("hello" + "foo bar") + "\n\t")) {
                    case "\n\t":
                        if (1e10 <= ((0.1 + 1.5) * -2.25)) {
                            v7 -= 4294967295
                        }
                    case "":
                        v18 := GenS1{F0: 1000}
                        _ = v18
                    case "hello":
                        fmt.Println(fi8, fi8)
                    }
                    _ = v17
                }
                v19 := GenS2{F0: "\n\t", F1: "hello"}
                v7 = 486
                _ = v19
            }
            _ = fi8
        }
        var v20 string = ("hello" + "")
        _ = fi6
        _ = v20
        _ = v7
    }
    _ = v5
    return 634
}

func main() {
    var v1 float64 = 0.0
    var v2 int8 = -1
    var v3 uint16 = 0
    var v4 bool = true
    for fi21 := 0; fi21 < 3; fi21++ {
        var v22 uint32 = 904
        if v4 {
            v4 = true
            v22 += v22
            for fi23 := 0; fi23 < 20; fi23++ {
                v4 = (0 >= 1)
                _ = fi23
            }
        } else {
            v4 = v4
            for fi24 := 0; fi24 < 3; fi24++ {
                switch fi24 {
                case 4294967295:
                    var v25 int64 = -1
                    _ = v25
                case 0:
                    v26 := GenS1{F0: 1}
                    _ = v26
                case 42:
                    var v27 int = 128
                    _ = v27
                default:
                    v28 := []int{(fi21 << ((457) & 63)), fi21, fi24, (^(fi24))}
                    _ = v28
                }
                _ = fi24
            }
            for fi29 := 0; fi29 < 5; fi29++ {
                if true {
                    if v4 {
                        var v30 int64 = 1
                        var v31 int64 = v30
                        switch 1 {
                        case -1:
                            v32 := []int{128, -320}
                            _ = v32
                        default:
                            v3 = 1
                        }
                        _ = v30
                        _ = v31
                    }
                    var v33 int16 = -1
                    fmt.Println(fi21, v1)
                    _ = v33
                } else {
                    var v34 float64 = v1
                    _ = v34
                }
                _ = fi29
            }
        }
        v35 := GenS3{F0: 248}
        _ = fi21
        _ = v22
        _ = v35
    }
    var v36 int = 528
    v37 := GenS4{F0: (-32768 <= -1)}
    v38 := GenS5{F0: ("a" >= "\n\t"), F1: v2, F2: v4}
    var v39 float32 = 3.14159
    switch (v36 ^ v36) {
    case 32767:
        v4 = false
    }
    for fi40 := 0; fi40 < 5; fi40++ {
        for fi41 := 0; fi41 < 15; fi41++ {
            fmt.Println(v39, fi41, fi41)
            if (v4 || true) {
                var v42 uint = 4294967295
                _ = v42
            } else {
                v3 += (^(v3))
                v4 = v4
            }
            v43 := GenS6{F0: 780}
            fmt.Println(v1, v2, fi40)
            var v44 uint8 = 1
            _ = fi41
            _ = v43
            _ = v44
        }
        var v45 string = ""
        v46 := []int{(v36 / (((v36 | (v36 * 42))) | 1)), (v36 & (fi40 / (((fi40 % ((fi40) | 1))) | 1))), 913, 127}
        _ = fi40
        _ = v45
        _ = v46
    }
    var v47 int16 = 1
    switch (v36 >> ((0) & 63)) {
    case -701:
        v47 -= (-(v47))
    }
    v48 := []int{127, (^(v36))}
    fmt.Println(v2, v1)
    fmt.Println(v39, v47, v4)
    v36 -= v36
    v36 = 32768
    for fi49 := 0; fi49 < 2; fi49++ {
        var v50 uint64 = 4294967295
        v3 -= 0
        var v51 string = (("hello" + "a") + ("foo bar" + ("" + "")))
        _ = fi49
        _ = v50
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v48:", len(v48), v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v47
    _ = v48
}
