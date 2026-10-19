package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 int16; F1 int16; F2 uint8 }
type GenS4 struct { F0 int32; F1 uint64 }
func genFn0(p0 int32, p1 string, p2 uint) int64 {
    v7 := GenS1{F0: 1, F1: p2}
    v8 := []int{-2147483648, 1, 4294967295, 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -9223372036854775808
}
func genFn1(p0 uint64) uint16 {
    var v9 string = "foo bar"
    v10 := []int{307}
    _ = p0
    _ = v10
    _ = v9
    return 32767
}

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 284
    v3 := []int{0, -9223372036854775808, 128}
    var v4 int8 = 7
    v5 := make(map[string]int)
    var v6 int64 = 0
    v2 = (v2 - v2)
    v3 = append(v3, 4294967295)
    fmt.Println(v4, v4)
    switch "a" {
    case "":
        switch ("\n\t" + "hello") {
        case "a":
            v4 += 42
        }
    case "a":
        v6 -= v6
    case "\n\t":
        switch 0 {
        case 625:
            switch 9223372036854775807 {
            case 32767:
                v11 := GenS1{F0: -640, F1: 65535}
                _ = v11
            default:
                fmt.Println(v4)
            }
        case -83:
            fmt.Println(v4, v4, v4)
        }
    }
    switch "\n\t" {
    case "a":
        fmt.Println(v2, v4, v4)
    case "hello":
        for fi12 := 0; fi12 < 1; fi12++ {
            var v13 int64 = v6
            _ = fi12
            _ = v13
        }
    case "foo bar":
        fmt.Println(v4)
    default:
        var v14 uint16 = 32767
        _ = v14
    }
    var v15 float64 = ((-2.25 - 1e10) / 0.1)
    var v16 bool = false
    fmt.Println(v6, v15)
    switch -1 {
    case 846:
        delete(v1, "")
    case -74:
        for fi17 := 0; fi17 < 18; fi17++ {
            v15 = v15
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if ((v6 | -1) >= genFn0(2147483647, "a", 255)) {
                if (0 < 42) {
                    v4 -= 127
                }
                for fi18 := 0; fi18 < 16; fi18++ {
                    switch (fi18 + fi18) {
                    case 32768:
                        if (float64(float32(((v15 - 1e10) - float64(0.0)))) > v15) {
                            v19 := GenS2{F0: ("hello" + "foo bar"), F1: v2}
                            if true {
                                fmt.Println(v2)
                                var v20 uint16 = 32767
                                _ = v20
                            } else {
                                fmt.Println(fi17)
                                delete(v5, (("" + ("hello" + "")) + ("" + (("" + "\n\t") + "hello"))))
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v21 := GenS2{F0: "hello", F1: (v2 << ((65535) & 63))}
                                switch fi17 {
                                case 128:
                                    v15 = v15
                                case 752:
                                    v16 = (206 >= 27)
                                default:
                                    v22 := GenS3{F0: -439, F1: 1, F2: 255}
                                    _ = v22
                                }
                                var v23 float32 = 0.1
                                v24 := GenS4{F0: -155, F1: 24}
                                _ = v21
                                _ = v23
                                _ = v24
                            }
                            v16 = (111 < -441)
                            _ = v19
                        } else {
                            var v25 int = 2147483647
                            var v26 string = "a"
                            fmt.Println(v15, v4, v16)
                            var v27 uint = 1
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                    case 1:
                        v6 = v6
                    default:
                        var v28 string = "hello"
                        _ = v28
                    }
                    v6 -= -416
                    _ = fi18
                }
            }
            _ = fi17
        }
    default:
        v4 += (^(v4))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
