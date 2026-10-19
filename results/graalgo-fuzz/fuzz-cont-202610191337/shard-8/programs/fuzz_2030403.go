package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 uint16 }
func genFn0(p0 bool, p1 int16) int32 {
    v5 := []int{}
    p1 += 0
    _ = p0
    _ = p1
    _ = v5
    return -801
}

func main() {
    var v1 uint16 = 255
    var v2 uint64 = 2
    var v3 uint = 431
    v4 := []int{}
    v2 = (v2 & (^(v2)))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v6 int64 = 9223372036854775807
    if false {
        if len(v4) > 0 {
            v4[0] = -1
        }
        for fi7 := 0; fi7 < 12; fi7++ {
            if (!(false || false)) {
                fmt.Println(v2, v3, fi7)
                v8 := GenS1{F0: (v6 >> ((4294967295) & 63))}
                var v9 uint32 = 2147483647
                _ = v8
                _ = v9
            } else {
                for fi10 := 0; fi10 < 6; fi10++ {
                    if (!true) {
                        switch "\n\t" {
                        case "":
                            _ = len(v4)
                        }
                        v6 = 9223372036854775807
                    }
                    v3 -= 419
                    _ = fi10
                }
            }
            if false {
                v6 += v6
            }
            v1 -= v1
            _ = fi7
        }
        v11 := GenS1{F0: v6}
        _ = v11
    }
    v12 := GenS1{F0: 1}
    if len(v4) > 0 {
        v4[0] = -2147483648
    }
    var v13 uint32 = 689
    _ = len(v4)
    fmt.Println(v6, v6, v13)
    switch ("foo bar" + "a") {
    case "a":
        v14 := GenS2{F0: ("hello" != "hello"), F1: 255}
        _ = v14
    case "":
        if ("\n\t" != "foo bar") {
            if ("\n\t" == ("" + "")) {
                v15 := GenS1{F0: (^(v6))}
                v6 = (v6 | (v6 + v6))
                _ = v15
            } else {
                var v16 float64 = 3.14159
                _ = v16
            }
        } else {
            var v17 string = ("" + "\n\t")
            _ = v17
        }
    case "foo bar":
        var v18 int64 = v6
        _ = v18
    }
    v4 = append(v4, -518)
    fmt.Println(v1, v3, v3)
    v1 = uint16(v3)
    v13 = v13
    var v19 uint8 = 127
    if true {
        _ = len(v4)
        var v20 float64 = 0.1
        _ = v20
    } else {
        v6 += v6
        var v21 bool = ("\n\t" > ("a" + "foo bar"))
        v3 -= v3
        _ = v21
    }
    v3 = 512
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
