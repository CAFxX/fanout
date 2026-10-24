package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint8; F1 int; F2 bool }
func genFn0(p0 uint) uint64 {
    var v6 int32 = -1
    if (-75 >= -1) {
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 uint32 = 1
            p0 = 275
            fmt.Println(v6, v6)
            _ = fi7
            _ = v8
        }
        var v9 string = ""
        v10 := []int{127, 0}
        _ = v10
        _ = v9
    }
    var v11 float64 = 3.14159
    v12 := []int{-902, -352, -2147483648, -819}
    _ = p0
    _ = v11
    _ = v12
    _ = v6
    return 18446744073709551615
}
func genFn1(p0 int32, p1 uint) int {
    p0 += p0
    var v13 uint8 = 201
    _ = p0
    _ = p1
    _ = v13
    return 2147483647
}

func main() {
    var v1 int16 = -619
    var v2 bool = (-484 > 32767)
    var v3 int = -186
    v4 := []string{"", (("a" + "a") + ("hello" + (("\n\t" + ("a" + "\n\t")) + "\n\t"))), ""}
    var v5 float64 = 3.14159
    var v14 int8 = 100
    var v15 int8 = v14
    fmt.Println(v2, v15, v14)
    v4 = append(v4, "hello")
    switch genFn1(-976, 302) {
    case 9223372036854775807:
        v5 -= (v5 + ((((v5 + (3.14159 / 0.1)) * v5) - float64(1.5)) + (((v5 / v5) - v5) / (v5 - v5))))
    case -2147483648:
        for fi16 := 0; fi16 < 11; fi16++ {
            switch v3 {
            case 932:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case 42:
                var v17 int = genFn1(281, 2)
                _ = v17
            }
            v5 += ((v5 / v5) / v5)
            if (42 < 4294967295) {
                if len(v4) > 0 {
                    v4[0] = "foo bar"
                }
                var v18 int32 = 0
                v2 = (!("a" != "hello"))
                _ = v18
            }
            _ = fi16
        }
    case 566:
        if false {
            fmt.Println(v15, v3)
        } else {
            var v19 int = 128
            if v2 {
                v1 -= v1
                if ((-2.25 * 0.1) > 1.5) {
                    var v20 int = int(v1)
                    v3 += v19
                    var v21 uint16 = 1
                    _ = v20
                    _ = v21
                }
            } else {
                for fi22 := 0; fi22 < 14; fi22++ {
                    var v23 uint = 384
                    v4 = append(v4, ("\n\t" + "a"))
                    _ = fi22
                    _ = v23
                }
            }
            if (!v2) {
                fmt.Println(v3)
                v24 := GenS1{F0: 4294967295, F1: 0}
                for fi25 := 0; fi25 < 7; fi25++ {
                    var v26 string = ("hello" + "\n\t")
                    for fi27 := 0; fi27 < 5; fi27++ {
                        var v28 float64 = v5
                        _ = fi27
                        _ = v28
                    }
                    var v29 int16 = (v1 * v1)
                    _ = fi25
                    _ = v26
                    _ = v29
                }
                _ = v24
            }
            _ = v19
        }
    default:
        if true {
            var v30 int32 = 2147483647
            _ = v30
        } else {
            v2 = (4294967295 == genFn0(4294967295))
            _ = len(v4)
        }
    }
    v2 = (v2 || v2)
    v31 := GenS2{F0: 103, F1: genFn1(1, 698), F2: v2}
    if (v2 || (false && v2)) {
        var v32 int16 = v1
        _ = v32
    }
    var v33 int = genFn1(0, 255)
    switch "" {
    case "a":
        switch ("foo bar" + "foo bar") {
        case "hello":
            v2 = (true && false)
        case "":
            _ = len(v4)
        default:
            var v34 float32 = -2.25
            _ = v34
        }
    case "\n\t":
        v5 = v5
    case "foo bar":
        var v35 string = ((("hello" + "hello") + "") + "a")
        _ = v35
    }
    var v36 uint16 = 42
    v5 -= (-2.25 + 0.1)
    var v37 int8 = (v14 ^ v14)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
