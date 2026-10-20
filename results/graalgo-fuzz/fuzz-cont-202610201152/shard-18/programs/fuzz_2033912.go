package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 uint32 }

func main() {
    var v1 int8 = 1
    var v2 string = (("foo bar" + "") + "foo bar")
    var v3 float32 = float32((1.5 * 1.5))
    var v4 string = (("hello" + v2) + v2)
    v5 := make(map[uint32]bool)
    fmt.Println(v3, v4)
    if true {
        switch v2 {
        case "hello":
            fmt.Println(v3, v1)
        case "":
            var v6 int16 = 0
            _ = v6
        case "foo bar":
            if (42 != 2) {
                var v7 float64 = 1e10
                v3 += float32(v7)
                _ = v7
            }
        }
        v8 := v5[0]
        _ = v8
    } else {
        v5[4294967295] = (-254 != 351)
    }
    v3 -= ((v3 * float32(-2.25)) * v3)
    switch (("" + (v2 + ("\n\t" + ("\n\t" + v4)))) + "a") {
    case "foo bar":
        v3 = float32(0.1)
    case "hello":
        var v9 uint = 408
        _ = v9
    case "":
        if (true && true) {
            for fi10 := 0; fi10 < 2; fi10++ {
                for fi11 := 0; fi11 < 19; fi11++ {
                    var v12 uint = 255
                    _ = fi11
                    _ = v12
                }
                _ = fi10
            }
        }
    default:
        if false {
            var v13 int16 = -1
            v2 = ((v2 + v4) + v4)
            _ = v13
        }
    }
    var v14 uint32 = 155
    var v15 uint16 = 32767
    var v16 int64 = 462
    fmt.Println(v4, v15)
    if (false && false) {
        for fi17 := 0; fi17 < 12; fi17++ {
            if (true && ((2147483647 > v14) && false)) {
                v18 := GenS1{F0: 1, F1: -880, F2: 0}
                _ = v18
            }
            _ = fi17
        }
        v5[466] = true
    }
    var v19 int16 = -1
    if true {
        var v20 float32 = (float32((0.0 / float64(-2.25))) * v3)
        switch ("" + v2) {
        case "\n\t":
            if true {
                fmt.Println(v1, v4)
            } else {
                for fi21 := 0; fi21 < 6; fi21++ {
                    if ((!false) || (4294967295 >= v14)) {
                        var v22 bool = (!true)
                        v23 := GenS1{F0: v19, F1: (v19 + 42), F2: 1}
                        var v24 int64 = int64(fi21)
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    _ = fi21
                }
                var v25 uint = 255
                v26 := GenS2{F0: 42}
                _ = v25
                _ = v26
            }
        default:
            v27 := []int{42, -397, 2147483647}
            _ = v27
        }
        _ = v20
    }
    switch v4 {
    case "":
        var v28 uint64 = 0
        _ = v28
    case "\n\t":
        fmt.Println(v19, v15, v2)
    }
    v4 = v4
    if true {
        v3 -= 0.1
    } else {
        fmt.Println(v3, v14, v4)
        if true {
            fmt.Println(v2, v15, v4)
        } else {
            v19 = -32768
        }
    }
    v29 := GenS3{F0: v19, F1: (^(v14))}
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
