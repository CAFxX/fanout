package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool; F1 int16 }
type GenS6 struct { F0 string; F1 uint8 }
func genFn0(p0 int16, p1 uint8) int16 {
    p1 = 2
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    return (p0 * p0)
}
func genFn1(p0 uint64, p1 int16) int32 {
    var v7 uint16 = 255
    var v8 int32 = -2147483648
    v9 := GenS1{F0: true, F1: -1, F2: 63}
    v10 := GenS1{F0: (!((-2.25 - (-2.25 * 1.5)) == 0.1)), F1: genFn0((p1 << ((76) & 63)), 127), F2: 107}
    for fi11 := 0; fi11 < 18; fi11++ {
        if (!true) {
            v12 := GenS1{F0: true, F1: p1, F2: 0}
            var v13 int64 = -1
            v14 := GenS1{F0: true, F1: (p1 / ((p1) | 1)), F2: 202}
            _ = v12
            _ = v13
            _ = v14
        } else {
            if (v8 < (v8 % ((v8) | 1))) {
                switch ((("\n\t" + (("" + "hello") + "\n\t")) + ("a" + "foo bar")) + "foo bar") {
                case "\n\t":
                    var v15 string = "a"
                    _ = v15
                case "":
                    v16 := GenS2{F0: (p1 | (p1 * -32768))}
                    _ = v16
                case "foo bar":
                    v17 := GenS3{F0: (p1 >> ((255) & 63))}
                    _ = v17
                default:
                    p0 = 677
                }
                var v18 int64 = int64(p0)
                _ = v18
            } else {
                v7 = v7
                fmt.Println(fi11)
                v19 := GenS3{F0: (p1 / ((genFn0((^(p1)), 0)) | 1))}
                _ = v19
            }
        }
        p0 = p0
        v7 = (^(v7))
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return v8
}

func main() {
    var v1 uint64 = 4294967295
    var v2 uint = 964
    var v3 uint8 = 1
    v4 := make(map[int]int)
    v5 := make(map[uint32]uint32)
    var v6 bool = false
    if (uint32(v2) != 2147483647) {
        v1 += (^(v1))
        var v20 int32 = 348
        _ = v20
    }
    v21 := GenS4{F0: 0}
    if (("foo bar" < (("" + ("\n\t" + ("hello" + ("foo bar" + "foo bar")))) + "")) && (!((v1 << ((80) & 63)) != v1))) {
        var v22 float64 = ((0.0 / 0.1) * -2.25)
        _ = v22
    } else {
        var v23 bool = false
        switch 43 {
        case 128:
            v24 := v4[4294967295]
            _ = v24
        case 554:
            fmt.Println(v2, v3, v1)
        case 1:
            v25 := []int{-1, 245}
            _ = v25
        default:
            if (v2 == (v2 & (v2 * (v2 - v2)))) {
                var v26 uint = 255
                _ = v26
            }
        }
        _ = v23
    }
    _ = len(v4)
    v27 := GenS4{F0: 42}
    for fi28 := 0; fi28 < 12; fi28++ {
        if true {
            v29 := GenS4{F0: -323}
            _ = v29
        } else {
            fmt.Println(v2, v1)
            fmt.Println(v6, fi28, v6)
        }
        _ = fi28
    }
    switch "foo bar" {
    case "a":
        v30 := GenS2{F0: 1000}
        _ = v30
    }
    v31 := []int{9223372036854775807, -9223372036854775808, 761}
    fmt.Println(v1, v6, v6)
    if (4294967295 >= 2) {
        var v32 bool = (32767 >= 255)
        if (v6 || (92 != 100)) {
            fmt.Println(v1, v3, v6)
            v33 := GenS5{F0: false, F1: 42}
            if (v32 || false) {
                fmt.Println(v3, v6)
                var v34 int8 = -1
                for fi35 := 0; fi35 < 16; fi35++ {
                    v32 = (42 == (-(v34)))
                    v34 += v34
                    v36 := GenS3{F0: 140}
                    _ = fi35
                    _ = v36
                }
                _ = v34
            } else {
                fmt.Println(v32, v2, v1)
            }
            _ = v33
        } else {
            v37 := GenS5{F0: v32, F1: genFn0(genFn0(630, 0), v3)}
            v31 = append(v31, 2147483647)
            _ = v37
        }
        var v38 uint8 = v3
        _ = v32
        _ = v38
    } else {
        for fi39 := 0; fi39 < 6; fi39++ {
            var v40 int8 = -128
            v6 = (42 >= v40)
            _ = fi39
            _ = v40
        }
        var v41 uint16 = 2
        _ = v41
    }
    var v42 string = "foo bar"
    delete(v4, 128)
    var v43 int64 = 42
    for fi44 := 0; fi44 < 15; fi44++ {
        if (v6 && false) {
            if len(v31) >= 2 {
                v31 = v31[1:len(v31)]
            }
            if (!(!v6)) {
                v45 := GenS6{F0: ("\n\t" + ""), F1: v3}
                v42 = "hello"
                var v46 uint32 = 65535
                _ = v45
                _ = v46
            } else {
                v3 -= 127
                fmt.Println(v42, v2)
            }
            fmt.Println(fi44)
        }
        _ = fi44
    }
    for fi47 := 0; fi47 < 1; fi47++ {
        for fi48 := 0; fi48 < 10; fi48++ {
            var v49 uint16 = 255
            fmt.Println(v42, fi47)
            v50 := GenS1{F0: true, F1: -1, F2: v3}
            _ = fi48
            _ = v49
            _ = v50
        }
        v2 = 369
        _ = fi47
    }
    var v51 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v51
    _ = v6
}
