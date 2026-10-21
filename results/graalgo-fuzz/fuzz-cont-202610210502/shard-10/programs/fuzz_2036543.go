package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int32 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int8; F1 string; F2 int16 }
func genFn0(p0 bool) string {
    var v5 int32 = -165
    if p0 {
        fmt.Println(p0, v5)
        if (false || p0) {
            v5 = v5
            for fi6 := 0; fi6 < 3; fi6++ {
                var v7 bool = p0
                var v8 uint8 = 0
                v9 := GenS1{F0: 0, F1: 888, F2: v5}
                _ = fi6
                _ = v7
                _ = v8
                _ = v9
            }
            v10 := []int{}
            _ = v10
        }
    } else {
        if true {
            p0 = ((!p0) || p0)
            var v11 string = ""
            _ = v11
        }
        p0 = ((!((p0 && (!(v5 < v5))) || false)) || false)
    }
    v12 := GenS1{F0: 1, F1: 339, F2: (v5 - v5)}
    _ = p0
    _ = v12
    _ = v5
    return ("foo bar" + "\n\t")
}

func main() {
    v1 := []string{(("a" + ("" + "hello")) + ((("" + "foo bar") + "foo bar") + ("hello" + (("" + ("a" + (("\n\t" + "\n\t") + ""))) + "foo bar")))), "hello", ("hello" + ("hello" + ("hello" + "\n\t"))), "\n\t"}
    var v2 int32 = 1
    var v3 uint16 = 42
    var v4 int8 = 42
    var v13 int16 = -404
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        if ((3.14159 >= 3.14159) || (!(!(42 >= -1)))) {
            var v15 bool = false
            v1 = append(v1, "a")
            _ = v15
        }
        for fi16 := 0; fi16 < 10; fi16++ {
            var v17 float64 = 1.5
            if (2 != uint8(v3)) {
                var v18 uint32 = 1
                var v19 uint64 = 1
                switch "foo bar" {
                case "\n\t":
                    fmt.Println(v18, fi16, v18)
                case "a":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "":
                    var v20 string = "a"
                    _ = v20
                }
                v1 = append(v1, "\n\t")
                _ = len(v1)
                _ = v18
                _ = v19
            }
            switch (("foo bar" + ("hello" + (("foo bar" + "foo bar") + "\n\t"))) + "\n\t") {
            case "a":
                v1 = append(v1, ("foo bar" + "\n\t"))
            case "foo bar":
                var v21 bool = ("hello" < (("hello" + "") + ("a" + "")))
                if (18446744073709551615 == 4294967295) {
                    if (v21 && (974 >= 2)) {
                        var v22 bool = v21
                        var v23 uint8 = 1
                        if v22 {
                            var v24 uint = 4294967295
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v25 := GenS2{F0: (v3 ^ 77), F1: (v24 << ((v24) & 63))}
                            _ = v24
                            _ = v25
                        }
                        _ = v22
                        _ = v23
                    } else {
                        var v26 int = (fi16 / ((42) | 1))
                        _ = v26
                    }
                    v27 := GenS2{F0: v3, F1: 1}
                    v21 = v21
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v27
                }
                if false {
                    v2 = 2147483647
                    var v28 uint8 = 42
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v28
                }
                _ = v21
            case "":
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
            }
            _ = fi16
            _ = v17
        }
        var v29 uint = 65535
        _ = fi14
        _ = v29
    }
    if (383 < 2147483648) {
        if (599 >= 18446744073709551615) {
            v4 = (v4 / (((v4 | (v4 ^ (v4 % ((int8(v3)) | 1))))) | 1))
            v30 := GenS3{F0: (-(v4)), F1: "a", F2: (v13 & (v13 | v13))}
            _ = v30
        }
    } else {
        v31 := GenS3{F0: (v4 << ((0) & 63)), F1: "\n\t", F2: (-(v13))}
        v32 := GenS3{F0: v4, F1: (("\n\t" + "") + "foo bar"), F2: (v13 | 954)}
        _ = v31
        _ = v32
    }
    var v33 int = -727
    v1 = append(v1, "hello")
    for fi34 := 0; fi34 < 17; fi34++ {
        v2 += -593
        var v35 uint8 = 131
        fmt.Println(v4, fi34, v2)
        _ = fi34
        _ = v35
    }
    for fi36 := 0; fi36 < 13; fi36++ {
        v4 += 0
        v37 := GenS1{F0: -9223372036854775808, F1: 463, F2: 42}
        _ = fi36
        _ = v37
    }
    var v38 int16 = (v13 - v13)
    var v39 uint8 = 128
    var v40 float32 = 1e10
    var v41 uint = uint(v39)
    var v42 int8 = (v4 % (((v4 * 42)) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v33
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
}
