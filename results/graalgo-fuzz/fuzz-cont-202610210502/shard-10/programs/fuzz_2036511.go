package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 bool; F2 string }
type GenS3 struct { F0 int64 }
func genFn0(p0 int, p1 int) uint64 {
    var v7 string = "a"
    v8 := GenS1{F0: p0}
    v9 := GenS1{F0: (-(p1))}
    var v10 float32 = 0.0
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 96
}
func genFn1(p0 uint) uint16 {
    if (1 != -2147483648) {
        switch 128 {
        case 760:
            fmt.Println(p0)
        case 42:
            if (255 != 42) {
                var v11 float32 = 1e10
                _ = v11
            }
        default:
            v12 := []int{170, 128, 42}
            _ = v12
        }
    }
    var v13 int16 = -275
    p0 += p0
    v13 = v13
    _ = p0
    _ = v13
    return 65535
}

func main() {
    var v1 float64 = (1.5 * (1e10 * 0.0))
    var v2 int16 = -32768
    var v3 uint16 = 32767
    v4 := make(map[string]bool)
    var v5 int = 32767
    var v6 int32 = -1
    v14 := GenS1{F0: (v5 >> ((992) & 63))}
    var v15 int16 = (v2 - v2)
    if ((v6 >= (v6 ^ (v6 | (v6 - (^(v6)))))) && (-2147483648 >= 2147483647)) {
        switch "" {
        case "hello":
            v5 = (v5 - (v5 & (v5 * v5)))
        }
        var v16 int32 = (-(v6))
        switch ("" + "hello") {
        case "hello":
            v17 := GenS2{F0: 91, F1: (((((1 < 942) && false) || true) && true) && false), F2: ""}
            _ = v17
        case "a":
            v18 := []int{(v5 / (((v5 >> ((255) & 63))) | 1)), (v5 / ((251) | 1))}
            _ = v18
        case "\n\t":
            var v19 bool = true
            _ = v19
        default:
            delete(v4, ("foo bar" + ("foo bar" + (("foo bar" + "foo bar") + ""))))
        }
        _ = v16
    } else {
        v2 = 32767
        switch (("foo bar" + "foo bar") + "\n\t") {
        case "foo bar":
            for fi20 := 0; fi20 < 17; fi20++ {
                var v21 uint64 = 775
                _ = fi20
                _ = v21
            }
        case "hello":
            if true {
                v15 = (v2 % ((v2) | 1))
                v22 := []int{-573, v5, v5}
                _ = v22
            } else {
                v23 := []int{(v5 * v5), (v5 << ((188) & 63))}
                _ = v23
            }
        }
        switch (v5 ^ (v5 << ((405) & 63))) {
        case 328:
            fmt.Println(v3)
        case -1:
            for fi24 := 0; fi24 < 1; fi24++ {
                fmt.Println(v2, v5)
                v25 := GenS3{F0: 9223372036854775807}
                v26 := []int{}
                _ = fi24
                _ = v25
                _ = v26
            }
        default:
            v15 = -32768
        }
    }
    var v27 uint16 = v3
    delete(v4, "")
    for fi28 := 0; fi28 < 3; fi28++ {
        var v29 float32 = 1.5
        _ = fi28
        _ = v29
    }
    switch "a" {
    case "hello":
        v5 -= 128
    case "\n\t":
        switch "a" {
        case "":
            if (float64((0.1 / 1e10)) >= (v1 * 0.0)) {
                v5 += v5
            }
        case "foo bar":
            var v30 int16 = (v15 / (((v2 % (((-(v2))) | 1))) | 1))
            _ = v30
        case "hello":
            v15 = 1
        default:
            for fi31 := 0; fi31 < 5; fi31++ {
                var v32 int32 = int32(fi31)
                for fi33 := 0; fi33 < 14; fi33++ {
                    v34 := []int{(v5 / ((fi31) | 1)), (fi33 >> ((65535) & 63)), fi33}
                    _ = fi33
                    _ = v34
                }
                v5 -= int(v6)
                _ = fi31
                _ = v32
            }
        }
    case "a":
        var v35 uint8 = 2
        _ = v35
    default:
        var v36 uint64 = 4294967295
        _ = v36
    }
    v27 += v27
    v37 := GenS1{F0: -9223372036854775808}
    var v38 string = (("\n\t" + "hello") + "a")
    var v39 uint = 18446744073709551615
    var v40 uint32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v27
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
