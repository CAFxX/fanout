package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int16; F1 int32; F2 int8 }
func genFn0(p0 uint8) uint64 {
    for fi4 := 0; fi4 < 4; fi4++ {
        for fi5 := 0; fi5 < 4; fi5++ {
            p0 = 2
            v6 := GenS1{F0: 9223372036854775807, F1: 100}
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    var v7 uint32 = 178
    _ = p0
    _ = v7
    return uint64(p0)
}
func genFn1() int32 {
    switch -9223372036854775808 {
    case 4294967295:
        v8 := []int{475, -80, 294}
        _ = v8
    }
    var v9 uint16 = 644
    v10 := GenS1{F0: 0, F1: 0}
    v11 := GenS1{F0: 0, F1: 41}
    if ((^(v9)) <= v9) {
        if false {
            fmt.Println(v9)
            switch "\n\t" {
            case "\n\t":
                if true {
                    var v12 int16 = -1
                    _ = v12
                }
            case "foo bar":
                v9 = v9
            default:
                v9 -= 23
            }
        }
    }
    _ = v10
    _ = v11
    _ = v9
    return -143
}

func main() {
    var v1 int32 = 0
    var v2 bool = false
    var v3 int32 = v1
    v13 := []int{}
    switch ("a" + "\n\t") {
    case "a":
        var v14 string = "hello"
        _ = v14
    case "":
        v1 -= 2147483647
    case "foo bar":
        v3 = 179
    default:
        var v15 int8 = 1
        _ = v15
    }
    var v16 string = "foo bar"
    var v17 uint16 = 2
    if len(v13) > 0 {
        v13[0] = 9223372036854775807
    }
    var v18 uint8 = 128
    for fi19 := 0; fi19 < 20; fi19++ {
        for fi20 := 0; fi20 < 17; fi20++ {
            switch 0 {
            case 128:
                if ((false || ((fi20 * int(v3)) >= 127)) && true) {
                    v21 := GenS1{F0: -759, F1: 0}
                    v2 = (!v2)
                    if v2 {
                        if ((v18 + 42) == v18) {
                            var v22 bool = v2
                            _ = v22
                        }
                        v23 := GenS1{F0: 9223372036854775807, F1: 100}
                        _ = v23
                    }
                    _ = v21
                } else {
                    v24 := GenS1{F0: 1, F1: int8(v3)}
                    v1 = 2147483647
                    _ = v24
                }
            case -9223372036854775808:
                var v25 int64 = 0
                _ = v25
            case 32767:
                var v26 float32 = 0.0
                _ = v26
            default:
                v17 = v17
            }
            if (!false) {
                v1 -= genFn1()
                v16 = "a"
            } else {
                var v27 uint16 = v17
                v28 := GenS2{F0: int16(fi20), F1: (v1 | (^(v1))), F2: 42}
                _ = v27
                _ = v28
            }
            var v29 string = (v16 + ("" + v16))
            _ = fi20
            _ = v29
        }
        var v30 uint32 = 2147483647
        _ = fi19
        _ = v30
    }
    switch 2147483647 {
    case 1:
        v13 = append(v13, -535)
    case -9223372036854775808:
        switch v16 {
        case "\n\t":
            if len(v13) > 0 {
                v13[0] = 128
            }
        default:
            if false {
                var v31 uint16 = (v17 ^ 255)
                v2 = (1 == 42)
                if (32767 >= 0) {
                    v18 = (v18 | 66)
                    for fi32 := 0; fi32 < 9; fi32++ {
                        var v33 uint32 = 2147483647
                        var v34 int64 = -9223372036854775808
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                } else {
                    var v35 int = 32768
                    if v2 {
                        var v36 string = "foo bar"
                        _ = v36
                    } else {
                        var v37 uint = 433
                        fmt.Println(v37, v37, v18)
                        var v38 int64 = 478
                        _ = v37
                        _ = v38
                    }
                    for fi39 := 0; fi39 < 15; fi39++ {
                        fmt.Println(fi39)
                        for fi40 := 0; fi40 < 10; fi40++ {
                            v2 = (v31 > (v17 | v17))
                            _ = fi40
                        }
                        _ = fi39
                    }
                    _ = v35
                }
                _ = v31
            } else {
                _ = len(v13)
            }
        }
    }
    for fi41 := 0; fi41 < 1; fi41++ {
        var v42 int = (fi41 * 0)
        _ = fi41
        _ = v42
    }
    if len(v13) > 0 {
        v13[0] = -248
    }
    var v43 int = -9223372036854775808
    if len(v13) > 0 {
        v13[0] = -482
    }
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    var v44 uint16 = 2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v43
    _ = v44
}
