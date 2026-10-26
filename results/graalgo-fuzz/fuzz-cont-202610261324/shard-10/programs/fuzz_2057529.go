package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint32 }

func main() {
    var v1 uint8 = 128
    v2 := []string{"a"}
    var v3 float32 = 0.1
    var v4 float32 = v3
    var v5 int64 = 1
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := GenS1{F0: "hello", F1: 42}
        _ = fi6
        _ = v7
    }
    var v8 uint32 = 0
    var v9 int64 = (-(v5))
    var v10 int16 = -32768
    if false {
        switch ("foo bar" + "hello") {
        case "a":
            for fi11 := 0; fi11 < 9; fi11++ {
                v2 = append(v2, "hello")
                switch fi11 {
                case -1:
                    switch ("a" + "foo bar") {
                    case "a":
                        var v12 int64 = 482
                        _ = v12
                    case "":
                        v1 -= (v1 - 255)
                    case "hello":
                        var v13 int = fi11
                        _ = v13
                    }
                case 32768:
                    v9 -= (v9 + v9)
                case 128:
                    switch 4294967295 {
                    case -9223372036854775808:
                        _ = len(v2)
                    }
                default:
                    v14 := GenS1{F0: "a", F1: 2}
                    _ = v14
                }
                var v15 bool = (475 < v8)
                _ = fi11
                _ = v15
            }
        case "":
            for fi16 := 0; fi16 < 17; fi16++ {
                var v17 int8 = 42
                _ = len(v2)
                fmt.Println(v9, v9)
                _ = fi16
                _ = v17
            }
        case "\n\t":
            v18 := GenS2{F0: (v8 >> ((42) & 63)), F1: v10, F2: 338}
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        if true {
            switch "foo bar" {
            case "":
                var v20 uint = 412
                _ = v20
            case "hello":
                for fi21 := 0; fi21 < 8; fi21++ {
                    var v22 int16 = -940
                    _ = fi21
                    _ = v22
                }
            default:
                fmt.Println(v4, v5, v9)
            }
            for fi23 := 0; fi23 < 11; fi23++ {
                v4 = 0.1
                _ = fi23
            }
            v8 += (v8 * v8)
        }
        _ = fi19
    }
    _ = len(v2)
    if (786 < 42) {
        for fi24 := 0; fi24 < 20; fi24++ {
            fmt.Println(v8, v10)
            var v25 bool = true
            v8 -= 65535
            _ = fi24
            _ = v25
        }
    } else {
        if true {
            switch "hello" {
            case "a":
                v2 = append(v2, "a")
            case "foo bar":
                v3 = (1.5 / ((float32(float64((1e10 - v3))) - 0.1) - v4))
            }
            fmt.Println(v10)
        } else {
            v5 = v5
            if false {
                v4 = v4
                for fi26 := 0; fi26 < 20; fi26++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if len(v2) > 0 {
                        v2[0] = ("foo bar" + "")
                    }
                    _ = fi26
                }
                for fi27 := 0; fi27 < 11; fi27++ {
                    var v28 int16 = 32767
                    v29 := GenS1{F0: "hello", F1: 2}
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            } else {
                v1 += (v1 ^ v1)
                _ = len(v2)
                if false {
                    if (604 != 42) {
                        v1 += 223
                    }
                }
            }
            v30 := GenS1{F0: "", F1: 18446744073709551615}
            _ = v30
        }
    }
    if len(v2) > 0 {
        v2[0] = ("\n\t" + "\n\t")
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
