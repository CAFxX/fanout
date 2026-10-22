package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS2 struct { F0 int32; F1 uint64; F2 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint32; F1 int32; F2 bool }
type GenS7 struct { F0 uint32; F1 uint }

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[string]string)
    v3 := make(map[int]int)
    v4 := []string{"\n\t", ("hello" + "")}
    var v5 string = (("" + (("a" + "\n\t") + ("foo bar" + "foo bar"))) + "hello")
    var v6 bool = (v5 != (v5 + v5))
    v7 := GenS1{F0: 1, F1: 65535, F2: 901}
    if v6 {
        if len(v4) > 0 {
            v4[0] = "a"
        }
        v8 := GenS2{F0: 42, F1: 18446744073709551615, F2: 2}
        _ = v8
    }
    if (127 >= 42) {
        switch (v5 + ((v5 + ((v5 + (v5 + v5)) + v5)) + ("\n\t" + "\n\t"))) {
        case "hello":
            if (228 < 4294967295) {
                v6 = (!true)
                fmt.Println(v5, v1)
                for fi9 := 0; fi9 < 20; fi9++ {
                    v10 := GenS1{F0: 1, F1: 247, F2: 2147483647}
                    _ = fi9
                    _ = v10
                }
            } else {
                var v11 uint8 = 1
                _ = v11
            }
        case "\n\t":
            for fi12 := 0; fi12 < 2; fi12++ {
                var v13 int16 = -145
                _ = fi12
                _ = v13
            }
        default:
            v14 := GenS1{F0: 65535, F1: 2, F2: 163}
            _ = v14
        }
        _ = len(v2)
        if (v6 && (((372 >= 255) && (1 >= 0)) || (0 == 42))) {
            v4 = append(v4, (v5 + "\n\t"))
            _ = len(v2)
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    } else {
        fmt.Println(v6, v6)
        var v15 uint8 = 127
        for fi16 := 0; fi16 < 18; fi16++ {
            if v6 {
                v6 = true
                var v17 int64 = 925
                for fi18 := 0; fi18 < 9; fi18++ {
                    v19 := GenS3{F0: v17}
                    v1 = (v1 - ((1e10 * 0.1) + ((v1 * 0.0) * 1.5)))
                    v6 = v6
                    _ = fi18
                    _ = v19
                }
                _ = v17
            } else {
                v20 := GenS4{F0: (fi16 | fi16), F1: 68}
                var v21 int16 = -67
                _ = v20
                _ = v21
            }
            for fi22 := 0; fi22 < 1; fi22++ {
                if ((v15 - (^(v15))) == 1) {
                    var v23 int16 = 1000
                    switch (v5 + "a") {
                    case "hello":
                        v24 := GenS3{F0: -9223372036854775808}
                        _ = v24
                    }
                    switch (fi22 + 4294967295) {
                    case 2147483647:
                        var v25 bool = (!(10 != 754))
                        _ = v25
                    case 42:
                        fmt.Println(fi16)
                    }
                    _ = v23
                }
                if (4294967295 <= 743) {
                    if (42 == -1) {
                        v26 := GenS1{F0: 1, F1: 42, F2: 0}
                        v5 = v5
                        _ = v26
                    } else {
                        v27 := GenS5{F0: 2147483647}
                        v6 = false
                        fmt.Println(v5)
                        _ = v27
                    }
                    var v28 uint16 = 65535
                    _ = v28
                } else {
                    _ = len(v2)
                }
                var v29 uint16 = 32767
                _ = fi22
                _ = v29
            }
            _ = fi16
        }
        _ = v15
    }
    var v30 int8 = -1
    for fi31 := 0; fi31 < 10; fi31++ {
        for fi32 := 0; fi32 < 18; fi32++ {
            v6 = false
            v6 = (true || v6)
            v6 = (!v6)
            var v33 uint = 18446744073709551615
            fmt.Println(v1, fi32, v30)
            _ = fi32
            _ = v33
        }
        v3[(fi31 - fi31)] = fi31
        var v34 bool = v6
        _ = fi31
        _ = v34
    }
    switch v5 {
    case "\n\t":
        if v6 {
            var v35 uint64 = 1
            var v36 float64 = float64((v1 + -2.25))
            _ = v35
            _ = v36
        } else {
            v37 := GenS2{F0: 1, F1: 870, F2: 0}
            var v38 uint = 65535
            for fi39 := 0; fi39 < 8; fi39++ {
                for fi40 := 0; fi40 < 15; fi40++ {
                    if v6 {
                        v30 = 100
                    }
                    var v41 uint16 = 42
                    v38 = v38
                    _ = fi40
                    _ = v41
                }
                _ = fi39
            }
            _ = v37
            _ = v38
        }
    case "a":
        _ = len(v4)
    default:
        var v42 uint16 = 18
        _ = v42
    }
    switch v5 {
    case "hello":
        v43 := GenS6{F0: 42, F1: 420, F2: ((((v5 + v5) + (("a" + v5) + "\n\t")) + v5) <= (((v5 + v5) + (v5 + ((v5 + (v5 + v5)) + (((v5 + (("a" + (v5 + ((v5 + v5) + "hello"))) + (v5 + ""))) + "") + (((("hello" + ("" + v5)) + v5) + v5) + (v5 + v5)))))) + (v5 + v5)))}
        _ = v43
    }
    v44 := GenS7{F0: 42, F1: 395}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v44
    _ = v5
    _ = v6
    _ = v7
}
