package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 int16; F2 string }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 uint; F1 uint }

func main() {
    var v1 uint32 = 42
    var v2 int8 = 100
    var v3 uint16 = 610
    var v4 uint32 = (v1 / ((1) | 1))
    v5 := make(map[uint32]string)
    if true {
        switch (("a" + ("a" + ((("" + ("a" + "")) + "foo bar") + ("" + "foo bar")))) + ("" + ("hello" + "a"))) {
        case "hello":
            switch "\n\t" {
            case "hello":
                if false {
                    switch ("foo bar" + ("" + ("\n\t" + ""))) {
                    case "a":
                        switch ("" + "hello") {
                        case "":
                            v1 = (v1 - 65535)
                        }
                    }
                } else {
                    v4 = (v4 % ((432) | 1))
                }
            default:
                v1 += 0
            }
        case "\n\t":
            v6 := GenS1{F0: (v2 % ((v2) | 1))}
            _ = v6
        default:
            for fi7 := 0; fi7 < 4; fi7++ {
                var v8 uint = 2
                v3 += 1
                fmt.Println(v1)
                _ = fi7
                _ = v8
            }
        }
        fmt.Println(v1, v2, v3)
        v9 := GenS2{F0: "foo bar", F1: 583, F2: ("foo bar" + ((("\n\t" + ("hello" + (("a" + "a") + ""))) + "hello") + "\n\t"))}
        _ = v9
    } else {
        v10 := GenS3{F0: -1, F1: (true || (false || (255 != 42)))}
        for fi11 := 0; fi11 < 7; fi11++ {
            fmt.Println(v2, fi11, v4)
            _ = fi11
        }
        v5[v4] = "a"
        _ = v10
    }
    v12 := []int{2147483647, -9223372036854775808, 42, 611}
    v4 = (v1 + (^(v1)))
    var v13 float64 = 0.1
    v14 := GenS4{F0: 798, F1: 255}
    v15 := GenS2{F0: ((((("hello" + "") + "foo bar") + "a") + ("\n\t" + "a")) + (((("foo bar" + "hello") + "hello") + (("" + ("" + ("foo bar" + "hello"))) + ("a" + "hello"))) + "\n\t")), F1: 0, F2: "\n\t"}
    switch -808 {
    case 32767:
        for fi16 := 0; fi16 < 11; fi16++ {
            switch ((("hello" + "\n\t") + "\n\t") + "hello") {
            case "a":
                if (true || (66 < -637)) {
                    var v17 uint64 = 253
                    var v18 int8 = (v2 + (v2 + v2))
                    _ = v17
                    _ = v18
                }
            default:
                _ = len(v5)
            }
            _ = fi16
        }
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        v20 := GenS4{F0: 65535, F1: 584}
        _ = fi19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
