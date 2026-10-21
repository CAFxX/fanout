package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 int8; F2 uint32 }

func main() {
    var v1 int = -260
    var v2 int8 = 127
    v3 := []string{"a", ("foo bar" + "\n\t"), ("" + ""), "\n\t"}
    v4 := make(map[int]string)
    v5 := []string{"", (("" + ((("\n\t" + "a") + "\n\t") + "a")) + "a"), "foo bar", ""}
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 int8 = -1
        _ = fi6
        _ = v7
    }
    switch "\n\t" {
    case "":
        v1 -= v1
    }
    if (2147483647 >= 1) {
        if (true || true) {
            v8 := GenS1{F0: ((!(true && true)) || (1.5 == ((0.0 * ((1e10 / ((-2.25 - 3.14159) - float64(0.1))) + (float64(((0.1 / float32((float64(3.14159) / 1e10))) + 3.14159)) * float64(3.14159)))) + (-2.25 * 1.5)))), F1: 42}
            _ = v8
        } else {
            v9 := GenS2{F0: v1}
            _ = v9
        }
        v1 = 2147483647
        var v10 int32 = -217
        _ = v10
    } else {
        switch (v1 % (((v1 | -611)) | 1)) {
        case -219:
            fmt.Println(v2)
        }
        v11 := GenS1{F0: (32767 > (v1 | (-(v1)))), F1: 18446744073709551615}
        var v12 int64 = 42
        _ = v11
        _ = v12
    }
    v1 = (-(v1))
    v13 := GenS2{F0: (v1 ^ (^(v1)))}
    v3 = append(v3, "")
    var v14 int8 = v2
    var v15 int = v1
    switch ((((("foo bar" + "a") + "hello") + "hello") + ("a" + "foo bar")) + "foo bar") {
    case "hello":
        for fi16 := 0; fi16 < 15; fi16++ {
            var v17 uint16 = 65535
            for fi18 := 0; fi18 < 3; fi18++ {
                switch v15 {
                case 9223372036854775807:
                    v17 += v17
                case 32767:
                    var v19 int = v1
                    _ = v19
                }
                _ = fi18
            }
            _ = fi16
            _ = v17
        }
    case "foo bar":
        v2 = 0
    case "a":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    v15 -= -2147483648
    v4[(v1 + v1)] = (("" + "") + "\n\t")
    v20 := GenS2{F0: 1}
    v2 += -60
    _ = len(v4)
    if true {
        v1 = 127
        if ((!(255 == 2)) || true) {
            fmt.Println(v2, v1, v1)
            v21 := GenS3{F0: 4294967295}
            _ = v21
        }
        for fi22 := 0; fi22 < 6; fi22++ {
            for fi23 := 0; fi23 < 6; fi23++ {
                var v24 int32 = 21
                _ = len(v3)
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
    } else {
        v25 := GenS4{F0: 2, F1: -128, F2: 65535}
        for fi26 := 0; fi26 < 12; fi26++ {
            var v27 uint32 = 42
            var v28 int64 = 276
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = v25
    }
    var v29 float64 = 0.0
    if false {
        v4[v1] = (("a" + (("foo bar" + ("" + ("a" + ""))) + ("hello" + (("" + ("a" + "a")) + "foo bar")))) + "")
        v30 := GenS2{F0: (^(v15))}
        _ = v30
    } else {
        if false {
            var v31 uint32 = 0
            v32 := GenS1{F0: false, F1: 0}
            _ = v31
            _ = v32
        }
        v2 += (v2 % (((-(v2))) | 1))
        var v33 string = ""
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
