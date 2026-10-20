package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 int64 = 0
    v2 := []int{42, 0, 717}
    var v3 string = ("foo bar" + "foo bar")
    var v4 float64 = (0.0 / float64(float32((-2.25 - (0.1 / 1.0)))))
    var v5 string = ""
    var v6 int = 32767
    var v7 uint16 = 2
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v8 uint8 = 1
    v9 := GenS1{F0: -32768, F1: (v5 + (("\n\t" + (("hello" + (v5 + v5)) + "")) + "hello"))}
    if false {
        v4 -= (1e10 / ((v4 / v4) - (((v4 + (float64(float32((v4 / v4))) * v4)) + float64(1e10)) * 1e10)))
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi10 := 0; fi10 < 6; fi10++ {
            for fi11 := 0; fi11 < 17; fi11++ {
                fmt.Println(v1, v7, v8)
                var v12 string = "foo bar"
                fmt.Println(v3, v4)
                _ = fi11
                _ = v12
            }
            for fi13 := 0; fi13 < 9; fi13++ {
                if false {
                    v7 += (v7 * v7)
                    v14 := GenS2{F0: 1}
                    fmt.Println(v7)
                    _ = v14
                } else {
                    switch ("a" + "\n\t") {
                    case "\n\t":
                        var v15 uint16 = 42
                        _ = v15
                    }
                    v7 -= 366
                }
                _ = len(v2)
                if len(v2) > 0 {
                    v2[0] = fi10
                }
                _ = fi13
            }
            v3 = v3
            _ = fi10
        }
        v2 = append(v2, v6)
    }
    fmt.Println(v7, v3, v5)
    var v16 uint64 = 2
    v7 = v7
    var v17 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
