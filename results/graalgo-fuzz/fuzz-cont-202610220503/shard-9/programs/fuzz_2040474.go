package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 string }

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{-352, 415}
    var v3 int = -627
    v4 := make(map[string]bool)
    if (90 >= 2) {
        v3 -= (v3 / ((v3) | 1))
    }
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 bool = false
        fmt.Println(v3, fi5)
        var v7 int64 = -368
        _ = fi5
        _ = v6
        _ = v7
    }
    v3 = 42
    v8 := GenS1{F0: true, F1: 631, F2: "a"}
    v9 := GenS1{F0: (true || ((42 < 42) && (("a" == (("\n\t" + "hello") + "a")) && false))), F1: 1, F2: "hello"}
    var v10 int16 = 0
    v2 = append(v2, (v3 ^ 0))
    fmt.Println(v10, v3)
    for fi11 := 0; fi11 < 9; fi11++ {
        var v12 uint16 = 1
        _ = fi11
        _ = v12
    }
    if true {
        var v13 int64 = 1
        for fi14 := 0; fi14 < 11; fi14++ {
            v15 := GenS1{F0: true, F1: 65535, F2: "hello"}
            v3 = v3
            if len(v2) > 0 {
                v2[0] = -2147483648
            }
            _ = fi14
            _ = v15
        }
        _ = v13
    } else {
        for fi16 := 0; fi16 < 9; fi16++ {
            var v17 uint8 = 113
            if true {
                fmt.Println(v3, v3)
                var v18 string = ("hello" + "foo bar")
                _ = v18
            } else {
                switch v3 {
                case -327:
                    for fi19 := 0; fi19 < 2; fi19++ {
                        v10 = v10
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = len(v4)
                        _ = fi19
                    }
                default:
                    if true {
                        for fi20 := 0; fi20 < 11; fi20++ {
                            v21 := GenS1{F0: false, F1: 2, F2: ""}
                            _ = fi20
                            _ = v21
                        }
                        v4[("foo bar" + ("a" + ("hello" + "hello")))] = ((("foo bar" + ("\n\t" + ("" + "foo bar"))) + "") >= "")
                    } else {
                        switch 9223372036854775807 {
                        case 128:
                            switch "a" {
                            case "hello":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        case 42:
                            v3 = fi16
                        default:
                            v3 = v3
                        }
                    }
                }
                var v22 uint32 = 2
                _ = len(v1)
                _ = v22
            }
            v10 -= 32767
            _ = fi16
            _ = v17
        }
        var v23 uint8 = 181
        for fi24 := 0; fi24 < 9; fi24++ {
            for fi25 := 0; fi25 < 20; fi25++ {
                var v26 float64 = 1.5
                _ = fi25
                _ = v26
            }
            _ = fi24
        }
        _ = v23
    }
    var v27 uint = 255
    v10 = (v10 % (((v10 ^ (v10 / ((v10) | 1)))) | 1))
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
