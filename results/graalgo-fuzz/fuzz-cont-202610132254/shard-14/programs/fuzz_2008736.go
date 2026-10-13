package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 int32; F2 uint64 }
type GenS3 struct { F0 uint32; F1 uint32 }

func main() {
    var v1 int8 = 0
    v2 := make(map[string]int)
    var v3 uint16 = 42
    v4 := GenS1{F0: "foo bar"}
    if true {
        var v5 float32 = (3.14159 + 0.0)
        v5 -= (1e10 + v5)
        _ = v5
    } else {
        v6 := v2["hello"]
        _ = v6
    }
    switch "a" {
    case "":
        v7 := GenS1{F0: ("hello" + (("" + "foo bar") + ""))}
        _ = v7
    default:
        if false {
            v3 = v3
            v3 += 255
            v3 += 32767
        }
    }
    v1 -= v1
    v1 = -128
    v3 = (v3 ^ v3)
    switch 1 {
    case 986:
        var v8 int64 = 0
        _ = v8
    case 42:
        if true {
            v3 = 255
            v1 += v1
            for fi9 := 0; fi9 < 15; fi9++ {
                _ = len(v2)
                if false {
                    v10 := []int{fi9, fi9, fi9}
                    _ = v10
                }
                if (false || false) {
                    switch (fi9 - 4294967295) {
                    case 32767:
                        var v11 float64 = -2.25
                        _ = v11
                    case 4294967295:
                        v1 = (v1 + (-(v1)))
                    default:
                        var v12 int8 = v1
                        _ = v12
                    }
                    var v13 uint32 = 834
                    var v14 int8 = (-(v1))
                    _ = v13
                    _ = v14
                }
                _ = fi9
            }
        }
    case 4294967295:
        var v15 uint8 = 1
        _ = v15
    default:
        var v16 float64 = 3.14159
        _ = v16
    }
    switch "foo bar" {
    case "a":
        var v17 string = "\n\t"
        _ = v17
    case "hello":
        v3 = v3
    default:
        var v18 uint = 455
        _ = v18
    }
    switch int(v1) {
    case -769:
        var v19 uint64 = 394
        _ = v19
    case -2147483648:
        v1 = (v1 - v1)
    case -953:
        v2[("" + ("a" + "hello"))] = 0
    }
    var v20 float32 = 1.5
    switch ("hello" + "\n\t") {
    case "foo bar":
        v21 := v2[("foo bar" + "\n\t")]
        _ = v21
    case "":
        switch 42 {
        case -44:
            fmt.Println(v1, v1)
        case 325:
            v1 += v1
        case 0:
            v1 += v1
        }
    case "\n\t":
        fmt.Println(v1)
    default:
        var v22 int64 = -500
        _ = v22
    }
    var v23 uint32 = 1
    for fi24 := 0; fi24 < 18; fi24++ {
        v25 := []int{fi24}
        v1 = (v1 | int8(fi24))
        _ = fi24
        _ = v25
    }
    var v26 bool = true
    v27 := []int{-655, 127, 1}
    fmt.Println(v23, v1, v1)
    if v26 {
        v28 := GenS2{F0: 626, F1: 566, F2: 244}
        fmt.Println(v1, v3)
        _ = v28
    } else {
        switch 32767 {
        case 552:
            switch -559 {
            case 32768:
                v29 := GenS1{F0: (("" + ("hello" + "\n\t")) + "\n\t")}
                _ = v29
            case 42:
                var v30 string = "a"
                _ = v30
            default:
                switch (("hello" + (("\n\t" + ("hello" + "")) + (("hello" + ("a" + "hello")) + ""))) + "") {
                case "a":
                    if (v26 && true) {
                        v23 = 288
                        for fi31 := 0; fi31 < 15; fi31++ {
                            for fi32 := 0; fi32 < 8; fi32++ {
                                v33 := GenS3{F0: v23, F1: 4294967295}
                                _ = fi32
                                _ = v33
                            }
                            if len(v27) > 0 {
                                v27[0] = -2147483648
                            }
                            fmt.Println(v26, fi31, v23)
                            _ = fi31
                        }
                    }
                default:
                    delete(v2, "a")
                }
            }
        case 2147483647:
            switch -2147483648 {
            case 4294967295:
                for fi34 := 0; fi34 < 5; fi34++ {
                    var v35 float64 = float64(v20)
                    switch 0 {
                    case 32767:
                        switch fi34 {
                        case -514:
                            for fi36 := 0; fi36 < 8; fi36++ {
                                var v37 uint64 = 18446744073709551615
                                _ = fi36
                                _ = v37
                            }
                        default:
                            var v38 float64 = ((v35 * v35) + v35)
                            _ = v38
                        }
                    case 728:
                        v26 = v26
                    default:
                        if false {
                            if (42 < 249) {
                                v23 += 640
                                var v39 string = "\n\t"
                                _ = v39
                            }
                            if len(v27) > 0 {
                                v27[0] = (fi34 >> ((18446744073709551615) & 63))
                            }
                        } else {
                            v40 := GenS2{F0: 1, F1: 0, F2: 1}
                            v1 = v1
                            v3 = (v3 + 32767)
                            _ = v40
                        }
                    }
                    _ = fi34
                    _ = v35
                }
            case 42:
                switch "hello" {
                case "a":
                    for fi41 := 0; fi41 < 3; fi41++ {
                        v3 = v3
                        for fi42 := 0; fi42 < 14; fi42++ {
                            var v43 int = (fi41 >> ((42) & 63))
                            _ = fi42
                            _ = v43
                        }
                        _ = fi41
                    }
                }
            case -1:
                v1 = (^(v1))
            }
        }
    }
    if v26 {
        v20 = -2.25
    } else {
        var v44 int = 4294967295
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
