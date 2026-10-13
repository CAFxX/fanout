package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string }

func main() {
    v1 := []uint32{0, 2147483647, 4294967295, 0}
    var v2 uint64 = 1
    var v3 int16 = 42
    v4 := make(map[uint32]int)
    var v5 float64 = 3.14159
    v2 = (^(v2))
    v2 -= (v2 * 599)
    var v6 uint = 629
    switch 32767 {
    case 9223372036854775807:
        switch ("\n\t" + ((("hello" + "a") + "foo bar") + "hello")) {
        case "\n\t":
            v7 := GenS1{F0: (v6 >> ((v6) & 63)), F1: (v6 & v6)}
            _ = v7
        }
    case 1:
        for fi8 := 0; fi8 < 4; fi8++ {
            fmt.Println(v2, v6, fi8)
            v1 = append(v1, 42)
            var v9 float32 = 1.5
            _ = fi8
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        if true {
            var v11 string = ("\n\t" + ("foo bar" + ("" + "\n\t")))
            _ = v11
        }
        if ("hello" != "") {
            if (v2 == 4294967295) {
                var v12 int8 = int8(fi10)
                fmt.Println(v2)
                _ = v12
            } else {
                for fi13 := 0; fi13 < 2; fi13++ {
                    v2 += (^(v2))
                    switch "\n\t" {
                    case "a":
                        v1 = append(v1, 2147483647)
                    case "":
                        fmt.Println(v3, fi10, v6)
                    default:
                        _ = len(v1)
                    }
                    _ = fi13
                }
                if (-1 != 2147483647) {
                    if true {
                        for fi14 := 0; fi14 < 18; fi14++ {
                            var v15 uint8 = 152
                            var v16 uint = v6
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                    } else {
                        for fi17 := 0; fi17 < 3; fi17++ {
                            var v18 uint32 = 2147483648
                            v2 = v2
                            if true {
                                var v19 string = ("foo bar" + "a")
                                _ = v19
                            }
                            delete(v4, (^(v18)))
                            _ = fi17
                            _ = v18
                        }
                        if true {
                            v20 := v4[0]
                            _ = v20
                        } else {
                            v2 = v2
                            var v21 uint64 = (v2 & v2)
                            _ = v21
                        }
                    }
                    switch ("" + "foo bar") {
                    case "a":
                        for fi22 := 0; fi22 < 3; fi22++ {
                            var v23 bool = false
                            _ = fi22
                            _ = v23
                        }
                    default:
                        if true {
                            if false {
                                _ = len(v4)
                                v5 -= v5
                            }
                            for fi24 := 0; fi24 < 7; fi24++ {
                                if true {
                                    v25 := GenS2{F0: -128}
                                    _ = v25
                                } else {
                                    fmt.Println(v5, fi10)
                                }
                                v26 := GenS1{F0: v6, F1: v6}
                                var v27 int32 = 2147483647
                                _ = fi24
                                _ = v26
                                _ = v27
                            }
                            var v28 uint32 = 382
                            _ = v28
                        }
                    }
                } else {
                    var v29 uint16 = 2
                    _ = v29
                }
            }
        }
        var v30 uint64 = uint64(v6)
        _ = fi10
        _ = v30
    }
    fmt.Println(v6, v3, v5)
    for fi31 := 0; fi31 < 16; fi31++ {
        if (-1 != 1) {
            var v32 uint = (v6 + v6)
            fmt.Println(v6, v32, v3)
            _ = v32
        }
        if false {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch "\n\t" {
            case "\n\t":
                v3 = v3
            case "a":
                fmt.Println(v3, fi31)
            default:
                fmt.Println(v2)
            }
        }
        for fi33 := 0; fi33 < 20; fi33++ {
            v34 := GenS3{F0: ((("foo bar" + "\n\t") + "\n\t") + ("foo bar" + "hello"))}
            _ = fi33
            _ = v34
        }
        _ = fi31
    }
    fmt.Println(v5, v2)
    if false {
        _ = len(v4)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
