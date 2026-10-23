package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS2 struct { F0 string; F1 int64 }
func genFn0() int {
    switch ("hello" + ("foo bar" + "a")) {
    case "":
        var v4 uint64 = 42
        _ = v4
    default:
        var v5 int32 = 42
        _ = v5
    }
    v6 := GenS1{F0: 32767, F1: 526, F2: 32767}
    var v7 string = (("a" + ("\n\t" + "a")) + ("foo bar" + ""))
    _ = v6
    _ = v7
    return -9223372036854775808
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 177
    var v3 int8 = 127
    v3 += v3
    v3 = (v3 % (((v3 * (v3 - -42))) | 1))
    var v8 int64 = 307
    if (!true) {
        var v9 uint16 = 678
        if (!(165 > 2)) {
            var v10 uint32 = 0
            _ = v10
        } else {
            fmt.Println(v8)
        }
        _ = v9
    } else {
        for fi11 := 0; fi11 < 4; fi11++ {
            var v12 bool = (((v8 * (v8 + v8)) >= v8) || false)
            for fi13 := 0; fi13 < 3; fi13++ {
                fmt.Println(fi11, v3)
                var v14 float64 = -2.25
                _ = fi13
                _ = v14
            }
            _ = fi11
            _ = v12
        }
    }
    var v15 bool = (1 >= 2147483647)
    var v16 float32 = -2.25
    switch ("a" + "foo bar") {
    case "\n\t":
        if v15 {
            v8 = v8
            for fi17 := 0; fi17 < 11; fi17++ {
                fmt.Println(v2, fi17)
                v18 := GenS1{F0: 1000, F1: v2, F2: 255}
                _ = fi17
                _ = v18
            }
        } else {
            var v19 float32 = 0.0
            var v20 uint16 = 42
            switch (("\n\t" + ("\n\t" + "")) + ("a" + "\n\t")) {
            case "\n\t":
                fmt.Println(v3, v15)
            }
            _ = v19
            _ = v20
        }
    case "a":
        var v21 uint32 = 4294967295
        _ = v21
    default:
        if true {
            v3 += (v3 & v3)
        } else {
            v22 := GenS2{F0: "hello", F1: v8}
            _ = v22
        }
    }
    switch (("\n\t" + (("hello" + "foo bar") + "\n\t")) + ("foo bar" + "\n\t")) {
    case "foo bar":
        var v23 float64 = float64(float32(0.1))
        _ = v23
    case "a":
        switch 346 {
        case 1:
            var v24 int8 = v3
            _ = v24
        }
    default:
        var v25 uint32 = 218
        _ = v25
    }
    var v26 int64 = 1
    var v27 uint64 = (v2 - 0)
    v28 := GenS1{F0: -924, F1: (v2 ^ v2), F2: 255}
    v15 = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v8
}
