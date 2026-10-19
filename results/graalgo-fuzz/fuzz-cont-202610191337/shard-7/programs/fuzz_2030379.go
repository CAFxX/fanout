package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint16; F1 string }
func genFn0(p0 int64, p1 uint8, p2 int8) int32 {
    v4 := []int{-9223372036854775808, 4294967295, 127}
    v4 = append(v4, -1)
    if len(v4) > 0 {
        v4[0] = 32767
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 323
}

func main() {
    var v1 uint32 = 42
    v2 := []int{-603}
    var v3 int8 = -119
    if false {
        if (((42 >= -32768) && true) && (32767 <= 2)) {
            fmt.Println(v3, v1)
        } else {
            switch -9223372036854775808 {
            case -1:
                v2 = append(v2, 0)
            case 4294967295:
                var v5 uint32 = v1
                _ = v5
            default:
                v1 += (v1 << ((18446744073709551615) & 63))
            }
        }
        switch "" {
        case "a":
            switch "foo bar" {
            case "\n\t":
                switch 0 {
                case 397:
                    v1 -= v1
                case 42:
                    v3 = (v3 & v3)
                case 0:
                    v3 = -1
                default:
                    fmt.Println(v3, v1)
                }
            case "foo bar":
                var v6 bool = true
                _ = v6
            case "":
                var v7 int = 515
                _ = v7
            default:
                var v8 string = "a"
                _ = v8
            }
        case "foo bar":
            v9 := GenS1{F0: false}
            _ = v9
        case "":
            if (-2147483648 != genFn0(42, 42, (v3 * (v3 - 127)))) {
                switch -635 {
                case 128:
                    for fi10 := 0; fi10 < 11; fi10++ {
                        if (false && false) {
                            switch int(v1) {
                            case 907:
                                v3 = 83
                            }
                            v3 += v3
                            fmt.Println(v3)
                        }
                        if (262 == 18446744073709551615) {
                            v3 = v3
                            var v11 string = ("a" + "hello")
                            _ = v11
                        }
                        _ = fi10
                    }
                case 127:
                    if false {
                        var v12 int16 = -741
                        _ = v12
                    } else {
                        if (!(1 != int32(v3))) {
                            var v13 uint16 = 65535
                            switch "hello" {
                            case "":
                                v13 = 42
                            default:
                                var v14 int64 = -274
                                _ = v14
                            }
                            var v15 bool = (false || false)
                            _ = v13
                            _ = v15
                        }
                    }
                }
            } else {
                v16 := GenS1{F0: (v3 < 107)}
                if false {
                    switch ((("\n\t" + "\n\t") + ("" + "\n\t")) + (("hello" + "hello") + (("hello" + (("a" + "\n\t") + "")) + "\n\t"))) {
                    case "\n\t":
                        switch 889 {
                        case -174:
                            v3 = (v3 ^ int8(v1))
                        default:
                            v17 := GenS2{F0: 1, F1: 1}
                            _ = v17
                        }
                    default:
                        fmt.Println(v1)
                    }
                    v18 := GenS1{F0: false}
                    _ = v18
                }
                v1 -= (v1 % (((v1 * 42)) | 1))
                _ = v16
            }
        default:
            v1 -= 680
        }
        v3 -= (v3 * (v3 * v3))
    }
    v1 = 974
    v2 = append(v2, -990)
    var v19 float64 = float64((1e10 - 1.5))
    fmt.Println(v1)
    var v20 int8 = (^(v3))
    v3 += v20
    var v21 uint32 = v1
    v3 -= -128
    var v22 float32 = (0.0 * (((float32(1e10) * 1e10) * 1e10) + 1e10))
    var v23 float64 = v19
    if true {
        if false {
            for fi24 := 0; fi24 < 17; fi24++ {
                v25 := GenS3{F0: 144}
                for fi26 := 0; fi26 < 9; fi26++ {
                    v23 += float64((v22 * ((v22 / (0.1 + v22)) - 0.0)))
                    v1 = (v1 - (v1 << ((286) & 63)))
                    var v27 float32 = ((-2.25 / 1.0) + 0.0)
                    _ = fi26
                    _ = v27
                }
                _ = fi24
                _ = v25
            }
        }
    }
    for fi28 := 0; fi28 < 10; fi28++ {
        if true {
            if ((v3 + v3) <= v20) {
                v29 := GenS4{F0: "hello"}
                var v30 uint64 = 660
                _ = v29
                _ = v30
            }
            switch (-(fi28)) {
            case 70:
                switch ("" + "hello") {
                case "hello":
                    v1 += v1
                case "\n\t":
                    v31 := GenS5{F0: 4294967295}
                    _ = v31
                default:
                    switch fi28 {
                    case -2147483648:
                        var v32 uint8 = 177
                        _ = v32
                    }
                }
            }
        }
        var v33 int32 = -1
        var v34 uint16 = 1
        _ = fi28
        _ = v33
        _ = v34
    }
    switch 4294967295 {
    case -282:
        if true {
            var v35 float64 = v19
            v36 := GenS6{F0: 2}
            _ = v35
            _ = v36
        } else {
            v22 = v22
            v37 := GenS7{F0: 32767, F1: (((("foo bar" + "") + ("hello" + ("\n\t" + "foo bar"))) + "") + "")}
            _ = v37
        }
    }
    var v38 uint8 = 2
    v22 = (v22 + v22)
    var v39 int = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v38
    _ = v39
}
