package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS3 struct { F0 bool; F1 uint }

func main() {
    v1 := make(map[int]string)
    var v2 float64 = (-2.25 * (-2.25 / 1e10))
    var v3 int16 = 32767
    v4 := []uint32{0}
    var v5 bool = false
    v3 = v3
    fmt.Println(v3, v2)
    var v6 uint = 957
    v3 = v3
    switch "\n\t" {
    case "\n\t":
        v2 = -2.25
    }
    fmt.Println(v5)
    switch 32768 {
    case 4294967295:
        if true {
            var v7 string = ("a" + ("\n\t" + (("foo bar" + "a") + ("a" + ("\n\t" + "foo bar")))))
            v8 := GenS1{F0: 42, F1: 171, F2: -1}
            _ = v7
            _ = v8
        } else {
            v5 = (-1 == -380)
            for fi9 := 0; fi9 < 14; fi9++ {
                _ = len(v1)
                v3 = (v3 % (((v3 + (^(v3)))) | 1))
                _ = fi9
            }
            switch "" {
            case "":
                if len(v4) > 0 {
                    v4[0] = 65535
                }
            case "\n\t":
                if v5 {
                    switch ("foo bar" + "foo bar") {
                    case "":
                        v6 = 698
                    case "hello":
                        if v5 {
                            v10 := GenS1{F0: 2147483647, F1: 1, F2: 1}
                            _ = v10
                        } else {
                            delete(v1, 9223372036854775807)
                            var v11 int32 = -2147483648
                            switch ("" + ((("foo bar" + "") + ("" + "a")) + ("hello" + "a"))) {
                            case "":
                                if ((-976 <= 0) || ("" > ("a" + "a"))) {
                                    v12 := v1[128]
                                    var v13 uint = v6
                                    _ = v12
                                    _ = v13
                                }
                            case "a":
                                var v14 int32 = (v11 >> (((v6 | 2)) & 63))
                                _ = v14
                            case "\n\t":
                                if (("" + (("a" + "foo bar") + "")) <= ((((("\n\t" + ("hello" + "\n\t")) + "hello") + "") + (("foo bar" + "\n\t") + ("foo bar" + "a"))) + "\n\t")) {
                                    v11 += (v11 >> ((369) & 63))
                                    v15 := GenS1{F0: (v11 - 0), F1: 905, F2: v11}
                                    switch -1 {
                                    case 127:
                                        fmt.Println(v5)
                                    case 2147483647:
                                        for fi16 := 0; fi16 < 3; fi16++ {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            var v17 uint64 = 852
                                            if v5 {
                                                var v18 uint16 = 42
                                                v6 = 896
                                                var v19 int16 = v3
                                                _ = v18
                                                _ = v19
                                            }
                                            _ = fi16
                                            _ = v17
                                        }
                                    default:
                                        var v20 uint8 = uint8(v6)
                                        _ = v20
                                    }
                                    _ = v15
                                } else {
                                    for fi21 := 0; fi21 < 10; fi21++ {
                                        switch ("a" + "hello") {
                                        case "hello":
                                            v4 = append(v4, 2147483648)
                                        case "":
                                            v22 := GenS2{F0: v6, F1: 618, F2: 1}
                                            _ = v22
                                        case "a":
                                            if false {
                                                fmt.Println(v2)
                                                for fi23 := 0; fi23 < 5; fi23++ {
                                                    v11 = v11
                                                    v5 = (v3 == -851)
                                                    v24 := GenS2{F0: v6, F1: 278, F2: 42}
                                                    _ = fi23
                                                    _ = v24
                                                }
                                            }
                                        }
                                        _ = fi21
                                    }
                                    for fi25 := 0; fi25 < 15; fi25++ {
                                        fmt.Println(v6, v6, v6)
                                        _ = fi25
                                    }
                                    switch "hello" {
                                    case "a":
                                        v26 := v1[-574]
                                        _ = v26
                                    default:
                                        v27 := GenS3{F0: (128 < 0), F1: 1}
                                        _ = v27
                                    }
                                }
                            default:
                                var v28 uint = v6
                                _ = v28
                            }
                            _ = v11
                        }
                    default:
                        v5 = false
                    }
                    if (("foo bar" + ("hello" + "hello")) >= ("foo bar" + "hello")) {
                        var v29 uint16 = 42
                        v29 += 255
                        if v5 {
                            v5 = ((9223372036854775807 != 42) || ((v5 || (1 < 449)) || v5))
                        }
                        _ = v29
                    }
                }
            case "a":
                v4 = append(v4, 430)
            }
        }
    }
    v30 := GenS2{F0: v6, F1: 65535, F2: 604}
    switch "hello" {
    case "\n\t":
        v4 = append(v4, 4294967295)
    case "":
        v5 = v5
    case "foo bar":
        var v31 uint32 = 87
        _ = v31
    default:
        v5 = v5
    }
    fmt.Println(v6, v5, v6)
    v3 -= v3
    var v32 uint = v6
    for fi33 := 0; fi33 < 18; fi33++ {
        v34 := GenS3{F0: v5, F1: (v32 & (v32 - v6))}
        v6 = (v32 & (v32 & (v32 | 723)))
        _ = fi33
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
