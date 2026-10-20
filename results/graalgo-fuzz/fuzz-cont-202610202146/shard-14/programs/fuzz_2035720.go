package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 int8; F1 uint64; F2 int64 }
func genFn0(p0 uint) int64 {
    var v5 int64 = 42
    var v6 float32 = 3.14159
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 int32 = 42
        _ = fi7
        _ = v8
    }
    p0 += p0
    v9 := []int{32767, int(v5), -983, 494}
    _ = p0
    _ = v5
    _ = v6
    _ = v9
    return 0
}

func main() {
    v1 := []uint32{2147483648, 129, 65535}
    var v2 string = "a"
    var v3 uint = 666
    var v4 int16 = 42
    for fi10 := 0; fi10 < 8; fi10++ {
        v11 := GenS1{F0: 32767, F1: ((true && (((-2.25 + -2.25) + 1.5) == 0.0)) || false)}
        v2 = "a"
        _ = fi10
        _ = v11
    }
    var v12 int16 = 1000
    for fi13 := 0; fi13 < 20; fi13++ {
        v12 = v4
        _ = fi13
    }
    if false {
        var v14 uint8 = 128
        v15 := GenS2{F0: 42, F1: 100}
        _ = v14
        _ = v15
    } else {
        for fi16 := 0; fi16 < 14; fi16++ {
            fmt.Println(v2, v2, v3)
            var v17 uint = (v3 * 1)
            _ = fi16
            _ = v17
        }
        v12 = 1
        v2 = "hello"
    }
    if ((2147483648 >= 2147483647) && false) {
        v3 = (v3 * (v3 + (v3 << ((v3) & 63))))
    }
    if true {
        switch -152 {
        case 1:
            var v18 uint32 = 1
            _ = v18
        case 4294967295:
            var v19 uint16 = 0
            _ = v19
        case -1:
            switch (("" + ("\n\t" + (("foo bar" + "a") + "hello"))) + "foo bar") {
            case "a":
                v20 := GenS1{F0: (v4 * -32768), F1: false}
                _ = v20
            case "\n\t":
                for fi21 := 0; fi21 < 7; fi21++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if len(v1) > 0 {
                        v1[0] = 977
                    }
                    _ = fi21
                }
            }
        }
    }
    fmt.Println(v3)
    var v22 uint64 = 0
    var v23 bool = false
    var v24 float32 = (-2.25 / -2.25)
    var v25 int8 = 104
    var v26 uint64 = 109
    v27 := GenS3{F0: v25, F1: (v3 + v3)}
    v28 := GenS4{F0: v25, F1: v26, F2: -1}
    v26 -= (^(v22))
    if v23 {
        var v29 uint8 = 58
        _ = v29
    }
    var v30 int = -201
    v23 = v23
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v4
}
