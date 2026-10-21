package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int16 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 string; F1 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = 0
    v3 := make(map[string]bool)
    v4 := make(map[int]string)
    v5 := make(map[uint32]int)
    var v6 uint64 = 971
    fmt.Println(v2, v6)
    switch (v2 | v2) {
    case -2147483648:
        fmt.Println(v2, v2)
    case 4294967295:
        v7 := GenS1{F0: "hello", F1: 1, F2: 1000}
        _ = v7
    case 128:
        if true {
            if true {
                var v8 int8 = 0
                v8 += v8
                for fi9 := 0; fi9 < 13; fi9++ {
                    v10 := GenS2{F0: 65535, F1: int32(v2)}
                    var v11 uint32 = 4294967295
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                _ = v8
            } else {
                v12 := []int{}
                _ = v12
            }
            v13 := []int{}
            _ = v13
        }
    }
    delete(v3, "")
    v14 := GenS2{F0: 4294967295, F1: 42}
    fmt.Println(v6, v2)
    v15 := []int{}
    if (0 <= 1) {
        for fi16 := 0; fi16 < 17; fi16++ {
            var v17 float32 = 1.5
            if false {
                var v18 float32 = (((v17 / v17) + v17) + float32(((((float64(3.14159) - 1.5) * (1e10 - 3.14159)) / float64((v17 / v17))) - (((1.5 + -2.25) + (-2.25 * 1.5)) / float64(3.14159)))))
                _ = v18
            } else {
                switch ("hello" + "hello") {
                case "\n\t":
                    _ = len(v15)
                case "a":
                    if (!(2147483647 == fi16)) {
                        delete(v3, ("" + "a"))
                    } else {
                        var v19 float32 = v17
                        var v20 uint = 273
                        for fi21 := 0; fi21 < 16; fi21++ {
                            var v22 float32 = (float32(1e10) + float32((-2.25 / 1.5)))
                            v23 := v5[0]
                            if ((!(false || false)) && false) {
                                v24 := GenS3{F0: "a", F1: -1}
                                var v25 int32 = 0
                                switch ("\n\t" + ("\n\t" + "\n\t")) {
                                case "a":
                                    var v26 float32 = v19
                                    _ = v26
                                case "\n\t":
                                    v22 = 1.5
                                }
                                switch v23 {
                                case 9223372036854775807:
                                    fmt.Println(v23)
                                case 127:
                                    var v27 uint8 = 1
                                    _ = v27
                                }
                                _ = v24
                                _ = v25
                            } else {
                                if false {
                                    v28 := GenS2{F0: (v20 << ((v20) & 63)), F1: 112}
                                    var v29 string = "a"
                                    _ = v28
                                    _ = v29
                                } else {
                                    switch ("foo bar" + "foo bar") {
                                    case "a":
                                        fmt.Println(fi16, fi21)
                                    case "":
                                        fmt.Println(v20, v22)
                                    }
                                }
                            }
                            _ = fi21
                            _ = v22
                            _ = v23
                        }
                        _ = v19
                        _ = v20
                    }
                case "hello":
                    for fi30 := 0; fi30 < 12; fi30++ {
                        fmt.Println(fi30)
                        _ = fi30
                    }
                }
            }
            delete(v3, "")
            _ = fi16
            _ = v17
        }
    } else {
        var v31 int64 = 1
        var v32 int64 = v31
        _ = v31
        _ = v32
    }
    var v33 int64 = 42
    v6 = (^(v6))
    for fi34 := 0; fi34 < 13; fi34++ {
        v33 = (v33 % (((v33 / (((v33 - v33)) | 1))) | 1))
        if true {
            var v35 uint16 = 255
            v36 := GenS1{F0: (((("a" + "") + ("" + ("a" + "a"))) + ("foo bar" + ((("" + "\n\t") + "hello") + ((((("\n\t" + "") + "\n\t") + (("a" + "hello") + ("a" + ("hello" + "")))) + "foo bar") + "")))) + "hello"), F1: (^(v33)), F2: 42}
            _ = v35
            _ = v36
        } else {
            _ = len(v15)
            if false {
                var v37 uint8 = 255
                v33 += (v33 & v33)
                var v38 int64 = v33
                _ = v37
                _ = v38
            } else {
                fmt.Println(fi34)
            }
            v6 += (^(v6))
        }
        _ = fi34
    }
    var v39 bool = ((v2 ^ v2) > v2)
    v40 := GenS1{F0: ("a" + "a"), F1: v33, F2: int16(v6)}
    if ("\n\t" >= "") {
        var v41 int16 = 1
        fmt.Println(v41, v39)
        _ = v41
    }
    var v42 uint64 = 0
    switch 42 {
    case 4294967295:
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
    }
    v42 = v42
    if len(v15) > 0 {
        v15[0] = (v2 << ((4294967295) & 63))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v33
    _ = v39
    _ = v4
    _ = v40
    _ = v42
    _ = v5
    _ = v6
}
