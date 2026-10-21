package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint }
func genFn0() uint8 {
    var v7 string = ""
    var v8 uint = 628
    v8 = (v8 % ((v8) | 1))
    _ = v7
    _ = v8
    return 128
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = -1
    var v3 uint32 = 2
    var v4 uint64 = 4294967295
    v5 := []string{("a" + "foo bar"), (("foo bar" + "\n\t") + "foo bar"), "\n\t"}
    var v6 int16 = 384
    if true {
        v9 := GenS1{F0: ((!(42 != -8)) || (((("a" + (("" + ("hello" + "a")) + ("foo bar" + "\n\t"))) + "foo bar") + ("\n\t" + "a")) >= "\n\t"))}
        _ = v9
    } else {
        v5 = append(v5, (("foo bar" + ("a" + ((("a" + "a") + "") + ("\n\t" + "\n\t")))) + "a"))
        fmt.Println(v4, v4)
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 int16 = int16(fi10)
        _ = fi10
        _ = v11
    }
    if (float64(((3.14159 / 1.0) * -2.25)) != float64(3.14159)) {
        v6 -= v6
    } else {
        v12 := GenS1{F0: (!true)}
        _ = v12
    }
    if false {
        switch (("foo bar" + "foo bar") + "") {
        case "hello":
            switch "" {
            case "foo bar":
                v1[208] = 2147483648
            case "\n\t":
                switch 42 {
                case 9223372036854775807:
                    v4 -= uint64(v3)
                case 2147483647:
                    fmt.Println(v3)
                }
            }
        }
        v13 := GenS2{F0: 42, F1: 710, F2: 626}
        _ = v13
    }
    fmt.Println(v2, v6)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v14 := GenS1{F0: (v6 <= (v6 | 1))}
    var v15 uint = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
