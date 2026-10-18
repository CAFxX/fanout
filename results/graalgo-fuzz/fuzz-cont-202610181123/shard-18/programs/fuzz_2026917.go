package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool; F1 bool }
type GenS6 struct { F0 int }
type GenS7 struct { F0 string; F1 int32 }

func main() {
    var v1 float64 = float64(0.0)
    v2 := make(map[string]bool)
    var v3 int16 = 493
    v3 = (v3 ^ 75)
    v4 := v2["foo bar"]
    for fi5 := 0; fi5 < 17; fi5++ {
        if (((v3 | v3) >= -32768) || (!true)) {
            v6 := GenS1{F0: -2147483648}
            _ = v6
        } else {
            for fi7 := 0; fi7 < 17; fi7++ {
                delete(v2, "foo bar")
                _ = fi7
            }
            if true {
                var v8 bool = v4
                for fi9 := 0; fi9 < 9; fi9++ {
                    v4 = (2147483647 <= -458)
                    v8 = (3.14159 < 0.0)
                    _ = fi9
                }
                _ = v8
            }
        }
        var v10 int8 = 1
        _ = fi5
        _ = v10
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        fmt.Println(v4, fi11)
        fmt.Println(fi11)
        _ = fi11
    }
    fmt.Println(v4, v4, v1)
    for fi12 := 0; fi12 < 6; fi12++ {
        if (fi12 == (-(fi12))) {
            _ = len(v2)
            for fi13 := 0; fi13 < 10; fi13++ {
                var v14 float64 = v1
                v15 := []int{fi13, fi12, (fi12 ^ fi13), fi13}
                v16 := GenS2{F0: 11}
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
            if (!(1.5 >= float32(-2.25))) {
                v4 = (18446744073709551615 == 0)
            } else {
                switch fi12 {
                case 4294967295:
                    for fi17 := 0; fi17 < 9; fi17++ {
                        if v4 {
                            var v18 string = "hello"
                            var v19 uint32 = 453
                            _ = v18
                            _ = v19
                        } else {
                            v20 := GenS3{F0: -6}
                            _ = v20
                        }
                        v2[(((("foo bar" + ("\n\t" + "a")) + ("foo bar" + ("hello" + (("hello" + "\n\t") + "")))) + "") + "foo bar")] = v4
                        switch "\n\t" {
                        case "":
                            var v21 uint64 = 18446744073709551615
                            _ = v21
                        case "hello":
                            if (v4 || v4) {
                                v3 = v3
                                fmt.Println(v4, v3)
                            }
                        case "foo bar":
                            fmt.Println(fi12, v4, v1)
                        }
                        _ = fi17
                    }
                case 2147483647:
                    v22 := []int{fi12, fi12, (fi12 % ((fi12) | 1))}
                    _ = v22
                }
            }
        } else {
            v1 = (float64(0.1) + v1)
        }
        if (fi12 == 353) {
            if (!v4) {
                for fi23 := 0; fi23 < 11; fi23++ {
                    var v24 int32 = -727
                    _ = fi23
                    _ = v24
                }
            }
        } else {
            var v25 int = (^(fi12))
            v26 := GenS3{F0: 42}
            v27 := GenS2{F0: -1}
            _ = v25
            _ = v26
            _ = v27
        }
        _ = fi12
    }
    fmt.Println(v4, v1, v3)
    v28 := GenS4{F0: -9223372036854775808}
    var v29 int32 = 2147483647
    var v30 bool = v4
    v31 := GenS5{F0: ((v30 || (false && (true && (0.1 <= 1.5)))) && true), F1: (512 < -9223372036854775808)}
    if (!false) {
        for fi32 := 0; fi32 < 14; fi32++ {
            v3 = (v3 | (v3 % (((v3 - v3)) | 1)))
            if ((v3 ^ 730) > v3) {
                v4 = v4
                switch "foo bar" {
                case "":
                    var v33 int8 = -40
                    _ = v33
                case "\n\t":
                    switch 32768 {
                    case 0:
                        fmt.Println(v4)
                    case 42:
                        var v34 string = "hello"
                        _ = v34
                    }
                default:
                    if (1 <= 1) {
                        var v35 uint32 = 42
                        var v36 float32 = 3.14159
                        _ = v35
                        _ = v36
                    }
                }
            } else {
                v37 := GenS6{F0: (fi32 & (fi32 & (fi32 & fi32)))}
                v30 = ((-(v3)) == v3)
                _ = v37
            }
            _ = fi32
        }
        switch 1 {
        case 32768:
            for fi38 := 0; fi38 < 11; fi38++ {
                if false {
                    if v30 {
                        v39 := GenS4{F0: fi38}
                        v40 := []int{fi38, (fi38 & (fi38 + (fi38 << ((4294967295) & 63)))), fi38}
                        _ = len(v40)
                        _ = v39
                        _ = v40
                    } else {
                        var v41 int8 = 33
                        v42 := []int{0, (fi38 + 42), 127, fi38}
                        _ = v41
                        _ = v42
                    }
                    var v43 int = 1
                    _ = v43
                } else {
                    v44 := GenS4{F0: (fi38 ^ fi38)}
                    _ = v44
                }
                v45 := GenS7{F0: (("a" + ("a" + ("hello" + "\n\t"))) + ""), F1: v29}
                if (!(v3 == (v3 & -32768))) {
                    var v46 uint16 = 1
                    _ = v46
                } else {
                    if v4 {
                        var v47 uint16 = 65535
                        v1 = (v1 + 1.5)
                        _ = v47
                    }
                    v3 = v3
                    var v48 bool = v30
                    _ = v48
                }
                _ = fi38
                _ = v45
            }
        case -2147483648:
            var v49 uint64 = 18446744073709551615
            _ = v49
        case 1:
            var v50 int16 = (v3 + 32767)
            _ = v50
        }
        var v51 uint64 = 4294967295
        _ = v51
    }
    v4 = v4
    switch 366 {
    case 984:
        var v52 int = 0
        _ = v52
    case 128:
        for fi53 := 0; fi53 < 5; fi53++ {
            if (v30 || v30) {
                var v54 string = (("hello" + "hello") + ("a" + ("foo bar" + (("foo bar" + "") + ""))))
                _ = v54
            } else {
                v55 := GenS1{F0: v29}
                _ = v55
            }
            var v56 uint = 0
            var v57 int8 = 127
            _ = fi53
            _ = v56
            _ = v57
        }
    case 1:
        if v30 {
            for fi58 := 0; fi58 < 10; fi58++ {
                switch "hello" {
                case "foo bar":
                    switch (fi58 % ((127) | 1)) {
                    case 9223372036854775807:
                        v4 = ((fi58 & fi58) < (fi58 | fi58))
                    default:
                        v59 := GenS2{F0: -1}
                        _ = v59
                    }
                }
                var v60 int8 = 124
                _ = fi58
                _ = v60
            }
        } else {
            var v61 string = ((("\n\t" + "foo bar") + "a") + (("\n\t" + "\n\t") + ("a" + "\n\t")))
            v62 := GenS5{F0: (!true), F1: (65535 <= 853)}
            var v63 int = 128
            _ = v61
            _ = v62
            _ = v63
        }
    default:
        _ = len(v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
