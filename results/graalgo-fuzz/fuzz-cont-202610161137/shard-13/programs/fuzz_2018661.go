package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }

func main() {
    var v1 uint64 = 4294967295
    v2 := []int{}
    v3 := []string{("hello" + "foo bar"), "foo bar", ("" + "hello"), ((("a" + "hello") + (("\n\t" + ((("foo bar" + ("" + "a")) + (("" + ("foo bar" + ("foo bar" + "foo bar"))) + "")) + ("foo bar" + ""))) + "hello")) + "")}
    var v4 uint64 = (v1 / ((v1) | 1))
    v5 := make(map[string]string)
    for fi6 := 0; fi6 < 8; fi6++ {
        v4 = (v4 - 908)
        for fi7 := 0; fi7 < 7; fi7++ {
            _ = len(v5)
            var v8 int16 = -1
            _ = fi7
            _ = v8
        }
        var v9 int = fi6
        _ = fi6
        _ = v9
    }
    if (!true) {
        v5["\n\t"] = "hello"
        if true {
            v10 := GenS1{F0: ((127 >= 1) && true), F1: int64(v4)}
            _ = v10
        }
    }
    _ = len(v5)
    var v11 int32 = 125
    if (!(0 == 65535)) {
        var v12 bool = (((1.5 / 0.1) == 0.0) && true)
        var v13 uint16 = 704
        if (v12 || (-2.25 < 0.0)) {
            for fi14 := 0; fi14 < 5; fi14++ {
                var v15 uint8 = 182
                v15 = 1
                _ = fi14
                _ = v15
            }
        } else {
            for fi16 := 0; fi16 < 1; fi16++ {
                fmt.Println(v4, v1)
                v1 = 2
                v1 = v4
                _ = fi16
            }
        }
        _ = v12
        _ = v13
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        v4 = 18446744073709551615
        for fi18 := 0; fi18 < 2; fi18++ {
            switch "hello" {
            case "":
                v1 += 4294967295
            case "foo bar":
                var v19 int = 2147483647
                _ = v19
            default:
                v11 = (v11 & (-(v11)))
            }
            _ = fi18
        }
        var v20 bool = (!true)
        _ = fi17
        _ = v20
    }
    v1 -= 86
    fmt.Println(v1, v4, v4)
    v11 = v11
    var v21 float32 = 0.1
    for fi22 := 0; fi22 < 2; fi22++ {
        if (float64(float32(0.1)) == (1e10 / (-2.25 - 1e10))) {
            var v23 uint8 = 42
            for fi24 := 0; fi24 < 13; fi24++ {
                if true {
                    var v25 uint32 = 1
                    _ = v25
                } else {
                    v5["\n\t"] = ("a" + ("a" + (("" + ("foo bar" + "")) + "foo bar")))
                }
                _ = fi24
            }
            var v26 int8 = -102
            _ = v23
            _ = v26
        }
        _ = fi22
    }
    v1 = v4
    var v27 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
