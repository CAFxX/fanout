package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint32; F2 int8 }

func main() {
    var v1 string = "a"
    var v2 string = ((v1 + "foo bar") + "hello")
    v3 := []string{(v1 + v2), v2}
    v4 := GenS1{F0: true, F1: -461}
    fmt.Println(v2, v2)
    if false {
        if ((false || true) || ((1 >= 759) && true)) {
            var v5 int16 = 489
            if len(v3) > 0 {
                v3[0] = v1
            }
            _ = len(v3)
            _ = v5
        }
    }
    for fi6 := 0; fi6 < 14; fi6++ {
        for fi7 := 0; fi7 < 5; fi7++ {
            if (-9223372036854775808 != -9223372036854775808) {
                v2 = (v1 + ("\n\t" + ("hello" + (v2 + v1))))
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v8 uint16 = 2
                var v9 int32 = 860
                v10 := GenS2{F0: 2147483647, F1: 407, F2: 100}
                _ = v10
                _ = v8
                _ = v9
            }
            _ = len(v3)
            v1 = v2
            _ = fi7
        }
        _ = fi6
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        fmt.Println(v1, fi11)
        if (0 == 0) {
            v3 = append(v3, (("hello" + ("" + ((v2 + "hello") + (((v1 + "a") + (((v1 + "foo bar") + v2) + "")) + v2)))) + ("a" + "foo bar")))
            v2 = ((("hello" + (v2 + ("a" + ("hello" + (v2 + v2))))) + "hello") + v1)
        } else {
            if true {
                fmt.Println(fi11, fi11)
                for fi12 := 0; fi12 < 5; fi12++ {
                    if (630 < 1) {
                        v1 = (v2 + v2)
                    }
                    fmt.Println(fi11)
                    _ = fi12
                }
            } else {
                if (13 == 0) {
                    var v13 float64 = 0.1
                    var v14 uint = 4294967295
                    var v15 int8 = 42
                    _ = v13
                    _ = v14
                    _ = v15
                } else {
                    v1 = (("\n\t" + "hello") + v2)
                }
            }
            var v16 uint8 = 132
            _ = v16
        }
        if false {
            switch fi11 {
            case 4294967295:
                var v17 int16 = -1
                _ = v17
            }
        }
        _ = fi11
    }
    v1 = (v1 + ((("\n\t" + "\n\t") + v1) + v1))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch v1 {
    case "hello":
        v1 = v2
    default:
        v18 := GenS1{F0: (!true), F1: 42}
        _ = v18
    }
    v2 = "hello"
    if true {
        fmt.Println(v2, v1)
        fmt.Println(v2, v1)
        _ = len(v3)
    } else {
        fmt.Println(v1, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
