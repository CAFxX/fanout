package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 int64; F2 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS6 struct { F0 uint32; F1 int32 }
func genFn0(p0 uint32) uint16 {
    if ((3.14159 >= (1e10 + (float32((-2.25 * (-2.25 / 1.0))) * 3.14159))) && (1 == -2147483648)) {
        if true {
            switch "hello" {
            case "hello":
                if ((0.0 != 3.14159) || (118 < 32767)) {
                    var v4 bool = (!((-11 >= -26) && ("hello" < (("\n\t" + "") + "\n\t"))))
                    _ = v4
                } else {
                    v5 := []int{9223372036854775807}
                    _ = v5
                }
            case "":
                v6 := GenS1{F0: 49, F1: 127, F2: "\n\t"}
                _ = v6
            }
            var v7 int8 = 1
            switch "a" {
            case "":
                var v8 int64 = -745
                _ = v8
            case "hello":
                v7 = v7
            }
            _ = v7
        }
        switch ("\n\t" + "a") {
        case "a":
            for fi9 := 0; fi9 < 4; fi9++ {
                var v10 float64 = (((float64((1.5 * (1e10 * (3.14159 + (float32(0.1) + 0.1))))) / (1e10 / -2.25)) / 1.5) / ((0.1 + 3.14159) * 1e10))
                var v11 int = fi9
                _ = fi9
                _ = v10
                _ = v11
            }
        default:
            var v12 uint16 = 2
            _ = v12
        }
    }
    v13 := []int{-2147483648, 4294967295, -1, 128}
    p0 -= p0
    if false {
        p0 += p0
        v14 := GenS1{F0: 42, F1: 42, F2: "\n\t"}
        switch "foo bar" {
        case "a":
            p0 = p0
        }
        _ = v14
    }
    _ = p0
    _ = v13
    return 2
}

func main() {
    var v1 uint16 = 236
    var v2 uint16 = v1
    var v3 int16 = -143
    v15 := GenS2{F0: "\n\t"}
    for fi16 := 0; fi16 < 13; fi16++ {
        if (1 < 225) {
            fmt.Println(fi16, v1, v3)
            switch "foo bar" {
            case "\n\t":
                v17 := GenS1{F0: 2, F1: 1, F2: ("\n\t" + (("foo bar" + ("" + "a")) + "\n\t"))}
                _ = v17
            default:
                v2 = v1
            }
            if (!("a" >= ("foo bar" + "foo bar"))) {
                if true {
                    for fi18 := 0; fi18 < 9; fi18++ {
                        v19 := []int{(fi18 + 867), (fi18 + -2147483648)}
                        _ = fi18
                        _ = v19
                    }
                }
                fmt.Println(v2)
            }
        }
        if false {
            v20 := GenS1{F0: 42, F1: 1, F2: (("a" + ("hello" + "")) + (("foo bar" + (("a" + "hello") + "")) + ("hello" + ((("hello" + "foo bar") + ("" + (("hello" + "foo bar") + ("foo bar" + "hello")))) + ("a" + (("\n\t" + "") + ("foo bar" + "")))))))}
            _ = v20
        } else {
            for fi21 := 0; fi21 < 12; fi21++ {
                switch "foo bar" {
                case "hello":
                    var v22 int64 = -9223372036854775808
                    _ = v22
                case "":
                    var v23 float64 = 0.0
                    _ = v23
                default:
                    switch "a" {
                    case "\n\t":
                        switch 32767 {
                        case 9223372036854775807:
                            v3 = v3
                            if (("foo bar" + "hello") == "") {
                                var v24 uint64 = 18446744073709551615
                                v25 := []int{fi16, fi16}
                                var v26 uint8 = 42
                                _ = v24
                                _ = v25
                                _ = v26
                            } else {
                                fmt.Println(v3, v3, v2)
                            }
                        }
                    }
                }
                _ = fi21
            }
            v3 += (-(v3))
        }
        var v27 int16 = 42
        _ = fi16
        _ = v27
    }
    if (!(false || (4294967295 >= 2))) {
        for fi28 := 0; fi28 < 10; fi28++ {
            if (true && true) {
                for fi29 := 0; fi29 < 20; fi29++ {
                    v30 := GenS3{F0: fi29, F1: -911, F2: true}
                    _ = fi29
                    _ = v30
                }
                v31 := GenS4{F0: 65535}
                _ = v31
            }
            _ = fi28
        }
        v32 := GenS5{F0: 330, F1: 18446744073709551615, F2: -46}
        _ = v32
    } else {
        v3 = v3
        v2 -= 255
        v33 := GenS4{F0: v2}
        _ = v33
    }
    v34 := GenS1{F0: 255, F1: 0, F2: ("\n\t" + "hello")}
    if false {
        for fi35 := 0; fi35 < 15; fi35++ {
            fmt.Println(v1)
            _ = fi35
        }
        var v36 int32 = 1
        _ = v36
    } else {
        var v37 uint8 = 127
        _ = v37
    }
    if true {
        fmt.Println(v1)
    } else {
        v38 := GenS6{F0: 4294967295, F1: 1}
        _ = v38
    }
    var v39 uint32 = 606
    var v40 uint8 = 81
    var v41 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v34
    _ = v39
    _ = v40
    _ = v41
}
