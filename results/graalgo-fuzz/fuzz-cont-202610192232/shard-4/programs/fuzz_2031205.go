package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint32; F1 uint64 }
type GenS7 struct { F0 string; F1 uint16 }

func main() {
    var v1 uint16 = 32767
    v2 := make(map[string]int)
    var v3 int32 = -2147483648
    var v4 uint32 = 749
    var v5 uint64 = 42
    v6 := make(map[string]string)
    v7 := GenS1{F0: 100, F1: v5}
    delete(v2, "hello")
    var v8 uint8 = uint8(v4)
    delete(v6, "hello")
    if true {
        v9 := GenS1{F0: 2, F1: (v5 | (^(v5)))}
        v10 := GenS2{F0: v8, F1: 1, F2: 0}
        _ = v10
        _ = v9
    }
    v4 = 2147483648
    for fi11 := 0; fi11 < 1; fi11++ {
        v12 := GenS3{F0: v3, F1: 0}
        v13 := GenS2{F0: (v8 + 95), F1: v1, F2: v4}
        fmt.Println(v4, fi11)
        _ = fi11
        _ = v12
        _ = v13
    }
    v5 = (v5 % ((4294967295) | 1))
    for fi14 := 0; fi14 < 14; fi14++ {
        var v15 uint64 = v5
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        fmt.Println(v4, fi16)
        v1 = v1
        for fi17 := 0; fi17 < 20; fi17++ {
            v18 := v6["\n\t"]
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        var v20 string = ("foo bar" + "")
        switch fi19 {
        case -149:
            v4 = (v4 ^ (v4 << ((681) & 63)))
        case 32768:
            v21 := GenS4{F0: v20}
            _ = v21
        }
        _ = fi19
        _ = v20
    }
    v22 := []int{}
    v4 -= (v4 ^ 539)
    for fi23 := 0; fi23 < 9; fi23++ {
        for fi24 := 0; fi24 < 12; fi24++ {
            var v25 float32 = 3.14159
            _ = fi24
            _ = v25
        }
        if false {
            if ("\n\t" != "a") {
                if (!false) {
                    v8 = (v8 >> ((2) & 63))
                    if false {
                        var v26 float64 = 0.0
                        v8 += (v8 + (^(v8)))
                        _ = v26
                    } else {
                        v27 := v6["a"]
                        _ = v27
                    }
                    _ = len(v22)
                } else {
                    v4 += (v4 + 2)
                }
                var v28 bool = false
                if (((("foo bar" + (("" + "") + "\n\t")) + "foo bar") == ("foo bar" + (("hello" + "foo bar") + (("foo bar" + ("" + "hello")) + (("a" + ("hello" + "\n\t")) + "hello"))))) && ((v3 + v3) == v3)) {
                    delete(v6, "")
                    v4 = (v4 / ((65535) | 1))
                }
                _ = v28
            } else {
                for fi29 := 0; fi29 < 17; fi29++ {
                    v30 := GenS3{F0: (v3 / ((1) | 1)), F1: (v5 % ((v5) | 1))}
                    v31 := GenS5{F0: "\n\t"}
                    var v32 int8 = 1
                    _ = fi29
                    _ = v30
                    _ = v31
                    _ = v32
                }
                _ = len(v6)
                for fi33 := 0; fi33 < 15; fi33++ {
                    _ = len(v6)
                    switch "foo bar" {
                    case "foo bar":
                        switch ("\n\t" + "\n\t") {
                        case "foo bar":
                            v4 = (^(v4))
                        }
                    case "\n\t":
                        if (!true) {
                            if false {
                                v4 = v4
                            } else {
                                v3 -= -645
                                var v34 uint = 760
                                v35 := GenS2{F0: (^(v8)), F1: uint16(v34), F2: (^(v4))}
                                _ = v34
                                _ = v35
                            }
                            var v36 int16 = 42
                            var v37 float64 = 1.5
                            _ = v36
                            _ = v37
                        }
                    default:
                        if false {
                            v1 = 255
                            if true {
                                fmt.Println(v3, v1, v5)
                                var v38 int32 = v3
                                _ = v38
                            }
                            if true {
                                fmt.Println(fi23, v8)
                                v22 = append(v22, 252)
                            } else {
                                v4 += (v4 & (v4 << ((18446744073709551615) & 63)))
                            }
                        } else {
                            switch ((("hello" + "hello") + "a") + (("\n\t" + ("hello" + "\n\t")) + ("\n\t" + ""))) {
                            case "foo bar":
                                fmt.Println(v1, v8, fi23)
                            }
                            var v39 int64 = int64(v3)
                            _ = v39
                        }
                    }
                    v1 -= 32767
                    _ = fi33
                }
            }
            v40 := GenS6{F0: (v4 - 4294967295), F1: v5}
            v2["a"] = -60
            _ = v40
        } else {
            for fi41 := 0; fi41 < 2; fi41++ {
                fmt.Println(fi23, v3, v1)
                _ = fi41
            }
            v1 += 32767
            if (1 != v3) {
                fmt.Println(v8)
                fmt.Println(v4, v3)
                v8 = (v8 * v8)
            } else {
                v42 := GenS7{F0: "hello", F1: v1}
                _ = v42
            }
        }
        _ = fi23
    }
    var v43 bool = (!(v4 < 784))
    v5 += 541
    v43 = false
    for fi44 := 0; fi44 < 12; fi44++ {
        var v45 float32 = 0.0
        _ = fi44
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
