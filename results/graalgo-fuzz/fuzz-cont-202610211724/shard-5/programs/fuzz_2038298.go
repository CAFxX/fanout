package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint64 }

func main() {
    v1 := make(map[int]string)
    var v2 uint = 2
    v3 := []int{338}
    v4 := GenS1{F0: 68, F1: (!((false && (3.14159 <= 3.14159)) && ((-2.25 - 3.14159) != (-2.25 * (((-2.25 * 1.5) / 3.14159) * 1.5))))), F2: 1}
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 bool = true
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 3; fi7++ {
        v2 = 65535
        delete(v1, (fi7 >> ((v2) & 63)))
        v8 := GenS1{F0: 4294967295, F1: true, F2: 26}
        _ = fi7
        _ = v8
    }
    v1[2147483647] = "a"
    v3 = append(v3, 834)
    var v9 int16 = 0
    v2 = v2
    v3 = append(v3, 273)
    if (false && true) {
        v9 = v9
        fmt.Println(v9, v9)
    } else {
        if false {
            fmt.Println(v9, v9)
            switch "" {
            case "":
                if ((((((1.5 - (1e10 * -2.25)) / 1.0) != ((1e10 / 3.14159) / float32((1e10 * 0.1)))) && true) && true) && false) {
                    v9 += 1000
                    v2 = (v2 & 42)
                }
            case "foo bar":
                switch -705 {
                case 32767:
                    v1[0] = "foo bar"
                case -386:
                    v9 += v9
                case -294:
                    var v10 uint64 = 0
                    _ = v10
                default:
                    var v11 uint32 = 875
                    _ = v11
                }
            case "\n\t":
                switch "a" {
                case "\n\t":
                    if false {
                        fmt.Println(v9, v9)
                    }
                case "foo bar":
                    for fi12 := 0; fi12 < 19; fi12++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v9 -= -371
                        _ = fi12
                    }
                case "a":
                    for fi13 := 0; fi13 < 1; fi13++ {
                        fmt.Println(v2, fi13, v9)
                        _ = fi13
                    }
                }
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
