package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 int64; F1 int16 }
func genFn0(p0 int64) bool {
    var v4 bool = ((p0 / ((p0) | 1)) > 1)
    v5 := GenS1{F0: -119, F1: 100, F2: 369}
    v4 = (-32768 < 169)
    _ = p0
    _ = v4
    _ = v5
    return v4
}

func main() {
    var v1 int64 = 0
    var v2 float32 = -2.25
    var v3 int16 = 1000
    v6 := GenS2{F0: 180}
    v2 = (v2 / 1.5)
    var v7 uint = 65535
    var v8 uint16 = 255
    fmt.Println(v1, v7)
    var v9 float32 = float32(0.1)
    v2 += -2.25
    if (!true) {
        v3 = v3
        v10 := GenS1{F0: int8(v7), F1: 100, F2: 18446744073709551615}
        _ = v10
    } else {
        v11 := GenS3{F0: -56, F1: v3}
        if (42 > int32(v7)) {
            v12 := []int{9223372036854775807}
            fmt.Println(v7, v2)
            _ = v12
        } else {
            for fi13 := 0; fi13 < 8; fi13++ {
                v7 = v7
                _ = fi13
            }
            if true {
                var v14 int16 = 270
                fmt.Println(v8, v1)
                if false {
                    if (!(!false)) {
                        v7 += v7
                        v15 := []int{}
                        _ = v15
                    } else {
                        var v16 uint = (v7 - v7)
                        if true {
                            v7 += (v16 + 0)
                            switch 9223372036854775807 {
                            case 32768:
                                v17 := []int{1, 127, -563, -135}
                                _ = v17
                            case 42:
                                var v18 bool = (((1.5 / 1.5) / 0.1) <= (3.14159 + 3.14159))
                                _ = v18
                            case 1:
                                switch 2147483647 {
                                case 144:
                                    switch "hello" {
                                    case "hello":
                                        v19 := GenS3{F0: -1, F1: 1}
                                        _ = v19
                                    default:
                                        fmt.Println(v7, v8)
                                    }
                                case -208:
                                    var v20 uint = v7
                                    _ = v20
                                }
                            }
                            var v21 uint64 = 42
                            _ = v21
                        }
                        _ = v16
                    }
                    var v22 uint = v7
                    switch ("a" + "foo bar") {
                    case "\n\t":
                        var v23 int8 = 78
                        _ = v23
                    case "foo bar":
                        if false {
                            v2 = (v2 - v2)
                            var v24 int = -282
                            for fi25 := 0; fi25 < 3; fi25++ {
                                var v26 int32 = -2147483648
                                switch "hello" {
                                case "a":
                                    v27 := []int{fi25}
                                    _ = v27
                                case "hello":
                                    for fi28 := 0; fi28 < 20; fi28++ {
                                        fmt.Println(fi25, fi25)
                                        var v29 float32 = v9
                                        _ = fi28
                                        _ = v29
                                    }
                                case "\n\t":
                                    if ((^(v26)) < (v26 / (((v26 % (((^(v26))) | 1))) | 1))) {
                                        v30 := []int{(fi25 ^ (fi25 * v24)), (fi25 - v24)}
                                        fmt.Println(v1)
                                        _ = v30
                                    }
                                }
                                _ = fi25
                                _ = v26
                            }
                            _ = v24
                        }
                    case "hello":
                        for fi31 := 0; fi31 < 7; fi31++ {
                            for fi32 := 0; fi32 < 12; fi32++ {
                                switch ("hello" + "a") {
                                case "foo bar":
                                    v33 := GenS2{F0: 58}
                                    _ = v33
                                case "hello":
                                    v34 := GenS1{F0: 55, F1: 1, F2: 912}
                                    _ = v34
                                }
                                var v35 uint64 = 1
                                switch "\n\t" {
                                case "hello":
                                    v36 := []int{(fi31 & (fi32 ^ 2147483647)), (-(fi31)), (fi32 << ((v7) & 63)), (fi31 - fi32)}
                                    _ = v36
                                case "\n\t":
                                    var v37 string = ((("hello" + "hello") + ("" + ("hello" + ("" + ("" + "foo bar"))))) + "\n\t")
                                    _ = v37
                                default:
                                    var v38 uint16 = (^(v8))
                                    _ = v38
                                }
                                _ = fi32
                                _ = v35
                            }
                            v39 := []int{fi31, fi31, fi31, (fi31 + fi31)}
                            _ = fi31
                            _ = v39
                        }
                    default:
                        v40 := []int{4294967295, 4294967295, -662, 32767}
                        _ = v40
                    }
                    _ = v22
                }
                _ = v14
            } else {
                v41 := GenS2{F0: 128}
                _ = v41
            }
            v42 := []int{413}
            _ = v42
        }
        for fi43 := 0; fi43 < 9; fi43++ {
            var v44 uint64 = 2
            if true {
                if (!false) {
                    var v45 uint8 = 42
                    _ = v45
                }
                v7 = (v7 ^ v7)
                v46 := []int{}
                _ = v46
            }
            var v47 uint16 = v8
            _ = fi43
            _ = v44
            _ = v47
        }
        _ = v11
    }
    v8 = v8
    v48 := GenS4{F0: 42, F1: v3}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v48
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
