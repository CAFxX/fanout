package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int32 }

func main() {
    var v1 uint16 = 255
    v2 := make(map[string]bool)
    var v3 uint16 = 255
    v4 := v2[(((("a" + ("" + (("\n\t" + ("foo bar" + "")) + ("foo bar" + "")))) + "a") + "") + ("a" + ("a" + (("foo bar" + ("foo bar" + (((("hello" + ("\n\t" + "")) + "foo bar") + "") + (("" + "a") + "foo bar")))) + "hello"))))]
    fmt.Println(v3, v1, v4)
    v1 = (v3 + (^(v1)))
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 uint64 = 4294967295
        fmt.Println(v1, v3)
        _ = fi5
        _ = v6
    }
    v4 = ((false || true) && ((true && v4) && true))
    fmt.Println(v3)
    var v7 uint64 = 18446744073709551615
    fmt.Println(v7)
    if ((v1 - v3) < (v1 >> ((744) & 63))) {
        var v8 float32 = 1.5
        for fi9 := 0; fi9 < 7; fi9++ {
            v1 -= 1
            var v10 int32 = 997
            v11 := []int{}
            _ = fi9
            _ = v10
            _ = v11
        }
        v2["a"] = v4
        _ = v8
    } else {
        for fi12 := 0; fi12 < 10; fi12++ {
            v13 := GenS1{F0: v7, F1: "\n\t", F2: -2147483648}
            var v14 int8 = 15
            for fi15 := 0; fi15 < 18; fi15++ {
                var v16 uint64 = (v7 >> ((18446744073709551615) & 63))
                switch (fi15 / (((^(fi15))) | 1)) {
                case 127:
                    var v17 bool = v4
                    _ = v17
                case -957:
                    fmt.Println(v14, v1)
                }
                var v18 int32 = -1
                var v19 uint8 = 203
                v20 := v2["hello"]
                _ = fi15
                _ = v16
                _ = v18
                _ = v19
                _ = v20
            }
            _ = fi12
            _ = v13
            _ = v14
        }
        v2["a"] = ("foo bar" < "foo bar")
        if v4 {
            for fi21 := 0; fi21 < 4; fi21++ {
                for fi22 := 0; fi22 < 12; fi22++ {
                    var v23 int32 = 0
                    v1 += (v1 >> ((406) & 63))
                    _ = fi22
                    _ = v23
                }
                v24 := []int{(fi21 - (fi21 & (fi21 * (fi21 + (fi21 + fi21))))), fi21, fi21, (fi21 * fi21)}
                _ = fi21
                _ = v24
            }
            v25 := v2[("a" + "foo bar")]
            _ = v25
        } else {
            for fi26 := 0; fi26 < 20; fi26++ {
                v1 = (v1 & 65535)
                _ = fi26
            }
            var v27 uint = 42
            _ = v27
        }
    }
    var v28 string = "hello"
    var v29 uint64 = (^(v7))
    switch (v28 + "hello") {
    case "":
        for fi30 := 0; fi30 < 7; fi30++ {
            fmt.Println(v1, v28, v7)
            v31 := GenS1{F0: uint64(v1), F1: (((("a" + v28) + v28) + v28) + (v28 + v28)), F2: 2147483647}
            _ = fi30
            _ = v31
        }
    case "a":
        for fi32 := 0; fi32 < 19; fi32++ {
            if (1 < 42) {
                v28 = v28
            }
            fmt.Println(fi32)
            _ = fi32
        }
    default:
        v7 += 2
    }
    for fi33 := 0; fi33 < 10; fi33++ {
        var v34 int32 = 42
        _ = fi33
        _ = v34
    }
    var v35 string = v28
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v35
    _ = v4
    _ = v7
}
