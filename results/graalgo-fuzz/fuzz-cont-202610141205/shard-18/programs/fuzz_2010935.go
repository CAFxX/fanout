package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }

func main() {
    var v1 uint16 = 255
    var v2 uint8 = 255
    v3 := make(map[int]string)
    v4 := make(map[int]bool)
    var v5 string = "a"
    v1 -= 32767
    var v6 uint64 = uint64(v2)
    fmt.Println(v1, v1, v5)
    var v7 uint32 = 130
    v3[558] = "hello"
    v5 = v5
    v5 = v5
    v8 := []int{2147483647, 32767}
    v9 := v3[512]
    v6 -= v6
    var v10 float64 = 3.14159
    fmt.Println(v2, v2, v1)
    fmt.Println(v2, v6)
    v11 := GenS1{F0: 32767, F1: 1}
    var v12 float32 = 0.0
    switch -471 {
    case 4294967295:
        var v13 int8 = 1
        _ = v13
    case 32768:
        for fi14 := 0; fi14 < 18; fi14++ {
            v15 := GenS1{F0: 51, F1: int64(v7)}
            if true {
                if true {
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    for fi16 := 0; fi16 < 5; fi16++ {
                        if false {
                            v9 = v5
                            fmt.Println(v10, v9)
                        } else {
                            var v17 float64 = (v10 - float64(v12))
                            _ = v17
                        }
                        switch fi16 {
                        case 128:
                            fmt.Println(v1)
                        case -226:
                            fmt.Println(v5)
                        }
                        _ = fi16
                    }
                    if (!(true || false)) {
                        var v18 uint16 = (^(v1))
                        var v19 bool = false
                        for fi20 := 0; fi20 < 6; fi20++ {
                            if (0 > 504) {
                                v9 = "\n\t"
                                var v21 uint32 = (v7 - 2)
                                var v22 string = ("hello" + v5)
                                _ = v21
                                _ = v22
                            } else {
                                fmt.Println(v6, v1)
                            }
                            if v19 {
                                v8 = append(v8, (-(fi20)))
                            }
                            _ = fi20
                        }
                        _ = v18
                        _ = v19
                    }
                }
            }
            switch -1 {
            case -160:
                var v23 int32 = 1
                _ = v23
            case -2147483648:
                switch v9 {
                case "foo bar":
                    fmt.Println(v7, v9)
                case "\n\t":
                    v24 := GenS1{F0: -1, F1: -1}
                    _ = v24
                default:
                    switch ("hello" + ((v5 + v9) + ((v5 + (("\n\t" + ("hello" + "hello")) + v5)) + "foo bar"))) {
                    case "\n\t":
                        for fi25 := 0; fi25 < 4; fi25++ {
                            var v26 uint = 255
                            var v27 uint16 = (v1 + 712)
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                    case "foo bar":
                        v1 = uint16(v2)
                    case "a":
                        fmt.Println(v5)
                    }
                }
            default:
                v28 := GenS1{F0: int16(v6), F1: -9223372036854775808}
                _ = v28
            }
            _ = fi14
            _ = v15
        }
    case -9223372036854775808:
        fmt.Println(v12)
    }
    v9 = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
