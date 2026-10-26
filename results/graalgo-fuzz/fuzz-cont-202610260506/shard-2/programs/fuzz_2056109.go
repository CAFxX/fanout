package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 string; F2 bool }
type GenS3 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS4 struct { F0 uint; F1 bool }
func genFn0(p0 bool, p1 uint32) uint64 {
    var v4 bool = p0
    v4 = p0
    var v5 bool = false
    p1 = p1
    for fi6 := 0; fi6 < 19; fi6++ {
        if v4 {
            if p0 {
                v7 := []int{}
                _ = v7
            }
            v8 := GenS1{F0: 236}
            var v9 uint = 18446744073709551615
            _ = v8
            _ = v9
        } else {
            fmt.Println(p1, v5, v4)
        }
        for fi10 := 0; fi10 < 17; fi10++ {
            fmt.Println(p0)
            _ = fi10
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 42
}
func genFn1() uint8 {
    v11 := GenS1{F0: 42}
    var v12 uint64 = 513
    v13 := []int{32767, 0}
    if len(v13) > 0 {
        v13[0] = 128
    }
    _ = v11
    _ = v12
    _ = v13
    return 0
}

func main() {
    var v1 uint16 = 0
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]string)
    v1 = 1
    v14 := GenS2{F0: -1, F1: "hello", F2: (310 > 4294967295)}
    v1 = v1
    switch -2147483648 {
    case 2147483647:
        for fi15 := 0; fi15 < 5; fi15++ {
            v16 := GenS2{F0: 1, F1: ("\n\t" + ("a" + "\n\t")), F2: (1e10 == float64((0.1 - (1.5 - -2.25))))}
            var v17 string = (("foo bar" + "hello") + "foo bar")
            _ = fi15
            _ = v16
            _ = v17
        }
    default:
        var v18 int64 = 0
        _ = v18
    }
    var v19 float64 = 0.0
    v20 := []int{}
    delete(v2, 554)
    var v21 uint16 = 255
    switch "hello" {
    case "":
        if false {
            v22 := GenS3{F0: 1, F1: false, F2: 42}
            if (false && (true && true)) {
                var v23 bool = (((("" + ("hello" + (("hello" + "") + "foo bar"))) < "a") || false) && (-2147483648 >= 1))
                _ = len(v20)
                _ = v23
            } else {
                for fi24 := 0; fi24 < 10; fi24++ {
                    v20 = append(v20, 157)
                    _ = fi24
                }
                var v25 float32 = 1.5
                v25 -= ((v25 * 0.0) + v25)
                _ = v25
            }
            _ = v22
        }
    case "a":
        v1 += v21
    case "hello":
        v1 = 42
    }
    if len(v20) > 0 {
        v20[0] = 803
    }
    switch 42 {
    case 32768:
        v20 = append(v20, 32768)
    default:
        if false {
            if (450 < 18446744073709551615) {
                v26 := GenS1{F0: 2}
                _ = v26
            } else {
                for fi27 := 0; fi27 < 11; fi27++ {
                    var v28 uint64 = 742
                    v28 -= genFn0(true, 2147483647)
                    _ = fi27
                    _ = v28
                }
                if true {
                    switch ("hello" + ("hello" + ("" + "hello"))) {
                    case "a":
                        for fi29 := 0; fi29 < 16; fi29++ {
                            v30 := GenS3{F0: 782, F1: true, F2: 0}
                            if len(v20) > 0 {
                                v20[0] = -9223372036854775808
                            }
                            _ = fi29
                            _ = v30
                        }
                    case "":
                        if (45 <= 1) {
                            fmt.Println(v21, v19, v21)
                        }
                    default:
                        if true {
                            for fi31 := 0; fi31 < 1; fi31++ {
                                v1 = v1
                                _ = fi31
                            }
                            v32 := GenS1{F0: 30}
                            if true {
                                fmt.Println(v19, v1)
                                fmt.Println(v1, v21, v1)
                                var v33 uint8 = genFn1()
                                _ = v33
                            } else {
                                v2[42] = 42
                                v1 += (v1 >> ((2) & 63))
                            }
                            _ = v32
                        }
                    }
                    for fi34 := 0; fi34 < 18; fi34++ {
                        v19 -= v19
                        _ = fi34
                    }
                }
            }
        }
    }
    if false {
        var v35 int32 = 0
        var v36 int = -9223372036854775808
        for fi37 := 0; fi37 < 14; fi37++ {
            v19 = float64(3.14159)
            if false {
                if (!((v36 | v36) <= (fi37 | (v36 + (fi37 - v36))))) {
                    v21 += 1
                    switch fi37 {
                    case 128:
                        v35 = 0
                    }
                    for fi38 := 0; fi38 < 8; fi38++ {
                        var v39 string = ""
                        _ = fi38
                        _ = v39
                    }
                }
                if len(v20) >= 2 {
                    v20 = v20[1:len(v20)]
                }
            } else {
                if ((v1 << ((18446744073709551615) & 63)) <= (v21 & v1)) {
                    v19 += 3.14159
                    switch "hello" {
                    case "":
                        for fi40 := 0; fi40 < 15; fi40++ {
                            fmt.Println(fi37, v35)
                            var v41 uint64 = 176
                            _ = fi40
                            _ = v41
                        }
                    case "foo bar":
                        v42 := GenS1{F0: genFn1()}
                        _ = v42
                    case "a":
                        v20 = append(v20, -2147483648)
                    }
                    delete(v2, 42)
                }
                v43 := GenS4{F0: 0, F1: (!(!(!(("foo bar" != "foo bar") && (false && (false || (255 < 2)))))))}
                switch "\n\t" {
                case "\n\t":
                    if true {
                        v1 = v21
                    }
                case "a":
                    switch "" {
                    case "\n\t":
                        _ = len(v20)
                    }
                }
                _ = v43
            }
            _ = fi37
        }
        _ = v35
        _ = v36
    }
    v44 := GenS1{F0: genFn1()}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v44
}
