package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 uint8 }

func main() {
    v1 := make(map[int]bool)
    var v2 bool = (1e10 == 3.14159)
    var v3 string = "a"
    var v4 uint64 = 2
    v5 := make(map[string]string)
    v3 = v3
    if ((v2 || (-432 == 1000)) && true) {
        fmt.Println(v4)
        if ((!(false || v2)) || true) {
            var v6 float64 = 1.5
            _ = v6
        }
        v7 := GenS1{F0: 42}
        _ = v7
    } else {
        if (-307 == -1) {
            v8 := GenS1{F0: -9223372036854775808}
            _ = v8
        } else {
            delete(v1, -451)
            var v9 uint64 = 4294967295
            for fi10 := 0; fi10 < 5; fi10++ {
                fmt.Println(v2, v9, v3)
                _ = fi10
            }
            _ = v9
        }
        var v11 int16 = 42
        v2 = false
        _ = v11
    }
    switch 5 {
    case 32768:
        var v12 int = 0
        _ = v12
    case 128:
        v4 = (v4 % ((v4) | 1))
    }
    if false {
        var v13 int16 = 42
        _ = v13
    } else {
        v4 -= v4
    }
    switch v3 {
    case "foo bar":
        if true {
            for fi14 := 0; fi14 < 19; fi14++ {
                var v15 int8 = 1
                for fi16 := 0; fi16 < 20; fi16++ {
                    v4 += v4
                    var v17 int16 = 1
                    v18 := GenS2{F0: 1, F1: 194}
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                _ = fi14
                _ = v15
            }
        }
    case "a":
        v19 := v1[688]
        _ = v19
    case "":
        v20 := []int{42, 50, 4294967295, 765}
        _ = v20
    default:
        if v2 {
            v2 = (227 >= 152)
            var v21 uint64 = (v4 / ((2) | 1))
            _ = v21
        }
    }
    switch (v3 + "hello") {
    case "hello":
        v3 = v3
    case "foo bar":
        v22 := GenS1{F0: -9223372036854775808}
        _ = v22
    case "\n\t":
        var v23 uint64 = v4
        _ = v23
    default:
        var v24 int16 = -446
        _ = v24
    }
    fmt.Println(v2)
    switch v3 {
    case "\n\t":
        var v25 uint8 = 0
        _ = v25
    case "":
        v3 = (((v3 + ("\n\t" + ("hello" + "foo bar"))) + ("a" + ("foo bar" + "hello"))) + "a")
    case "foo bar":
        v4 += (v4 / ((v4) | 1))
    }
    if v2 {
        v2 = false
        v4 = 18446744073709551615
    }
    var v26 int64 = 0
    var v27 int8 = 0
    var v28 int64 = (v26 | v26)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
