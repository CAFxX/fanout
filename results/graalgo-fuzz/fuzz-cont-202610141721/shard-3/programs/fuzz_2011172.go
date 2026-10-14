package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 uint; F1 uint8 }
func genFn0(p0 bool, p1 uint32, p2 uint16) int8 {
    v5 := []int{127, 9223372036854775807}
    p0 = p0
    if false {
        var v6 int64 = 9223372036854775807
        _ = v6
    }
    p0 = ((0.1 < ((-2.25 - float64(0.1)) - 1e10)) && false)
    p2 = (^(p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -128
}
func genFn1() string {
    for fi7 := 0; fi7 < 16; fi7++ {
        if false {
            if (-128 != 1) {
                var v8 int32 = -2147483648
                _ = v8
            }
        } else {
            switch "" {
            case "foo bar":
                var v9 int8 = 100
                _ = v9
            case "a":
                switch fi7 {
                case 882:
                    v10 := GenS1{F0: -1, F1: 644}
                    _ = v10
                case -289:
                    var v11 bool = (!(!false))
                    _ = v11
                case 32768:
                    for fi12 := 0; fi12 < 4; fi12++ {
                        if true {
                            var v13 int8 = 1
                            var v14 int8 = v13
                            _ = v13
                            _ = v14
                        }
                        _ = fi12
                    }
                }
            case "hello":
                if false {
                    var v15 uint64 = 18446744073709551615
                    v16 := []int{2147483647}
                    _ = v15
                    _ = v16
                }
            }
            var v17 int64 = -533
            _ = v17
        }
        _ = fi7
    }
    var v18 int = 4294967295
    var v19 uint16 = 255
    _ = v18
    _ = v19
    return "a"
}

func main() {
    var v1 float64 = float64(-2.25)
    var v2 string = "\n\t"
    var v3 bool = false
    var v4 int = 2147483647
    v20 := []int{(v4 >> ((0) & 63))}
    var v21 uint = 785
    var v22 uint16 = 2
    v3 = (1e10 <= float64(float32(1.5)))
    for fi23 := 0; fi23 < 15; fi23++ {
        switch (v4 % (((v4 * 0)) | 1)) {
        case -1:
            for fi24 := 0; fi24 < 19; fi24++ {
                if (-114 == 42) {
                    if false {
                        if ((v21 * (v21 % ((uint(v22)) | 1))) == v21) {
                            v22 = (v22 & 698)
                        } else {
                            v20 = append(v20, fi23)
                        }
                    }
                    var v25 int16 = -32768
                    _ = v25
                } else {
                    switch (v4 & (fi24 & v4)) {
                    case 407:
                        v1 = v1
                        _ = len(v20)
                        if (498 > 18446744073709551615) {
                            switch "\n\t" {
                            case "":
                                if (((0.1 / 1e10) + (0.0 - 3.14159)) <= 0.1) {
                                    switch ("hello" + "hello") {
                                    case "a":
                                        v22 = v22
                                    }
                                }
                            case "foo bar":
                                fmt.Println(v2)
                            case "a":
                                if (4294967295 < 65535) {
                                    if true {
                                        v3 = false
                                    }
                                }
                            default:
                                v22 -= v22
                            }
                        } else {
                            var v26 uint8 = 255
                            _ = v26
                        }
                        v3 = (v3 || v3)
                    case -779:
                        switch fi24 {
                        case -745:
                            var v27 uint32 = 2
                            _ = v27
                        }
                    case 32767:
                        v20 = append(v20, 42)
                    }
                }
                var v28 uint = (^(v21))
                _ = fi24
                _ = v28
            }
        case 850:
            v2 = v2
        case -2147483648:
            fmt.Println(v3)
        }
        if (true && (v3 || true)) {
            v20 = append(v20, (fi23 % (((^(fi23))) | 1)))
            v2 = (v2 + "")
            fmt.Println(v3)
        } else {
            v29 := GenS2{F0: 100}
            _ = v29
        }
        fmt.Println(v4, v3)
        _ = fi23
    }
    v1 += (1e10 - float64(float32((v1 * (v1 * v1)))))
    switch (v4 | 0) {
    case 128:
        fmt.Println(v2, v3, v21)
    case 4294967295:
        var v30 float32 = ((1.5 * float32(1.5)) / 1.5)
        _ = v30
    case 32768:
        for fi31 := 0; fi31 < 3; fi31++ {
            if v3 {
                for fi32 := 0; fi32 < 8; fi32++ {
                    var v33 int32 = -1
                    _ = fi32
                    _ = v33
                }
            }
            _ = fi31
        }
    }
    v34 := GenS3{F0: (!(!(!(v2 >= "a")))), F1: 42}
    v35 := GenS2{F0: 0}
    var v36 uint16 = v22
    var v37 string = (("" + (v2 + "")) + v2)
    for fi38 := 0; fi38 < 10; fi38++ {
        v2 = v37
        _ = fi38
    }
    fmt.Println(v36)
    v39 := GenS3{F0: (true || v3), F1: 341}
    v40 := GenS4{F0: uint(v22), F1: 127}
    var v41 string = ("hello" + v37)
    v22 = 710
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
