package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 uint8 }

func main() {
    var v1 int8 = 42
    var v2 string = "foo bar"
    v3 := make(map[string]uint32)
    v4 := make(map[string]uint32)
    var v5 uint64 = 2
    v6 := []int{-517, 127, 42, 1}
    for fi7 := 0; fi7 < 16; fi7++ {
        if true {
            var v8 uint8 = 0
            _ = v8
        }
        _ = fi7
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        if true {
            switch (v2 + ("a" + "hello")) {
            case "\n\t":
                fmt.Println(v1, fi9)
            }
            v2 = v2
            _ = len(v6)
        } else {
            var v10 uint64 = (v5 >> ((0) & 63))
            _ = v10
        }
        _ = fi9
    }
    v11 := GenS1{F0: (true && (0.1 >= 0.0)), F1: 0}
    var v12 int8 = (v1 >> ((181) & 63))
    var v13 uint64 = 4294967295
    v2 = "hello"
    v14 := GenS1{F0: (!true), F1: 127}
    var v15 uint32 = 2147483648
    v5 += 1
    for fi16 := 0; fi16 < 11; fi16++ {
        var v17 bool = true
        fmt.Println(v13)
        for fi18 := 0; fi18 < 3; fi18++ {
            if v17 {
                delete(v4, "")
                if (0.1 != 1e10) {
                    v15 += (v15 % ((80) | 1))
                    switch v2 {
                    case "foo bar":
                        v12 += 0
                    case "":
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                } else {
                    switch "a" {
                    case "foo bar":
                        switch "a" {
                        case "foo bar":
                            switch ("foo bar" + ((v2 + "\n\t") + "")) {
                            case "\n\t":
                                v5 = 418
                                var v19 bool = (!v17)
                                var v20 uint64 = (^(v5))
                                _ = v19
                                _ = v20
                            case "foo bar":
                                if false {
                                    fmt.Println(v13)
                                } else {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    _ = len(v6)
                                }
                            case "a":
                                v1 = v12
                            }
                        }
                    default:
                        var v21 bool = ((true && v17) || v17)
                        _ = v21
                    }
                    v22 := GenS1{F0: (((!v17) || ((v2 + ("\n\t" + (v2 + v2))) == v2)) || (42 >= 302)), F1: 128}
                    _ = v22
                }
                v15 = v15
            } else {
                v23 := GenS2{F0: 521}
                v24 := GenS3{F0: (v12 >> ((361) & 63)), F1: 0}
                _ = v23
                _ = v24
            }
            _ = fi18
        }
        _ = fi16
        _ = v17
    }
    v25 := GenS1{F0: (v13 == (v13 * 42)), F1: 148}
    v26 := GenS4{F0: 38}
    v6 = append(v6, 277)
    var v27 uint = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
