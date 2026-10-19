package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint64; F1 int }
type GenS7 struct { F0 int16; F1 uint; F2 uint32 }
func genFn0(p0 int64, p1 uint8) int32 {
    v4 := GenS1{F0: ""}
    if false {
        var v5 string = "\n\t"
        var v6 bool = (!false)
        switch ("foo bar" + v5) {
        case "a":
            fmt.Println(p1)
        case "":
            v7 := []int{1, 1, -9223372036854775808}
            _ = v7
        default:
            var v8 uint = 1
            _ = v8
        }
        _ = v5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    return -2147483648
}

func main() {
    var v1 int = -1
    var v2 uint = 18446744073709551615
    v3 := []string{"\n\t", "foo bar", ""}
    switch "" {
    case "\n\t":
        _ = len(v3)
    case "hello":
        v2 = 45
    case "":
        v9 := GenS1{F0: (("hello" + "\n\t") + ("a" + ""))}
        _ = v9
    default:
        v1 += (v1 ^ (-(v1)))
    }
    v10 := GenS2{F0: 32767}
    v2 -= (^(v2))
    fmt.Println(v1)
    v11 := GenS3{F0: (((v1 << ((v2) & 63)) >= v1) && (true && (0 > -9223372036854775808))), F1: 267, F2: 32767}
    fmt.Println(v1, v1)
    v3 = append(v3, "")
    v1 = (-(v1))
    fmt.Println(v2, v2)
    if true {
        if false {
            for fi12 := 0; fi12 < 13; fi12++ {
                switch "hello" {
                case "":
                    var v13 int = (v1 << (((v2 - v2)) & 63))
                    _ = v13
                case "hello":
                    var v14 string = "hello"
                    _ = v14
                case "foo bar":
                    switch fi12 {
                    case 740:
                        fmt.Println(fi12)
                    default:
                        switch "foo bar" {
                        case "foo bar":
                            var v15 string = ("" + "hello")
                            _ = v15
                        }
                    }
                }
                _ = fi12
            }
        }
        v2 -= v2
    } else {
        for fi16 := 0; fi16 < 18; fi16++ {
            for fi17 := 0; fi17 < 10; fi17++ {
                v18 := GenS4{F0: v1, F1: 2147483647}
                v2 = v2
                _ = fi17
                _ = v18
            }
            _ = fi16
        }
        for fi19 := 0; fi19 < 5; fi19++ {
            v2 += v2
            fmt.Println(v1)
            switch ("\n\t" + "foo bar") {
            case "":
                v20 := GenS5{F0: 18446744073709551615}
                _ = v20
            case "foo bar":
                for fi21 := 0; fi21 < 14; fi21++ {
                    var v22 int32 = -2147483648
                    _ = len(v3)
                    v2 = 42
                    _ = fi21
                    _ = v22
                }
            case "\n\t":
                if (!true) {
                    switch v1 {
                    case 1:
                        var v23 string = ("" + (("hello" + "\n\t") + ("" + "a")))
                        _ = v23
                    default:
                        v3 = append(v3, "\n\t")
                    }
                }
            }
            _ = fi19
        }
        var v24 string = "hello"
        _ = v24
    }
    v1 += (v1 + v1)
    if true {
        fmt.Println(v2)
        switch "\n\t" {
        case "a":
            v1 = 42
        case "foo bar":
            for fi25 := 0; fi25 < 5; fi25++ {
                var v26 uint8 = 16
                switch (fi25 / ((32767) | 1)) {
                case -1:
                    for fi27 := 0; fi27 < 16; fi27++ {
                        if ("a" >= "\n\t") {
                            if false {
                                var v28 uint32 = 42
                                var v29 uint = (^(v2))
                                _ = v28
                                _ = v29
                            } else {
                                if (false || (true && true)) {
                                    v30 := GenS6{F0: 96, F1: (v1 & 42)}
                                    _ = v30
                                }
                            }
                            switch v1 {
                            case 2147483647:
                                v26 = (v26 - v26)
                            case 0:
                                var v31 int16 = 0
                                _ = v31
                            }
                        } else {
                            _ = len(v3)
                        }
                        _ = fi27
                    }
                }
                v2 += 255
                _ = fi25
                _ = v26
            }
        case "hello":
            var v32 uint32 = 525
            _ = v32
        default:
            v2 = (v2 >> (((v2 ^ v2)) & 63))
        }
        var v33 float32 = 1.5
        _ = v33
    } else {
        v34 := GenS7{F0: 0, F1: (v2 - v2), F2: 65535}
        v3 = append(v3, "foo bar")
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
}
