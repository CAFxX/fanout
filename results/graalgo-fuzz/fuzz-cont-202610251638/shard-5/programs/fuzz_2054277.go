package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS3 struct { F0 string; F1 int }
func genFn0(p0 int) int16 {
    switch (-(p0)) {
    case 42:
        p0 -= p0
    case 0:
        v6 := GenS1{F0: 1000}
        _ = v6
    case -703:
        p0 -= (p0 ^ (p0 * 4294967295))
    }
    if (1 <= 380) {
        p0 = (p0 - 9223372036854775807)
    }
    _ = p0
    return 425
}
func genFn1() uint8 {
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 float64 = (-2.25 - 1e10)
        v9 := GenS2{F0: 255, F1: 65535, F2: 4294967295}
        _ = fi7
        _ = v8
        _ = v9
    }
    if (-2147483648 > -2147483648) {
        for fi10 := 0; fi10 < 3; fi10++ {
            v11 := GenS3{F0: "", F1: fi10}
            v12 := GenS1{F0: -1}
            var v13 float64 = float64((float32(3.14159) / float32(((float64((0.0 / 1.5)) - 1e10) - (float64(1e10) * 1e10)))))
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        var v14 float32 = float32((0.1 + 1.5))
        _ = v14
    } else {
        var v15 float64 = -2.25
        var v16 int16 = -32768
        if (42 < 1) {
            var v17 int = 2147483647
            _ = v17
        } else {
            if (!(true && (0 == 100))) {
                fmt.Println(v16)
                var v18 uint = 65535
                _ = v18
            } else {
                v16 += (v16 * -1)
                switch 770 {
                case -9223372036854775808:
                    var v19 uint = 65535
                    _ = v19
                }
                v20 := []int{4294967295}
                _ = v20
            }
            var v21 uint8 = 121
            _ = v21
        }
        _ = v15
        _ = v16
    }
    return 255
}

func main() {
    var v1 string = "a"
    v2 := []int{}
    var v3 string = v1
    var v4 int32 = -1
    var v5 uint8 = 128
    var v22 int32 = 820
    v23 := GenS2{F0: 628, F1: 2, F2: 4294967295}
    switch "foo bar" {
    case "hello":
        switch 1 {
        case 32767:
            v4 = v4
        case 128:
            fmt.Println(v1, v22)
        default:
            v4 += (-(v4))
        }
    case "\n\t":
        v5 -= v5
    case "foo bar":
        v4 -= 1
    }
    if false {
        var v24 float32 = -2.25
        _ = v24
    }
    var v25 uint64 = 340
    var v26 uint = 73
    var v27 uint64 = 42
    var v28 int32 = (v22 % (((v4 >> (((v26 << (((v26 * v26)) & 63))) & 63))) | 1))
    v2 = append(v2, 692)
    var v29 int8 = -46
    for fi30 := 0; fi30 < 17; fi30++ {
        for fi31 := 0; fi31 < 3; fi31++ {
            switch (fi31 >> ((v26) & 63)) {
            case 127:
                v22 += v22
            case -2147483648:
                switch "a" {
                case "hello":
                    v3 = (v1 + "a")
                case "a":
                    var v32 uint8 = (v5 | v5)
                    _ = v32
                case "foo bar":
                    switch (^(fi31)) {
                    case 2147483647:
                        var v33 float32 = -2.25
                        _ = v33
                    case -9223372036854775808:
                        if false {
                            v34 := GenS1{F0: -32768}
                            _ = v34
                        }
                    default:
                        v4 = -1
                    }
                }
            case 2147483647:
                fmt.Println(v3, v25, v3)
            default:
                v35 := GenS2{F0: uint16(v26), F1: (^(v26)), F2: uint32(v5)}
                _ = v35
            }
            v29 += v29
            _ = fi31
        }
        for fi36 := 0; fi36 < 4; fi36++ {
            if len(v2) > 0 {
                v2[0] = 128
            }
            if ((^(v22)) <= (^(v4))) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if true {
                    v2 = append(v2, fi36)
                    var v37 bool = false
                    _ = v37
                } else {
                    switch (v3 + (v3 + "foo bar")) {
                    case "":
                        v5 -= (^(v5))
                        v25 = v27
                    }
                }
            } else {
                v2 = append(v2, fi30)
                v25 = (^(v25))
            }
            v5 = 42
            _ = fi36
        }
        _ = fi30
    }
    var v38 float64 = float64(1e10)
    fmt.Println(v4, v28)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
