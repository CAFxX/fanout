package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 uint8; F1 uint32 }
func genFn0(p0 uint64, p1 int8) uint8 {
    p1 = p1
    p1 += p1
    p0 = p0
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return 127
}
func genFn1(p0 int16, p1 int32) int64 {
    v4 := []int{-1, 127}
    var v5 uint = 65535
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 9223372036854775807
}

func main() {
    var v1 uint16 = 213
    v2 := []int{-9223372036854775808}
    var v3 int64 = -9223372036854775808
    v6 := GenS1{F0: ("a" + "\n\t"), F1: 127}
    v3 -= (v3 ^ v3)
    var v7 int32 = -78
    v3 = v3
    switch ("a" + (("\n\t" + "foo bar") + "foo bar")) {
    case "":
        v8 := GenS1{F0: "foo bar", F1: 2}
        _ = v8
    default:
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 int8 = 100
            v3 = v3
            switch "" {
            case "":
                v11 := GenS1{F0: ("foo bar" + "a"), F1: 128}
                _ = v11
            default:
                if (!false) {
                    switch ("foo bar" + "\n\t") {
                    case "foo bar":
                        var v12 float64 = (1e10 * 1e10)
                        _ = v12
                    case "a":
                        var v13 int32 = (v7 << ((997) & 63))
                        _ = v13
                    default:
                        switch "\n\t" {
                        case "foo bar":
                            v3 += genFn1(0, (v7 << ((2) & 63)))
                        case "hello":
                            var v14 bool = true
                            _ = v14
                        case "":
                            for fi15 := 0; fi15 < 10; fi15++ {
                                if (817 == 18446744073709551615) {
                                    v10 = v10
                                    if false {
                                        var v16 int16 = 1000
                                        _ = v16
                                    }
                                } else {
                                    switch "foo bar" {
                                    case "\n\t":
                                        if (127 >= 127) {
                                            v10 -= 0
                                            v3 -= 0
                                        }
                                    }
                                }
                                v17 := GenS1{F0: "foo bar", F1: 42}
                                _ = fi15
                                _ = v17
                            }
                        default:
                            if true {
                                for fi18 := 0; fi18 < 10; fi18++ {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = fi18
                                }
                            } else {
                                v19 := GenS2{F0: (v7 | v7), F1: genFn0(0, 29)}
                                if len(v2) > 0 {
                                    v2[0] = (fi9 >> ((4294967295) & 63))
                                }
                                _ = v19
                            }
                        }
                    }
                    for fi20 := 0; fi20 < 19; fi20++ {
                        v21 := GenS2{F0: (v7 << ((4294967295) & 63)), F1: 127}
                        _ = fi20
                        _ = v21
                    }
                }
            }
            _ = fi9
            _ = v10
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if ((v1 - (^(v1))) == (v1 * 65535)) {
        for fi22 := 0; fi22 < 1; fi22++ {
            v23 := GenS3{F0: 15, F1: 42}
            _ = fi22
            _ = v23
        }
        for fi24 := 0; fi24 < 10; fi24++ {
            switch ((("" + ("hello" + "\n\t")) + "\n\t") + (((("\n\t" + "hello") + ("\n\t" + "\n\t")) + ("\n\t" + ("\n\t" + ("hello" + "\n\t")))) + ("\n\t" + ""))) {
            case "hello":
                var v25 bool = false
                _ = v25
            case "foo bar":
                _ = len(v2)
            }
            var v26 uint8 = 127
            switch "" {
            case "\n\t":
                v3 += v3
            default:
                v1 = v1
            }
            _ = fi24
            _ = v26
        }
        var v27 string = "foo bar"
        _ = v27
    }
    var v28 uint = 18446744073709551615
    if (127 < 1) {
        for fi29 := 0; fi29 < 18; fi29++ {
            if false {
                var v30 int16 = 0
                fmt.Println(v7)
                _ = v30
            }
            _ = fi29
        }
    } else {
        v28 -= (v28 * (^(v28)))
        var v31 int16 = 0
        v1 = v1
        _ = v31
    }
    v7 += 1
    var v32 uint32 = 2
    if true {
        var v33 uint = (v28 + 435)
        var v34 int8 = 39
        if (2 < (v33 << ((v28) & 63))) {
            var v35 int64 = v3
            v34 = (v34 & v34)
            _ = v35
        } else {
            var v36 int64 = v3
            v2 = append(v2, 128)
            _ = v36
        }
        _ = v33
        _ = v34
    }
    if (true && true) {
        var v37 int64 = (-(v3))
        _ = v37
    }
    fmt.Println(v7, v1)
    for fi38 := 0; fi38 < 2; fi38++ {
        for fi39 := 0; fi39 < 5; fi39++ {
            if (!true) {
                var v40 uint16 = v1
                _ = v40
            } else {
                v2 = append(v2, 42)
                var v41 float32 = 1e10
                v1 = 255
                _ = v41
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = len(v2)
            _ = len(v2)
            _ = fi39
        }
        switch (fi38 | fi38) {
        case 415:
            fmt.Println(v3)
        }
        _ = fi38
    }
    var v42 uint16 = (^(v1))
    fmt.Println(v3, v32)
    if len(v2) > 0 {
        v2[0] = 145
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v42:", v42)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v42
    _ = v6
    _ = v7
}
