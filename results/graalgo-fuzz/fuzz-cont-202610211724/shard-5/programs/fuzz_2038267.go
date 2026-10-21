package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 string; F2 int64 }
type GenS4 struct { F0 int64; F1 int32; F2 int }
type GenS5 struct { F0 int32; F1 uint }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint32; F1 int16 }
type GenS8 struct { F0 uint; F1 int }
func genFn0(p0 uint64) uint {
    v5 := GenS1{F0: 128}
    v6 := GenS2{F0: 42}
    if true {
        if true {
            var v7 float64 = 1.5
            v8 := []int{}
            _ = v7
            _ = v8
        } else {
            switch 4294967295 {
            case -999:
                p0 = (p0 | 663)
            case 633:
                v9 := GenS3{F0: 2, F1: "foo bar", F2: 42}
                _ = v9
            default:
                v10 := GenS4{F0: 9223372036854775807, F1: 42, F2: 0}
                _ = v10
            }
            for fi11 := 0; fi11 < 6; fi11++ {
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(p0)
                case "":
                    var v12 uint8 = 255
                    _ = v12
                default:
                    for fi13 := 0; fi13 < 16; fi13++ {
                        v14 := GenS5{F0: 0, F1: 2}
                        switch -934 {
                        case 42:
                            var v15 uint16 = 255
                            _ = v15
                        case 652:
                            v16 := []int{fi13, (fi13 | 741)}
                            _ = v16
                        case -587:
                            if (514 <= 580) {
                                p0 = (p0 + (p0 & p0))
                                var v17 float64 = 3.14159
                                _ = v17
                            }
                        default:
                            var v18 uint8 = 42
                            _ = v18
                        }
                        var v19 float32 = 0.1
                        _ = fi13
                        _ = v14
                        _ = v19
                    }
                }
                _ = fi11
            }
            var v20 float32 = 1.5
            _ = v20
        }
        if (64 == 42) {
            if (!(57 <= 100)) {
                fmt.Println(p0)
                if false {
                    p0 = (p0 | (p0 - (p0 * p0)))
                } else {
                    v21 := []int{32767}
                    fmt.Println(p0)
                    _ = v21
                }
            } else {
                v22 := []int{0, -9223372036854775808, -9223372036854775808}
                var v23 uint32 = 654
                _ = v22
                _ = v23
            }
        }
        p0 -= (p0 % ((p0) | 1))
    } else {
        var v24 uint64 = (p0 | p0)
        _ = v24
    }
    _ = p0
    _ = v5
    _ = v6
    return 42
}
func genFn1() uint16 {
    var v25 int = 127
    for fi26 := 0; fi26 < 1; fi26++ {
        v27 := GenS6{F0: 787}
        _ = fi26
        _ = v27
    }
    var v28 int64 = 42
    v29 := []int{v25, v25, (v25 % (((v25 % (((-(v25))) | 1))) | 1))}
    _ = v25
    _ = v28
    _ = v29
    return 255
}

func main() {
    v1 := []uint32{2147483647, 2}
    var v2 int16 = -32768
    v3 := make(map[string]string)
    var v4 int16 = v2
    for fi30 := 0; fi30 < 7; fi30++ {
        if (874 >= 1) {
            if (18446744073709551615 == genFn0(18446744073709551615)) {
                for fi31 := 0; fi31 < 10; fi31++ {
                    switch 9223372036854775807 {
                    case 0:
                        v2 = (v2 & (-(v2)))
                    case 371:
                        v32 := GenS2{F0: -34}
                        _ = v32
                    case 127:
                        var v33 string = "foo bar"
                        _ = v33
                    }
                    if (true || false) {
                        v34 := GenS7{F0: 597, F1: (v2 % ((v2) | 1))}
                        v35 := GenS4{F0: 0, F1: 1, F2: 259}
                        _ = v34
                        _ = v35
                    }
                    if (!false) {
                        if false {
                            v4 += (^(v2))
                            v2 = (-(v4))
                        }
                        if true {
                            fmt.Println(fi31, fi30, v4)
                        } else {
                            v4 += (-(v2))
                            var v36 uint8 = 133
                            _ = v36
                        }
                        v4 = (v4 & (v2 >> ((65535) & 63)))
                    }
                    _ = fi31
                }
                v2 = -32768
            } else {
                if (-2.25 <= 0.1) {
                    var v37 int8 = -103
                    _ = v37
                } else {
                    v1 = append(v1, 1)
                    _ = len(v3)
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v4 += v4
            }
            var v38 uint32 = 0
            _ = v38
        }
        v3["hello"] = "foo bar"
        v39 := GenS8{F0: 1, F1: (fi30 + (fi30 | (fi30 * fi30)))}
        _ = fi30
        _ = v39
    }
    v2 = (^(v2))
    delete(v3, (("" + (("" + "hello") + "\n\t")) + (("foo bar" + ("hello" + "")) + ("\n\t" + ((("a" + "hello") + "\n\t") + "a")))))
    v2 = -32768
    v4 = (v4 >> ((genFn0(2)) & 63))
    v40 := v3["hello"]
    switch (("a" + ("" + v40)) + "hello") {
    case "\n\t":
        switch 128 {
        case -963:
            fmt.Println(v40, v4)
        default:
            for fi41 := 0; fi41 < 12; fi41++ {
                var v42 uint64 = 0
                _ = fi41
                _ = v42
            }
        }
    }
    v43 := GenS1{F0: 127}
    if len(v1) > 0 {
        v1[0] = 87
    }
    switch (v40 + (((v40 + v40) + ("hello" + "a")) + "hello")) {
    case "foo bar":
        if (true && false) {
            fmt.Println(v40, v40, v40)
        }
    case "hello":
        fmt.Println(v2, v2)
    case "a":
        v2 += v4
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v43
}
