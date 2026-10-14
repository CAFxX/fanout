package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int; F1 int64; F2 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 int16; F2 string }
type GenS5 struct { F0 uint64; F1 uint64; F2 uint64 }
type GenS6 struct { F0 uint8; F1 int64 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]bool)
    v3 := []string{("" + "\n\t"), "a", ("hello" + ((("" + ("" + ("\n\t" + ""))) + "") + "")), (("hello" + ("foo bar" + "hello")) + "\n\t")}
    v4 := make(map[uint32]uint32)
    var v5 bool = true
    var v6 float64 = 1.5
    switch ("\n\t" + "foo bar") {
    case "hello":
        if v5 {
            v7 := GenS1{F0: 1, F1: 42}
            var v8 int16 = -32768
            _ = v7
            _ = v8
        }
    case "a":
        v6 += v6
    case "":
        if (127 == -2147483648) {
            switch ("\n\t" + "") {
            case "\n\t":
                for fi9 := 0; fi9 < 12; fi9++ {
                    if (false && (v6 != v6)) {
                        var v10 float32 = 1.5
                        var v11 uint32 = 1
                        switch "foo bar" {
                        case "":
                            var v12 uint32 = 4294967295
                            _ = v12
                        case "hello":
                            var v13 bool = v5
                            _ = v13
                        }
                        _ = v10
                        _ = v11
                    }
                    if true {
                        fmt.Println(v5)
                        fmt.Println(fi9)
                    } else {
                        if v5 {
                            v5 = (1 < 0)
                            v5 = false
                        }
                        var v14 uint8 = 127
                        _ = v14
                    }
                    _ = fi9
                }
            case "foo bar":
                var v15 uint32 = 944
                _ = v15
            case "a":
                for fi16 := 0; fi16 < 10; fi16++ {
                    v6 = v6
                    v17 := GenS1{F0: -9223372036854775808, F1: -1}
                    _ = fi16
                    _ = v17
                }
            default:
                v5 = v5
            }
        } else {
            _ = len(v3)
            var v18 int = 881
            _ = v18
        }
    default:
        v6 = (v6 - 1.5)
    }
    fmt.Println(v6, v6)
    if (128 <= 128) {
        v19 := GenS2{F0: 4294967295, F1: -1, F2: 0}
        v3 = append(v3, "hello")
        _ = v19
    } else {
        v5 = false
        _ = len(v4)
        if len(v3) > 0 {
            v3[0] = ("" + "hello")
        }
    }
    v6 += v6
    var v20 float64 = v6
    v21 := GenS2{F0: 127, F1: 42, F2: 42}
    v5 = true
    v5 = v5
    for fi22 := 0; fi22 < 20; fi22++ {
        var v23 uint16 = 32767
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        _ = len(v3)
        v5 = v5
        _ = fi24
    }
    v20 = ((v6 - 3.14159) / v6)
    var v25 uint8 = 1
    for fi26 := 0; fi26 < 4; fi26++ {
        for fi27 := 0; fi27 < 2; fi27++ {
            if (("foo bar" + (("a" + "\n\t") + "\n\t")) >= "foo bar") {
                if false {
                    var v28 int16 = -462
                    _ = v28
                } else {
                    v29 := GenS1{F0: -467, F1: 2147483647}
                    _ = v29
                }
                v30 := GenS3{F0: -9223372036854775808}
                v31 := GenS4{F0: 891, F1: -32768, F2: "hello"}
                _ = v30
                _ = v31
            } else {
                var v32 int8 = 45
                if v5 {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    v5 = ("a" == "hello")
                    v33 := GenS3{F0: -9223372036854775808}
                    if (!true) {
                        v4[2147483648] = 2147483647
                        if false {
                            var v34 float32 = 1e10
                            _ = v34
                        }
                        v35 := GenS5{F0: 169, F1: 1, F2: 0}
                        _ = v35
                    } else {
                        fmt.Println(v5)
                        var v36 int16 = -328
                        _ = v36
                    }
                    switch ("foo bar" + ("foo bar" + (("hello" + "a") + ((("foo bar" + "") + ("a" + ("\n\t" + ("foo bar" + "foo bar")))) + "")))) {
                    case "foo bar":
                        var v37 int16 = 1
                        _ = v37
                    case "hello":
                        var v38 uint32 = 4294967295
                        _ = v38
                    }
                    if (!(false && (v5 && ((42 < -226) || v5)))) {
                        v25 = 245
                        v25 = v25
                    }
                    _ = v33
                }
                _ = v32
            }
            _ = fi27
        }
        _ = fi26
    }
    switch 229 {
    case 128:
        fmt.Println(v20, v6, v25)
    default:
        if v5 {
            v5 = false
            fmt.Println(v5)
        }
    }
    for fi39 := 0; fi39 < 16; fi39++ {
        v6 += v20
        for fi40 := 0; fi40 < 17; fi40++ {
            var v41 bool = v5
            _ = fi40
            _ = v41
        }
        v42 := GenS6{F0: v25, F1: 9223372036854775807}
        _ = fi39
        _ = v42
    }
    _ = len(v3)
    for fi43 := 0; fi43 < 20; fi43++ {
        v20 = v6
        if (!(false && (963 < 0))) {
            v5 = (!v5)
            fmt.Println(fi43)
        }
        _ = fi43
    }
    v3 = append(v3, "a")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
