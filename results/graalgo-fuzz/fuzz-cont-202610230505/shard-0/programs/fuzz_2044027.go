package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 uint8; F1 int }
func genFn0(p0 int32, p1 uint) int8 {
    if false {
        p0 = (p0 % ((p0) | 1))
        for fi6 := 0; fi6 < 12; fi6++ {
            v7 := []int{}
            switch (("foo bar" + "a") + ("\n\t" + "a")) {
            case "\n\t":
                switch (fi6 % ((fi6) | 1)) {
                case 119:
                    var v8 int8 = 1
                    _ = v8
                case -1:
                    var v9 uint = 1
                    _ = v9
                }
            case "hello":
                p0 -= (p0 | (-(p0)))
            }
            _ = fi6
            _ = v7
        }
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        p0 -= (p0 | (p0 - 2147483647))
        _ = fi10
    }
    p0 = 528
    _ = p0
    _ = p1
    return -1
}
func genFn1() uint16 {
    if ((!(0.0 == (0.1 * (1e10 - (((0.1 + (3.14159 / (0.1 + (1.5 * 0.0)))) / 3.14159) + 1e10))))) && false) {
        if true {
            var v11 float64 = float64((1.5 / 1.0))
            fmt.Println(v11)
            v11 += 3.14159
            _ = v11
        } else {
            var v12 string = "foo bar"
            _ = v12
        }
    } else {
        for fi13 := 0; fi13 < 10; fi13++ {
            fmt.Println(fi13)
            v14 := []int{-440, (fi13 * 0), fi13}
            _ = fi13
            _ = v14
        }
        v15 := []int{}
        var v16 float64 = (float64(0.1) - 1.5)
        _ = v15
        _ = v16
    }
    if false {
        var v17 uint = 255
        v18 := []int{-1, 128, 127}
        v17 = v17
        _ = v17
        _ = v18
    }
    if false {
        v19 := []int{-932, -2147483648}
        if len(v19) > 0 {
            v19[0] = 191
        }
        var v20 int = 996
        _ = v19
        _ = v20
    } else {
        if ((((0.1 + -2.25) / 1.0) >= 1e10) || (false || (1 < 2147483647))) {
            var v21 string = ("hello" + "foo bar")
            var v22 uint8 = 1
            _ = v21
            _ = v22
        }
        for fi23 := 0; fi23 < 7; fi23++ {
            var v24 int = fi23
            _ = fi23
            _ = v24
        }
    }
    return 286
}

func main() {
    var v1 uint8 = 42
    var v2 float64 = 3.14159
    var v3 uint32 = 0
    var v4 bool = (-717 < -1)
    v5 := []string{"a", "foo bar", (("a" + ("foo bar" + (("\n\t" + "a") + "\n\t"))) + ("foo bar" + (("foo bar" + (((("foo bar" + ("a" + "foo bar")) + ("foo bar" + "")) + (("hello" + "foo bar") + "\n\t")) + "a")) + ("a" + ("hello" + "foo bar")))))}
    var v25 string = ("" + "foo bar")
    var v26 float64 = ((-2.25 - 0.1) / (v2 / v2))
    v27 := GenS1{F0: 9223372036854775807}
    var v28 string = "hello"
    var v29 int32 = 42
    for fi30 := 0; fi30 < 12; fi30++ {
        v2 += (1e10 - v2)
        var v31 uint16 = 2
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 7; fi32++ {
        v33 := GenS1{F0: fi32}
        if (-1 == v29) {
            v34 := GenS1{F0: fi32}
            switch v25 {
            case "":
                v35 := GenS1{F0: (fi32 ^ 831)}
                _ = v35
            }
            _ = v34
        } else {
            var v36 uint16 = 255
            v37 := GenS2{F0: 763}
            if ((!(0 < 0)) && (!false)) {
                fmt.Println(v4, fi32)
                fmt.Println(v3, v2)
            }
            _ = v36
            _ = v37
        }
        for fi38 := 0; fi38 < 2; fi38++ {
            v1 = v1
            if v4 {
                if ((v28 + v28) > (v28 + "a")) {
                    v39 := GenS3{F0: 42, F1: -721}
                    v4 = (42 > 1)
                    v5 = append(v5, "hello")
                    _ = v39
                }
                if false {
                    if v4 {
                        v40 := GenS3{F0: 0, F1: 1000}
                        if len(v5) > 0 {
                            v5[0] = "\n\t"
                        }
                        _ = v40
                    } else {
                        switch v25 {
                        case "a":
                            var v41 uint16 = 1
                            _ = v41
                        case "\n\t":
                            v4 = (float32(float64(3.14159)) < (-2.25 / (((0.0 / (float32(v2) / ((-2.25 * -2.25) - (1e10 + 3.14159)))) * (0.0 - (1.5 / 1.5))) * (3.14159 / 3.14159))))
                        case "hello":
                            var v42 bool = (!(575 != 0))
                            _ = v42
                        default:
                            v43 := GenS4{F0: v1, F1: 42}
                            _ = v43
                        }
                        fmt.Println(v1, v28)
                    }
                    v26 += v2
                }
                var v44 float64 = 0.1
                _ = v44
            } else {
                if false {
                    v25 = ""
                    v45 := GenS3{F0: 1000, F1: 0}
                    v1 = (^(v1))
                    _ = v45
                }
                if v4 {
                    v2 = v2
                    var v46 uint8 = v1
                    fmt.Println(v29)
                    _ = v46
                }
            }
            _ = fi38
        }
        _ = fi32
        _ = v33
    }
    v2 = (v2 + v26)
    _ = len(v5)
    _ = len(v5)
    var v47 int = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v47
    _ = v5
}
