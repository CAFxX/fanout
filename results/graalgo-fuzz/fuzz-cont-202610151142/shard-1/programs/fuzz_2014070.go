package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint64 }
type GenS2 struct { F0 int8; F1 int16; F2 int8 }

func main() {
    var v1 int16 = 1000
    var v2 int = 127
    v3 := []string{("\n\t" + ""), "\n\t"}
    v4 := make(map[string]uint32)
    var v5 uint16 = 65535
    v6 := make(map[string]bool)
    v2 -= 26
    switch "\n\t" {
    case "a":
        for fi7 := 0; fi7 < 19; fi7++ {
            _ = len(v3)
            var v8 int16 = v1
            _ = fi7
            _ = v8
        }
    case "foo bar":
        var v9 bool = (1 <= 4294967295)
        _ = v9
    default:
        delete(v6, (((("hello" + "hello") + "\n\t") + "\n\t") + "foo bar"))
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        for fi11 := 0; fi11 < 13; fi11++ {
            v12 := GenS1{F0: "a", F1: (v2 != 32767), F2: 2}
            v2 += v2
            var v13 int = (fi10 ^ v2)
            _ = fi11
            _ = v12
            _ = v13
        }
        if (127 <= uint8(v5)) {
            v2 -= (-(v2))
            v14 := v4["\n\t"]
            _ = v14
        }
        _ = fi10
    }
    fmt.Println(v2)
    v2 = 606
    fmt.Println(v2, v1)
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    var v15 uint64 = 42
    v15 += (v15 << ((18446744073709551615) & 63))
    v5 = (v5 >> ((69) & 63))
    fmt.Println(v1, v15)
    v16 := GenS1{F0: "foo bar", F1: (v2 >= v2), F2: v15}
    for fi17 := 0; fi17 < 20; fi17++ {
        if (-1 < 0) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if (!false) {
                if (100 <= 42) {
                    var v18 int = v2
                    if true {
                        v15 = 2
                        switch "foo bar" {
                        case "foo bar":
                            v15 -= (v15 | (v15 | v15))
                        }
                    }
                    _ = v18
                }
                fmt.Println(v15, v2)
                _ = len(v3)
            }
        } else {
            if (!(!(22 >= 70))) {
                var v19 uint = 4294967295
                v15 = 4294967295
                _ = v19
            } else {
                v2 += fi17
            }
        }
        v3 = append(v3, ("a" + ("" + "hello")))
        _ = fi17
    }
    for fi20 := 0; fi20 < 7; fi20++ {
        v21 := GenS2{F0: -128, F1: (v1 << ((65535) & 63)), F2: 0}
        _ = fi20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
