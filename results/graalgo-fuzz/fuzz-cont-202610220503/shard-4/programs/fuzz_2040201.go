package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 int; F2 uint8 }
func genFn0(p0 uint, p1 bool, p2 uint8) int {
    v7 := []int{9223372036854775807, 32767}
    v8 := GenS1{F0: p1, F1: 2}
    var v9 float32 = 0.1
    for fi10 := 0; fi10 < 13; fi10++ {
        if ("hello" > "a") {
            switch (fi10 >> ((uint(p2)) & 63)) {
            case -889:
                if ("foo bar" == (("hello" + "") + ("foo bar" + (("foo bar" + "hello") + "foo bar")))) {
                    var v11 uint16 = 0
                    p0 += 168
                    _ = v11
                } else {
                    for fi12 := 0; fi12 < 5; fi12++ {
                        v13 := GenS1{F0: (((!((-1 <= -97) || p1)) && true) && p1), F1: 734}
                        var v14 uint = p0
                        var v15 float32 = 1.5
                        _ = fi12
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                }
            case 1:
                var v16 float64 = ((3.14159 - (-2.25 - 1e10)) - 1e10)
                _ = v16
            case -2147483648:
                v9 += (3.14159 + -2.25)
            }
        } else {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            var v17 int = fi10
            _ = v17
        }
        if true {
            v9 = float32((-2.25 * (-2.25 - 1e10)))
        } else {
            var v18 uint64 = 510
            _ = v18
        }
        p1 = ("a" < (("\n\t" + "a") + "\n\t"))
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return -541
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = 42
    var v3 int = 1
    var v4 float64 = (1.5 * (((1e10 + (1e10 / 0.1)) + 0.1) - 0.0))
    var v5 string = ""
    var v6 float64 = float64(0.1)
    var v19 uint8 = 0
    v20 := GenS1{F0: true, F1: uint16(v19)}
    if true {
        var v21 string = v5
        v22 := []int{(^(v3)), v3}
        _ = v21
        _ = v22
    }
    var v23 string = v5
    switch v23 {
    case "a":
        var v24 bool = true
        _ = v24
    case "\n\t":
        if (2147483648 < 628) {
            v23 = v5
        }
    default:
        switch "\n\t" {
        case "foo bar":
            v19 = (^(v19))
        case "hello":
            v23 = v23
        default:
            fmt.Println(v2, v23, v23)
        }
    }
    for fi25 := 0; fi25 < 19; fi25++ {
        var v26 uint32 = 2147483647
        v27 := GenS2{F0: (("hello" + "") > "\n\t")}
        _ = fi25
        _ = v26
        _ = v27
    }
    v28 := GenS3{F0: 905, F1: v3, F2: 42}
    v6 = v4
    v2 = 100
    v1[""] = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
