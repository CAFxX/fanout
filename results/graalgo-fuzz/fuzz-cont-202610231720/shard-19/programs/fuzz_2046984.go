package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 bool; F2 uint }

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = 1
    v3 := make(map[uint32]string)
    var v4 bool = ((-128 > 0) && false)
    var v5 float64 = (((1.5 * (1e10 + ((3.14159 + (((3.14159 + (1.5 * (0.0 + 3.14159))) - ((0.0 - 0.0) + 0.0)) * 1e10)) + (1e10 + 1e10)))) / 1.0) / 1e10)
    v4 = ("a" <= ((((("foo bar" + "hello") + "hello") + (((("hello" + "") + "") + "\n\t") + ("hello" + "a"))) + ("" + ("\n\t" + ("\n\t" + "")))) + ""))
    v6 := []int{-2147483648, 127}
    v7 := GenS1{F0: 193, F1: "foo bar", F2: (2147483647 >= 0)}
    var v8 float32 = 1.5
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := GenS1{F0: 77, F1: "foo bar", F2: ((^(fi9)) < 128)}
        for fi11 := 0; fi11 < 8; fi11++ {
            v5 = (((((v5 + 3.14159) - (v5 - 0.1)) / v5) * v5) + 0.1)
            _ = fi11
        }
        v8 = v8
        _ = fi9
        _ = v10
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        fmt.Println(v8)
        _ = fi12
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        v14 := GenS1{F0: 161, F1: "foo bar", F2: v4}
        _ = fi13
        _ = v14
    }
    if v4 {
        var v15 uint = 1
        v16 := GenS1{F0: 255, F1: ("hello" + ((("hello" + (((("\n\t" + "") + ("foo bar" + "\n\t")) + "\n\t") + "hello")) + "") + ("\n\t" + "\n\t"))), F2: v4}
        _ = v15
        _ = v16
    } else {
        v17 := v3[753]
        for fi18 := 0; fi18 < 10; fi18++ {
            var v19 uint16 = 32767
            if (314 != 1) {
                if v4 {
                    v8 += v8
                    switch (-(fi18)) {
                    case 951:
                        if (v19 >= v19) {
                            fmt.Println(v8, v4)
                        }
                    case 32767:
                        v20 := GenS2{F0: (!v4)}
                        _ = v20
                    case 497:
                        v5 = v5
                    }
                } else {
                    v5 = v5
                    switch (v17 + v17) {
                    case "\n\t":
                        v3[0] = v17
                    case "a":
                        if true {
                            if v4 {
                                v8 -= 0.0
                                v8 -= (v8 / 0.1)
                            } else {
                                for fi21 := 0; fi21 < 13; fi21++ {
                                    switch (fi18 << ((65535) & 63)) {
                                    case 797:
                                        if v4 {
                                            if v4 {
                                                _ = len(v1)
                                            }
                                            if v4 {
                                                var v22 int32 = -964
                                                fmt.Println(v2)
                                                v23 := GenS3{F0: 2, F1: 2}
                                                _ = v22
                                                _ = v23
                                            }
                                            v2 = v2
                                            if (((!(!v4)) || v4) && (v2 == (v2 * (v2 * 42)))) {
                                                if true {
                                                    v2 -= 1000
                                                    switch (v17 + (v17 + "a")) {
                                                    case "hello":
                                                        v24 := GenS2{F0: (46 != 100)}
                                                        _ = v24
                                                    case "":
                                                        _ = len(v6)
                                                    default:
                                                        v25 := v1[42]
                                                        _ = v25
                                                    }
                                                    var v26 uint64 = 42
                                                    _ = v26
                                                } else {
                                                    v8 -= ((float32((((v5 - v5) * v5) / (v5 + float64(v8)))) / v8) * 0.1)
                                                    fmt.Println(v2, v2)
                                                    var v27 string = "foo bar"
                                                    if false {
                                                        v8 = v8
                                                    }
                                                    if v4 {
                                                        _ = len(v3)
                                                        v4 = v4
                                                    } else {
                                                        v8 -= v8
                                                    }
                                                    _ = v27
                                                }
                                            } else {
                                                var v28 bool = v4
                                                fmt.Println(v28)
                                                _ = v28
                                            }
                                            if false {
                                                v29 := GenS4{F0: 1}
                                                switch 1 {
                                                case -846:
                                                    if true {
                                                        var v30 int64 = 104
                                                        v31 := v1[4294967295]
                                                        v19 -= 1
                                                        _ = v30
                                                        _ = v31
                                                    }
                                                }
                                                if v4 {
                                                    v32 := GenS3{F0: 127, F1: 17}
                                                    var v33 string = v17
                                                    v4 = (!false)
                                                    _ = v32
                                                    _ = v33
                                                } else {
                                                    var v34 bool = v4
                                                    _ = v34
                                                }
                                                _ = v29
                                            }
                                            v17 = v17
                                        }
                                    default:
                                        var v35 int = (fi21 + 32767)
                                        _ = v35
                                    }
                                    v36 := GenS5{F0: 1, F1: (true && v4), F2: 255}
                                    switch fi18 {
                                    case 127:
                                        switch 128 {
                                        case -9223372036854775808:
                                            v2 -= v2
                                        }
                                    case 4294967295:
                                        var v37 int64 = 1
                                        _ = v37
                                    }
                                    _ = fi21
                                    _ = v36
                                }
                                for fi38 := 0; fi38 < 9; fi38++ {
                                    v17 = v17
                                    _ = len(v6)
                                    var v39 int32 = 42
                                    _ = fi38
                                    _ = v39
                                }
                            }
                            v19 -= 1
                            v40 := GenS5{F0: 127, F1: v4, F2: 1}
                            _ = v40
                        } else {
                            v41 := GenS1{F0: 125, F1: "\n\t", F2: false}
                            _ = v41
                        }
                    }
                    v8 = (1e10 * 1.5)
                }
                v17 = ("hello" + v17)
            }
            _ = fi18
            _ = v19
        }
        _ = v17
    }
    var v42 int32 = 1
    v8 += v8
    if false {
        v5 -= (1e10 - v5)
    }
    v43 := v3[0]
    v44 := GenS1{F0: 93, F1: (v43 + v43), F2: (!(127 >= 42))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
