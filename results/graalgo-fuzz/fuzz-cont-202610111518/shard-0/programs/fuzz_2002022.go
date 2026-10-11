package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool; F1 int64; F2 int8 }
func genFn0(p0 int, p1 string, p2 bool) uint8 {
    var v7 int64 = -54
    v8 := GenS1{F0: 92, F1: 32767}
    p2 = p2
    v7 = (v7 ^ 1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 142
}
func genFn1() int8 {
    var v9 uint = 2
    for fi10 := 0; fi10 < 17; fi10++ {
        v11 := GenS1{F0: 1, F1: 929}
        var v12 string = "a"
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 string = ("" + ("foo bar" + ("hello" + "\n\t")))
    v9 = 694
    v13 = (v13 + (v13 + (("hello" + (v13 + v13)) + "hello")))
    _ = v13
    _ = v9
    return 127
}

func main() {
    v1 := []int{-568, -1, -634}
    v2 := make(map[string]string)
    v3 := []string{"foo bar"}
    var v4 uint32 = 1
    var v5 uint64 = 42
    var v6 string = ""
    var v14 string = (v6 + v6)
    v5 -= v5
    if false {
        switch v6 {
        case "hello":
            fmt.Println(v14, v6, v6)
        case "a":
            if false {
                delete(v2, (v6 + v6))
                var v15 bool = false
                _ = v15
            } else {
                var v16 float32 = 0.0
                switch 128 {
                case -2147483648:
                    for fi17 := 0; fi17 < 14; fi17++ {
                        switch ("a" + (v14 + ("" + (((((("foo bar" + (v6 + ("a" + v14))) + "a") + v14) + v14) + (v6 + v14)) + "")))) {
                        case "foo bar":
                            if true {
                                fmt.Println(v14, v16, fi17)
                            } else {
                                fmt.Println(v14)
                                v6 = (v6 + v14)
                            }
                        default:
                            v18 := GenS1{F0: 29, F1: 248}
                            _ = v18
                        }
                        v16 -= v16
                        switch 2147483647 {
                        case 2147483647:
                            fmt.Println(v5, v6, v5)
                        case 127:
                            v16 = v16
                        case -1:
                            v19 := GenS1{F0: 2, F1: 255}
                            _ = v19
                        }
                        _ = fi17
                    }
                case 42:
                    v1 = append(v1, 184)
                case 32768:
                    switch 4294967295 {
                    case -1:
                        switch v14 {
                        case "hello":
                            if false {
                                var v20 int64 = -264
                                var v21 uint32 = (v4 & 12)
                                _ = v20
                                _ = v21
                            }
                        default:
                            if false {
                                v16 = -2.25
                            }
                        }
                    case 86:
                        v22 := GenS2{F0: 32767, F1: 2}
                        _ = v22
                    default:
                        v5 = v5
                    }
                default:
                    switch -9223372036854775808 {
                    case 127:
                        if true {
                            v5 -= v5
                            if true {
                                if false {
                                    var v23 uint8 = 2
                                    _ = v23
                                }
                                fmt.Println(v6)
                                var v24 uint64 = uint64(v4)
                                _ = v24
                            } else {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                for fi25 := 0; fi25 < 18; fi25++ {
                                    var v26 bool = (!(65535 == uint16(v5)))
                                    v14 = ("" + "\n\t")
                                    _ = fi25
                                    _ = v26
                                }
                                fmt.Println(v14, v6)
                            }
                            v14 = v14
                        } else {
                            v6 = (v14 + (v6 + v14))
                        }
                    case 4294967295:
                        fmt.Println(v4, v14)
                    }
                }
                v6 = v6
                _ = v16
            }
        }
        for fi27 := 0; fi27 < 13; fi27++ {
            var v28 string = "foo bar"
            _ = fi27
            _ = v28
        }
        switch 2147483647 {
        case -681:
            var v29 uint16 = 32767
            _ = v29
        }
    }
    var v30 int16 = -546
    v31 := GenS3{F0: uint(v4)}
    var v32 uint64 = 1
    _ = len(v2)
    v33 := GenS4{F0: true, F1: -9223372036854775808, F2: 127}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
