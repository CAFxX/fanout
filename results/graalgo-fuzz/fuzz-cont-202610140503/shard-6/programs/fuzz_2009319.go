package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint8; F1 int8; F2 string }
type GenS3 struct { F0 uint16; F1 uint16 }
func genFn0(p0 int64, p1 int64, p2 int16) uint {
    var v4 int64 = (^(p1))
    v5 := GenS1{F0: ("a" + "hello"), F1: 450}
    p1 = (v4 + -61)
    var v6 int = -2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 142
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []string{"\n\t", "\n\t"}
    v3 := []uint32{65535}
    v7 := GenS2{F0: 32, F1: 1, F2: ("hello" + "\n\t")}
    if true {
        v1 -= v1
        var v8 int64 = 9223372036854775807
        for fi9 := 0; fi9 < 12; fi9++ {
            switch fi9 {
            case 32767:
                for fi10 := 0; fi10 < 10; fi10++ {
                    _ = len(v2)
                    _ = fi10
                }
            }
            var v11 int16 = 42
            var v12 float64 = 0.1
            _ = fi9
            _ = v11
            _ = v12
        }
        _ = v8
    } else {
        if true {
            for fi13 := 0; fi13 < 4; fi13++ {
                v1 += 1
                fmt.Println(v1, v1)
                _ = fi13
            }
        } else {
            for fi14 := 0; fi14 < 11; fi14++ {
                v15 := GenS2{F0: 127, F1: 75, F2: ("hello" + "\n\t")}
                v1 += (v1 - 0)
                if false {
                    v1 = 2
                    _ = len(v3)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                _ = fi14
                _ = v15
            }
        }
        v16 := GenS2{F0: 42, F1: -128, F2: "hello"}
        _ = v16
    }
    v1 += v1
    switch ("hello" + ("a" + "a")) {
    case "foo bar":
        for fi17 := 0; fi17 < 8; fi17++ {
            var v18 bool = false
            switch ("foo bar" + ("foo bar" + "\n\t")) {
            case "foo bar":
                if len(v3) > 0 {
                    v3[0] = 2147483648
                }
            }
            if false {
                v19 := GenS3{F0: 2, F1: 32767}
                for fi20 := 0; fi20 < 2; fi20++ {
                    _ = len(v3)
                    var v21 uint16 = 255
                    _ = fi20
                    _ = v21
                }
                _ = v19
            }
            _ = fi17
            _ = v18
        }
    }
    v22 := GenS2{F0: 236, F1: -10, F2: "a"}
    fmt.Println(v1)
    var v23 uint = 865
    if false {
        var v24 bool = (!false)
        var v25 int64 = 9223372036854775807
        _ = v24
        _ = v25
    } else {
        v1 = 82
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
    }
    fmt.Println(v1)
    if len(v3) > 0 {
        v3[0] = 1
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v7
}
