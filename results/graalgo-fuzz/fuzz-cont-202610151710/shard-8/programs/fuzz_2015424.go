package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int8 }
type GenS2 struct { F0 bool; F1 uint; F2 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 int8; F2 uint32 }
type GenS5 struct { F0 int16; F1 uint }
func genFn0(p0 uint32) string {
    v6 := GenS1{F0: 4294967295, F1: "hello", F2: -26}
    switch "foo bar" {
    case "foo bar":
        fmt.Println(p0)
    default:
        v7 := GenS2{F0: (true && (!(!false))), F1: 1, F2: false}
        _ = v7
    }
    _ = p0
    _ = v6
    return (("foo bar" + (("\n\t" + "hello") + ("\n\t" + ("\n\t" + ("" + ("foo bar" + ((("\n\t" + "hello") + "foo bar") + ("a" + "a")))))))) + "hello")
}
func genFn1() uint32 {
    v8 := []int{-1}
    var v9 float64 = 3.14159
    if (((0.1 > 0.0) && false) && (!((!((4294967295 != 929) || true)) || false))) {
        v9 = ((1.5 + 1e10) / 3.14159)
    } else {
        for fi10 := 0; fi10 < 6; fi10++ {
            v11 := GenS3{F0: -116}
            if (false && true) {
                if false {
                    for fi12 := 0; fi12 < 11; fi12++ {
                        v13 := GenS1{F0: 364, F1: "", F2: 42}
                        v9 = 0.1
                        _ = fi12
                        _ = v13
                    }
                }
            }
            _ = fi10
            _ = v11
        }
        fmt.Println(v9)
    }
    v9 = 0.0
    var v14 uint32 = 65535
    _ = v14
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]bool)
    var v3 int32 = 1
    v4 := []string{""}
    var v5 int = 32768
    v15 := GenS4{F0: 906, F1: 100, F2: 2147483647}
    var v16 uint16 = 438
    if false {
        v16 = 557
        var v17 bool = (!(int32(v5) != v3))
        _ = v17
    } else {
        v16 = 65535
        fmt.Println(v16)
        for fi18 := 0; fi18 < 8; fi18++ {
            if true {
                fmt.Println(v5)
            }
            _ = fi18
        }
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        var v20 uint32 = 42
        var v21 string = "hello"
        _ = fi19
        _ = v20
        _ = v21
    }
    v5 = (v5 & v5)
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    switch (("a" + "\n\t") + "") {
    case "a":
        var v22 float32 = (3.14159 / -2.25)
        _ = v22
    case "hello":
        v3 = (v3 | v3)
    case "\n\t":
        if (32767 < 32767) {
            switch 1 {
            case 230:
                var v23 int32 = (v3 % (((v3 + v3)) | 1))
                _ = v23
            default:
                for fi24 := 0; fi24 < 1; fi24++ {
                    for fi25 := 0; fi25 < 8; fi25++ {
                        if (-9223372036854775808 > -1) {
                            var v26 int64 = -305
                            var v27 uint = 255
                            var v28 int8 = 48
                            _ = v26
                            _ = v27
                            _ = v28
                        } else {
                            fmt.Println(v16, v16, v5)
                            if true {
                                var v29 string = ""
                                _ = v29
                            } else {
                                switch ("\n\t" + "foo bar") {
                                case "":
                                    fmt.Println(fi25)
                                case "hello":
                                    if len(v4) > 0 {
                                        v4[0] = "foo bar"
                                    }
                                default:
                                    var v30 uint16 = 32767
                                    _ = v30
                                }
                            }
                            var v31 int = (fi24 / (((-(fi24))) | 1))
                            _ = v31
                        }
                        fmt.Println(v3)
                        _ = fi25
                    }
                    _ = fi24
                }
            }
            v3 += -2147483648
        }
    }
    fmt.Println(v3, v16)
    v32 := GenS4{F0: 0, F1: 42, F2: 330}
    for fi33 := 0; fi33 < 11; fi33++ {
        delete(v1, (-(fi33)))
        _ = fi33
    }
    v16 += 419
    if true {
        v5 = v5
    }
    for fi34 := 0; fi34 < 8; fi34++ {
        if (32767 >= 42) {
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
        }
        v35 := GenS5{F0: int16(v5), F1: 326}
        _ = fi34
        _ = v35
    }
    var v36 float64 = 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v32
    _ = v36
    _ = v4
    _ = v5
}
