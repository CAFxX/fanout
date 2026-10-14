package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int64 }
type GenS2 struct { F0 int; F1 uint; F2 int64 }
type GenS3 struct { F0 uint64; F1 uint16; F2 int32 }
func genFn0(p0 int, p1 int8) uint {
    v6 := []int{(p0 >> ((794) & 63)), (-(p0)), (p0 * (p0 % ((2147483647) | 1)))}
    if false {
        switch "a" {
        case "foo bar":
            fmt.Println(p0, p0)
        }
        p1 += p1
    } else {
        p0 += p0
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 uint64 = 0
            if false {
                v8 = 2
            }
            v9 := GenS1{F0: 1, F1: "foo bar", F2: -621}
            _ = fi7
            _ = v8
            _ = v9
        }
        fmt.Println(p0, p0)
    }
    var v10 float32 = 0.1
    v11 := GenS2{F0: (^(p0)), F1: 65535, F2: 297}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v6
    return 4294967295
}
func genFn1() uint32 {
    var v12 uint32 = 34
    var v13 float64 = (float64(3.14159) + ((0.0 - (1.5 / 3.14159)) / 1.0))
    _ = v12
    _ = v13
    return 4294967295
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 32767
    var v3 int = 1
    var v4 bool = true
    var v5 uint16 = 2
    var v14 float32 = 1e10
    var v15 int16 = 0
    v16 := GenS3{F0: 6, F1: v2, F2: 915}
    for fi17 := 0; fi17 < 16; fi17++ {
        var v18 string = (("hello" + (((("\n\t" + "a") + "foo bar") + "foo bar") + "")) + ("hello" + "a"))
        v19 := v1[2]
        _ = fi17
        _ = v18
        _ = v19
    }
    switch ("foo bar" + ("\n\t" + "")) {
    case "a":
        var v20 int64 = -963
        _ = v20
    default:
        v2 += (v2 % ((2) | 1))
    }
    v21 := GenS2{F0: v3, F1: genFn0(32767, 42), F2: -9223372036854775808}
    switch ("a" + "\n\t") {
    case "a":
        var v22 uint32 = genFn1()
        _ = v22
    case "foo bar":
        var v23 string = ("hello" + "")
        _ = v23
    case "hello":
        delete(v1, 1)
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        delete(v1, 2)
        _ = fi24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
