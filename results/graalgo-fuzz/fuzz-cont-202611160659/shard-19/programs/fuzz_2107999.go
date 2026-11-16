package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 int8; F2 string }
func genFn0(p0 int, p1 int32, p2 int8) uint {
    v5 := GenS1{F0: "hello"}
    v6 := GenS2{F0: (p0 * -2147483648), F1: p2, F2: ""}
    fmt.Println(p2, p2)
    var v7 int8 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 747
}

func main() {
    var v1 int16 = 42
    v2 := []int{42, 128, -9223372036854775808, 32768}
    var v3 int64 = 0
    v4 := make(map[string]string)
    fmt.Println(v3, v3)
    if false {
        v8 := GenS2{F0: 127, F1: 0, F2: "a"}
        if (!true) {
            if true {
                v4[("foo bar" + (("" + ("\n\t" + ("" + ("foo bar" + "foo bar")))) + ("a" + "a")))] = (("a" + "hello") + "foo bar")
            }
        }
        _ = v8
    } else {
        v9 := GenS1{F0: ("" + "hello")}
        _ = v9
    }
    if (true && true) {
        v10 := GenS2{F0: 2147483647, F1: 42, F2: ""}
        _ = v10
    } else {
        var v11 bool = true
        _ = v11
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        fmt.Println(fi12)
        delete(v4, "\n\t")
        var v13 int32 = 1
        _ = fi12
        _ = v13
    }
    delete(v4, ("\n\t" + "foo bar"))
    v1 = 42
    switch "foo bar" {
    case "hello":
        v14 := GenS1{F0: ""}
        _ = v14
    }
    if true {
        var v15 int = 0
        for fi16 := 0; fi16 < 18; fi16++ {
            var v17 int64 = (v3 << ((240) & 63))
            v17 += (v17 % (((v17 << ((18446744073709551615) & 63))) | 1))
            _ = fi16
            _ = v17
        }
        var v18 int8 = 1
        _ = v15
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
