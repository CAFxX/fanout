package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int16; F2 uint64 }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 int64; F1 int32; F2 uint64 }
type GenS5 struct { F0 int; F1 uint8 }
func genFn0() string {
    v6 := []int{9223372036854775807, 4294967295, 42, 42}
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 uint = 859
        _ = fi7
        _ = v8
    }
    _ = v6
    return ("\n\t" + (((("" + (((((((("\n\t" + "") + ((("" + ("foo bar" + "\n\t")) + "a") + "a")) + "foo bar") + "foo bar") + ("\n\t" + ("foo bar" + ("hello" + "foo bar")))) + "\n\t") + "hello") + ("a" + "a"))) + "") + ("a" + "")) + "\n\t"))
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = -1
    v3 := []uint32{}
    v4 := make(map[uint32]bool)
    var v5 float64 = 0.0
    v9 := GenS1{F0: 695}
    v5 += v5
    if (!(true && true)) {
        fmt.Println(v5, v2)
        for fi10 := 0; fi10 < 11; fi10++ {
            v5 -= (float64(float32(v5)) + (v5 - 3.14159))
            if true {
                fmt.Println(v2)
                for fi11 := 0; fi11 < 19; fi11++ {
                    switch "hello" {
                    case "foo bar":
                        v5 = float64(3.14159)
                    case "":
                        fmt.Println(v5, v2, v5)
                    case "\n\t":
                        var v12 bool = true
                        _ = v12
                    default:
                        if (127 < 42) {
                            fmt.Println(fi11)
                            v13 := GenS2{F0: 255, F1: -32768, F2: 0}
                            _ = v13
                        }
                    }
                    switch fi11 {
                    case 2147483647:
                        var v14 float64 = v5
                        _ = v14
                    case -9223372036854775808:
                        switch fi10 {
                        case 633:
                            v5 += 0.0
                            v5 = v5
                            _ = len(v3)
                        case 699:
                            var v15 string = ""
                            _ = v15
                        case -1:
                            var v16 float64 = 3.14159
                            _ = v16
                        }
                    }
                    v17 := GenS3{F0: false, F1: -1}
                    if (0 == 42) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    } else {
                        v2 = v2
                        if (false && false) {
                            v5 = v5
                            switch fi10 {
                            case 32767:
                                v18 := GenS4{F0: -1, F1: 2147483647, F2: 4294967295}
                                _ = v18
                            case 340:
                                v19 := GenS5{F0: 797, F1: 85}
                                _ = v19
                            default:
                                switch fi11 {
                                case 127:
                                    var v20 uint64 = 151
                                    _ = v20
                                default:
                                    v5 += v5
                                }
                            }
                            v2 = v2
                        }
                        fmt.Println(v2)
                        var v21 uint64 = 618
                        _ = v21
                    }
                    v22 := GenS2{F0: 32767, F1: 1, F2: 619}
                    _ = len(v3)
                    _ = fi11
                    _ = v17
                    _ = v22
                }
            }
            v5 -= v5
            _ = fi10
        }
    } else {
        v3 = append(v3, 65535)
    }
    fmt.Println(v2, v2)
    _ = len(v3)
    if ("hello" == "foo bar") {
        if (true && false) {
            var v23 bool = false
            var v24 uint16 = 32767
            v3 = append(v3, 1)
            _ = v23
            _ = v24
        }
    }
    for fi25 := 0; fi25 < 9; fi25++ {
        switch "" {
        case "a":
            switch 375 {
            case 9223372036854775807:
                for fi26 := 0; fi26 < 17; fi26++ {
                    v27 := GenS3{F0: false, F1: -32768}
                    _ = fi26
                    _ = v27
                }
            case 0:
                v2 = v2
            }
        }
        _ = fi25
    }
    v5 = v5
    _ = len(v3)
    delete(v4, 713)
    var v28 bool = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
