package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16 }
func genFn0(p0 int8) string {
    p0 -= p0
    v4 := []int{0, -9223372036854775808}
    _ = len(v4)
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := GenS1{F0: ("" + (("hello" + "") + ("\n\t" + ("" + ("" + "a"))))), F1: 725}
        var v7 float32 = float32((3.14159 - (3.14159 + (1e10 - 1.5))))
        _ = fi5
        _ = v6
        _ = v7
    }
    switch 127 {
    case 32767:
        var v8 int64 = -9223372036854775808
        _ = v8
    case -2147483648:
        if ((!((true || (18446744073709551615 >= 357)) && (2147483648 >= 1))) && ((4294967295 != 710) && true)) {
            var v9 string = ((("hello" + ((((((("a" + "hello") + "\n\t") + "foo bar") + "foo bar") + ("" + "")) + ("hello" + "hello")) + ("\n\t" + "\n\t"))) + "\n\t") + "foo bar")
            p0 = -88
            for fi10 := 0; fi10 < 1; fi10++ {
                for fi11 := 0; fi11 < 13; fi11++ {
                    if (!false) {
                        var v12 int32 = -114
                        v9 = ""
                        _ = v12
                    } else {
                        fmt.Println(v9, fi11, fi10)
                        v13 := GenS1{F0: v9, F1: 4294967295}
                        _ = v13
                    }
                    v4 = append(v4, 127)
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v14 uint = 1
                    _ = fi11
                    _ = v14
                }
                switch fi10 {
                case 32768:
                    var v15 int = -2147483648
                    _ = v15
                case 0:
                    _ = len(v4)
                default:
                    v16 := GenS2{F0: 255}
                    _ = v16
                }
                _ = fi10
            }
            _ = v9
        } else {
            p0 = (p0 % ((p0) | 1))
            p0 = (p0 ^ -128)
            switch "hello" {
            case "a":
                if len(v4) > 0 {
                    v4[0] = 700
                }
            case "\n\t":
                var v17 uint = 522
                _ = v17
            }
        }
    }
    _ = p0
    _ = v4
    return ""
}

func main() {
    var v1 bool = (3.14159 < 1e10)
    var v2 uint = 42
    v3 := make(map[uint32]int)
    v1 = v1
    v18 := GenS3{F0: -32768}
    fmt.Println(v2, v2)
    if v1 {
        v1 = v1
    }
    fmt.Println(v2, v2)
    var v19 uint32 = 42
    var v20 int = -2147483648
    v20 -= (-(v20))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
