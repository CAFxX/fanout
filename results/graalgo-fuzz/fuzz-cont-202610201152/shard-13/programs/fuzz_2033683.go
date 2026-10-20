package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint8 }
func genFn0() int16 {
    v6 := GenS1{F0: "", F1: -1, F2: 26}
    var v7 int = 0
    for fi8 := 0; fi8 < 1; fi8++ {
        if (((-(v7)) == (v7 | (-(fi8)))) && false) {
            var v9 float32 = 3.14159
            var v10 int32 = -478
            _ = v10
            _ = v9
        }
        _ = fi8
    }
    v11 := []int{v7, 648, 598}
    var v12 int32 = 0
    _ = v11
    _ = v12
    _ = v6
    _ = v7
    return -32768
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]int)
    v3 := make(map[int]int)
    v4 := make(map[string]bool)
    var v5 int64 = 42
    var v13 int64 = v5
    var v14 uint16 = 1
    for fi15 := 0; fi15 < 2; fi15++ {
        v16 := []int{(fi15 ^ fi15), 1, (fi15 + -458)}
        for fi17 := 0; fi17 < 18; fi17++ {
            v18 := v4["foo bar"]
            _ = fi17
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    v19 := []int{-1, 864}
    var v20 float64 = float64(1.5)
    var v21 uint64 = 1
    var v22 uint16 = (v14 & (v14 & v14))
    v23 := GenS1{F0: "foo bar", F1: (v5 & (v5 + v13)), F2: 236}
    switch -196 {
    case -2147483648:
        _ = len(v19)
    }
    fmt.Println(v13)
    v21 += v21
    for fi24 := 0; fi24 < 13; fi24++ {
        fmt.Println(v14)
        if len(v19) > 0 {
            v19[0] = 128
        }
        _ = fi24
    }
    if true {
        var v25 string = ((("" + (((("\n\t" + (("\n\t" + "\n\t") + "foo bar")) + "\n\t") + "hello") + "hello")) + "") + ("" + ("\n\t" + "foo bar")))
        if (!((-2.25 > ((v20 * -2.25) + (v20 + 0.1))) && (false && (!true)))) {
            var v26 uint64 = (v21 - (v21 % (((^(v21))) | 1)))
            if (!(uint(v14) != 1)) {
                v21 = (v26 & (v21 >> ((1) & 63)))
                for fi27 := 0; fi27 < 1; fi27++ {
                    var v28 int8 = 100
                    _ = fi27
                    _ = v28
                }
                var v29 uint64 = v21
                _ = v29
            } else {
                for fi30 := 0; fi30 < 18; fi30++ {
                    for fi31 := 0; fi31 < 5; fi31++ {
                        var v32 float32 = 0.1
                        v2[0] = fi31
                        _ = fi31
                        _ = v32
                    }
                    fmt.Println(v14, fi30)
                    _ = fi30
                }
                v33 := v1[65535]
                v33 += 830
                _ = v33
            }
            _ = v26
        }
        _ = v25
    } else {
        for fi34 := 0; fi34 < 10; fi34++ {
            var v35 float32 = 0.1
            var v36 uint8 = 236
            _ = fi34
            _ = v35
            _ = v36
        }
        switch 127 {
        case -507:
            var v37 int64 = (^(v5))
            _ = v37
        case -2147483648:
            var v38 uint = 4294967295
            _ = v38
        }
    }
    var v39 int16 = genFn0()
    v13 = (v13 * v5)
    switch "hello" {
    case "hello":
        v22 += (^(v14))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v39
    _ = v4
    _ = v5
}
