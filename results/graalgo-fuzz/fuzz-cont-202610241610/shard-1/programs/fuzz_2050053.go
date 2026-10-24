package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 uint8; F1 int64; F2 int }
func genFn0(p0 int8, p1 uint8) uint {
    fmt.Println(p0, p1)
    v7 := GenS1{F0: 44, F1: "hello"}
    _ = p0
    _ = p1
    _ = v7
    return 255
}

func main() {
    var v1 int8 = 100
    var v2 int64 = 42
    var v3 uint8 = 127
    v4 := make(map[string]uint32)
    v5 := make(map[int]int)
    var v6 int16 = -32768
    var v8 uint8 = (v3 * v3)
    v3 -= 255
    v6 = -32768
    v5[0] = -9223372036854775808
    var v9 float32 = (1.5 + (1.5 / ((-2.25 - float32(0.1)) * 0.1)))
    var v10 float64 = ((0.1 - 0.1) * 1e10)
    var v11 uint = 2
    v12 := GenS1{F0: 94, F1: ("" + ("\n\t" + ""))}
    delete(v4, ("a" + (("hello" + (("\n\t" + "a") + "\n\t")) + "hello")))
    switch (("hello" + "") + "") {
    case "foo bar":
        var v13 uint = (v11 * v11)
        _ = v13
    case "":
        var v14 int32 = 42
        _ = v14
    default:
        for fi15 := 0; fi15 < 9; fi15++ {
            for fi16 := 0; fi16 < 12; fi16++ {
                fmt.Println(v2, v1)
                v3 -= v3
                _ = fi16
            }
            v17 := []int{1, fi15, fi15}
            if true {
                if false {
                    var v18 bool = (4294967295 <= 4294967295)
                    fmt.Println(v1)
                    for fi19 := 0; fi19 < 10; fi19++ {
                        v17 = append(v17, fi15)
                        _ = fi19
                    }
                    _ = v18
                }
            } else {
                v10 = (v10 - v10)
                for fi20 := 0; fi20 < 19; fi20++ {
                    v11 = 18446744073709551615
                    if false {
                        v5[(fi15 | fi15)] = -2147483648
                    } else {
                        v21 := GenS1{F0: (^(v8)), F1: "hello"}
                        v22 := GenS1{F0: 255, F1: "foo bar"}
                        _ = v21
                        _ = v22
                    }
                    v10 = 1e10
                    _ = fi20
                }
            }
            _ = fi15
            _ = v17
        }
    }
    var v23 bool = (!true)
    var v24 int32 = -1
    var v25 float64 = v10
    for fi26 := 0; fi26 < 20; fi26++ {
        switch -416 {
        case 128:
            for fi27 := 0; fi27 < 5; fi27++ {
                if false {
                    var v28 int8 = -128
                    v29 := GenS1{F0: 13, F1: "hello"}
                    _ = v28
                    _ = v29
                }
                _ = fi27
            }
        case -109:
            v30 := GenS1{F0: v3, F1: ((("foo bar" + "") + ("a" + "a")) + ("hello" + (("\n\t" + "a") + "\n\t")))}
            _ = v30
        }
        _ = fi26
    }
    switch ("foo bar" + "hello") {
    case "\n\t":
        if v23 {
            for fi31 := 0; fi31 < 20; fi31++ {
                var v32 float32 = 1.5
                _ = fi31
                _ = v32
            }
            v2 = v2
        } else {
            delete(v5, -465)
        }
    case "foo bar":
        v4["hello"] = 65535
    default:
        v33 := GenS2{F0: v3, F1: v2, F2: -1}
        _ = v33
    }
    v23 = (476 == 18446744073709551615)
    v8 = v3
    var v34 float64 = (((v25 + v10) / float64((v9 - float32(v25)))) / v25)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
