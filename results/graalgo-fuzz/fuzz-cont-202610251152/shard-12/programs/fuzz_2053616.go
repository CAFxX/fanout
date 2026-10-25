package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 int16; F1 uint8; F2 int16 }
type GenS5 struct { F0 uint64; F1 int16; F2 int64 }
type GenS6 struct { F0 bool; F1 string }
func genFn0(p0 int8, p1 int32, p2 bool) int32 {
    v4 := []int{-9223372036854775808}
    if false {
        p1 -= (p1 + p1)
        for fi5 := 0; fi5 < 10; fi5++ {
            var v6 int16 = 385
            p0 += p0
            switch fi5 {
            case 32767:
                var v7 float64 = 0.0
                _ = v7
            case 4294967295:
                for fi8 := 0; fi8 < 4; fi8++ {
                    if (((^(fi8)) <= fi8) && (("a" + "a") == "")) {
                        if (!(false && p2)) {
                            var v9 uint32 = 42
                            if true {
                                p1 = p1
                                var v10 float64 = (-2.25 - 1.5)
                                _ = v10
                            }
                            v11 := GenS1{F0: 255}
                            var v12 uint16 = 65535
                            v6 = (v6 * (v6 * 1))
                            var v13 bool = ((p1 * (p1 & p1)) < p1)
                            var v14 string = (("" + "foo bar") + "\n\t")
                            _ = v11
                            _ = v12
                            _ = v13
                            _ = v14
                            _ = v9
                        }
                        if true {
                            var v15 int = fi5
                            _ = v15
                        }
                        var v16 string = ""
                        _ = v16
                    }
                    _ = fi8
                }
            case 2147483647:
                p2 = ((v6 - v6) < v6)
            }
            _ = fi5
            _ = v6
        }
    } else {
        for fi17 := 0; fi17 < 12; fi17++ {
            v4 = append(v4, (fi17 * fi17))
            _ = fi17
        }
        var v18 int8 = 0
        if p2 {
            fmt.Println(p0, p1, v18)
            switch ("\n\t" + "a") {
            case "hello":
                var v19 int64 = 0
                _ = v19
            case "":
                v20 := GenS1{F0: 1}
                _ = v20
            case "foo bar":
                v21 := GenS2{F0: 128, F1: 0}
                _ = v21
            }
        }
        _ = v18
    }
    switch "hello" {
    case "\n\t":
        v22 := GenS3{F0: -800, F1: 255}
        _ = v22
    }
    v23 := GenS1{F0: 42}
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    _ = v4
    return (p1 >> ((4294967295) & 63))
}
func genFn1(p0 int16, p1 int16) int64 {
    var v24 int16 = (p1 | p1)
    for fi25 := 0; fi25 < 3; fi25++ {
        var v26 bool = (fi25 <= 32768)
        v27 := GenS4{F0: (^(p0)), F1: 127, F2: (-(v24))}
        var v28 int = 2147483647
        _ = fi25
        _ = v26
        _ = v27
        _ = v28
    }
    switch -1 {
    case -810:
        p1 += (v24 * v24)
    case -9223372036854775808:
        if false {
            fmt.Println(v24, v24)
        } else {
            v29 := GenS4{F0: v24, F1: 127, F2: (v24 - (p1 % ((p0) | 1)))}
            var v30 uint16 = 65535
            _ = v29
            _ = v30
        }
    default:
        if true {
            var v31 uint = 0
            v32 := []int{int(v31), 32768, 32768}
            for fi33 := 0; fi33 < 6; fi33++ {
                var v34 float32 = 0.1
                v35 := GenS2{F0: 106, F1: -1}
                _ = fi33
                _ = v34
                _ = v35
            }
            _ = v31
            _ = v32
        }
    }
    for fi36 := 0; fi36 < 2; fi36++ {
        p1 -= v24
        _ = fi36
    }
    _ = p0
    _ = p1
    _ = v24
    return 42
}

func main() {
    var v1 int32 = -1
    var v2 int32 = (v1 & (v1 & v1))
    v3 := []string{}
    var v37 uint8 = 27
    if true {
        fmt.Println(v37)
    } else {
        v37 = v37
    }
    for fi38 := 0; fi38 < 17; fi38++ {
        fmt.Println(v2, v37)
        _ = fi38
    }
    v39 := GenS5{F0: 2, F1: 42, F2: -494}
    var v40 float32 = 0.0
    v41 := GenS5{F0: 4294967295, F1: 32767, F2: genFn1(-1, -32768)}
    var v42 float32 = (v40 + 3.14159)
    if true {
        var v43 int64 = 311
        _ = v43
    } else {
        var v44 int8 = -111
        fmt.Println(v40)
        var v45 int = 4294967295
        _ = v44
        _ = v45
    }
    v40 = (0.0 + v42)
    var v46 uint8 = 141
    if false {
        v47 := GenS4{F0: 32767, F1: v37, F2: 1000}
        v46 += (v46 % ((v46) | 1))
        _ = v47
    }
    var v48 int8 = 100
    _ = len(v3)
    if ((v42 == v42) && true) {
        var v49 int64 = -1
        for fi50 := 0; fi50 < 7; fi50++ {
            v51 := GenS6{F0: (399 > 0), F1: "a"}
            var v52 uint32 = 2147483647
            _ = fi50
            _ = v51
            _ = v52
        }
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v39
    _ = v40
    _ = v41
    _ = v42
    _ = v46
    _ = v48
}
