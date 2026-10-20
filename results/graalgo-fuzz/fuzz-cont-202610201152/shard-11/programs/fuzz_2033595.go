package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 uint8; F2 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 int64; F2 int16 }

func main() {
    var v1 uint32 = 2147483648
    var v2 uint16 = 1
    var v3 int32 = 33
    var v4 int32 = (v3 % ((42) | 1))
    var v5 uint16 = v2
    fmt.Println(v3, v4, v5)
    var v6 uint8 = 255
    var v7 uint32 = v1
    var v8 float32 = (1.5 - 1.5)
    if (!true) {
        var v9 uint = 42
        switch (("a" + (("foo bar" + ("a" + "foo bar")) + "\n\t")) + "\n\t") {
        case "a":
            if true {
                for fi10 := 0; fi10 < 19; fi10++ {
                    v11 := []int{2147483647}
                    switch (("" + "hello") + "") {
                    case "\n\t":
                        var v12 uint = (^(v9))
                        _ = v12
                    case "hello":
                        v1 = (v7 ^ v7)
                    default:
                        var v13 int16 = 0
                        _ = v13
                    }
                    _ = fi10
                    _ = v11
                }
                v9 = v9
            }
        default:
            v14 := []int{42, int(v6), 1, -438}
            _ = v14
        }
        v9 += 42
        _ = v9
    }
    if (-2.25 != (0.1 + -2.25)) {
        var v15 int64 = -745
        switch 1 {
        case 2147483647:
            if false {
                var v16 uint8 = uint8(v7)
                switch ("a" + ("hello" + "a")) {
                case "hello":
                    switch "" {
                    case "hello":
                        v17 := []int{-1, -1, 2147483647}
                        _ = v17
                    case "foo bar":
                        switch 128 {
                        case 128:
                            fmt.Println(v3)
                        }
                    }
                case "foo bar":
                    fmt.Println(v7, v4, v4)
                }
                _ = v16
            } else {
                v18 := []int{128, -9223372036854775808, 797, 1}
                _ = v18
            }
        case 0:
            switch ("\n\t" + "hello") {
            case "":
                var v19 bool = (1 < 2)
                _ = v19
            case "foo bar":
                if false {
                    v8 += (v8 * -2.25)
                    for fi20 := 0; fi20 < 18; fi20++ {
                        if (335 != 2) {
                            for fi21 := 0; fi21 < 17; fi21++ {
                                v22 := []int{fi21, fi21, fi21, 9223372036854775807}
                                _ = fi21
                                _ = v22
                            }
                        } else {
                            switch "\n\t" {
                            case "a":
                                if (0 >= 433) {
                                    switch 42 {
                                    case -9223372036854775808:
                                        v23 := GenS1{F0: 2}
                                        _ = v23
                                    }
                                    for fi24 := 0; fi24 < 8; fi24++ {
                                        var v25 string = "\n\t"
                                        v26 := []int{9223372036854775807, (^(fi20))}
                                        var v27 uint64 = 0
                                        if len(v26) >= 2 {
                                            v26 = v26[1:len(v26)]
                                        }
                                        if true {
                                            fmt.Println(v27)
                                            v25 = v25
                                            v2 += (v2 % (((v5 | v2)) | 1))
                                        } else {
                                            v28 := GenS1{F0: 0}
                                            switch 462 {
                                            case -528:
                                                v29 := GenS2{F0: (v25 + ((v25 + v25) + v25)), F1: (v6 & 127), F2: v2}
                                                _ = v29
                                            case -806:
                                                v30 := GenS1{F0: 4294967295}
                                                _ = v30
                                            default:
                                                var v31 uint8 = v6
                                                _ = v31
                                            }
                                            var v32 float32 = v8
                                            _ = v28
                                            _ = v32
                                        }
                                        v33 := GenS1{F0: 0}
                                        _ = len(v26)
                                        var v34 int = fi20
                                        _ = fi24
                                        _ = v25
                                        _ = v26
                                        _ = v27
                                        _ = v33
                                        _ = v34
                                    }
                                    v35 := GenS2{F0: "", F1: (v6 & uint8(v7)), F2: 874}
                                    _ = v35
                                } else {
                                    switch (fi20 + (fi20 << ((2) & 63))) {
                                    case 2147483647:
                                        v36 := GenS2{F0: "foo bar", F1: (^(v6)), F2: (v2 >> ((42) & 63))}
                                        _ = v36
                                    }
                                }
                            case "":
                                v37 := GenS3{F0: ((((("" + ("foo bar" + "a")) + ("\n\t" + (("" + ("\n\t" + "foo bar")) + "hello"))) + ("hello" + (("foo bar" + ("\n\t" + "foo bar")) + "\n\t"))) + "hello") + "a")}
                                _ = v37
                            default:
                                v15 = v15
                            }
                            for fi38 := 0; fi38 < 7; fi38++ {
                                var v39 string = "\n\t"
                                _ = fi38
                                _ = v39
                            }
                        }
                        _ = fi20
                    }
                } else {
                    fmt.Println(v4, v5, v6)
                    v40 := []int{152, -735, -1, 1}
                    _ = v40
                }
            case "hello":
                v41 := GenS4{F0: v2, F1: 782, F2: 0}
                _ = v41
            }
        case 127:
            v42 := GenS1{F0: 4294967295}
            _ = v42
        }
        _ = v15
    }
    fmt.Println(v2, v8, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
