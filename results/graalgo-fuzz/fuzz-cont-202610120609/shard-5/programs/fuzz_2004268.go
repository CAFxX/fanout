package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 int; F2 int64 }
type GenS4 struct { F0 int16; F1 int; F2 uint32 }
type GenS5 struct { F0 bool; F1 int8 }
type GenS6 struct { F0 int; F1 int16 }
func genFn0(p0 int32, p1 string) uint16 {
    for fi4 := 0; fi4 < 16; fi4++ {
        p0 += p0
        v5 := GenS1{F0: 42}
        _ = fi4
        _ = v5
    }
    if (1 >= 2) {
        if (false || (181 != 2)) {
            var v6 int64 = 42
            p0 += p0
            v7 := GenS1{F0: 656}
            _ = v6
            _ = v7
        }
        var v8 int8 = 70
        p0 = p0
        _ = v8
    } else {
        v9 := GenS2{F0: 0}
        _ = v9
    }
    _ = p0
    _ = p1
    return 0
}
func genFn1() uint64 {
    if (!(2 > 1)) {
        if (164 >= 0) {
            v10 := []int{0, -1, -1, 42}
            for fi11 := 0; fi11 < 19; fi11++ {
                fmt.Println(fi11)
                switch fi11 {
                case 32768:
                    fmt.Println(fi11)
                case 128:
                    switch ("\n\t" + ("foo bar" + "hello")) {
                    case "a":
                        if false {
                            fmt.Println(fi11)
                            v12 := GenS2{F0: 0}
                            if (1.5 >= -2.25) {
                                var v13 uint8 = 42
                                _ = v13
                            }
                            _ = v12
                        } else {
                            v14 := GenS3{F0: 401, F1: fi11, F2: 9223372036854775807}
                            _ = v14
                        }
                    case "":
                        for fi15 := 0; fi15 < 18; fi15++ {
                            v16 := GenS4{F0: 1000, F1: fi11, F2: 4294967295}
                            var v17 uint16 = genFn0(1, "hello")
                            if len(v10) > 0 {
                                v10[0] = (fi11 * (-(fi11)))
                            }
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                    case "foo bar":
                        switch ("\n\t" + ("\n\t" + "foo bar")) {
                        case "\n\t":
                            var v18 int16 = 497
                            _ = v18
                        case "":
                            fmt.Println(fi11)
                        }
                    }
                case 32767:
                    fmt.Println(fi11)
                }
                for fi19 := 0; fi19 < 12; fi19++ {
                    switch fi11 {
                    case -606:
                        var v20 int16 = 1
                        _ = v20
                    case -359:
                        var v21 uint64 = 169
                        _ = v21
                    case -939:
                        _ = len(v10)
                    }
                    var v22 uint32 = 0
                    if ((false && (1e10 != (float64((0.1 + 0.1)) + 3.14159))) || false) {
                        var v23 uint = 65535
                        _ = v23
                    } else {
                        v22 = v22
                    }
                    _ = fi19
                    _ = v22
                }
                _ = fi11
            }
            _ = v10
        } else {
            switch -451 {
            case 924:
                v24 := []int{}
                _ = v24
            default:
                var v25 int16 = -1
                _ = v25
            }
            v26 := []int{-2147483648}
            _ = v26
        }
        var v27 int64 = 0
        _ = v27
    } else {
        var v28 string = ("" + (("" + "a") + ("foo bar" + "")))
        _ = v28
    }
    v29 := GenS5{F0: true, F1: 42}
    var v30 string = ("" + "foo bar")
    _ = v29
    _ = v30
    return 2
}

func main() {
    var v1 int16 = -418
    var v2 int32 = 1
    var v3 int32 = v2
    switch "foo bar" {
    case "a":
        v3 -= v3
    case "\n\t":
        if (v2 != (-(v2))) {
            var v31 int32 = v3
            _ = v31
        } else {
            fmt.Println(v1)
            var v32 int32 = (v2 | v3)
            switch (("a" + (("hello" + "foo bar") + "\n\t")) + "hello") {
            case "hello":
                fmt.Println(v1, v32)
            }
            _ = v32
        }
    }
    var v33 int32 = (v2 - v3)
    v3 = -2147483648
    var v34 int = 4294967295
    for fi35 := 0; fi35 < 18; fi35++ {
        v36 := GenS6{F0: (v34 | (^(fi35))), F1: v1}
        _ = fi35
        _ = v36
    }
    v2 -= v33
    var v37 string = "a"
    v33 = (v33 % ((v33) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v37
}
