package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := []string{"\n\t", ("foo bar" + "foo bar")}
    var v2 string = (("a" + ("\n\t" + "a")) + "\n\t")
    v3 := make(map[uint32]uint32)
    var v4 uint16 = 2
    for fi5 := 0; fi5 < 14; fi5++ {
        switch (fi5 * fi5) {
        case -450:
            var v6 uint32 = 196
            _ = v6
        case -2147483648:
            var v7 int = fi5
            _ = v7
        case 9223372036854775807:
            v8 := GenS1{F0: 617}
            _ = v8
        }
        _ = fi5
    }
    v9 := v3[2]
    if true {
        if (!false) {
            switch v2 {
            case "hello":
                v1 = append(v1, "")
            default:
                if (-885 == -1) {
                    if (("\n\t" + v2) < "foo bar") {
                        for fi10 := 0; fi10 < 18; fi10++ {
                            var v11 int16 = 840
                            fmt.Println(fi10, v11, v4)
                            _ = fi10
                            _ = v11
                        }
                    }
                } else {
                    _ = len(v1)
                }
            }
            if false {
                var v12 uint8 = 255
                if (false && (1 < -532)) {
                    if true {
                        if true {
                            v13 := GenS1{F0: v9}
                            _ = v13
                        }
                    }
                }
                _ = v12
            }
            v2 = "hello"
        }
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 string = (v2 + ((v2 + "\n\t") + "foo bar"))
        fmt.Println(fi14)
        _ = fi14
        _ = v15
    }
    switch -2147483648 {
    case 572:
        _ = len(v3)
    case 0:
        var v16 int = -9223372036854775808
        _ = v16
    case -183:
        var v17 int64 = -1
        _ = v17
    }
    v2 = "\n\t"
    fmt.Println(v9, v4)
    if (uint8(v9) >= 1) {
        var v18 uint = 324
        _ = v18
    }
    if true {
        var v19 int = 32768
        if (v4 <= 576) {
            delete(v3, v9)
            switch v2 {
            case "foo bar":
                delete(v3, (v9 ^ 65535))
            case "":
                var v20 string = "a"
                _ = v20
            }
            v2 = v2
        }
        fmt.Println(v9)
        _ = v19
    } else {
        var v21 int = 128
        _ = v21
    }
    var v22 int64 = 147
    v9 = v9
    if true {
        var v23 bool = false
        v22 = v22
        _ = v23
    } else {
        if ((-2 < 510) && true) {
            v9 += (^(v9))
            fmt.Println(v9, v22, v9)
            if (!false) {
                var v24 uint8 = 1
                var v25 int = 819
                _ = v24
                _ = v25
            } else {
                switch 42 {
                case -2147483648:
                    var v26 string = v2
                    _ = v26
                }
            }
        } else {
            v9 = v9
            v22 = -9223372036854775808
        }
        fmt.Println(v22, v2)
        var v27 uint8 = 255
        _ = v27
    }
    for fi28 := 0; fi28 < 1; fi28++ {
        var v29 int64 = v22
        var v30 int32 = 42
        _ = fi28
        _ = v29
        _ = v30
    }
    fmt.Println(v9, v4, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v9
}
