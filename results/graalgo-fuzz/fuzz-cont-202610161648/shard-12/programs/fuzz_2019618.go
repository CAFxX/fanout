package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 uint32; F2 int }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 uint16 = 42
    v2 := make(map[int]uint32)
    v3 := []string{"hello", "\n\t", "a", "hello"}
    var v4 uint16 = (v1 - (v1 * (v1 / (((^(v1))) | 1))))
    var v5 bool = true
    v6 := v2[-982]
    switch -2147483648 {
    case -1:
        fmt.Println(v5, v1)
    case 211:
        if (((!true) || (722 <= 255)) && (!(40 == -37))) {
            var v7 int64 = -9223372036854775808
            _ = v7
        } else {
            var v8 int64 = 42
            var v9 int = -2147483648
            _ = len(v3)
            _ = v8
            _ = v9
        }
    case -2147483648:
        v10 := GenS1{F0: (("a" + (("\n\t" + "a") + ("\n\t" + "foo bar"))) + ("\n\t" + "\n\t"))}
        _ = v10
    }
    var v11 int32 = -973
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 uint32 = 2147483647
        _ = fi12
        _ = v13
    }
    v2[4294967295] = v6
    fmt.Println(v11, v6, v4)
    if (!(!v5)) {
        switch 32768 {
        case 2147483647:
            v14 := GenS1{F0: ("\n\t" + "")}
            _ = v14
        }
        switch ("\n\t" + ("" + "\n\t")) {
        case "\n\t":
            switch "hello" {
            case "\n\t":
                for fi15 := 0; fi15 < 12; fi15++ {
                    v16 := GenS2{F0: fi15, F1: (v6 | uint32(v1)), F2: (fi15 / (((-(fi15))) | 1))}
                    _ = fi15
                    _ = v16
                }
            case "foo bar":
                for fi17 := 0; fi17 < 12; fi17++ {
                    v3 = append(v3, "hello")
                    switch fi17 {
                    case -9223372036854775808:
                        if v5 {
                            v1 = 65535
                        }
                    case 9223372036854775807:
                        var v18 int32 = (v11 - (v11 >> ((42) & 63)))
                        _ = v18
                    case 2147483647:
                        fmt.Println(v6)
                    }
                    fmt.Println(v1, v4, v4)
                    _ = fi17
                }
            case "a":
                v1 = (v1 | 2)
            default:
                v1 += v1
            }
        }
    } else {
        var v19 string = "a"
        _ = v19
    }
    v3 = append(v3, "a")
    for fi20 := 0; fi20 < 5; fi20++ {
        v21 := GenS1{F0: "a"}
        switch ("foo bar" + "a") {
        case "a":
            var v22 float64 = 3.14159
            _ = v22
        case "hello":
            for fi23 := 0; fi23 < 20; fi23++ {
                var v24 uint32 = v6
                _ = fi23
                _ = v24
            }
        case "foo bar":
            v1 += (v1 % ((v1) | 1))
        }
        _ = fi20
        _ = v21
    }
    v11 -= 42
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    _ = len(v2)
    var v25 int64 = -1
    for fi26 := 0; fi26 < 3; fi26++ {
        v27 := GenS3{F0: (v1 & (v4 % ((975) | 1)))}
        _ = fi26
        _ = v27
    }
    if true {
        v2[167] = (v6 / ((v6) | 1))
        for fi28 := 0; fi28 < 19; fi28++ {
            if ((45 == 42) && v5) {
                switch "" {
                case "":
                    v25 -= v25
                }
                if v5 {
                    var v29 uint8 = 2
                    v11 += (v11 ^ v11)
                    _ = v29
                }
            }
            _ = fi28
        }
    }
    var v30 uint32 = 2
    switch "\n\t" {
    case "\n\t":
        _ = len(v3)
    case "hello":
        for fi31 := 0; fi31 < 14; fi31++ {
            v6 = 65535
            _ = fi31
        }
    default:
        for fi32 := 0; fi32 < 3; fi32++ {
            v11 = v11
            _ = fi32
        }
    }
    fmt.Println(v25)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
