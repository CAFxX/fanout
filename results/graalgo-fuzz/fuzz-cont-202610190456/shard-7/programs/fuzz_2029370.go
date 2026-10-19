package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 4294967295
    var v3 float32 = ((0.1 - 0.0) * -2.25)
    v4 := []uint32{2147483648, 1, 0}
    v5 := []string{"hello", "a"}
    var v6 bool = true
    fmt.Println(v2)
    if v6 {
        switch -15 {
        case -1:
            if v6 {
                var v7 uint64 = 4294967295
                _ = v7
            } else {
                v8 := GenS1{F0: 335}
                var v9 uint = (v2 * v2)
                switch ("hello" + (("\n\t" + (("\n\t" + "foo bar") + "\n\t")) + ("foo bar" + "foo bar"))) {
                case "":
                    var v10 int32 = -2147483648
                    _ = v10
                }
                _ = v8
                _ = v9
            }
        default:
            for fi11 := 0; fi11 < 13; fi11++ {
                switch fi11 {
                case -1:
                    var v12 int64 = -487
                    _ = v12
                case 127:
                    v13 := GenS2{F0: 86}
                    _ = v13
                }
                _ = fi11
            }
        }
        for fi14 := 0; fi14 < 9; fi14++ {
            v3 = v3
            v3 = 1.5
            _ = fi14
        }
    } else {
        switch -9223372036854775808 {
        case 127:
            if ("foo bar" < (("\n\t" + (("" + "foo bar") + "hello")) + ("hello" + ((("" + "a") + "hello") + ("\n\t" + "\n\t"))))) {
                for fi15 := 0; fi15 < 16; fi15++ {
                    var v16 string = ("\n\t" + "\n\t")
                    _ = fi15
                    _ = v16
                }
            }
        case 32768:
            v6 = (v6 && true)
        default:
            if len(v5) > 0 {
                v5[0] = "\n\t"
            }
        }
        if (false && false) {
            _ = len(v4)
        }
    }
    v1["\n\t"] = ((^(v2)) != v2)
    v17 := GenS1{F0: 0}
    for fi18 := 0; fi18 < 11; fi18++ {
        v3 = ((v3 - ((-2.25 - float32(3.14159)) - 0.0)) * 1.5)
        delete(v1, "a")
        var v19 bool = (-2.25 != v3)
        _ = fi18
        _ = v19
    }
    fmt.Println(v6, v2)
    v20 := GenS1{F0: 2147483647}
    var v21 float32 = v3
    v22 := v1["a"]
    var v23 uint = (v2 | v2)
    for fi24 := 0; fi24 < 3; fi24++ {
        v21 = (0.1 * (v21 * v21))
        _ = fi24
    }
    switch "a" {
    case "foo bar":
        v3 += ((v21 * v3) / ((v3 * 1e10) - v21))
    case "":
        for fi25 := 0; fi25 < 13; fi25++ {
            var v26 float64 = (float64(0.0) / (float64(v3) * 0.0))
            fmt.Println(v3)
            _ = fi25
            _ = v26
        }
    }
    for fi27 := 0; fi27 < 6; fi27++ {
        for fi28 := 0; fi28 < 12; fi28++ {
            var v29 uint64 = 18446744073709551615
            _ = fi28
            _ = v29
        }
        switch 32767 {
        case 42:
            switch "" {
            case "a":
                v22 = (fi27 < (fi27 * (fi27 & (fi27 ^ int(v2)))))
            case "hello":
                fmt.Println(v23)
            default:
                var v30 int64 = 42
                _ = v30
            }
        }
        var v31 float32 = 1.5
        _ = fi27
        _ = v31
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
