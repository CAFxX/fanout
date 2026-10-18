package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 uint }
func genFn0(p0 uint) uint32 {
    fmt.Println(p0)
    for fi6 := 0; fi6 < 14; fi6++ {
        for fi7 := 0; fi7 < 16; fi7++ {
            fmt.Println(p0, fi6, fi7)
            if (1e10 < 3.14159) {
                var v8 float32 = (3.14159 + ((-2.25 - float32(1e10)) - 1e10))
                switch -9223372036854775808 {
                case 2147483647:
                    var v9 uint32 = 42
                    _ = v9
                case 32767:
                    var v10 uint8 = 241
                    _ = v10
                default:
                    var v11 int8 = -1
                    _ = v11
                }
                fmt.Println(fi7)
                p0 = 4294967295
                var v12 uint16 = 1
                v13 := GenS1{F0: "a"}
                var v14 uint64 = 18446744073709551615
                _ = v12
                _ = v13
                _ = v14
                _ = v8
            } else {
                p0 -= (p0 >> ((p0) & 63))
            }
            _ = fi7
        }
        var v15 uint32 = 562
        _ = fi6
        _ = v15
    }
    v16 := GenS1{F0: ""}
    var v17 float64 = (-2.25 / 1.0)
    _ = p0
    _ = v16
    _ = v17
    return 2
}
func genFn1() uint8 {
    v18 := GenS2{F0: (true || false), F1: 220}
    switch 0 {
    case 4:
        for fi19 := 0; fi19 < 11; fi19++ {
            for fi20 := 0; fi20 < 18; fi20++ {
                var v21 uint = 189
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
    case 32767:
        var v22 uint64 = 42
        _ = v22
    }
    var v23 int8 = -1
    _ = v18
    _ = v23
    return 128
}

func main() {
    var v1 float64 = -2.25
    var v2 float32 = ((0.0 / (((1.5 + 0.1) + float32((v1 + ((float64(3.14159) - float64(0.1)) * v1)))) * 1.5)) - 0.0)
    var v3 int8 = 42
    var v4 string = ("\n\t" + "foo bar")
    var v5 int32 = -2147483648
    switch 0 {
    case 4294967295:
        var v24 string = v4
        _ = v24
    case 434:
        var v25 int8 = v3
        _ = v25
    }
    v26 := []int{-2147483648}
    _ = len(v26)
    var v27 uint16 = 255
    for fi28 := 0; fi28 < 11; fi28++ {
        v29 := GenS3{F0: 65535}
        v30 := GenS4{F0: 42, F1: ((v5 & v5) > v5)}
        _ = fi28
        _ = v29
        _ = v30
    }
    var v31 int16 = 243
    v32 := GenS3{F0: v27}
    switch -9223372036854775808 {
    case 4294967295:
        var v33 uint8 = genFn1()
        _ = v33
    case 128:
        v5 += v5
    }
    v34 := GenS4{F0: genFn1(), F1: (!(2147483647 <= (^(v5))))}
    switch -831 {
    case -1:
        var v35 uint32 = 0
        _ = v35
    case 1:
        v27 -= v27
    }
    var v36 float64 = v1
    v2 = (v2 - v2)
    var v37 bool = false
    if false {
        if len(v26) >= 2 {
            v26 = v26[1:len(v26)]
        }
        var v38 uint32 = 42
        _ = v38
    }
    for fi39 := 0; fi39 < 18; fi39++ {
        if true {
            if v37 {
                for fi40 := 0; fi40 < 9; fi40++ {
                    var v41 bool = false
                    v27 = (^(v27))
                    _ = fi40
                    _ = v41
                }
            }
            v37 = (!(!(v37 && (!false))))
            for fi42 := 0; fi42 < 3; fi42++ {
                var v43 bool = (v2 < v2)
                _ = fi42
                _ = v43
            }
        }
        var v44 uint16 = (v27 << ((152) & 63))
        _ = fi39
        _ = v44
    }
    v45 := GenS5{F0: 1}
    var v46 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v34
    _ = v36
    _ = v37
    _ = v4
    _ = v45
    _ = v46
    _ = v5
}
