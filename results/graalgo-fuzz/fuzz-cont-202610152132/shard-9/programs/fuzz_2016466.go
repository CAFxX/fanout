package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }

func main() {
    var v1 float64 = -2.25
    var v2 uint32 = 0
    var v3 int = 844
    var v4 float32 = 0.1
    var v5 int8 = 127
    v6 := []int{(v3 & v3), (v3 & v3), v3, (-(v3))}
    if false {
        v7 := GenS1{F0: (^(v2)), F1: "foo bar"}
        if (!((v2 % ((2) | 1)) > (v2 | (v2 % ((v2) | 1))))) {
            var v8 float32 = v4
            if true {
                v9 := GenS1{F0: v2, F1: "foo bar"}
                var v10 float64 = v1
                _ = v10
                _ = v9
            } else {
                var v11 uint = 65535
                _ = v11
            }
            _ = v8
        }
        _ = len(v6)
        _ = v7
    } else {
        var v12 float32 = 1.5
        if true {
            v13 := GenS1{F0: (^(v2)), F1: "\n\t"}
            _ = v13
        } else {
            if (true && ((v5 - (v5 << ((42) & 63))) >= (v5 | v5))) {
                switch ("hello" + "foo bar") {
                case "\n\t":
                    var v14 float64 = 1.5
                    _ = v14
                case "a":
                    fmt.Println(v5)
                }
                var v15 int32 = 42
                _ = v15
            }
        }
        _ = v12
    }
    if (23 == 189) {
        for fi16 := 0; fi16 < 8; fi16++ {
            v5 -= v5
            _ = fi16
        }
        for fi17 := 0; fi17 < 14; fi17++ {
            fmt.Println(fi17)
            v3 = -9223372036854775808
            v2 = v2
            _ = fi17
        }
        fmt.Println(v3)
    } else {
        if len(v6) > 0 {
            v6[0] = (^(v3))
        }
    }
    v6 = append(v6, (-(v3)))
    var v18 uint32 = (v2 & (v2 | (v2 & 0)))
    switch (("\n\t" + ("a" + "\n\t")) + "hello") {
    case "":
        if (((v1 - ((v1 + ((((-2.25 * v1) / ((v1 + v1) - (3.14159 / v1))) + 0.1) / 1.0)) * ((v1 / v1) + v1))) + v1) <= (v1 - float64(1e10))) {
            var v19 uint16 = 1
            if ((("a" + ("\n\t" + ("\n\t" + "hello"))) <= ("\n\t" + "foo bar")) || ((v18 << ((0) & 63)) < (^(v18)))) {
                if len(v6) > 0 {
                    v6[0] = (v3 ^ -9223372036854775808)
                }
                fmt.Println(v3, v1)
            }
            _ = v19
        } else {
            var v20 uint8 = 0
            v3 -= v3
            v21 := GenS1{F0: v18, F1: "a"}
            _ = v20
            _ = v21
        }
    case "a":
        if len(v6) > 0 {
            v6[0] = v3
        }
    default:
        fmt.Println(v1)
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        _ = len(v6)
        v3 = fi22
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi22
    }
    v23 := GenS1{F0: (^(v2)), F1: (("hello" + ("" + "hello")) + "foo bar")}
    for fi24 := 0; fi24 < 2; fi24++ {
        for fi25 := 0; fi25 < 17; fi25++ {
            v26 := GenS1{F0: v18, F1: ""}
            fmt.Println(v4, fi25, fi25)
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    if true {
        if (!true) {
            v1 = -2.25
        }
        v1 += ((v1 - v1) + float64((v4 - (v4 - (v4 / v4)))))
        switch "" {
        case "a":
            v5 += int8(v3)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
