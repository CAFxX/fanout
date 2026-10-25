package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 bool }
func genFn0(p0 int32) int32 {
    for fi4 := 0; fi4 < 14; fi4++ {
        if true {
            var v5 float32 = 1.5
            _ = v5
        } else {
            for fi6 := 0; fi6 < 4; fi6++ {
                if (!(false || (4294967295 <= 131))) {
                    if (((40 > 2) && true) || (true && ((((1e10 + ((3.14159 + 0.0) / 1.0)) * 3.14159) != 1e10) && true))) {
                        v7 := []int{(fi4 + 0)}
                        if ("foo bar" == "hello") {
                            p0 -= p0
                        } else {
                            var v8 float32 = (((0.1 + (0.0 + -2.25)) - 1.5) - 3.14159)
                            var v9 int32 = 2147483647
                            v10 := GenS1{F0: 44}
                            _ = v10
                            _ = v8
                            _ = v9
                        }
                        _ = v7
                    } else {
                        var v11 float64 = -2.25
                        if (((0.1 < 0.1) && true) || true) {
                            var v12 int64 = 0
                            var v13 bool = true
                            var v14 uint16 = 663
                            v15 := []int{(fi6 - 9223372036854775807)}
                            _ = v12
                            _ = v13
                            _ = v14
                            _ = v15
                        }
                        p0 -= -1
                        _ = v11
                    }
                    var v16 uint32 = 97
                    var v17 bool = true
                    _ = v16
                    _ = v17
                } else {
                    var v18 uint16 = 0
                    _ = v18
                }
                fmt.Println(p0, fi6)
                if true {
                    var v19 int8 = -99
                    switch "foo bar" {
                    case "foo bar":
                        var v20 int64 = 42
                        _ = v20
                    case "a":
                        v21 := GenS1{F0: 55}
                        _ = v21
                    default:
                        v19 = (v19 % (((^(v19))) | 1))
                    }
                    p0 = p0
                    v22 := []int{(-(fi4)), (^(fi6))}
                    _ = v19
                    _ = v22
                }
                _ = fi6
            }
        }
        var v23 int16 = 885
        for fi24 := 0; fi24 < 9; fi24++ {
            v23 = 1
            v25 := []int{fi24}
            if len(v25) >= 2 {
                v25 = v25[1:len(v25)]
            }
            _ = fi24
            _ = v25
        }
        _ = fi4
        _ = v23
    }
    if false {
        fmt.Println(p0)
        fmt.Println(p0)
    } else {
        v26 := GenS1{F0: 127}
        var v27 float32 = 0.0
        _ = v26
        _ = v27
    }
    _ = p0
    return p0
}

func main() {
    var v1 int16 = 670
    v2 := make(map[string]bool)
    v3 := make(map[string]uint32)
    for fi28 := 0; fi28 < 2; fi28++ {
        fmt.Println(v1, fi28)
        if (1 >= 127) {
            for fi29 := 0; fi29 < 5; fi29++ {
                if (!false) {
                    v3["foo bar"] = 980
                    if ((!((1e10 != ((0.1 - 0.1) * 0.1)) && false)) || (fi29 > 9223372036854775807)) {
                        v1 += v1
                        var v30 float64 = 1e10
                        var v31 float64 = v30
                        _ = v30
                        _ = v31
                    }
                    v1 -= 32767
                }
                _ = fi29
            }
            var v32 bool = (!false)
            var v33 uint64 = 2
            _ = v32
            _ = v33
        }
        for fi34 := 0; fi34 < 20; fi34++ {
            fmt.Println(fi34, fi28)
            _ = fi34
        }
        _ = fi28
    }
    for fi35 := 0; fi35 < 18; fi35++ {
        if false {
            if (!true) {
                v1 = 1000
            } else {
                v1 -= v1
                var v36 uint16 = 32767
                _ = v36
            }
            v1 = v1
            var v37 uint32 = 0
            _ = v37
        }
        var v38 bool = false
        var v39 string = ("\n\t" + ("a" + "a"))
        _ = fi35
        _ = v38
        _ = v39
    }
    var v40 uint32 = 2147483647
    switch ("" + "\n\t") {
    case "hello":
        for fi41 := 0; fi41 < 11; fi41++ {
            v42 := []int{}
            v1 = -313
            _ = fi41
            _ = v42
        }
    case "a":
        v43 := GenS1{F0: 217}
        _ = v43
    case "":
        if ((0 <= 4294967295) && (false && true)) {
            v44 := []int{-9223372036854775808, -2147483648, -1}
            v45 := GenS2{F0: 65535, F1: true}
            _ = v44
            _ = v45
        }
    }
    v1 += v1
    var v46 int64 = -1
    v47 := GenS1{F0: uint8(v40)}
    fmt.Println(v40, v1)
    v1 -= (-(v1))
    var v48 string = ""
    if (!(!true)) {
        var v49 uint32 = 0
        var v50 int32 = 212
        _ = v49
        _ = v50
    } else {
        v51 := []int{32768, -509}
        _ = v51
    }
    fmt.Println(v48, v40)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v40:", v40)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v46
    _ = v47
    _ = v48
}
