package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint16 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int }
func genFn0() uint {
    var v5 float32 = (1.5 - ((-2.25 * (-2.25 * (3.14159 / -2.25))) + 0.0))
    fmt.Println(v5)
    _ = v5
    return 0
}
func genFn1(p0 int8) uint16 {
    var v6 uint32 = 2
    p0 -= 1
    _ = p0
    _ = v6
    return 922
}

func main() {
    var v1 int = 0
    var v2 string = ""
    var v3 string = (("foo bar" + "a") + "")
    var v4 int8 = -128
    if (!false) {
        fmt.Println(v1, v1)
        var v7 float64 = 0.1
        _ = v7
    }
    if (true && (!true)) {
        var v8 int32 = -2147483648
        var v9 uint16 = genFn1((v4 ^ 100))
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 15; fi10++ {
            for fi11 := 0; fi11 < 18; fi11++ {
                var v12 uint64 = 2
                var v13 float32 = (1e10 / 1e10)
                fmt.Println(v12, v3, v12)
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = fi10
        }
    }
    v14 := []int{v1}
    var v15 uint = 595
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    v16 := GenS1{F0: (v15 / (((^(v15))) | 1)), F1: "foo bar", F2: 354}
    switch v1 {
    case 0:
        fmt.Println(v15)
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 int64 = 620
        v19 := GenS1{F0: genFn0(), F1: v2, F2: 32767}
        _ = fi17
        _ = v18
        _ = v19
    }
    v20 := GenS2{F0: genFn0(), F1: (!false)}
    if false {
        fmt.Println(v15, v15)
        for fi21 := 0; fi21 < 15; fi21++ {
            switch (fi21 << ((v15) & 63)) {
            case 9223372036854775807:
                for fi22 := 0; fi22 < 19; fi22++ {
                    if len(v14) > 0 {
                        v14[0] = 251
                    }
                    _ = fi22
                }
            }
            _ = fi21
        }
        v2 = (v2 + v2)
    } else {
        if (0.1 == float32(1.5)) {
            var v23 uint16 = 1
            _ = v23
        } else {
            for fi24 := 0; fi24 < 15; fi24++ {
                var v25 string = v2
                v26 := GenS3{F0: -9223372036854775808}
                _ = fi24
                _ = v25
                _ = v26
            }
            var v27 uint = v15
            _ = v27
        }
        for fi28 := 0; fi28 < 14; fi28++ {
            v14 = append(v14, fi28)
            v1 -= (v1 + 0)
            _ = fi28
        }
    }
    if true {
        v15 -= (v15 - v15)
        v29 := GenS4{F0: (^(v1))}
        _ = v29
    }
    switch v2 {
    case "foo bar":
        v4 = (v4 >> ((v15) & 63))
    case "hello":
        v1 -= v1
    case "a":
        v14 = append(v14, (v1 ^ v1))
    default:
        v15 += (^(v15))
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
