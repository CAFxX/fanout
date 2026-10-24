package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string; F1 int; F2 int32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int; F1 uint64 }
type GenS7 struct { F0 uint; F1 string; F2 uint }
func genFn0(p0 string, p1 uint16, p2 bool) int {
    v6 := []int{32767, 127, -9223372036854775808}
    v7 := GenS1{F0: 127}
    fmt.Println(p0)
    fmt.Println(p2)
    var v8 int16 = 986
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 int16, p1 int32, p2 uint16) int {
    for fi9 := 0; fi9 < 10; fi9++ {
        p1 += (p1 + (p1 + -1))
        if (!(uint8(p2) != 1)) {
            p2 += p2
        }
        _ = fi9
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        v11 := GenS1{F0: -128}
        p0 = (p0 / (((-(p0))) | 1))
        var v12 bool = false
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 int64 = 42
    if false {
        fmt.Println(v13, p1)
        p2 = (p2 | (p2 / ((p2) | 1)))
        for fi14 := 0; fi14 < 10; fi14++ {
            v15 := GenS2{F0: (p1 << ((588) & 63)), F1: "foo bar"}
            switch "hello" {
            case "\n\t":
                var v16 uint32 = 2
                _ = v16
            case "foo bar":
                var v17 int = fi14
                _ = v17
            }
            p1 = (p1 >> ((160) & 63))
            _ = fi14
            _ = v15
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return -9223372036854775808
}

func main() {
    v1 := []uint32{}
    v2 := []int{128}
    v3 := []int{127}
    v4 := make(map[uint32]uint32)
    v5 := make(map[string]uint32)
    _ = len(v2)
    v18 := v4[42]
    for fi19 := 0; fi19 < 8; fi19++ {
        switch ("\n\t" + ((((("foo bar" + "foo bar") + "") + ("a" + "\n\t")) + "\n\t") + "hello")) {
        case "foo bar":
            if (!(65535 != v18)) {
                delete(v5, "hello")
                v18 += 2
                var v20 int8 = -128
                _ = v20
            }
        }
        _ = fi19
    }
    if (v18 > 65535) {
        v21 := GenS3{F0: 2}
        var v22 int16 = -1
        fmt.Println(v22, v18)
        _ = v21
        _ = v22
    } else {
        v23 := GenS1{F0: 42}
        v24 := GenS1{F0: 77}
        _ = v23
        _ = v24
    }
    if (132 == 2) {
        if false {
            var v25 uint = 2
            if false {
                var v26 uint32 = v18
                if (1 == -254) {
                    v18 = 490
                    if len(v1) > 0 {
                        v1[0] = 2147483647
                    }
                    for fi27 := 0; fi27 < 11; fi27++ {
                        for fi28 := 0; fi28 < 6; fi28++ {
                            fmt.Println(fi28)
                            v29 := GenS4{F0: "", F1: -1, F2: 126}
                            _ = fi28
                            _ = v29
                        }
                        _ = fi27
                    }
                }
                fmt.Println(v26)
                _ = v26
            } else {
                var v30 int16 = 57
                switch -137 {
                case 9223372036854775807:
                    v31 := GenS5{F0: 2}
                    _ = v31
                default:
                    if true {
                        if true {
                            for fi32 := 0; fi32 < 15; fi32++ {
                                switch ("hello" + ("foo bar" + ("a" + ("a" + (("\n\t" + ("\n\t" + "")) + ("foo bar" + "a")))))) {
                                case "a":
                                    var v33 uint8 = 127
                                    _ = v33
                                default:
                                    v30 = (v30 - v30)
                                }
                                _ = fi32
                            }
                            v34 := GenS4{F0: "foo bar", F1: 32768, F2: -299}
                            v18 += 65535
                            _ = v34
                        }
                        v18 = v18
                        v30 = -537
                    }
                }
                v25 = 922
                _ = v30
            }
            _ = v25
        }
        var v35 float32 = (float32(-2.25) - 0.1)
        for fi36 := 0; fi36 < 19; fi36++ {
            if (genFn1(1, -216, 2) == fi36) {
                v37 := GenS5{F0: 122}
                v38 := GenS1{F0: -1}
                var v39 uint64 = 991
                _ = v37
                _ = v38
                _ = v39
            } else {
                v40 := GenS3{F0: 1}
                fmt.Println(fi36, v18, fi36)
                _ = v40
            }
            for fi41 := 0; fi41 < 6; fi41++ {
                if true {
                    switch "" {
                    case "":
                        _ = len(v3)
                    }
                    v35 -= v35
                }
                switch "hello" {
                case "foo bar":
                    switch 4294967295 {
                    case 4294967295:
                        v35 = (-2.25 / (1e10 - ((-2.25 - 1e10) * (v35 - (v35 + 1e10)))))
                    case 32768:
                        switch ("a" + "a") {
                        case "\n\t":
                            v35 = -2.25
                        case "foo bar":
                            _ = len(v4)
                        case "a":
                            var v42 string = "hello"
                            _ = v42
                        default:
                            v43 := GenS6{F0: (fi41 | fi41), F1: 380}
                            _ = v43
                        }
                    case -9223372036854775808:
                        var v44 int32 = 818
                        _ = v44
                    }
                }
                _ = fi41
            }
            _ = fi36
        }
        _ = v35
    } else {
        for fi45 := 0; fi45 < 7; fi45++ {
            fmt.Println(v18, fi45)
            _ = fi45
        }
        var v46 string = ""
        var v47 uint64 = 4294967295
        _ = v46
        _ = v47
    }
    v48 := GenS4{F0: "", F1: 32767, F2: 170}
    var v49 float32 = -2.25
    if (-1 <= -128) {
        v50 := GenS7{F0: 1, F1: (("foo bar" + "\n\t") + "a"), F2: 866}
        _ = v50
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
}
