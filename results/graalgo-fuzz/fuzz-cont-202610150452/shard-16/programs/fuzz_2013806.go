package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 uint32; F2 int16 }
type GenS4 struct { F0 bool; F1 int8; F2 bool }
type GenS5 struct { F0 int16 }
func genFn0() int16 {
    var v7 string = (("foo bar" + "") + "")
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 float32 = 0.1
        _ = fi8
        _ = v9
    }
    if false {
        if false {
            v7 = v7
        }
        v7 = (((v7 + v7) + (v7 + (("hello" + "hello") + "hello"))) + (v7 + "foo bar"))
        v7 = "hello"
    }
    _ = v7
    return -110
}
func genFn1() string {
    v10 := GenS1{F0: (232 <= 0), F1: 255, F2: 42}
    switch "foo bar" {
    case "hello":
        for fi11 := 0; fi11 < 1; fi11++ {
            switch (-(fi11)) {
            case -858:
                v12 := []int{fi11}
                _ = v12
            default:
                v13 := GenS1{F0: true, F1: 32767, F2: 65535}
                _ = v13
            }
            if (true || false) {
                for fi14 := 0; fi14 < 20; fi14++ {
                    var v15 string = "hello"
                    switch 32768 {
                    case -2147483648:
                        v16 := []int{1, (fi11 - fi14), fi11, 128}
                        _ = v16
                    }
                    v15 = ((v15 + "foo bar") + (("hello" + (v15 + v15)) + (v15 + v15)))
                    _ = fi14
                    _ = v15
                }
                if false {
                    if true {
                        var v17 uint16 = 42
                        _ = v17
                    }
                }
            }
            _ = fi11
        }
    case "foo bar":
        var v18 string = "a"
        _ = v18
    default:
        v19 := GenS1{F0: true, F1: 2, F2: 65535}
        _ = v19
    }
    _ = v10
    return ("a" + "\n\t")
}

func main() {
    var v1 float32 = float32(float64(3.14159))
    v2 := []int{2147483647, 0}
    var v3 string = "hello"
    var v4 int32 = -1
    var v5 int = -2147483648
    var v6 uint = 65535
    v6 -= (^(v6))
    fmt.Println(v1, v6)
    var v20 uint32 = 0
    var v21 int64 = -725
    var v22 uint16 = 524
    v23 := GenS2{F0: 0}
    for fi24 := 0; fi24 < 19; fi24++ {
        v5 = -9223372036854775808
        switch ("" + v3) {
        case "hello":
            switch v3 {
            case "hello":
                switch (fi24 + v5) {
                case 42:
                    var v25 uint32 = 0
                    _ = v25
                case -425:
                    var v26 float32 = v1
                    _ = v26
                case -530:
                    v21 = (v21 & 42)
                }
            case "":
                v27 := GenS3{F0: (v21 % (((^(v21))) | 1)), F1: (v20 - v20), F2: -513}
                _ = v27
            case "a":
                var v28 int = (fi24 & fi24)
                _ = v28
            }
        case "":
            for fi29 := 0; fi29 < 17; fi29++ {
                v5 -= fi24
                v3 = (v3 + "\n\t")
                _ = fi29
            }
        case "foo bar":
            _ = len(v2)
        }
        _ = fi24
    }
    var v30 bool = true
    for fi31 := 0; fi31 < 17; fi31++ {
        fmt.Println(v21)
        _ = fi31
    }
    switch (v5 >> (((v6 % ((35) | 1))) & 63)) {
    case -283:
        var v32 uint16 = v22
        _ = v32
    case -2147483648:
        if v30 {
            v2 = append(v2, 127)
            var v33 uint64 = 18446744073709551615
            _ = v33
        }
    case -1:
        var v34 float32 = 0.1
        _ = v34
    }
    if len(v2) > 0 {
        v2[0] = v5
    }
    if v30 {
        if (!true) {
            v35 := GenS4{F0: (!false), F1: 0, F2: (false || (!(v1 == -2.25)))}
            _ = v35
        } else {
            var v36 int64 = v21
            _ = v36
        }
        for fi37 := 0; fi37 < 7; fi37++ {
            for fi38 := 0; fi38 < 6; fi38++ {
                var v39 float64 = ((3.14159 - 0.0) - 0.0)
                switch (v5 >> (((^(v6))) & 63)) {
                case -36:
                    v2 = append(v2, v5)
                default:
                    switch v3 {
                    case "foo bar":
                        v40 := GenS5{F0: 1}
                        _ = v40
                    case "\n\t":
                        switch v3 {
                        case "hello":
                            fmt.Println(v20, v30, v3)
                        case "\n\t":
                            v1 += float32(v39)
                            v4 = -1
                            v2 = append(v2, (fi38 ^ int(v22)))
                        }
                    default:
                        fmt.Println(v3)
                    }
                }
                _ = fi38
                _ = v39
            }
            _ = fi37
        }
    } else {
        v22 = v22
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v20 += v20
    }
    fmt.Println(v5, v21, v20)
    var v41 int = (v5 / ((-1) | 1))
    v2 = append(v2, v5)
    v30 = false
    v22 -= v22
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v30
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
