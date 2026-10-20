package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint64 }
type GenS2 struct { F0 int; F1 int8; F2 int }

func main() {
    v1 := []string{("" + ""), ("\n\t" + "hello")}
    var v2 string = ((("\n\t" + "") + "\n\t") + (("a" + "hello") + "hello"))
    var v3 uint = 0
    for fi4 := 0; fi4 < 17; fi4++ {
        v2 = v2
        _ = fi4
    }
    v3 += (v3 >> ((v3) & 63))
    v2 = "hello"
    for fi5 := 0; fi5 < 20; fi5++ {
        v2 = v2
        var v6 float64 = 3.14159
        _ = fi5
        _ = v6
    }
    var v7 int16 = 0
    fmt.Println(v3)
    if (true && false) {
        fmt.Println(v7, v3)
        switch 9223372036854775807 {
        case 423:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case 313:
            v8 := GenS1{F0: 2, F1: 128, F2: 42}
            _ = v8
        default:
            for fi9 := 0; fi9 < 9; fi9++ {
                v1 = append(v1, "hello")
                switch "a" {
                case "\n\t":
                    fmt.Println(v2, fi9, v7)
                case "a":
                    v3 += v3
                default:
                    if (!("hello" < (v2 + ("hello" + v2)))) {
                        switch v2 {
                        case "foo bar":
                            if true {
                                var v10 float64 = 0.0
                                if false {
                                    v11 := GenS1{F0: 773, F1: fi9, F2: 243}
                                    _ = v11
                                } else {
                                    if ((v3 >> (((^(v3))) & 63)) <= v3) {
                                        var v12 bool = ((v2 + "hello") > v2)
                                        _ = len(v1)
                                        _ = v12
                                    }
                                    var v13 bool = ((1e10 != (v10 - (((3.14159 + 1e10) * v10) + v10))) || (false || ((false && true) && ("" > v2))))
                                    if v13 {
                                        var v14 string = v2
                                        _ = v14
                                    }
                                    _ = v13
                                }
                                switch v2 {
                                case "foo bar":
                                    for fi15 := 0; fi15 < 4; fi15++ {
                                        v16 := GenS2{F0: (fi15 - 42), F1: 42, F2: 42}
                                        _ = fi15
                                        _ = v16
                                    }
                                case "a":
                                    var v17 string = ("foo bar" + v2)
                                    _ = v17
                                case "\n\t":
                                    switch (fi9 ^ fi9) {
                                    case 32767:
                                        v1 = append(v1, v2)
                                    case 1:
                                        switch "\n\t" {
                                        case "\n\t":
                                            _ = len(v1)
                                        case "foo bar":
                                            var v18 uint = 1
                                            _ = v18
                                        case "":
                                            v3 -= v3
                                        default:
                                            if len(v1) > 0 {
                                                v1[0] = ("foo bar" + (v2 + v2))
                                            }
                                        }
                                    }
                                default:
                                    v2 = v2
                                }
                                _ = v10
                            }
                        }
                        fmt.Println(v3)
                    } else {
                        fmt.Println(v3, fi9, v3)
                    }
                }
                _ = fi9
            }
        }
    } else {
        var v19 uint16 = 280
        _ = v19
    }
    if (4294967295 >= 4294967295) {
        for fi20 := 0; fi20 < 18; fi20++ {
            var v21 string = ((v2 + v2) + v2)
            v22 := GenS2{F0: fi20, F1: 42, F2: 32768}
            var v23 string = ("\n\t" + v21)
            _ = fi20
            _ = v21
            _ = v22
            _ = v23
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = len(v1)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
}
