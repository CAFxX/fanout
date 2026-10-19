package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint16, p1 uint32) bool {
    fmt.Println(p0, p0)
    var v5 uint = 1
    fmt.Println(v5)
    switch -2147483648 {
    case 127:
        var v6 int16 = 32767
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v5
    return ("a" < "foo bar")
}

func main() {
    var v1 uint16 = 0
    v2 := []string{"", "hello", "", (("hello" + "a") + "foo bar")}
    v3 := make(map[int]bool)
    var v4 float64 = (-2.25 * float64((1.5 - -2.25)))
    var v7 float32 = 1e10
    _ = len(v3)
    for fi8 := 0; fi8 < 7; fi8++ {
        for fi9 := 0; fi9 < 2; fi9++ {
            if (1 < 1) {
                var v10 uint8 = 125
                _ = len(v3)
                _ = len(v3)
                _ = v10
            } else {
                v11 := GenS1{F0: (-1 < -887)}
                if len(v2) > 0 {
                    v2[0] = "\n\t"
                }
                _ = v11
            }
            v1 += (v1 ^ (^(v1)))
            v7 = (v7 - v7)
            _ = fi9
        }
        _ = fi8
    }
    v12 := GenS1{F0: true}
    switch 866 {
    case -894:
        if false {
            for fi13 := 0; fi13 < 19; fi13++ {
                if true {
                    v14 := GenS2{F0: -688}
                    _ = v14
                } else {
                    if len(v2) > 0 {
                        v2[0] = "foo bar"
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v7 += 3.14159
                }
                v7 = float32(1.5)
                _ = fi13
            }
        }
    default:
        if (4294967295 != 2147483647) {
            v1 = 843
            fmt.Println(v1)
        }
    }
    v1 = (^(v1))
    switch 4294967295 {
    case 523:
        fmt.Println(v4, v7)
    case 32767:
        switch 718 {
        case 128:
            v4 = v4
        case -9223372036854775808:
            v4 = v4
        case 425:
            var v15 int = -2147483648
            _ = v15
        }
    default:
        v1 += v1
    }
    var v16 int16 = -533
    v4 -= (v4 * -2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
