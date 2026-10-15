package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }

func main() {
    var v1 uint16 = 0
    v2 := []int{9223372036854775807, 570}
    v3 := make(map[string]int)
    v4 := []string{"foo bar", ("a" + (("a" + "a") + ("hello" + (("" + "") + ("foo bar" + ("hello" + "hello")))))), ""}
    if (!true) {
        v5 := GenS1{F0: 51}
        v6 := GenS1{F0: 393}
        for fi7 := 0; fi7 < 12; fi7++ {
            fmt.Println(fi7, fi7)
            var v8 float64 = 3.14159
            _ = fi7
            _ = v8
        }
        _ = v5
        _ = v6
    } else {
        v1 = (v1 & v1)
        v1 += (^(v1))
    }
    v1 -= (^(v1))
    switch ((("foo bar" + "hello") + "a") + ("\n\t" + "foo bar")) {
    case "hello":
        switch "a" {
        case "":
            v9 := v3["hello"]
            _ = v9
        }
    case "\n\t":
        for fi10 := 0; fi10 < 19; fi10++ {
            for fi11 := 0; fi11 < 9; fi11++ {
                var v12 string = "foo bar"
                _ = fi11
                _ = v12
            }
            v1 = 255
            if (915 >= 1) {
                v1 = (v1 - (v1 % ((v1) | 1)))
                for fi13 := 0; fi13 < 2; fi13++ {
                    switch fi10 {
                    case -9:
                        if (true || false) {
                            v14 := GenS2{F0: "foo bar"}
                            _ = v14
                        }
                    }
                    v3[""] = 412
                    switch (("hello" + "a") + ("foo bar" + "")) {
                    case "":
                        var v15 float64 = -2.25
                        _ = v15
                    case "hello":
                        var v16 int16 = -1
                        _ = v16
                    default:
                        fmt.Println(fi10, fi10, fi10)
                    }
                    _ = fi13
                }
                var v17 bool = true
                _ = v17
            }
            _ = fi10
        }
    case "":
        var v18 string = "a"
        _ = v18
    }
    var v19 int8 = 127
    v20 := GenS2{F0: ("hello" + "hello")}
    for fi21 := 0; fi21 < 1; fi21++ {
        var v22 uint64 = 18446744073709551615
        _ = fi21
        _ = v22
    }
    if len(v2) > 0 {
        v2[0] = 127
    }
    var v23 int = 128
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
}
