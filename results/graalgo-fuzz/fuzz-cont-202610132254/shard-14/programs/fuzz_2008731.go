package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 uint16; F2 int8 }
type GenS3 struct { F0 int8; F1 int64; F2 uint8 }
type GenS4 struct { F0 int; F1 string; F2 uint64 }
func genFn0(p0 int8, p1 int8, p2 int16) int16 {
    var v6 float32 = (3.14159 / float32(1e10))
    var v7 int = -40
    v8 := GenS1{F0: 127}
    v7 -= (v7 ^ v7)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return (p2 * p2)
}

func main() {
    var v1 string = "\n\t"
    var v2 int = 32768
    v3 := []string{}
    var v4 uint8 = 1
    var v5 int16 = -32768
    fmt.Println(v4, v2, v2)
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := GenS2{F0: -1, F1: 374, F2: 42}
        if true {
            v3 = append(v3, "hello")
        }
        switch v1 {
        case "foo bar":
            fmt.Println(v2, v5)
        case "":
            v3 = append(v3, (v1 + v1))
        case "a":
            var v11 int8 = -91
            _ = v11
        }
        _ = fi9
        _ = v10
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        for fi13 := 0; fi13 < 1; fi13++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v2 += (v2 * v2)
            v4 = (v4 ^ (v4 / (((v4 & v4)) | 1)))
            _ = fi13
        }
        _ = fi12
    }
    v14 := GenS2{F0: -1, F1: 65535, F2: -82}
    for fi15 := 0; fi15 < 18; fi15++ {
        for fi16 := 0; fi16 < 8; fi16++ {
            switch 485 {
            case -1:
                v17 := GenS3{F0: 1, F1: 42, F2: 222}
                _ = v17
            case -707:
                var v18 bool = true
                _ = v18
            case 684:
                var v19 bool = (2 >= 2)
                _ = v19
            }
            fmt.Println(v4, fi15, v4)
            _ = fi16
        }
        _ = fi15
    }
    switch ("\n\t" + "") {
    case "foo bar":
        var v20 int8 = 1
        _ = v20
    }
    var v21 uint16 = 480
    for fi22 := 0; fi22 < 7; fi22++ {
        switch (fi22 % ((127) | 1)) {
        case 4294967295:
            if (v4 > 153) {
                var v23 int16 = v5
                _ = v23
            }
        default:
            switch "hello" {
            case "\n\t":
                v24 := GenS1{F0: 2}
                _ = v24
            case "a":
                _ = len(v3)
            default:
                v25 := GenS4{F0: v2, F1: v1, F2: 529}
                _ = v25
            }
        }
        switch "a" {
        case "foo bar":
            v21 = v21
        case "hello":
            fmt.Println(v1)
        case "a":
            v21 -= 143
        }
        _ = fi22
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        if (true || (true && ((0.1 / 1.5) >= 0.1))) {
            fmt.Println(v1, fi26, v4)
            switch "" {
            case "":
                switch v2 {
                case -2147483648:
                    switch v2 {
                    case 582:
                        for fi27 := 0; fi27 < 8; fi27++ {
                            switch (("\n\t" + ("a" + v1)) + (("" + "hello") + v1)) {
                            case "\n\t":
                                var v28 int32 = -194
                                _ = v28
                            default:
                                if true {
                                    v5 = 42
                                    fmt.Println(fi27)
                                } else {
                                    var v29 uint64 = 42
                                    _ = v29
                                }
                            }
                            var v30 uint = 2
                            v21 = v21
                            _ = fi27
                            _ = v30
                        }
                    }
                }
            case "\n\t":
                for fi31 := 0; fi31 < 9; fi31++ {
                    fmt.Println(fi26)
                    _ = fi31
                }
            case "foo bar":
                if (!(v5 >= (v5 & (^(v5))))) {
                    v5 = 42
                } else {
                    v1 = ""
                }
            }
        } else {
            switch ("a" + (v1 + v1)) {
            case "foo bar":
                for fi32 := 0; fi32 < 20; fi32++ {
                    v5 = (-(v5))
                    v21 = (v21 % ((v21) | 1))
                    _ = fi32
                }
            case "\n\t":
                v33 := GenS2{F0: 490, F1: (^(v21)), F2: 100}
                _ = v33
            case "a":
                for fi34 := 0; fi34 < 7; fi34++ {
                    fmt.Println(fi34)
                    _ = fi34
                }
            default:
                v4 = v4
            }
            v35 := GenS3{F0: 42, F1: 0, F2: (^(v4))}
            _ = v35
        }
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
