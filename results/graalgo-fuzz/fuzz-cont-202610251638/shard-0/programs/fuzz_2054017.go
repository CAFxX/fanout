package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint64 }
type GenS2 struct { F0 int; F1 int32; F2 string }
type GenS3 struct { F0 int8; F1 int8; F2 int16 }

func main() {
    var v1 uint16 = 215
    var v2 int32 = 42
    var v3 uint = 509
    var v4 uint32 = 1
    v2 = v2
    fmt.Println(v2, v3, v2)
    var v5 uint64 = 247
    v6 := GenS1{F0: (("" + (("a" + ("\n\t" + ((("\n\t" + "hello") + "a") + ("a" + ("foo bar" + "\n\t"))))) + "hello")) + ("" + "foo bar")), F1: -128, F2: 0}
    var v7 uint = v3
    var v8 int16 = 1
    switch ("foo bar" + (("hello" + "foo bar") + "a")) {
    case "hello":
        if ((v8 >> (((v7 - uint(v5))) & 63)) > (^(v8))) {
            v3 -= v7
            v1 = 0
        }
    case "a":
        if (44 >= 46) {
            if false {
                var v9 int = int(v7)
                _ = v9
            }
            for fi10 := 0; fi10 < 17; fi10++ {
                v11 := GenS1{F0: "", F1: -128, F2: 42}
                _ = fi10
                _ = v11
            }
        } else {
            v12 := GenS1{F0: ("foo bar" + "hello"), F1: -112, F2: (v5 & (v5 & 411))}
            _ = v12
        }
    default:
        v13 := []int{}
        _ = v13
    }
    switch "" {
    case "\n\t":
        for fi14 := 0; fi14 < 20; fi14++ {
            switch int(v1) {
            case -406:
                v15 := GenS2{F0: fi14, F1: (^(v2)), F2: ""}
                _ = v15
            case 1:
                if ((true && (((!false) || false) && false)) || (!false)) {
                    v16 := []int{32767, (fi14 & (fi14 | fi14)), fi14}
                    _ = v16
                }
            case -562:
                var v17 bool = false
                _ = v17
            default:
                var v18 uint64 = 0
                _ = v18
            }
            v19 := GenS3{F0: 24, F1: 73, F2: (-(v8))}
            var v20 uint32 = (^(v4))
            _ = fi14
            _ = v19
            _ = v20
        }
    case "hello":
        v21 := []int{42}
        _ = v21
    case "foo bar":
        var v22 int16 = v8
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
