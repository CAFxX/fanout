package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32; F1 uint64; F2 int }
type GenS5 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS6 struct { F0 uint16; F1 uint32; F2 int16 }
func genFn0() bool {
    v7 := []int{127, 32768}
    v8 := GenS1{F0: -78, F1: 17}
    for fi9 := 0; fi9 < 17; fi9++ {
        _ = len(v7)
        _ = fi9
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        _ = len(v7)
        v11 := GenS1{F0: 127, F1: 127}
        _ = fi10
        _ = v11
    }
    var v12 int16 = -1
    _ = v12
    _ = v7
    _ = v8
    return ("foo bar" == "")
}

func main() {
    var v1 uint16 = 42
    var v2 int16 = -1
    v3 := []int{199}
    var v4 float32 = 3.14159
    var v5 uint = 701
    var v6 string = ""
    switch (v6 + "\n\t") {
    case "\n\t":
        v3 = append(v3, 1)
    case "foo bar":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "a":
        v13 := GenS1{F0: 111, F1: -33}
        _ = v13
    default:
        var v14 float64 = -2.25
        _ = v14
    }
    if false {
        for fi15 := 0; fi15 < 5; fi15++ {
            var v16 int32 = 1
            switch 127 {
            case -143:
                for fi17 := 0; fi17 < 19; fi17++ {
                    v5 = (v5 - v5)
                    if (true && true) {
                        v18 := GenS2{F0: v5, F1: 9223372036854775807}
                        switch 898 {
                        case 4294967295:
                            if false {
                                var v19 int16 = (v2 & (v2 ^ (v2 - (v2 ^ v2))))
                                fmt.Println(v19, v6)
                                _ = v19
                            }
                        case 9223372036854775807:
                            v20 := GenS3{F0: v1}
                            _ = v20
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        v6 = v6
                        _ = v18
                    }
                    _ = fi17
                }
            default:
                for fi21 := 0; fi21 < 3; fi21++ {
                    var v22 float64 = (-2.25 / 1e10)
                    _ = fi21
                    _ = v22
                }
            }
            _ = fi15
            _ = v16
        }
        v23 := GenS4{F0: 2147483647, F1: 42, F2: 4294967295}
        v3 = append(v3, 128)
        _ = v23
    }
    v2 += (-(v2))
    if (true || (-9223372036854775808 < 9223372036854775807)) {
        var v24 int8 = 106
        fmt.Println(v1)
        _ = v24
    }
    var v25 bool = ((v1 & (v1 << ((18446744073709551615) & 63))) != v1)
    v2 += 32767
    _ = len(v3)
    for fi26 := 0; fi26 < 5; fi26++ {
        if false {
            fmt.Println(v25, v2, v6)
            switch fi26 {
            case 2147483647:
                v3 = append(v3, (-(fi26)))
            case -313:
                var v27 uint16 = (^(v1))
                _ = v27
            case -6:
                var v28 float64 = 3.14159
                _ = v28
            default:
                v25 = v25
            }
            switch (fi26 | 4294967295) {
            case 32768:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case 9223372036854775807:
                if (1 >= 42) {
                    v29 := GenS4{F0: 59, F1: 0, F2: 0}
                    var v30 float64 = 1e10
                    v31 := GenS5{F0: 2147483648, F1: -2147483648, F2: uint8(v1)}
                    _ = v29
                    _ = v30
                    _ = v31
                }
            case 255:
                v1 = v1
            default:
                var v32 int8 = 100
                _ = v32
            }
        }
        v33 := GenS6{F0: v1, F1: 1, F2: 32767}
        var v34 string = ((v6 + (v6 + v6)) + "foo bar")
        _ = fi26
        _ = v33
        _ = v34
    }
    v1 -= v1
    fmt.Println(v6, v4, v1)
    var v35 int16 = 324
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (v25 || (1 != -2147483648)) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        for fi36 := 0; fi36 < 16; fi36++ {
            var v37 int64 = 42
            _ = fi36
            _ = v37
        }
    }
    v25 = (42 != 139)
    var v38 uint32 = 2147483647
    var v39 bool = ((42 >= 1) && (1 >= 255))
    var v40 int16 = (v35 ^ v35)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v35
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
