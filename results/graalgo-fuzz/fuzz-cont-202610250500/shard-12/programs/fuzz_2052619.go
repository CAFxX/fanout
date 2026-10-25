package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0() uint32 {
    switch "a" {
    case "hello":
        var v7 uint16 = 255
        _ = v7
    case "foo bar":
        var v8 uint32 = 2147483647
        _ = v8
    }
    v9 := []int{}
    _ = v9
    return 65535
}

func main() {
    var v1 int16 = 32767
    var v2 uint = 631
    v3 := []string{"a"}
    var v4 float32 = (0.0 + (3.14159 - -2.25))
    var v5 float64 = 1.5
    var v6 float64 = float64(1.5)
    v3 = append(v3, ("a" + ""))
    var v10 bool = false
    var v11 float32 = (v4 / float32((0.1 / ((v6 - 1e10) / v6))))
    switch "a" {
    case "foo bar":
        var v12 uint = (v2 % ((v2) | 1))
        _ = v12
    default:
        var v13 int = 42
        _ = v13
    }
    v5 += (v6 * (v5 - v6))
    v14 := GenS1{F0: 530}
    v15 := GenS1{F0: 42}
    for fi16 := 0; fi16 < 19; fi16++ {
        fmt.Println(v1, v2)
        v3 = append(v3, "foo bar")
        _ = fi16
    }
    v10 = ("foo bar" == ("foo bar" + "foo bar"))
    for fi17 := 0; fi17 < 6; fi17++ {
        v3 = append(v3, ("" + (("\n\t" + ("\n\t" + "\n\t")) + ("a" + ("\n\t" + "a")))))
        v6 = v5
        _ = len(v3)
        _ = fi17
    }
    for fi18 := 0; fi18 < 18; fi18++ {
        var v19 uint8 = 255
        fmt.Println(v2)
        v19 -= v19
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 14; fi20++ {
        for fi21 := 0; fi21 < 6; fi21++ {
            switch (-(fi21)) {
            case 1:
                v22 := GenS1{F0: genFn0()}
                _ = v22
            case -2147483648:
                var v23 int16 = (v1 | -32768)
                _ = v23
            }
            _ = fi21
        }
        var v24 int16 = int16(fi20)
        _ = fi20
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
