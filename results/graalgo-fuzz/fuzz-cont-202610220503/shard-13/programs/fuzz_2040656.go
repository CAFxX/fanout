package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool; F1 uint }
func genFn0(p0 int32) int16 {
    var v6 uint = 42
    fmt.Println(v6, p0)
    v7 := []int{-9223372036854775808, -483, 42, 4294967295}
    _ = p0
    _ = v6
    _ = v7
    return -1
}

func main() {
    var v1 bool = true
    var v2 int16 = -954
    var v3 int64 = -9223372036854775808
    v4 := make(map[string]bool)
    var v5 int8 = -1
    var v8 uint64 = 1
    switch -1 {
    case 129:
        if (false && true) {
            v5 = (v5 % (((v5 & 127)) | 1))
            for fi9 := 0; fi9 < 13; fi9++ {
                if (42 > 459) {
                    v10 := GenS1{F0: fi9}
                    v8 -= (v8 * 1)
                    for fi11 := 0; fi11 < 4; fi11++ {
                        switch "" {
                        case "\n\t":
                            var v12 int8 = -1
                            _ = v12
                        }
                        _ = fi11
                    }
                    _ = v10
                } else {
                    v1 = true
                    v13 := v4[""]
                    v14 := GenS1{F0: fi9}
                    _ = v13
                    _ = v14
                }
                _ = fi9
            }
            if (((!(42 != -855)) || (52 > 1)) || (0.0 <= float64((3.14159 - float32(1.5))))) {
                if v1 {
                    var v15 float32 = (0.1 + (0.0 + (0.0 / -2.25)))
                    fmt.Println(v5, v2)
                    var v16 float32 = v15
                    _ = v15
                    _ = v16
                } else {
                    var v17 uint8 = 128
                    _ = v17
                }
                var v18 uint64 = v8
                _ = v18
            } else {
                v19 := []int{-404, -489}
                v20 := GenS2{F0: 42, F1: (((("a" + ((("foo bar" + "\n\t") + ("" + "a")) + ((("\n\t" + "foo bar") + ("hello" + "")) + ("\n\t" + "\n\t")))) + (("hello" + "foo bar") + ("hello" + ("\n\t" + ("\n\t" + ""))))) + "a") + ((("foo bar" + ("" + "")) + "") + "foo bar"))}
                v2 -= 1000
                _ = v19
                _ = v20
            }
        } else {
            switch -724 {
            case 42:
                v21 := []int{416}
                _ = v21
            }
            v2 -= v2
            v22 := []int{-376}
            _ = v22
        }
    }
    fmt.Println(v2, v2)
    switch 991 {
    case -815:
        if ((-634 == (-(v2))) || (!((float32(1e10) - (-2.25 * float32(float64((0.1 + 0.0))))) >= -2.25))) {
            for fi23 := 0; fi23 < 20; fi23++ {
                v2 -= v2
                fmt.Println(v5, fi23)
                _ = fi23
            }
            v24 := []int{127}
            var v25 bool = (!true)
            _ = v24
            _ = v25
        }
    default:
        v26 := GenS2{F0: 2147483647, F1: "a"}
        _ = v26
    }
    _ = len(v4)
    for fi27 := 0; fi27 < 14; fi27++ {
        fmt.Println(fi27, v5)
        _ = fi27
    }
    fmt.Println(v2)
    for fi28 := 0; fi28 < 1; fi28++ {
        switch (^(fi28)) {
        case 361:
            if (2147483647 != 1) {
                v5 = 1
                v29 := GenS2{F0: (^(fi28)), F1: "a"}
                if true {
                    for fi30 := 0; fi30 < 3; fi30++ {
                        v31 := []int{fi28, (fi30 + (fi30 & -177))}
                        v32 := GenS3{F0: 2}
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                    if (("\n\t" != "\n\t") || (104 >= 224)) {
                        if v1 {
                            var v33 float32 = 1.5
                            v8 -= v8
                            switch "a" {
                            case "hello":
                                v5 = v5
                            case "":
                                v8 = v8
                            }
                            _ = v33
                        }
                        switch ("a" + "") {
                        case "\n\t":
                            v34 := GenS4{F0: v1, F1: 467}
                            _ = v34
                        case "foo bar":
                            if v1 {
                                v35 := GenS4{F0: v1, F1: 1}
                                _ = v35
                            }
                        }
                        switch ("" + ("" + "foo bar")) {
                        case "a":
                            var v36 int32 = 2147483647
                            _ = v36
                        default:
                            v5 = v5
                        }
                    }
                } else {
                    var v37 int8 = (v5 / ((v5) | 1))
                    _ = v37
                }
                _ = v29
            } else {
                switch fi28 {
                case 128:
                    if true {
                        if true {
                            v8 = 0
                            v5 += (v5 / ((-1) | 1))
                        } else {
                            v3 += (v3 ^ v3)
                        }
                    }
                }
            }
        case 319:
            for fi38 := 0; fi38 < 16; fi38++ {
                var v39 uint = 1
                _ = fi38
                _ = v39
            }
        case 127:
            v5 += (v5 * (v5 - int8(v8)))
        default:
            v1 = v1
        }
        v40 := GenS4{F0: (v1 && (!((((v5 != v5) && v1) || v1) && v1))), F1: 42}
        _ = fi28
        _ = v40
    }
    var v41 int64 = (-(v3))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v8
}
