package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 uint }

func main() {
    v1 := make(map[string]string)
    var v2 bool = (0.0 >= (3.14159 + 1.5))
    v3 := []int{689, 606, 4294967295, -70}
    var v4 uint16 = 825
    var v5 uint = 0
    var v6 int64 = -9223372036854775808
    var v7 int32 = 0
    fmt.Println(v6)
    v6 -= (v6 / (((v6 % ((v6) | 1))) | 1))
    if true {
        v2 = v2
    } else {
        fmt.Println(v5, v4)
        if false {
            var v8 int64 = -9223372036854775808
            if v2 {
                var v9 uint64 = 1
                v6 = (^(v8))
                _ = v9
            }
            _ = v8
        }
        v2 = ((!((v4 >> ((v5) & 63)) > (v4 * (v4 - v4)))) || true)
    }
    v4 = 216
    var v10 float32 = 1.5
    switch 42 {
    case 650:
        for fi11 := 0; fi11 < 5; fi11++ {
            var v12 bool = ("\n\t" > "a")
            var v13 uint64 = 2
            _ = fi11
            _ = v12
            _ = v13
        }
    case 358:
        v6 += -1
    case 32768:
        var v14 bool = false
        _ = v14
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        v16 := GenS1{F0: ("" >= "hello"), F1: fi15}
        switch 1 {
        case 0:
            if true {
                v17 := GenS1{F0: (1 > (v7 / (((v7 >> (((v5 & v5)) & 63))) | 1))), F1: (-(fi15))}
                v18 := GenS2{F0: (v7 << (((v5 >> (((v5 * (v5 - (v5 - 65535)))) & 63))) & 63)), F1: 57}
                v19 := GenS3{F0: v6, F1: 719}
                _ = v17
                _ = v18
                _ = v19
            }
        case 42:
            v5 = v5
        case -2147483648:
            var v20 string = "foo bar"
            _ = v20
        default:
            fmt.Println(v10, v4)
        }
        _ = fi15
        _ = v16
    }
    fmt.Println(v10, v7)
    if ((0.0 > (0.0 - 1.5)) && true) {
        switch ((("hello" + "\n\t") + "\n\t") + "hello") {
        case "hello":
            var v21 float64 = 1.5
            _ = v21
        }
        delete(v1, "a")
        var v22 uint16 = 255
        _ = v22
    } else {
        var v23 float64 = 1e10
        v24 := GenS4{F0: 65535}
        _ = v23
        _ = v24
    }
    var v25 float64 = 1.5
    v25 += v25
    var v26 int64 = (^(v6))
    if v2 {
        v6 -= (v26 | v6)
        for fi27 := 0; fi27 < 4; fi27++ {
            delete(v1, ("a" + ("a" + ("\n\t" + "foo bar"))))
            _ = fi27
        }
    } else {
        if v2 {
            delete(v1, "\n\t")
        }
        var v28 int64 = v6
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
