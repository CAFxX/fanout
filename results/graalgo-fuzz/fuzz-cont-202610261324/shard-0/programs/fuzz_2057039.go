package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int16; F1 uint64; F2 bool }
func genFn0() string {
    if true {
        switch ("" + "\n\t") {
        case "":
            var v4 int64 = -1
            _ = v4
        default:
            var v5 uint32 = 2147483648
            _ = v5
        }
        var v6 uint32 = 968
        var v7 int16 = 42
        _ = v6
        _ = v7
    }
    var v8 uint32 = 7
    for fi9 := 0; fi9 < 2; fi9++ {
        v8 -= (v8 << ((75) & 63))
        if false {
            v8 = v8
            var v10 uint64 = 291
            _ = v10
        } else {
            if ("hello" != "\n\t") {
                v11 := GenS1{F0: (("\n\t" + "foo bar") + "foo bar"), F1: (fi9 - fi9)}
                _ = v11
            } else {
                for fi12 := 0; fi12 < 13; fi12++ {
                    v13 := GenS2{F0: -1, F1: 1, F2: false}
                    if (56 >= 42) {
                        v14 := []int{fi12}
                        var v15 float32 = 0.0
                        switch fi12 {
                        case -1:
                            if (fi9 == fi9) {
                                fmt.Println(v8, v8)
                            } else {
                                switch "a" {
                                case "hello":
                                    fmt.Println(fi12, fi9)
                                }
                                v16 := GenS1{F0: ((("hello" + ("hello" + ("foo bar" + ("hello" + (("foo bar" + (((("" + "hello") + "") + ("a" + ((("" + "\n\t") + "") + "\n\t"))) + ((("a" + (("" + ("foo bar" + "\n\t")) + ((((("foo bar" + ("\n\t" + "")) + ((("" + ("\n\t" + ("\n\t" + "a"))) + ("foo bar" + ("a" + "hello"))) + "")) + "foo bar") + "a") + ("\n\t" + "\n\t")))) + ("\n\t" + ("hello" + ("a" + ("a" + "a"))))) + "a"))) + "a"))))) + "a") + ""), F1: (^(fi12))}
                                v15 -= (((-2.25 / 1e10) + (v15 * (v15 - v15))) + v15)
                                _ = v16
                            }
                        default:
                            switch "\n\t" {
                            case "":
                                fmt.Println(v15)
                            case "\n\t":
                                var v17 int64 = -617
                                _ = v17
                            case "foo bar":
                                var v18 int64 = -278
                                _ = v18
                            }
                        }
                        _ = v14
                        _ = v15
                    }
                    v8 += v8
                    _ = fi12
                    _ = v13
                }
            }
            var v19 bool = ((!(false || (true && (-1 == 100)))) || (!false))
            _ = v19
        }
        _ = fi9
    }
    _ = v8
    return "foo bar"
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]bool)
    var v3 string = "\n\t"
    var v20 float32 = 3.14159
    var v21 uint32 = 0
    switch v3 {
    case "hello":
        v20 += 1.5
    case "\n\t":
        v22 := GenS1{F0: "", F1: 534}
        _ = v22
    case "":
        var v23 uint16 = 113
        _ = v23
    default:
        var v24 int8 = 39
        _ = v24
    }
    v25 := GenS2{F0: 1, F1: 1, F2: false}
    var v26 float32 = float32((float64(((v20 - (v20 * (v20 * (v20 * 3.14159)))) * (v20 / float32(float64(1e10))))) + 1.5))
    for fi27 := 0; fi27 < 7; fi27++ {
        v28 := GenS1{F0: ("\n\t" + "hello"), F1: fi27}
        _ = fi27
        _ = v28
    }
    for fi29 := 0; fi29 < 16; fi29++ {
        v26 = -2.25
        fmt.Println(fi29, v26, v21)
        _ = fi29
    }
    if false {
        switch 0 {
        case 32767:
            for fi30 := 0; fi30 < 16; fi30++ {
                v21 -= (v21 << ((2) & 63))
                v3 = (v3 + (v3 + v3))
                v21 = 2147483647
                _ = fi30
            }
        case 0:
            for fi31 := 0; fi31 < 20; fi31++ {
                var v32 bool = true
                _ = fi31
                _ = v32
            }
        }
    }
    var v33 int8 = 42
    fmt.Println(v3, v20)
    fmt.Println(v26, v21)
    v34 := GenS1{F0: v3, F1: 32768}
    fmt.Println(v33, v21, v20)
    v26 = (v20 + v26)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v33
    _ = v34
}
