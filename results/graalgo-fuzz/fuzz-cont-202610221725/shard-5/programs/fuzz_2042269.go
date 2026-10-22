package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 uint8 }

func main() {
    v1 := []int{9223372036854775807, 2147483647}
    v2 := []string{"foo bar", "foo bar"}
    v3 := make(map[string]bool)
    var v4 string = (("foo bar" + "hello") + ("" + "\n\t"))
    var v5 uint64 = 339
    v6 := make(map[uint32]string)
    switch ((v4 + v4) + (v4 + "")) {
    case "":
        fmt.Println(v4)
    case "hello":
        v7 := GenS1{F0: 127}
        _ = v7
    }
    switch 1 {
    case 127:
        v8 := GenS2{F0: 4294967295, F1: -278}
        _ = v8
    case -1:
        v4 = (v4 + "\n\t")
    case -271:
        v5 = v5
    }
    var v9 int16 = -1
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 int = (fi10 << ((18446744073709551615) & 63))
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        if false {
            v5 = (v5 << ((4294967295) & 63))
        } else {
            for fi13 := 0; fi13 < 6; fi13++ {
                v5 -= v5
                var v14 int32 = -2147483648
                _ = fi13
                _ = v14
            }
        }
        _ = fi12
    }
    v15 := GenS1{F0: 127}
    var v16 int16 = 1
    switch 32768 {
    case -2147483648:
        var v17 float32 = float32((1.5 + ((0.0 * 1e10) + 1.5)))
        _ = v17
    case 127:
        if true {
            v18 := GenS3{F0: 2147483647, F1: -9223372036854775808}
            for fi19 := 0; fi19 < 4; fi19++ {
                v20 := GenS1{F0: 127}
                _ = fi19
                _ = v20
            }
            for fi21 := 0; fi21 < 19; fi21++ {
                switch "a" {
                case "hello":
                    v5 = v5
                case "\n\t":
                    var v22 int64 = 801
                    _ = v22
                case "foo bar":
                    var v23 float64 = 0.1
                    _ = v23
                }
                _ = fi21
            }
            _ = v18
        } else {
            if true {
                v9 -= (-(v16))
                v5 = (v5 ^ 307)
            }
        }
    }
    var v24 string = ((("" + "a") + v4) + v4)
    var v25 int8 = 0
    if len(v2) > 0 {
        v2[0] = v24
    }
    if false {
        switch -2147483648 {
        case -9223372036854775808:
            for fi26 := 0; fi26 < 13; fi26++ {
                for fi27 := 0; fi27 < 7; fi27++ {
                    if (-9223372036854775808 < 407) {
                        switch 32767 {
                        case 857:
                            v28 := v3["\n\t"]
                            _ = v28
                        case 32767:
                            v29 := GenS4{F0: 255}
                            _ = v29
                        case 42:
                            v16 = v16
                        }
                        v16 -= 731
                        switch (^(fi27)) {
                        case 4294967295:
                            v30 := v3[""]
                            _ = v30
                        case 128:
                            if len(v2) > 0 {
                                v2[0] = v4
                            }
                        default:
                            v5 = (^(v5))
                            if false {
                                var v31 float32 = 1e10
                                fmt.Println(v9, fi27, v4)
                                _ = v31
                            } else {
                                var v32 float64 = (float64(1e10) / (-2.25 - ((-2.25 - (3.14159 / 1e10)) / 3.14159)))
                                var v33 float64 = float64(0.0)
                                _ = v32
                                _ = v33
                            }
                        }
                    }
                    _ = fi27
                }
                _ = fi26
            }
        default:
            switch 0 {
            case 408:
                for fi34 := 0; fi34 < 2; fi34++ {
                    switch v24 {
                    case "\n\t":
                        v9 += (^(v16))
                    case "":
                        var v35 uint8 = 42
                        _ = v35
                    case "a":
                        _ = len(v3)
                    default:
                        var v36 int8 = (^(v25))
                        _ = v36
                    }
                    _ = fi34
                }
            case 128:
                v24 = (((v24 + "a") + v24) + v24)
            }
        }
        switch "hello" {
        case "hello":
            fmt.Println(v5, v5, v24)
        case "":
            v9 -= (v16 & v16)
        default:
            fmt.Println(v9, v5)
        }
    } else {
        var v37 int8 = (v25 << ((42) & 63))
        fmt.Println(v5, v37)
        var v38 uint8 = 179
        _ = v37
        _ = v38
    }
    v39 := GenS2{F0: 2, F1: 470}
    v40 := GenS3{F0: -97, F1: 0}
    var v41 uint16 = 65535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
    _ = v9
}
