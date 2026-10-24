package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS3 struct { F0 int; F1 int; F2 uint16 }
type GenS4 struct { F0 int64 }

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 127
    v3 := []string{(("" + "foo bar") + "a"), "foo bar", (("foo bar" + "\n\t") + (("foo bar" + "foo bar") + ("\n\t" + "hello")))}
    var v4 uint8 = 50
    _ = len(v1)
    if true {
        var v5 int32 = 2147483647
        var v6 int8 = -103
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        v2 = (^(v2))
        if true {
            for fi8 := 0; fi8 < 19; fi8++ {
                v4 += 238
                _ = fi8
            }
        } else {
            v2 = v2
            var v9 uint = 255
            _ = v9
        }
        _ = fi7
    }
    var v10 float64 = 0.1
    _ = len(v1)
    var v11 int8 = v2
    for fi12 := 0; fi12 < 10; fi12++ {
        v11 -= (v11 << ((547) & 63))
        _ = fi12
    }
    v13 := GenS1{F0: 193, F1: -1, F2: 18446744073709551615}
    if false {
        v4 = v4
        delete(v1, (("a" + "\n\t") + "hello"))
    } else {
        if true {
            if (42 >= 18446744073709551615) {
                _ = len(v3)
                var v14 int32 = 1
                _ = v14
            }
            for fi15 := 0; fi15 < 12; fi15++ {
                v16 := v1["a"]
                if (32767 >= -719) {
                    for fi17 := 0; fi17 < 15; fi17++ {
                        v4 += (^(v4))
                        _ = fi17
                    }
                    if true {
                        v18 := GenS2{F0: 1, F1: -863, F2: 32767}
                        v1["a"] = (fi15 | fi15)
                        _ = v18
                    } else {
                        var v19 int8 = v2
                        if (!false) {
                            if (false && false) {
                                var v20 int = v16
                                _ = v20
                            } else {
                                for fi21 := 0; fi21 < 4; fi21++ {
                                    switch "hello" {
                                    case "":
                                        v1["a"] = (fi15 + 32768)
                                    }
                                    if (false || false) {
                                        v16 += fi21
                                        var v22 int64 = 42
                                        _ = len(v3)
                                        _ = v22
                                    }
                                    v11 += v19
                                    var v23 float32 = 1e10
                                    _ = fi21
                                    _ = v23
                                }
                                switch "foo bar" {
                                case "hello":
                                    var v24 int32 = 0
                                    _ = v24
                                case "a":
                                    var v25 float32 = 1.5
                                    _ = v25
                                case "":
                                    var v26 int = -2147483648
                                    _ = v26
                                default:
                                    for fi27 := 0; fi27 < 2; fi27++ {
                                        var v28 float64 = v10
                                        v16 = fi27
                                        fmt.Println(v16)
                                        fmt.Println(v10)
                                        _ = fi27
                                        _ = v28
                                    }
                                }
                            }
                        }
                        if false {
                            v29 := GenS3{F0: (v16 >> ((255) & 63)), F1: (fi15 - fi15), F2: 708}
                            delete(v1, "hello")
                            fmt.Println(v2, v10, v16)
                            _ = v29
                        }
                        _ = v19
                    }
                } else {
                    v4 += v4
                    if true {
                        v2 = v2
                        v30 := GenS4{F0: -1}
                        _ = v30
                    } else {
                        v10 = -2.25
                        fmt.Println(v4, v2)
                    }
                    switch "foo bar" {
                    case "foo bar":
                        v31 := GenS2{F0: 2147483647, F1: 937, F2: 920}
                        _ = v31
                    case "\n\t":
                        fmt.Println(v11, v16)
                    case "":
                        v4 = v4
                    }
                }
                delete(v1, "hello")
                _ = fi15
                _ = v16
            }
        }
        fmt.Println(v2, v2)
    }
    if (!true) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        v11 -= 100
        fmt.Println(v4)
    }
    v2 = (v2 | v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
