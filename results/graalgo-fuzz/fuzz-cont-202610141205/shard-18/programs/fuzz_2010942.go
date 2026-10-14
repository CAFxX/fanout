package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 uint16; F1 int8; F2 uint32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint; F1 int64 }
type GenS7 struct { F0 uint8; F1 string; F2 uint16 }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int32 = 1
    var v3 uint64 = 0
    var v4 uint64 = (v3 % ((v3) | 1))
    if (v3 < v3) {
        var v5 float32 = 3.14159
        v6 := []int{}
        _ = v5
        _ = v6
    } else {
        var v7 int8 = 127
        fmt.Println(v2, v2)
        if true {
            v8 := GenS1{F0: 42, F1: 255, F2: 1}
            for fi9 := 0; fi9 < 6; fi9++ {
                if (1.5 <= (1e10 / 3.14159)) {
                    fmt.Println(v1, v3, v1)
                    switch ("" + "hello") {
                    case "":
                        v2 = v2
                    case "\n\t":
                        for fi10 := 0; fi10 < 14; fi10++ {
                            v11 := GenS2{F0: v3}
                            var v12 int = fi9
                            _ = fi10
                            _ = v11
                            _ = v12
                        }
                    case "hello":
                        v13 := GenS3{F0: 2, F1: 65535}
                        _ = v13
                    }
                }
                var v14 float32 = float32(1e10)
                for fi15 := 0; fi15 < 10; fi15++ {
                    if (1e10 == float64(0.0)) {
                        var v16 float32 = (float32(0.1) / -2.25)
                        var v17 uint8 = 173
                        var v18 int8 = 1
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    v7 -= 108
                    _ = fi15
                }
                _ = fi9
                _ = v14
            }
            _ = v8
        } else {
            v3 = v3
            if false {
                v19 := GenS2{F0: v3}
                v20 := []int{32767}
                var v21 int64 = v1
                _ = v19
                _ = v20
                _ = v21
            } else {
                var v22 uint = 2
                for fi23 := 0; fi23 < 9; fi23++ {
                    if true {
                        if true {
                            if true {
                                for fi24 := 0; fi24 < 15; fi24++ {
                                    var v25 uint = v22
                                    v2 = (-(v2))
                                    _ = fi24
                                    _ = v25
                                }
                                for fi26 := 0; fi26 < 12; fi26++ {
                                    switch ("" + "\n\t") {
                                    case "":
                                        fmt.Println(fi23, v4, fi26)
                                    case "hello":
                                        v27 := GenS1{F0: 128, F1: 1, F2: v22}
                                        var v28 int8 = 42
                                        fmt.Println(fi26, v2)
                                        _ = v27
                                        _ = v28
                                    case "foo bar":
                                        var v29 int16 = 32767
                                        _ = v29
                                    }
                                    var v30 bool = true
                                    _ = fi26
                                    _ = v30
                                }
                                var v31 int32 = (v2 ^ (-(v2)))
                                _ = v31
                            }
                            v32 := GenS4{F0: 1, F1: -111, F2: 997}
                            for fi33 := 0; fi33 < 11; fi33++ {
                                v7 += v7
                                v34 := []int{fi23, (-(fi23))}
                                _ = fi33
                                _ = v34
                            }
                            _ = v32
                        }
                        v35 := []int{(fi23 | (fi23 + fi23)), fi23, -9223372036854775808, (-(fi23))}
                        if len(v35) >= 2 {
                            v35 = v35[1:len(v35)]
                        }
                        _ = v35
                    } else {
                        fmt.Println(v2, v7, v2)
                    }
                    v36 := GenS4{F0: 42, F1: 1, F2: 2}
                    if false {
                        v3 -= (v3 | (v4 | 0))
                        var v37 bool = true
                        _ = v37
                    }
                    _ = fi23
                    _ = v36
                }
                v22 -= (v22 % ((v22) | 1))
                _ = v22
            }
            v38 := []int{}
            _ = v38
        }
        _ = v7
    }
    v2 += v2
    switch 32768 {
    case 32768:
        var v39 int8 = -1
        _ = v39
    case 4294967295:
        var v40 uint32 = 4294967295
        _ = v40
    }
    v41 := GenS5{F0: false}
    v42 := []int{}
    if true {
        switch -98 {
        case 87:
            fmt.Println(v2, v1)
        default:
            var v43 uint64 = 18446744073709551615
            _ = v43
        }
    } else {
        v4 -= v4
        if len(v42) >= 2 {
            v42 = v42[1:len(v42)]
        }
        if len(v42) >= 2 {
            v42 = v42[1:len(v42)]
        }
    }
    fmt.Println(v1)
    var v44 uint64 = 804
    v4 = (^(v3))
    if (0.0 == (float64((1e10 * 3.14159)) + float64(((1.5 - float32(1e10)) - 1e10)))) {
        v4 = v4
        var v45 uint8 = 127
        v46 := GenS6{F0: 255, F1: (v1 * (v1 * (v1 >> ((42) & 63))))}
        _ = v45
        _ = v46
    } else {
        for fi47 := 0; fi47 < 12; fi47++ {
            var v48 string = "foo bar"
            v49 := GenS7{F0: 0, F1: "foo bar", F2: 287}
            _ = fi47
            _ = v48
            _ = v49
        }
    }
    v50 := GenS2{F0: v4}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v44
    _ = v50
}
