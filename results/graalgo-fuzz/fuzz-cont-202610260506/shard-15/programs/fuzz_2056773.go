package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int8 }
type GenS2 struct { F0 int8; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 int32; F1 int16 }
func genFn0(p0 int8, p1 uint32) bool {
    p0 = p0
    p0 = 127
    if false {
        var v4 uint32 = p1
        var v5 string = ((("" + "foo bar") + ("foo bar" + ("foo bar" + "\n\t"))) + "\n\t")
        _ = v4
        _ = v5
    }
    _ = p0
    _ = p1
    return (false || (false && false))
}
func genFn1(p0 int, p1 uint32, p2 uint) uint64 {
    var v6 bool = (1 == 1)
    v7 := []int{p0, (p0 >> (((p2 & (^(p2)))) & 63)), (p0 / ((-9223372036854775808) | 1))}
    var v8 int64 = 9223372036854775807
    var v9 uint64 = 18446744073709551615
    if (v8 >= v8) {
        fmt.Println(v9, p1, p2)
        p2 -= uint(p1)
        if (((uint16(v9) >= 616) || (!v6)) || ((v6 && (!true)) && (!true))) {
            for fi10 := 0; fi10 < 12; fi10++ {
                v11 := GenS1{F0: (fi10 - p0), F1: 1, F2: 63}
                _ = fi10
                _ = v11
            }
            if v6 {
                v7 = append(v7, (p0 + p0))
            } else {
                switch ("" + "") {
                case "a":
                    var v12 float64 = (3.14159 + 0.0)
                    _ = v12
                case "foo bar":
                    var v13 uint16 = 2
                    _ = v13
                default:
                    fmt.Println(p2)
                }
                for fi14 := 0; fi14 < 6; fi14++ {
                    fmt.Println(fi14, p2, p1)
                    _ = fi14
                }
            }
            fmt.Println(v6)
        }
    } else {
        v15 := GenS2{F0: 1, F1: 66, F2: 42}
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 int8 = 127
            v18 := GenS1{F0: (^(fi16)), F1: 32767, F2: (v17 ^ int8(p0))}
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 468
}

func main() {
    v1 := []int{}
    v2 := []uint32{1, 4294967295}
    var v3 uint = 18446744073709551615
    v19 := GenS2{F0: 31, F1: 0, F2: 1}
    var v20 int32 = 42
    var v21 int32 = (v20 << (((v3 - 987)) & 63))
    if ((!((0 >= -9223372036854775808) || true)) || true) {
        v22 := GenS2{F0: -68, F1: 128, F2: 48}
        switch 0 {
        case -1:
            var v23 int64 = int64(v21)
            _ = v23
        case 32767:
            v20 += v20
        case -9223372036854775808:
            if true {
                var v24 string = ""
                v25 := GenS3{F0: 1}
                _ = v24
                _ = v25
            } else {
                var v26 uint = v3
                _ = v26
            }
        }
        _ = v22
    }
    fmt.Println(v3, v3)
    switch (("hello" + ((("a" + ("hello" + "\n\t")) + "foo bar") + "")) + ("hello" + ("hello" + "a"))) {
    case "":
        for fi27 := 0; fi27 < 5; fi27++ {
            v3 = v3
            if (42 < 127) {
                for fi28 := 0; fi28 < 3; fi28++ {
                    v3 = (v3 / (((v3 << (((v3 >> ((104) & 63))) & 63))) | 1))
                    if false {
                        var v29 uint = (v3 >> ((18446744073709551615) & 63))
                        v30 := GenS4{F0: -901, F1: (fi27 - fi27)}
                        var v31 int = fi28
                        _ = v29
                        _ = v30
                        _ = v31
                    } else {
                        if true {
                            var v32 bool = (1 > 32767)
                            v3 = (^(v3))
                            var v33 uint16 = 1
                            _ = v32
                            _ = v33
                        }
                        var v34 bool = false
                        var v35 bool = (42 == 36)
                        _ = v34
                        _ = v35
                    }
                    _ = fi28
                }
            }
            var v36 int16 = -620
            _ = fi27
            _ = v36
        }
    case "\n\t":
        var v37 uint16 = 631
        _ = v37
    case "foo bar":
        v38 := GenS5{F0: (v21 % ((42) | 1)), F1: -32768}
        _ = v38
    }
    var v39 float64 = -2.25
    var v40 uint64 = 2
    if len(v2) > 0 {
        v2[0] = 0
    }
    var v41 float32 = 0.1
    fmt.Println(v41, v20)
    switch 32768 {
    case 755:
        _ = len(v2)
    default:
        var v42 uint8 = 42
        _ = v42
    }
    var v43 uint32 = 193
    var v44 string = ("\n\t" + "\n\t")
    if true {
        fmt.Println(v3, v41, v41)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v39
    _ = v40
    _ = v41
    _ = v43
    _ = v44
}
