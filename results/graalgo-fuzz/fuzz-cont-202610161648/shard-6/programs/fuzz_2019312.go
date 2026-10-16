package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 bool }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := []string{((((((("" + ("\n\t" + "foo bar")) + "") + "hello") + (("\n\t" + "\n\t") + ("" + ""))) + "hello") + ("" + ("a" + "foo bar"))) + ""), "foo bar", "hello", "hello"}
    var v2 float32 = (0.1 - float32((((float64(float32(3.14159)) / 1.0) / 1.0) * 1e10)))
    var v3 float32 = v2
    v4 := []uint32{0, 441, 0, 151}
    var v5 uint32 = 4294967295
    var v6 uint32 = (v5 & (v5 - v5))
    _ = len(v4)
    _ = len(v4)
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 int64 = 0
        v2 = 1e10
        if (128 == -618) {
            if false {
                fmt.Println(v2, v2)
            }
        } else {
            if len(v1) > 0 {
                v1[0] = "a"
            }
            var v9 uint32 = v5
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    v5 -= 1
    v10 := GenS1{F0: 880, F1: 255, F2: (-1 > 1)}
    for fi11 := 0; fi11 < 13; fi11++ {
        v6 = 2147483648
        _ = fi11
    }
    var v12 int32 = 0
    switch ((("hello" + "foo bar") + "foo bar") + (("foo bar" + "a") + (("\n\t" + "") + ("" + "")))) {
    case "\n\t":
        v13 := GenS2{F0: 1}
        _ = v13
    case "":
        var v14 uint16 = 347
        _ = v14
    }
    v4 = append(v4, 2)
    if false {
        _ = len(v4)
        for fi15 := 0; fi15 < 1; fi15++ {
            v16 := GenS1{F0: 0, F1: 1, F2: (0.0 <= 3.14159)}
            _ = fi15
            _ = v16
        }
    } else {
        v1 = append(v1, (("a" + ("" + "\n\t")) + ("a" + "hello")))
        var v17 uint32 = (^(v5))
        _ = v17
    }
    var v18 bool = (((!false) && ((43 <= -1) || ((!(!false)) || (((("" + "hello") != "a") && false) && (-649 == 42))))) && (false || (766 == 978)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
