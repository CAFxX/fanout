package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int; F1 uint8; F2 int8 }
type GenS3 struct { F0 bool; F1 int16; F2 int32 }
type GenS4 struct { F0 uint16; F1 uint32; F2 uint32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int32; F1 int }

func main() {
    var v1 int32 = 42
    v2 := []int{}
    v3 := make(map[string]uint32)
    var v4 uint64 = 798
    v5 := GenS1{F0: "foo bar", F1: ("foo bar" + "a")}
    v6 := GenS1{F0: "\n\t", F1: "\n\t"}
    if true {
        var v7 uint8 = 87
        _ = v7
    }
    var v8 int32 = (v1 << ((18446744073709551615) & 63))
    var v9 string = ("foo bar" + (("a" + "\n\t") + (("hello" + ((("" + "\n\t") + ("" + "a")) + "")) + ("\n\t" + "a"))))
    if true {
        if false {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi10 := 0; fi10 < 19; fi10++ {
                v11 := GenS2{F0: fi10, F1: 42, F2: 126}
                if (true || false) {
                    for fi12 := 0; fi12 < 3; fi12++ {
                        if true {
                            if (!true) {
                                v13 := GenS3{F0: ((false && true) && true), F1: -1, F2: v8}
                                fmt.Println(v4, v8, v8)
                                v14 := GenS3{F0: true, F1: -1, F2: 2147483647}
                                _ = v13
                                _ = v14
                            } else {
                                if true {
                                    var v15 bool = false
                                    if ((4294967295 == 2) && (v15 && v15)) {
                                        fmt.Println(fi10, fi10)
                                        v15 = ((v9 + "hello") <= (v9 + "hello"))
                                    } else {
                                        v16 := GenS4{F0: 65535, F1: 42, F2: 633}
                                        if v15 {
                                            fmt.Println(v1, v9, v1)
                                            var v17 int64 = 767
                                            _ = v17
                                        } else {
                                            _ = len(v2)
                                            v2 = append(v2, fi10)
                                        }
                                        _ = v16
                                    }
                                    if v15 {
                                        if (!v15) {
                                            v18 := GenS2{F0: (-(fi10)), F1: 127, F2: 0}
                                            _ = v18
                                        } else {
                                            fmt.Println(v1)
                                            v9 = v9
                                            var v19 uint16 = 32767
                                            _ = v19
                                        }
                                    }
                                    _ = v15
                                }
                                v20 := v3[(v9 + v9)]
                                var v21 int64 = -878
                                _ = v20
                                _ = v21
                            }
                        }
                        v22 := GenS1{F0: v9, F1: ("hello" + v9)}
                        var v23 uint64 = (v4 | (v4 ^ v4))
                        _ = fi12
                        _ = v22
                        _ = v23
                    }
                    fmt.Println(v1, v9)
                }
                _ = len(v3)
                _ = fi10
                _ = v11
            }
        }
        var v24 bool = false
        _ = v24
    } else {
        v1 += v1
    }
    if len(v2) > 0 {
        v2[0] = 1
    }
    if (false || (65535 >= 612)) {
        for fi25 := 0; fi25 < 1; fi25++ {
            v26 := GenS4{F0: 682, F1: 65535, F2: 1}
            _ = fi25
            _ = v26
        }
    } else {
        fmt.Println(v4, v8, v9)
        switch "foo bar" {
        case "foo bar":
            for fi27 := 0; fi27 < 1; fi27++ {
                for fi28 := 0; fi28 < 20; fi28++ {
                    v8 -= 42
                    var v29 int32 = (v1 & -1)
                    _ = fi28
                    _ = v29
                }
                _ = fi27
            }
        case "":
            var v30 uint = 1
            _ = v30
        case "\n\t":
            var v31 uint64 = v4
            _ = v31
        default:
            for fi32 := 0; fi32 < 10; fi32++ {
                for fi33 := 0; fi33 < 10; fi33++ {
                    if true {
                        v3["\n\t"] = 42
                        fmt.Println(fi33, fi33)
                        fmt.Println(fi33)
                        if (true || false) {
                            switch v9 {
                            case "hello":
                                var v34 float32 = 0.1
                                _ = v34
                            case "foo bar":
                                if (true || (fi32 > fi33)) {
                                    if ((-(v8)) > 1) {
                                        var v35 int64 = 514
                                        v1 -= -325
                                        _ = v35
                                    } else {
                                        v36 := GenS1{F0: v9, F1: "\n\t"}
                                        v4 += (v4 & v4)
                                        _ = len(v2)
                                        v37 := GenS5{F0: 85}
                                        _ = v36
                                        _ = v37
                                    }
                                } else {
                                    if true {
                                        v3[("foo bar" + v9)] = 682
                                        var v38 uint8 = 255
                                        if false {
                                            fmt.Println(v38)
                                            v39 := GenS3{F0: false, F1: -32768, F2: v8}
                                            if (!(((0 == 556) && true) || false)) {
                                                fmt.Println(fi33, v4)
                                                _ = len(v3)
                                                v9 = ""
                                            }
                                            _ = v39
                                        } else {
                                            switch (fi33 * (fi33 * 42)) {
                                            case 227:
                                                v8 = (-(v8))
                                            }
                                            _ = len(v2)
                                            v9 = "foo bar"
                                        }
                                        var v40 uint = 734
                                        _ = v38
                                        _ = v40
                                    }
                                }
                            }
                            _ = len(v2)
                            v41 := GenS6{F0: -2147483648, F1: 498}
                            _ = v41
                        }
                    } else {
                        fmt.Println(fi33, v4, fi32)
                    }
                    _ = fi33
                }
                _ = fi32
            }
        }
        v1 = (v1 >> ((310) & 63))
    }
    var v42 uint = 255
    v43 := GenS1{F0: (v9 + ("\n\t" + ("" + v9))), F1: ("" + "a")}
    fmt.Println(v1, v8)
    fmt.Println(v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
