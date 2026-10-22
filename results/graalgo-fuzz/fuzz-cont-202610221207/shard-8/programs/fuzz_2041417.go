package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{}
    var v3 uint16 = 1
    var v4 string = ("" + "hello")
    fmt.Println(v4)
    fmt.Println(v4)
    var v5 int = -419
    var v6 float64 = (1.5 - 1e10)
    var v7 uint8 = 254
    for fi8 := 0; fi8 < 4; fi8++ {
        v7 += v7
        _ = fi8
    }
    v9 := GenS1{F0: 4294967295, F1: 57}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        for fi11 := 0; fi11 < 5; fi11++ {
            v3 = v3
            switch v4 {
            case "a":
                switch (((((v4 + v4) + v4) + v4) + "hello") + "") {
                case "hello":
                    v12 := GenS1{F0: 678, F1: 1}
                    _ = v12
                case "\n\t":
                    v7 = v7
                case "foo bar":
                    if false {
                        fmt.Println(v4)
                        v7 -= (v7 >> ((18446744073709551615) & 63))
                        v3 = v3
                    }
                }
                if false {
                    switch (("\n\t" + (("foo bar" + (v4 + "a")) + "a")) + (v4 + "\n\t")) {
                    case "":
                        var v13 float32 = 1e10
                        _ = v13
                    case "hello":
                        if (((true || false) && (!(9223372036854775807 <= -621))) || (47 > 1)) {
                            if true {
                                v6 = v6
                                var v14 uint = 2
                                _ = v14
                            } else {
                                if len(v2) > 0 {
                                    v2[0] = v4
                                }
                                var v15 string = v4
                                _ = v15
                            }
                        }
                    default:
                        fmt.Println(v3, v6, fi11)
                    }
                }
                v5 = 4294967295
            }
            _ = fi11
        }
        _ = fi10
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
