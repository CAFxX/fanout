package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 string }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]uint32)
    var v3 int8 = 42
    var v4 uint8 = 0
    v5 := []string{"hello"}
    var v6 bool = false
    if ((!(!((42 == 1) && v6))) && ("hello" == "\n\t")) {
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 int8 = v3
            _ = fi7
            _ = v8
        }
    } else {
        switch -573 {
        case 307:
            for fi9 := 0; fi9 < 4; fi9++ {
                v4 = (v4 << ((18446744073709551615) & 63))
                _ = fi9
            }
        case 128:
            var v10 bool = (-2.25 > (-2.25 / 1.0))
            _ = v10
        }
    }
    v3 = -128
    v2["a"] = 42
    v11 := GenS1{F0: 449}
    if (0 > 2147483647) {
        var v12 int16 = -32768
        for fi13 := 0; fi13 < 20; fi13++ {
            v12 -= v12
            _ = fi13
        }
        _ = v12
    } else {
        switch "foo bar" {
        case "\n\t":
            if len(v5) > 0 {
                v5[0] = "foo bar"
            }
        case "foo bar":
            if true {
                fmt.Println(v4, v6)
                var v14 bool = (v6 && v6)
                if v6 {
                    for fi15 := 0; fi15 < 11; fi15++ {
                        var v16 bool = v14
                        var v17 float64 = 3.14159
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                }
                _ = v14
            }
        case "a":
            switch 4294967295 {
            case -9223372036854775808:
                if (!v6) {
                    v18 := GenS1{F0: 1}
                    v6 = (!(255 == 2))
                    _ = v18
                }
            default:
                v19 := GenS2{F0: true, F1: 100}
                _ = v19
            }
        default:
            v20 := GenS1{F0: -440}
            _ = v20
        }
        v4 = v4
        v3 = v3
    }
    v6 = (0.0 > -2.25)
    var v21 int64 = 0
    fmt.Println(v3, v6)
    for fi22 := 0; fi22 < 4; fi22++ {
        v23 := v1["\n\t"]
        var v24 int16 = 19
        v23 = v23
        _ = fi22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 10; fi25++ {
        for fi26 := 0; fi26 < 8; fi26++ {
            v3 = (v3 % ((v3) | 1))
            v2["foo bar"] = 0
            v27 := GenS1{F0: -2147483648}
            v28 := v2["hello"]
            _ = fi26
            _ = v27
            _ = v28
        }
        v29 := GenS1{F0: 9223372036854775807}
        _ = fi25
        _ = v29
    }
    var v30 bool = ((!v6) && v6)
    if true {
        v5 = append(v5, ("foo bar" + "\n\t"))
    } else {
        switch -706 {
        case -694:
            if v30 {
                v30 = ((534 > 65535) || (4294967295 < 413))
            }
        case -221:
            var v31 uint32 = 1
            _ = v31
        case 196:
            fmt.Println(v21, v21)
        default:
            fmt.Println(v4, v3)
        }
        v4 = (v4 - v4)
        if true {
            v32 := GenS2{F0: (v21 >= (v21 >> ((811) & 63))), F1: v3}
            for fi33 := 0; fi33 < 6; fi33++ {
                if false {
                    var v34 int16 = 1000
                    v34 -= -32768
                    var v35 int32 = 597
                    _ = v34
                    _ = v35
                } else {
                    delete(v1, ("a" + "a"))
                    v5 = append(v5, "\n\t")
                }
                _ = len(v1)
                if (!true) {
                    if true {
                        v6 = v30
                        v36 := v2[("hello" + "hello")]
                        var v37 uint16 = 2
                        _ = v36
                        _ = v37
                    }
                } else {
                    v3 = (-(v3))
                    v38 := GenS3{F0: "\n\t"}
                    _ = v38
                }
                _ = fi33
            }
            _ = v32
        }
    }
    v39 := GenS1{F0: 128}
    fmt.Println(v30)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
