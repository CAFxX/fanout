package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint; F1 int16; F2 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int; F1 int16 }
type GenS7 struct { F0 int8 }
func genFn0(p0 int8) string {
    var v7 int8 = (p0 ^ 45)
    var v8 uint8 = 255
    var v9 float32 = 1e10
    v10 := GenS1{F0: (("hello" + "") + (("hello" + "foo bar") + "\n\t")), F1: ("hello" + "a")}
    if (v9 == v9) {
        if ((v9 == v9) || false) {
            fmt.Println(v9, p0, v8)
        } else {
            var v11 int16 = -339
            _ = v11
        }
    }
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return ("" + ("" + ""))
}
func genFn1() int {
    for fi12 := 0; fi12 < 7; fi12++ {
        v13 := []int{-693, 9223372036854775807, (fi12 | (fi12 * fi12))}
        var v14 int8 = 42
        _ = fi12
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: "a", F1: (("foo bar" + ("" + "")) + "hello")}
    v16 := GenS1{F0: "a", F1: "foo bar"}
    var v17 int = 9223372036854775807
    _ = v15
    _ = v16
    _ = v17
    return (v17 | (v17 | (v17 % ((v17) | 1))))
}

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 110
    var v3 uint64 = 793
    var v4 float32 = 0.1
    v5 := make(map[string]bool)
    var v6 uint8 = (v2 >> ((65535) & 63))
    var v18 uint16 = 2
    var v19 float64 = 3.14159
    v20 := []int{-986, 4294967295, -771, genFn1()}
    var v21 uint16 = (v18 << ((272) & 63))
    if true {
        var v22 int64 = 1
        v23 := GenS1{F0: "hello", F1: "\n\t"}
        _ = len(v5)
        _ = v22
        _ = v23
    } else {
        for fi24 := 0; fi24 < 9; fi24++ {
            v3 = v3
            _ = fi24
        }
        if (true || ((!true) || (true && true))) {
            v25 := GenS2{F0: 1, F1: -709, F2: 2147483647}
            for fi26 := 0; fi26 < 3; fi26++ {
                var v27 string = ""
                fmt.Println(v18)
                _ = fi26
                _ = v27
            }
            _ = v25
        }
        v19 -= (float64(v4) - -2.25)
    }
    v28 := GenS1{F0: ("hello" + "a"), F1: "hello"}
    var v29 string = "foo bar"
    v30 := GenS3{F0: 0}
    switch v29 {
    case "":
        switch (v29 + (v29 + "\n\t")) {
        case "":
            switch v29 {
            case "":
                v31 := GenS3{F0: 0}
                _ = v31
            case "hello":
                v32 := GenS1{F0: "a", F1: "foo bar"}
                _ = v32
            case "a":
                if (-1 > 127) {
                    var v33 uint16 = (v18 - (v21 % (((^(v18))) | 1)))
                    var v34 float32 = v4
                    _ = v33
                    _ = v34
                }
            }
        case "a":
            var v35 int = -74
            _ = v35
        default:
            var v36 uint = 65535
            _ = v36
        }
    case "foo bar":
        v37 := GenS4{F0: 1, F1: v6}
        _ = v37
    default:
        for fi38 := 0; fi38 < 11; fi38++ {
            switch v29 {
            case "":
                for fi39 := 0; fi39 < 7; fi39++ {
                    v40 := GenS2{F0: 0, F1: 32767, F2: 32767}
                    switch genFn1() {
                    case 4294967295:
                        if true {
                            fmt.Println(v6)
                            var v41 uint64 = 343
                            _ = v41
                        } else {
                            switch fi39 {
                            case -274:
                                var v42 bool = true
                                _ = v42
                            case -802:
                                if (-128 == -128) {
                                    v43 := v1[fi39]
                                    _ = v43
                                } else {
                                    v44 := GenS2{F0: 255, F1: 1000, F2: 308}
                                    var v45 int16 = 32767
                                    _ = v44
                                    _ = v45
                                }
                            default:
                                var v46 float32 = v4
                                _ = v46
                            }
                            var v47 string = ((("" + "\n\t") + "a") + (v29 + "\n\t"))
                            _ = v47
                        }
                    default:
                        v48 := GenS5{F0: -1}
                        _ = v48
                    }
                    _ = fi39
                    _ = v40
                }
            }
            _ = fi38
        }
    }
    v49 := GenS6{F0: 563, F1: 32767}
    for fi50 := 0; fi50 < 9; fi50++ {
        v51 := GenS3{F0: -1}
        fmt.Println(v3, v4, v29)
        if (false || (true || false)) {
            v52 := GenS7{F0: -105}
            delete(v1, -170)
            _ = v52
        }
        _ = fi50
        _ = v51
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v49
    _ = v5
    _ = v6
}
