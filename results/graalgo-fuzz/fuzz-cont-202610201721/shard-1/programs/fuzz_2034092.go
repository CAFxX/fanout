package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }

func main() {
    v1 := []string{"foo bar"}
    v2 := make(map[int]bool)
    v3 := make(map[uint32]string)
    v4 := v3[2147483647]
    switch 0 {
    case 32768:
        v4 = ""
    case -662:
        v1 = append(v1, v4)
    case -325:
        var v5 string = v4
        _ = v5
    }
    switch 9223372036854775807 {
    case 4294967295:
        if (2147483647 == 2) {
            switch "foo bar" {
            case "\n\t":
                if (1 != -128) {
                    v4 = ("foo bar" + "\n\t")
                    for fi6 := 0; fi6 < 7; fi6++ {
                        var v7 int64 = 42
                        v4 = ((v4 + v4) + "foo bar")
                        _ = fi6
                        _ = v7
                    }
                    v4 = (v4 + (((v4 + "a") + v4) + (v4 + ("foo bar" + v4))))
                } else {
                    v4 = ("\n\t" + (v4 + v4))
                }
            }
            v4 = ("\n\t" + v4)
            var v8 int8 = 0
            _ = v8
        }
    default:
        v4 = ((v4 + "hello") + v4)
    }
    fmt.Println(v4)
    if true {
        var v9 int16 = 42
        var v10 float64 = 1e10
        _ = v10
        _ = v9
    }
    v4 = "foo bar"
    if false {
        switch v4 {
        case "\n\t":
            delete(v3, 50)
        default:
            for fi11 := 0; fi11 < 1; fi11++ {
                v12 := GenS1{F0: 4294967295, F1: 1}
                _ = fi11
                _ = v12
            }
        }
    } else {
        v13 := GenS1{F0: 939, F1: -672}
        _ = v13
    }
    var v14 string = v4
    v15 := GenS1{F0: 1, F1: -1}
    for fi16 := 0; fi16 < 7; fi16++ {
        v14 = ("foo bar" + v14)
        v1 = append(v1, (v4 + ("\n\t" + (((v4 + v14) + "a") + v14))))
        v4 = ""
        _ = fi16
    }
    v14 = ("a" + "\n\t")
    var v17 int = -88
    var v18 int64 = -1
    for fi19 := 0; fi19 < 10; fi19++ {
        var v20 uint32 = 434
        _ = fi19
        _ = v20
    }
    v21 := v3[4294967295]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
