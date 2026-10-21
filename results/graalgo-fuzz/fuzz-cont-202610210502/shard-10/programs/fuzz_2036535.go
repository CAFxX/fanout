package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 int; F1 int8; F2 uint64 }
type GenS4 struct { F0 int; F1 string }
func genFn0() int {
    for fi7 := 0; fi7 < 9; fi7++ {
        fmt.Println(fi7)
        _ = fi7
    }
    var v8 int16 = 1
    var v9 uint16 = 235
    _ = v8
    _ = v9
    return 205
}

func main() {
    v1 := make(map[string]string)
    var v2 int = -21
    var v3 bool = false
    var v4 uint16 = 32767
    v5 := []string{"a", "a"}
    v6 := make(map[uint32]int)
    for fi10 := 0; fi10 < 11; fi10++ {
        v11 := v1["hello"]
        _ = len(v5)
        _ = fi10
        _ = v11
    }
    var v12 uint = 693
    v13 := GenS1{F0: 1, F1: -32768}
    v3 = v3
    if v3 {
        for fi14 := 0; fi14 < 16; fi14++ {
            v4 = 1
            _ = fi14
        }
        if v3 {
            var v15 uint = 65535
            for fi16 := 0; fi16 < 12; fi16++ {
                switch ((("a" + "a") + (("foo bar" + "\n\t") + "")) + "a") {
                case "":
                    switch "\n\t" {
                    case "foo bar":
                        if v3 {
                            var v17 int32 = 2147483647
                            v15 -= v15
                            var v18 float64 = 0.1
                            _ = v17
                            _ = v18
                        }
                    }
                default:
                    for fi19 := 0; fi19 < 18; fi19++ {
                        switch (v2 >> (((v15 >> (((v12 % ((544) | 1))) & 63))) & 63)) {
                        case 128:
                            v15 = 770
                        case 127:
                            _ = len(v1)
                        default:
                            v20 := GenS1{F0: 1, F1: 806}
                            _ = v20
                        }
                        v5 = append(v5, ("\n\t" + ("hello" + "a")))
                        var v21 float32 = 0.1
                        _ = fi19
                        _ = v21
                    }
                }
                switch 32767 {
                case 127:
                    v15 += 18446744073709551615
                case -950:
                    fmt.Println(v15, v12, v4)
                case 42:
                    v22 := GenS1{F0: 42, F1: 42}
                    _ = v22
                }
                for fi23 := 0; fi23 < 1; fi23++ {
                    var v24 string = ("\n\t" + "foo bar")
                    if v3 {
                        var v25 bool = true
                        v26 := GenS1{F0: 4294967295, F1: 0}
                        v12 = v15
                        v2 -= (fi16 >> (((v15 >> ((v15) & 63))) & 63))
                        _ = v25
                        _ = v26
                    } else {
                        var v27 uint32 = 2147483648
                        v28 := GenS1{F0: 42, F1: 335}
                        _ = v27
                        _ = v28
                    }
                    v24 = (v24 + (v24 + "a"))
                    _ = fi23
                    _ = v24
                }
                _ = fi16
            }
            _ = v15
        }
    } else {
        fmt.Println(v3, v4)
        for fi29 := 0; fi29 < 13; fi29++ {
            var v30 string = (("\n\t" + ("hello" + "hello")) + (("foo bar" + "a") + ("\n\t" + "\n\t")))
            if (true && v3) {
                switch "\n\t" {
                case "a":
                    if v3 {
                        fmt.Println(v4)
                        _ = len(v5)
                    } else {
                        for fi31 := 0; fi31 < 15; fi31++ {
                            v5 = append(v5, v30)
                            _ = fi31
                        }
                        var v32 uint64 = 42
                        v32 = v32
                        _ = v32
                    }
                default:
                    switch v2 {
                    case 128:
                        for fi33 := 0; fi33 < 4; fi33++ {
                            var v34 uint8 = 40
                            _ = fi33
                            _ = v34
                        }
                    }
                }
            } else {
                if v3 {
                    var v35 int32 = -2147483648
                    switch (v30 + v30) {
                    case "foo bar":
                        if len(v5) > 0 {
                            v5[0] = "\n\t"
                        }
                    }
                    _ = v35
                } else {
                    v3 = (4294967295 < 2)
                    var v36 string = "foo bar"
                    _ = v36
                }
            }
            switch -245 {
            case -910:
                switch "hello" {
                case "foo bar":
                    var v37 float64 = float64(0.1)
                    _ = v37
                case "":
                    for fi38 := 0; fi38 < 15; fi38++ {
                        var v39 uint32 = 909
                        _ = fi38
                        _ = v39
                    }
                case "hello":
                    v40 := GenS2{F0: -32768, F1: (v30 + v30)}
                    _ = v40
                default:
                    if true {
                        for fi41 := 0; fi41 < 15; fi41++ {
                            v3 = v3
                            _ = fi41
                        }
                    }
                }
            default:
                v12 = (v12 << ((v12) & 63))
            }
            _ = fi29
            _ = v30
        }
        if (false && false) {
            if v3 {
                var v42 string = "hello"
                switch v42 {
                case "":
                    v43 := GenS2{F0: -32768, F1: (v42 + "foo bar")}
                    _ = v43
                case "\n\t":
                    for fi44 := 0; fi44 < 14; fi44++ {
                        delete(v6, 675)
                        _ = fi44
                    }
                default:
                    v45 := GenS3{F0: genFn0(), F1: 42, F2: 4294967295}
                    _ = v45
                }
                _ = v42
            }
        }
    }
    var v46 uint32 = 822
    v47 := GenS4{F0: (v2 | 618), F1: "\n\t"}
    fmt.Println(v4)
    v3 = false
    fmt.Println(v3, v4, v3)
    v48 := GenS3{F0: v2, F1: 42, F2: 18446744073709551615}
    fmt.Println(v46, v3, v12)
    v3 = ((((((("foo bar" + "a") + "") + "hello") + (("hello" + "hello") + "")) + ("foo bar" + "a")) + "") == ("foo bar" + ""))
    if false {
        for fi49 := 0; fi49 < 12; fi49++ {
            var v50 int64 = 0
            _ = fi49
            _ = v50
        }
    } else {
        v3 = (231 > 1)
        var v51 int = (v2 % (((v2 * genFn0())) | 1))
        v4 = v4
        _ = v51
    }
    if (!v3) {
        v12 = (v12 - v12)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
