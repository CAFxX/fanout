package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 string; F2 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16; F1 int64 }
func genFn0() int16 {
    v5 := GenS1{F0: 643}
    var v6 uint64 = 687
    for fi7 := 0; fi7 < 5; fi7++ {
        fmt.Println(fi7, v6)
        v8 := GenS1{F0: 0}
        _ = fi7
        _ = v8
    }
    var v9 uint16 = 32767
    v10 := GenS2{F0: 0, F1: "foo bar", F2: 777}
    _ = v10
    _ = v5
    _ = v6
    _ = v9
    return 185
}
func genFn1(p0 int64, p1 int32, p2 uint32) string {
    for fi11 := 0; fi11 < 5; fi11++ {
        if ("\n\t" > "a") {
            var v12 int32 = 2147483647
            p1 = (v12 / (((v12 - p1)) | 1))
            _ = v12
        }
        p1 = p1
        _ = fi11
    }
    if false {
        fmt.Println(p1, p2)
    }
    v13 := GenS3{F0: false}
    p2 -= p2
    if (false && true) {
        p2 = 134
    } else {
        p0 = p0
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return "foo bar"
}

func main() {
    var v1 int16 = 0
    v2 := make(map[string]string)
    var v3 int8 = -128
    v4 := []string{"hello", "\n\t", "\n\t"}
    for fi14 := 0; fi14 < 15; fi14++ {
        var v15 int64 = 9223372036854775807
        for fi16 := 0; fi16 < 19; fi16++ {
            switch ("hello" + ("hello" + "\n\t")) {
            case "hello":
                if false {
                    fmt.Println(v1)
                } else {
                    v17 := GenS1{F0: 751}
                    switch "hello" {
                    case "a":
                        var v18 int32 = -2147483648
                        _ = v18
                    }
                    _ = v17
                }
            case "foo bar":
                var v19 float32 = 1e10
                _ = v19
            }
            _ = fi16
        }
        fmt.Println(fi14, v15)
        _ = fi14
        _ = v15
    }
    for fi20 := 0; fi20 < 10; fi20++ {
        v21 := GenS1{F0: 42}
        _ = fi20
        _ = v21
    }
    if false {
        if (!(!true)) {
            for fi22 := 0; fi22 < 11; fi22++ {
                v23 := GenS4{F0: (1 == 42), F1: 42}
                v1 = v1
                _ = fi22
                _ = v23
            }
            v24 := GenS3{F0: true}
            _ = v24
        }
        var v25 int16 = v1
        fmt.Println(v3, v3, v3)
        _ = v25
    } else {
        v1 += genFn0()
        if true {
            switch "foo bar" {
            case "\n\t":
                for fi26 := 0; fi26 < 3; fi26++ {
                    var v27 float32 = ((0.1 + (1e10 / (1e10 - -2.25))) / 0.1)
                    var v28 uint = 4294967295
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            case "hello":
                v29 := GenS5{F0: 2147483647}
                _ = v29
            }
            var v30 int = 9223372036854775807
            _ = v30
        }
        var v31 int64 = 0
        _ = v31
    }
    var v32 string = "hello"
    for fi33 := 0; fi33 < 14; fi33++ {
        if true {
            _ = len(v4)
        }
        v3 -= v3
        _ = fi33
    }
    if ((-(v3)) > v3) {
        switch -422 {
        case -1:
            var v34 string = v32
            _ = v34
        default:
            var v35 float32 = 1e10
            _ = v35
        }
    }
    if true {
        v3 += (v3 & v3)
    } else {
        switch (v32 + v32) {
        case "foo bar":
            v1 += genFn0()
        default:
            for fi36 := 0; fi36 < 8; fi36++ {
                v1 -= (v1 + 0)
                for fi37 := 0; fi37 < 16; fi37++ {
                    v3 += (v3 & v3)
                    var v38 float32 = (1e10 - (3.14159 / float32(3.14159)))
                    _ = fi37
                    _ = v38
                }
                switch v32 {
                case "a":
                    v39 := GenS6{F0: 2, F1: 9223372036854775807}
                    _ = v39
                default:
                    fmt.Println(v1, v1, v1)
                }
                _ = fi36
            }
        }
        var v40 uint16 = 32767
        _ = v40
    }
    fmt.Println(v32, v32, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
}
