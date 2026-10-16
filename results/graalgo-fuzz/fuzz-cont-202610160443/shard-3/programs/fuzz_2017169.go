package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
func genFn0() uint8 {
    if ("a" < "hello") {
        for fi5 := 0; fi5 < 10; fi5++ {
            v6 := GenS1{F0: "", F1: 4294967295}
            _ = fi5
            _ = v6
        }
    } else {
        if true {
            var v7 string = ("a" + "")
            v8 := []int{-317, -1, 42}
            _ = v7
            _ = v8
        } else {
            var v9 int64 = -132
            switch "hello" {
            case "hello":
                v10 := []int{-565, -2147483648, 1}
                _ = v10
            }
            _ = v9
        }
        switch "hello" {
        case "\n\t":
            for fi11 := 0; fi11 < 1; fi11++ {
                var v12 uint16 = 2
                _ = fi11
                _ = v12
            }
        }
        var v13 int8 = 108
        _ = v13
    }
    if false {
        var v14 float64 = (((float64((1e10 * float32(-2.25))) + -2.25) / 1.5) / float64(((1.5 / 3.14159) - float32(float64(1e10)))))
        _ = v14
    }
    var v15 float64 = (float64((((3.14159 * float32(1.5)) / 0.1) - (0.0 / -2.25))) + (0.1 + 3.14159))
    if (v15 < v15) {
        if (32767 != -1) {
            for fi16 := 0; fi16 < 13; fi16++ {
                v15 = (float64(0.0) + (float64(0.1) * -2.25))
                fmt.Println(v15, fi16)
                if false {
                    for fi17 := 0; fi17 < 4; fi17++ {
                        if ("foo bar" >= "") {
                            v15 += (1e10 + (1e10 * float64(1e10)))
                            fmt.Println(v15, v15, fi16)
                        }
                        fmt.Println(fi17, fi16)
                        _ = fi17
                    }
                } else {
                    v15 = v15
                    for fi18 := 0; fi18 < 16; fi18++ {
                        if (false && (2147483647 > 1)) {
                            var v19 int8 = 100
                            _ = v19
                        } else {
                            if true {
                                var v20 float64 = (v15 * v15)
                                _ = v20
                            }
                        }
                        v21 := []int{fi16}
                        v15 -= float64(3.14159)
                        _ = fi18
                        _ = v21
                    }
                    v15 += v15
                }
                _ = fi16
            }
            v15 = (v15 * float64((0.0 / (3.14159 * (1.5 / (-2.25 / 0.1))))))
            v15 = v15
        }
        for fi22 := 0; fi22 < 12; fi22++ {
            v15 += v15
            v23 := []int{691, fi22, 0}
            _ = fi22
            _ = v23
        }
        switch 2147483647 {
        case -142:
            v15 = v15
        }
    } else {
        for fi24 := 0; fi24 < 15; fi24++ {
            v15 -= float64(1.5)
            _ = fi24
        }
    }
    if false {
        var v25 bool = (true || false)
        _ = v25
    }
    _ = v15
    return 0
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 641
    v3 := make(map[string]string)
    var v4 int64 = -9223372036854775808
    var v26 uint32 = 0
    for fi27 := 0; fi27 < 7; fi27++ {
        var v28 uint = (v2 & v2)
        switch "" {
        case "\n\t":
            v29 := []int{}
            _ = v29
        case "foo bar":
            v28 += 1
        default:
            v30 := GenS1{F0: "foo bar", F1: v26}
            _ = v30
        }
        fmt.Println(v2, fi27)
        _ = fi27
        _ = v28
    }
    var v31 bool = (1e10 >= 0.1)
    v32 := []int{32767, 817, 42}
    v33 := GenS1{F0: "foo bar", F1: v26}
    v26 = 1
    for fi34 := 0; fi34 < 5; fi34++ {
        v35 := GenS1{F0: ("hello" + "a"), F1: v26}
        _ = fi34
        _ = v35
    }
    v36 := GenS1{F0: "", F1: v26}
    v37 := GenS1{F0: "\n\t", F1: v26}
    for fi38 := 0; fi38 < 12; fi38++ {
        fmt.Println(v26, v26)
        _ = fi38
    }
    if false {
        var v39 int = 0
        if false {
            fmt.Println(v4, v39, v31)
            v4 += v4
            for fi40 := 0; fi40 < 3; fi40++ {
                fmt.Println(v39)
                _ = fi40
            }
        }
        _ = v39
    }
    v4 += -9223372036854775808
    v2 = (^(v2))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v36
    _ = v37
    _ = v4
}
