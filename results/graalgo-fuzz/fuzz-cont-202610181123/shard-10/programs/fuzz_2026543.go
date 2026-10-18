package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 uint64 }
func genFn0() bool {
    var v6 int32 = -2147483648
    v7 := []int{128, -1}
    if (127 <= -340) {
        switch ("a" + ("\n\t" + "foo bar")) {
        case "a":
            var v8 uint32 = 948
            _ = v8
        default:
            var v9 bool = ((v6 >> ((18446744073709551615) & 63)) < (v6 & (-(v6))))
            _ = v9
        }
    }
    _ = v6
    _ = v7
    return (false && true)
}
func genFn1(p0 int8, p1 uint8) uint {
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 int32 = 42
        switch (fi10 * fi10) {
        case -836:
            v11 -= v11
        }
        switch "\n\t" {
        case "":
            p1 += 157
        default:
            var v12 uint8 = p1
            _ = v12
        }
        _ = fi10
        _ = v11
    }
    var v13 uint64 = 0
    var v14 uint32 = 4294967295
    for fi15 := 0; fi15 < 2; fi15++ {
        var v16 int = fi15
        if (128 <= v16) {
            v16 -= (fi15 | (fi15 << ((42) & 63)))
        }
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return 80
}

func main() {
    var v1 string = "hello"
    v2 := []string{}
    var v3 int64 = 0
    var v4 uint64 = 0
    v5 := []string{}
    var v17 int64 = 9223372036854775807
    v18 := GenS1{F0: v1, F1: 65535, F2: 127}
    v19 := GenS1{F0: ("" + v1), F1: 65535, F2: 127}
    switch -246 {
    case 32768:
        if ((false && false) && true) {
            v3 = (^(v3))
        }
    }
    v1 = ("\n\t" + (("a" + v1) + ""))
    if true {
        v1 = (v1 + "foo bar")
        for fi20 := 0; fi20 < 4; fi20++ {
            v3 += (v17 << ((1) & 63))
            for fi21 := 0; fi21 < 10; fi21++ {
                if ((v1 + v1) >= "a") {
                    fmt.Println(v4)
                    if true {
                        v1 = v1
                        v22 := GenS2{F0: 83}
                        _ = v22
                    } else {
                        fmt.Println(v4, v1)
                    }
                } else {
                    if ((32767 <= 32767) || (42 <= 0)) {
                        v17 += v17
                        var v23 uint = 255
                        v17 += (-(v17))
                        _ = v23
                    } else {
                        _ = len(v5)
                        var v24 int32 = 0
                        _ = v24
                    }
                    switch v1 {
                    case "":
                        switch fi20 {
                        case 9223372036854775807:
                            if (!((23 >= 18446744073709551615) && (((v17 / ((v3) | 1)) > v17) || true))) {
                                v2 = append(v2, v1)
                                v25 := GenS1{F0: "", F1: 42, F2: 127}
                                var v26 bool = (((1 == 839) && false) && true)
                                _ = v25
                                _ = v26
                            } else {
                                if len(v2) > 0 {
                                    v2[0] = (v1 + (v1 + ("" + v1)))
                                }
                                v3 += (v17 & v17)
                                switch fi20 {
                                case 404:
                                    v27 := GenS1{F0: (((v1 + "\n\t") + (("\n\t" + v1) + "foo bar")) + v1), F1: 95, F2: -1}
                                    _ = v27
                                case 9223372036854775807:
                                    v4 = 2
                                case -9223372036854775808:
                                    switch (v1 + ((v1 + (v1 + ((v1 + "foo bar") + v1))) + v1)) {
                                    case "foo bar":
                                        fmt.Println(fi20, fi21, v1)
                                    default:
                                        var v28 int64 = v3
                                        _ = v28
                                    }
                                }
                            }
                        case -373:
                            v29 := GenS1{F0: "\n\t", F1: 4294967295, F2: -1}
                            _ = v29
                        case 1:
                            v3 -= (v3 ^ (v3 | 42))
                        }
                    case "\n\t":
                        switch "\n\t" {
                        case "hello":
                            v30 := GenS2{F0: 42}
                            v17 = int64(fi20)
                            var v31 bool = false
                            _ = v30
                            _ = v31
                        }
                    }
                    if (!(1e10 >= 1e10)) {
                        v32 := GenS3{F0: true, F1: v4}
                        switch (v1 + v1) {
                        case "\n\t":
                            v33 := GenS1{F0: v1, F1: 2147483648, F2: 56}
                            _ = v33
                        }
                        _ = v32
                    } else {
                        var v34 int64 = -445
                        v17 -= 42
                        v1 = ("hello" + v1)
                        _ = v34
                    }
                }
                var v35 uint32 = 1
                _ = fi21
                _ = v35
            }
            _ = fi20
        }
    }
    v5 = append(v5, (((v1 + (v1 + v1)) + ("a" + v1)) + v1))
    v17 -= v17
    var v36 int32 = -2147483648
    v3 = -9223372036854775808
    fmt.Println(v4, v3, v36)
    if (42 >= 32767) {
        if len(v5) > 0 {
            v5[0] = (v1 + v1)
        }
        v3 = 0
        var v37 int32 = v36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
