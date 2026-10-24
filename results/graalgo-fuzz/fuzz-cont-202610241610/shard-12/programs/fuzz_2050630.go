package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }

func main() {
    var v1 string = ("a" + "foo bar")
    var v2 int32 = -618
    var v3 string = v1
    var v4 string = (v3 + "hello")
    var v5 int8 = 1
    if false {
        v6 := []int{}
        _ = v6
    } else {
        var v7 uint64 = 0
        _ = v7
    }
    var v8 int = -9223372036854775808
    v9 := []int{}
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    v3 = v3
    switch v1 {
    case "a":
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    case "\n\t":
        if true {
            v10 := GenS1{F0: 42, F1: 9}
            switch (v8 | -887) {
            case 4294967295:
                v4 = ("" + v4)
            case -708:
                switch "hello" {
                case "foo bar":
                    if false {
                        var v11 int8 = (v5 ^ (^(v5)))
                        var v12 uint = 1
                        for fi13 := 0; fi13 < 2; fi13++ {
                            v9 = append(v9, 4294967295)
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                            v14 := GenS1{F0: 0, F1: (v12 - v12)}
                            _ = fi13
                            _ = v14
                        }
                        _ = v11
                        _ = v12
                    } else {
                        var v15 int64 = 502
                        _ = v15
                    }
                case "\n\t":
                    _ = len(v9)
                }
            }
            _ = v10
        }
    case "hello":
        v16 := GenS1{F0: -942, F1: 65535}
        _ = v16
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 bool = (false && (v5 <= 42))
        _ = fi17
        _ = v18
    }
    var v19 int = v8
    if (3.14159 <= 3.14159) {
        if (v19 != v19) {
            v8 -= v19
        }
        switch v3 {
        case "foo bar":
            if (!false) {
                for fi20 := 0; fi20 < 12; fi20++ {
                    var v21 int16 = -32768
                    _ = fi20
                    _ = v21
                }
                var v22 string = "foo bar"
                _ = v22
            }
        }
        v8 += (v19 + (v19 - (v8 - 151)))
    } else {
        v2 = -944
    }
    for fi23 := 0; fi23 < 18; fi23++ {
        if (((1.5 - (0.0 - float64((1e10 / -2.25)))) + 1.5) != 1.5) {
            v8 += (v19 - (v19 >> ((18446744073709551615) & 63)))
            v3 = ("" + v3)
        } else {
            if false {
                v2 = v2
                v19 = (v19 & (v8 ^ (-(v8))))
                for fi24 := 0; fi24 < 2; fi24++ {
                    v5 = (v5 | v5)
                    v8 += (v19 % ((fi24) | 1))
                    _ = fi24
                }
            }
            var v25 int32 = (^(v2))
            _ = v25
        }
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
