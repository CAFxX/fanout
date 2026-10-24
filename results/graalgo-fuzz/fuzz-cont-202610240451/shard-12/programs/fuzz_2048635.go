package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 uint32; F2 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8 }

func main() {
    var v1 int8 = 17
    var v2 float64 = (3.14159 / 1.5)
    v3 := []string{"a", ("hello" + "a")}
    var v4 uint16 = 42
    v5 := make(map[string]int)
    var v6 int8 = (v1 | (v1 ^ 42))
    v1 += (v1 - 100)
    if false {
        for fi7 := 0; fi7 < 18; fi7++ {
            v8 := GenS1{F0: 255}
            v9 := v5["hello"]
            v10 := GenS2{F0: -39}
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        v3 = append(v3, "hello")
        for fi11 := 0; fi11 < 1; fi11++ {
            if (!true) {
                var v12 uint8 = 127
                _ = v12
            } else {
                fmt.Println(v4, v6)
                v13 := GenS3{F0: v1, F1: 2147483648, F2: fi11}
                _ = v13
            }
            _ = fi11
        }
    } else {
        if (((((!true) || (false && (1 > 127))) && false) || (false && ("\n\t" <= ("" + "\n\t")))) && ((32768 >= 32768) || false)) {
            if false {
                var v14 int32 = 2147483647
                var v15 uint16 = 0
                _ = v14
                _ = v15
            }
            switch 107 {
            case 32767:
                v4 -= 255
            case -1:
                v16 := GenS4{F0: 1}
                _ = v16
            }
            v4 -= 2
        } else {
            if (32767 <= 42) {
                v17 := GenS1{F0: 0}
                v5[("\n\t" + "")] = 128
                if ("foo bar" >= "\n\t") {
                    switch ("\n\t" + "\n\t") {
                    case "":
                        for fi18 := 0; fi18 < 18; fi18++ {
                            fmt.Println(v4)
                            if (!(float64((((3.14159 / (float32(v2) * 0.1)) + 3.14159) + 0.1)) == (v2 / v2))) {
                                switch ("a" + "hello") {
                                case "foo bar":
                                    if ((v2 / v2) < 0.1) {
                                        fmt.Println(v4)
                                    }
                                case "":
                                    if true {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = len(v3)
                                    }
                                }
                                for fi19 := 0; fi19 < 4; fi19++ {
                                    v20 := GenS1{F0: 158}
                                    var v21 string = "a"
                                    _ = fi19
                                    _ = v20
                                    _ = v21
                                }
                            } else {
                                var v22 float32 = 0.1
                                _ = v22
                            }
                            var v23 int64 = 9223372036854775807
                            _ = fi18
                            _ = v23
                        }
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                _ = v17
            }
        }
        var v24 float32 = 0.1
        _ = v24
    }
    v2 -= v2
    fmt.Println(v2)
    var v25 int = 4294967295
    v26 := GenS1{F0: 0}
    var v27 uint = 255
    if (uint(v4) <= v27) {
        fmt.Println(v6, v1)
        _ = len(v5)
    }
    if true {
        switch "" {
        case "a":
            if false {
                v3 = append(v3, "")
                v28 := GenS5{F0: v1}
                _ = v28
            } else {
                var v29 uint8 = 1
                if false {
                    fmt.Println(v25)
                } else {
                    var v30 int = v25
                    _ = v30
                }
                _ = v29
            }
        }
        var v31 int32 = -1
        _ = v31
    } else {
        for fi32 := 0; fi32 < 9; fi32++ {
            var v33 int = (fi32 / (((fi32 % (((v25 >> ((v27) & 63))) | 1))) | 1))
            _ = fi32
            _ = v33
        }
        switch "\n\t" {
        case "\n\t":
            var v34 uint32 = 43
            _ = v34
        default:
            var v35 int = -2147483648
            _ = v35
        }
        if false {
            for fi36 := 0; fi36 < 1; fi36++ {
                v2 -= (v2 * ((-2.25 / v2) + v2))
                _ = fi36
            }
            v27 += uint(v4)
        } else {
            v6 += (v1 * v1)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
