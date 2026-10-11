package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 string; F1 int8; F2 int16 }
func genFn0(p0 bool) bool {
    switch 4294967295 {
    case 42:
        p0 = true
    case 434:
        for fi7 := 0; fi7 < 18; fi7++ {
            var v8 uint64 = 835
            for fi9 := 0; fi9 < 9; fi9++ {
                p0 = false
                _ = fi9
            }
            _ = fi7
            _ = v8
        }
    default:
        p0 = (!p0)
    }
    var v10 string = "hello"
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 float32 = float32((0.1 / (-2.25 * 1e10)))
        v13 := GenS1{F0: p0}
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: p0}
    _ = p0
    _ = v10
    _ = v14
    return (v10 > v10)
}
func genFn1() int8 {
    var v15 uint32 = 2147483648
    v15 = 65535
    v16 := []int{2147483647, -798}
    v17 := GenS1{F0: false}
    v15 = 2
    _ = v15
    _ = v16
    _ = v17
    return 42
}

func main() {
    var v1 float32 = (1.5 + -2.25)
    var v2 int = 128
    var v3 int = v2
    var v4 float32 = v1
    var v5 uint64 = 18446744073709551615
    var v6 uint8 = 127
    switch ("hello" + ("foo bar" + "a")) {
    case "hello":
        fmt.Println(v1, v2, v6)
    case "\n\t":
        v3 += v3
    case "":
        v18 := GenS2{F0: 664, F1: -2147483648}
        _ = v18
    }
    v19 := GenS3{F0: "\n\t", F1: 1, F2: -32768}
    for fi20 := 0; fi20 < 11; fi20++ {
        var v21 uint = 2
        _ = fi20
        _ = v21
    }
    if (1 >= 42) {
        v22 := GenS3{F0: ("a" + "foo bar"), F1: 100, F2: 42}
        if (42 > 32767) {
            var v23 int16 = 0
            v23 -= (^(v23))
            switch "" {
            case "foo bar":
                var v24 float64 = -2.25
                _ = v24
            }
            _ = v23
        }
        fmt.Println(v2)
        _ = v22
    }
    switch ("\n\t" + "hello") {
    case "\n\t":
        for fi25 := 0; fi25 < 11; fi25++ {
            fmt.Println(fi25)
            _ = fi25
        }
    case "a":
        if (504 < 42) {
            v6 -= v6
            fmt.Println(v6)
        } else {
            for fi26 := 0; fi26 < 1; fi26++ {
                v2 += (^(v3))
                _ = fi26
            }
            var v27 int64 = int64(v6)
            _ = v27
        }
    case "hello":
        var v28 int = v3
        _ = v28
    default:
        switch (v2 & v2) {
        case -511:
            if ((v6 % ((v6) | 1)) < (v6 >> ((255) & 63))) {
                for fi29 := 0; fi29 < 10; fi29++ {
                    for fi30 := 0; fi30 < 13; fi30++ {
                        v1 -= v1
                        _ = fi30
                    }
                    v4 = 1.5
                    v4 = v4
                    _ = fi29
                }
                var v31 uint32 = 0
                _ = v31
            }
        case 127:
            v1 += (((0.1 * v4) * (v4 * float32(float64(v1)))) - v1)
        }
    }
    fmt.Println(v3)
    var v32 bool = (("\n\t" < "") || true)
    v3 = v3
    v5 += 42
    var v33 int64 = 9223372036854775807
    v3 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
