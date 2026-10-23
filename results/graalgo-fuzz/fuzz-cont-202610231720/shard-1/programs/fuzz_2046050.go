package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 bool }

func main() {
    v1 := make(map[int]uint32)
    var v2 int = 128
    var v3 uint8 = 1
    v4 := make(map[int]uint32)
    v5 := []string{("a" + "a"), "a", ("\n\t" + "")}
    var v6 uint8 = v3
    v2 = v2
    switch ("\n\t" + "a") {
    case "\n\t":
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 float32 = 1.5
            _ = fi7
            _ = v8
        }
    case "hello":
        v9 := GenS1{F0: 1}
        _ = v9
    case "foo bar":
        if true {
            v6 += (v3 / (((v3 ^ (v3 - 155))) | 1))
            v6 = v3
            v6 += v3
        }
    }
    v2 += v2
    switch ("foo bar" + "foo bar") {
    case "a":
        for fi10 := 0; fi10 < 4; fi10++ {
            switch (v2 % ((v2) | 1)) {
            case 32767:
                if false {
                    var v11 float64 = 1.5
                    _ = v11
                } else {
                    v12 := GenS2{F0: 42, F1: ((("a" + "") + ("foo bar" + ("\n\t" + "hello"))) <= "hello")}
                    switch fi10 {
                    case -962:
                        if (true || true) {
                            for fi13 := 0; fi13 < 10; fi13++ {
                                var v14 float32 = 0.1
                                _ = fi13
                                _ = v14
                            }
                            var v15 bool = true
                            var v16 string = "hello"
                            _ = v15
                            _ = v16
                        } else {
                            _ = len(v5)
                            v3 -= (v3 - v3)
                        }
                    case 1000:
                        if false {
                            for fi17 := 0; fi17 < 5; fi17++ {
                                var v18 float32 = 0.0
                                fmt.Println(v3, fi17, fi17)
                                if false {
                                    fmt.Println(v3, v6, v6)
                                }
                                _ = fi17
                                _ = v18
                            }
                        } else {
                            if (18446744073709551615 != 42) {
                                for fi19 := 0; fi19 < 19; fi19++ {
                                    var v20 string = (("\n\t" + "foo bar") + "foo bar")
                                    if true {
                                        var v21 int = 9223372036854775807
                                        _ = v21
                                    } else {
                                        if false {
                                            if (!true) {
                                                v22 := v4[fi19]
                                                _ = v22
                                            }
                                            var v23 float32 = 3.14159
                                            if false {
                                                var v24 uint16 = 65535
                                                _ = v24
                                            }
                                            _ = v23
                                        }
                                        v20 = ""
                                        var v25 string = (v20 + ((v20 + (("foo bar" + "") + v20)) + v20))
                                        _ = v25
                                    }
                                    var v26 uint16 = 1
                                    _ = fi19
                                    _ = v20
                                    _ = v26
                                }
                            }
                        }
                    case 274:
                        v3 = (v3 ^ (v3 | 0))
                    default:
                        fmt.Println(fi10)
                    }
                    _ = v12
                }
            case 2147483647:
                if true {
                    switch "a" {
                    case "hello":
                        fmt.Println(v6, fi10, v2)
                    case "":
                        fmt.Println(fi10)
                    }
                    for fi27 := 0; fi27 < 20; fi27++ {
                        var v28 uint8 = 0
                        _ = fi27
                        _ = v28
                    }
                }
            default:
                v29 := GenS1{F0: 42}
                _ = v29
            }
            _ = fi10
        }
    case "":
        for fi30 := 0; fi30 < 3; fi30++ {
            fmt.Println(fi30, v2)
            _ = fi30
        }
    }
    fmt.Println(v2)
    v1[v2] = 4294967295
    var v31 int32 = 0
    var v32 int = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
