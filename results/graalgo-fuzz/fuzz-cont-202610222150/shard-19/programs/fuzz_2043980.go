package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 int8; F1 uint64 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = -1
    v3 := []string{(("hello" + ("foo bar" + "\n\t")) + ("foo bar" + "\n\t")), "foo bar", ""}
    switch (("foo bar" + (("a" + "foo bar") + "\n\t")) + "hello") {
    case "hello":
        if (!(1 >= 127)) {
            for fi4 := 0; fi4 < 16; fi4++ {
                fmt.Println(fi4, fi4)
                if len(v3) > 0 {
                    v3[0] = "a"
                }
                _ = fi4
            }
            v2 += v2
            fmt.Println(v2)
        } else {
            v2 -= (v2 | 0)
        }
    default:
        if false {
            if false {
                _ = len(v3)
            }
            v2 = 42
            for fi5 := 0; fi5 < 12; fi5++ {
                fmt.Println(v2, v2)
                if true {
                    v2 += (v2 >> ((8) & 63))
                }
                _ = fi5
            }
        }
    }
    v6 := GenS1{F0: "foo bar", F1: false}
    for fi7 := 0; fi7 < 3; fi7++ {
        v2 = v2
        switch fi7 {
        case -9223372036854775808:
            var v8 int32 = 0
            _ = v8
        case 1:
            for fi9 := 0; fi9 < 17; fi9++ {
                fmt.Println(v2, fi7, fi7)
                _ = fi9
            }
        default:
            for fi10 := 0; fi10 < 10; fi10++ {
                if true {
                    if true {
                        var v11 uint = 42
                        v2 += (v2 + (v2 ^ (v2 * v2)))
                        _ = v11
                    } else {
                        var v12 uint64 = 42
                        _ = v12
                    }
                }
                v3 = append(v3, "")
                _ = fi10
            }
        }
        _ = fi7
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        for fi14 := 0; fi14 < 6; fi14++ {
            v2 = 42
            _ = fi14
        }
        v3 = append(v3, "a")
        _ = fi13
    }
    if false {
        var v15 bool = true
        if (0 <= 4294967295) {
            fmt.Println(v2, v2)
        } else {
            v2 = v2
        }
        var v16 uint = 0
        _ = v15
        _ = v16
    } else {
        delete(v1, "hello")
        for fi17 := 0; fi17 < 12; fi17++ {
            fmt.Println(fi17, v2)
            for fi18 := 0; fi18 < 3; fi18++ {
                if (fi17 < fi17) {
                    var v19 float32 = 0.0
                    var v20 bool = (!(926 > 2147483648))
                    fmt.Println(fi18, v2)
                    _ = v19
                    _ = v20
                }
                _ = fi18
            }
            _ = fi17
        }
    }
    v21 := GenS2{F0: -1}
    _ = len(v3)
    var v22 uint32 = 0
    var v23 string = ""
    v24 := GenS2{F0: 2147483647}
    v25 := GenS3{F0: 0, F1: 2}
    switch -488 {
    case 4294967295:
        for fi26 := 0; fi26 < 11; fi26++ {
            if true {
                v27 := GenS4{F0: 100, F1: 42}
                v28 := GenS3{F0: -1, F1: 1}
                _ = v27
                _ = v28
            }
            _ = fi26
        }
    case -9223372036854775808:
        for fi29 := 0; fi29 < 12; fi29++ {
            var v30 uint = 736
            for fi31 := 0; fi31 < 15; fi31++ {
                if false {
                    var v32 int64 = v2
                    switch (("a" + (v23 + "a")) + (("foo bar" + (v23 + v23)) + (("foo bar" + (v23 + v23)) + v23))) {
                    case "foo bar":
                        var v33 float32 = ((3.14159 + (3.14159 + float32(0.0))) - 3.14159)
                        _ = v33
                    default:
                        var v34 int8 = 97
                        _ = v34
                    }
                    _ = v32
                } else {
                    v22 += (v22 & v22)
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v1["hello"] = v22
                _ = fi31
            }
            _ = fi29
            _ = v30
        }
    }
    var v35 uint16 = 255
    if (1 > 1) {
        if ((42 > v35) && false) {
            v23 = (v23 + ("a" + ("" + "")))
        }
        v35 = v35
    } else {
        v35 = 686
        v22 -= (^(v22))
    }
    v36 := GenS1{F0: "foo bar", F1: (237 <= 2)}
    v35 -= v35
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v35
    _ = v36
    _ = v6
}
