package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8; F1 bool; F2 bool }

func main() {
    var v1 string = ("" + ((("hello" + ("\n\t" + "foo bar")) + "") + ""))
    var v2 uint16 = 407
    v3 := make(map[string]string)
    switch "" {
    case "hello":
        v4 := GenS1{F0: 32767}
        _ = v4
    case "a":
        for fi5 := 0; fi5 < 20; fi5++ {
            var v6 uint8 = 36
            _ = fi5
            _ = v6
        }
    case "\n\t":
        v7 := GenS1{F0: (^(v2))}
        _ = v7
    }
    v8 := GenS2{F0: 127, F1: 18446744073709551615, F2: 42}
    for fi9 := 0; fi9 < 9; fi9++ {
        v3[v1] = "foo bar"
        _ = fi9
    }
    v10 := GenS3{F0: -2147483648}
    v1 = v1
    for fi11 := 0; fi11 < 13; fi11++ {
        v1 = "a"
        if true {
            _ = len(v3)
        }
        if (494 >= -9223372036854775808) {
            v1 = v1
            if false {
                var v12 uint64 = 113
                var v13 float32 = 0.1
                _ = v12
                _ = v13
            }
            if false {
                if (100 >= -35) {
                    delete(v3, "foo bar")
                    v14 := GenS1{F0: 42}
                    _ = v14
                } else {
                    if (128 >= 2) {
                        var v15 uint = 563
                        _ = v15
                    }
                    v16 := GenS4{F0: 120}
                    _ = v16
                }
                fmt.Println(fi11, v1)
                var v17 uint32 = 1
                _ = v17
            } else {
                for fi18 := 0; fi18 < 13; fi18++ {
                    var v19 uint64 = 965
                    v2 = (v2 >> ((151) & 63))
                    v20 := GenS5{F0: 2, F1: (!(!((fi11 - (^(fi18))) <= fi18))), F2: ((0.0 + (0.1 + 3.14159)) <= (0.0 + 0.0))}
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                switch v1 {
                case "a":
                    fmt.Println(v2, fi11)
                case "hello":
                    if ((0 < 370) && true) {
                        switch "\n\t" {
                        case "foo bar":
                            v21 := GenS3{F0: -642}
                            _ = v21
                        }
                    } else {
                        fmt.Println(v2, v2)
                        var v22 float32 = 1e10
                        _ = v22
                    }
                }
            }
        }
        _ = fi11
    }
    v2 += v2
    fmt.Println(v1, v2)
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v8
}
