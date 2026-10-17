package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 int8; F2 bool }
func genFn0(p0 int32, p1 int, p2 int32) string {
    var v5 int8 = -128
    p0 += p2
    if (994 == -1) {
        v6 := []int{(p1 / (((p1 % ((p1) | 1))) | 1)), -382}
        _ = v6
    }
    var v7 string = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v7
    return v7
}
func genFn1(p0 uint32) int32 {
    var v8 int64 = 1
    v9 := GenS1{F0: 1, F1: (("" + "\n\t") + "foo bar")}
    if false {
        switch "foo bar" {
        case "":
            v10 := []int{2147483647, -9223372036854775808}
            _ = v10
        case "a":
            if true {
                fmt.Println(p0, v8)
                v11 := GenS1{F0: -985, F1: ((("hello" + ("\n\t" + "foo bar")) + "a") + "hello")}
                var v12 uint64 = 4294967295
                _ = v11
                _ = v12
            }
        case "\n\t":
            v13 := GenS2{F0: 970}
            _ = v13
        default:
            for fi14 := 0; fi14 < 16; fi14++ {
                if (0 < 1) {
                    var v15 int8 = 0
                    _ = v15
                }
                _ = fi14
            }
        }
        for fi16 := 0; fi16 < 1; fi16++ {
            for fi17 := 0; fi17 < 11; fi17++ {
                v18 := GenS1{F0: 0, F1: ("foo bar" + "hello")}
                v19 := []int{(fi17 % (((fi17 * 32768)) | 1)), 127, (fi17 - (fi17 * fi17))}
                p0 -= (p0 * 4294967295)
                _ = fi17
                _ = v18
                _ = v19
            }
            var v20 float32 = (0.0 * (0.0 - 0.1))
            _ = fi16
            _ = v20
        }
        p0 += p0
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        for fi22 := 0; fi22 < 17; fi22++ {
            if (-1 <= -522) {
                switch "" {
                case "foo bar":
                    if (!(false && true)) {
                        v8 -= v8
                    }
                case "hello":
                    fmt.Println(fi22, p0, p0)
                case "":
                    if (false || (-1 > 0)) {
                        v23 := GenS2{F0: 867}
                        _ = v23
                    }
                }
                var v24 bool = (!((!false) || false))
                _ = v24
            } else {
                v25 := GenS1{F0: 42, F1: "a"}
                v26 := []int{(-(fi21)), 42}
                v8 += 401
                _ = v25
                _ = v26
            }
            var v27 int8 = -1
            v27 = v27
            _ = fi22
            _ = v27
        }
        p0 = p0
        var v28 uint8 = 122
        _ = fi21
        _ = v28
    }
    var v29 bool = true
    _ = p0
    _ = v29
    _ = v8
    _ = v9
    return -2147483648
}

func main() {
    var v1 float64 = 0.1
    v2 := []string{"foo bar", "foo bar", ("hello" + ((("\n\t" + "\n\t") + "a") + (("\n\t" + "a") + "a")))}
    var v3 bool = (4294967295 < 255)
    var v4 int = 356
    var v30 string = ""
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi31 := 0; fi31 < 1; fi31++ {
        var v32 int16 = 42
        _ = fi31
        _ = v32
    }
    v4 -= v4
    var v33 uint16 = 118
    fmt.Println(v4)
    var v34 uint16 = (v33 - v33)
    _ = len(v2)
    if (!("hello" <= v30)) {
        v34 = (^(v33))
        v35 := GenS1{F0: genFn1(2147483648), F1: v30}
        switch v30 {
        case "hello":
            for fi36 := 0; fi36 < 16; fi36++ {
                for fi37 := 0; fi37 < 17; fi37++ {
                    v33 += (^(v34))
                    _ = fi37
                }
                var v38 int32 = -218
                _ = fi36
                _ = v38
            }
        case "":
            if len(v2) > 0 {
                v2[0] = ""
            }
        default:
            _ = len(v2)
        }
        _ = v35
    } else {
        v39 := GenS3{F0: v33, F1: 127, F2: v3}
        var v40 float64 = ((v1 / v1) + v1)
        var v41 string = v30
        _ = v39
        _ = v40
        _ = v41
    }
    v33 = (v33 ^ v34)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v4
}
