package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := []int{32768, 127}
    v2 := []string{("a" + (("" + "foo bar") + "")), "", ("foo bar" + "hello"), "foo bar"}
    var v3 int32 = -421
    if false {
        v3 += (v3 % (((v3 & v3)) | 1))
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
    } else {
        fmt.Println(v3)
    }
    v3 = v3
    for fi4 := 0; fi4 < 7; fi4++ {
        switch "a" {
        case "foo bar":
            for fi5 := 0; fi5 < 14; fi5++ {
                var v6 int32 = (v3 | -1)
                _ = fi5
                _ = v6
            }
        case "":
            switch "" {
            case "":
                v7 := GenS1{F0: 42}
                _ = v7
            case "hello":
                if false {
                    switch "\n\t" {
                    case "a":
                        for fi8 := 0; fi8 < 20; fi8++ {
                            if len(v1) > 0 {
                                v1[0] = fi8
                            }
                            _ = fi8
                        }
                    case "":
                        var v9 bool = ((false && true) || ((("\n\t" + ("foo bar" + ("hello" + ("a" + ("" + "foo bar"))))) + ("foo bar" + "foo bar")) <= ((((("hello" + "a") + "hello") + "foo bar") + ((("hello" + "foo bar") + "\n\t") + ("\n\t" + ""))) + ((("\n\t" + ("hello" + ("" + ("" + ((("hello" + "") + "\n\t") + "hello"))))) + "") + "a"))))
                        _ = v9
                    }
                } else {
                    if false {
                        if len(v2) > 0 {
                            v2[0] = ("a" + "")
                        }
                        fmt.Println(fi4, v3)
                    }
                    for fi10 := 0; fi10 < 16; fi10++ {
                        switch "a" {
                        case "":
                            v11 := GenS1{F0: 312}
                            _ = v11
                        default:
                            v3 -= (v3 / ((0) | 1))
                        }
                        var v12 string = "\n\t"
                        _ = len(v1)
                        _ = fi10
                        _ = v12
                    }
                }
            default:
                v13 := GenS1{F0: -32768}
                _ = v13
            }
        case "hello":
            var v14 uint64 = 144
            _ = v14
        default:
            switch 2147483647 {
            case -639:
                v3 += v3
            case 0:
                for fi15 := 0; fi15 < 5; fi15++ {
                    var v16 uint8 = 0
                    _ = fi15
                    _ = v16
                }
            case 204:
                _ = len(v2)
            }
        }
        _ = fi4
    }
    fmt.Println(v3)
    for fi17 := 0; fi17 < 4; fi17++ {
        v3 = v3
        switch fi17 {
        case 1:
            fmt.Println(fi17, v3)
        case 32768:
            fmt.Println(v3, fi17)
        case 9223372036854775807:
            v3 = (v3 << ((18446744073709551615) & 63))
        }
        _ = fi17
    }
    var v18 uint16 = 42
    v3 -= v3
    v19 := GenS1{F0: 1000}
    if true {
        var v20 bool = (1e10 != ((-2.25 * 1e10) + 1.5))
        if len(v1) > 0 {
            v1[0] = 459
        }
        fmt.Println(v20, v18)
        _ = v20
    } else {
        var v21 string = "foo bar"
        v21 = v21
        v22 := GenS1{F0: 1}
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 6; fi23++ {
        v18 = (v18 * 0)
        v24 := GenS1{F0: 490}
        v25 := GenS1{F0: 394}
        _ = fi23
        _ = v24
        _ = v25
    }
    var v26 float32 = (float32(0.1) - 0.1)
    for fi27 := 0; fi27 < 18; fi27++ {
        v18 = v18
        v26 += v26
        if (!false) {
            var v28 uint64 = 18446744073709551615
            switch (fi27 | 9223372036854775807) {
            case -117:
                v29 := GenS1{F0: -32768}
                _ = v29
            }
            _ = v28
        }
        _ = fi27
    }
    fmt.Println(v18, v3)
    switch "" {
    case "":
        v18 += (v18 ^ v18)
    case "\n\t":
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v30 float64 = 1.5
    v26 -= (v26 + v26)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v30
}
