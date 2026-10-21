package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }

func main() {
    var v1 float32 = 1.5
    v2 := []string{"foo bar"}
    v3 := []string{"foo bar", "hello"}
    for fi4 := 0; fi4 < 2; fi4++ {
        var v5 string = "a"
        v1 -= ((v1 * (v1 / ((0.0 / v1) + v1))) - (v1 * v1))
        switch fi4 {
        case 32768:
            if false {
                var v6 uint32 = 1
                _ = len(v2)
                if len(v3) > 0 {
                    v3[0] = v5
                }
                _ = v6
            }
        case 128:
            for fi7 := 0; fi7 < 10; fi7++ {
                if len(v2) > 0 {
                    v2[0] = ("\n\t" + "")
                }
                if (42 <= 65535) {
                    v5 = v5
                    v8 := GenS1{F0: (v5 + v5), F1: -25}
                    _ = v8
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                _ = fi7
            }
        case 127:
            v5 = (v5 + (v5 + v5))
        }
        _ = fi4
        _ = v5
    }
    v1 = 3.14159
    v1 = ((1e10 - v1) - v1)
    var v9 uint16 = 360
    if true {
        fmt.Println(v1, v9)
    }
    var v10 int32 = -384
    var v11 uint8 = 2
    if false {
        v2 = append(v2, ("hello" + ("a" + ("" + "hello"))))
        if (-71 != -1) {
            var v12 int64 = 866
            _ = v12
        } else {
            var v13 int = 32768
            v11 -= 2
            _ = v13
        }
        fmt.Println(v9, v10, v11)
    } else {
        fmt.Println(v11)
    }
    v1 -= (float32(1e10) * float32(1e10))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v9
}
