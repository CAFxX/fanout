package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint8 }
type GenS2 struct { F0 int; F1 uint16; F2 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 int8; F2 string }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 1
    v3 := make(map[string]int)
    v4 := make(map[uint32]uint32)
    v5 := make(map[string]int)
    v2 += 40
    fmt.Println(v2)
    fmt.Println(v2)
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 string = "a"
        switch v7 {
        case "hello":
            v8 := GenS1{F0: (("foo bar" + v7) + v7), F1: 18446744073709551615, F2: 1}
            _ = v8
        case "\n\t":
            fmt.Println(v7)
        case "foo bar":
            v2 -= (-(v2))
        }
        v7 = ("\n\t" + "\n\t")
        _ = fi6
        _ = v7
    }
    var v9 int16 = -32768
    v10 := v3[(("a" + "") + "hello")]
    v2 = -128
    if false {
        v11 := GenS1{F0: ("" + "a"), F1: 0, F2: 214}
        switch "hello" {
        case "a":
            v12 := []int{(v10 - 32768), (v10 << ((65535) & 63)), 4294967295}
            _ = v12
        }
        if (-2147483648 > -15) {
            v10 = 614
        } else {
            for fi13 := 0; fi13 < 17; fi13++ {
                v14 := []int{}
                _ = fi13
                _ = v14
            }
            delete(v1, 2147483648)
        }
        _ = v11
    } else {
        for fi15 := 0; fi15 < 5; fi15++ {
            fmt.Println(fi15)
            if (0 == (v10 % (((v10 | fi15)) | 1))) {
                var v16 string = "hello"
                if (0 >= int64(v10)) {
                    v9 = (v9 ^ int16(fi15))
                    v9 = v9
                }
                _ = v16
            } else {
                v17 := GenS2{F0: v10, F1: 1, F2: 177}
                if ((1e10 != (1e10 / 1.0)) || false) {
                    for fi18 := 0; fi18 < 1; fi18++ {
                        switch (("\n\t" + (("foo bar" + "") + ((("" + "a") + ("foo bar" + ((("a" + (("foo bar" + "a") + (("" + "\n\t") + "foo bar"))) + "") + "hello"))) + ""))) + "a") {
                        case "hello":
                            v3[""] = (^(fi18))
                        default:
                            var v19 uint = 233
                            _ = v19
                        }
                        v9 -= v9
                        v20 := GenS3{F0: "a"}
                        _ = fi18
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 16; fi21++ {
                        if false {
                            v22 := []int{(-(fi21)), -2147483648}
                            fmt.Println(v9, v9)
                            var v23 int64 = 217
                            _ = v22
                            _ = v23
                        }
                        switch (fi15 | 642) {
                        case 32768:
                            if (673 != 65535) {
                                v9 += (v9 << ((0) & 63))
                                if ((v2 | v2) == -1) {
                                    var v24 uint8 = 2
                                    if (42 < 0) {
                                        v25 := GenS4{F0: ((v2 < (v2 % ((0) | 1))) || false), F1: v2, F2: (("hello" + ("" + ("hello" + "foo bar"))) + "")}
                                        _ = v25
                                    } else {
                                        if (true || true) {
                                            var v26 bool = false
                                            v27 := GenS5{F0: 42}
                                            v28 := GenS6{F0: 4294967295}
                                            var v29 uint16 = 0
                                            switch "hello" {
                                            case "":
                                                if v26 {
                                                    var v30 bool = v26
                                                    _ = v30
                                                }
                                            case "foo bar":
                                                v29 = 2
                                            case "\n\t":
                                                var v31 bool = ((fi15 + 588) == v10)
                                                _ = v31
                                            default:
                                                v32 := GenS3{F0: "\n\t"}
                                                _ = v32
                                            }
                                            _ = v26
                                            _ = v27
                                            _ = v28
                                            _ = v29
                                        } else {
                                            v10 -= fi21
                                            if false {
                                                fmt.Println(v24)
                                                var v33 bool = (762 >= 1)
                                                v34 := GenS7{F0: 2}
                                                _ = v33
                                                _ = v34
                                            } else {
                                                v35 := []int{-565, (^(fi15))}
                                                v10 -= v10
                                                _ = v35
                                            }
                                            var v36 float64 = ((0.1 * (((-2.25 + 0.1) + (0.0 + 1e10)) * 1.5)) + ((0.1 / 3.14159) + 0.1))
                                            _ = v36
                                        }
                                        v37 := []int{}
                                        var v38 int16 = (^(v9))
                                        _ = v37
                                        _ = v38
                                    }
                                    _ = v24
                                }
                                var v39 uint32 = 939
                                _ = v39
                            } else {
                                _ = len(v3)
                            }
                        case 42:
                            var v40 int32 = 1
                            _ = v40
                        }
                        _ = fi21
                    }
                    var v41 uint64 = 18446744073709551615
                    _ = v41
                } else {
                    v42 := GenS6{F0: 65535}
                    v10 += 634
                    _ = v42
                }
                _ = v17
            }
            var v43 int = (fi15 ^ (fi15 * -747))
            _ = fi15
            _ = v43
        }
    }
    v44 := v5["hello"]
    switch "\n\t" {
    case "a":
        v45 := []int{v44, v10}
        _ = v45
    default:
        if true {
            fmt.Println(v44, v9)
        }
    }
    for fi46 := 0; fi46 < 18; fi46++ {
        var v47 float32 = (3.14159 - (float32(((1.5 - 0.1) * -2.25)) - ((3.14159 - 0.0) + (float32(float64(0.0)) - 1.5))))
        _ = fi46
        _ = v47
    }
    switch 32768 {
    case -1:
        var v48 float32 = 1e10
        _ = v48
    }
    v49 := GenS2{F0: int(v9), F1: 255, F2: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v49
    _ = v5
    _ = v9
}
