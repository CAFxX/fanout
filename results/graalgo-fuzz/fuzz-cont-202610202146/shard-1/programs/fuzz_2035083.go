package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int16 }
type GenS2 struct { F0 uint16; F1 string; F2 int }

func main() {
    v1 := []uint32{1, 0, 42, 861}
    var v2 string = ("" + "foo bar")
    v3 := []int{-454, 127}
    for fi4 := 0; fi4 < 8; fi4++ {
        switch ((v2 + "\n\t") + ("\n\t" + v2)) {
        case "hello":
            if false {
                v2 = ("foo bar" + "\n\t")
            }
        case "foo bar":
            fmt.Println(fi4)
        case "a":
            v5 := GenS1{F0: 42, F1: 2147483647, F2: 1}
            _ = v5
        }
        if (18446744073709551615 < 738) {
            v6 := GenS1{F0: 794, F1: fi4, F2: -886}
            if len(v3) > 0 {
                v3[0] = 436
            }
            _ = v6
        } else {
            v7 := GenS2{F0: 42, F1: v2, F2: -667}
            fmt.Println(fi4, v2)
            switch v2 {
            case "\n\t":
                if true {
                    var v8 float32 = 3.14159
                    var v9 string = ("\n\t" + "\n\t")
                    fmt.Println(fi4, v8, v9)
                    _ = v8
                    _ = v9
                }
            }
            _ = v7
        }
        _ = fi4
    }
    _ = len(v3)
    _ = len(v1)
    v1 = append(v1, 4294967295)
    v2 = (v2 + (v2 + ("a" + (v2 + "foo bar"))))
    v10 := GenS2{F0: 42, F1: "hello", F2: 0}
    fmt.Println(v2)
    v11 := GenS2{F0: 303, F1: (v2 + v2), F2: 127}
    var v12 uint64 = 440
    if true {
        v13 := GenS2{F0: 578, F1: "hello", F2: -9223372036854775808}
        switch 4294967295 {
        case 4294967295:
            if (-1 != -9223372036854775808) {
                if len(v1) > 0 {
                    v1[0] = 42
                }
                switch 9223372036854775807 {
                case 2147483647:
                    v12 -= (^(v12))
                default:
                    v2 = v2
                }
            }
        case 42:
            if false {
                var v14 uint8 = 0
                if (-643 == -269) {
                    if false {
                        var v15 uint32 = 115
                        var v16 string = v2
                        for fi17 := 0; fi17 < 14; fi17++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi17
                        }
                        _ = v15
                        _ = v16
                    } else {
                        var v18 uint = 0
                        if (false && (2 < 480)) {
                            if len(v1) > 0 {
                                v1[0] = 42
                            }
                        }
                        fmt.Println(v12, v18, v12)
                        _ = v18
                    }
                    for fi19 := 0; fi19 < 3; fi19++ {
                        fmt.Println(v14)
                        fmt.Println(v14, v2, v2)
                        v2 = v2
                        _ = fi19
                    }
                }
                _ = v14
            }
        }
        _ = v13
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
}
