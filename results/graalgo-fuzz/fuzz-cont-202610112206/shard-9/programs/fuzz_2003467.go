package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int8 }

func main() {
    var v1 int = 32768
    var v2 int = 384
    v3 := []uint32{830}
    var v4 int = -1
    v5 := make(map[uint32]bool)
    v4 -= (^(v1))
    v1 = -1
    fmt.Println(v2, v1)
    if true {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if (!(-128 > 100)) {
            v2 += 32768
            v2 += (-(v2))
        }
        v1 = -567
    } else {
        fmt.Println(v2, v2, v1)
    }
    v5[150] = true
    for fi6 := 0; fi6 < 7; fi6++ {
        if true {
            switch ("foo bar" + "") {
            case "\n\t":
                switch "\n\t" {
                case "hello":
                    v7 := GenS1{F0: 42, F1: fi6, F2: -1}
                    _ = v7
                case "foo bar":
                    if (341 > -2147483648) {
                        v1 = 2147483647
                        v4 += -333
                        var v8 int16 = -175
                        _ = v8
                    }
                case "a":
                    var v9 float32 = -2.25
                    _ = v9
                }
            case "foo bar":
                for fi10 := 0; fi10 < 19; fi10++ {
                    v2 -= 9223372036854775807
                    switch ("a" + (("a" + "foo bar") + ("a" + ("foo bar" + ((((("hello" + "") + "a") + "") + "a") + "a"))))) {
                    case "a":
                        if false {
                            var v11 uint32 = 644
                            _ = v11
                        } else {
                            v12 := GenS1{F0: 65535, F1: v2, F2: -17}
                            switch "" {
                            case "hello":
                                v2 = -938
                            case "foo bar":
                                if len(v3) > 0 {
                                    v3[0] = 865
                                }
                            }
                            _ = v12
                        }
                    case "hello":
                        if (!true) {
                            var v13 int = (-(v4))
                            if len(v3) > 0 {
                                v3[0] = 891
                            }
                            _ = v13
                        }
                    case "foo bar":
                        switch (v1 >> ((2) & 63)) {
                        case 2147483647:
                            var v14 uint64 = 2
                            _ = v14
                        case 691:
                            v2 -= v1
                        }
                    }
                    _ = fi10
                }
            }
            if true {
                var v15 float32 = 1e10
                _ = v15
            }
            v16 := GenS1{F0: 0, F1: v2, F2: -96}
            _ = v16
        }
        var v17 int = (v2 ^ v4)
        _ = fi6
        _ = v17
    }
    var v18 string = "a"
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
