package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint32; F1 string; F2 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int }
func genFn0() uint {
    v7 := []int{1}
    for fi8 := 0; fi8 < 9; fi8++ {
        if ((-2.25 - float64(-2.25)) >= float64(float32(-2.25))) {
            v7 = append(v7, (^(fi8)))
        }
        for fi9 := 0; fi9 < 15; fi9++ {
            switch -9223372036854775808 {
            case 32768:
                var v10 float32 = 3.14159
                _ = v10
            case 266:
                if false {
                    var v11 float32 = -2.25
                    v11 = v11
                    if len(v7) > 0 {
                        v7[0] = 4294967295
                    }
                    _ = v11
                }
            default:
                v12 := GenS1{F0: "a", F1: (fi9 & (fi8 % ((-1) | 1)))}
                var v13 uint32 = 4294967295
                var v14 float32 = ((1.5 * 1.5) / 0.1)
                _ = v12
                _ = v13
                _ = v14
            }
            var v15 int64 = 42
            _ = fi9
            _ = v15
        }
        _ = fi8
    }
    if true {
        for fi16 := 0; fi16 < 16; fi16++ {
            for fi17 := 0; fi17 < 11; fi17++ {
                var v18 string = "a"
                var v19 uint8 = 1
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                _ = fi17
                _ = v18
                _ = v19
            }
            var v20 string = "\n\t"
            _ = fi16
            _ = v20
        }
        var v21 int16 = 42
        _ = v21
    }
    _ = v7
    return 4294967295
}

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 82
    var v3 int16 = -32768
    v4 := []int{975, -1, 128}
    var v5 uint32 = 65535
    var v6 int32 = -631
    var v22 int32 = v6
    switch 32768 {
    case -231:
        var v23 bool = (!false)
        _ = v23
    case 32767:
        _ = len(v4)
    case 1:
        fmt.Println(v5)
    default:
        delete(v1, "")
    }
    _ = len(v1)
    v24 := GenS2{F0: uint32(v2), F1: "", F2: 0}
    if false {
        v3 = v3
        fmt.Println(v6, v2)
    } else {
        var v25 bool = false
        var v26 uint = 18446744073709551615
        v2 = v2
        _ = v25
        _ = v26
    }
    var v27 float32 = 3.14159
    v22 = (v22 & v22)
    var v28 bool = (v22 > v6)
    var v29 int = 1
    var v30 uint = 163
    switch "a" {
    case "foo bar":
        if v28 {
            for fi31 := 0; fi31 < 17; fi31++ {
                fmt.Println(v30, v5)
                _ = fi31
            }
            var v32 float32 = (v27 + v27)
            _ = v32
        } else {
            for fi33 := 0; fi33 < 15; fi33++ {
                var v34 int32 = (v6 * (v22 % ((-2147483648) | 1)))
                _ = fi33
                _ = v34
            }
            for fi35 := 0; fi35 < 11; fi35++ {
                v36 := GenS3{F0: (v5 >= (v5 / ((v5) | 1)))}
                if (int64(fi35) >= 42) {
                    var v37 uint32 = 2147483648
                    v37 += 2147483648
                    for fi38 := 0; fi38 < 17; fi38++ {
                        var v39 string = "a"
                        _ = fi38
                        _ = v39
                    }
                    _ = v37
                }
                _ = fi35
                _ = v36
            }
        }
    case "hello":
        var v40 uint16 = 260
        _ = v40
    case "":
        v6 = (^(v6))
    default:
        switch "a" {
        case "\n\t":
            if ((!v28) || v28) {
                v22 += (v6 >> ((v30) & 63))
                var v41 int = v29
                v42 := GenS3{F0: v28}
                _ = v41
                _ = v42
            } else {
                v28 = (("foo bar" + "foo bar") > "hello")
            }
        case "":
            fmt.Println(v3)
        case "a":
            for fi43 := 0; fi43 < 2; fi43++ {
                var v44 int = fi43
                if v28 {
                    fmt.Println(v6, v28)
                    for fi45 := 0; fi45 < 9; fi45++ {
                        var v46 string = "\n\t"
                        v47 := GenS4{F0: fi45}
                        _ = fi45
                        _ = v46
                        _ = v47
                    }
                }
                _ = fi43
                _ = v44
            }
        }
    }
    v3 = int16(v2)
    if true {
        var v48 int = (v29 >> (((v30 + 808)) & 63))
        var v49 uint32 = 2147483647
        for fi50 := 0; fi50 < 1; fi50++ {
            var v51 int32 = 2147483647
            _ = fi50
            _ = v51
        }
        _ = v48
        _ = v49
    }
    v22 = -647
    var v52 int64 = 42
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v52
    _ = v6
}
