package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 string; F2 int }
func genFn0() bool {
    var v7 uint64 = 4294967295
    v8 := GenS1{F0: "", F1: v7, F2: "\n\t"}
    for fi9 := 0; fi9 < 16; fi9++ {
        if ((true && (!true)) && (!true)) {
            if (!false) {
                switch "" {
                case "foo bar":
                    fmt.Println(v7, v7)
                case "a":
                    switch "\n\t" {
                    case "a":
                        v10 := []int{fi9, -910, (fi9 & fi9)}
                        _ = v10
                    case "\n\t":
                        v11 := GenS2{F0: "\n\t"}
                        _ = v11
                    case "hello":
                        v12 := []int{127, (-(fi9)), (fi9 % ((fi9) | 1))}
                        _ = v12
                    }
                default:
                    v13 := []int{-9223372036854775808, (fi9 ^ (fi9 * -891))}
                    _ = v13
                }
                if (2147483647 == 65535) {
                    var v14 uint8 = 255
                    _ = v14
                }
            } else {
                for fi15 := 0; fi15 < 13; fi15++ {
                    fmt.Println(v7, fi15, fi15)
                    v16 := []int{fi15, -783, fi15, fi15}
                    _ = fi15
                    _ = v16
                }
                var v17 uint = 227
                _ = v17
            }
            v7 += (v7 & (v7 + v7))
            var v18 float64 = -2.25
            _ = v18
        }
        var v19 int8 = -35
        _ = fi9
        _ = v19
    }
    _ = v7
    _ = v8
    return false
}

func main() {
    v1 := []string{(("a" + "\n\t") + "")}
    v2 := make(map[int]bool)
    v3 := []uint32{48}
    var v4 uint = 42
    v5 := []string{(("a" + "") + "hello"), "a"}
    var v6 uint64 = 777
    fmt.Println(v6, v6)
    for fi20 := 0; fi20 < 20; fi20++ {
        for fi21 := 0; fi21 < 11; fi21++ {
            switch "a" {
            case "foo bar":
                switch "foo bar" {
                case "\n\t":
                    v2[(fi20 & (fi20 ^ fi20))] = false
                case "a":
                    var v22 uint32 = 42
                    _ = v22
                case "":
                    if ((true || (99 == 42)) && false) {
                        v6 = 0
                        v4 -= v4
                        switch fi21 {
                        case 0:
                            var v23 int32 = 495
                            _ = v23
                        default:
                            v6 -= 2
                        }
                    } else {
                        var v24 uint8 = 2
                        v24 = uint8(v6)
                        var v25 float32 = (3.14159 - -2.25)
                        v25 += (v25 - (3.14159 - v25))
                        var v26 float32 = (0.1 * float32((0.0 + 0.1)))
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                default:
                    var v27 float32 = ((3.14159 + -2.25) - 0.0)
                    _ = v27
                }
            }
            _ = fi21
        }
        _ = fi20
    }
    v2[128] = (!((v6 - (v6 * (^(v6)))) > (v6 | (v6 | v6))))
    switch "\n\t" {
    case "a":
        for fi28 := 0; fi28 < 14; fi28++ {
            switch fi28 {
            case 128:
                switch ("hello" + "") {
                case "hello":
                    if true {
                        v6 += v6
                        if false {
                            fmt.Println(fi28, fi28)
                            v29 := GenS1{F0: ("" + ("a" + "hello")), F1: (v6 + v6), F2: "hello"}
                            _ = v29
                        }
                        var v30 int8 = -1
                        _ = v30
                    } else {
                        v6 -= (v6 & 18446744073709551615)
                        v6 += (^(v6))
                    }
                default:
                    delete(v2, (fi28 >> ((404) & 63)))
                }
            }
            _ = fi28
        }
    case "\n\t":
        var v31 uint16 = 32767
        _ = v31
    case "":
        v32 := GenS3{F0: false}
        _ = v32
    }
    var v33 uint16 = 1
    var v34 int = 984
    var v35 float64 = -2.25
    v36 := GenS4{F0: 151, F1: ("a" + "\n\t"), F2: 0}
    var v37 uint64 = v6
    _ = len(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
