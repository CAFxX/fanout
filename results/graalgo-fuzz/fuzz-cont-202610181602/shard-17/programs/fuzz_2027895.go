package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint8) uint {
    var v4 uint8 = p0
    v5 := GenS1{F0: 1}
    _ = p0
    _ = v4
    _ = v5
    return 814
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]string)
    var v3 int64 = 0
    switch (("hello" + "a") + "\n\t") {
    case "foo bar":
        v6 := []int{}
        _ = v6
    default:
        v3 = v3
    }
    var v7 int64 = (v3 & (v3 & (v3 >> ((1) & 63))))
    var v8 int16 = 32767
    switch ("" + "hello") {
    case "hello":
        for fi9 := 0; fi9 < 7; fi9++ {
            v10 := []int{fi9}
            switch ("hello" + ("\n\t" + (((((("" + "a") + "") + "foo bar") + ("" + ((("" + ((("\n\t" + (("a" + "") + "a")) + "a") + ("" + "hello"))) + "foo bar") + (("a" + "a") + "foo bar")))) + "foo bar") + "hello"))) {
            case "\n\t":
                var v11 int64 = -1
                _ = v11
            default:
                v7 += -539
            }
            v7 = -1
            _ = fi9
            _ = v10
        }
    }
    v12 := GenS1{F0: (v3 ^ v3)}
    var v13 uint32 = 2
    v3 += -1
    v8 = (v8 + 32767)
    if (-2147483648 >= 1) {
        v14 := GenS1{F0: (v7 * 1)}
        _ = v14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
