package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 uint16; F1 uint32 }
func genFn0(p0 uint, p1 int) string {
    var v6 int16 = -32768
    v7 := GenS1{F0: 42, F1: 42}
    if false {
        var v8 float32 = 0.1
        v6 = (v6 * (v6 ^ (v6 * (v6 / ((v6) | 1)))))
        if false {
            fmt.Println(v8, v8, v8)
            fmt.Println(p1, v6, p0)
            fmt.Println(p0, v8, p1)
        }
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return ""
}
func genFn1() uint64 {
    v9 := []int{}
    var v10 int8 = 31
    if (!(float64(0.1) <= (-2.25 - 1.5))) {
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 int32 = 771
            fmt.Println(v10, fi11, v10)
            _ = len(v9)
            _ = fi11
            _ = v12
        }
    } else {
        switch 32768 {
        case 32767:
            if true {
                if true {
                    v10 -= -1
                }
                if true {
                    if true {
                        var v13 uint32 = 4294967295
                        var v14 float64 = -2.25
                        v9 = append(v9, 32768)
                        _ = v13
                        _ = v14
                    }
                    v10 = v10
                }
                v10 -= -123
            }
        case 2147483647:
            v15 := GenS2{F0: 0}
            _ = v15
        case 1:
            for fi16 := 0; fi16 < 15; fi16++ {
                v17 := GenS2{F0: 876}
                _ = fi16
                _ = v17
            }
        default:
            fmt.Println(v10)
        }
        fmt.Println(v10)
        v10 = (v10 | (v10 << ((4294967295) & 63)))
    }
    switch ("foo bar" + ((("\n\t" + "hello") + (("a" + "\n\t") + "hello")) + "a")) {
    case "\n\t":
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    case "foo bar":
        v10 -= (v10 << ((42) & 63))
    }
    v18 := GenS1{F0: 128, F1: -2147483648}
    _ = v10
    _ = v18
    _ = v9
    return 1
}

func main() {
    var v1 float32 = 1.5
    var v2 int32 = 719
    var v3 uint64 = 448
    v4 := make(map[int]string)
    v5 := make(map[uint32]string)
    v1 = 1.5
    for fi19 := 0; fi19 < 16; fi19++ {
        if ((!(267 >= 576)) || (!true)) {
            if true {
                v4[324] = "hello"
                if ((65535 <= 539) || (!(959 < 42))) {
                    var v20 int = (fi19 << ((2) & 63))
                    fmt.Println(v20, v1, v3)
                    v20 = v20
                    _ = v20
                }
                var v21 int32 = v2
                _ = v21
            } else {
                var v22 string = ""
                _ = v22
            }
        }
        _ = fi19
    }
    if true {
        v23 := []int{-9223372036854775808, -579, 128, 9223372036854775807}
        v1 = 0.0
        var v24 uint8 = 103
        _ = v23
        _ = v24
    }
    var v25 int8 = 0
    var v26 uint = 1
    var v27 int64 = 340
    v28 := GenS3{F0: (("" + "\n\t") + "hello"), F1: v27}
    fmt.Println(v3, v1)
    var v29 int = int(v2)
    var v30 string = ("hello" + ("\n\t" + (("hello" + "hello") + "")))
    fmt.Println(v25, v25)
    var v31 uint32 = 0
    if ((v31 & v31) >= 1) {
        fmt.Println(v30, v2)
        v29 -= (v29 / (((v29 >> ((347) & 63))) | 1))
        if false {
            v1 += -2.25
            if (false && (true && (v25 < (^(v25))))) {
                v32 := GenS4{F0: 65535, F1: (v31 & v31)}
                _ = v32
            } else {
                var v33 int = (-(v29))
                _ = v33
            }
        }
    } else {
        _ = len(v4)
        v34 := v5[2147483648]
        v1 = v1
        _ = v34
    }
    v35 := []int{v29, v29}
    var v36 int8 = (-(v25))
    if false {
        var v37 float32 = (v1 + 0.0)
        v38 := GenS1{F0: 0, F1: (v2 + -2147483648)}
        _ = v37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
