package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int16 }

func main() {
    var v1 bool = true
    var v2 bool = (18446744073709551615 < 0)
    v3 := make(map[uint32]uint32)
    v4 := []int{-1, 127}
    fmt.Println(v1)
    var v5 int64 = 9223372036854775807
    if true {
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v5, v1, v2)
        case "":
            if v1 {
                switch ("a" + "a") {
                case "\n\t":
                    v5 = v5
                case "foo bar":
                    var v6 uint32 = 0
                    _ = v6
                }
            }
        default:
            switch "" {
            case "a":
                if (!false) {
                    for fi7 := 0; fi7 < 17; fi7++ {
                        if false {
                            fmt.Println(v1, v2, v5)
                            v2 = v2
                            var v8 float64 = 1e10
                            _ = v8
                        } else {
                            v1 = v2
                        }
                        _ = fi7
                    }
                    for fi9 := 0; fi9 < 17; fi9++ {
                        var v10 uint8 = 2
                        var v11 float64 = (0.1 / 1.0)
                        _ = len(v4)
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                    fmt.Println(v5)
                } else {
                    for fi12 := 0; fi12 < 3; fi12++ {
                        var v13 int8 = 0
                        for fi14 := 0; fi14 < 4; fi14++ {
                            var v15 uint = 4294967295
                            _ = fi14
                            _ = v15
                        }
                        _ = fi12
                        _ = v13
                    }
                    if true {
                        var v16 int64 = (v5 << ((348) & 63))
                        _ = v16
                    }
                }
            }
        }
        v1 = true
    } else {
        for fi17 := 0; fi17 < 10; fi17++ {
            var v18 float64 = 3.14159
            _ = len(v4)
            _ = fi17
            _ = v18
        }
        fmt.Println(v5, v5, v2)
    }
    v2 = (42 != 0)
    v19 := GenS1{F0: "foo bar", F1: (("a" + "foo bar") + (("" + "") + ("" + "hello"))), F2: 1}
    var v20 int16 = 32767
    v21 := GenS1{F0: ((("\n\t" + (("a" + "hello") + "a")) + ((("\n\t" + "a") + "hello") + "a")) + "foo bar"), F1: "foo bar", F2: v20}
    var v22 int8 = -1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
