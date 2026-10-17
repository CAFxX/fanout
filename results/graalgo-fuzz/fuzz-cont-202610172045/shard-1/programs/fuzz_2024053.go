package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 int }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]string)
    var v3 uint32 = 0
    v4 := []string{("a" + "a"), "hello"}
    if true {
        v3 += (v3 % (((v3 - (v3 % ((189) | 1)))) | 1))
        fmt.Println(v3)
    }
    switch ("a" + (("" + "foo bar") + "foo bar")) {
    case "\n\t":
        v5 := GenS1{F0: 1}
        _ = v5
    default:
        var v6 string = ((("a" + ("foo bar" + "")) + ("foo bar" + "\n\t")) + (("hello" + (("\n\t" + "") + ("foo bar" + ("foo bar" + "\n\t")))) + "hello"))
        _ = v6
    }
    switch 1 {
    case -9223372036854775808:
        for fi7 := 0; fi7 < 18; fi7++ {
            var v8 int8 = 0
            if (!(false && true)) {
                v9 := GenS2{F0: 42}
                _ = v9
            }
            _ = fi7
            _ = v8
        }
    }
    v3 = (v3 + v3)
    fmt.Println(v3)
    var v10 string = ((("" + "") + "\n\t") + "\n\t")
    if (42 != 1) {
        delete(v2, v3)
        var v11 int = 9223372036854775807
        _ = v11
    }
    var v12 uint8 = 189
    var v13 int32 = -2147483648
    var v14 uint64 = 2
    var v15 uint32 = (v3 | 862)
    v16 := GenS3{F0: 255, F1: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
