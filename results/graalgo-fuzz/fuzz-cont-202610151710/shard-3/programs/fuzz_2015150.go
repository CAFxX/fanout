package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int16; F1 uint16; F2 int64 }
type GenS3 struct { F0 uint32; F1 string; F2 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int; F1 int64; F2 int16 }
func genFn0(p0 string, p1 bool, p2 uint) uint64 {
    v4 := GenS1{F0: -156, F1: p2}
    fmt.Println(p2, p1, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 135
}
func genFn1() int16 {
    if false {
        for fi5 := 0; fi5 < 9; fi5++ {
            var v6 float64 = 3.14159
            _ = fi5
            _ = v6
        }
        var v7 int32 = -762
        if (47 != 128) {
            if true {
                v8 := []int{252, -2147483648, 32768}
                _ = v8
            }
        } else {
            switch 0 {
            case 32767:
                var v9 float32 = (-2.25 + 1e10)
                _ = v9
            case -288:
                if true {
                    v10 := []int{101, -289, 0}
                    if false {
                        fmt.Println(v7)
                        var v11 float32 = ((((3.14159 * 1.5) * float32(1.5)) / ((((-2.25 / -2.25) + ((3.14159 + 3.14159) / ((0.1 * -2.25) + (0.0 - (3.14159 * (float32(0.1) - -2.25)))))) + 1.5) - float32(1.5))) + 0.0)
                        _ = v11
                    } else {
                        var v12 int64 = 42
                        switch "" {
                        case "foo bar":
                            var v13 int64 = (v12 / (((v12 & 9223372036854775807)) | 1))
                            _ = v13
                        case "":
                            v7 = (v7 << ((18446744073709551615) & 63))
                        case "hello":
                            v7 = (^(v7))
                        }
                        if false {
                            var v14 uint = 36
                            v14 = (^(v14))
                            _ = v14
                        } else {
                            v12 -= (v12 / ((-426) | 1))
                            v15 := GenS1{F0: 0, F1: 1}
                            _ = v15
                        }
                        _ = v12
                    }
                    var v16 uint16 = 1
                    _ = v10
                    _ = v16
                }
            case 4294967295:
                if true {
                    v17 := GenS2{F0: 1000, F1: 2, F2: 9223372036854775807}
                    v18 := GenS3{F0: 875, F1: ("foo bar" + (("a" + "hello") + "foo bar")), F2: "\n\t"}
                    _ = v17
                    _ = v18
                }
            }
            for fi19 := 0; fi19 < 20; fi19++ {
                var v20 uint8 = 1
                fmt.Println(v20, v20)
                _ = fi19
                _ = v20
            }
            var v21 uint16 = 255
            _ = v21
        }
        _ = v7
    }
    if true {
        for fi22 := 0; fi22 < 13; fi22++ {
            if false {
                var v23 bool = false
                if false {
                    v23 = (!(!(!true)))
                }
                _ = v23
            }
            v24 := GenS4{F0: (fi22 & fi22)}
            var v25 uint8 = 33
            _ = fi22
            _ = v24
            _ = v25
        }
        var v26 float32 = 0.0
        var v27 float32 = (v26 * v26)
        _ = v26
        _ = v27
    }
    switch 42 {
    case 272:
        var v28 uint64 = genFn0(("a" + ("foo bar" + "")), ((!true) && true), 536)
        _ = v28
    case 543:
        var v29 uint64 = 0
        _ = v29
    case -574:
        v30 := GenS5{F0: 9223372036854775807}
        _ = v30
    default:
        var v31 float32 = 0.1
        _ = v31
    }
    var v32 int32 = -2147483648
    for fi33 := 0; fi33 < 12; fi33++ {
        if true {
            v34 := GenS6{F0: 206}
            _ = v34
        }
        var v35 int = (fi33 % (((fi33 / ((fi33) | 1))) | 1))
        _ = fi33
        _ = v35
    }
    _ = v32
    return -1
}

func main() {
    var v1 int32 = 1
    v2 := []uint32{42, 2, 2147483648}
    var v3 bool = false
    v3 = (false && false)
    var v36 float64 = 0.0
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    var v37 int8 = 100
    v1 = v1
    v3 = (2 <= 4294967295)
    var v38 string = "a"
    var v39 uint64 = 18446744073709551615
    if v3 {
        v38 = "\n\t"
        for fi40 := 0; fi40 < 12; fi40++ {
            v39 = (^(v39))
            _ = fi40
        }
    } else {
        v37 = v37
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v37 = 0
    }
    if len(v2) > 0 {
        v2[0] = 681
    }
    if v3 {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if (v1 != v1) {
            fmt.Println(v38, v38, v37)
            var v41 int64 = 1
            var v42 uint16 = 388
            _ = v41
            _ = v42
        } else {
            v2 = append(v2, 2)
            fmt.Println(v3)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        v36 = 0.0
    }
    for fi43 := 0; fi43 < 9; fi43++ {
        fmt.Println(v39)
        for fi44 := 0; fi44 < 10; fi44++ {
            if (v1 == v1) {
                if v3 {
                    if (((fi43 + (fi44 + fi44)) == (fi44 & 35)) && true) {
                        v39 += v39
                        _ = len(v2)
                    }
                    v45 := GenS4{F0: (fi44 / ((-1) | 1))}
                    _ = v45
                } else {
                    v38 = (v38 + (v38 + v38))
                }
                v46 := GenS7{F0: (fi44 & fi43), F1: 1, F2: -980}
                _ = v46
            }
            _ = fi44
        }
        _ = fi43
    }
    var v47 string = (v38 + "hello")
    switch ("" + v38) {
    case "":
        fmt.Println(v39, v1, v38)
    case "hello":
        var v48 uint8 = 254
        _ = v48
    case "\n\t":
        fmt.Println(v3)
    }
    if len(v2) > 0 {
        v2[0] = 42
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v47
}
