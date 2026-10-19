package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string; F1 int8 }
func genFn0() uint16 {
    var v7 uint32 = 2
    var v8 bool = (972 < 9223372036854775807)
    _ = v7
    _ = v8
    return 32767
}

func main() {
    var v1 uint16 = 255
    v2 := make(map[string]uint32)
    var v3 bool = (42 <= 1)
    var v4 uint = 663
    var v5 uint64 = 934
    v6 := []int{}
    fmt.Println(v4, v5)
    v9 := GenS1{F0: "\n\t"}
    if v3 {
        v10 := GenS1{F0: ("\n\t" + "\n\t")}
        if len(v6) > 0 {
            v6[0] = -445
        }
        _ = v10
    }
    _ = len(v6)
    switch "a" {
    case "foo bar":
        v2["hello"] = 14
    case "\n\t":
        var v11 int32 = -1
        _ = v11
    }
    var v12 bool = (42 == uint32(v4))
    if v3 {
        var v13 int = 54
        switch ("" + "") {
        case "hello":
            for fi14 := 0; fi14 < 10; fi14++ {
                v15 := GenS1{F0: "\n\t"}
                _ = fi14
                _ = v15
            }
        case "a":
            fmt.Println(v12, v13)
        case "\n\t":
            v12 = v3
        default:
            switch (v13 >> ((v4) & 63)) {
            case -266:
                var v16 uint32 = 2147483647
                _ = v16
            case 382:
                if ((2 <= 0) || (0 < -2147483648)) {
                    v12 = ((!v3) && (548 == 188))
                    v2["a"] = 207
                }
            case 396:
                if v12 {
                    if (!((127 == 1) && (!(!v12)))) {
                        v13 += -9223372036854775808
                        for fi17 := 0; fi17 < 11; fi17++ {
                            v18 := v2["foo bar"]
                            var v19 uint16 = (^(v1))
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                    if len(v6) > 0 {
                        v6[0] = (^(v13))
                    }
                    fmt.Println(v13, v4)
                } else {
                    delete(v2, ((("" + "foo bar") + "") + ("foo bar" + ("\n\t" + "foo bar"))))
                }
            default:
                fmt.Println(v4, v13, v13)
            }
        }
        switch (v13 + v13) {
        case 229:
            var v20 float64 = (float64(3.14159) - (0.1 * 1.5))
            _ = v20
        case 619:
            var v21 float64 = 3.14159
            _ = v21
        default:
            if len(v6) > 0 {
                v6[0] = 32767
            }
        }
        _ = v13
    }
    switch (("\n\t" + (("a" + ("\n\t" + "")) + ("hello" + "hello"))) + "foo bar") {
    case "hello":
        _ = len(v6)
    case "\n\t":
        switch -9223372036854775808 {
        case -2147483648:
            var v22 int8 = -110
            _ = v22
        case -932:
            var v23 uint64 = (v5 + (v5 << ((v4) & 63)))
            _ = v23
        case 1:
            v5 = (v5 * (v5 << ((v4) & 63)))
        default:
            var v24 int = 9223372036854775807
            _ = v24
        }
    }
    v25 := GenS2{F0: -813, F1: (50 > 0)}
    v1 = v1
    v26 := GenS3{F0: -1}
    fmt.Println(v4, v12)
    for fi27 := 0; fi27 < 19; fi27++ {
        v1 += 32767
        if v3 {
            for fi28 := 0; fi28 < 10; fi28++ {
                switch (fi27 - (^(fi28))) {
                case 42:
                    v29 := GenS2{F0: 42, F1: (!((!false) || v12))}
                    _ = v29
                case -9223372036854775808:
                    var v30 float64 = 1.5
                    _ = v30
                }
                v3 = v3
                _ = fi28
            }
            var v31 uint = 736
            var v32 int32 = 2147483647
            _ = v31
            _ = v32
        } else {
            for fi33 := 0; fi33 < 5; fi33++ {
                v5 = (v5 << (((^(v4))) & 63))
                var v34 bool = ((0 == -128) && ((!(0.0 != 0.1)) || ((2147483648 >= 40) && (v5 > (v5 - 2)))))
                v1 += genFn0()
                v5 -= v5
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = fi33
                _ = v34
            }
            v1 = v1
            v12 = v3
        }
        _ = fi27
    }
    v35 := GenS4{F0: "foo bar", F1: -1}
    v12 = v12
    var v36 uint64 = (v5 - 18446744073709551615)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
