package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int32 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int64; F1 uint64 }
func genFn0(p0 int64) uint16 {
    if false {
        var v7 float32 = (0.1 - -2.25)
        var v8 uint16 = 0
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 2; fi9++ {
            p0 = (p0 & (^(p0)))
            _ = fi9
        }
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        if false {
            p0 -= 0
            p0 -= p0
            if (18446744073709551615 < 42) {
                v11 := GenS1{F0: "foo bar", F1: 18446744073709551615, F2: 42}
                v12 := GenS2{F0: (fi10 ^ fi10), F1: ("hello" + "a")}
                for fi13 := 0; fi13 < 15; fi13++ {
                    v14 := GenS2{F0: 127, F1: ((("" + (("hello" + "a") + "a")) + ("hello" + "")) + "")}
                    var v15 string = ("foo bar" + "a")
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                _ = v11
                _ = v12
            }
        } else {
            var v16 uint16 = 32767
            v16 -= 0
            p0 = (p0 % (((-(p0))) | 1))
            _ = v16
        }
        v17 := GenS1{F0: "\n\t", F1: 18446744073709551615, F2: 599}
        _ = fi10
        _ = v17
    }
    if true {
        v18 := []int{-334, 2147483647}
        _ = v18
    } else {
        v19 := []int{-592}
        var v20 int64 = (p0 - -566)
        for fi21 := 0; fi21 < 19; fi21++ {
            if false {
                for fi22 := 0; fi22 < 20; fi22++ {
                    switch 32767 {
                    case 6:
                        v23 := GenS3{F0: 30}
                        _ = v23
                    }
                    p0 = (v20 - 0)
                    _ = fi22
                }
                for fi24 := 0; fi24 < 20; fi24++ {
                    fmt.Println(fi21)
                    p0 = v20
                    _ = fi24
                }
            }
            _ = fi21
        }
        _ = v19
        _ = v20
    }
    v25 := GenS1{F0: "foo bar", F1: 723, F2: 270}
    _ = p0
    _ = v25
    return 675
}

func main() {
    var v1 string = ("foo bar" + "")
    var v2 bool = false
    var v3 bool = (2147483647 != 0)
    var v4 int64 = -1
    var v5 bool = (2 == 705)
    v6 := make(map[uint32]int)
    v26 := v6[1]
    v27 := GenS3{F0: 255}
    v2 = v2
    switch (v26 | (v26 ^ 4294967295)) {
    case 4294967295:
        fmt.Println(v1, v1)
    case 42:
        v28 := []int{v26}
        _ = v28
    case -1:
        var v29 float64 = (1e10 * 0.1)
        _ = v29
    default:
        v30 := GenS4{F0: -536}
        _ = v30
    }
    var v31 int16 = 32767
    v32 := []int{(v26 + (v26 >> ((255) & 63)))}
    v6[509] = 2147483647
    var v33 int32 = 597
    fmt.Println(v33, v1)
    v34 := GenS5{F0: (!false)}
    v2 = ((false && (!true)) && v3)
    switch v1 {
    case "":
        switch v1 {
        case "foo bar":
            v35 := GenS6{F0: (v4 >> ((287) & 63)), F1: 42}
            _ = v35
        case "":
            fmt.Println(v26, v31)
        }
    case "hello":
        v36 := GenS2{F0: -740, F1: (v1 + (("" + v1) + v1))}
        _ = v36
    }
    var v37 int = v26
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
