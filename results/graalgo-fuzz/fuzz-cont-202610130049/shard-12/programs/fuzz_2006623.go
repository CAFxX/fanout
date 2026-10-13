package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64; F1 uint; F2 int }
type GenS5 struct { F0 int8; F1 uint8 }
func genFn0(p0 uint32, p1 uint64) string {
    p0 -= p0
    v5 := []int{}
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := GenS1{F0: "hello"}
        switch (fi6 ^ (-(fi6))) {
        case 1:
            if true {
                v5 = append(v5, fi6)
                p1 += p1
            } else {
                p1 += (p1 - p1)
                if (42 == -96) {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    switch 32768 {
                    case 32767:
                        fmt.Println(p0, fi6)
                    }
                } else {
                    for fi8 := 0; fi8 < 12; fi8++ {
                        var v9 uint = 65535
                        _ = fi8
                        _ = v9
                    }
                    fmt.Println(p1)
                    fmt.Println(fi6, fi6)
                }
                if false {
                    v10 := GenS2{F0: ("\n\t" + "hello"), F1: 65535}
                    _ = v10
                }
            }
        }
        _ = fi6
        _ = v7
    }
    v11 := GenS1{F0: "a"}
    _ = p0
    _ = p1
    _ = v11
    _ = v5
    return ""
}
func genFn1() uint16 {
    if (!(!false)) {
        var v12 float32 = (0.1 / (float32(((-2.25 + (0.1 - 0.0)) / 1.0)) / 1.5))
        _ = v12
    }
    var v13 int8 = 42
    v14 := []int{}
    for fi15 := 0; fi15 < 15; fi15++ {
        if true {
            var v16 float64 = 0.1
            _ = v16
        } else {
            if false {
                var v17 float32 = 0.0
                _ = v17
            }
        }
        v13 += (^(v13))
        if false {
            v18 := GenS2{F0: "hello", F1: 18446744073709551615}
            fmt.Println(v13)
            v13 = (v13 - (v13 >> ((1) & 63)))
            _ = v18
        } else {
            v13 = v13
            _ = len(v14)
            for fi19 := 0; fi19 < 7; fi19++ {
                if ((!true) && false) {
                    v13 = 0
                    _ = len(v14)
                } else {
                    switch "\n\t" {
                    case "foo bar":
                        var v20 int8 = v13
                        _ = v20
                    case "":
                        v13 = v13
                    case "a":
                        var v21 int8 = 21
                        _ = v21
                    default:
                        v13 += -33
                    }
                    if true {
                        var v22 int8 = (v13 % (((v13 | (v13 >> ((1) & 63)))) | 1))
                        _ = v22
                    } else {
                        v13 = v13
                    }
                }
                var v23 uint = 2
                v23 -= v23
                v23 = (v23 * v23)
                if false {
                    v24 := GenS3{F0: 34}
                    v14 = append(v14, (fi15 | 2147483647))
                    v13 = 3
                    _ = v24
                } else {
                    if true {
                        if false {
                            var v25 uint16 = 348
                            _ = v25
                        }
                    }
                    v23 += (v23 & (v23 - (v23 * v23)))
                }
                _ = fi19
                _ = v23
            }
        }
        _ = fi15
    }
    _ = v13
    _ = v14
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = 1.5
    var v3 int64 = 42
    var v4 int8 = -1
    for fi26 := 0; fi26 < 12; fi26++ {
        switch "\n\t" {
        case "":
            var v27 int32 = -142
            _ = v27
        }
        _ = fi26
    }
    var v28 uint16 = 2
    switch ("\n\t" + ("\n\t" + "\n\t")) {
    case "foo bar":
        var v29 uint16 = (v28 & (v28 >> ((65535) & 63)))
        _ = v29
    case "hello":
        var v30 int8 = v4
        _ = v30
    default:
        var v31 uint32 = 4294967295
        _ = v31
    }
    v4 = 100
    v32 := v1[0]
    v33 := GenS1{F0: ("\n\t" + "hello")}
    v34 := []int{127, -1, 4294967295, -366}
    var v35 uint16 = 32767
    var v36 int = 9223372036854775807
    v37 := GenS4{F0: 2, F1: 42, F2: 1}
    v38 := GenS5{F0: (v4 / (((v4 - (v4 - v4))) | 1)), F1: 0}
    if v32 {
        var v39 uint8 = 127
        v32 = (!true)
        v35 -= v28
        _ = v39
    } else {
        for fi40 := 0; fi40 < 4; fi40++ {
            var v41 uint32 = uint32(v35)
            _ = fi40
            _ = v41
        }
    }
    v4 += v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
