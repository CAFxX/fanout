package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint64 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8; F1 int; F2 uint8 }
func genFn0(p0 uint8, p1 int) uint {
    var v4 float64 = -2.25
    fmt.Println(p1, p0, p0)
    v5 := []int{p1, p1}
    p0 -= 2
    var v6 int32 = 1
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 605
}

func main() {
    var v1 uint16 = 170
    var v2 uint64 = 0
    var v3 uint8 = 42
    switch ((((("foo bar" + "a") + "foo bar") + "") + ("\n\t" + ("\n\t" + "hello"))) + "") {
    case "\n\t":
        if (false || true) {
            if true {
                v1 = (v1 % ((v1) | 1))
                v7 := []int{}
                _ = v7
            } else {
                v8 := GenS1{F0: "hello", F1: "foo bar", F2: 1}
                for fi9 := 0; fi9 < 14; fi9++ {
                    if ("a" >= ("foo bar" + ("a" + "a"))) {
                        v2 += 318
                        var v10 int16 = 361
                        var v11 int64 = 266
                        _ = v10
                        _ = v11
                    } else {
                        var v12 int16 = -40
                        _ = v12
                    }
                    fmt.Println(v1, v2, v3)
                    v13 := []int{341, fi9, -9223372036854775808}
                    _ = fi9
                    _ = v13
                }
                v14 := GenS1{F0: "", F1: "a", F2: (^(v2))}
                _ = v14
                _ = v8
            }
            v15 := GenS2{F0: -91, F1: 2147483647}
            _ = v15
        } else {
            if false {
                v2 += (v2 % (((v2 + (v2 * (v2 % ((v2) | 1))))) | 1))
            }
            var v16 uint64 = (v2 - 967)
            _ = v16
        }
    }
    if (42 >= -9) {
        v17 := []int{4294967295, 32767, 32767, 32768}
        var v18 uint64 = v2
        _ = v17
        _ = v18
    }
    v19 := []int{}
    v20 := GenS3{F0: -128}
    var v21 uint16 = v1
    v2 = (v2 >> ((2) & 63))
    v2 += v2
    var v22 uint = 222
    v1 -= uint16(v3)
    v1 += (v21 * (^(v1)))
    var v23 float32 = 0.1
    for fi24 := 0; fi24 < 20; fi24++ {
        fmt.Println(v22, v23)
        if len(v19) >= 2 {
            v19 = v19[1:len(v19)]
        }
        for fi25 := 0; fi25 < 4; fi25++ {
            v26 := GenS4{F0: (1000 != 928)}
            var v27 int32 = -412
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = fi24
    }
    for fi28 := 0; fi28 < 12; fi28++ {
        if false {
            var v29 string = "hello"
            var v30 float64 = 0.0
            _ = v29
            _ = v30
        }
        if (true || (v2 <= (v2 * v2))) {
            v31 := GenS5{F0: (v3 & (v3 + (v3 * v3))), F1: fi28, F2: 128}
            _ = v31
        } else {
            v32 := GenS5{F0: v3, F1: int(v21), F2: (v3 * (v3 / ((v3) | 1)))}
            var v33 uint8 = v3
            var v34 uint = (v22 & 42)
            _ = v32
            _ = v33
            _ = v34
        }
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
