package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 uint16 }
func genFn0(p0 uint8, p1 int, p2 uint32) string {
    for fi7 := 0; fi7 < 7; fi7++ {
        p1 = (fi7 | p1)
        switch "\n\t" {
        case "":
            var v8 int64 = 315
            _ = v8
        case "a":
            var v9 bool = false
            _ = v9
        }
        _ = fi7
    }
    v10 := GenS1{F0: -97, F1: -65, F2: ("foo bar" + "\n\t")}
    var v11 int32 = 568
    if (!(true && (29 != (p0 / ((42) | 1))))) {
        v12 := GenS1{F0: -1, F1: 127, F2: ("hello" + "")}
        if false {
            p0 += 42
        } else {
            p1 = (-(p1))
        }
        v13 := []int{p1}
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return "a"
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]int)
    v3 := make(map[string]string)
    v4 := make(map[uint32]uint32)
    var v5 uint = 42
    var v6 int = 9223372036854775807
    v5 -= (v5 & 267)
    if (42 != 0) {
        v14 := GenS1{F0: 4, F1: 0, F2: ""}
        for fi15 := 0; fi15 < 5; fi15++ {
            switch "a" {
            case "\n\t":
                if false {
                    var v16 int16 = 42
                    v1[-902] = "hello"
                    _ = v16
                } else {
                    delete(v1, (fi15 * fi15))
                    for fi17 := 0; fi17 < 1; fi17++ {
                        v5 = v5
                        v18 := []int{}
                        switch "hello" {
                        case "\n\t":
                            if true {
                                if true {
                                    _ = len(v18)
                                    v5 += 750
                                } else {
                                    fmt.Println(v5)
                                }
                            } else {
                                var v19 bool = (false || true)
                                if v19 {
                                    var v20 uint8 = 128
                                    _ = v20
                                }
                                delete(v2, "a")
                                _ = v19
                            }
                        default:
                            v21 := GenS2{F0: 18446744073709551615}
                            _ = v21
                        }
                        v22 := GenS2{F0: 18446744073709551615}
                        var v23 int = v6
                        switch v6 {
                        case 42:
                            var v24 int16 = -1
                            _ = v24
                        }
                        var v25 float32 = 0.1
                        _ = fi17
                        _ = v18
                        _ = v22
                        _ = v23
                        _ = v25
                    }
                }
            }
            if false {
                v5 = (^(v5))
            }
            if false {
                delete(v4, 302)
                v26 := GenS3{F0: "a", F1: 32767}
                var v27 uint16 = 32767
                _ = v26
                _ = v27
            } else {
                var v28 string = ((((("a" + "\n\t") + "") + "") + "hello") + "\n\t")
                v28 = v28
                v29 := GenS3{F0: v28, F1: 32767}
                _ = v28
                _ = v29
            }
            _ = fi15
        }
        _ = v14
    }
    v6 = v6
    v30 := []int{(v6 & (^(v6))), 1, (^(v6)), (^(v6))}
    v6 = (v6 % ((128) | 1))
    var v31 float64 = ((1.5 / 1.0) - 1.5)
    var v32 int = (v6 - v6)
    var v33 uint32 = 2147483648
    v34 := GenS3{F0: ("" + ("a" + "a")), F1: 2}
    v35 := GenS1{F0: 0, F1: 127, F2: (("foo bar" + "") + "\n\t")}
    var v36 uint = 0
    var v37 uint64 = 18446744073709551615
    if true {
        v33 = (v33 ^ v33)
        switch (((("foo bar" + "foo bar") + "a") + ("\n\t" + "foo bar")) + "\n\t") {
        case "":
            v38 := GenS3{F0: "hello", F1: 65535}
            _ = v38
        case "a":
            v5 = 18446744073709551615
        }
        var v39 uint16 = 255
        _ = v39
    }
    var v40 int64 = -1
    var v41 int = 4294967295
    v33 += (v33 << ((v5) & 63))
    v33 -= (v33 & (v33 - (v33 >> (((v36 / (((^(v5))) | 1))) & 63))))
    for fi42 := 0; fi42 < 19; fi42++ {
        var v43 int64 = v40
        _ = fi42
        _ = v43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
