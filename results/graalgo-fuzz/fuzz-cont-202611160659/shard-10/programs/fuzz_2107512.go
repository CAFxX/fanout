package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 int8; F2 uint16 }

func main() {
    v1 := []string{("\n\t" + "foo bar"), ""}
    v2 := make(map[int]uint32)
    var v3 int16 = 32767
    v4 := []string{"a"}
    switch 2147483647 {
    case 42:
        v5 := GenS1{F0: 65535}
        _ = v5
    case 170:
        _ = len(v1)
    case 2147483647:
        var v6 uint32 = 82
        _ = v6
    }
    fmt.Println(v3)
    v3 = (v3 & v3)
    if len(v1) > 0 {
        v1[0] = ""
    }
    fmt.Println(v3)
    v3 = (^(v3))
    fmt.Println(v3)
    var v7 int = 127
    delete(v2, (-(v7)))
    var v8 bool = false
    if (("a" + "foo bar") != ((("a" + (("" + "hello") + "")) + (("foo bar" + "foo bar") + (((((("hello" + (("" + "") + "a")) + ("a" + "foo bar")) + (("foo bar" + "") + "a")) + "") + "\n\t") + "a"))) + "")) {
        var v9 uint32 = 618
        _ = v9
    }
    var v10 uint = 42
    fmt.Println(v8, v7, v7)
    if false {
        if (32767 == 1) {
            for fi11 := 0; fi11 < 10; fi11++ {
                v12 := v2[fi11]
                _ = fi11
                _ = v12
            }
        }
    }
    switch "hello" {
    case "\n\t":
        for fi13 := 0; fi13 < 2; fi13++ {
            v7 = 32767
            _ = fi13
        }
    }
    v8 = v8
    if true {
        var v14 bool = false
        _ = v14
    } else {
        for fi15 := 0; fi15 < 3; fi15++ {
            if (!true) {
                switch 32768 {
                case 1:
                    if len(v4) > 0 {
                        v4[0] = "foo bar"
                    }
                }
                var v16 uint16 = 0
                _ = v16
            } else {
                for fi17 := 0; fi17 < 5; fi17++ {
                    switch "hello" {
                    case "foo bar":
                        if len(v4) > 0 {
                            v4[0] = ("" + ("a" + "a"))
                        }
                    }
                    var v18 float64 = 1e10
                    v10 = v10
                    fmt.Println(v18, v10, fi15)
                    v7 += v7
                    _ = fi17
                    _ = v18
                }
                v19 := GenS2{F0: "foo bar", F1: 102, F2: 808}
                _ = v19
            }
            _ = fi15
        }
        v8 = (823 == 2)
        fmt.Println(v7)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
