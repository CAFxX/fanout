package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 int16 }
func genFn0(p0 int64) int {
    p0 = -392
    for fi4 := 0; fi4 < 10; fi4++ {
        if true {
            fmt.Println(p0, fi4)
            var v5 float64 = 1e10
            _ = v5
        }
        _ = fi4
    }
    var v6 uint64 = 18446744073709551615
    _ = p0
    _ = v6
    return 0
}
func genFn1() string {
    var v7 uint = 453
    v7 = (^(v7))
    fmt.Println(v7)
    _ = v7
    return "\n\t"
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[string]string)
    var v3 uint16 = 978
    v3 = v3
    fmt.Println(v3, v3)
    var v8 uint = 18446744073709551615
    v8 -= v8
    switch genFn0(v1) {
    case 1:
        v1 += 9223372036854775807
    case -1:
        v2[("hello" + "hello")] = ("a" + (("a" + "foo bar") + "foo bar"))
    }
    switch "hello" {
    case "\n\t":
        if false {
            var v9 string = ((("foo bar" + "\n\t") + ("foo bar" + "\n\t")) + "a")
            for fi10 := 0; fi10 < 1; fi10++ {
                v2[v9] = v9
                v9 = ((((v9 + (v9 + "\n\t")) + v9) + v9) + v9)
                _ = fi10
            }
            _ = v9
        } else {
            for fi11 := 0; fi11 < 3; fi11++ {
                fmt.Println(v3)
                _ = fi11
            }
            for fi12 := 0; fi12 < 16; fi12++ {
                v13 := GenS1{F0: (-(fi12))}
                var v14 uint = 4294967295
                _ = fi12
                _ = v13
                _ = v14
            }
            v8 = v8
        }
    case "foo bar":
        var v15 uint8 = 183
        _ = v15
    }
    fmt.Println(v3)
    v16 := GenS2{F0: (^(v3))}
    switch -638 {
    case -2147483648:
        v17 := []int{950}
        _ = v17
    }
    v18 := GenS3{F0: "", F1: int16(v8)}
    v8 = v8
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v8
}
