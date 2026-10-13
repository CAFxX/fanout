package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int64 }
type GenS2 struct { F0 int }

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 905
    var v3 int16 = 1
    v1[(("\n\t" + "\n\t") + "foo bar")] = (true && (v2 > v2))
    for fi4 := 0; fi4 < 4; fi4++ {
        for fi5 := 0; fi5 < 17; fi5++ {
            if false {
                switch fi4 {
                case -332:
                    switch ((((("\n\t" + "foo bar") + "foo bar") + ("foo bar" + ((("hello" + "foo bar") + "foo bar") + (("" + ("\n\t" + ("foo bar" + "\n\t"))) + "foo bar")))) + "hello") + "foo bar") {
                    case "foo bar":
                        v6 := v1["\n\t"]
                        _ = v6
                    case "\n\t":
                        fmt.Println(v2, v2)
                    }
                case 128:
                    var v7 int64 = -1
                    _ = v7
                case -2147483648:
                    v8 := v1[("hello" + (("" + ("" + ("foo bar" + "hello"))) + ""))]
                    _ = v8
                default:
                    delete(v1, ("" + "foo bar"))
                }
            }
            var v9 float32 = 0.0
            v1[""] = (1 != 65535)
            _ = fi5
            _ = v9
        }
        var v10 uint = uint(v2)
        _ = fi4
        _ = v10
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        v12 := GenS1{F0: fi11, F1: 537, F2: -9223372036854775808}
        _ = fi11
        _ = v12
    }
    if false {
        v1[""] = (42 == 288)
    } else {
        v13 := []int{2147483647}
        var v14 uint16 = 935
        fmt.Println(v3, v14, v3)
        _ = v13
        _ = v14
    }
    v1["hello"] = (!(127 == 2147483647))
    for fi15 := 0; fi15 < 5; fi15++ {
        switch "a" {
        case "a":
            v2 -= (^(v2))
        default:
            var v16 uint16 = 32767
            _ = v16
        }
        _ = fi15
    }
    if (2147483647 >= 0) {
        for fi17 := 0; fi17 < 20; fi17++ {
            v18 := GenS2{F0: fi17}
            _ = fi17
            _ = v18
        }
        var v19 float64 = float64(((0.0 - -2.25) - (float32(((0.1 / ((3.14159 / -2.25) * (1.5 + (-2.25 - (1.5 + float64(3.14159)))))) * (1.5 * 1.5))) - 0.0)))
        for fi20 := 0; fi20 < 18; fi20++ {
            if false {
                v1[("hello" + "a")] = (125 <= -1)
            } else {
                fmt.Println(v3)
                v21 := []int{4294967295}
                _ = v21
            }
            _ = fi20
        }
        _ = v19
    }
    var v22 uint8 = 128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
}
