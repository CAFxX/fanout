package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS4 struct { F0 int8; F1 uint32 }
type GenS5 struct { F0 uint64 }
func genFn0() int {
    var v4 uint = 65535
    if true {
        var v5 int8 = 100
        fmt.Println(v5)
        _ = v5
    }
    var v6 uint32 = 369
    var v7 uint16 = 537
    _ = v4
    _ = v6
    _ = v7
    return 630
}
func genFn1(p0 string) int8 {
    v8 := []int{32768, 32767, -1}
    p0 = p0
    p0 = p0
    _ = p0
    _ = v8
    return 0
}

func main() {
    var v1 int64 = 0
    var v2 int64 = (v1 - 139)
    var v3 int8 = 1
    var v9 float64 = (1.5 / ((3.14159 + 1e10) * (0.1 - 3.14159)))
    var v10 int16 = 183
    var v11 uint32 = 4294967295
    v10 = 42
    if (18446744073709551615 >= 42) {
        if true {
            v12 := []int{2147483647, -1, -1, -1}
            var v13 uint64 = 4294967295
            fmt.Println(v9, v10, v9)
            _ = v12
            _ = v13
        } else {
            switch "\n\t" {
            case "\n\t":
                for fi14 := 0; fi14 < 2; fi14++ {
                    fmt.Println(v1, v3, v3)
                    var v15 uint32 = (v11 << ((1) & 63))
                    v16 := GenS1{F0: true, F1: (^(v10)), F2: ("foo bar" + "")}
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            case "hello":
                v1 = -1
            case "foo bar":
                if (325 == 1) {
                    if false {
                        var v17 uint16 = 744
                        for fi18 := 0; fi18 < 13; fi18++ {
                            for fi19 := 0; fi19 < 6; fi19++ {
                                v20 := GenS2{F0: (v10 * 81)}
                                if (0 == 220) {
                                    v9 += v9
                                    var v21 uint64 = 42
                                    v22 := []int{9223372036854775807, (fi18 << ((65535) & 63)), fi19, (-(fi18))}
                                    v11 = v11
                                    v3 -= (^(v3))
                                    _ = v21
                                    _ = v22
                                }
                                _ = fi19
                                _ = v20
                            }
                            _ = fi18
                        }
                        v2 -= v1
                        _ = v17
                    } else {
                        v23 := []int{genFn0(), -946, genFn0(), genFn0()}
                        _ = v23
                    }
                    v10 = v10
                }
            default:
                fmt.Println(v3)
            }
        }
    } else {
        for fi24 := 0; fi24 < 20; fi24++ {
            v3 += (v3 - (v3 * (v3 / ((genFn1(("foo bar" + "hello"))) | 1))))
            fmt.Println(v3, v10, v2)
            if (42 < 42) {
                v25 := GenS3{F0: 32767, F1: int64(v11), F2: 549}
                _ = v25
            }
            _ = fi24
        }
        v26 := GenS4{F0: (v3 / (((v3 % ((-128) | 1))) | 1)), F1: (v11 | (v11 >> ((2) & 63)))}
        v1 = (v2 + (v1 - 0))
        _ = v26
    }
    for fi27 := 0; fi27 < 15; fi27++ {
        if false {
            fmt.Println(v11)
            v28 := GenS2{F0: 346}
            var v29 int = (fi27 % ((fi27) | 1))
            _ = v28
            _ = v29
        } else {
            v30 := GenS1{F0: false, F1: v10, F2: ("\n\t" + ("" + "foo bar"))}
            _ = v30
        }
        fmt.Println(fi27, v11)
        switch "foo bar" {
        case "a":
            fmt.Println(v2, v1, v10)
        case "\n\t":
            v31 := GenS5{F0: 0}
            _ = v31
        case "foo bar":
            if true {
                var v32 int32 = -397
                _ = v32
            }
        }
        _ = fi27
    }
    fmt.Println(v9)
    v10 = int16(v11)
    v33 := []int{}
    fmt.Println(v1, v11, v1)
    v10 += 1
    switch (("foo bar" + "a") + "foo bar") {
    case "":
        v33 = append(v33, 127)
    case "hello":
        v10 = -61
    case "a":
        if true {
            v10 = (v10 * (v10 / (((v10 + -455)) | 1)))
            var v34 uint = 1
            _ = v34
        } else {
            var v35 int32 = -2147483648
            _ = v35
        }
    }
    fmt.Println(v11)
    v36 := GenS1{F0: ("a" < "foo bar"), F1: v10, F2: "foo bar"}
    var v37 uint64 = 17
    var v38 int64 = (v2 - v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v33
    _ = v36
    _ = v37
    _ = v38
    _ = v9
}
