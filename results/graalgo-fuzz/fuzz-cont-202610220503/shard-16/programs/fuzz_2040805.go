package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint16; F1 uint16 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]string)
    var v3 float64 = 1.5
    var v4 float64 = 0.0
    for fi5 := 0; fi5 < 2; fi5++ {
        var v6 uint = 4294967295
        v6 += v6
        v3 = v3
        _ = fi5
        _ = v6
    }
    fmt.Println(v4, v3)
    if (false && true) {
        v4 -= v3
        fmt.Println(v3)
        v7 := v2[("foo bar" + ("hello" + ""))]
        _ = v7
    }
    v3 += (1e10 + v4)
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 float32 = (0.0 * 0.1)
        v10 := GenS1{F0: 2147483648, F1: 1000}
        v3 -= v3
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS2{F0: 65535, F1: 32767}
    var v12 uint16 = 42
    if false {
        v4 -= v4
        var v13 int32 = -1
        var v14 string = "foo bar"
        _ = v13
        _ = v14
    } else {
        switch "" {
        case "":
            v15 := v1[(("foo bar" + ("a" + ("a" + ""))) + ("foo bar" + ""))]
            _ = v15
        case "\n\t":
            fmt.Println(v4)
        }
        for fi16 := 0; fi16 < 8; fi16++ {
            v17 := []int{(fi16 ^ (fi16 << ((761) & 63)))}
            v4 -= v4
            switch ("\n\t" + (("\n\t" + "") + "foo bar")) {
            case "\n\t":
                var v18 int16 = 944
                _ = v18
            case "hello":
                if (false || (false && (!(!true)))) {
                    _ = len(v1)
                    v12 -= (v12 / (((v12 - (v12 * (v12 % ((v12) | 1))))) | 1))
                    v19 := GenS2{F0: v12, F1: 620}
                    _ = v19
                } else {
                    for fi20 := 0; fi20 < 1; fi20++ {
                        if len(v17) >= 2 {
                            v17 = v17[1:len(v17)]
                        }
                        _ = fi20
                    }
                    var v21 int8 = 102
                    _ = v21
                }
            }
            _ = fi16
            _ = v17
        }
        if true {
            var v22 uint64 = 4294967295
            v3 = v3
            _ = v22
        } else {
            var v23 uint64 = 2
            for fi24 := 0; fi24 < 12; fi24++ {
                v23 = 4294967295
                _ = fi24
            }
            _ = v23
        }
    }
    for fi25 := 0; fi25 < 6; fi25++ {
        var v26 int16 = 42
        var v27 float32 = (1e10 / 3.14159)
        _ = fi25
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 17; fi28++ {
        var v29 uint8 = 2
        _ = fi28
        _ = v29
    }
    if true {
        var v30 string = "foo bar"
        _ = v30
    } else {
        for fi31 := 0; fi31 < 20; fi31++ {
            for fi32 := 0; fi32 < 3; fi32++ {
                if (!("hello" != ("a" + (("\n\t" + "a") + ("a" + "hello"))))) {
                    var v33 string = "foo bar"
                    v34 := GenS2{F0: v12, F1: v12}
                    _ = v33
                    _ = v34
                }
                switch fi31 {
                case -1:
                    if false {
                        if (("" + "") != "foo bar") {
                            v35 := GenS1{F0: 2, F1: 0}
                            _ = v35
                        }
                        v12 += v12
                        v3 += float64(1e10)
                    }
                }
                _ = len(v1)
                _ = fi32
            }
            _ = fi31
        }
        v36 := []int{559, 32767, 836, -1}
        fmt.Println(v12)
        _ = v36
    }
    for fi37 := 0; fi37 < 6; fi37++ {
        var v38 uint32 = 1
        if true {
            v12 = (v12 | 2)
            v4 = (-2.25 * v3)
        }
        _ = fi37
        _ = v38
    }
    v3 = (((v3 - float64(((1.5 - (3.14159 - (float32(((float64(((1e10 * (0.0 + ((0.0 * 1e10) + 1e10))) * 1e10)) + v3) - v3)) + 0.1))) - 0.1))) - v3) / 1.0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
