package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 uint8; F1 bool }
func genFn0() uint32 {
    if ("" == "foo bar") {
        for fi5 := 0; fi5 < 17; fi5++ {
            switch (fi5 ^ fi5) {
            case 365:
                switch fi5 {
                case 1:
                    switch ("\n\t" + "\n\t") {
                    case "foo bar":
                        switch ("a" + "\n\t") {
                        case "foo bar":
                            var v6 uint8 = 42
                            _ = v6
                        }
                    case "hello":
                        for fi7 := 0; fi7 < 7; fi7++ {
                            var v8 uint32 = 390
                            _ = fi7
                            _ = v8
                        }
                    case "a":
                        var v9 uint16 = 2
                        _ = v9
                    default:
                        var v10 int8 = 127
                        _ = v10
                    }
                case 128:
                    if false {
                        if true {
                            v11 := GenS1{F0: 172, F1: "a"}
                            var v12 float64 = (-2.25 + 1e10)
                            _ = v11
                            _ = v12
                        } else {
                            fmt.Println(fi5)
                            var v13 int8 = -1
                            _ = v13
                        }
                        v14 := []int{(fi5 * (^(fi5))), (fi5 & fi5), (fi5 & 112), (fi5 | (^(fi5)))}
                        _ = v14
                    }
                }
            case 0:
                v15 := GenS1{F0: 65535, F1: "\n\t"}
                _ = v15
            }
            v16 := GenS2{F0: (("a" + ("\n\t" + ("" + "a"))) + "foo bar")}
            _ = fi5
            _ = v16
        }
        var v17 string = ""
        _ = v17
    }
    v18 := GenS3{F0: 42, F1: 0}
    v19 := GenS3{F0: 4294967295, F1: 0}
    _ = v18
    _ = v19
    return 547
}
func genFn1(p0 uint8, p1 int64) bool {
    v20 := []int{4294967295, -639, -2147483648}
    var v21 uint64 = 257
    if true {
        v22 := GenS2{F0: "foo bar"}
        v23 := GenS4{F0: (p0 | p0), F1: ((p1 ^ p1) == p1)}
        _ = v22
        _ = v23
    }
    _ = p0
    _ = p1
    _ = v20
    _ = v21
    return (!(!false))
}

func main() {
    var v1 int = -807
    var v2 float64 = (0.0 * 1.5)
    var v3 int = 0
    var v4 uint64 = 0
    v24 := []int{(v1 / ((v1) | 1))}
    for fi25 := 0; fi25 < 1; fi25++ {
        if true {
            v26 := GenS1{F0: 225, F1: "hello"}
            _ = v26
        } else {
            for fi27 := 0; fi27 < 13; fi27++ {
                switch ((("\n\t" + ("hello" + "foo bar")) + "") + "\n\t") {
                case "\n\t":
                    v3 += 442
                case "":
                    v28 := GenS1{F0: 65535, F1: ""}
                    _ = v28
                default:
                    switch fi25 {
                    case 0:
                        if len(v24) > 0 {
                            v24[0] = 865
                        }
                    case -2147483648:
                        var v29 uint32 = 816
                        _ = v29
                    case -472:
                        fmt.Println(v2)
                    }
                }
                if (v4 == v4) {
                    var v30 uint8 = 23
                    if len(v24) > 0 {
                        v24[0] = 32767
                    }
                    var v31 float64 = v2
                    _ = v30
                    _ = v31
                }
                _ = fi27
            }
        }
        _ = fi25
    }
    fmt.Println(v4, v3)
    for fi32 := 0; fi32 < 12; fi32++ {
        var v33 int32 = 750
        _ = fi32
        _ = v33
    }
    _ = len(v24)
    v34 := GenS3{F0: v4, F1: 0}
    var v35 int8 = -128
    fmt.Println(v35, v4, v1)
    var v36 string = "foo bar"
    if (false && (false && (563 < 1))) {
        if len(v24) > 0 {
            v24[0] = v1
        }
    }
    var v37 uint8 = 133
    v38 := GenS3{F0: v4, F1: -1}
    fmt.Println(v37)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
