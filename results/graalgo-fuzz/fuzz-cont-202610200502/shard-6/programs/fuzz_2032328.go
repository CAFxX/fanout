package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint64 }
func genFn0() int8 {
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 int16 = -204
        var v8 int = (fi6 | 2147483647)
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 int8 = 71
    if true {
        var v10 float64 = 1e10
        for fi11 := 0; fi11 < 15; fi11++ {
            v10 = (((float64(3.14159) / v10) - ((1.5 + float64(float32(v10))) / v10)) + (v10 - v10))
            v9 = v9
            _ = fi11
        }
        v12 := []int{32767, -868, -2147483648}
        _ = v10
        _ = v12
    }
    _ = v9
    return -128
}
func genFn1() string {
    var v13 uint64 = 18446744073709551615
    switch (("hello" + "\n\t") + "\n\t") {
    case "":
        v14 := []int{679}
        _ = v14
    case "foo bar":
        for fi15 := 0; fi15 < 12; fi15++ {
            var v16 float32 = 3.14159
            _ = fi15
            _ = v16
        }
    }
    if (127 >= 0) {
        v17 := []int{-675}
        for fi18 := 0; fi18 < 11; fi18++ {
            var v19 int = fi18
            _ = fi18
            _ = v19
        }
        _ = v17
    }
    var v20 string = ""
    _ = v13
    _ = v20
    return v20
}

func main() {
    var v1 uint = 4294967295
    var v2 int16 = 1000
    var v3 uint32 = 475
    v4 := make(map[string]uint32)
    var v5 int32 = -1
    v1 -= v1
    switch 0 {
    case 243:
        var v21 bool = (true && ("" != ("" + (("foo bar" + ("hello" + ("" + ("foo bar" + "hello")))) + "hello"))))
        _ = v21
    case 834:
        v3 -= uint32(v1)
    default:
        _ = len(v4)
    }
    v22 := GenS1{F0: (("foo bar" + "foo bar") + ("a" + (("a" + "\n\t") + "hello"))), F1: false, F2: 93}
    if false {
        v3 = v3
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        var v24 uint32 = uint32(v1)
        if (((((v2 / ((v2) | 1)) != v2) || (v3 >= 65535)) || true) && false) {
            if true {
                delete(v4, ("hello" + "a"))
                for fi25 := 0; fi25 < 5; fi25++ {
                    v3 += v3
                    _ = fi25
                }
            } else {
                var v26 int16 = -32768
                v27 := []int{(fi23 ^ fi23), (fi23 >> ((255) & 63))}
                var v28 int = fi23
                _ = v26
                _ = v27
                _ = v28
            }
        }
        if ((^(v5)) == v5) {
            var v29 int8 = 55
            fmt.Println(v2)
            _ = v29
        } else {
            if (!true) {
                if (v1 > v1) {
                    switch (fi23 & 680) {
                    case 32767:
                        switch "hello" {
                        case "":
                            if (2 >= 5) {
                                v30 := GenS1{F0: "hello", F1: (0.1 != ((0.1 - (-2.25 / (1.5 - 1e10))) - 3.14159)), F2: 2}
                                _ = v30
                            } else {
                                v4["hello"] = (v24 - (v24 ^ (v3 >> (((v1 & v1)) & 63))))
                            }
                        case "foo bar":
                            v31 := []int{9223372036854775807}
                            _ = v31
                        }
                    }
                }
                fmt.Println(v2)
            }
            if true {
                for fi32 := 0; fi32 < 14; fi32++ {
                    if true {
                        var v33 int64 = 1
                        _ = v33
                    } else {
                        v1 -= (v1 >> ((0) & 63))
                    }
                    if (((!((v1 - 0) < (v1 / (((v1 | 4294967295)) | 1)))) || (!(!(!true)))) && true) {
                        v4["foo bar"] = 42
                        v2 += -32768
                    } else {
                        v3 += (v24 << (((^(v1))) & 63))
                        var v34 float32 = 1e10
                        _ = v34
                    }
                    _ = fi32
                }
            }
        }
        _ = fi23
        _ = v24
    }
    if (true && ((!(!false)) && (-911 >= 504))) {
        for fi35 := 0; fi35 < 14; fi35++ {
            _ = len(v4)
            _ = fi35
        }
        v36 := []int{1, -788, 1}
        _ = v36
    }
    var v37 uint = (v1 << ((42) & 63))
    v3 = 42
    _ = len(v4)
    var v38 float64 = 0.1
    var v39 int64 = 1
    switch ("\n\t" + ("foo bar" + "\n\t")) {
    case "\n\t":
        v4[(("" + ("\n\t" + "foo bar")) + "")] = (v3 % ((v3) | 1))
    case "a":
        v39 = (v39 & (-(v39)))
    case "foo bar":
        if true {
            v40 := []int{-9223372036854775808, 2147483647, 1}
            _ = v40
        } else {
            v39 -= (v39 / (((v39 >> (((v37 - 42)) & 63))) | 1))
        }
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
