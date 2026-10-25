package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 int8 }
func genFn0(p0 int, p1 int8, p2 uint32) uint {
    var v7 bool = false
    var v8 uint16 = 539
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 848
}

func main() {
    var v1 float32 = 3.14159
    var v2 uint32 = 2
    var v3 float64 = 1e10
    v4 := make(map[int]bool)
    var v5 uint32 = (v2 * (v2 | (v2 + (v2 & (^(v2))))))
    var v6 int = 9223372036854775807
    v5 += (v5 + (^(v5)))
    switch ("" + ((("" + "foo bar") + "a") + (("hello" + "a") + "hello"))) {
    case "hello":
        var v9 string = ("hello" + "hello")
        _ = v9
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        fmt.Println(v2)
        _ = fi10
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        v6 = v6
        _ = fi11
    }
    var v12 uint32 = v5
    v4[v6] = (-1 != 127)
    v3 -= v3
    v13 := GenS1{F0: 2, F1: 1}
    var v14 uint8 = 1
    v15 := []int{(v6 * v6), -987}
    if len(v15) > 0 {
        v15[0] = 32768
    }
    var v16 int16 = 0
    _ = len(v4)
    var v17 int64 = -1
    switch "hello" {
    case "hello":
        delete(v4, v6)
    case "":
        if len(v15) > 0 {
            v15[0] = v6
        }
    case "foo bar":
        for fi18 := 0; fi18 < 10; fi18++ {
            switch fi18 {
            case 42:
                if false {
                    if (18446744073709551615 < 42) {
                        v12 = v12
                        for fi19 := 0; fi19 < 3; fi19++ {
                            switch "hello" {
                            case "":
                                if (!true) {
                                    v2 = v2
                                    var v20 int = (fi18 >> ((75) & 63))
                                    fmt.Println(v2, v1)
                                    _ = v20
                                }
                            default:
                                v17 = (v17 * int64(v12))
                            }
                            v21 := v4[(fi19 << ((801) & 63))]
                            _ = fi19
                            _ = v21
                        }
                    }
                    if false {
                        v22 := GenS1{F0: 288, F1: 0}
                        _ = v22
                    } else {
                        v23 := GenS2{F0: 100, F1: ("\n\t" + ("hello" + ("hello" + "")))}
                        _ = v23
                    }
                    switch 4294967295 {
                    case 4294967295:
                        fmt.Println(v12, v6)
                    case 9223372036854775807:
                        switch "" {
                        case "":
                            if (4294967295 >= 65535) {
                                var v24 float32 = 1.5
                                if len(v15) > 0 {
                                    v15[0] = v6
                                }
                                _ = v24
                            }
                        }
                    default:
                        switch "\n\t" {
                        case "a":
                            var v25 int32 = 0
                            _ = v25
                        default:
                            fmt.Println(v2, v2)
                        }
                    }
                } else {
                    var v26 string = "a"
                    v27 := GenS3{F0: -128}
                    _ = v26
                    _ = v27
                }
            case -261:
                var v28 uint8 = v14
                _ = v28
            case -9223372036854775808:
                switch (fi18 + v6) {
                case 13:
                    v16 = -32768
                default:
                    for fi29 := 0; fi29 < 5; fi29++ {
                        fmt.Println(v17, v17)
                        v6 = (v6 << ((166) & 63))
                        _ = fi29
                    }
                }
            default:
                v30 := GenS1{F0: 141, F1: 65535}
                _ = v30
            }
            var v31 float64 = (v3 - (((-2.25 / v3) - 3.14159) - (((-2.25 - float64(v1)) - v3) / v3)))
            for fi32 := 0; fi32 < 5; fi32++ {
                v4[(fi18 + (^(v6)))] = false
                v6 = (v6 | 4294967295)
                fmt.Println(v17)
                fmt.Println(v6)
                var v33 int8 = -1
                v6 += v6
                _ = fi32
                _ = v33
            }
            _ = fi18
            _ = v31
        }
    default:
        var v34 uint = 1
        _ = v34
    }
    var v35 uint8 = v14
    if (v1 < float32(v3)) {
        for fi36 := 0; fi36 < 1; fi36++ {
            if (!(!(false && false))) {
                if (!(v6 != v6)) {
                    fmt.Println(v17, v16, v6)
                    v14 = 127
                    v5 = v2
                }
            }
            v17 -= (v17 << ((65535) & 63))
            var v37 uint8 = v14
            _ = fi36
            _ = v37
        }
        if false {
            var v38 int = v6
            delete(v4, v38)
            v39 := GenS3{F0: 1}
            _ = v38
            _ = v39
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
