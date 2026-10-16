package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 uint; F2 int16 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]string)
    v3 := []int{128, 128, 4294967295}
    var v4 uint32 = 775
    var v5 uint16 = 65535
    fmt.Println(v5, v5)
    if false {
        if (false || ((v4 | (v4 << ((207) & 63))) == v4)) {
            switch (("foo bar" + "a") + (("hello" + "a") + "\n\t")) {
            case "hello":
                v6 := GenS1{F0: 18446744073709551615}
                _ = v6
            case "foo bar":
                v5 += v5
            case "\n\t":
                switch "hello" {
                case "foo bar":
                    var v7 int = -1
                    _ = v7
                case "":
                    v8 := v1["hello"]
                    _ = v8
                case "a":
                    switch -9223372036854775808 {
                    case -9223372036854775808:
                        var v9 int64 = int64(v4)
                        _ = v9
                    case -229:
                        switch 32767 {
                        case -1:
                            var v10 uint16 = uint16(v4)
                            _ = v10
                        case -121:
                            if ("foo bar" == "hello") {
                                switch "a" {
                                case "a":
                                    var v11 uint8 = 80
                                    _ = v11
                                default:
                                    var v12 int16 = 0
                                    _ = v12
                                }
                                var v13 int16 = -90
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v13
                            }
                        case 32767:
                            var v14 uint = 1
                            _ = v14
                        default:
                            v5 -= v5
                        }
                    case 128:
                        for fi15 := 0; fi15 < 11; fi15++ {
                            fmt.Println(v5, fi15, v4)
                            fmt.Println(v4, v5)
                            var v16 uint64 = 0
                            _ = fi15
                            _ = v16
                        }
                    }
                }
            }
            v5 = 255
            v5 = v5
        } else {
            var v17 float32 = 0.0
            _ = v17
        }
    } else {
        var v18 int64 = -53
        v19 := GenS2{F0: 18446744073709551615, F1: uint(v5), F2: 42}
        if ((v5 >> ((18446744073709551615) & 63)) < v5) {
            if ((1.5 - 1e10) >= 3.14159) {
                fmt.Println(v4, v18, v18)
                var v20 float32 = 3.14159
                switch ("hello" + "a") {
                case "\n\t":
                    for fi21 := 0; fi21 < 11; fi21++ {
                        var v22 int32 = -1
                        _ = fi21
                        _ = v22
                    }
                case "a":
                    _ = len(v2)
                case "":
                    if true {
                        for fi23 := 0; fi23 < 3; fi23++ {
                            var v24 int = fi23
                            v5 -= (v5 % (((v5 / ((558) | 1))) | 1))
                            _ = fi23
                            _ = v24
                        }
                    } else {
                        var v25 float64 = -2.25
                        _ = v25
                    }
                }
                _ = v20
            }
        }
        _ = v18
        _ = v19
    }
    v26 := GenS1{F0: 18446744073709551615}
    if true {
        v4 += (^(v4))
    }
    v27 := v1["\n\t"]
    v5 -= (v5 >> ((2) & 63))
    v4 = v27
    fmt.Println(v5)
    v28 := v1["\n\t"]
    var v29 uint = 0
    var v30 string = "a"
    v28 = (v27 - (v4 & (v27 * uint32(v5))))
    var v31 uint64 = 4294967295
    v32 := GenS1{F0: (^(v29))}
    v1[v30] = 813
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
