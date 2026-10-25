package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 uint32; F2 uint32 }
func genFn0(p0 int) int {
    for fi5 := 0; fi5 < 3; fi5++ {
        for fi6 := 0; fi6 < 3; fi6++ {
            switch (("" + (("\n\t" + ("a" + "hello")) + "hello")) + ("a" + "\n\t")) {
            case "foo bar":
                fmt.Println(fi6, p0)
            case "":
                fmt.Println(p0)
            case "\n\t":
                p0 += 1
            }
            _ = fi6
        }
        if ((p0 * p0) >= (p0 | p0)) {
            if true {
                var v7 uint8 = 2
                v8 := []int{fi5, fi5}
                _ = v7
                _ = v8
            } else {
                var v9 int16 = -1
                var v10 uint32 = 398
                fmt.Println(v9, v9, fi5)
                _ = v10
                _ = v9
            }
        } else {
            v11 := []int{p0, (fi5 ^ (p0 | 652)), -51}
            fmt.Println(p0, fi5)
            _ = v11
        }
        p0 = fi5
        _ = fi5
    }
    var v12 bool = (0.0 < 0.0)
    _ = p0
    _ = v12
    return -165
}

func main() {
    v1 := []int{128}
    var v2 int64 = 470
    var v3 int8 = -1
    v4 := []string{("\n\t" + ("\n\t" + "foo bar")), "", (("a" + "a") + ("hello" + "foo bar")), (("a" + ("\n\t" + "a")) + "hello")}
    v13 := GenS1{F0: "", F1: 2147483647}
    var v14 uint8 = 0
    v3 += v3
    var v15 int64 = 172
    v16 := GenS1{F0: ("\n\t" + ("hello" + "")), F1: -964}
    v15 -= 214
    if (uint32(v14) < 2147483647) {
        if (((!true) && true) || false) {
            if (923 >= 32767) {
                for fi17 := 0; fi17 < 3; fi17++ {
                    switch "hello" {
                    case "":
                        v18 := GenS1{F0: "hello", F1: 42}
                        _ = v18
                    case "foo bar":
                        var v19 float32 = 3.14159
                        _ = v19
                    case "\n\t":
                        var v20 float64 = 3.14159
                        _ = v20
                    }
                    v2 = 488
                    v2 = v2
                    _ = fi17
                }
                if false {
                    for fi21 := 0; fi21 < 9; fi21++ {
                        if true {
                            v2 = v2
                        }
                        v15 = int64(v3)
                        _ = fi21
                    }
                    v3 += (v3 - v3)
                    if ((v3 % (((v3 + -1)) | 1)) <= -121) {
                        fmt.Println(v14, v2)
                        if false {
                            var v22 uint64 = 0
                            v14 -= v14
                            v14 = (v14 * (v14 >> ((731) & 63)))
                            _ = v22
                        } else {
                            var v23 uint32 = 42
                            _ = v23
                        }
                    }
                }
            }
            v24 := GenS1{F0: "\n\t", F1: 522}
            v1 = append(v1, -1)
            _ = v24
        } else {
            if true {
                v2 += (v2 * (v15 % ((v2) | 1)))
                var v25 uint32 = 0
                if len(v1) > 0 {
                    v1[0] = genFn0(128)
                }
                _ = v25
            }
            v3 = v3
        }
        switch 32767 {
        case 9223372036854775807:
            if (!(!false)) {
                fmt.Println(v3, v15)
                v15 -= (^(v2))
            } else {
                var v26 uint = 310
                v27 := GenS1{F0: ("a" + "hello"), F1: 42}
                _ = v26
                _ = v27
            }
        case 211:
            var v28 string = ("foo bar" + (("" + "hello") + "foo bar"))
            _ = v28
        }
    }
    v3 = (v3 | v3)
    v29 := GenS2{F0: int16(v3)}
    v30 := GenS1{F0: "foo bar", F1: -2147483648}
    var v31 int32 = -2147483648
    v32 := GenS3{F0: -32768, F1: 4294967295, F2: 247}
    v2 = (v2 ^ v2)
    var v33 int32 = int32(v14)
    var v34 string = ("a" + "hello")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
