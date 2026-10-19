package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 uint16; F1 uint; F2 string }
func genFn0(p0 int32, p1 uint16, p2 int16) int8 {
    for fi5 := 0; fi5 < 16; fi5++ {
        v6 := GenS1{F0: 467, F1: (p2 - p2)}
        p0 += p0
        if (42 == 255) {
            var v7 uint8 = 1
            _ = v7
        } else {
            if true {
                switch ("" + ((((("foo bar" + (((("foo bar" + "a") + "hello") + "foo bar") + "hello")) + "foo bar") + "") + ("hello" + ("a" + "\n\t"))) + "")) {
                case "foo bar":
                    if true {
                        p0 -= p0
                    }
                }
            } else {
                p0 += p0
            }
        }
        _ = fi5
        _ = v6
    }
    v8 := []int{-258, 292}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 1
}
func genFn1(p0 uint16, p1 uint32) int8 {
    if true {
        p1 -= (p1 - p1)
        p1 = p1
    }
    var v9 uint = 839
    switch 686 {
    case -1:
        var v10 uint64 = uint64(p0)
        _ = v10
    case 32767:
        switch 9223372036854775807 {
        case 472:
            for fi11 := 0; fi11 < 1; fi11++ {
                for fi12 := 0; fi12 < 7; fi12++ {
                    if true {
                        if (!((v9 & 4294967295) >= (v9 + v9))) {
                            var v13 float64 = (0.1 - 1e10)
                            var v14 uint8 = 96
                            _ = v13
                            _ = v14
                        } else {
                            v9 = (^(v9))
                            p0 = p0
                            p0 -= 0
                            v15 := GenS1{F0: -820, F1: -448}
                            _ = v15
                        }
                        v16 := GenS1{F0: 0, F1: 32767}
                        fmt.Println(p1, p1)
                        _ = v16
                    } else {
                        var v17 string = "\n\t"
                        v18 := GenS1{F0: 9223372036854775807, F1: 1000}
                        _ = v17
                        _ = v18
                    }
                    _ = fi12
                }
                for fi19 := 0; fi19 < 7; fi19++ {
                    v20 := GenS1{F0: 1, F1: -32768}
                    p1 -= (p1 / (((p1 & 0)) | 1))
                    _ = fi19
                    _ = v20
                }
                p0 = 2
                _ = fi11
            }
        default:
            var v21 uint64 = 42
            _ = v21
        }
    default:
        v9 -= (v9 % ((198) | 1))
    }
    fmt.Println(v9)
    v22 := []int{}
    _ = p0
    _ = p1
    _ = v22
    _ = v9
    return -128
}

func main() {
    var v1 float64 = 1.5
    var v2 string = "hello"
    v3 := make(map[int]string)
    v4 := make(map[int]bool)
    delete(v4, 128)
    if false {
        v1 = -2.25
        v23 := GenS1{F0: 42, F1: 42}
        switch v2 {
        case "\n\t":
            var v24 int16 = 32767
            _ = v24
        case "hello":
            if (!(!false)) {
                var v25 uint32 = 725
                fmt.Println(v25)
                _ = v25
            } else {
                fmt.Println(v2)
                delete(v3, -678)
            }
        case "":
            v2 = ""
        default:
            v26 := GenS1{F0: -967, F1: 42}
            _ = v26
        }
        _ = v23
    }
    v27 := v4[-9223372036854775808]
    for fi28 := 0; fi28 < 14; fi28++ {
        v29 := GenS2{F0: 65535, F1: 801, F2: v2}
        v30 := []int{(fi28 % ((42) | 1)), (fi28 + 127), -252}
        for fi31 := 0; fi31 < 9; fi31++ {
            fmt.Println(fi28, fi28)
            _ = fi31
        }
        _ = fi28
        _ = v29
        _ = v30
    }
    switch 32767 {
    case 9223372036854775807:
        v32 := []int{}
        _ = v32
    case -9223372036854775808:
        v33 := []int{-308, -563, 4294967295, 730}
        _ = v33
    }
    v27 = true
    var v34 int32 = -221
    var v35 uint16 = 32767
    v34 += (v34 >> ((822) & 63))
    v36 := GenS1{F0: -9223372036854775808, F1: int16(v34)}
    _ = len(v4)
    var v37 uint32 = 42
    var v38 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
