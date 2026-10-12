package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 uint; F2 int8 }
type GenS4 struct { F0 uint; F1 int64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS7 struct { F0 int16; F1 string; F2 int8 }
func genFn0(p0 uint) uint8 {
    var v6 uint16 = 1
    v7 := []int{-413}
    _ = p0
    _ = v6
    _ = v7
    return 153
}

func main() {
    var v1 uint16 = 399
    var v2 uint32 = 4294967295
    var v3 int = 42
    var v4 uint32 = (^(v2))
    v5 := make(map[string]string)
    fmt.Println(v4)
    for fi8 := 0; fi8 < 8; fi8++ {
        switch "" {
        case "a":
            v9 := GenS1{F0: 471, F1: -1}
            _ = v9
        case "\n\t":
            if true {
                v10 := GenS2{F0: -1}
                _ = v10
            } else {
                if (false || false) {
                    v11 := GenS2{F0: 127}
                    _ = v11
                } else {
                    v1 -= (^(v1))
                }
                for fi12 := 0; fi12 < 4; fi12++ {
                    var v13 int = v3
                    if true {
                        if false {
                            v14 := GenS3{F0: 255, F1: 2, F2: -1}
                            _ = v14
                        } else {
                            v15 := GenS4{F0: 4294967295, F1: -1}
                            fmt.Println(fi8)
                            _ = v15
                        }
                    }
                    v2 = 422
                    switch "hello" {
                    case "\n\t":
                        var v16 uint32 = (v4 | v2)
                        _ = v16
                    case "foo bar":
                        fmt.Println(v3)
                    default:
                        v1 = v1
                    }
                    v17 := GenS3{F0: 255, F1: 1, F2: 0}
                    _ = fi12
                    _ = v13
                    _ = v17
                }
                fmt.Println(v2, v4, v3)
            }
        case "":
            var v18 string = "\n\t"
            _ = v18
        default:
            v19 := GenS5{F0: 1}
            _ = v19
        }
        var v20 float64 = 3.14159
        _ = fi8
        _ = v20
    }
    v3 = (v3 & -2147483648)
    for fi21 := 0; fi21 < 13; fi21++ {
        switch v3 {
        case 128:
            v22 := GenS1{F0: 0, F1: 42}
            _ = v22
        default:
            var v23 int8 = -1
            _ = v23
        }
        v24 := v5["hello"]
        v24 = v24
        _ = fi21
        _ = v24
    }
    var v25 int64 = 139
    var v26 uint32 = (v4 >> ((4294967295) & 63))
    for fi27 := 0; fi27 < 9; fi27++ {
        var v28 string = "foo bar"
        var v29 int = 32768
        _ = fi27
        _ = v28
        _ = v29
    }
    var v30 int32 = 42
    delete(v5, "")
    if true {
        if true {
            fmt.Println(v1, v25)
            v3 = (v3 % (((v3 / ((v3) | 1))) | 1))
        } else {
            v31 := GenS3{F0: 1, F1: 0, F2: 42}
            _ = v31
        }
    } else {
        for fi32 := 0; fi32 < 18; fi32++ {
            var v33 uint = 255
            if ((!(true || true)) && ((42 == -1) && true)) {
                v34 := GenS5{F0: 18446744073709551615}
                var v35 string = "foo bar"
                _ = v34
                _ = v35
            } else {
                fmt.Println(v33, v4)
                v25 = v25
                var v36 int = (fi32 & (v3 & (v3 - (-(v3)))))
                _ = v36
            }
            _ = fi32
            _ = v33
        }
    }
    v37 := GenS6{F0: ((true || false) || false), F1: 128, F2: 2147483647}
    var v38 uint64 = 18446744073709551615
    v39 := GenS7{F0: -1, F1: ("a" + ("" + (("hello" + ("a" + "foo bar")) + "hello"))), F2: 42}
    for fi40 := 0; fi40 < 16; fi40++ {
        var v41 int32 = (v30 + int32(v3))
        v25 = (v25 >> ((591) & 63))
        if true {
            v42 := v5[""]
            _ = v42
        }
        _ = fi40
        _ = v41
    }
    v43 := GenS2{F0: -128}
    v38 -= 480
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v43
    _ = v5
}
