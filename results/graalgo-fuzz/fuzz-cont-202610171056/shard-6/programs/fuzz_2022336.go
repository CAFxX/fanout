package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 uint64 = 63
    v2 := make(map[string]int)
    var v3 uint64 = v1
    v4 := make(map[string]string)
    for fi5 := 0; fi5 < 13; fi5++ {
        v6 := []int{fi5}
        switch 9223372036854775807 {
        case -727:
            v3 = v3
        case 9223372036854775807:
            v3 += 4294967295
        case -351:
            switch 127 {
            case -1:
                fmt.Println(v3, v1, v1)
            case 148:
                var v7 string = "a"
                _ = v7
            case 651:
                switch (("foo bar" + "\n\t") + "") {
                case "":
                    for fi8 := 0; fi8 < 10; fi8++ {
                        v9 := GenS1{F0: (2 < uint16(v3))}
                        var v10 float64 = 0.1
                        var v11 uint64 = (v3 ^ (v1 << ((65535) & 63)))
                        _ = fi8
                        _ = v10
                        _ = v11
                        _ = v9
                    }
                case "\n\t":
                    fmt.Println(v3, v3, v3)
                case "a":
                    v12 := v4[("foo bar" + "\n\t")]
                    _ = v12
                default:
                    for fi13 := 0; fi13 < 1; fi13++ {
                        var v14 float64 = float64((((((-2.25 + 1e10) * 1.5) * (1.5 / 0.1)) * -2.25) / 0.1))
                        _ = fi13
                        _ = v14
                    }
                }
            default:
                delete(v2, "a")
            }
        default:
            v1 = v1
        }
        _ = fi5
        _ = v6
    }
    v3 += 4294967295
    v3 = (v3 ^ 4294967295)
    var v15 uint16 = 255
    if false {
        switch "" {
        case "\n\t":
            switch (("foo bar" + "hello") + "hello") {
            case "hello":
                v16 := GenS2{F0: v1}
                _ = v16
            case "":
                var v17 float64 = (1e10 + (float64(float32((-2.25 / float64(-2.25)))) * (float64((float32(1.5) + 1.5)) * -2.25)))
                _ = v17
            }
        case "":
            fmt.Println(v1, v3, v15)
        default:
            fmt.Println(v3)
        }
        for fi18 := 0; fi18 < 2; fi18++ {
            fmt.Println(v3, v3, v1)
            _ = fi18
        }
        for fi19 := 0; fi19 < 1; fi19++ {
            v1 = v1
            _ = fi19
        }
    } else {
        var v20 uint8 = uint8(v15)
        _ = v20
    }
    v3 = 662
    if true {
        v4["foo bar"] = (("foo bar" + "hello") + ("hello" + ((("a" + "a") + (("\n\t" + ("a" + "\n\t")) + "")) + "")))
    } else {
        if false {
            var v21 uint16 = (v15 + (v15 >> ((2) & 63)))
            v22 := GenS2{F0: v1}
            _ = v21
            _ = v22
        } else {
            var v23 uint8 = 2
            _ = v23
        }
        if (!true) {
            v24 := GenS2{F0: (v1 + v1)}
            _ = v24
        }
    }
    for fi25 := 0; fi25 < 10; fi25++ {
        v26 := []int{int(v3), fi25, (fi25 + fi25), fi25}
        _ = len(v26)
        v27 := v4["foo bar"]
        _ = fi25
        _ = v26
        _ = v27
    }
    v3 += 18446744073709551615
    if ((!(89 != 127)) || (v15 <= v15)) {
        var v28 bool = (((v1 | (^(v3))) <= v3) || (!true))
        if v28 {
            var v29 int32 = -2147483648
            var v30 int64 = 876
            _ = v29
            _ = v30
        } else {
            if v28 {
                fmt.Println(v15, v15)
            } else {
                var v31 uint32 = uint32(v15)
                v32 := []int{-2147483648, 127, 541}
                _ = v31
                _ = v32
            }
            v1 = v3
        }
        var v33 uint64 = v1
        _ = v28
        _ = v33
    } else {
        v3 += v3
    }
    v15 += 466
    for fi34 := 0; fi34 < 11; fi34++ {
        switch "a" {
        case "foo bar":
            for fi35 := 0; fi35 < 3; fi35++ {
                _ = len(v2)
                switch fi34 {
                case 219:
                    v3 = 705
                case 0:
                    switch ("foo bar" + ("\n\t" + "a")) {
                    case "a":
                        var v36 int32 = 1
                        _ = v36
                    case "hello":
                        v37 := []int{fi35, (fi34 % (((^(fi34))) | 1)), fi34}
                        _ = v37
                    case "":
                        var v38 int8 = -128
                        _ = v38
                    }
                }
                _ = fi35
            }
        case "\n\t":
            _ = len(v2)
        case "a":
            var v39 uint8 = uint8(v15)
            _ = v39
        }
        for fi40 := 0; fi40 < 18; fi40++ {
            if true {
                switch -9223372036854775808 {
                case 42:
                    var v41 uint8 = 42
                    _ = v41
                default:
                    var v42 int8 = -128
                    _ = v42
                }
                v43 := []int{(^(fi40)), (-(fi40))}
                _ = v43
            }
            v15 += (v15 + 946)
            _ = fi40
        }
        _ = fi34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
