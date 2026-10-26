package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 string }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 int; F1 uint64; F2 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int; F1 uint8 }
type GenS6 struct { F0 uint32 }

func main() {
    var v1 uint = 42
    v2 := []string{"a"}
    var v3 uint = 18446744073709551615
    var v4 uint32 = 134
    var v5 int64 = 1
    v3 -= (v3 & (^(v1)))
    var v6 bool = (!true)
    if len(v2) > 0 {
        v2[0] = ""
    }
    v7 := GenS1{F0: -32768, F1: 1, F2: "foo bar"}
    v8 := GenS2{F0: v6, F1: 242}
    fmt.Println(v6)
    v9 := GenS1{F0: -1, F1: 42, F2: ("" + (("\n\t" + "a") + ("" + "hello")))}
    for fi10 := 0; fi10 < 20; fi10++ {
        for fi11 := 0; fi11 < 5; fi11++ {
            var v12 int8 = 127
            var v13 float64 = 0.1
            v12 -= v12
            _ = fi11
            _ = v12
            _ = v13
        }
        var v14 uint16 = 32767
        for fi15 := 0; fi15 < 6; fi15++ {
            var v16 int8 = int8(v1)
            _ = fi15
            _ = v16
        }
        _ = fi10
        _ = v14
    }
    switch "a" {
    case "hello":
        switch (("a" + ("a" + "hello")) + ("a" + "a")) {
        case "\n\t":
            v17 := GenS3{F0: 42, F1: 0, F2: 4294967295}
            _ = v17
        case "a":
            v6 = (42 == -79)
        case "foo bar":
            if len(v2) > 0 {
                v2[0] = ("\n\t" + "\n\t")
            }
        default:
            if len(v2) > 0 {
                v2[0] = "hello"
            }
        }
    case "":
        var v18 bool = (255 == 164)
        _ = v18
    case "a":
        for fi19 := 0; fi19 < 15; fi19++ {
            switch fi19 {
            case 1:
                v20 := GenS2{F0: false, F1: 42}
                _ = v20
            default:
                v21 := GenS4{F0: 32767}
                _ = v21
            }
            v2 = append(v2, ("hello" + "a"))
            for fi22 := 0; fi22 < 17; fi22++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi22
            }
            _ = fi19
        }
    default:
        var v23 bool = v6
        _ = v23
    }
    if (!v6) {
        if (127 <= 1) {
            var v24 uint32 = (v4 % ((v4) | 1))
            _ = v24
        }
        v3 -= v1
    }
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    for fi25 := 0; fi25 < 9; fi25++ {
        v26 := GenS5{F0: fi25, F1: 92}
        v5 += v5
        _ = fi25
        _ = v26
    }
    for fi27 := 0; fi27 < 17; fi27++ {
        v28 := GenS6{F0: (v4 | (v4 + v4))}
        _ = fi27
        _ = v28
    }
    switch 42 {
    case 32767:
        var v29 float32 = 0.1
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
