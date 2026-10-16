package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 uint; F2 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int }
func genFn0() uint8 {
    var v5 uint16 = 2
    v6 := []int{42}
    for fi7 := 0; fi7 < 3; fi7++ {
        v5 += 65535
        _ = fi7
    }
    var v8 uint16 = v5
    _ = v5
    _ = v6
    _ = v8
    return 42
}

func main() {
    var v1 int = -118
    var v2 uint16 = 65535
    var v3 int16 = -1
    v4 := make(map[string]bool)
    v9 := []int{(v1 & (v1 >> ((42) & 63)))}
    for fi10 := 0; fi10 < 1; fi10++ {
        for fi11 := 0; fi11 < 5; fi11++ {
            switch (((("a" + "hello") + "\n\t") + "a") + "foo bar") {
            case "\n\t":
                if (true || false) {
                    var v12 string = "hello"
                    var v13 bool = false
                    delete(v4, v12)
                    _ = v12
                    _ = v13
                } else {
                    var v14 int32 = 641
                    var v15 uint64 = 1
                    fmt.Println(v14)
                    _ = v14
                    _ = v15
                }
            case "":
                fmt.Println(fi11)
                fmt.Println(v3, v1, fi10)
                v3 -= (v3 ^ 1000)
                v1 += (fi10 / (((^(v1))) | 1))
                switch (fi10 % (((fi11 / ((v1) | 1))) | 1)) {
                case 32768:
                    v3 = 0
                    if false {
                        v16 := GenS1{F0: 401}
                        var v17 int32 = 670
                        if true {
                            fmt.Println(fi10, fi10)
                            switch -1 {
                            case -9223372036854775808:
                                fmt.Println(v3)
                            case 4294967295:
                                if (!((-2.25 != 0.0) || (v17 >= -2147483648))) {
                                    var v18 int32 = 256
                                    var v19 uint8 = 127
                                    _ = v18
                                    _ = v19
                                }
                            }
                            var v20 float32 = (-2.25 * 1e10)
                            _ = v20
                        }
                        v21 := GenS1{F0: 642}
                        if (!(!true)) {
                            switch (("\n\t" + "foo bar") + (("hello" + (("a" + "a") + "")) + ("foo bar" + "foo bar"))) {
                            case "":
                                v3 = (v3 | 32767)
                            case "\n\t":
                                v1 = -1
                            case "hello":
                                v3 = 1
                                switch fi10 {
                                case 128:
                                    var v22 uint = 1
                                    _ = v22
                                case 0:
                                    v23 := v4["a"]
                                    _ = v23
                                case 2147483647:
                                    fmt.Println(fi11)
                                default:
                                    var v24 uint16 = (v2 / ((v2) | 1))
                                    _ = v24
                                }
                            }
                            switch fi11 {
                            case -558:
                                if false {
                                    _ = len(v9)
                                    var v25 int16 = 1
                                    _ = v25
                                }
                            case -9223372036854775808:
                                v2 = v2
                            default:
                                if (true && false) {
                                    v26 := GenS1{F0: 42}
                                    _ = v26
                                } else {
                                    _ = len(v9)
                                    v27 := GenS2{F0: 42, F1: 0, F2: 1}
                                    _ = v27
                                }
                            }
                            var v28 uint32 = 448
                            v2 = (^(v2))
                            _ = v28
                        } else {
                            switch "\n\t" {
                            case "a":
                                switch (("\n\t" + "hello") + "\n\t") {
                                case "":
                                    v2 = v2
                                case "\n\t":
                                    v29 := GenS1{F0: 0}
                                    _ = v29
                                }
                                var v30 string = "\n\t"
                                _ = v30
                            case "hello":
                                var v31 int = (^(fi10))
                                var v32 uint32 = uint32(v2)
                                _ = v31
                                _ = v32
                            }
                            v3 -= -1
                            var v33 float64 = -2.25
                            _ = v33
                        }
                        _ = v16
                        _ = v17
                        _ = v21
                    }
                case 2147483647:
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                }
            case "foo bar":
                if (!false) {
                    var v34 bool = (false || (false || (!true)))
                    _ = v34
                } else {
                    v35 := GenS2{F0: 1, F1: 152, F2: 469}
                    _ = v35
                }
            }
            _ = fi11
        }
        v3 = v3
        v36 := v4["\n\t"]
        _ = fi10
        _ = v36
    }
    v1 -= 4294967295
    fmt.Println(v3, v3)
    var v37 int = (v1 >> ((42) & 63))
    for fi38 := 0; fi38 < 15; fi38++ {
        v39 := GenS2{F0: genFn0(), F1: 65535, F2: 1}
        _ = fi38
        _ = v39
    }
    v40 := GenS2{F0: genFn0(), F1: 36, F2: 255}
    switch (v1 >> ((464) & 63)) {
    case 2147483647:
        v41 := GenS1{F0: 18446744073709551615}
        _ = v41
    case 4294967295:
        for fi42 := 0; fi42 < 17; fi42++ {
            var v43 uint16 = 42
            var v44 uint64 = 2
            var v45 uint16 = 32767
            _ = fi42
            _ = v43
            _ = v44
            _ = v45
        }
    case 9223372036854775807:
        v37 -= -876
    }
    var v46 int16 = v3
    var v47 bool = false
    v48 := GenS3{F0: 127}
    _ = len(v9)
    v47 = true
    for fi49 := 0; fi49 < 13; fi49++ {
        v50 := GenS4{F0: (-(fi49))}
        _ = fi49
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v40
    _ = v46
    _ = v47
    _ = v48
    _ = v9
}
