package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }

func main() {
    v1 := []string{"a", "\n\t"}
    var v2 uint32 = 769
    var v3 string = ("foo bar" + "\n\t")
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    var v4 uint16 = 63
    v1 = append(v1, "")
    var v5 uint16 = v4
    switch ("\n\t" + (v3 + ("hello" + v3))) {
    case "foo bar":
        for fi6 := 0; fi6 < 12; fi6++ {
            v7 := GenS1{F0: 2147483647, F1: v4}
            v5 += 549
            switch (^(fi6)) {
            case 128:
                var v8 uint32 = (v2 + v2)
                _ = v8
            case 127:
                switch v3 {
                case "hello":
                    if (v3 <= v3) {
                        v2 = 1
                    } else {
                        switch ("hello" + ((("hello" + ((((v3 + v3) + "") + "hello") + v3)) + ((v3 + "hello") + "foo bar")) + v3)) {
                        case "hello":
                            var v9 bool = false
                            _ = v9
                        case "":
                            v10 := GenS1{F0: 42, F1: 32767}
                            _ = v10
                        case "a":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        default:
                            _ = len(v1)
                        }
                    }
                case "":
                    v11 := GenS1{F0: -2147483648, F1: 42}
                    _ = v11
                default:
                    v4 = 2
                }
            }
            _ = fi6
            _ = v7
        }
    case "a":
        v2 += v2
    }
    var v12 uint64 = 127
    fmt.Println(v5, v2, v12)
    var v13 float64 = 3.14159
    if true {
        if len(v1) > 0 {
            v1[0] = ((v3 + v3) + ((v3 + (v3 + "foo bar")) + "foo bar"))
        }
    }
    v5 -= (v4 + (^(v4)))
    for fi14 := 0; fi14 < 1; fi14++ {
        var v15 uint32 = v2
        for fi16 := 0; fi16 < 11; fi16++ {
            fmt.Println(v12, v3)
            _ = fi16
        }
        _ = fi14
        _ = v15
    }
    var v17 int32 = -14
    var v18 int64 = 42
    var v19 int8 = int8(v12)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v20 uint64 = (^(v12))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
