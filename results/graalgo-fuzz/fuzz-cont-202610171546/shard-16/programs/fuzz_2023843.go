package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int }

func main() {
    v1 := make(map[string]int)
    var v2 bool = true
    var v3 int64 = 225
    var v4 int32 = -2147483648
    var v5 int8 = 1
    v6 := make(map[string]uint32)
    fmt.Println(v3, v4, v3)
    var v7 uint16 = 0
    switch 128 {
    case 4294967295:
        fmt.Println(v7, v5, v4)
    case 128:
        fmt.Println(v5)
    default:
        v4 = 471
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 int8 = (v5 / ((v5) | 1))
        var v10 bool = (!(0 == -32768))
        fmt.Println(v2, v2)
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := []int{1, 4294967295}
    if v2 {
        v5 += v5
        switch 9223372036854775807 {
        case 128:
            if (false || (!v2)) {
                fmt.Println(v3, v2)
                v11 = append(v11, -9223372036854775808)
                switch ("foo bar" + ("\n\t" + ("hello" + ("" + ("hello" + ("hello" + ("\n\t" + ("foo bar" + ("a" + ("\n\t" + (("\n\t" + "foo bar") + ("foo bar" + "\n\t")))))))))))) {
                case "a":
                    v12 := GenS1{F0: 0}
                    _ = v12
                case "foo bar":
                    if false {
                        v4 += (^(v4))
                        v3 += int64(v4)
                    }
                }
            }
        default:
            v5 = (v5 * (-(v5)))
        }
    }
    v7 -= 255
    for fi13 := 0; fi13 < 7; fi13++ {
        var v14 int64 = v3
        v5 -= (v5 / (((v5 + (v5 | v5))) | 1))
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 6; fi15++ {
        switch "hello" {
        case "hello":
            var v16 uint16 = (v7 ^ (v7 * 901))
            _ = v16
        case "foo bar":
            var v17 int8 = (v5 % ((int8(v7)) | 1))
            _ = v17
        default:
            var v18 uint64 = 42
            _ = v18
        }
        var v19 float32 = (-2.25 / 0.1)
        _ = fi15
        _ = v19
    }
    if ((42 > 18446744073709551615) && false) {
        if (("hello" + "\n\t") == "\n\t") {
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
        } else {
            v20 := GenS1{F0: int64(v7)}
            var v21 int8 = v5
            _ = len(v11)
            _ = v20
            _ = v21
        }
        v22 := GenS1{F0: v3}
        _ = v22
    } else {
        v4 -= v4
    }
    for fi23 := 0; fi23 < 3; fi23++ {
        for fi24 := 0; fi24 < 14; fi24++ {
            if ((!true) || v2) {
                if true {
                    v5 = (v5 % ((42) | 1))
                    if (!false) {
                        if true {
                            var v25 bool = (0 >= 65535)
                            v26 := GenS1{F0: (v3 >> ((2) & 63))}
                            var v27 uint64 = 42
                            var v28 uint16 = v7
                            v29 := GenS2{F0: fi23}
                            v30 := GenS1{F0: v3}
                            _ = v25
                            _ = v26
                            _ = v27
                            _ = v28
                            _ = v29
                            _ = v30
                        } else {
                            switch fi23 {
                            case 1:
                                _ = len(v11)
                            default:
                                var v31 int32 = (v4 << ((315) & 63))
                                _ = v31
                            }
                        }
                    }
                    v4 -= (v4 - (v4 >> ((481) & 63)))
                }
                var v32 uint8 = 0
                _ = v32
            } else {
                v7 -= (v7 << ((0) & 63))
                switch "foo bar" {
                case "\n\t":
                    if true {
                        v33 := GenS2{F0: fi23}
                        _ = v33
                    } else {
                        switch (("\n\t" + "foo bar") + "") {
                        case "a":
                            v3 = v3
                        case "hello":
                            var v34 float32 = float32(1e10)
                            _ = v34
                        default:
                            var v35 bool = (v4 <= (v4 | v4))
                            _ = v35
                        }
                    }
                default:
                    v2 = v2
                }
                var v36 int8 = (v5 - -128)
                _ = v36
            }
            _ = fi24
        }
        var v37 uint8 = 42
        fmt.Println(v4, v5)
        _ = fi23
        _ = v37
    }
    var v38 int = 32767
    v39 := GenS1{F0: -9223372036854775808}
    if (-2147483648 != (v4 ^ -925)) {
        v40 := GenS2{F0: (v38 | int(v5))}
        _ = v40
    } else {
        v11 = append(v11, (-(v38)))
        v1[((("hello" + ("\n\t" + "\n\t")) + ("" + (("\n\t" + "\n\t") + ("\n\t" + "a")))) + "\n\t")] = (v38 * v38)
    }
    if ("a" >= "a") {
        switch int(v4) {
        case 127:
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
        case -856:
            var v41 float32 = -2.25
            _ = v41
        }
    }
    v42 := v1["a"]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v7
}
