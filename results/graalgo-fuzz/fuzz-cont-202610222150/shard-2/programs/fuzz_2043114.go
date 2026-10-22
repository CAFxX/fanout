package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 string; F1 string; F2 bool }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 string; F1 int8; F2 int16 }

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 429
    v3 := make(map[uint32]int)
    v4 := make(map[string]bool)
    var v5 int16 = 0
    v6 := make(map[string]string)
    switch 2147483647 {
    case 128:
        v2 += 2147483648
    default:
        for fi7 := 0; fi7 < 7; fi7++ {
            switch "a" {
            case "hello":
                v8 := []int{42}
                _ = v8
            default:
                fmt.Println(fi7, v2, v5)
            }
            var v9 int8 = -128
            for fi10 := 0; fi10 < 10; fi10++ {
                v2 = (v2 * (^(v2)))
                var v11 uint64 = 1
                _ = fi10
                _ = v11
            }
            _ = fi7
            _ = v9
        }
    }
    var v12 uint = 0
    v13 := v6[("hello" + "")]
    v14 := []int{0, 128}
    for fi15 := 0; fi15 < 8; fi15++ {
        if (false && (true && (v2 == 658))) {
            var v16 uint16 = 2
            _ = v16
        } else {
            delete(v1, (v13 + (v13 + ((v13 + (v13 + v13)) + v13))))
            for fi17 := 0; fi17 < 7; fi17++ {
                v18 := GenS1{F0: 897, F1: v13}
                var v19 float32 = 0.0
                if true {
                    if true {
                        if (2147483647 == 9223372036854775807) {
                            if false {
                                v5 = (v5 | 0)
                            }
                        } else {
                            if false {
                                _ = len(v14)
                                v14 = append(v14, (fi17 << ((v12) & 63)))
                                var v20 int32 = 2147483647
                                _ = v20
                            } else {
                                v19 = v19
                            }
                            v21 := GenS1{F0: 53, F1: v13}
                            switch (fi17 | (fi17 | fi17)) {
                            case 9223372036854775807:
                                var v22 bool = true
                                _ = v22
                            case 42:
                                v23 := GenS2{F0: v13, F1: "a", F2: true}
                                _ = v23
                            default:
                                fmt.Println(v19, v19)
                            }
                            _ = v21
                        }
                        v13 = v13
                    } else {
                        v4[v13] = true
                        if ((v12 >= (v12 & v12)) && (32767 <= 721)) {
                            v5 = v5
                            v12 = v12
                            var v24 uint16 = 1
                            _ = v24
                        }
                    }
                    switch (v13 + v13) {
                    case "foo bar":
                        var v25 uint8 = 255
                        _ = v25
                    case "\n\t":
                        var v26 int8 = -1
                        _ = v26
                    case "a":
                        if true {
                            if false {
                                if (false || false) {
                                    if (204 <= 1) {
                                        var v27 bool = (fi17 > -1)
                                        _ = v27
                                    }
                                    v13 = (((v13 + v13) + ((("foo bar" + ((v13 + (v13 + v13)) + v13)) + "foo bar") + ((("hello" + "\n\t") + "hello") + ""))) + "a")
                                }
                            }
                            v13 = v13
                            switch (^(fi15)) {
                            case 758:
                                v28 := GenS3{F0: 0, F1: -9223372036854775808}
                                _ = v28
                            case 32767:
                                var v29 uint = 0
                                _ = v29
                            }
                            switch fi15 {
                            case -2147483648:
                                v13 = (("a" + (("\n\t" + "a") + "")) + "a")
                            }
                            v5 += 32767
                        }
                    }
                    switch (v13 + "foo bar") {
                    case "a":
                        if true {
                            fmt.Println(v19)
                            fmt.Println(v13, v5)
                        }
                    case "\n\t":
                        var v30 uint16 = 65535
                        _ = v30
                    default:
                        v5 -= 642
                        var v31 int = fi17
                        var v32 string = "\n\t"
                        v33 := GenS4{F0: (v32 + ((v13 + "a") + v13)), F1: 55, F2: (v5 / ((v5) | 1))}
                        v34 := GenS4{F0: ("foo bar" + v32), F1: -47, F2: -32768}
                        _ = v31
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                }
                _ = fi17
                _ = v18
                _ = v19
            }
        }
        _ = fi15
    }
    for fi35 := 0; fi35 < 7; fi35++ {
        var v36 uint8 = 42
        var v37 int64 = 0
        _ = len(v4)
        _ = fi35
        _ = v36
        _ = v37
    }
    var v38 int = 1
    var v39 uint8 = 128
    v13 = (v13 + ("a" + v13))
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    var v40 uint64 = 451
    _ = len(v14)
    for fi41 := 0; fi41 < 16; fi41++ {
        if (!((((243 == 37) && false) && (!true)) && false)) {
            v2 -= (v2 ^ v2)
            v42 := v1[""]
            _ = v42
        }
        _ = fi41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
