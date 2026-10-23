package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int32 }
type GenS2 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS3 struct { F0 int8; F1 bool; F2 int }
type GenS4 struct { F0 int64; F1 int64; F2 uint16 }
func genFn0(p0 int16) string {
    v6 := GenS1{F0: "\n\t", F1: ("a" + "\n\t"), F2: 1}
    var v7 float32 = (-2.25 * -2.25)
    for fi8 := 0; fi8 < 10; fi8++ {
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 string = "foo bar"
            switch v10 {
            case "hello":
                var v11 bool = true
                _ = v11
            default:
                v10 = v10
            }
            var v12 uint16 = 2
            _ = fi9
            _ = v10
            _ = v12
        }
        _ = fi8
    }
    fmt.Println(p0)
    _ = p0
    _ = v6
    _ = v7
    return ""
}

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 255
    var v3 int32 = 2147483647
    v4 := make(map[int]int)
    v5 := make(map[int]bool)
    _ = len(v1)
    v13 := []int{311, 32768}
    fmt.Println(v3)
    if (127 < 42) {
        var v14 bool = false
        for fi15 := 0; fi15 < 19; fi15++ {
            v16 := GenS1{F0: ("\n\t" + "foo bar"), F1: ((("\n\t" + ("a" + "\n\t")) + "hello") + ("foo bar" + (("" + ("a" + "hello")) + (("hello" + "hello") + ("foo bar" + "hello"))))), F2: (v3 & v3)}
            v3 -= -2147483648
            _ = fi15
            _ = v16
        }
        v17 := GenS2{F0: v2, F1: 42, F2: -9223372036854775808}
        _ = v14
        _ = v17
    }
    v18 := GenS2{F0: (v2 * (v2 * v2)), F1: 65535, F2: 42}
    for fi19 := 0; fi19 < 16; fi19++ {
        fmt.Println(fi19)
        _ = fi19
    }
    switch "a" {
    case "a":
        v20 := GenS3{F0: 1, F1: (true && (!true)), F2: 32767}
        _ = v20
    }
    v21 := GenS4{F0: -741, F1: -938, F2: (v2 | (v2 * 32767))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
