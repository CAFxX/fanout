package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 bool }

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 65535
    v3 := make(map[string]int)
    v2 = (v2 | 0)
    v4 := GenS1{F0: "\n\t"}
    var v5 int8 = 1
    var v6 int = 32767
    v7 := []int{(-(v6))}
    if true {
        if (v6 > v6) {
            for fi8 := 0; fi8 < 11; fi8++ {
                v6 = 42
                if (true || (!false)) {
                    for fi9 := 0; fi9 < 9; fi9++ {
                        v2 += 705
                        _ = fi9
                    }
                    var v10 int16 = 0
                    _ = v10
                } else {
                    var v11 bool = (-2147483648 >= 0)
                    var v12 int = (v6 - (fi8 - fi8))
                    _ = v11
                    _ = v12
                }
                _ = fi8
            }
            if ((true || (0.1 > (1.5 + (1e10 - (0.0 / float32((0.1 + 3.14159))))))) && ((!false) && true)) {
                var v13 string = "foo bar"
                _ = v13
            }
            for fi14 := 0; fi14 < 12; fi14++ {
                v5 = v5
                var v15 int64 = -1
                _ = fi14
                _ = v15
            }
        } else {
            fmt.Println(v2)
            v2 = v2
        }
        for fi16 := 0; fi16 < 7; fi16++ {
            var v17 int32 = 42
            for fi18 := 0; fi18 < 16; fi18++ {
                if true {
                    if (((v5 | (^(v5))) != -81) || (1 == 0)) {
                        v2 += (v2 & (^(v2)))
                        var v19 int = (fi18 >> ((18446744073709551615) & 63))
                        _ = v19
                    }
                }
                if true {
                    var v20 float32 = (1.5 * 1e10)
                    if false {
                        var v21 int8 = -1
                        var v22 float32 = (1.5 - 0.0)
                        _ = v21
                        _ = v22
                    }
                    v17 += -1
                    v23 := GenS2{F0: -797, F1: false}
                    switch ("\n\t" + "\n\t") {
                    case "a":
                        var v24 int = (^(v6))
                        _ = v24
                    default:
                        if false {
                            v5 -= (v5 % ((v5) | 1))
                        } else {
                            var v25 uint64 = 515
                            v5 += v5
                            _ = v25
                        }
                    }
                    v3["hello"] = fi16
                    v2 -= (v2 / ((v2) | 1))
                    _ = v20
                    _ = v23
                } else {
                    var v26 string = ""
                    var v27 uint64 = 2
                    _ = v26
                    _ = v27
                }
                v28 := GenS1{F0: "\n\t"}
                var v29 string = (("\n\t" + ("foo bar" + "\n\t")) + (("hello" + ("\n\t" + "hello")) + (("a" + "") + (("" + "foo bar") + (((("\n\t" + "a") + "hello") + "\n\t") + (("" + "hello") + ("" + "\n\t")))))))
                _ = fi18
                _ = v28
                _ = v29
            }
            _ = fi16
            _ = v17
        }
        for fi30 := 0; fi30 < 16; fi30++ {
            var v31 uint16 = v2
            _ = fi30
            _ = v31
        }
    }
    v2 = (v2 >> ((18446744073709551615) & 63))
    v32 := GenS1{F0: "\n\t"}
    var v33 bool = (38 != (v5 / ((v5) | 1)))
    v34 := GenS1{F0: ""}
    switch "foo bar" {
    case "foo bar":
        var v35 uint32 = 42
        _ = v35
    case "\n\t":
        v33 = true
    case "hello":
        v33 = v33
    default:
        _ = len(v7)
    }
    var v36 float64 = 3.14159
    fmt.Println(v36, v5)
    if (true && v33) {
        v36 += (float64((((-2.25 - 1.5) - 1e10) / (1e10 + (1.5 / (1e10 + float32((((3.14159 * v36) - 3.14159) * 1e10))))))) * 0.0)
        if ("hello" <= "hello") {
            var v37 uint = 18446744073709551615
            _ = v37
        }
        var v38 uint = 2
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
