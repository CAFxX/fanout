package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }

func main() {
    var v1 int = 392
    v2 := make(map[string]string)
    var v3 int64 = 9223372036854775807
    for fi4 := 0; fi4 < 8; fi4++ {
        for fi5 := 0; fi5 < 10; fi5++ {
            delete(v2, "foo bar")
            v1 += v1
            _ = fi5
        }
        _ = len(v2)
        v6 := []int{int(v3), -1}
        _ = fi4
        _ = v6
    }
    v1 -= v1
    for fi7 := 0; fi7 < 6; fi7++ {
        fmt.Println(v1, v1, v3)
        var v8 uint8 = 1
        switch 9223372036854775807 {
        case -941:
            if false {
                fmt.Println(fi7, v3, v1)
            } else {
                var v9 uint16 = 939
                _ = v9
            }
        case 9223372036854775807:
            switch ("a" + "") {
            case "":
                switch (v1 | (fi7 | (fi7 & 32768))) {
                case -1:
                    var v10 uint64 = 0
                    _ = v10
                default:
                    switch ((("\n\t" + "a") + "a") + ("a" + "")) {
                    case "foo bar":
                        var v11 int = 127
                        _ = v11
                    case "hello":
                        var v12 int16 = 1
                        _ = v12
                    case "":
                        var v13 int = (v1 | 0)
                        _ = v13
                    default:
                        for fi14 := 0; fi14 < 14; fi14++ {
                            v3 -= (-(v3))
                            var v15 int8 = -1
                            _ = fi14
                            _ = v15
                        }
                    }
                }
            case "\n\t":
                switch "\n\t" {
                case "":
                    fmt.Println(fi7)
                case "foo bar":
                    var v16 int = (fi7 << ((1) & 63))
                    _ = v16
                default:
                    var v17 int32 = 0
                    _ = v17
                }
            case "foo bar":
                var v18 float64 = (3.14159 - 0.0)
                _ = v18
            }
        case -2147483648:
            var v19 uint16 = 255
            _ = v19
        }
        _ = fi7
        _ = v8
    }
    var v20 int32 = -1
    var v21 uint64 = 1
    v3 = 1
    switch (v1 >> ((255) & 63)) {
    case -980:
        v22 := []int{(v1 + (v1 | (-(v1))))}
        _ = v22
    default:
        v3 += (^(v3))
    }
    fmt.Println(v3)
    var v23 string = ("\n\t" + "foo bar")
    v23 = (((v23 + "hello") + v23) + (v23 + "foo bar"))
    switch (v1 * int(v3)) {
    case 0:
        v2[v23] = ((v23 + "a") + "a")
    case -162:
        for fi24 := 0; fi24 < 6; fi24++ {
            var v25 uint8 = 119
            v26 := GenS1{F0: (v20 >> ((255) & 63)), F1: 427}
            v25 = 0
            _ = fi24
            _ = v25
            _ = v26
        }
    case 32767:
        v27 := []int{v1, (v1 | v1), 77, (v1 & (-(v1)))}
        _ = v27
    }
    v23 = v23
    for fi28 := 0; fi28 < 13; fi28++ {
        v23 = ("foo bar" + "hello")
        fmt.Println(v23, v1, v23)
        _ = fi28
    }
    if false {
        for fi29 := 0; fi29 < 10; fi29++ {
            var v30 bool = false
            _ = fi29
            _ = v30
        }
    } else {
        v23 = (v23 + "\n\t")
        var v31 uint16 = 1
        v21 -= (v21 + (v21 >> ((211) & 63)))
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
}
