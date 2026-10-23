package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint8 }
func genFn0() bool {
    for fi7 := 0; fi7 < 2; fi7++ {
        if ((65535 >= 0) && (42 > 2)) {
            for fi8 := 0; fi8 < 13; fi8++ {
                v9 := []int{(fi7 >> ((2) & 63)), (fi7 * (^(fi8))), fi7}
                _ = fi8
                _ = v9
            }
            var v10 string = ""
            _ = v10
        }
        v11 := GenS1{F0: 229, F1: (((((("foo bar" + "a") + "a") + "") + ("hello" + (("" + "\n\t") + "foo bar"))) + ("hello" + "")) + "\n\t"), F2: 1}
        var v12 int64 = 42
        _ = fi7
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 float32 = 1.5
        v14 = 0.0
        _ = fi13
        _ = v14
    }
    v15 := []int{128, 128, -9223372036854775808}
    var v16 uint32 = 2147483648
    _ = v15
    _ = v16
    return (!false)
}

func main() {
    v1 := []string{}
    var v2 uint8 = 2
    var v3 int8 = 106
    v4 := []int{2147483647, 32767}
    var v5 float64 = ((((3.14159 / (0.0 + -2.25)) - 1.5) * 0.1) / 0.1)
    v6 := []uint32{}
    for fi17 := 0; fi17 < 5; fi17++ {
        v5 = ((-2.25 - v5) + v5)
        v4 = append(v4, int(v3))
        switch (("" + (("a" + "") + ("" + (("foo bar" + "") + "hello")))) + ("\n\t" + "\n\t")) {
        case "":
            fmt.Println(fi17)
        }
        _ = fi17
    }
    fmt.Println(v2, v3)
    v5 -= (v5 * (0.0 - v5))
    if false {
        v3 = v3
    }
    fmt.Println(v2, v2)
    v5 += (v5 + (1.5 * (((v5 / v5) * v5) * (0.0 / ((0.1 - (v5 - v5)) * (-2.25 + (1.5 - (v5 * 0.0))))))))
    var v18 string = ""
    for fi19 := 0; fi19 < 6; fi19++ {
        for fi20 := 0; fi20 < 2; fi20++ {
            var v21 int8 = (v3 << ((255) & 63))
            v22 := GenS1{F0: 18446744073709551615, F1: v18, F2: (^(v2))}
            if true {
                var v23 int16 = 0
                fmt.Println(v23)
                fmt.Println(fi19, fi20)
                _ = v23
            }
            _ = fi20
            _ = v21
            _ = v22
        }
        for fi24 := 0; fi24 < 5; fi24++ {
            v5 = ((v5 / v5) + v5)
            _ = fi24
        }
        _ = fi19
    }
    v3 = (-(v3))
    if (0 > 2147483648) {
        v3 = (v3 | (^(v3)))
        fmt.Println(v2, v3, v18)
    } else {
        var v25 string = v18
        var v26 int = 2147483647
        _ = v25
        _ = v26
    }
    if false {
        var v27 uint32 = 0
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v27
    }
    for fi28 := 0; fi28 < 20; fi28++ {
        var v29 uint16 = 255
        _ = fi28
        _ = v29
    }
    fmt.Println(v2, v3)
    v18 = v18
    if true {
        v18 = (v18 + v18)
        var v30 bool = (42 >= 252)
        _ = v30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
