package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() string {
    var v5 bool = false
    var v6 float32 = (((0.1 * 1.5) + 1.5) + (-2.25 / 1.5))
    fmt.Println(v5)
    _ = v5
    _ = v6
    return "hello"
}
func genFn1(p0 uint16, p1 int) int {
    v7 := []int{2147483647, 2147483647}
    var v8 int = 438
    switch "hello" {
    case "foo bar":
        for fi9 := 0; fi9 < 11; fi9++ {
            for fi10 := 0; fi10 < 17; fi10++ {
                if (0.0 == 0.0) {
                    v8 -= (v8 % ((v8) | 1))
                }
                switch "foo bar" {
                case "a":
                    p0 = p0
                    fmt.Println(p0, fi10)
                    if len(v7) > 0 {
                        v7[0] = fi9
                    }
                    var v11 int64 = 9223372036854775807
                    var v12 uint16 = (p0 >> ((724) & 63))
                    _ = v11
                    _ = v12
                default:
                    var v13 uint32 = 429
                    _ = v13
                }
                var v14 uint64 = 1
                _ = fi10
                _ = v14
            }
            _ = fi9
        }
    case "hello":
        var v15 uint = 27
        _ = v15
    case "":
        var v16 int32 = -340
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 594
}

func main() {
    var v1 string = (("foo bar" + ((("a" + (((("\n\t" + "foo bar") + (((("foo bar" + ("hello" + "hello")) + "foo bar") + (("\n\t" + "hello") + (("\n\t" + (("a" + "\n\t") + "a")) + "a"))) + (("foo bar" + "a") + "foo bar"))) + "hello") + "\n\t")) + ("\n\t" + "foo bar")) + ((("\n\t" + "a") + (("\n\t" + ("hello" + "a")) + "foo bar")) + "\n\t"))) + "foo bar")
    var v2 bool = (0.1 >= ((0.1 + 1.5) / 1.0))
    v3 := make(map[uint32]string)
    v4 := make(map[string]int)
    fmt.Println(v2, v2)
    v1 = ("a" + "a")
    v2 = (2 < 2)
    for fi17 := 0; fi17 < 13; fi17++ {
        v18 := GenS1{F0: -128}
        v1 = "a"
        _ = fi17
        _ = v18
    }
    _ = len(v4)
    v2 = ((!v2) || true)
    v1 = ""
    fmt.Println(v1, v2)
    v19 := GenS1{F0: 1}
    v20 := []int{}
    var v21 uint = 4294967295
    delete(v4, (v1 + ((v1 + "foo bar") + v1)))
    v21 -= 0
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
