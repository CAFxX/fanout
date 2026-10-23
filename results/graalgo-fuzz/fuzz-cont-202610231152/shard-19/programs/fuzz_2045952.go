package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 int; F2 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64; F1 int8 }
func genFn0() bool {
    switch ("\n\t" + "") {
    case "":
        var v5 uint16 = 2
        _ = v5
    case "a":
        for fi6 := 0; fi6 < 19; fi6++ {
            switch (fi6 / (((fi6 << ((65535) & 63))) | 1)) {
            case 42:
                switch fi6 {
                case 0:
                    v7 := GenS1{F0: false}
                    _ = v7
                }
            case 2147483647:
                var v8 uint16 = 2
                _ = v8
            case 501:
                var v9 int32 = 0
                _ = v9
            }
            fmt.Println(fi6)
            _ = fi6
        }
    default:
        switch ("\n\t" + (("a" + "a") + ("a" + ((("a" + "hello") + "\n\t") + "hello")))) {
        case "foo bar":
            if true {
                var v10 uint8 = 1
                switch -9223372036854775808 {
                case -9223372036854775808:
                    if false {
                        if false {
                            var v11 int = 0
                            _ = v11
                        }
                        for fi12 := 0; fi12 < 2; fi12++ {
                            v13 := []int{}
                            _ = fi12
                            _ = v13
                        }
                    } else {
                        for fi14 := 0; fi14 < 1; fi14++ {
                            v15 := GenS2{F0: ("foo bar" + "\n\t"), F1: (-(fi14)), F2: 353}
                            v10 = v10
                            _ = fi14
                            _ = v15
                        }
                        var v16 uint16 = 1
                        _ = v16
                    }
                default:
                    switch -304 {
                    case -2147483648:
                        v17 := GenS3{F0: (32767 >= 1)}
                        _ = v17
                    case 32768:
                        fmt.Println(v10)
                    case 127:
                        fmt.Println(v10)
                    }
                }
                var v18 uint32 = 2147483648
                _ = v10
                _ = v18
            } else {
                var v19 uint8 = 128
                var v20 uint8 = 181
                _ = v19
                _ = v20
            }
        case "hello":
            v21 := []int{32768, 1}
            _ = v21
        }
    }
    var v22 string = ("\n\t" + ("\n\t" + "foo bar"))
    switch v22 {
    case "":
        v23 := []int{-526, -9223372036854775808, 9223372036854775807, -340}
        _ = v23
    case "a":
        fmt.Println(v22)
    case "foo bar":
        var v24 int32 = -2147483648
        _ = v24
    default:
        v22 = v22
    }
    if false {
        for fi25 := 0; fi25 < 8; fi25++ {
            var v26 float32 = float32(0.1)
            _ = fi25
            _ = v26
        }
    } else {
        v27 := []int{4294967295}
        var v28 uint64 = 2
        _ = v27
        _ = v28
    }
    v22 = ("" + ("\n\t" + ("foo bar" + "hello")))
    _ = v22
    return (982 <= 2147483648)
}

func main() {
    var v1 int16 = -852
    var v2 float64 = float64((0.0 / 1.5))
    var v3 bool = (4294967295 < 651)
    v4 := make(map[string]string)
    var v29 int16 = -64
    var v30 uint16 = 40
    v31 := GenS2{F0: "hello", F1: -670, F2: v30}
    v3 = true
    v32 := GenS1{F0: ("foo bar" >= "hello")}
    if false {
        v2 += -2.25
        switch "\n\t" {
        case "hello":
            var v33 float32 = 1.5
            _ = v33
        case "a":
            var v34 uint8 = 128
            _ = v34
        }
        if (int32(v30) != -2147483648) {
            for fi35 := 0; fi35 < 7; fi35++ {
                v36 := v4[(("" + ("" + "hello")) + ("hello" + "\n\t"))]
                _ = fi35
                _ = v36
            }
            if (!(!false)) {
                v37 := GenS4{F0: 4294967295}
                v4["foo bar"] = "hello"
                fmt.Println(v29, v1)
                _ = v37
            } else {
                if false {
                    v3 = v3
                } else {
                    v38 := []int{}
                    v2 = float64(0.0)
                    var v39 float64 = 1.5
                    _ = v38
                    _ = v39
                }
            }
        } else {
            switch -2147483648 {
            case 32768:
                for fi40 := 0; fi40 < 1; fi40++ {
                    fmt.Println(v29, v3)
                    for fi41 := 0; fi41 < 6; fi41++ {
                        var v42 uint8 = 255
                        var v43 uint16 = (v30 / ((0) | 1))
                        var v44 uint64 = 42
                        _ = fi41
                        _ = v42
                        _ = v43
                        _ = v44
                    }
                    v3 = (v3 && v3)
                    _ = fi40
                }
            default:
                for fi45 := 0; fi45 < 17; fi45++ {
                    v46 := GenS5{F0: 0, F1: -128}
                    _ = fi45
                    _ = v46
                }
            }
            v47 := GenS3{F0: (v3 || true)}
            _ = v47
        }
    }
    v48 := GenS5{F0: -1, F1: -128}
    var v49 bool = (0 == 1)
    var v50 bool = (true || v3)
    var v51 uint8 = 0
    v52 := GenS5{F0: 519, F1: -100}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v48
    _ = v49
    _ = v50
    _ = v51
    _ = v52
}
