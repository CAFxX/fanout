package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 string; F2 int32 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 int; F1 uint; F2 uint32 }
func genFn0(p0 int, p1 int8) int32 {
    if false {
        for fi7 := 0; fi7 < 10; fi7++ {
            p0 = int(p1)
            _ = fi7
        }
        v8 := GenS1{F0: 0}
        _ = v8
    }
    var v9 int64 = -155
    for fi10 := 0; fi10 < 2; fi10++ {
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 string = ("hello" + "\n\t")
            if (((float32(0.0) / 1.0) == 0.0) || ((p1 & (p1 + p1)) > 123)) {
                v13 := []int{}
                v12 = ""
                v14 := GenS1{F0: 800}
                _ = v13
                _ = v14
            } else {
                v15 := GenS2{F0: -32768, F1: "\n\t", F2: 2147483647}
                fmt.Println(v12, v9)
                v16 := []int{}
                _ = v15
                _ = v16
            }
            _ = fi11
            _ = v12
        }
        v17 := GenS2{F0: 42, F1: (((("foo bar" + "") + "a") + "") + ""), F2: 492}
        _ = fi10
        _ = v17
    }
    v9 -= (v9 & 42)
    var v18 int16 = 1000
    _ = p0
    _ = p1
    _ = v18
    _ = v9
    return -611
}

func main() {
    v1 := make(map[string]int)
    var v2 float32 = (-2.25 - (3.14159 - (-2.25 - -2.25)))
    v3 := make(map[uint32]bool)
    v4 := []int{-2147483648, -1, -296, 9223372036854775807}
    var v5 int32 = 0
    var v6 string = (("a" + ("hello" + "foo bar")) + ("foo bar" + "\n\t"))
    if true {
        if (true || (42 > 42)) {
            if len(v4) > 0 {
                v4[0] = 2147483647
            }
        }
        if true {
            if ((625 >= 65535) || false) {
                v19 := GenS1{F0: (v5 ^ v5)}
                var v20 uint64 = 0
                _ = v19
                _ = v20
            } else {
                v2 -= ((((v2 * 1e10) + 1e10) + v2) + v2)
            }
            v6 = v6
            switch 9223372036854775807 {
            case 32767:
                for fi21 := 0; fi21 < 2; fi21++ {
                    v6 = v6
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v6 = ("\n\t" + "\n\t")
                    _ = fi21
                }
            }
        }
    }
    if true {
        _ = len(v3)
        v22 := GenS3{F0: 4294967295, F1: 32768}
        var v23 int32 = genFn0(32767, 100)
        _ = v22
        _ = v23
    } else {
        v5 = genFn0(2147483647, 127)
        switch 5 {
        case 1:
            v24 := v1[v6]
            _ = v24
        }
    }
    var v25 uint32 = 608
    for fi26 := 0; fi26 < 18; fi26++ {
        v27 := v1[v6]
        var v28 float64 = float64(((float32(-2.25) + ((v2 * 1.5) - (0.0 + ((0.1 / v2) + float32(1.5))))) * v2))
        v6 = v6
        _ = fi26
        _ = v27
        _ = v28
    }
    var v29 uint32 = 2
    fmt.Println(v25, v5)
    for fi30 := 0; fi30 < 10; fi30++ {
        v29 -= v29
        for fi31 := 0; fi31 < 20; fi31++ {
            var v32 int32 = -96
            v33 := GenS4{F0: uint64(v25), F1: 1}
            _ = fi31
            _ = v32
            _ = v33
        }
        var v34 int = (fi30 + fi30)
        _ = fi30
        _ = v34
    }
    switch 676 {
    case 128:
        var v35 uint32 = (v29 | (^(v29)))
        _ = v35
    case 32767:
        v36 := GenS5{F0: 876, F1: 18446744073709551615, F2: v25}
        _ = v36
    case 32768:
        v2 = 3.14159
    }
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
