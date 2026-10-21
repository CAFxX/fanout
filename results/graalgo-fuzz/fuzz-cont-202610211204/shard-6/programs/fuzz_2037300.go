package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 uint }

func main() {
    var v1 int16 = 1000
    var v2 float64 = -2.25
    var v3 uint = 65535
    var v4 float64 = (v2 * v2)
    var v5 string = ("foo bar" + "a")
    var v6 int64 = 9223372036854775807
    v4 -= 3.14159
    for fi7 := 0; fi7 < 14; fi7++ {
        v5 = v5
        fmt.Println(v2)
        v8 := GenS1{F0: "", F1: 42}
        _ = fi7
        _ = v8
    }
    var v9 int16 = (^(v1))
    if false {
        v10 := []int{2147483647, 786}
        for fi11 := 0; fi11 < 12; fi11++ {
            for fi12 := 0; fi12 < 17; fi12++ {
                var v13 bool = ((!true) && (-9223372036854775808 < -1))
                v4 = v4
                var v14 int = (fi12 * int(v6))
                var v15 uint16 = 42
                switch ("" + v5) {
                case "\n\t":
                    var v16 int64 = -1
                    _ = v16
                case "a":
                    switch "hello" {
                    case "a":
                        if v13 {
                            v17 := GenS2{F0: (("hello" + (v5 + v5)) + "foo bar")}
                            _ = v17
                        } else {
                            var v18 uint32 = 0
                            v19 := GenS2{F0: "foo bar"}
                            if ((v13 && (v15 == 1)) || v13) {
                                v20 := GenS2{F0: "\n\t"}
                                v21 := GenS2{F0: (((v5 + "hello") + "") + v5)}
                                v22 := GenS3{F0: (v3 >> ((v3) & 63)), F1: v3}
                                switch ((v5 + ((v5 + v5) + v5)) + "foo bar") {
                                case "hello":
                                    v15 -= 0
                                default:
                                    var v23 uint32 = (v18 ^ 2)
                                    _ = v23
                                }
                                _ = v20
                                _ = v21
                                _ = v22
                            } else {
                                var v24 int64 = v6
                                _ = v24
                            }
                            _ = v18
                            _ = v19
                        }
                    case "foo bar":
                        v15 = 42
                    }
                }
                v15 -= (v15 * v15)
                fmt.Println(v2, v15, v4)
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
            if len(v10) > 0 {
                v10[0] = fi11
            }
            switch v5 {
            case "hello":
                for fi25 := 0; fi25 < 4; fi25++ {
                    v3 -= 255
                    _ = fi25
                }
            case "foo bar":
                switch 305 {
                case 4294967295:
                    fmt.Println(fi11)
                case 0:
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                case 127:
                    if false {
                        var v26 uint16 = 2
                        _ = v26
                    } else {
                        var v27 string = "foo bar"
                        var v28 float64 = ((v4 / (float64((1e10 * 3.14159)) * (v4 - (v4 - -2.25)))) - -2.25)
                        fmt.Println(v1, v5, v1)
                        _ = v27
                        _ = v28
                    }
                }
            }
            _ = fi11
        }
        _ = v10
    } else {
        var v29 uint16 = 255
        _ = v29
    }
    for fi30 := 0; fi30 < 4; fi30++ {
        var v31 bool = (v6 == v6)
        v9 -= v9
        _ = fi30
        _ = v31
    }
    v6 = v6
    fmt.Println(v6)
    var v32 uint32 = 42
    v33 := GenS3{F0: (v3 * v3), F1: uint(v32)}
    var v34 int64 = v6
    v4 += ((0.1 - v2) - v2)
    v5 = v5
    v4 += v2
    v35 := GenS1{F0: "foo bar", F1: 569}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
