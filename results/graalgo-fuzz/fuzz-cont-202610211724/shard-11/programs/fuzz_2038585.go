package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 int32; F2 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32; F1 uint64 }
func genFn0() int {
    var v7 uint64 = 979
    var v8 uint8 = 1
    v8 += v8
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 int = -2147483648
    v2 := make(map[string]string)
    var v3 uint8 = 1
    var v4 float64 = ((0.1 + 3.14159) * 1.5)
    var v5 int8 = int8(v3)
    v6 := make(map[string]uint32)
    delete(v2, ((((("foo bar" + "foo bar") + (("\n\t" + ("\n\t" + "hello")) + "a")) + "") + "hello") + ""))
    v4 = (v4 * v4)
    for fi9 := 0; fi9 < 17; fi9++ {
        for fi10 := 0; fi10 < 13; fi10++ {
            v11 := GenS1{F0: -618}
            v5 -= (v5 | (v5 + v5))
            v3 -= 0
            _ = fi10
            _ = v11
        }
        v5 = v5
        _ = fi9
    }
    v4 = float64(float32((0.1 - (v4 / v4))))
    if true {
        v1 = (-(v1))
        _ = len(v6)
    }
    v3 = (v3 << ((1) & 63))
    if false {
        v4 -= (-2.25 + -2.25)
        for fi12 := 0; fi12 < 8; fi12++ {
            v3 = (v3 / ((v3) | 1))
            _ = fi12
        }
        fmt.Println(v4)
    } else {
        var v13 uint16 = 32767
        _ = v13
    }
    fmt.Println(v5, v4)
    for fi14 := 0; fi14 < 14; fi14++ {
        for fi15 := 0; fi15 < 2; fi15++ {
            v16 := GenS2{F0: 42, F1: 42, F2: 0}
            var v17 int = (v1 >> ((42) & 63))
            _ = fi15
            _ = v16
            _ = v17
        }
        v1 = fi14
        v1 += v1
        _ = fi14
    }
    var v18 int8 = (v5 ^ -71)
    v1 = (v1 - v1)
    switch "" {
    case "a":
        switch (v1 << ((0) & 63)) {
        case -2147483648:
            v18 += (v5 << ((255) & 63))
        case 443:
            if true {
                var v19 int8 = (v18 | (v18 | (v5 * 17)))
                for fi20 := 0; fi20 < 18; fi20++ {
                    switch "a" {
                    case "\n\t":
                        v21 := GenS3{F0: 0}
                        _ = v21
                    case "":
                        fmt.Println(v3)
                    case "a":
                        for fi22 := 0; fi22 < 19; fi22++ {
                            _ = len(v2)
                            _ = fi22
                        }
                    default:
                        if (0 == 42) {
                            fmt.Println(v18, v19, v19)
                        }
                    }
                    if (!false) {
                        fmt.Println(v19, fi20, v3)
                        v23 := GenS4{F0: 42, F1: 962}
                        v24 := GenS4{F0: -1, F1: 538}
                        _ = v23
                        _ = v24
                    }
                    var v25 int32 = 0
                    _ = fi20
                    _ = v25
                }
                _ = v19
            }
        default:
            var v26 int16 = 253
            _ = v26
        }
    case "hello":
        if ((-447 != 400) || false) {
            v27 := []int{v1}
            switch "hello" {
            case "hello":
                switch "hello" {
                case "a":
                    if len(v27) >= 2 {
                        v27 = v27[1:len(v27)]
                    }
                default:
                    switch "a" {
                    case "foo bar":
                        var v28 float64 = 1.5
                        _ = v28
                    case "\n\t":
                        v3 -= v3
                    }
                }
            case "\n\t":
                var v29 string = ("a" + "foo bar")
                _ = v29
            case "":
                _ = len(v27)
            }
            _ = v27
        } else {
            switch ((("hello" + ((("\n\t" + ("" + "")) + "") + "hello")) + "hello") + "a") {
            case "foo bar":
                for fi30 := 0; fi30 < 10; fi30++ {
                    for fi31 := 0; fi31 < 7; fi31++ {
                        var v32 int64 = 9223372036854775807
                        var v33 int32 = 1
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                    v3 += v3
                    delete(v2, "foo bar")
                    _ = fi30
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
