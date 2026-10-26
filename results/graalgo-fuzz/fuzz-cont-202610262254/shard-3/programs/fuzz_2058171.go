package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 int }
func genFn0(p0 uint, p1 string, p2 bool) bool {
    p2 = true
    var v5 int8 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p2
}
func genFn1(p0 bool) int {
    v6 := GenS1{F0: 42}
    p0 = p0
    v7 := GenS1{F0: 826}
    switch 2147483647 {
    case -768:
        p0 = (-1 == 86)
    }
    _ = p0
    _ = v6
    _ = v7
    return 32768
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]bool)
    var v3 uint8 = 167
    var v4 int32 = int32(v3)
    fmt.Println(v3)
    switch "hello" {
    case "\n\t":
        var v8 uint64 = 890
        _ = v8
    case "hello":
        fmt.Println(v4, v4)
    default:
        v3 -= 128
    }
    fmt.Println(v3, v3)
    v3 = 255
    _ = len(v2)
    if ((v3 / ((v3) | 1)) <= (v3 + 127)) {
        var v9 int32 = (v4 / (((v4 - -1)) | 1))
        var v10 int = -2147483648
        v11 := []int{v10, (v10 | -216)}
        _ = v10
        _ = v11
        _ = v9
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        v4 -= (^(v4))
        v13 := GenS2{F0: (0.0 >= float32(((-2.25 / 1.5) - 3.14159))), F1: fi12}
        for fi14 := 0; fi14 < 17; fi14++ {
            v3 = (v3 % (((^(v3))) | 1))
            v15 := GenS2{F0: (!(true || (true && (((0.1 != -2.25) && (true || (-2.25 < (0.1 * -2.25)))) && (!true))))), F1: fi12}
            v4 = (v4 ^ v4)
            _ = fi14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    v3 = 42
    v3 = (^(v3))
    if true {
        if true {
            v4 -= v4
            _ = len(v2)
        }
    }
    switch ("foo bar" + "\n\t") {
    case "foo bar":
        var v16 uint16 = 1
        _ = v16
    case "":
        switch "hello" {
        case "foo bar":
            switch ("a" + "\n\t") {
            case "a":
                fmt.Println(v3, v3)
            case "":
                switch 1 {
                case 32768:
                    for fi17 := 0; fi17 < 1; fi17++ {
                        v3 = v3
                        var v18 bool = true
                        _ = fi17
                        _ = v18
                    }
                default:
                    for fi19 := 0; fi19 < 12; fi19++ {
                        v3 -= (v3 & v3)
                        for fi20 := 0; fi20 < 3; fi20++ {
                            var v21 int64 = -1
                            switch fi20 {
                            case -985:
                                delete(v1, fi20)
                            case 607:
                                var v22 float32 = float32((float64(3.14159) * 1.5))
                                _ = v22
                            }
                            v23 := []int{(fi19 | fi20), fi20}
                            _ = fi20
                            _ = v21
                            _ = v23
                        }
                        _ = fi19
                    }
                }
            }
        default:
            var v24 uint32 = 65535
            _ = v24
        }
    }
    for fi25 := 0; fi25 < 1; fi25++ {
        if ((fi25 << ((2) & 63)) > fi25) {
            for fi26 := 0; fi26 < 1; fi26++ {
                fmt.Println(fi26)
                _ = fi26
            }
            fmt.Println(v3, v3)
            var v27 float32 = -2.25
            _ = v27
        } else {
            v28 := []int{}
            var v29 string = "hello"
            _ = v28
            _ = v29
        }
        fmt.Println(v3)
        var v30 uint32 = 2147483647
        _ = fi25
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
