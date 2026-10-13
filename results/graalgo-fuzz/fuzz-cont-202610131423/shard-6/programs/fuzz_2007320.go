package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int; F1 uint8 }
func genFn0(p0 int64, p1 string) int {
    v7 := []int{-1, -486, 59, 128}
    var v8 int8 = -52
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 int32, p1 uint16, p2 int32) int16 {
    v9 := GenS1{F0: -32768}
    p2 = p2
    for fi10 := 0; fi10 < 19; fi10++ {
        fmt.Println(p0)
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 0
}

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 42
    var v3 string = ("" + "\n\t")
    var v4 int = -341
    var v5 uint64 = 4294967295
    var v6 uint64 = v5
    fmt.Println(v2)
    if (v4 > (v4 ^ v4)) {
        v4 = v4
        v5 = v5
    }
    var v11 uint64 = (v6 / (((v6 >> ((2) & 63))) | 1))
    v12 := []int{(^(v4)), (-(v4)), v4}
    if false {
        switch v3 {
        case "":
            v11 = (v11 >> ((42) & 63))
        }
    } else {
        v13 := GenS1{F0: 0}
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        _ = v13
    }
    v14 := GenS2{F0: 0, F1: 1}
    if (!true) {
        fmt.Println(v11, v4)
        delete(v1, v4)
    } else {
        var v15 float64 = 0.1
        fmt.Println(v4, v15, v11)
        if (v6 <= (v6 ^ (^(v6)))) {
            v5 += 18446744073709551615
            v16 := GenS3{F0: (v4 << ((0) & 63)), F1: 131}
            _ = v16
        }
        _ = v15
    }
    if false {
        var v17 uint16 = (v2 ^ v2)
        var v18 float32 = (0.0 * (1e10 / ((3.14159 * 3.14159) - 1e10)))
        var v19 int = (v4 - v4)
        _ = v17
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 11; fi20++ {
        for fi21 := 0; fi21 < 5; fi21++ {
            var v22 uint16 = (v2 >> ((1) & 63))
            switch v3 {
            case "":
                switch v3 {
                case "":
                    v2 -= (v2 / ((v2) | 1))
                case "foo bar":
                    _ = len(v12)
                case "a":
                    v23 := GenS3{F0: (^(fi20)), F1: 2}
                    _ = v23
                }
            case "foo bar":
                var v24 float32 = 1e10
                _ = v24
            }
            v5 += 1
            if true {
                v3 = v3
                var v25 int8 = -37
                _ = v25
            } else {
                v2 -= v2
                _ = len(v1)
                v4 = genFn0(-9223372036854775808, "a")
                fmt.Println(v22, v11)
                fmt.Println(v11, v11, fi21)
            }
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    v11 += (v11 / (((v5 ^ (v11 & v5))) | 1))
    if (false && ((!(!true)) || (true && true))) {
        if false {
            _ = len(v12)
            switch (v4 - genFn0(607, ((v3 + (v3 + "a")) + ""))) {
            case 168:
                fmt.Println(v6, v3)
            case 0:
                delete(v1, (v4 & v4))
            case 30:
                var v26 float32 = (-2.25 + float32(-2.25))
                _ = v26
            default:
                for fi27 := 0; fi27 < 6; fi27++ {
                    var v28 string = v3
                    for fi29 := 0; fi29 < 12; fi29++ {
                        v11 = v5
                        v30 := GenS1{F0: 807}
                        _ = fi29
                        _ = v30
                    }
                    _ = fi27
                    _ = v28
                }
            }
        } else {
            v2 -= (v2 % ((v2) | 1))
            fmt.Println(v5, v5)
            var v31 int = genFn0(9223372036854775807, "foo bar")
            _ = v31
        }
        var v32 uint16 = (v2 - v2)
        v5 += (v5 >> ((0) & 63))
        _ = v32
    }
    if len(v12) > 0 {
        v12[0] = 32768
    }
    v3 = ("\n\t" + "\n\t")
    if (!((float64(-2.25) * float64((3.14159 * 0.1))) == 1e10)) {
        v33 := GenS2{F0: 614, F1: v2}
        _ = v33
    } else {
        var v34 string = "a"
        fmt.Println(v3, v2, v4)
        v3 = ("hello" + v34)
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
