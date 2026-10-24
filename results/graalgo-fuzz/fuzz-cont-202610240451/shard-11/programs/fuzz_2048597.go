package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int }

func main() {
    v1 := []int{127, 127}
    v2 := []string{(("a" + "foo bar") + ""), (((("a" + ("hello" + "hello")) + "a") + ("" + "\n\t")) + "hello"), "a"}
    var v3 uint16 = 255
    for fi4 := 0; fi4 < 4; fi4++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if (!true) {
            var v5 uint = 945
            v3 = (v3 * v3)
            var v6 int64 = 42
            _ = v5
            _ = v6
        } else {
            var v7 string = ("\n\t" + "foo bar")
            for fi8 := 0; fi8 < 3; fi8++ {
                var v9 uint8 = 42
                _ = fi8
                _ = v9
            }
            _ = v7
        }
        switch "a" {
        case "foo bar":
            v10 := GenS1{F0: 100}
            _ = v10
        case "hello":
            if (!false) {
                v11 := GenS2{F0: fi4}
                v3 -= (v3 ^ 255)
                _ = v11
            } else {
                for fi12 := 0; fi12 < 14; fi12++ {
                    var v13 float64 = -2.25
                    v14 := GenS1{F0: -128}
                    v15 := GenS1{F0: 100}
                    _ = fi12
                    _ = v13
                    _ = v14
                    _ = v15
                }
                if true {
                    v3 = v3
                    v3 -= v3
                    var v16 uint16 = (^(v3))
                    _ = v16
                }
                var v17 float64 = -2.25
                _ = v17
            }
        case "a":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = fi4
    }
    var v18 string = (("" + "hello") + (("a" + (("hello" + "") + "\n\t")) + (("a" + "a") + "hello")))
    fmt.Println(v3, v3)
    var v19 int16 = -1
    var v20 uint8 = 40
    if true {
        if false {
            v18 = "\n\t"
            if false {
                var v21 int8 = 42
                _ = v21
            } else {
                var v22 int8 = -128
                var v23 float32 = 0.0
                var v24 bool = ((("" <= v18) || ((uint64(v20) > 906) || true)) && true)
                _ = v22
                _ = v23
                _ = v24
            }
        }
    }
    v25 := GenS2{F0: 42}
    if true {
        v18 = v18
    } else {
        switch 32767 {
        case 9223372036854775807:
            for fi26 := 0; fi26 < 19; fi26++ {
                _ = len(v2)
                if (false || (!(-1 > 9223372036854775807))) {
                    switch (fi26 - 2147483647) {
                    case 9223372036854775807:
                        if (2147483647 >= 0) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        } else {
                            if (v18 < v18) {
                                v20 -= 74
                                v3 += (^(v3))
                            }
                            v27 := GenS1{F0: 127}
                            if len(v1) > 0 {
                                v1[0] = fi26
                            }
                            _ = v27
                        }
                    case 2147483647:
                        var v28 int8 = -128
                        _ = v28
                    case -9223372036854775808:
                        switch 906 {
                        case -9223372036854775808:
                            var v29 int64 = -1
                            _ = v29
                        case 2147483647:
                            for fi30 := 0; fi30 < 16; fi30++ {
                                v1 = append(v1, -376)
                                if len(v1) > 0 {
                                    v1[0] = (fi30 & fi30)
                                }
                                _ = fi30
                            }
                        default:
                            var v31 bool = (-1 != 389)
                            _ = v31
                        }
                    }
                    v18 = (((v18 + "foo bar") + "\n\t") + v18)
                }
                v20 += v20
                _ = fi26
            }
        case -2147483648:
            fmt.Println(v19, v18, v20)
        case 206:
            v19 = -120
        }
        for fi32 := 0; fi32 < 10; fi32++ {
            v3 -= uint16(v20)
            var v33 string = v18
            v3 -= v3
            _ = fi32
            _ = v33
        }
    }
    var v34 uint32 = 378
    var v35 int8 = 1
    _ = len(v2)
    _ = len(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v34
    _ = v35
}
