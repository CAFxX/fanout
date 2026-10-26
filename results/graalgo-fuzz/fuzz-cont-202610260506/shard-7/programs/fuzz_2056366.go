package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 int32; F2 string }
type GenS4 struct { F0 string }

func main() {
    var v1 int32 = -1
    var v2 uint32 = 42
    v3 := make(map[uint32]int)
    if false {
        for fi4 := 0; fi4 < 19; fi4++ {
            for fi5 := 0; fi5 < 11; fi5++ {
                var v6 float64 = 1e10
                _ = fi5
                _ = v6
            }
            for fi7 := 0; fi7 < 18; fi7++ {
                _ = len(v3)
                _ = fi7
            }
            _ = fi4
        }
        var v8 int64 = -9223372036854775808
        v1 -= (-(v1))
        _ = v8
    }
    v9 := GenS1{F0: true}
    fmt.Println(v2, v2)
    if (!true) {
        _ = len(v3)
        v2 += v2
        v2 += (^(v2))
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        if false {
            var v11 float32 = (float32(-2.25) - (((3.14159 + -2.25) - (0.1 * 0.0)) * -2.25))
            _ = v11
        } else {
            switch (("foo bar" + ("" + ("a" + "hello"))) + ((("\n\t" + ("foo bar" + "foo bar")) + ("hello" + "foo bar")) + "a")) {
            case "":
                fmt.Println(v2, v1)
            default:
                if false {
                    v12 := GenS1{F0: true}
                    switch ("hello" + "") {
                    case "":
                        var v13 float32 = 1e10
                        _ = v13
                    case "foo bar":
                        var v14 int32 = 0
                        _ = v14
                    case "\n\t":
                        v1 = v1
                    }
                    _ = v12
                }
            }
            for fi15 := 0; fi15 < 1; fi15++ {
                if true {
                    v16 := []int{(fi10 + (fi10 / ((-9223372036854775808) | 1))), (fi15 | fi10)}
                    _ = v16
                }
                fmt.Println(v1, fi10, v1)
                _ = fi15
            }
            switch "\n\t" {
            case "a":
                var v17 uint64 = 253
                _ = v17
            case "foo bar":
                if (false && true) {
                    var v18 uint8 = 0
                    _ = v18
                } else {
                    v19 := GenS2{F0: 42}
                    for fi20 := 0; fi20 < 8; fi20++ {
                        v2 = v2
                        v2 += v2
                        v21 := v3[(v2 | 2)]
                        switch (-(fi20)) {
                        case 2147483647:
                            var v22 string = "foo bar"
                            delete(v3, v2)
                            _ = v22
                        case 32767:
                            switch (fi10 * 2147483647) {
                            case -95:
                                v2 += (v2 ^ 2147483647)
                            case 42:
                                var v23 float64 = 1.5
                                _ = v23
                            }
                        case -1:
                            var v24 int = fi10
                            _ = v24
                        }
                        v25 := GenS1{F0: (false || false)}
                        _ = fi20
                        _ = v21
                        _ = v25
                    }
                    fmt.Println(fi10, v2, v2)
                    _ = v19
                }
            default:
                if (42 <= -550) {
                    switch "\n\t" {
                    case "hello":
                        for fi26 := 0; fi26 < 20; fi26++ {
                            if false {
                                var v27 int16 = -32768
                                v27 = int16(v1)
                                var v28 uint16 = 860
                                _ = v27
                                _ = v28
                            }
                            _ = fi26
                        }
                    default:
                        v1 = 42
                    }
                    switch "a" {
                    case "a":
                        for fi29 := 0; fi29 < 17; fi29++ {
                            v1 -= 881
                            var v30 uint = 965
                            v30 -= (v30 + 943)
                            v31 := GenS3{F0: true, F1: v1, F2: "hello"}
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                    }
                    v1 = v1
                }
            }
        }
        _ = fi10
    }
    v32 := []int{42, 127, 32767, 1}
    for fi33 := 0; fi33 < 5; fi33++ {
        v1 = (v1 - 115)
        for fi34 := 0; fi34 < 13; fi34++ {
            var v35 uint16 = 1
            v35 = v35
            _ = fi34
            _ = v35
        }
        _ = fi33
    }
    if false {
        if false {
            var v36 int = 0
            var v37 int64 = -9223372036854775808
            switch "a" {
            case "\n\t":
                fmt.Println(v36)
            }
            _ = v36
            _ = v37
        } else {
            v38 := v3[(v2 * (^(v2)))]
            _ = v38
        }
    } else {
        v39 := GenS4{F0: "foo bar"}
        v1 -= (v1 + v1)
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v32:", len(v32), v32)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v9
}
