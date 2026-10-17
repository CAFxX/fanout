package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 int; F1 uint64; F2 uint16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint; F1 int64; F2 int8 }
func genFn0(p0 uint8) string {
    fmt.Println(p0)
    if true {
        for fi5 := 0; fi5 < 10; fi5++ {
            var v6 float32 = 0.1
            _ = fi5
            _ = v6
        }
    } else {
        v7 := GenS1{F0: 1}
        var v8 int8 = -1
        _ = v7
        _ = v8
    }
    fmt.Println(p0)
    switch -9223372036854775808 {
    case 978:
        switch (((("" + "") + "") + "a") + (("foo bar" + "foo bar") + (("hello" + ("a" + ("" + (("hello" + "foo bar") + "")))) + ("" + ("" + ("\n\t" + ("" + "\n\t"))))))) {
        case "a":
            switch "a" {
            case "a":
                p0 += (p0 & 155)
            case "hello":
                var v9 uint8 = 128
                _ = v9
            case "foo bar":
                v10 := GenS1{F0: 105}
                _ = v10
            }
        case "hello":
            v11 := GenS2{F0: 793, F1: ((("hello" + "") + "hello") + "a")}
            _ = v11
        }
    case 4294967295:
        if false {
            v12 := []int{619, -9223372036854775808, 32768, 32768}
            _ = len(v12)
            _ = v12
        } else {
            switch 0 {
            case 349:
                if (!false) {
                    for fi13 := 0; fi13 < 2; fi13++ {
                        if false {
                            fmt.Println(fi13)
                            p0 = 240
                            var v14 int32 = 2147483647
                            _ = v14
                        } else {
                            v15 := GenS3{F0: 1, F1: 42, F2: 65535}
                            v16 := GenS4{F0: 1}
                            var v17 string = "\n\t"
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                        _ = fi13
                    }
                    var v18 uint = 395
                    v19 := []int{-399, 32767, -439}
                    _ = v18
                    _ = v19
                }
            }
        }
    }
    p0 = p0
    _ = p0
    return "foo bar"
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = 0.1
    var v3 int64 = 304
    var v4 float64 = 1e10
    var v20 int = -9223372036854775808
    for fi21 := 0; fi21 < 19; fi21++ {
        var v22 uint64 = 4294967295
        _ = fi21
        _ = v22
    }
    if (-2147483648 >= 0) {
        fmt.Println(v4)
    }
    v23 := GenS4{F0: 686}
    for fi24 := 0; fi24 < 12; fi24++ {
        if ((3.14159 / v4) == v4) {
            v20 = -1
            var v25 uint64 = 4294967295
            _ = v25
        }
        if false {
            for fi26 := 0; fi26 < 8; fi26++ {
                if (255 > 255) {
                    if false {
                        if true {
                            v27 := v1["hello"]
                            v3 = v3
                            _ = v27
                        }
                        var v28 int = v20
                        v28 += (fi24 - (v28 % (((v28 << ((951) & 63))) | 1)))
                        _ = v28
                    } else {
                        fmt.Println(fi26, v3, v4)
                    }
                    v4 -= (v4 - v4)
                }
                if ((42 <= 2) || (895 != 1)) {
                    var v29 float64 = 1.5
                    _ = v29
                }
                delete(v1, "hello")
                _ = fi26
            }
            fmt.Println(v3, v20, fi24)
            v30 := []int{(v20 >> ((865) & 63)), v20, -471, 128}
            _ = v30
        } else {
            var v31 uint32 = 42
            v32 := GenS5{F0: 18446744073709551615, F1: (v3 & -106), F2: 95}
            _ = v31
            _ = v32
        }
        _ = fi24
    }
    v33 := []int{v20, v20, v20}
    if true {
        v2 = v2
    } else {
        v4 = float64(v2)
        _ = len(v1)
    }
    var v34 int32 = -1
    var v35 uint16 = 0
    switch v20 {
    case 4294967295:
        switch 1 {
        case 32767:
            var v36 string = ""
            _ = v36
        case 904:
            for fi37 := 0; fi37 < 3; fi37++ {
                v20 = (^(fi37))
                var v38 int64 = -713
                _ = fi37
                _ = v38
            }
        }
    case 42:
        v39 := GenS4{F0: 2}
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
