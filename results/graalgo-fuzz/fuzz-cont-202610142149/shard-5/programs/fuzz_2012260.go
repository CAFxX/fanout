package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8 }

func main() {
    v1 := []int{}
    var v2 uint64 = 18446744073709551615
    v3 := make(map[string]string)
    var v4 uint16 = 2
    v5 := make(map[uint32]bool)
    var v6 int16 = -1
    switch ("" + "\n\t") {
    case "":
        for fi7 := 0; fi7 < 10; fi7++ {
            if true {
                v8 := GenS1{F0: (v6 & (^(v6)))}
                _ = v8
            }
            _ = fi7
        }
    case "\n\t":
        switch 128 {
        case 2147483647:
            v6 = (v6 / ((-1) | 1))
        case 128:
            var v9 uint8 = 1
            _ = v9
        default:
            var v10 int16 = (^(v6))
            _ = v10
        }
    case "a":
        if len(v1) > 0 {
            v1[0] = 9223372036854775807
        }
    default:
        switch 845 {
        case 4294967295:
            var v11 int64 = -1
            _ = v11
        case 616:
            var v12 int64 = -9223372036854775808
            _ = v12
        default:
            if false {
                v13 := GenS2{F0: -35}
                _ = v13
            }
        }
    }
    fmt.Println(v2, v4, v2)
    for fi14 := 0; fi14 < 12; fi14++ {
        for fi15 := 0; fi15 < 2; fi15++ {
            var v16 float64 = (3.14159 * ((float64(0.0) - float64(3.14159)) * ((float64((-2.25 + 0.1)) + 1e10) + (0.1 + 3.14159))))
            _ = fi15
            _ = v16
        }
        for fi17 := 0; fi17 < 5; fi17++ {
            var v18 bool = ((v6 - (^(v6))) == v6)
            v19 := GenS2{F0: 127}
            v18 = (false || (-116 != 2147483647))
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi14
    }
    if len(v1) > 0 {
        v1[0] = -9223372036854775808
    }
    fmt.Println(v6, v4, v6)
    if (!false) {
        if false {
            var v20 uint = 749
            _ = v20
        }
        for fi21 := 0; fi21 < 9; fi21++ {
            var v22 bool = ((false || (!(125 >= -128))) && ((true || false) || true))
            var v23 int64 = 9223372036854775807
            _ = len(v3)
            _ = fi21
            _ = v22
            _ = v23
        }
    }
    var v24 int8 = 100
    var v25 uint = 1
    if (int64(v25) < -9223372036854775808) {
        fmt.Println(v6, v2)
        for fi26 := 0; fi26 < 3; fi26++ {
            v5[112] = false
            if (42 >= -2147483648) {
                for fi27 := 0; fi27 < 20; fi27++ {
                    if false {
                        v28 := GenS2{F0: (v24 << ((v25) & 63))}
                        switch (((("a" + "") + "") + ("hello" + "\n\t")) + (("\n\t" + "a") + "hello")) {
                        case "\n\t":
                            var v29 int32 = -2147483648
                            _ = v29
                        case "foo bar":
                            fmt.Println(v24)
                        case "a":
                            if false {
                                v30 := v5[42]
                                v4 += v4
                                v30 = v30
                                _ = v30
                            } else {
                                v6 = -1
                                var v31 uint8 = 127
                                _ = v31
                            }
                        default:
                            var v32 uint = v25
                            _ = v32
                        }
                        switch (^(fi27)) {
                        case 848:
                            var v33 uint32 = uint32(v25)
                            _ = v33
                        case 9223372036854775807:
                            if false {
                                v34 := v3[""]
                                _ = v34
                            }
                        case -9223372036854775808:
                            v35 := v3["a"]
                            _ = v35
                        }
                        v24 = 76
                        _ = v28
                    }
                    if true {
                        v4 = (v4 ^ (v4 >> ((659) & 63)))
                    } else {
                        v25 -= 18446744073709551615
                        var v36 int8 = v24
                        _ = v36
                    }
                    switch ((("\n\t" + "a") + "hello") + "\n\t") {
                    case "\n\t":
                        if (1 < (v2 + v2)) {
                            var v37 uint16 = 1
                            _ = v37
                        }
                    case "hello":
                        var v38 uint = 255
                        _ = v38
                    case "foo bar":
                        var v39 bool = false
                        _ = v39
                    }
                    _ = fi27
                }
            }
            _ = fi26
        }
        v40 := GenS2{F0: (v24 & v24)}
        _ = v40
    } else {
        if true {
            v4 -= 0
        }
        var v41 float64 = (((-2.25 + ((0.0 + (0.0 + 1e10)) * 3.14159)) - 1.5) - (float64(3.14159) - 0.1))
        _ = v41
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
