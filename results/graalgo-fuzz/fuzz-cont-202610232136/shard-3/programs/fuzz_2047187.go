package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 uint32 }

func main() {
    var v1 int = 1
    v2 := []string{("\n\t" + ((("" + ("hello" + (("\n\t" + "foo bar") + ("hello" + "\n\t")))) + ("" + "\n\t")) + (("a" + "a") + "hello"))), "a", "foo bar", "foo bar"}
    var v3 bool = true
    v4 := make(map[int]string)
    var v5 float64 = (-2.25 - 0.1)
    if false {
        var v6 int8 = 1
        _ = v6
    } else {
        v4[128] = ("" + "")
        switch (v1 + v1) {
        case 128:
            v7 := GenS1{F0: "hello"}
            _ = v7
        }
    }
    var v8 uint = 65535
    if true {
        _ = len(v2)
    }
    var v9 uint = v8
    var v10 uint = (v9 % ((v8) | 1))
    if v3 {
        v10 = v8
        for fi11 := 0; fi11 < 11; fi11++ {
            v12 := GenS2{F0: -33, F1: 697}
            _ = fi11
            _ = v12
        }
    } else {
        for fi13 := 0; fi13 < 1; fi13++ {
            var v14 float32 = -2.25
            var v15 uint = v10
            if false {
                v5 -= (0.0 - v5)
            }
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    if v3 {
        if (42 < -564) {
            fmt.Println(v9)
            v10 += 255
            fmt.Println(v10, v3)
        } else {
            v8 = 4294967295
            v9 = (v8 * 255)
        }
        var v16 int16 = -81
        for fi17 := 0; fi17 < 10; fi17++ {
            v2 = append(v2, ("foo bar" + ""))
            for fi18 := 0; fi18 < 2; fi18++ {
                v3 = (-1 == 951)
                _ = fi18
            }
            for fi19 := 0; fi19 < 14; fi19++ {
                var v20 string = ("foo bar" + "\n\t")
                if (v16 > v16) {
                    var v21 uint8 = 255
                    _ = v21
                } else {
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                }
                _ = fi19
                _ = v20
            }
            _ = fi17
        }
        _ = v16
    } else {
        v1 = (v1 / ((-2147483648) | 1))
        if (!false) {
            if len(v2) > 0 {
                v2[0] = "hello"
            }
            if v3 {
                v5 += (v5 - ((3.14159 / (1.5 + 1.5)) / v5))
            } else {
                fmt.Println(v3, v3)
            }
        }
    }
    var v22 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
