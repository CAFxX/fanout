package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int8 }
func genFn0(p0 int16, p1 int32) uint64 {
    v5 := GenS1{F0: 1, F1: 356, F2: 127}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v5
    return 4294967295
}

func main() {
    var v1 bool = false
    var v2 bool = true
    var v3 uint16 = 42
    v4 := make(map[string]bool)
    fmt.Println(v3, v3)
    var v6 bool = v2
    switch -671 {
    case 32768:
        switch 32768 {
        case -1:
            switch "\n\t" {
            case "":
                switch -2147483648 {
                case 837:
                    fmt.Println(v2)
                case 4294967295:
                    v7 := GenS1{F0: 405, F1: 1, F2: 0}
                    _ = v7
                case 0:
                    var v8 float64 = 1e10
                    _ = v8
                }
            case "foo bar":
                var v9 uint = 4294967295
                _ = v9
            }
        default:
            var v10 uint = 2
            _ = v10
        }
    case -9223372036854775808:
        var v11 float32 = float32((-2.25 / (3.14159 / (1e10 / float64((0.1 - (0.1 / 1e10)))))))
        _ = v11
    case 157:
        v12 := GenS1{F0: 42, F1: 65535, F2: -128}
        _ = v12
    default:
        if v1 {
            for fi13 := 0; fi13 < 4; fi13++ {
                v6 = (372 >= 0)
                switch ("a" + ((("foo bar" + (("hello" + "a") + "hello")) + ("foo bar" + "a")) + "")) {
                case "hello":
                    var v14 uint8 = 156
                    _ = v14
                }
                _ = fi13
            }
        }
    }
    if (1 == 128) {
        switch 0 {
        case 42:
            v3 = (v3 & 779)
        case -897:
            var v15 bool = true
            _ = v15
        }
    } else {
        if v1 {
            for fi16 := 0; fi16 < 2; fi16++ {
                fmt.Println(v2)
                _ = fi16
            }
            var v17 int = -2147483648
            _ = v17
        }
    }
    v18 := v4["a"]
    var v19 int8 = -128
    var v20 float64 = -2.25
    v21 := GenS1{F0: 2, F1: 42, F2: (v19 ^ (-(v19)))}
    for fi22 := 0; fi22 < 3; fi22++ {
        v23 := GenS1{F0: 42, F1: 239, F2: v19}
        switch (fi22 >> ((65535) & 63)) {
        case 32767:
            switch ("a" + "foo bar") {
            case "hello":
                var v24 uint16 = (^(v3))
                _ = v24
            case "\n\t":
                v3 -= v3
            }
        case 2147483647:
            var v25 float32 = (1e10 - (float32((v20 + 0.0)) + (-2.25 - 1.5)))
            _ = v25
        case 42:
            var v26 uint64 = 42
            _ = v26
        }
        var v27 int8 = v19
        _ = fi22
        _ = v23
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v6
}
