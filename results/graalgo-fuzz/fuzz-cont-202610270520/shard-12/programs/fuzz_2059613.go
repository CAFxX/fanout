package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 uint32 }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []string{("foo bar" + "hello"), (("a" + "a") + "a"), ("\n\t" + ("foo bar" + "")), ("\n\t" + "a")}
    v3 := []string{"", (("hello" + (("" + "\n\t") + ("a" + ""))) + "hello"), ""}
    v1 -= 904
    switch "" {
    case "\n\t":
        v3 = append(v3, ("hello" + "\n\t"))
    case "a":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "":
        var v4 float32 = (3.14159 - 0.0)
        _ = v4
    default:
        if (32767 >= 635) {
            v5 := GenS1{F0: ((1.5 + 0.1) <= 3.14159), F1: 4294967295, F2: ("" + "hello")}
            _ = v5
        } else {
            for fi6 := 0; fi6 < 19; fi6++ {
                var v7 uint64 = (v1 & 0)
                _ = fi6
                _ = v7
            }
            for fi8 := 0; fi8 < 17; fi8++ {
                for fi9 := 0; fi9 < 11; fi9++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    fmt.Println(fi9, fi8)
                    var v10 int32 = 42
                    _ = fi9
                    _ = v10
                }
                if (42 <= 4294967295) {
                    var v11 int16 = -16
                    for fi12 := 0; fi12 < 13; fi12++ {
                        var v13 uint16 = 32767
                        v1 = v1
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi12
                        _ = v13
                    }
                    var v14 float32 = 3.14159
                    _ = v11
                    _ = v14
                }
                _ = fi8
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    }
    v15 := GenS2{F0: ("hello" + "foo bar")}
    fmt.Println(v1)
    v1 = 18446744073709551615
    var v16 int = 42
    if len(v2) > 0 {
        v2[0] = "a"
    }
    var v17 uint8 = 42
    if false {
        v2 = append(v2, "\n\t")
        var v18 int = (v16 & 2147483647)
        _ = v18
    }
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    var v19 string = "\n\t"
    switch "a" {
    case "\n\t":
        var v20 bool = true
        _ = v20
    case "a":
        fmt.Println(v17)
    default:
        fmt.Println(v17, v16)
    }
    v21 := GenS3{F0: 0, F1: 219}
    v19 = "a"
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi22 := 0; fi22 < 8; fi22++ {
        v1 += 674
        for fi23 := 0; fi23 < 17; fi23++ {
            v3 = append(v3, "hello")
            _ = fi23
        }
        switch "foo bar" {
        case "a":
            if (65535 <= 761) {
                var v24 uint8 = v17
                if len(v2) > 0 {
                    v2[0] = (("hello" + v19) + v19)
                }
                _ = v24
            }
        case "hello":
            v16 += fi22
        case "\n\t":
            for fi25 := 0; fi25 < 2; fi25++ {
                fmt.Println(v16)
                fmt.Println(v17)
                _ = fi25
            }
        default:
            v26 := GenS3{F0: 42, F1: 2}
            _ = v26
        }
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v21
    _ = v3
}
