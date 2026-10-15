package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 bool; F2 uint }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 int64; F1 int32; F2 int32 }
type GenS6 struct { F0 int; F1 int32 }
type GenS7 struct { F0 int8 }
func genFn0() uint8 {
    v5 := []int{42}
    var v6 int = 42
    var v7 float64 = 0.1
    var v8 int32 = -2147483648
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 198
}
func genFn1(p0 string) uint {
    v9 := GenS1{F0: 0}
    v10 := GenS2{F0: ((true || true) || (1 == -128)), F1: true, F2: 2}
    _ = p0
    _ = v10
    _ = v9
    return 65535
}

func main() {
    v1 := []string{"foo bar", "a", "a", "\n\t"}
    var v2 uint64 = 42
    v3 := make(map[uint32]bool)
    var v4 float64 = float64(0.1)
    for fi11 := 0; fi11 < 10; fi11++ {
        v12 := v3[365]
        v13 := GenS3{F0: (fi11 << ((0) & 63)), F1: 388}
        _ = fi11
        _ = v12
        _ = v13
    }
    if true {
        for fi14 := 0; fi14 < 1; fi14++ {
            for fi15 := 0; fi15 < 11; fi15++ {
                if (!(!true)) {
                    v4 = (v4 - v4)
                    switch fi15 {
                    case -344:
                        var v16 int64 = 173
                        _ = v16
                    case -2147483648:
                        v17 := GenS1{F0: 0}
                        switch "hello" {
                        case "a":
                            var v18 int = (^(fi15))
                            var v19 int64 = -9223372036854775808
                            _ = v18
                            _ = v19
                        default:
                            v3[65535] = (2147483647 < -2147483648)
                        }
                        _ = v17
                    case 32768:
                        v20 := GenS4{F0: 2, F1: 1}
                        _ = v20
                    default:
                        var v21 int16 = 332
                        _ = v21
                    }
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v2 = v2
                _ = fi15
            }
            switch fi14 {
            case 32767:
                if false {
                    v22 := GenS2{F0: false, F1: true, F2: genFn1("")}
                    _ = v22
                }
            }
            v4 += -2.25
            _ = fi14
        }
    } else {
        var v23 uint32 = 65535
        fmt.Println(v2)
        _ = v23
    }
    if len(v1) > 0 {
        v1[0] = ("foo bar" + ("\n\t" + ((("" + "a") + "a") + "\n\t")))
    }
    if len(v1) > 0 {
        v1[0] = ("foo bar" + "hello")
    }
    switch "" {
    case "hello":
        fmt.Println(v4, v2)
    case "foo bar":
        fmt.Println(v4, v4)
    case "\n\t":
        if false {
            for fi24 := 0; fi24 < 18; fi24++ {
                for fi25 := 0; fi25 < 15; fi25++ {
                    var v26 int32 = 2147483647
                    v27 := GenS5{F0: 812, F1: int32(v2), F2: v26}
                    var v28 int16 = -1
                    v2 = v2
                    _ = fi25
                    _ = v26
                    _ = v27
                    _ = v28
                }
                _ = fi24
            }
            var v29 string = ("foo bar" + "hello")
            _ = v29
        }
    }
    switch 126 {
    case 203:
        v2 = v2
    default:
        v30 := GenS6{F0: 625, F1: -748}
        _ = v30
    }
    for fi31 := 0; fi31 < 3; fi31++ {
        v32 := GenS7{F0: 42}
        fmt.Println(fi31)
        _ = fi31
        _ = v32
    }
    delete(v3, 42)
    switch 9223372036854775807 {
    case -1:
        for fi33 := 0; fi33 < 17; fi33++ {
            for fi34 := 0; fi34 < 8; fi34++ {
                fmt.Println(fi34, v4, v2)
                _ = fi34
            }
            v35 := GenS7{F0: 100}
            _ = fi33
            _ = v35
        }
    case 2147483647:
        v1 = append(v1, "hello")
    default:
        switch "a" {
        case "":
            switch ("\n\t" + "\n\t") {
            case "\n\t":
                _ = len(v3)
            case "a":
                var v36 uint32 = 2147483647
                _ = v36
            default:
                fmt.Println(v2)
            }
        default:
            v2 += v2
        }
    }
    var v37 float64 = (v4 - (v4 / 0.1))
    fmt.Println(v4, v4, v37)
    var v38 float64 = v37
    v37 += float64((3.14159 + (((-2.25 * (1.5 * (0.1 - (1.5 - float32(v38))))) / (0.0 + 3.14159)) + -2.25)))
    var v39 bool = true
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
