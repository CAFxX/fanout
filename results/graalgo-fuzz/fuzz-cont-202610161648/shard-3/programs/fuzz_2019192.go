package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 int32 }

func main() {
    v1 := make(map[string]bool)
    v2 := []int{32768, 169}
    var v3 uint8 = 193
    var v4 uint8 = 22
    var v5 int8 = -1
    switch -674 {
    case 127:
        v6 := GenS1{F0: false, F1: 1}
        _ = v6
    }
    var v7 bool = true
    switch (("foo bar" + ((("foo bar" + ("\n\t" + ("" + ("hello" + "a")))) + ("foo bar" + ("a" + "hello"))) + "hello")) + "\n\t") {
    case "":
        if v7 {
            v4 = v4
            v4 = v4
            v8 := GenS1{F0: true, F1: 2147483647}
            _ = v8
        } else {
            if ((v4 * v3) != v4) {
                v4 -= v3
            } else {
                if (137 == 0) {
                    fmt.Println(v4, v4)
                } else {
                    v9 := v1["\n\t"]
                    for fi10 := 0; fi10 < 17; fi10++ {
                        for fi11 := 0; fi11 < 18; fi11++ {
                            var v12 uint16 = 42
                            var v13 uint16 = v12
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                        v4 += (v4 * 53)
                        _ = fi10
                    }
                    _ = v9
                }
                var v14 uint = 65535
                _ = v14
            }
            if true {
                v5 += (^(v5))
                if false {
                    fmt.Println(v4)
                    v2 = append(v2, -70)
                } else {
                    var v15 uint16 = 421
                    _ = v15
                }
            }
        }
    case "\n\t":
        if v7 {
            for fi16 := 0; fi16 < 20; fi16++ {
                var v17 string = "a"
                var v18 float32 = (1.5 + 1.5)
                var v19 bool = v7
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
            v3 = v3
        } else {
            v3 = (v3 & (v3 / (((v4 * v4)) | 1)))
        }
    case "foo bar":
        fmt.Println(v4, v5, v4)
    }
    var v20 uint32 = 718
    delete(v1, "\n\t")
    if (false && (606 != 0)) {
        var v21 uint = 0
        v22 := GenS1{F0: false, F1: -922}
        _ = v21
        _ = v22
    } else {
        var v23 int8 = v5
        var v24 uint64 = 4294967295
        _ = v23
        _ = v24
    }
    if v7 {
        switch -258 {
        case -9223372036854775808:
            v20 -= (v20 | v20)
        default:
            v7 = v7
        }
        if (!(4294967295 < 685)) {
            var v25 int16 = 508
            switch "foo bar" {
            case "":
                v4 += v4
            case "\n\t":
                v25 = (v25 | (-(v25)))
            default:
                v26 := GenS2{F0: "a"}
                _ = v26
            }
            _ = v25
        } else {
            if (v7 && true) {
                if v7 {
                    v27 := GenS3{F0: -1, F1: -2147483648}
                    switch "\n\t" {
                    case "":
                        v28 := GenS3{F0: 1, F1: -343}
                        _ = v28
                    case "hello":
                        fmt.Println(v4, v4, v3)
                    }
                    _ = v27
                } else {
                    if (-9223372036854775808 != 120) {
                        delete(v1, ("hello" + ("" + "")))
                        fmt.Println(v7, v7)
                        switch "a" {
                        case "\n\t":
                            var v29 string = ""
                            _ = v29
                        case "hello":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case "":
                            v30 := GenS1{F0: v7, F1: 0}
                            _ = v30
                        }
                    }
                    switch ("hello" + "foo bar") {
                    case "":
                        v20 = 2147483648
                    case "hello":
                        var v31 int64 = 1
                        _ = v31
                    case "\n\t":
                        v32 := GenS3{F0: -482, F1: -906}
                        _ = v32
                    }
                    var v33 uint32 = (v20 >> ((65535) & 63))
                    _ = v33
                }
            } else {
                for fi34 := 0; fi34 < 5; fi34++ {
                    v7 = v7
                    v20 += 2
                    for fi35 := 0; fi35 < 19; fi35++ {
                        v5 = 1
                        switch 32768 {
                        case 1:
                            v20 = 0
                        }
                        v36 := GenS2{F0: (((("\n\t" + ("" + "\n\t")) + (("foo bar" + "") + "a")) + (("a" + "hello") + "foo bar")) + "foo bar")}
                        _ = fi35
                        _ = v36
                    }
                    _ = fi34
                }
            }
            if v7 {
                v1["hello"] = v7
            }
        }
        v37 := GenS2{F0: "foo bar"}
        _ = v37
    }
    v1[("" + "\n\t")] = false
    var v38 int8 = v5
    var v39 uint8 = v4
    var v40 uint8 = v4
    v40 = (v39 << ((65535) & 63))
    v4 = v39
    var v41 int16 = -32768
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v7
}
