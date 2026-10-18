package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint32 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 int16; F1 uint; F2 uint }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]bool)
    v3 := make(map[uint32]uint32)
    var v4 float64 = 1e10
    var v5 int16 = 0
    var v6 string = ("\n\t" + "")
    v6 = ("hello" + (v6 + v6))
    v7 := GenS1{F0: 0, F1: -9223372036854775808, F2: 2147483647}
    v8 := []int{4294967295, 32768}
    fmt.Println(v6)
    if (!(false || (!true))) {
        v9 := GenS1{F0: 4294967295, F1: -9223372036854775808, F2: 2}
        _ = v9
    }
    v8 = append(v8, -932)
    v10 := v1[v6]
    if (v5 <= v5) {
        v11 := GenS2{F0: 65535, F1: 97}
        _ = v11
    }
    v12 := v1[(v10 + "\n\t")]
    for fi13 := 0; fi13 < 10; fi13++ {
        v5 -= 1
        _ = fi13
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        v15 := GenS1{F0: 42, F1: (fi14 - -789), F2: 65535}
        _ = fi14
        _ = v15
    }
    switch v12 {
    case "a":
        var v16 int = 1
        _ = v16
    default:
        if (v12 <= (v12 + ((v10 + "") + (v10 + v12)))) {
            switch (((v12 + "") + "a") + v12) {
            case "\n\t":
                for fi17 := 0; fi17 < 5; fi17++ {
                    if true {
                        v18 := GenS3{F0: -868, F1: 1}
                        v5 -= -1
                        _ = v18
                    }
                    for fi19 := 0; fi19 < 19; fi19++ {
                        v5 -= (v5 * (v5 << ((2) & 63)))
                        fmt.Println(fi17, fi17, v12)
                        switch v10 {
                        case "a":
                            var v20 int = fi17
                            _ = v20
                        }
                        var v21 int64 = 0
                        var v22 uint16 = 2
                        switch ((v12 + "") + v10) {
                        case "a":
                            _ = len(v8)
                        }
                        switch "a" {
                        case "foo bar":
                            v3[537] = 0
                        case "":
                            if false {
                                switch ("a" + ("hello" + v6)) {
                                case "\n\t":
                                    v23 := GenS4{F0: ((v12 + "\n\t") + ((("a" + (v12 + "\n\t")) + "\n\t") + "a")), F1: 2}
                                    _ = v23
                                case "a":
                                    var v24 int8 = -128
                                    _ = v24
                                }
                                var v25 int32 = 2147483647
                                _ = v25
                            }
                        case "\n\t":
                            var v26 int = int(v21)
                            _ = v26
                        }
                        _ = fi19
                        _ = v21
                        _ = v22
                    }
                    _ = fi17
                }
            case "hello":
                for fi27 := 0; fi27 < 10; fi27++ {
                    var v28 bool = true
                    v12 = v10
                    _ = fi27
                    _ = v28
                }
            case "foo bar":
                var v29 bool = (!false)
                _ = v29
            default:
                var v30 uint32 = 0
                _ = v30
            }
            fmt.Println(v12, v12, v6)
        }
    }
    v31 := GenS5{F0: 1000, F1: 51, F2: 42}
    if ((((false || (false || true)) || (false || (!(!false)))) || (true || false)) || (-9223372036854775808 == -675)) {
        v6 = v12
        for fi32 := 0; fi32 < 13; fi32++ {
            if false {
                switch "a" {
                case "\n\t":
                    v4 = float64(3.14159)
                case "a":
                    v6 = v6
                }
                v5 += v5
            }
            if false {
                switch (v6 + (v6 + (v12 + "a"))) {
                case "a":
                    if true {
                        var v33 uint32 = 4294967295
                        _ = len(v8)
                        var v34 uint = 65535
                        _ = v33
                        _ = v34
                    }
                case "foo bar":
                    var v35 int = (fi32 << ((4294967295) & 63))
                    _ = v35
                case "hello":
                    switch ("hello" + "foo bar") {
                    case "a":
                        var v36 int = 32767
                        _ = v36
                    case "\n\t":
                        for fi37 := 0; fi37 < 13; fi37++ {
                            var v38 float32 = 3.14159
                            _ = fi37
                            _ = v38
                        }
                    case "":
                        var v39 uint8 = 255
                        _ = v39
                    default:
                        v5 = v5
                    }
                default:
                    if (0 != 0) {
                        for fi40 := 0; fi40 < 18; fi40++ {
                            v41 := GenS5{F0: v5, F1: 4294967295, F2: 42}
                            _ = fi40
                            _ = v41
                        }
                    } else {
                        v8 = append(v8, (fi32 + fi32))
                        v12 = (v10 + (("\n\t" + "") + "hello"))
                        for fi42 := 0; fi42 < 7; fi42++ {
                            var v43 int64 = 9223372036854775807
                            v43 = (v43 - v43)
                            var v44 int32 = -1
                            v4 = float64((-2.25 + 0.0))
                            _ = fi42
                            _ = v43
                            _ = v44
                        }
                    }
                }
            } else {
                var v45 int16 = (v5 & v5)
                _ = v45
            }
            _ = fi32
        }
        var v46 bool = true
        _ = v46
    } else {
        var v47 uint16 = 881
        for fi48 := 0; fi48 < 20; fi48++ {
            v6 = ("foo bar" + ("\n\t" + "\n\t"))
            v8 = append(v8, (^(fi48)))
            if (false && false) {
                v5 = (v5 * v5)
                v5 = 32767
            }
            _ = fi48
        }
        _ = v47
    }
    v12 = ("foo bar" + "hello")
    _ = len(v8)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
