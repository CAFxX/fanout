package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 string }
func genFn0() uint {
    switch 1 {
    case 9223372036854775807:
        switch (("\n\t" + "a") + "hello") {
        case "hello":
            var v4 int8 = 42
            _ = v4
        case "":
            if true {
                v5 := GenS1{F0: 100}
                var v6 int8 = -1
                if true {
                    if (false && true) {
                        var v7 float64 = ((float64((0.0 + 0.0)) * -2.25) + 1.5)
                        _ = v7
                    } else {
                        switch 273 {
                        case 4294967295:
                            switch (("\n\t" + "") + ("hello" + (("hello" + ("" + ("\n\t" + "\n\t"))) + ("\n\t" + "a")))) {
                            case "":
                                fmt.Println(v6)
                            }
                        default:
                            v6 = v6
                        }
                    }
                    for fi8 := 0; fi8 < 6; fi8++ {
                        v9 := GenS1{F0: v6}
                        v10 := GenS1{F0: v6}
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                }
                _ = v5
                _ = v6
            } else {
                var v11 int8 = 100
                switch 309 {
                case 128:
                    for fi12 := 0; fi12 < 8; fi12++ {
                        var v13 int32 = 42
                        var v14 uint8 = 42
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                case 9223372036854775807:
                    switch 289 {
                    case 712:
                        var v15 uint = 255
                        _ = v15
                    case -24:
                        for fi16 := 0; fi16 < 3; fi16++ {
                            v17 := GenS2{F0: v11, F1: "\n\t"}
                            _ = fi16
                            _ = v17
                        }
                    }
                }
                _ = v11
            }
        case "\n\t":
            var v18 uint64 = 42
            _ = v18
        }
    default:
        v19 := GenS1{F0: -128}
        _ = v19
    }
    var v20 uint64 = 2
    for fi21 := 0; fi21 < 4; fi21++ {
        v20 -= v20
        var v22 int64 = 653
        _ = fi21
        _ = v22
    }
    v23 := []int{}
    _ = v20
    _ = v23
    return 18446744073709551615
}

func main() {
    v1 := []string{(("\n\t" + "a") + "a")}
    v2 := []string{(("a" + ("\n\t" + "a")) + "\n\t")}
    var v3 float32 = ((1.5 * 0.1) * -2.25)
    v3 += v3
    for fi24 := 0; fi24 < 19; fi24++ {
        v25 := GenS2{F0: 127, F1: ""}
        _ = fi24
        _ = v25
    }
    var v26 float64 = 0.1
    switch "hello" {
    case "":
        v1 = append(v1, ("" + "a"))
    case "foo bar":
        if true {
            v26 += (v26 + 0.0)
            for fi27 := 0; fi27 < 2; fi27++ {
                fmt.Println(v26)
                v3 = v3
                _ = fi27
            }
            fmt.Println(v26, v3)
        } else {
            var v28 bool = false
            for fi29 := 0; fi29 < 19; fi29++ {
                if v28 {
                    if ((float64((v3 * (v3 + (((v3 - v3) / 3.14159) * (v3 / v3))))) / v26) >= 0.1) {
                        var v30 bool = (1e10 < (3.14159 + ((((v26 + (v26 / float64((float32(0.1) + v3)))) / v26) * 1e10) - (v26 * v26))))
                        v28 = (!v30)
                        v28 = true
                        _ = v30
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        switch (("hello" + "\n\t") + (("" + "hello") + ((("a" + ("\n\t" + "a")) + "foo bar") + "a"))) {
                        case "hello":
                            fmt.Println(v26, v26, fi29)
                        case "a":
                            v26 = (0.0 / ((float64(((float32((1e10 - (v26 - v26))) + (0.0 * v3)) * (v3 / 0.1))) - v26) - (1e10 + float64(((v3 * v3) - v3)))))
                        }
                    }
                }
                var v31 string = ("\n\t" + "hello")
                if true {
                    if len(v2) > 0 {
                        v2[0] = (v31 + ("" + (v31 + "hello")))
                    }
                    var v32 uint16 = 1
                    _ = v32
                } else {
                    fmt.Println(v3, v3)
                    v3 = (((v3 * float32((1e10 + v26))) * (v3 - (0.0 / v3))) + (1e10 * v3))
                    v31 = v31
                }
                _ = fi29
                _ = v31
            }
            v33 := GenS2{F0: 1, F1: "hello"}
            _ = v28
            _ = v33
        }
    }
    v1 = append(v1, ("a" + "a"))
    for fi34 := 0; fi34 < 13; fi34++ {
        v26 -= (3.14159 * float64(v3))
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
        v3 = -2.25
        _ = fi34
    }
    var v35 uint = 255
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v35
}
