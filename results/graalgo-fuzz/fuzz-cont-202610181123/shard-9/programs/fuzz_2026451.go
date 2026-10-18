package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 string }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 int; F1 int8; F2 uint64 }
type GenS4 struct { F0 uint64; F1 uint16; F2 uint16 }
type GenS5 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS6 struct { F0 bool; F1 uint8; F2 int }
func genFn0(p0 bool, p1 int) uint32 {
    fmt.Println(p1, p1)
    p1 = (p1 * p1)
    _ = p0
    _ = p1
    return 4294967295
}
func genFn1(p0 uint32) int8 {
    v5 := []int{128}
    v5 = append(v5, 42)
    _ = p0
    _ = v5
    return 1
}

func main() {
    v1 := make(map[string]int)
    var v2 float32 = -2.25
    v3 := []string{"\n\t", "a"}
    v4 := make(map[uint32]uint32)
    if false {
        var v6 uint8 = 200
        for fi7 := 0; fi7 < 8; fi7++ {
            for fi8 := 0; fi8 < 20; fi8++ {
                if false {
                    v4[1] = 971
                }
                _ = fi8
            }
            v1[("hello" + "foo bar")] = 32768
            _ = len(v3)
            _ = fi7
        }
        fmt.Println(v6)
        _ = v6
    } else {
        v9 := GenS1{F0: false, F1: 65535, F2: "hello"}
        if true {
            v2 += ((v2 + float32((-2.25 - 3.14159))) - v2)
            var v10 float32 = 1e10
            _ = v10
        } else {
            var v11 uint16 = 32767
            var v12 uint32 = 42
            v11 = (v11 + v11)
            _ = v11
            _ = v12
        }
        _ = v9
    }
    if false {
        delete(v4, 2147483648)
    } else {
        var v13 uint = 42
        if true {
            v14 := GenS2{F0: "", F1: ""}
            _ = v14
        } else {
            switch ("hello" + "a") {
            case "\n\t":
                v2 = v2
            case "":
                fmt.Println(v13, v2)
            case "hello":
                var v15 int64 = 42
                _ = v15
            }
            switch "\n\t" {
            case "a":
                var v16 int32 = -2147483648
                _ = v16
            }
            var v17 uint8 = 42
            _ = v17
        }
        _ = v13
    }
    v2 = v2
    _ = len(v4)
    var v18 string = "hello"
    v19 := GenS3{F0: 127, F1: 100, F2: 42}
    v3 = append(v3, ((("a" + "") + v18) + (((v18 + "a") + (v18 + "\n\t")) + "a")))
    for fi20 := 0; fi20 < 5; fi20++ {
        if (!false) {
            switch "hello" {
            case "\n\t":
                if true {
                    v2 += v2
                }
            case "hello":
                switch ("" + "a") {
                case "\n\t":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case "":
                    if (false || false) {
                        v21 := GenS4{F0: 18446744073709551615, F1: 38, F2: 65535}
                        _ = v21
                    }
                case "hello":
                    v22 := GenS5{F0: 255, F1: 866, F2: -2147483648}
                    _ = v22
                default:
                    var v23 uint32 = 2147483647
                    _ = v23
                }
            default:
                if false {
                    for fi24 := 0; fi24 < 20; fi24++ {
                        var v25 string = ("foo bar" + v18)
                        v26 := GenS3{F0: fi24, F1: 42, F2: 550}
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                    var v27 string = v18
                    _ = v27
                }
            }
            var v28 float64 = -2.25
            if true {
                for fi29 := 0; fi29 < 15; fi29++ {
                    v28 -= 0.0
                    v18 = ""
                    v30 := GenS6{F0: false, F1: 255, F2: -449}
                    switch (fi29 << ((2) & 63)) {
                    case -418:
                        var v31 string = ((v18 + (v18 + v18)) + "\n\t")
                        _ = v31
                    case -912:
                        var v32 uint8 = 42
                        _ = v32
                    }
                    fmt.Println(v2, v2)
                    _ = fi29
                    _ = v30
                }
            }
            _ = v28
        } else {
            if false {
                var v33 int64 = 0
                v33 += v33
                v2 = v2
                _ = v33
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v34 := GenS4{F0: 4294967295, F1: 0, F2: 308}
            _ = v34
        }
        var v35 uint8 = 0
        _ = fi20
        _ = v35
    }
    for fi36 := 0; fi36 < 6; fi36++ {
        v18 = v18
        var v37 float32 = float32((1.5 / 0.1))
        _ = fi36
        _ = v37
    }
    for fi38 := 0; fi38 < 14; fi38++ {
        v2 = v2
        v2 = 0.0
        _ = fi38
    }
    for fi39 := 0; fi39 < 14; fi39++ {
        if (((fi39 >> ((42) & 63)) > (^(fi39))) || false) {
            _ = len(v4)
        } else {
            _ = len(v3)
            switch (v18 + v18) {
            case "":
                v18 = "hello"
            case "foo bar":
                v40 := GenS6{F0: (true && true), F1: 1, F2: 9223372036854775807}
                _ = v40
            case "a":
                v2 = v2
            default:
                v2 -= (1.5 + ((float32(-2.25) / v2) - v2))
            }
        }
        _ = fi39
    }
    for fi41 := 0; fi41 < 19; fi41++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v42 := GenS1{F0: false, F1: 255, F2: "\n\t"}
        _ = fi41
        _ = v42
    }
    if len(v3) > 0 {
        v3[0] = v18
    }
    var v43 int16 = 1000
    fmt.Println(v43, v2, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v43
}
