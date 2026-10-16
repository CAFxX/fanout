package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int8; F1 int32 }

func main() {
    var v1 float64 = -2.25
    var v2 int = -1
    v3 := make(map[string]bool)
    for fi4 := 0; fi4 < 4; fi4++ {
        var v5 float64 = (v1 - v1)
        for fi6 := 0; fi6 < 15; fi6++ {
            delete(v3, "a")
            switch (("a" + ((("hello" + "") + "a") + "a")) + "foo bar") {
            case "\n\t":
                fmt.Println(v1, v2)
            }
            _ = fi6
        }
        v7 := GenS1{F0: -32768, F1: 0}
        _ = fi4
        _ = v5
        _ = v7
    }
    var v8 int64 = 9223372036854775807
    v9 := GenS1{F0: 1000, F1: 1}
    if (true && (143 <= 2147483648)) {
        if (2 >= 65535) {
            v8 -= (-(v8))
            v10 := GenS1{F0: 1, F1: 42}
            for fi11 := 0; fi11 < 14; fi11++ {
                if true {
                    switch "foo bar" {
                    case "a":
                        var v12 float32 = 0.1
                        _ = v12
                    }
                }
                _ = fi11
            }
            _ = v10
        }
        v8 = v8
        for fi13 := 0; fi13 < 9; fi13++ {
            v8 -= (v8 / ((v8) | 1))
            v14 := []int{-530, fi13, (fi13 << ((4294967295) & 63))}
            _ = fi13
            _ = v14
        }
    } else {
        var v15 float64 = (v1 / v1)
        _ = v15
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        for fi17 := 0; fi17 < 20; fi17++ {
            var v18 float64 = ((v1 / v1) + -2.25)
            switch 32767 {
            case 127:
                var v19 int64 = 1
                _ = v19
            }
            _ = fi17
            _ = v18
        }
        v20 := GenS1{F0: 42, F1: 2}
        v8 -= -229
        _ = fi16
        _ = v20
    }
    fmt.Println(v8, v1, v8)
    if (!true) {
        if (!(false || true)) {
            fmt.Println(v1, v8, v1)
            for fi21 := 0; fi21 < 9; fi21++ {
                fmt.Println(fi21, fi21, v2)
                v1 = v1
                var v22 uint32 = 4294967295
                _ = fi21
                _ = v22
            }
        } else {
            v23 := GenS2{F0: -1, F1: 0}
            _ = v23
        }
        var v24 int64 = -9223372036854775808
        _ = v24
    }
    v1 = 3.14159
    var v25 uint8 = 178
    switch (^(v2)) {
    case -2147483648:
        v25 += v25
    case 2147483647:
        v1 = 1e10
    case -602:
        v1 = float64(1.5)
    }
    var v26 int32 = 0
    for fi27 := 0; fi27 < 2; fi27++ {
        v8 = (^(v8))
        var v28 string = ("hello" + "foo bar")
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v8
    _ = v9
}
