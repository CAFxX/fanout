package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 int32; F2 int8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint16; F1 uint64; F2 uint32 }
type GenS7 struct { F0 int8; F1 int32; F2 string }
func genFn0(p0 uint, p1 int) string {
    var v7 uint16 = 1
    if false {
        p0 += p0
        for fi8 := 0; fi8 < 1; fi8++ {
            v9 := []int{int(p0), (p1 ^ (^(p1))), (-(p1)), (p1 << (((p0 - p0)) & 63))}
            var v10 int8 = 1
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    v11 := []int{p1}
    for fi12 := 0; fi12 < 8; fi12++ {
        v11 = append(v11, -436)
        switch (("a" + ("" + "a")) + "") {
        case "":
            for fi13 := 0; fi13 < 15; fi13++ {
                if false {
                    var v14 int8 = int8(v7)
                    var v15 int = fi13
                    v14 -= v14
                    _ = v14
                    _ = v15
                } else {
                    v16 := GenS1{F0: (-(p1))}
                    v17 := GenS2{F0: (((("hello" + "hello") + (("hello" + "\n\t") + "hello")) + "\n\t") + ("hello" + "foo bar"))}
                    _ = v16
                    _ = v17
                }
                _ = fi13
            }
        case "\n\t":
            p0 = p0
        case "foo bar":
            p1 -= fi12
        default:
            if false {
                v18 := GenS2{F0: "hello"}
                p1 -= fi12
                _ = v18
            } else {
                v7 = v7
                if (!false) {
                    var v19 int = (fi12 ^ fi12)
                    v19 += (p1 | (fi12 << (((p0 >> ((p0) & 63))) & 63)))
                    _ = v19
                }
                if (!(false || true)) {
                    if true {
                        v7 = 42
                        var v20 uint = (p0 >> ((1) & 63))
                        _ = v20
                    } else {
                        v21 := GenS3{F0: 0}
                        v7 -= 65535
                        v22 := GenS4{F0: false, F1: -2147483648, F2: -1}
                        _ = v21
                        _ = v22
                    }
                    if (v7 >= (v7 - v7)) {
                        p0 += p0
                    }
                } else {
                    var v23 float64 = float64(((1e10 / (float32(3.14159) - -2.25)) + (1e10 + -2.25)))
                    fmt.Println(v23, v7)
                    _ = v23
                }
            }
        }
        v24 := GenS4{F0: false, F1: 1, F2: -6}
        _ = fi12
        _ = v24
    }
    var v25 bool = false
    _ = p0
    _ = p1
    _ = v11
    _ = v25
    _ = v7
    return ("a" + "a")
}
func genFn1(p0 int32, p1 int16) uint8 {
    var v26 int64 = -9223372036854775808
    p0 = (-(p0))
    var v27 int16 = p1
    for fi28 := 0; fi28 < 7; fi28++ {
        var v29 int64 = (-(v26))
        for fi30 := 0; fi30 < 5; fi30++ {
            var v31 int32 = (p0 << ((1) & 63))
            fmt.Println(v26, p0, v29)
            var v32 uint16 = 65535
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = fi28
        _ = v29
    }
    var v33 int16 = int16(v26)
    _ = p0
    _ = p1
    _ = v26
    _ = v27
    _ = v33
    return 128
}

func main() {
    var v1 int64 = 42
    var v2 uint64 = 678
    var v3 int16 = 1000
    var v4 int = 32768
    v5 := make(map[string]bool)
    var v6 uint16 = 32767
    var v34 int32 = 1
    v34 = v34
    switch "\n\t" {
    case "foo bar":
        if false {
            for fi35 := 0; fi35 < 2; fi35++ {
                v6 = v6
                if ((!(!(v2 != 4294967295))) && ((0.0 * 1.5) <= float64(((float32((0.0 + float64(((-2.25 + 3.14159) * (0.1 * 0.0))))) - float32((3.14159 + ((-2.25 * (3.14159 / 1e10)) / 1.0)))) + ((0.1 / -2.25) - 3.14159))))) {
                    if false {
                        v4 = (v4 ^ fi35)
                    } else {
                        v5[""] = true
                        var v36 uint64 = (v2 % ((v2) | 1))
                        _ = v36
                    }
                    switch int(v2) {
                    case -9223372036854775808:
                        v2 = v2
                    case -624:
                        var v37 int8 = 0
                        _ = v37
                    }
                }
                _ = fi35
            }
        }
    }
    v38 := []int{v4, (v4 >> ((0) & 63)), (v4 ^ (v4 % ((v4) | 1)))}
    var v39 int = (v4 + (v4 / ((int(v34)) | 1)))
    if true {
        v40 := GenS1{F0: v39}
        v41 := GenS5{F0: 42}
        _ = v40
        _ = v41
    }
    var v42 uint64 = v2
    v43 := GenS6{F0: 1, F1: 1, F2: 2147483647}
    var v44 int = (v39 << ((0) & 63))
    var v45 uint32 = 0
    var v46 int32 = v34
    var v47 uint = 42
    if true {
        _ = len(v5)
    }
    v48 := GenS7{F0: -14, F1: v46, F2: (("" + "\n\t") + "")}
    for fi49 := 0; fi49 < 4; fi49++ {
        v2 = v2
        _ = len(v38)
        for fi50 := 0; fi50 < 10; fi50++ {
            var v51 uint64 = (v42 ^ 2)
            _ = fi50
            _ = v51
        }
        _ = fi49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v38
    _ = v39
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
