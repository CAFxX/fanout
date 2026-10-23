package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 uint32; F2 int32 }
type GenS5 struct { F0 int64; F1 int }
func genFn0(p0 uint64) uint8 {
    var v5 float64 = 3.14159
    v6 := GenS1{F0: -558}
    _ = p0
    _ = v5
    _ = v6
    return 128
}
func genFn1() int64 {
    var v7 int8 = 1
    v8 := GenS2{F0: -226}
    var v9 uint8 = 0
    for fi10 := 0; fi10 < 3; fi10++ {
        for fi11 := 0; fi11 < 2; fi11++ {
            v9 = v9
            if (255 <= 617) {
                v7 += v7
                v12 := GenS2{F0: -9223372036854775808}
                switch ((("foo bar" + "") + "a") + ("a" + "")) {
                case "a":
                    fmt.Println(v7, fi10, v7)
                case "\n\t":
                    var v13 string = ("hello" + ("foo bar" + ("" + "foo bar")))
                    _ = v13
                }
                _ = v12
            } else {
                v14 := GenS2{F0: 581}
                _ = v14
            }
            _ = fi11
        }
        _ = fi10
    }
    switch -1 {
    case 9223372036854775807:
        if (!(float32(1e10) == ((3.14159 * 1.5) + float32((-2.25 / ((1.5 + -2.25) - 0.0)))))) {
            switch -277 {
            case 0:
                fmt.Println(v7, v7)
            case 127:
                switch "" {
                case "\n\t":
                    switch (("hello" + "foo bar") + "hello") {
                    case "":
                        for fi15 := 0; fi15 < 6; fi15++ {
                            v16 := GenS2{F0: 0}
                            _ = fi15
                            _ = v16
                        }
                    }
                case "foo bar":
                    for fi17 := 0; fi17 < 15; fi17++ {
                        var v18 string = "foo bar"
                        v9 = (v9 & (v9 * genFn0(4294967295)))
                        _ = fi17
                        _ = v18
                    }
                case "":
                    var v19 float64 = ((-2.25 / 1.0) * ((((-2.25 / -2.25) / 1.0) * 1.5) / -2.25))
                    _ = v19
                }
            case -1:
                if false {
                    var v20 int8 = (v7 / (((v7 - v7)) | 1))
                    _ = v20
                }
            }
        } else {
            v21 := []int{}
            _ = v21
        }
    case 128:
        for fi22 := 0; fi22 < 7; fi22++ {
            if (v9 != (v9 & v9)) {
                fmt.Println(fi22)
                v23 := []int{(fi22 % ((fi22) | 1))}
                _ = v23
            }
            _ = fi22
        }
    }
    _ = v7
    _ = v8
    _ = v9
    return -9223372036854775808
}

func main() {
    v1 := make(map[int]bool)
    v2 := []int{0, -1}
    var v3 string = "a"
    var v4 uint64 = 42
    fmt.Println(v4)
    fmt.Println(v4)
    v24 := GenS3{F0: genFn1()}
    v4 = v4
    v3 = "foo bar"
    v3 = ""
    var v25 uint16 = 65535
    var v26 int16 = -32768
    delete(v1, -869)
    fmt.Println(v25)
    v25 -= v25
    v1[789] = false
    _ = len(v2)
    v25 -= 2
    var v27 float64 = 0.0
    fmt.Println(v25, v3, v25)
    if false {
        if (-889 <= -169) {
            v28 := GenS4{F0: 1, F1: 42, F2: 1}
            _ = v28
        }
        if ((^(v4)) != (v4 | (v4 - (v4 / ((v4) | 1))))) {
            for fi29 := 0; fi29 < 13; fi29++ {
                v27 -= float64(1e10)
                v3 = (v3 + "")
                fmt.Println(fi29)
                _ = fi29
            }
        } else {
            var v30 uint32 = 565
            fmt.Println(v25, v25, v25)
            _ = v30
        }
        v31 := GenS4{F0: 1, F1: 1, F2: 2147483647}
        _ = v31
    }
    v32 := GenS5{F0: 285, F1: 976}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v4
}
