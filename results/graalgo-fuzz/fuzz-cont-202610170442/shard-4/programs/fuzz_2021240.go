package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 int = -569
    var v2 float32 = ((-2.25 / 1.0) - 1e10)
    v3 := []string{"\n\t", ("hello" + "foo bar")}
    var v4 string = ((("hello" + "foo bar") + "hello") + ("\n\t" + (("foo bar" + ("" + "a")) + "\n\t")))
    v5 := make(map[uint32]int)
    switch v4 {
    case "\n\t":
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
    default:
        v4 = ("hello" + v4)
    }
    v6 := GenS1{F0: 4294967295, F1: 128, F2: 1}
    v7 := v5[455]
    var v8 int32 = 0
    fmt.Println(v1)
    var v9 int64 = 0
    v1 = 674
    var v10 uint8 = 176
    var v11 int = v1
    fmt.Println(v9, v11)
    v8 = v8
    v3 = append(v3, ((v4 + (v4 + (v4 + ((((("\n\t" + v4) + v4) + "\n\t") + "a") + v4)))) + (v4 + ("\n\t" + (((("hello" + v4) + "") + "foo bar") + v4)))))
    if (!(!(v7 <= (v7 >> ((1) & 63))))) {
        for fi12 := 0; fi12 < 11; fi12++ {
            v8 += -767
            _ = fi12
        }
    }
    var v13 string = v4
    if true {
        switch v13 {
        case "foo bar":
            if (499 == 4294967295) {
                switch (^(v11)) {
                case 32767:
                    if false {
                        fmt.Println(v10)
                        var v14 uint64 = 1
                        var v15 uint32 = 263
                        _ = v14
                        _ = v15
                    }
                case -814:
                    var v16 uint16 = 355
                    _ = v16
                }
            } else {
                switch (-(v1)) {
                case 42:
                    var v17 uint16 = 42
                    _ = v17
                case 9223372036854775807:
                    v18 := GenS2{F0: -354}
                    _ = v18
                }
            }
        case "a":
            fmt.Println(v4)
        case "\n\t":
            if len(v3) > 0 {
                v3[0] = v13
            }
        default:
            v9 = 16
        }
        v19 := GenS1{F0: 0, F1: v10, F2: 32767}
        var v20 uint16 = 0
        _ = v19
        _ = v20
    }
    if true {
        v21 := v5[2]
        _ = v21
    }
    var v22 int64 = v9
    for fi23 := 0; fi23 < 3; fi23++ {
        switch v4 {
        case "":
            v11 = v1
        case "foo bar":
            switch (v11 * (-(v1))) {
            case 1:
                for fi24 := 0; fi24 < 16; fi24++ {
                    fmt.Println(v22)
                    if false {
                        var v25 uint64 = 0
                        fmt.Println(v8)
                        _ = v25
                    }
                    v8 = v8
                    _ = fi24
                }
            case 4294967295:
                for fi26 := 0; fi26 < 18; fi26++ {
                    var v27 int64 = -835
                    _ = fi26
                    _ = v27
                }
            case 916:
                var v28 int32 = 0
                _ = v28
            }
        }
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
