package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }

func main() {
    var v1 bool = (-9223372036854775808 == 335)
    v2 := make(map[string]string)
    var v3 float64 = 1e10
    var v4 int32 = -2147483648
    fmt.Println(v4, v1, v1)
    v1 = false
    var v5 string = "a"
    v5 = (((v5 + "foo bar") + ("\n\t" + ((("" + v5) + (((("hello" + v5) + (v5 + v5)) + "a") + (("\n\t" + ("foo bar" + (v5 + v5))) + ((v5 + (v5 + "foo bar")) + "foo bar")))) + v5))) + "hello")
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 uint16 = 94
        switch v5 {
        case "":
            switch ("hello" + ("a" + v5)) {
            case "":
                switch fi6 {
                case 2147483647:
                    v1 = (!v1)
                case 1:
                    var v8 uint16 = (v7 ^ v7)
                    _ = v8
                default:
                    var v9 float32 = 1e10
                    _ = v9
                }
            case "foo bar":
                var v10 int8 = 127
                _ = v10
            case "\n\t":
                v2[v5] = v5
            }
        case "a":
            _ = len(v2)
        case "hello":
            v5 = v5
        }
        _ = fi6
        _ = v7
    }
    v5 = v5
    switch v5 {
    case "hello":
        var v11 int16 = -800
        _ = v11
    case "":
        var v12 uint32 = 2147483648
        _ = v12
    }
    var v13 float64 = 1e10
    v14 := []int{128, 0, -1}
    if (v1 && (v1 && (1000 != 42))) {
        v5 = ("" + "\n\t")
        v13 += v13
        for fi15 := 0; fi15 < 8; fi15++ {
            var v16 uint64 = 2
            if false {
                if v1 {
                    switch fi15 {
                    case 4294967295:
                        v1 = (((!true) && true) && ((((v5 + v5) + v5) + "") > v5))
                    case 128:
                        if v1 {
                            switch (fi15 * fi15) {
                            case 42:
                                v17 := GenS1{F0: -9223372036854775808, F1: -1}
                                _ = v17
                            case 127:
                                v4 += v4
                            case -74:
                                fmt.Println(v16, v13, v4)
                            }
                            var v18 string = v5
                            _ = v18
                        }
                    }
                }
                if true {
                    var v19 bool = v1
                    _ = v19
                }
            } else {
                var v20 int8 = 86
                v3 -= (v3 * (v13 * v13))
                if (160 >= 0) {
                    if v1 {
                        v5 = ((v5 + "") + "foo bar")
                        _ = len(v2)
                    }
                    v5 = ("\n\t" + "foo bar")
                }
                _ = v20
            }
            _ = fi15
            _ = v16
        }
    } else {
        v21 := GenS1{F0: -955, F1: -1}
        fmt.Println(v4, v5, v13)
        _ = v21
    }
    var v22 uint = 1
    for fi23 := 0; fi23 < 8; fi23++ {
        var v24 int16 = 1
        v5 = v5
        _ = fi23
        _ = v24
    }
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    var v25 int8 = 58
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
