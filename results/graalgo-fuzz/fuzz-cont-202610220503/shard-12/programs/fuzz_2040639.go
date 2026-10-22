package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 int16 = -274
    var v2 int8 = 100
    var v3 uint = 2
    var v4 bool = (2147483647 <= uint32(v3))
    for fi5 := 0; fi5 < 1; fi5++ {
        v6 := []int{fi5, (fi5 - -852), 618}
        switch "a" {
        case "a":
            var v7 bool = v4
            _ = v7
        }
        _ = fi5
        _ = v6
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 bool = false
        for fi10 := 0; fi10 < 13; fi10++ {
            v11 := []int{fi10, (fi8 << ((112) & 63)), fi10, 738}
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    switch "a" {
    case "\n\t":
        var v12 uint = (^(v3))
        _ = v12
    case "":
        v13 := GenS1{F0: 18446744073709551615}
        _ = v13
    case "foo bar":
        var v14 uint8 = 1
        _ = v14
    }
    v4 = (v4 || v4)
    v15 := []int{-659}
    if false {
        var v16 float64 = 1.5
        if (!(!false)) {
            if v4 {
                fmt.Println(v2, v1, v4)
                switch ("foo bar" + "\n\t") {
                case "\n\t":
                    if len(v15) > 0 {
                        v15[0] = 9223372036854775807
                    }
                case "hello":
                    v15 = append(v15, 32768)
                }
            } else {
                for fi17 := 0; fi17 < 9; fi17++ {
                    var v18 uint = (v3 * v3)
                    fmt.Println(v2)
                    _ = fi17
                    _ = v18
                }
                var v19 float64 = v16
                _ = v19
            }
        } else {
            fmt.Println(v2, v4)
            var v20 string = "\n\t"
            _ = v20
        }
        _ = v16
    }
    var v21 uint64 = 42
    v4 = v4
    switch 65 {
    case -1:
        if (!v4) {
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
        }
    default:
        for fi22 := 0; fi22 < 9; fi22++ {
            if v4 {
                switch (-(fi22)) {
                case 127:
                    for fi23 := 0; fi23 < 12; fi23++ {
                        switch (fi23 << ((18446744073709551615) & 63)) {
                        case 42:
                            _ = len(v15)
                        case -2147483648:
                            if (0.1 != (1e10 - 3.14159)) {
                                fmt.Println(v1, v4, v1)
                                var v24 uint = (v3 & 2)
                                var v25 uint64 = 2
                                _ = v24
                                _ = v25
                            } else {
                                if false {
                                    if len(v15) > 0 {
                                        v15[0] = 42
                                    }
                                    var v26 bool = (1 == (v21 % (((^(v21))) | 1)))
                                    var v27 string = "\n\t"
                                    _ = v26
                                    _ = v27
                                }
                            }
                        }
                        var v28 uint16 = 78
                        _ = fi23
                        _ = v28
                    }
                default:
                    v21 = 290
                }
            } else {
                v15 = append(v15, fi22)
            }
            _ = fi22
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
