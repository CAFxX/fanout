package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint }

func main() {
    var v1 string = ((("" + "") + "\n\t") + (("hello" + (("foo bar" + (("hello" + ("\n\t" + "hello")) + ("\n\t" + "hello"))) + ("foo bar" + "foo bar"))) + ("a" + "\n\t")))
    var v2 float32 = (1.5 * ((1e10 + ((float32(1.5) + (1.5 / 0.1)) - float32((1e10 / 3.14159)))) / 1.5))
    var v3 uint = 1
    var v4 string = ("hello" + v1)
    var v5 uint32 = 2147483648
    var v6 int16 = 1
    v7 := []int{42, -300, -434}
    fmt.Println(v2, v6, v4)
    v5 = v5
    if false {
        for fi8 := 0; fi8 < 14; fi8++ {
            if false {
                for fi9 := 0; fi9 < 16; fi9++ {
                    v10 := GenS1{F0: uint32(v3), F1: 692}
                    _ = fi9
                    _ = v10
                }
                switch 32768 {
                case 9223372036854775807:
                    if true {
                        var v11 uint16 = 32767
                        for fi12 := 0; fi12 < 5; fi12++ {
                            switch v4 {
                            case "":
                                var v13 int = fi8
                                _ = v13
                            }
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            v14 := GenS2{F0: (^(v11))}
                            _ = fi12
                            _ = v14
                        }
                        for fi15 := 0; fi15 < 7; fi15++ {
                            v3 = 42
                            if len(v7) > 0 {
                                v7[0] = 280
                            }
                            if true {
                                switch fi8 {
                                case 0:
                                    var v16 bool = (false && (!(47 == 127)))
                                    _ = v16
                                }
                                var v17 int = -395
                                fmt.Println(v3, v11, v17)
                                _ = v17
                            }
                            _ = fi15
                        }
                        _ = v11
                    } else {
                        var v18 int = fi8
                        var v19 int32 = 1
                        _ = v18
                        _ = v19
                    }
                }
            }
            v5 = (v5 + uint32(v3))
            v6 -= 1000
            _ = fi8
        }
        var v20 float32 = -2.25
        _ = v20
    }
    v4 = "\n\t"
    fmt.Println(v6, v3)
    var v21 float32 = 1.5
    for fi22 := 0; fi22 < 4; fi22++ {
        if ((v5 % ((17) | 1)) >= v5) {
            v23 := GenS1{F0: v5, F1: v3}
            for fi24 := 0; fi24 < 11; fi24++ {
                switch v1 {
                case "\n\t":
                    switch v4 {
                    case "foo bar":
                        v4 = (("a" + "a") + "\n\t")
                    default:
                        fmt.Println(v4, v4, v3)
                    }
                case "":
                    v25 := GenS2{F0: 0}
                    _ = v25
                default:
                    switch (fi24 & fi24) {
                    case 32767:
                        v26 := GenS3{F0: v3}
                        _ = v26
                    case 42:
                        v5 -= uint32(v3)
                    case 168:
                        v7 = append(v7, 9223372036854775807)
                    }
                }
                _ = fi24
            }
            v1 = v4
            _ = v23
        }
        _ = fi22
    }
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
