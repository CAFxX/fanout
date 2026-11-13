package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
func genFn0(p0 uint8) int16 {
    for fi4 := 0; fi4 < 20; fi4++ {
        var v5 float64 = 3.14159
        _ = fi4
        _ = v5
    }
    if (448 < 18446744073709551615) {
        p0 -= p0
        var v6 uint16 = 255
        p0 -= p0
        _ = v6
    } else {
        p0 = (p0 >> ((42) & 63))
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 float32 = 1.5
            _ = fi7
            _ = v8
        }
        p0 = (p0 + p0)
    }
    _ = p0
    return 42
}
func genFn1(p0 string, p1 uint, p2 int16) int32 {
    p0 = (p0 + p0)
    var v9 int16 = p2
    fmt.Println(p0, v9, p0)
    for fi10 := 0; fi10 < 6; fi10++ {
        if (!true) {
            p2 += (v9 & -805)
            var v11 uint32 = 2147483648
            v11 = v11
            _ = v11
        } else {
            v12 := []int{(fi10 % ((fi10) | 1))}
            _ = v12
        }
        _ = fi10
    }
    fmt.Println(p1, p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 2147483647
}

func main() {
    v1 := []int{}
    v2 := make(map[string]int)
    v3 := []string{}
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 int64 = int64(fi13)
        _ = fi13
        _ = v14
    }
    var v15 string = ""
    v15 = v15
    v2[(v15 + v15)] = -1
    fmt.Println(v15)
    fmt.Println(v15)
    delete(v2, (("\n\t" + "foo bar") + (v15 + ("hello" + ("" + v15)))))
    switch ("" + v15) {
    case "foo bar":
        fmt.Println(v15)
    case "":
        var v16 uint64 = 840
        _ = v16
    }
    if (!true) {
        v15 = ("a" + (v15 + "a"))
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        for fi18 := 0; fi18 < 10; fi18++ {
            var v19 bool = true
            v19 = (1 != 1)
            if ((v19 || (42 == 18446744073709551615)) && ((0.1 / -2.25) == (1.5 + ((1.5 - (float64((0.1 * float32(1e10))) / 1e10)) + 0.0)))) {
                switch 157 {
                case -1:
                    var v20 int64 = 0
                    _ = v20
                default:
                    v15 = "hello"
                }
                switch v15 {
                case "foo bar":
                    v19 = (!((v19 && ((-2.25 <= 3.14159) && v19)) || v19))
                case "":
                    v19 = (3.14159 >= float32(-2.25))
                }
            }
            _ = fi18
            _ = v19
        }
        v15 = ((v15 + (v15 + "\n\t")) + v15)
        var v21 uint = 144
        _ = fi17
        _ = v21
    }
    var v22 int32 = 273
    switch v15 {
    case "":
        var v23 float64 = 3.14159
        _ = v23
    case "foo bar":
        delete(v2, (v15 + ""))
    case "\n\t":
        for fi24 := 0; fi24 < 14; fi24++ {
            switch (v15 + ("foo bar" + v15)) {
            case "a":
                fmt.Println(v22)
            case "":
                if (403 >= 2) {
                    if false {
                        if true {
                            if len(v3) > 0 {
                                v3[0] = ("hello" + v15)
                            }
                        } else {
                            var v25 uint64 = 1
                            fmt.Println(v15)
                            _ = v25
                        }
                    }
                    v22 = v22
                } else {
                    v15 = v15
                }
            }
            _ = fi24
        }
    }
    var v26 string = (v15 + v15)
    v27 := GenS1{F0: 18446744073709551615, F1: 65535}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v3
}
