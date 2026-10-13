package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint; F1 int; F2 uint16 }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 string; F1 bool; F2 string }
func genFn0(p0 uint) uint32 {
    p0 = (p0 * p0)
    v6 := []int{4294967295, 128, -530, 658}
    _ = p0
    _ = v6
    return 2
}
func genFn1(p0 uint32, p1 bool) uint16 {
    var v7 bool = p1
    fmt.Println(p0)
    v8 := []int{}
    p0 += 4294967295
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 255
}

func main() {
    v1 := make(map[int]string)
    var v2 int32 = -669
    var v3 int64 = 9223372036854775807
    v4 := make(map[string]int)
    var v5 uint64 = 654
    v9 := GenS1{F0: 65535, F1: 127}
    if true {
        v10 := []int{}
        _ = v10
    }
    if true {
        if true {
            v11 := GenS2{F0: 201, F1: 32767, F2: 508}
            _ = v11
        }
    } else {
        switch "foo bar" {
        case "a":
            if true {
                var v12 bool = (("a" + "a") <= "\n\t")
                for fi13 := 0; fi13 < 9; fi13++ {
                    switch "\n\t" {
                    case "\n\t":
                        v2 = v2
                    case "a":
                        v14 := GenS3{F0: 18446744073709551615, F1: "hello"}
                        _ = v14
                    case "hello":
                        for fi15 := 0; fi15 < 16; fi15++ {
                            if (v3 == v3) {
                                var v16 uint32 = 2
                                v2 -= v2
                                _ = v16
                            } else {
                                switch fi15 {
                                case -1:
                                    v3 += (v3 + v3)
                                case 127:
                                    v12 = (0 >= 42)
                                }
                                switch fi15 {
                                case 32767:
                                    v12 = (!(100 == 1))
                                case 0:
                                    v2 = -2147483648
                                }
                            }
                            _ = fi15
                        }
                    default:
                        for fi17 := 0; fi17 < 11; fi17++ {
                            v18 := v4[("a" + (("" + ((("foo bar" + "") + "a") + "hello")) + "hello"))]
                            fmt.Println(v12)
                            v19 := []int{}
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                    }
                    fmt.Println(fi13, v5)
                    if v12 {
                        fmt.Println(v2, fi13)
                        fmt.Println(fi13)
                    }
                    _ = fi13
                }
                fmt.Println(v3, v12)
                _ = v12
            } else {
                if (-123 <= -9223372036854775808) {
                    v3 += (v3 + (v3 + 9223372036854775807))
                } else {
                    fmt.Println(v3, v3, v3)
                }
            }
        case "":
            switch ("\n\t" + "hello") {
            case "":
                v20 := GenS4{F0: "hello", F1: ("\n\t" == "\n\t"), F2: (("" + "\n\t") + (((((("hello" + "a") + "foo bar") + "\n\t") + "hello") + "\n\t") + (("\n\t" + "\n\t") + "")))}
                _ = v20
            case "hello":
                v21 := []int{2147483647, 627, 9223372036854775807}
                _ = v21
            }
        case "foo bar":
            var v22 uint8 = 127
            _ = v22
        default:
            for fi23 := 0; fi23 < 4; fi23++ {
                var v24 float32 = (3.14159 - (-2.25 + -2.25))
                _ = fi23
                _ = v24
            }
        }
    }
    switch -2147483648 {
    case 0:
        if false {
            var v25 float64 = 0.0
            _ = v25
        } else {
            fmt.Println(v3)
            if (true && true) {
                v4[("a" + ("foo bar" + "hello"))] = -3
            } else {
                v26 := []int{}
                for fi27 := 0; fi27 < 8; fi27++ {
                    var v28 float64 = -2.25
                    fmt.Println(v28, v28)
                    _ = fi27
                    _ = v28
                }
                for fi29 := 0; fi29 < 10; fi29++ {
                    fmt.Println(fi29, v3)
                    v3 = 0
                    var v30 int64 = (v3 ^ v3)
                    _ = fi29
                    _ = v30
                }
                _ = v26
            }
            v4[(("foo bar" + "foo bar") + "hello")] = 9223372036854775807
        }
    case 116:
        if (("foo bar" + "\n\t") != "a") {
            var v31 float64 = float64(-2.25)
            var v32 int32 = v2
            v5 = (v5 / ((v5) | 1))
            _ = v31
            _ = v32
        }
    default:
        if ("hello" >= "") {
            v33 := v1[42]
            v33 = v33
            _ = v33
        }
    }
    for fi34 := 0; fi34 < 1; fi34++ {
        v5 = 1
        v35 := v4[("hello" + "hello")]
        _ = fi34
        _ = v35
    }
    var v36 float64 = 3.14159
    var v37 float32 = 3.14159
    fmt.Println(v36)
    var v38 int8 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v9
}
