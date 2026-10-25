package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int16; F1 string }
func genFn0() int8 {
    var v6 float32 = 1.5
    var v7 int16 = 1000
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 int64 = -343
            v10 = v10
            v11 := []int{(-(fi9)), fi8}
            _ = fi9
            _ = v10
            _ = v11
        }
        if false {
            v12 := []int{-9223372036854775808, fi8, fi8}
            v6 = (v6 - (float32(1e10) - v6))
            v6 += float32((((float64(v6) * ((float64(float32(1e10)) + ((1e10 + 1.5) * 1e10)) + 1.5)) / (3.14159 - 0.0)) / 1.0))
            _ = v12
        } else {
            var v13 uint8 = 127
            fmt.Println(v6, fi8, v6)
            if true {
                fmt.Println(v6, v13, v7)
            }
            _ = v13
        }
        var v14 float32 = 0.1
        _ = fi8
        _ = v14
    }
    var v15 float64 = 0.1
    v16 := GenS1{F0: (("foo bar" + "") + "a"), F1: "a"}
    _ = v15
    _ = v16
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 uint8, p1 uint32, p2 uint32) int8 {
    fmt.Println(p1, p1)
    fmt.Println(p2, p2, p1)
    for fi17 := 0; fi17 < 9; fi17++ {
        p0 = (^(p0))
        p2 -= p1
        _ = fi17
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    v1 := []uint32{2147483647, 2147483647, 263}
    v2 := make(map[uint32]string)
    var v3 int32 = -1
    v4 := []string{"foo bar"}
    v5 := []string{"\n\t", (("foo bar" + "\n\t") + "foo bar"), (((("\n\t" + (("" + "") + "foo bar")) + "a") + "foo bar") + "\n\t"), (((("\n\t" + "foo bar") + (("" + "") + "a")) + "") + ("a" + "foo bar"))}
    v3 += (v3 ^ (v3 - v3))
    v3 = v3
    for fi18 := 0; fi18 < 6; fi18++ {
        for fi19 := 0; fi19 < 11; fi19++ {
            v3 -= v3
            _ = fi19
        }
        var v20 uint16 = 255
        _ = fi18
        _ = v20
    }
    if (!(32767 <= 65535)) {
        switch 4294967295 {
        case 32767:
            v3 = -715
        case 46:
            var v21 float32 = 1e10
            _ = v21
        case 459:
            v22 := GenS1{F0: "a", F1: "a"}
            _ = v22
        default:
            switch "\n\t" {
            case "foo bar":
                v23 := GenS2{F0: 2, F1: 1}
                _ = v23
            case "\n\t":
                if false {
                    for fi24 := 0; fi24 < 9; fi24++ {
                        fmt.Println(v3, v3)
                        v3 = v3
                        if ((!(false && (!false))) || true) {
                            v25 := GenS1{F0: (("\n\t" + "hello") + "hello"), F1: "foo bar"}
                            if true {
                                switch 42 {
                                case 426:
                                    switch (^(fi24)) {
                                    case 9223372036854775807:
                                        var v26 uint64 = 4294967295
                                        _ = v26
                                    default:
                                        var v27 uint64 = 18446744073709551615
                                        _ = v27
                                    }
                                }
                                if ((v3 ^ 1) == v3) {
                                    var v28 int16 = -915
                                    _ = v28
                                } else {
                                    fmt.Println(fi24, fi24)
                                }
                            }
                            _ = v25
                        }
                        _ = fi24
                    }
                    v2[2147483647] = "hello"
                    if true {
                        v29 := GenS1{F0: "", F1: "\n\t"}
                        _ = v29
                    }
                }
            }
        }
    } else {
        for fi30 := 0; fi30 < 1; fi30++ {
            if true {
                v31 := GenS1{F0: "foo bar", F1: (("a" + (("foo bar" + "foo bar") + "")) + (("\n\t" + "foo bar") + ("\n\t" + ((((("" + ("a" + "a")) + "a") + "") + "a") + ("hello" + "foo bar")))))}
                var v32 bool = (true || false)
                _ = v31
                _ = v32
            } else {
                if (875 == 2) {
                    v3 += (v3 / (((v3 + v3)) | 1))
                    var v33 int8 = genFn0()
                    _ = v33
                }
            }
            _ = fi30
        }
    }
    switch "" {
    case "foo bar":
        switch "" {
        case "hello":
            var v34 float64 = float64((3.14159 + -2.25))
            _ = v34
        }
    case "":
        _ = len(v2)
    case "\n\t":
        var v35 float32 = -2.25
        _ = v35
    default:
        for fi36 := 0; fi36 < 10; fi36++ {
            var v37 string = "\n\t"
            switch (("foo bar" + (((v37 + v37) + "") + v37)) + (v37 + ((("a" + (v37 + v37)) + v37) + ""))) {
            case "foo bar":
                fmt.Println(v3, v3)
            case "\n\t":
                v38 := GenS2{F0: 937, F1: -9223372036854775808}
                _ = v38
            }
            _ = fi36
            _ = v37
        }
    }
    if ((v3 ^ (v3 % (((v3 * v3)) | 1))) < (-(v3))) {
        var v39 int8 = genFn0()
        _ = len(v2)
        _ = v39
    } else {
        var v40 int32 = (v3 + v3)
        _ = v40
    }
    v3 = v3
    var v41 string = "hello"
    for fi42 := 0; fi42 < 7; fi42++ {
        fmt.Println(fi42, fi42)
        _ = fi42
    }
    var v43 int = -2147483648
    v2[2] = v41
    v5 = append(v5, "hello")
    if ((v41 > ("\n\t" + v41)) && (!(32767 <= 1))) {
        v3 -= -1
        var v44 uint64 = 4294967295
        _ = v44
    } else {
        var v45 uint8 = 128
        for fi46 := 0; fi46 < 16; fi46++ {
            v47 := GenS3{F0: -32768, F1: "hello"}
            var v48 int8 = -48
            _ = fi46
            _ = v47
            _ = v48
        }
        _ = v45
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v43
    _ = v5
}
