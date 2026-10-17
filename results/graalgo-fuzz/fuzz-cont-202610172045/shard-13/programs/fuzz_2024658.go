package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint32; F1 int; F2 uint }
func genFn0() int16 {
    v6 := GenS1{F0: 0}
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := GenS1{F0: -1}
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 0}
    v10 := []int{-834, 42, -816, -204}
    var v11 int32 = -552
    _ = v10
    _ = v11
    _ = v6
    _ = v9
    return 925
}
func genFn1() int8 {
    v12 := []int{823, 0, 32768}
    var v13 int16 = 0
    if false {
        switch 1 {
        case 6:
            for fi14 := 0; fi14 < 7; fi14++ {
                var v15 int16 = (v13 | 1)
                var v16 float64 = 1.5
                var v17 string = "foo bar"
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
            }
        case 4294967295:
            var v18 uint16 = 65535
            _ = v18
        case 128:
            switch (("foo bar" + ("foo bar" + "\n\t")) + (("\n\t" + ("foo bar" + "\n\t")) + "")) {
            case "":
                for fi19 := 0; fi19 < 18; fi19++ {
                    for fi20 := 0; fi20 < 6; fi20++ {
                        var v21 int64 = 764
                        _ = fi20
                        _ = v21
                    }
                    switch fi19 {
                    case -1:
                        v22 := GenS1{F0: 42}
                        _ = v22
                    case -2147483648:
                        var v23 int8 = 0
                        _ = v23
                    case -107:
                        var v24 uint16 = 0
                        _ = v24
                    }
                    v25 := GenS1{F0: 1}
                    _ = fi19
                    _ = v25
                }
            case "\n\t":
                fmt.Println(v13)
            case "a":
                var v26 int8 = 0
                _ = v26
            }
        default:
            v12 = append(v12, 127)
        }
        for fi27 := 0; fi27 < 8; fi27++ {
            v28 := GenS1{F0: 660}
            if true {
                v13 = (v13 % ((v13) | 1))
            } else {
                for fi29 := 0; fi29 < 4; fi29++ {
                    v12 = append(v12, -1)
                    _ = fi29
                }
                var v30 float64 = 1.5
                var v31 int64 = 0
                _ = v30
                _ = v31
            }
            _ = fi27
            _ = v28
        }
    } else {
        var v32 int = 127
        v33 := GenS1{F0: -9223372036854775808}
        v34 := GenS1{F0: 9223372036854775807}
        _ = v32
        _ = v33
        _ = v34
    }
    if false {
        var v35 bool = (!(false || (-214 > int(v13))))
        if true {
            v36 := GenS1{F0: 9223372036854775807}
            switch -868 {
            case 197:
                var v37 uint8 = 196
                _ = v37
            case 973:
                if v35 {
                    switch (("" + ((((("hello" + ("foo bar" + "hello")) + (("hello" + "hello") + "a")) + ("foo bar" + ("a" + (("a" + ("hello" + "")) + "a")))) + "a") + ("a" + "a"))) + "a") {
                    case "":
                        switch "" {
                        case "foo bar":
                            v13 = (v13 - v13)
                        case "":
                            v13 -= v13
                        default:
                            v35 = (!v35)
                        }
                    case "foo bar":
                        var v38 float64 = -2.25
                        _ = v38
                    case "a":
                        fmt.Println(v35, v35)
                    }
                    v13 = genFn0()
                }
            case -1:
                var v39 string = "\n\t"
                _ = v39
            default:
                var v40 float32 = (3.14159 + (1e10 - 1e10))
                _ = v40
            }
            _ = v36
        } else {
            v41 := GenS1{F0: -9223372036854775808}
            _ = v41
        }
        v35 = (4294967295 <= 281)
        _ = v35
    } else {
        v13 = genFn0()
        var v42 uint64 = 1
        var v43 uint8 = 42
        _ = v42
        _ = v43
    }
    for fi44 := 0; fi44 < 9; fi44++ {
        v45 := GenS1{F0: 0}
        if len(v12) > 0 {
            v12[0] = 32767
        }
        _ = fi44
        _ = v45
    }
    _ = v12
    _ = v13
    return 100
}

func main() {
    var v1 uint32 = 99
    var v2 uint64 = 662
    v3 := []uint32{}
    var v4 string = ("" + "foo bar")
    var v5 string = v4
    v46 := GenS2{F0: int(v1), F1: 697}
    if (!(-2147483648 > -2147483648)) {
        var v47 uint16 = uint16(v2)
        v4 = "\n\t"
        for fi48 := 0; fi48 < 1; fi48++ {
            var v49 uint16 = uint16(v2)
            fmt.Println(v1, v49, v1)
            _ = fi48
            _ = v49
        }
        _ = v47
    }
    v5 = (((v5 + ("foo bar" + (v4 + "\n\t"))) + v5) + (v5 + "hello"))
    var v50 uint64 = v2
    fmt.Println(v1, v50)
    var v51 uint32 = 299
    for fi52 := 0; fi52 < 10; fi52++ {
        switch v5 {
        case "\n\t":
            v53 := GenS3{F0: v51, F1: (fi52 | fi52), F2: 255}
            _ = v53
        case "a":
            fmt.Println(fi52, v1, v50)
        default:
            if false {
                v50 -= v50
                var v54 int = (fi52 / (((fi52 | (fi52 ^ fi52))) | 1))
                _ = v54
            }
        }
        var v55 int8 = -128
        _ = fi52
        _ = v55
    }
    fmt.Println(v50, v2, v2)
    for fi56 := 0; fi56 < 1; fi56++ {
        v5 = v5
        _ = fi56
    }
    if true {
        var v57 string = (("hello" + ((("hello" + (v5 + (v4 + v5))) + ((((v5 + ((v4 + (("hello" + "") + "")) + (v4 + v4))) + "\n\t") + "\n\t") + (v4 + v4))) + ("foo bar" + "\n\t"))) + (v4 + ("" + "foo bar")))
        v57 = v57
        v1 = v51
        _ = v57
    }
    v2 = (^(v50))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v50
    _ = v51
}
