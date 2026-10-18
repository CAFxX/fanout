package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 uint = 418
    var v2 string = ""
    v3 := make(map[string]int)
    switch "\n\t" {
    case "a":
        v2 = ("\n\t" + v2)
    case "hello":
        fmt.Println(v1, v2)
    default:
        fmt.Println(v2)
    }
    v4 := []int{32768}
    for fi5 := 0; fi5 < 19; fi5++ {
        switch "a" {
        case "":
            var v6 uint64 = 4294967295
            _ = v6
        default:
            var v7 bool = (!true)
            _ = v7
        }
        if true {
            v1 -= (v1 | v1)
            fmt.Println(v2, fi5, fi5)
        }
        if true {
            v1 += 92
            v8 := GenS1{F0: 1}
            if true {
                v2 = (v2 + "")
                for fi9 := 0; fi9 < 5; fi9++ {
                    v10 := GenS1{F0: 638}
                    if true {
                        v2 = "foo bar"
                        var v11 float32 = ((3.14159 - (-2.25 + (1.5 / float32(0.1)))) / 1.5)
                        v1 = v1
                        v2 = "foo bar"
                        var v12 string = "foo bar"
                        _ = v11
                        _ = v12
                    } else {
                        var v13 uint8 = 2
                        _ = v13
                    }
                    _ = fi9
                    _ = v10
                }
                var v14 uint16 = 255
                _ = v14
            }
            _ = v8
        }
        _ = fi5
    }
    var v15 int32 = -560
    v16 := GenS1{F0: -1}
    v17 := GenS1{F0: 431}
    v15 += v15
    v15 -= -2147483648
    v15 = (v15 & 42)
    switch (v2 + v2) {
    case "hello":
        v15 = 1
    default:
        v2 = v2
    }
    v18 := GenS1{F0: 0}
    for fi19 := 0; fi19 < 1; fi19++ {
        _ = len(v3)
        var v20 uint32 = 4294967295
        if false {
            for fi21 := 0; fi21 < 17; fi21++ {
                switch fi21 {
                case -9223372036854775808:
                    var v22 string = v2
                    _ = v22
                case 128:
                    fmt.Println(v15)
                case 2147483647:
                    v15 -= v15
                default:
                    v2 = v2
                }
                switch (fi21 + -591) {
                case -9223372036854775808:
                    v23 := GenS1{F0: 1}
                    v24 := GenS1{F0: -9223372036854775808}
                    v15 += 0
                    v4 = append(v4, (fi21 >> (((v1 << ((v1) & 63))) & 63)))
                    if ((825 != (fi19 + (fi19 / (((-(fi21))) | 1)))) && (!(!(771 > v1)))) {
                        fmt.Println(fi21, fi21)
                        v15 = v15
                    }
                    _ = v23
                    _ = v24
                case 32767:
                    var v25 bool = (1.5 == 0.0)
                    _ = v25
                case -82:
                    switch 4294967295 {
                    case 1:
                        v26 := GenS1{F0: 42}
                        _ = v26
                    }
                default:
                    if true {
                        var v27 uint32 = (v20 * 418)
                        _ = v27
                    }
                }
                switch v2 {
                case "a":
                    v28 := GenS1{F0: 461}
                    _ = v28
                default:
                    fmt.Println(fi19)
                }
                _ = fi21
            }
            var v29 float32 = 1.5
            v30 := GenS1{F0: -9223372036854775808}
            _ = v29
            _ = v30
        }
        _ = fi19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
