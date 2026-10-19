package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint64 }

func main() {
    var v1 bool = true
    var v2 uint32 = 65535
    v3 := make(map[int]bool)
    switch 769 {
    case 32767:
        var v4 int = -2147483648
        _ = v4
    }
    v5 := GenS1{F0: 323, F1: -118, F2: 0}
    fmt.Println(v2, v1)
    for fi6 := 0; fi6 < 7; fi6++ {
        v7 := []int{fi6, fi6}
        _ = fi6
        _ = v7
    }
    var v8 uint64 = 0
    v2 += (v2 * v2)
    v2 = (^(v2))
    v9 := []int{}
    v8 += v8
    if (127 == 1) {
        v2 -= (v2 ^ (v2 >> ((4294967295) & 63)))
        v2 = (v2 * v2)
        if ((("a" + "hello") + ((("\n\t" + "\n\t") + "\n\t") + ("a" + "\n\t"))) == ("" + "foo bar")) {
            var v10 int16 = 1
            _ = v10
        } else {
            var v11 int8 = -1
            for fi12 := 0; fi12 < 16; fi12++ {
                if (v1 && true) {
                    v11 += -1
                    v8 -= v8
                    v1 = v1
                } else {
                    v13 := GenS1{F0: 9223372036854775807, F1: v11, F2: v8}
                    switch ("\n\t" + (("a" + (((((((((("hello" + ("" + ("foo bar" + ("hello" + "\n\t")))) + "foo bar") + "\n\t") + ("foo bar" + "foo bar")) + (((("\n\t" + ("\n\t" + "")) + ("" + (("\n\t" + ("\n\t" + "\n\t")) + "hello"))) + "foo bar") + "a")) + "hello") + "\n\t") + "hello") + "foo bar") + (("foo bar" + "") + ((((("foo bar" + "hello") + "") + "\n\t") + "foo bar") + "a")))) + ("\n\t" + "foo bar"))) {
                    case "a":
                        var v14 uint = uint(v8)
                        _ = v14
                    case "hello":
                        fmt.Println(fi12)
                    case "foo bar":
                        fmt.Println(fi12, v1, v2)
                    }
                    _ = v13
                }
                _ = fi12
            }
            _ = v11
        }
    } else {
        for fi15 := 0; fi15 < 9; fi15++ {
            var v16 int16 = 32767
            _ = fi15
            _ = v16
        }
        fmt.Println(v8, v1, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v8
    _ = v9
}
