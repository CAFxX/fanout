package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 int32 }

func main() {
    var v1 uint16 = 661
    var v2 float32 = float32(-2.25)
    var v3 float64 = 3.14159
    switch ("foo bar" + "hello") {
    case "":
        switch "hello" {
        case "foo bar":
            var v4 string = ("hello" + "a")
            _ = v4
        case "\n\t":
            var v5 uint16 = v1
            _ = v5
        case "hello":
            v6 := []int{-354, 32767, -2147483648, 128}
            _ = v6
        }
    case "\n\t":
        switch ("\n\t" + ("\n\t" + "a")) {
        case "":
            fmt.Println(v3, v2, v2)
        default:
            for fi7 := 0; fi7 < 11; fi7++ {
                switch ("hello" + "") {
                case "":
                    var v8 bool = (!(!true))
                    _ = v8
                default:
                    var v9 int64 = 42
                    _ = v9
                }
                _ = fi7
            }
        }
    default:
        v3 = -2.25
    }
    var v10 int16 = 0
    v11 := []int{32768, 128, 128}
    if false {
        if false {
            var v12 bool = true
            _ = v12
        }
    } else {
        _ = len(v11)
    }
    var v13 uint32 = 856
    for fi14 := 0; fi14 < 8; fi14++ {
        if false {
            v15 := GenS1{F0: "", F1: 42, F2: -933}
            v16 := GenS1{F0: "foo bar", F1: 385, F2: 1}
            _ = v15
            _ = v16
        }
        _ = fi14
    }
    v13 += (v13 | (v13 >> ((242) & 63)))
    fmt.Println(v3, v10)
    v13 = v13
    fmt.Println(v10)
    v2 = 3.14159
    for fi17 := 0; fi17 < 11; fi17++ {
        for fi18 := 0; fi18 < 13; fi18++ {
            v13 = 2147483648
            if len(v11) > 0 {
                v11[0] = (fi18 ^ fi17)
            }
            v11 = append(v11, fi18)
            var v19 int16 = -32768
            _ = fi18
            _ = v19
        }
        if (v13 >= 2) {
            v1 += (^(v1))
            v20 := GenS1{F0: "", F1: 0, F2: 1}
            if true {
                v3 = v3
            }
            _ = v20
        }
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        _ = fi17
    }
    v10 -= 0
    v1 = (^(v1))
    v3 += v3
    var v21 string = ("a" + ("hello" + ("" + (("foo bar" + (("a" + "\n\t") + "a")) + "hello"))))
    v22 := GenS1{F0: v21, F1: 4294967295, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
