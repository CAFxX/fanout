package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 bool }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 uint }

func main() {
    var v1 bool = true
    var v2 uint32 = 2
    var v3 uint16 = 1
    var v4 string = "hello"
    v4 = ((v4 + v4) + "hello")
    for fi5 := 0; fi5 < 15; fi5++ {
        for fi6 := 0; fi6 < 8; fi6++ {
            if v1 {
                var v7 uint8 = 175
                v3 -= (v3 + v3)
                _ = v7
            }
            v8 := GenS1{F0: ((v4 + (v4 + v4)) + v4), F1: 2, F2: (!v1)}
            _ = fi6
            _ = v8
        }
        v1 = v1
        switch v4 {
        case "a":
            var v9 int16 = -1
            _ = v9
        case "hello":
            switch fi5 {
            case 1:
                for fi10 := 0; fi10 < 18; fi10++ {
                    v1 = false
                    var v11 uint32 = (v2 | 0)
                    _ = fi10
                    _ = v11
                }
            case 0:
                v3 = 536
            }
        }
        _ = fi5
    }
    v4 = "hello"
    switch ("" + "") {
    case "hello":
        for fi12 := 0; fi12 < 17; fi12++ {
            if v1 {
                if (!v1) {
                    if v1 {
                        switch fi12 {
                        case 1:
                            switch v4 {
                            case "hello":
                                v4 = v4
                            case "":
                                if (3.14159 <= 1.5) {
                                    v1 = (42 != 128)
                                }
                            case "foo bar":
                                switch v4 {
                                case "hello":
                                    var v13 bool = v1
                                    _ = v13
                                case "foo bar":
                                    v14 := GenS1{F0: "foo bar", F1: 255, F2: true}
                                    _ = v14
                                default:
                                    switch fi12 {
                                    case 32767:
                                        v15 := GenS2{F0: (v2 - v2), F1: -2147483648}
                                        _ = v15
                                    case 388:
                                        v16 := GenS1{F0: v4, F1: 128, F2: ((v1 && ((589 < 32767) || (!false))) || (4294967295 != 18446744073709551615))}
                                        _ = v16
                                    default:
                                        fmt.Println(fi12, v1, v3)
                                    }
                                }
                            }
                        case 620:
                            v17 := []int{(fi12 - fi12), -9223372036854775808}
                            _ = v17
                        case 128:
                            v18 := GenS3{F0: 113}
                            _ = v18
                        }
                        v1 = (!false)
                    }
                    fmt.Println(v1, fi12)
                }
            }
            fmt.Println(v1, v1)
            v4 = ""
            _ = fi12
        }
    case "foo bar":
        for fi19 := 0; fi19 < 16; fi19++ {
            var v20 uint32 = (v2 * (^(v2)))
            _ = fi19
            _ = v20
        }
    case "a":
        var v21 float32 = 1e10
        _ = v21
    }
    v22 := []int{}
    if v1 {
        fmt.Println(v2, v2)
        for fi23 := 0; fi23 < 11; fi23++ {
            _ = len(v22)
            for fi24 := 0; fi24 < 15; fi24++ {
                switch v4 {
                case "a":
                    fmt.Println(v1, v3)
                default:
                    fmt.Println(fi24, fi24)
                    var v25 float64 = -2.25
                    _ = v25
                }
                v4 = v4
                _ = fi24
            }
            _ = fi23
        }
        fmt.Println(v2, v1)
    }
    v1 = v1
    v26 := GenS2{F0: uint32(v3), F1: 42}
    var v27 int = 32768
    switch "" {
    case "\n\t":
        fmt.Println(v2)
    }
    if false {
        v28 := GenS2{F0: (v2 + v2), F1: int32(v3)}
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
