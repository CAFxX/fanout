package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 uint64; F2 uint }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 int32; F1 uint8; F2 string }
func genFn0(p0 string, p1 uint64, p2 uint) int {
    v7 := []int{}
    switch -756 {
    case 183:
        if (false || (!(int16(p1) > -289))) {
            if true {
                var v8 uint16 = 0
                p1 = p1
                _ = v8
            } else {
                var v9 string = "a"
                if true {
                    fmt.Println(p2)
                    for fi10 := 0; fi10 < 7; fi10++ {
                        if true {
                            v9 = (v9 + v9)
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                        }
                        _ = fi10
                    }
                    v11 := GenS1{F0: -1, F1: 955}
                    _ = v11
                }
                _ = v9
            }
        }
    }
    v7 = append(v7, 1)
    var v12 uint = (p2 * (^(p2)))
    v13 := GenS2{F0: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v7
    return -9223372036854775808
}
func genFn1(p0 uint, p1 bool) uint32 {
    var v14 int8 = 0
    var v15 float32 = 3.14159
    fmt.Println(v15, v14, v15)
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return 367
}

func main() {
    var v1 uint16 = 42
    var v2 int8 = 100
    var v3 float32 = 3.14159
    var v4 int = 32767
    var v5 uint64 = 0
    var v6 float64 = ((-2.25 / 1e10) / -2.25)
    for fi16 := 0; fi16 < 20; fi16++ {
        var v17 int16 = -32768
        switch (^(v4)) {
        case 32767:
            var v18 uint64 = (v5 * (v5 << ((59) & 63)))
            _ = v18
        }
        v3 = v3
        _ = fi16
        _ = v17
    }
    fmt.Println(v2)
    v19 := []int{(v4 << ((628) & 63)), (v4 ^ 9223372036854775807)}
    var v20 float32 = float32(v6)
    switch "hello" {
    case "":
        var v21 bool = false
        _ = v21
    case "\n\t":
        v22 := GenS3{F0: v5, F1: (^(v5)), F2: 0}
        _ = v22
    }
    var v23 float64 = 1.5
    for fi24 := 0; fi24 < 8; fi24++ {
        switch "a" {
        case "hello":
            fmt.Println(v20, v20)
        case "foo bar":
            if false {
                switch -2147483648 {
                case 1:
                    var v25 uint32 = 2147483648
                    _ = v25
                case 9223372036854775807:
                    for fi26 := 0; fi26 < 12; fi26++ {
                        if true {
                            if ((v20 / (v20 - float32((-2.25 + (v6 + ((3.14159 * 1.5) + v23)))))) != (v20 + (((v20 * v3) + -2.25) * float32(1.5)))) {
                                var v27 int32 = 1
                                v4 -= fi26
                                v27 = v27
                                _ = v27
                            } else {
                                if ((!(127 <= v2)) || true) {
                                    v1 += (v1 / ((v1) | 1))
                                } else {
                                    var v28 uint32 = 0
                                    var v29 int = genFn0("foo bar", uint64(v1), 363)
                                    v5 += v5
                                    v6 += v6
                                    _ = v28
                                    _ = v29
                                }
                                fmt.Println(fi26, v4)
                                switch (-(v4)) {
                                case -2147483648:
                                    switch ("" + "") {
                                    case "foo bar":
                                        if len(v19) >= 2 {
                                            v19 = v19[1:len(v19)]
                                        }
                                    }
                                case 269:
                                    v5 += (v5 % ((v5) | 1))
                                default:
                                    fmt.Println(v2, v2, fi26)
                                }
                                var v30 string = "hello"
                                _ = v30
                            }
                        } else {
                            if (genFn0("foo bar", (v5 >> ((65535) & 63)), 42) < genFn0("hello", (^(v5)), 18446744073709551615)) {
                                var v31 int8 = v2
                                _ = v31
                            }
                        }
                        if len(v19) > 0 {
                            v19[0] = (fi26 - v4)
                        }
                        _ = fi26
                    }
                case -1:
                    if true {
                        if len(v19) > 0 {
                            v19[0] = v4
                        }
                    }
                default:
                    for fi32 := 0; fi32 < 6; fi32++ {
                        fmt.Println(v6, v20)
                        var v33 bool = true
                        _ = fi32
                        _ = v33
                    }
                }
                v23 = (v23 * v6)
            } else {
                var v34 float64 = v23
                var v35 int8 = v2
                if (-32768 == 32767) {
                    var v36 string = ""
                    var v37 int16 = 822
                    _ = v36
                    _ = v37
                } else {
                    for fi38 := 0; fi38 < 5; fi38++ {
                        fmt.Println(fi38)
                        if len(v19) > 0 {
                            v19[0] = -2147483648
                        }
                        _ = fi38
                    }
                }
                _ = v34
                _ = v35
            }
        case "\n\t":
            if false {
                _ = len(v19)
                v23 += float64((v20 - v20))
            }
        default:
            v1 += 255
        }
        _ = fi24
    }
    v39 := GenS4{F0: (v2 >> ((65535) & 63)), F1: 0}
    if len(v19) > 0 {
        v19[0] = (v4 + (v4 / (((v4 + v4)) | 1)))
    }
    var v40 uint16 = v1
    v40 += 482
    switch ("hello" + "\n\t") {
    case "hello":
        v4 = (v4 >> ((42) & 63))
    }
    v20 -= 3.14159
    v41 := GenS5{F0: -198, F1: 2, F2: "\n\t"}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
