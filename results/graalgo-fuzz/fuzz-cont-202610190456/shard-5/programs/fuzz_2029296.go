package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 uint; F2 string }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint64; F1 bool }
type GenS6 struct { F0 int64 }
func genFn0(p0 bool, p1 bool) int64 {
    v5 := GenS1{F0: 1}
    v6 := GenS1{F0: 32767}
    var v7 uint8 = 127
    var v8 int64 = 0
    fmt.Println(v8, p0)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return v8
}

func main() {
    var v1 int64 = 196
    var v2 uint8 = 56
    v3 := make(map[int]bool)
    v4 := make(map[string]int)
    v2 = v2
    if true {
        v9 := GenS2{F0: 1000, F1: 255, F2: ("foo bar" + "hello")}
        if (v2 >= v2) {
            v10 := []int{}
            v2 = v2
            if (false || (42 == 127)) {
                v4["foo bar"] = 2147483647
            }
            _ = v10
        }
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 int32 = -1
            _ = fi11
            _ = v12
        }
        _ = v9
    }
    if false {
        v13 := GenS2{F0: 1, F1: 65535, F2: "hello"}
        _ = v13
    }
    var v14 int64 = v1
    fmt.Println(v14, v1)
    v15 := GenS3{F0: 255, F1: (("hello" + "hello") + "a")}
    v16 := GenS4{F0: 0}
    for fi17 := 0; fi17 < 2; fi17++ {
        if ((!false) && false) {
            for fi18 := 0; fi18 < 3; fi18++ {
                v19 := GenS5{F0: 18446744073709551615, F1: (!true)}
                if false {
                    var v20 int64 = -687
                    switch "a" {
                    case "\n\t":
                        v14 = v14
                    case "hello":
                        v21 := v3[127]
                        var v22 uint8 = v2
                        switch "foo bar" {
                        case "foo bar":
                            if v21 {
                                var v23 uint32 = 42
                                v24 := GenS5{F0: 0, F1: (0 < 127)}
                                v25 := []int{}
                                _ = v23
                                _ = v24
                                _ = v25
                            }
                        case "hello":
                            v2 -= v2
                        case "":
                            fmt.Println(v14)
                            v1 = (v20 / (((-(v1))) | 1))
                        }
                        _ = v21
                        _ = v22
                    }
                    var v26 uint8 = (v2 | v2)
                    _ = v20
                    _ = v26
                } else {
                    if true {
                        v27 := v3[fi17]
                        v28 := GenS2{F0: 32767, F1: 2, F2: (("foo bar" + (("\n\t" + "foo bar") + "")) + "a")}
                        fmt.Println(fi18)
                        _ = v27
                        _ = v28
                    } else {
                        v29 := GenS3{F0: 65535, F1: (("a" + ("foo bar" + (("\n\t" + "a") + "\n\t"))) + "\n\t")}
                        switch ("a" + ("" + "a")) {
                        case "hello":
                            v30 := []int{}
                            fmt.Println(v14)
                            _ = v30
                        default:
                            switch ("\n\t" + "foo bar") {
                            case "":
                                var v31 float64 = 1.5
                                _ = v31
                            }
                        }
                        var v32 uint32 = 4294967295
                        _ = v29
                        _ = v32
                    }
                    fmt.Println(v1)
                }
                fmt.Println(v1)
                v2 -= v2
                var v33 uint16 = 941
                _ = fi18
                _ = v19
                _ = v33
            }
            v1 += (v14 & (v14 << ((272) & 63)))
            if (42 >= 856) {
                v34 := GenS2{F0: -481, F1: 42, F2: ""}
                v35 := v4[""]
                _ = v34
                _ = v35
            } else {
                var v36 uint = 44
                v37 := v4["hello"]
                _ = v36
                _ = v37
            }
        }
        var v38 float64 = (3.14159 / 1.0)
        var v39 float64 = v38
        _ = fi17
        _ = v38
        _ = v39
    }
    for fi40 := 0; fi40 < 2; fi40++ {
        v41 := GenS6{F0: v1}
        for fi42 := 0; fi42 < 11; fi42++ {
            v43 := GenS4{F0: v2}
            v14 = v14
            _ = fi42
            _ = v43
        }
        _ = fi40
        _ = v41
    }
    v14 = v14
    var v44 bool = true
    fmt.Println(v2, v1)
    v14 = (v14 ^ -550)
    for fi45 := 0; fi45 < 16; fi45++ {
        var v46 int = fi45
        if (!(v44 || (!v44))) {
            var v47 float32 = (3.14159 - -2.25)
            v14 = 1
            _ = v47
        } else {
            fmt.Println(v14)
        }
        v2 -= 1
        _ = fi45
        _ = v46
    }
    var v48 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v48
}
