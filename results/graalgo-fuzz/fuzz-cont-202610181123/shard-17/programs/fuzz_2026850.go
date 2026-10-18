package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 int16; F2 bool }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 int8; F1 int; F2 string }
type GenS5 struct { F0 uint32; F1 int; F2 string }
type GenS6 struct { F0 string }

func main() {
    var v1 int32 = 0
    var v2 string = ""
    var v3 int8 = 100
    v1 = (v1 & v1)
    for fi4 := 0; fi4 < 6; fi4++ {
        var v5 uint8 = 0
        _ = fi4
        _ = v5
    }
    fmt.Println(v3, v2)
    if (!false) {
        if true {
            var v6 int64 = 0
            v7 := GenS1{F0: 240}
            _ = v6
            _ = v7
        } else {
            if false {
                if (false || false) {
                    v8 := GenS2{F0: 1, F1: -1, F2: (true || (127 > 109))}
                    v3 = (^(v3))
                    _ = v8
                }
                v9 := GenS2{F0: v3, F1: -277, F2: (1 == 0)}
                v10 := GenS3{F0: -486, F1: (v2 + "")}
                _ = v10
                _ = v9
            } else {
                v11 := []int{9223372036854775807, 128, 42}
                var v12 int8 = (v3 | 127)
                _ = v11
                _ = v12
            }
            var v13 float32 = -2.25
            _ = v13
        }
        fmt.Println(v1, v2)
        switch (("hello" + v2) + "") {
        case "\n\t":
            v14 := GenS4{F0: v3, F1: 32767, F2: "hello"}
            _ = v14
        case "":
            for fi15 := 0; fi15 < 17; fi15++ {
                var v16 uint32 = 65535
                v17 := GenS1{F0: 571}
                _ = fi15
                _ = v16
                _ = v17
            }
        }
    } else {
        v18 := []int{42, -110, 822}
        switch v2 {
        case "hello":
            v18 = append(v18, -2147483648)
        default:
            fmt.Println(v2)
        }
        _ = v18
    }
    v19 := GenS3{F0: 32767, F1: "\n\t"}
    var v20 uint32 = 1
    if true {
        if true {
            v1 -= v1
        }
    } else {
        switch -1 {
        case 0:
            v21 := GenS4{F0: v3, F1: 0, F2: v2}
            _ = v21
        case 42:
            if (v1 < (v1 << ((255) & 63))) {
                v22 := GenS4{F0: v3, F1: 127, F2: (v2 + v2)}
                v3 = v3
                _ = v22
            }
        case 9223372036854775807:
            v23 := GenS5{F0: (v20 - (v20 + (v20 ^ (v20 - (v20 >> ((255) & 63)))))), F1: -445, F2: (((v2 + v2) + "\n\t") + v2)}
            _ = v23
        default:
            switch 42 {
            case 128:
                for fi24 := 0; fi24 < 14; fi24++ {
                    v25 := []int{1, (fi24 / ((42) | 1))}
                    _ = fi24
                    _ = v25
                }
            case 9223372036854775807:
                switch "foo bar" {
                case "foo bar":
                    for fi26 := 0; fi26 < 10; fi26++ {
                        v3 = -85
                        _ = fi26
                    }
                case "hello":
                    v27 := []int{9223372036854775807, 42, 42, 42}
                    _ = v27
                case "":
                    switch -616 {
                    case 9223372036854775807:
                        if true {
                            for fi28 := 0; fi28 < 18; fi28++ {
                                if (-2147483648 >= v1) {
                                    fmt.Println(fi28, v2, fi28)
                                    switch (fi28 >> ((2) & 63)) {
                                    case 42:
                                        v1 -= v1
                                    default:
                                        v1 -= 531
                                    }
                                    var v29 uint64 = 0
                                    _ = v29
                                }
                                v30 := []int{32767, (fi28 ^ -9223372036854775808), fi28}
                                v31 := GenS4{F0: v3, F1: (^(fi28)), F2: "foo bar"}
                                _ = fi28
                                _ = v30
                                _ = v31
                            }
                            for fi32 := 0; fi32 < 17; fi32++ {
                                v1 = 1
                                if (!(32767 <= 918)) {
                                    for fi33 := 0; fi33 < 14; fi33++ {
                                        var v34 uint16 = 2
                                        v35 := []int{fi33, fi33, (fi32 | fi32)}
                                        _ = fi33
                                        _ = v34
                                        _ = v35
                                    }
                                    v1 -= (v1 - v1)
                                } else {
                                    if false {
                                        v36 := GenS3{F0: 32767, F1: v2}
                                        fmt.Println(v1, fi32)
                                        var v37 int8 = -1
                                        _ = v36
                                        _ = v37
                                    }
                                }
                                var v38 int = fi32
                                _ = fi32
                                _ = v38
                            }
                            if ((!true) || true) {
                                if true {
                                    v39 := GenS5{F0: (v20 ^ v20), F1: 32768, F2: ((v2 + (v2 + v2)) + ((v2 + v2) + (("foo bar" + ("\n\t" + ("a" + "\n\t"))) + (v2 + "\n\t"))))}
                                    _ = v39
                                } else {
                                    var v40 float32 = 3.14159
                                    fmt.Println(v40)
                                    _ = v40
                                }
                            }
                        }
                    case 128:
                        for fi41 := 0; fi41 < 14; fi41++ {
                            v1 = (^(v1))
                            var v42 uint8 = 128
                            _ = fi41
                            _ = v42
                        }
                    default:
                        v43 := []int{}
                        _ = v43
                    }
                }
            case -2147483648:
                var v44 string = v2
                _ = v44
            }
        }
    }
    var v45 float32 = 1.5
    v46 := GenS5{F0: v20, F1: 0, F2: "a"}
    if false {
        v47 := GenS2{F0: 127, F1: 42, F2: true}
        if true {
            v1 -= (v1 * (^(v1)))
            v48 := GenS6{F0: ((("\n\t" + v2) + "hello") + "hello")}
            switch 0 {
            case 2147483647:
                v2 = (v2 + (("a" + "") + (v2 + (v2 + v2))))
            default:
                v49 := GenS1{F0: -426}
                _ = v49
            }
            _ = v48
        }
        v45 = 0.0
        _ = v47
    }
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v45
    _ = v46
}
