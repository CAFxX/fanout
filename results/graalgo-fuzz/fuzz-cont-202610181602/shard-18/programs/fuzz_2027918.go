package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 uint8 }
func genFn0(p0 bool, p1 int8) uint8 {
    var v6 int8 = 0
    v7 := GenS1{F0: 766}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 128
}

func main() {
    v1 := []string{("" + ((("\n\t" + (("foo bar" + "foo bar") + "a")) + "hello") + ("foo bar" + "a"))), ("a" + "foo bar"), "foo bar", "\n\t"}
    var v2 uint8 = 1
    var v3 bool = true
    v4 := make(map[uint32]bool)
    v5 := make(map[string]string)
    v4[940] = (((255 >= 783) || (42 <= 0)) && (!(int16(v2) >= -928)))
    var v8 int64 = -404
    switch 32768 {
    case -1:
        for fi9 := 0; fi9 < 13; fi9++ {
            _ = len(v4)
            switch ("foo bar" + "a") {
            case "":
                v2 -= genFn0(v3, -20)
            case "hello":
                switch (fi9 | -2147483648) {
                case 1:
                    v10 := GenS2{F0: 121}
                    _ = v10
                case 128:
                    v11 := v4[2]
                    _ = v11
                }
            case "foo bar":
                for fi12 := 0; fi12 < 1; fi12++ {
                    var v13 uint = 65535
                    _ = fi12
                    _ = v13
                }
            default:
                v1 = append(v1, (("hello" + "hello") + "a"))
            }
            var v14 string = "foo bar"
            _ = fi9
            _ = v14
        }
    case 42:
        switch "hello" {
        case "foo bar":
            var v15 int8 = -128
            _ = v15
        case "":
            fmt.Println(v8)
        case "\n\t":
            v16 := v4[4294967295]
            _ = v16
        }
    case 128:
        var v17 uint64 = 2
        _ = v17
    }
    _ = len(v4)
    v18 := v5["foo bar"]
    v18 = ((("hello" + (v18 + "\n\t")) + "foo bar") + v18)
    fmt.Println(v18, v18, v8)
    v19 := GenS3{F0: 94, F1: (v2 / ((v2) | 1))}
    var v20 int16 = 0
    v20 -= -1
    delete(v5, ("\n\t" + v18))
    var v21 int8 = 1
    v22 := v4[42]
    fmt.Println(v21, v18, v18)
    var v23 int64 = (v8 / ((v8) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
