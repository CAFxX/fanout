package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 uint16; F2 string }
type GenS3 struct { F0 int32; F1 uint8; F2 bool }
func genFn0(p0 uint32) int8 {
    var v7 string = ("a" + ("foo bar" + "\n\t"))
    fmt.Println(v7, p0)
    _ = p0
    _ = v7
    return -21
}
func genFn1() uint32 {
    var v8 uint16 = 42
    if false {
        v9 := []int{472, 127, 0}
        var v10 float32 = -2.25
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        v8 = 160
        _ = fi11
    }
    _ = v8
    return 252
}

func main() {
    var v1 uint16 = 1
    var v2 float32 = -2.25
    var v3 string = ("hello" + ("a" + ("" + "\n\t")))
    var v4 int16 = -32768
    var v5 float32 = 3.14159
    var v6 uint64 = 1
    fmt.Println(v6)
    var v12 uint8 = 136
    v1 = v1
    switch 42 {
    case 127:
        for fi13 := 0; fi13 < 12; fi13++ {
            if (true && (887 >= 2)) {
                v12 -= (^(v12))
                v5 -= v2
                var v14 uint = 42
                _ = v14
            } else {
                var v15 uint8 = v12
                for fi16 := 0; fi16 < 1; fi16++ {
                    v17 := GenS1{F0: 4294967295}
                    _ = fi16
                    _ = v17
                }
                v18 := []int{128, (-(fi13)), 32768, 9223372036854775807}
                _ = v15
                _ = v18
            }
            var v19 float64 = 1.5
            _ = fi13
            _ = v19
        }
    case -1:
        if (false && true) {
            v6 -= v6
        } else {
            for fi20 := 0; fi20 < 8; fi20++ {
                var v21 int64 = 9223372036854775807
                v22 := []int{(fi20 / (((fi20 & fi20)) | 1)), fi20, (fi20 - (fi20 % ((329) | 1))), (fi20 * fi20)}
                switch v3 {
                case "a":
                    v23 := GenS1{F0: 42}
                    _ = v23
                case "hello":
                    var v24 int64 = v21
                    _ = v24
                case "":
                    var v25 uint16 = 149
                    _ = v25
                }
                _ = fi20
                _ = v21
                _ = v22
            }
        }
    case -1000:
        if (genFn1() >= 2147483648) {
            var v26 uint16 = v1
            var v27 uint64 = (^(v6))
            _ = v26
            _ = v27
        }
    }
    var v28 int16 = int16(v6)
    var v29 uint = 598
    v2 = -2.25
    v30 := GenS2{F0: (v29 ^ v29), F1: v1, F2: (v3 + v3)}
    for fi31 := 0; fi31 < 10; fi31++ {
        v29 = v29
        switch (("\n\t" + "\n\t") + v3) {
        case "a":
            v32 := GenS3{F0: 42, F1: uint8(v1), F2: (-1 < 127)}
            _ = v32
        case "\n\t":
            for fi33 := 0; fi33 < 8; fi33++ {
                var v34 uint32 = 4294967295
                v35 := GenS3{F0: 884, F1: v12, F2: ((v4 * -32768) != v28)}
                fmt.Println(v5, v1, v28)
                if true {
                    v12 = (v12 % ((1) | 1))
                    var v36 bool = ((v34 + v34) <= (v34 ^ 90))
                    var v37 bool = v36
                    var v38 float64 = (0.0 * (3.14159 * 3.14159))
                    fmt.Println(v38, v34)
                    v12 -= v12
                    v5 -= (v5 + 1.5)
                    v39 := []int{}
                    fmt.Println(v3, v36)
                    fmt.Println(v4, v28)
                    v3 = (v3 + "hello")
                    _ = v36
                    _ = v37
                    _ = v38
                    _ = v39
                } else {
                    switch (fi31 ^ fi31) {
                    case -9223372036854775808:
                        v12 = v12
                    case 9223372036854775807:
                        if true {
                            fmt.Println(v6, v6)
                        }
                        if false {
                            v40 := []int{fi31, -36, fi33, fi31}
                            if len(v40) > 0 {
                                v40[0] = fi31
                            }
                            v5 += ((v5 - v5) / v2)
                            _ = v40
                        }
                    case -894:
                        var v41 uint32 = (^(v34))
                        _ = v41
                    default:
                        var v42 int64 = 9223372036854775807
                        _ = v42
                    }
                    switch "a" {
                    case "\n\t":
                        var v43 int = fi33
                        _ = v43
                    case "foo bar":
                        v44 := GenS2{F0: (^(v29)), F1: 562, F2: v3}
                        _ = v44
                    case "a":
                        if (true || true) {
                            var v45 uint = v29
                            var v46 string = "\n\t"
                            _ = v45
                            _ = v46
                        }
                    default:
                        fmt.Println(v34, v12)
                    }
                }
                v47 := []int{}
                _ = fi33
                _ = v34
                _ = v35
                _ = v47
            }
        }
        _ = fi31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
