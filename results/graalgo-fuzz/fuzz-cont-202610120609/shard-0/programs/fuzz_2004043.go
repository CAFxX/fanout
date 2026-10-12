package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int; F1 int32; F2 int32 }
type GenS3 struct { F0 int16; F1 int64; F2 uint64 }
func genFn0() uint64 {
    v6 := []int{}
    var v7 uint = 18446744073709551615
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    var v1 uint8 = 42
    var v2 int64 = -9223372036854775808
    var v3 float32 = -2.25
    var v4 uint32 = 4294967295
    var v5 uint32 = v4
    if true {
        v8 := GenS1{F0: "foo bar", F1: v1}
        _ = v8
    } else {
        switch "\n\t" {
        case "a":
            var v9 int32 = 42
            _ = v9
        case "":
            if false {
                var v10 uint16 = 65535
                v11 := []int{-2147483648, 996}
                _ = v10
                _ = v11
            } else {
                if (!(0 > 48)) {
                    v5 = 0
                    if true {
                        if (true && (!(!((((((v3 * v3) * ((float32(float64(((1e10 - v3) - v3))) - (v3 * (((v3 * v3) / v3) - -2.25))) - (v3 * v3))) - v3) + (1.5 + ((v3 * v3) + v3))) - float32(3.14159)) != v3)))) {
                            for fi12 := 0; fi12 < 9; fi12++ {
                                var v13 uint64 = 2
                                _ = fi12
                                _ = v13
                            }
                        } else {
                            var v14 float64 = (3.14159 - -2.25)
                            var v15 uint32 = (^(v5))
                            for fi16 := 0; fi16 < 7; fi16++ {
                                for fi17 := 0; fi17 < 1; fi17++ {
                                    v15 -= 122
                                    _ = fi17
                                }
                                switch "foo bar" {
                                case "hello":
                                    v2 = (v2 / (((v2 << ((255) & 63))) | 1))
                                default:
                                    var v18 int32 = 42
                                    _ = v18
                                }
                                _ = fi16
                            }
                            _ = v14
                            _ = v15
                        }
                    } else {
                        fmt.Println(v1)
                    }
                    switch "\n\t" {
                    case "\n\t":
                        v19 := []int{-2147483648, 127, int(v1)}
                        _ = v19
                    case "foo bar":
                        v20 := GenS1{F0: "a", F1: (^(v1))}
                        _ = v20
                    default:
                        var v21 uint = 2
                        _ = v21
                    }
                }
                v1 = (^(v1))
                var v22 float64 = 0.0
                _ = v22
            }
        default:
            var v23 string = (("\n\t" + "a") + ("hello" + ("" + "foo bar")))
            _ = v23
        }
        fmt.Println(v3)
        for fi24 := 0; fi24 < 20; fi24++ {
            v25 := []int{(^(fi24)), fi24, fi24, (fi24 ^ (fi24 << ((357) & 63)))}
            if len(v25) >= 2 {
                v25 = v25[1:len(v25)]
            }
            for fi26 := 0; fi26 < 4; fi26++ {
                fmt.Println(fi24, fi26, v1)
                switch fi26 {
                case -1:
                    v2 = 1
                case 935:
                    v27 := GenS2{F0: fi26, F1: 2147483647, F2: -2147483648}
                    _ = v27
                }
                _ = fi26
            }
            _ = fi24
            _ = v25
        }
    }
    v28 := GenS3{F0: 42, F1: (v2 * (v2 | (v2 | 42))), F2: 656}
    v2 = v2
    if true {
        switch ("foo bar" + ("a" + "hello")) {
        case "\n\t":
            var v29 string = (("hello" + "\n\t") + "")
            _ = v29
        case "":
            var v30 int64 = 516
            _ = v30
        case "foo bar":
            for fi31 := 0; fi31 < 3; fi31++ {
                v32 := GenS1{F0: "", F1: (v1 & v1)}
                v33 := []int{fi31, (-(fi31)), fi31}
                var v34 int16 = -396
                _ = fi31
                _ = v32
                _ = v33
                _ = v34
            }
        default:
            var v35 bool = ((("foo bar" + "hello") + "") > "foo bar")
            _ = v35
        }
        v36 := GenS2{F0: 4294967295, F1: 663, F2: -1}
        _ = v36
    } else {
        v3 -= (v3 + v3)
    }
    switch "hello" {
    case "hello":
        switch -365 {
        case -2147483648:
            v37 := GenS3{F0: 1, F1: (v2 & v2), F2: 2}
            _ = v37
        default:
            fmt.Println(v5, v3, v4)
        }
    default:
        for fi38 := 0; fi38 < 16; fi38++ {
            switch (fi38 ^ (fi38 ^ fi38)) {
            case 0:
                v3 += ((float32(1e10) - v3) * ((v3 * 0.0) / v3))
            case 4294967295:
                fmt.Println(v3, v2, v1)
            case 212:
                for fi39 := 0; fi39 < 6; fi39++ {
                    var v40 int16 = -32768
                    if false {
                        fmt.Println(v40)
                        fmt.Println(fi38, fi39, v40)
                    }
                    _ = fi39
                    _ = v40
                }
            default:
                v2 = -151
            }
            _ = fi38
        }
    }
    var v41 uint32 = v4
    switch 4294967295 {
    case 748:
        v1 = (v1 | 2)
    }
    v4 -= v4
    switch 177 {
    case 128:
        v42 := GenS3{F0: -32768, F1: (v2 + (v2 >> ((503) & 63))), F2: 2}
        _ = v42
    }
    switch 94 {
    case -1:
        v43 := GenS2{F0: 128, F1: -2147483648, F2: -2147483648}
        _ = v43
    case -147:
        var v44 int = -1
        _ = v44
    }
    var v45 int = -154
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v41
    _ = v45
    _ = v5
}
